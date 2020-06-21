package AWS

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"repo.nefrosovet.ru/maximus-platform/recognition/logger"
	"sync"
)

type awsRegion string

const (
	defaultSimilarity = 90.000000
	personMetaKey     = "PersonID"

	ImagePNG  = "image/png"
	ImageJPEG = "image/jpeg"

	usEast1      awsRegion = "us-east-1"
	usEast2      awsRegion = "us-east-2"
	usWest1      awsRegion = "us-west-1"
	usWest2      awsRegion = "us-west-2"
	apSouth1     awsRegion = "ap-south-1"
	apNorthEast1 awsRegion = "ap-northeast-1"
	apNorthEast2 awsRegion = "ap-northeast-2"
	apSouthEast1 awsRegion = "ap-southeast-1"
	apSouthEast2 awsRegion = "ap-southeast-2"
	euCentral1   awsRegion = "eu-central-1"
	euWest1      awsRegion = "eu-west-1"
	euWest2      awsRegion = "eu-west-2"
)

var (
	allowedImageType = map[string]struct{}{
		ImageJPEG: {},
		ImagePNG:  {},
	}
	allowedRegions = map[awsRegion]struct{}{
		usEast1:      {},
		usEast2:      {},
		usWest1:      {},
		usWest2:      {},
		apSouth1:     {},
		apNorthEast1: {},
		apNorthEast2: {},
		apSouthEast1: {},
		apSouthEast2: {},
		euCentral1:   {},
		euWest1:      {},
		euWest2:      {},
	}
)

var (
	ErrUnsupportedImage  = errors.New("unsupported image format")
	ErrUnrekognized      = errors.New("rekognition: images not found")
	ErrObjectNotFound    = errors.New("s3 object not found")
	ErrBucketDoesntExist = errors.New("s3 bucket doesn't exist")
	ErrNotStored         = errors.New("image wasn't stored")
	ErrUnsupportedRegion = fmt.Errorf(
		"Allowed regions: %v. Given ",
		reflect.ValueOf(allowedRegions).MapKeys(),
	)
	ErrPartiallyNotCompared = errors.New("comparing service sends error")
)

func ValidateImageType(img []byte) (format string, ok bool) {
	format = http.DetectContentType(img)
	_, ok = allowedImageType[format]
	return format, ok
}

type awsClient struct {
	bucket     string
	similarity float64
	name       string

	storage     *s3.S3
	recognition *rekognition.Rekognition

	wg *sync.WaitGroup
}

func NewAWSClient(
	similarity float64,
	bucket, accessID, accessSecret, region string,
	l logger.CoreEntrier) (_ *awsClient, err error) {
	if _, ok := allowedRegions[awsRegion(region)]; !ok {
		logrus.Fatal(ErrUnsupportedRegion, region)
	}

	if similarity == 0 {
		similarity = defaultSimilarity
	}

	var awsLogLevel aws.LogLevelType = aws.LogOff
	if l.Level() == "debug" {
		awsLogLevel = aws.LogDebug
	}

	var awsSess *session.Session
	awsSess, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewCredentials(
			&credentials.StaticProvider{Value: credentials.Value{
				AccessKeyID:     accessID,
				SecretAccessKey: accessSecret,
			}},
		),
		Region:   aws.String(region), // example "us-west-2"
		LogLevel: aws.LogLevel(awsLogLevel),
	})
	if err != nil {
		return nil, err
	}

	var s3Sess *s3.S3
	s3Sess = s3.New(awsSess)

	_, err = s3Sess.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				l.Debug(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				l.Debug(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				return nil, aerr
			}
		} else {
			return nil, aerr
		}
	}

	var rekSess *rekognition.Rekognition
	rekSess = rekognition.New(awsSess)
	return &awsClient{
		storage:     s3Sess,
		recognition: rekSess,
		bucket:      bucket,
		similarity:  similarity,
		name:        aws.SDKName,
		wg:          new(sync.WaitGroup),
	}, nil
}

