// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"

	"repo.nefrosovet.ru/go-lms/api-video/ent/webinaruser"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleAccountKey() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the accountkey's edges.

	// create accountkey vertex with its edges.
	ak := client.AccountKey.
		Create().
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetAccountID(1).
		SetKey("string").
		SetOptions("string").
		SetMetaData("string").
		SaveX(ctx)
	log.Println("accountkey created:", ak)

	// query edges.

	// Output:
}
func ExampleSubscriber() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the subscriber's edges.

	// create subscriber vertex with its edges.
	s := client.Subscriber.
		Create().
		SetUsername("string").
		SetDomain("string").
		SetHa1("string").
		SetHa1b("string").
		SaveX(ctx)
	log.Println("subscriber created:", s)

	// query edges.

	// Output:
}
func ExampleUser() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the user's edges.
	s0 := client.Subscriber.
		Create().
		SetUsername("string").
		SetDomain("string").
		SetHa1("string").
		SetHa1b("string").
		SaveX(ctx)
	log.Println("subscriber created:", s0)
	ua1 := client.UserAccount.
		Create().
		SetUsername("string").
		SetPassword("string").
		SetRememberToken("string").
		SetActive(1).
		SetEventChannel("string").
		SetDidPrefix("string").
		SetUseKamalio(1).
		SaveX(ctx)
	log.Println("useraccount created:", ua1)

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetMetaData("string").
		SetSubscriber(s0).
		SetUseraccount(ua1).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.
	s0, err = u.QuerySubscriber().First(ctx)
	if err != nil {
		log.Fatalf("failed querying subscriber: %v", err)
	}
	log.Println("subscriber found:", s0)

	ua1, err = u.QueryUseraccount().First(ctx)
	if err != nil {
		log.Fatalf("failed querying useraccount: %v", err)
	}
	log.Println("useraccount found:", ua1)

	// Output:
}
func ExampleUserAccount() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the useraccount's edges.
	ak1 := client.AccountKey.
		Create().
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetAccountID(1).
		SetKey("string").
		SetOptions("string").
		SetMetaData("string").
		SaveX(ctx)
	log.Println("accountkey created:", ak1)

	// create useraccount vertex with its edges.
	ua := client.UserAccount.
		Create().
		SetUsername("string").
		SetPassword("string").
		SetRememberToken("string").
		SetActive(1).
		SetEventChannel("string").
		SetDidPrefix("string").
		SetUseKamalio(1).
		AddAccountkeys(ak1).
		SaveX(ctx)
	log.Println("useraccount created:", ua)

	// query edges.

	ak1, err = ua.QueryAccountkeys().First(ctx)
	if err != nil {
		log.Fatalf("failed querying accountkeys: %v", err)
	}
	log.Println("accountkeys found:", ak1)

	// Output:
}
func ExampleWebinarUser() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the webinaruser's edges.

	// create webinaruser vertex with its edges.
	wu := client.WebinarUser.
		Create().
		SetUserID(1).
		SetWebinarID(1).
		SetStatus(webinaruser.StatusWAIT).
		SetMedoozeID(1).
		SetOldMedoozeID(1).
		SetMic(1).
		SetSound(1).
		SaveX(ctx)
	log.Println("webinaruser created:", wu)

	// query edges.

	// Output:
}