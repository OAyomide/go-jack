package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-jack/parser"
	"github.com/messenger"
)

var verifyT, accessT, appS = parser.GetTokens()

var (
	verifyToken = flag.String("verify-token", verifyT, "The token used to verify facebook (required)")
	verify      = flag.Bool("should-verify", false, "Whether or not the app should verify itself")
	pageToken   = flag.String("page-token", accessT, "The token that is used to verify the page on facebook")
	appSecret   = flag.String("app-secret", appS, "The app secret from the facebook developer portal (required)")
	host        = flag.String("host", "localhost", "The host used to serve the messenger bot")
	port        = flag.Int("port", 4000, "The port used to serve the messenger bot")
)

var bot = messenger.New(messenger.Options{
	Verify:      *verify,
	AppSecret:   *appSecret,
	VerifyToken: *verifyToken,
	Token:       *pageToken,
	WebhookURL:  "/api/webhook",
})

func main() {
	flag.Parse()

	if *verifyToken == "" || *appSecret == "" || *pageToken == "" {
		fmt.Println("There seem to be fields that are empty. Please make sure none is empty")
		fmt.Println()
		flag.Usage()

		os.Exit(-1)
	}

	// we want to see if the parser works and returns the persistent menu
	rs := parser.GetPersistentMenu()

	fmt.Println("HERE IS THE PERSISTENT MENU", rs)
	// Setup a handler to be triggered when a message is received
	bot.HandleMessage(func(m messenger.Message, r *messenger.Response) {
		fmt.Printf("%v (Sent, %v)\n", m.Text, m.Time.Format(time.UnixDate))

		p, err := bot.ProfileByID(m.Sender.ID, []string{"name", "first_name", "last_name", "profile_pic"})
		if err != nil {
			fmt.Println("Something went wrong!", err)
		}

		r.Text(fmt.Sprintf("Hello, %v!", p.FirstName), messenger.ResponseType)
	})

	// Setup a handler to be triggered when a message is delivered
	bot.HandleDelivery(func(d messenger.Delivery, r *messenger.Response) {
		fmt.Println("Delivered at:", d.Watermark().Format(time.UnixDate))
	})

	// Setup a handler to be triggered when a message is read
	bot.HandleRead(func(m messenger.Read, r *messenger.Response) {
		fmt.Println("Read at:", m.Watermark().Format(time.UnixDate))
	})

	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Println("Serving messenger bot on", addr)
	log.Fatal(http.ListenAndServe(addr, bot.Handler()))
}