func (c *awsClient) Set(personID string, format string, body []byte) (_ *strfmt.UUID, err error) {
	var photoID = new(strfmt.UUID)
	err = photoID.Scan(uuid.Must(uuid.NewV1()).String())
	if err != nil {
		return nil, err
	}

	var out *s3.PutObjectOutput
	out, err = c.storage.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(c.bucket),
		Key:         aws.String(photoID.String()),
		Body:        aws.ReadSeekCloser(bytes.NewReader(body)),
		ContentType: aws.String(format),
		Metadata:    map[string]*string{personMetaKey: aws.String(personID)},
	})
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrNotStored
	}

	return photoID, nil
}

// TODO cache
func (c *awsClient) Get(photoID string) (personID string, err error) {
	var out *s3.GetObjectOutput
	out, err = c.storage.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(photoID),
	})
	if err != nil {
		if err.Error() == s3.ErrCodeNoSuchKey {
			return "", ErrObjectNotFound
		}
		return "", err
	}
	if out == nil {
		return "", ErrObjectNotFound
	}

	if out.Metadata[personMetaKey] != nil {
		return *out.Metadata[personMetaKey], nil
	}

	return "", nil
}

func (c *awsClient) Delete(photoID string) (err error) {
	_, err = c.storage.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(photoID),
	})
	if err != nil {
		if err.Error() == s3.ErrCodeNoSuchKey {
			return ErrObjectNotFound
		}
		return err
	}
	return nil
}

// TODO cache
func (c *awsClient) List(limit *int64) (keys []*strfmt.UUID, err error) {
	var objects *s3.ListObjectsOutput
	objects, err = c.storage.ListObjects(&s3.ListObjectsInput{
		Bucket:  aws.String(c.bucket),
		MaxKeys: limit,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				return nil, ErrBucketDoesntExist
			default:
				return nil, aerr
			}
		}
		return nil, err
	}
	if len(objects.Contents) == 0 {
		return nil, ErrObjectNotFound
	}

	keys = make([]*strfmt.UUID, 0)
	for _, obj := range objects.Contents {
		if obj.Key != nil {
			key := strfmt.UUID(*obj.Key)
			keys = append(keys, &key)
		}
	}
	return keys, nil
}

type FaceMatch struct {
	Key        *strfmt.UUID
	Similarity *float64
}

type MatchedFaces []FaceMatch

func (c *awsClient) Rekognize(sourceImage []byte) (_ MatchedFaces, err error) {
	var keys []*strfmt.UUID
	keys, err = c.List(nil)
	if err != nil {
		return nil, err
	}

	var (
		rekErr error
		once = new(sync.Once)
		matchedChan  = make(chan FaceMatch, (len(keys) * 2)) // just in case
		matchedFaces = make(MatchedFaces, len(matchedChan))
	)

	for _, key := range keys {
		c.wg.Add(1)

		go func(k string) {
			defer c.wg.Done()
			input := &rekognition.CompareFacesInput{
				SimilarityThreshold: aws.Float64(c.similarity),
				SourceImage: &rekognition.Image{
					Bytes: sourceImage,
				},
				TargetImage: &rekognition.Image{
					S3Object: &rekognition.S3Object{
						Bucket: aws.String(c.bucket),
						Name:   aws.String(k),
					},
				},
			}

			var compareOut *rekognition.CompareFacesOutput
			compareOut, err = c.recognition.CompareFaces(input)
			if err != nil {
				once.Do(func() {
					rekErr = err // to be aware that at least one rekognition was failed and what error was
				})
				return
			}
			if len(compareOut.FaceMatches) == 0 {
				return
			}

			var mf *rekognition.CompareFacesMatch
			for _, mf = range compareOut.FaceMatches {
				if mf.Similarity == nil {
					continue
				}
				if *mf.Similarity >= c.similarity {
					matchedChan <- FaceMatch{
						Key:        key,
						Similarity: mf.Similarity,
					}
					return // for now we process only one person on a photo
				}
			}
		}(key.String())

	}
	c.wg.Wait()
	close(matchedChan)

	for k := range matchedChan {
		matchedFaces = append(matchedFaces, k)
	}

	if rekErr != nil {
		return matchedFaces, errors.Wrap(rekErr, ErrPartiallyNotCompared.Error())
	}

	if len(matchedFaces) == 0 {
		return nil, ErrUnrekognized
	}
	return matchedFaces, nil
}

func (c *awsClient) ServiceName() string {
	return c.name
}
