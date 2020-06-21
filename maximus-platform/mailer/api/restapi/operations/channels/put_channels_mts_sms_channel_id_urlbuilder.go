// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PutChannelsMtsSmsChannelIDURL generates an URL for the put channels mts sms channel ID operation
type PutChannelsMtsSmsChannelIDURL struct {
	ChannelID string

	AccessToken string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutChannelsMtsSmsChannelIDURL) WithBasePath(bp string) *PutChannelsMtsSmsChannelIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutChannelsMtsSmsChannelIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PutChannelsMtsSmsChannelIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/channels/mts_sms/{channelID}"

	channelID := o.ChannelID
	if channelID != "" {
		_path = strings.Replace(_path, "{channelID}", channelID, -1)
	} else {
		return nil, errors.New("channelId is required on PutChannelsMtsSmsChannelIDURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	accessTokenQ := o.AccessToken
	if accessTokenQ != "" {
		qs.Set("accessToken", accessTokenQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PutChannelsMtsSmsChannelIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PutChannelsMtsSmsChannelIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PutChannelsMtsSmsChannelIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PutChannelsMtsSmsChannelIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PutChannelsMtsSmsChannelIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PutChannelsMtsSmsChannelIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
