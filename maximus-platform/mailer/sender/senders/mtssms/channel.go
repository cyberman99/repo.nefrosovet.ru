package mtssms

import (
    "crypto/md5"
    "encoding/hex"
    "errors"
    "fmt"
    "net/http"
    "net/url"

    log "github.com/Sirupsen/logrus"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
)

const (
    mtsHostname = "www.mcommunicator.ru"
)

type message struct {
    DestinationNumber string
    Data string
}

type Channel struct {
    ID                string
    Login             string
    Password          string
    From              string
}

func (c *Channel) NormalizeDestination(dest string) (string, error) {
    return dest, nil
}

func (c *Channel) Send(destination, data string, _ ...senders.SendOption) error {
    m := message{
        DestinationNumber: destination,
        Data:              data,
    }
    return c.send(&m)
}

func (c *Channel) send(m *message) error {
    payload := url.Values{}
    payload.Add("msid", m.DestinationNumber)
    payload.Add("message", m.Data)
    payload.Add("naming", c.From)
    payload.Add("login", c.Login)
    payload.Add("password", getMD5Hash(c.Password))

    req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/m2m/m2m_api.asmx/SendMessage?%s", mtsHostname, payload.Encode()), nil)
    if err != nil {
        log.Error(err)
        return err
    }

    httpClient := http.DefaultClient
    resp, err := httpClient.Do(req)
    if err != nil {
        log.Error(err)
        return err
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return errors.New("connection to mts communicator - error")
    }
    return nil
}

// getMD5Hash returns MD5 of string
func getMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
