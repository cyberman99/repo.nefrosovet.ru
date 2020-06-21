package email

import (
    "crypto/tls"
    "fmt"
    "mime"
    "net/mail"
    "net/smtp"
    "net/textproto"
    "time"

    "github.com/asaskevich/govalidator"
    "github.com/jordan-wright/email"
    "github.com/pkg/errors"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
)

type Channel struct {
    ID          string
    From        string
    Login       string
    Password    string
    Port        int64
    Server      string
    SSL         bool
    ContentType string
}

func (ch *Channel) NormalizeDestination(destination string) (string, error) {
    if !govalidator.IsEmail(destination) {
        return "", senders.DestinationValidationError(senders.WrongDestinationMesssage)
    }
    return destination, nil
}

func (ch *Channel) Send(destination, data string, opts ...senders.SendOption) error {
    headers := make(textproto.MIMEHeader)
    headers.Set("Content-type", ch.ContentType)
    e := email.Email{
        From:    mime.QEncoding.Encode("utf-8", ch.From) + "<" + ch.Login + ">",
        To:      []string{destination},
        Text:    []byte(data),
        HTML:    []byte(data),
        Headers: headers,
    }
    for _, opt := range opts {
        if err := opt.Apply(&e); err != nil {
            return errors.Wrap(err, "unable to apply option to email")
        }
    }
    if ch.SSL {
        if err := ch.sendWithSSL(&e); err != nil {
            return errors.Wrap(err, "unable to send email with ssl")
        }
        return nil
    }
    pool, err := email.NewPool(
        fmt.Sprintf("%s:%d", ch.Server, ch.Port),
        1,
        smtp.PlainAuth("", ch.Login, ch.Password, ch.Server),
    )
    if err != nil {
        return errors.Wrap(err, "unable to create email pool")
    }
    if err := pool.Send(&e, 10*time.Second); err != nil {
        return errors.Wrap(err, "unable to send email with pool")
    }
    return nil

}

func (ch *Channel) sendWithSSL(e *email.Email) error {
    from := mail.Address{
        Name:    mime.QEncoding.Encode("utf-8", e.From),
        Address: ch.Login,
    }
    to := mail.Address{
        Name:    "",
        Address: e.To[0],
    }
    subj := mime.QEncoding.Encode("utf-8", e.Subject)
    body := e.HTML

    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subj
    headers["Content-type"] = e.Headers.Get("Content-type")

    // Setup message
    message := ""
    for k, v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + string(body)

    // Connect to the SMTP Server
    auth := smtp.PlainAuth("", ch.Login, ch.Password, ch.Server)

    // TLS config
    tlsconfig := &tls.Config{
        InsecureSkipVerify: true,
        ServerName:         ch.Server,
    }

    conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", ch.Server, ch.Port), tlsconfig)
    if err != nil {
        return errors.Wrap(err, "unable to dial tls")
    }

    c, err := smtp.NewClient(conn, ch.Server)
    if err != nil {
        return errors.Wrap(err, "unable to create smtp client")
    }

    // Auth
    if err = c.Auth(auth); err != nil {
        return errors.Wrap(err, "unable to auth")
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
        return err
    }

    if err = c.Rcpt(to.Address); err != nil {
        return err
    }

    // Data
    w, err := c.Data()
    if err != nil {
        return err
    }

    _, err = w.Write([]byte(message))
    if err != nil {
        return err
    }

    err = w.Close()
    if err != nil {
        return err
    }

    c.Quit()

    return nil

}
