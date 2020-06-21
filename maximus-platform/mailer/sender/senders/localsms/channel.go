package localsms

import (
    "database/sql"
    "fmt"
    "strconv"

    // mysql methods
    _ "github.com/go-sql-driver/mysql"
    "github.com/tsenart/nap"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders"
    "repo.nefrosovet.ru/maximus-platform/mailer/sender/senders/localsmsdb"
)

var localSMSConnects = map[string]*nap.DB{}

// Channel channel struct
type Channel struct {
    ID       string
    Login    string
    Password string
    Port     int64
    Server   string
    ModemID  int64
    Db       string
}

// message is message struct
type message struct {
    DestinationNumber string
    TextDecoded       string
    // SenderID          string
    DeliveryReport bool
    CreatorID      string
    Multipart      bool
    Coding         string
}

// coding is coding enum
var coding = map[string]localsmsdb.Coding{
    "Unicode_No_Compression": localsmsdb.CodingUnicodeNoCompression,
    "8bit":                   localsmsdb.CodingBit,
    "Default_Compression":    localsmsdb.CodingDefaultCompression,
    "Unicode_Compression":    localsmsdb.CodingUnicodeCompression,
}

func (c *Channel) NormalizeDestination(dest string) (string, error) {
    return dest, nil
}

func (c *Channel) Send(destination, data string, opts ...senders.SendOption) error {
    m := message{
        DestinationNumber: destination,
        TextDecoded:       data,
    }
    for _, opt := range opts {
        if err := opt.Apply(&m); err != nil {
            return err
        }
    }
    return c.send(&m)
}

// Send method sends message
func (c *Channel) send(message *message) error {
    if localSMSConnects[c.ID] == nil {
        err := c.Connect()
        if err != nil {
            return err
        }
    } else {
        err := c.Connect()
        if err != nil {
            return err
        }
    }

    sqlStr := `
		INSERT into sms.outbox (
			TextDecoded,
			DestinationNumber,
			SenderID,
			coding,
			CreatorID
		) VALUES (?, ?, ?, ?, ?)
	`
    res, err := localSMSConnects[c.ID].Exec(
        sqlStr,
        message.TextDecoded,
        message.DestinationNumber,
        sql.NullString{String: strconv.Itoa(int(c.ModemID)), Valid: true},
        coding["Unicode_No_Compression"],
        message.CreatorID,
    )
    if err != nil {
        return err
    }
    // retrieve id
    id, err := res.LastInsertId()
    if err != nil {
        return err
    }

    _, err = localsmsdb.OutboxByID(localSMSConnects[c.ID], uint(id))

    return err
}

// Connect establishes connection to db
func (c *Channel) Connect() error {
    // Construct DBString
    // Example: login:pass@tcp(addr:port)/db_name
    dbstring := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.Login, c.Password, c.Server, c.Port, c.Db)

    var err error
    localSMSConnects[c.ID], err = nap.Open("mysql", dbstring)

    return err
}
