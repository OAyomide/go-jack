package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-jack/parser"
	"github.com/go-jack/util"
)

type Messenger struct {
	mux           *http.ServeMux
	token         string
	verify        bool
	appsecret     string
	verifyhandler func(http.ResponseWriter, *http.Request)
	messagesHandlers []func(Messagestruct, *Response)
	postbackHandlers []func(Postbackstruct, *Response)
}

type Response struct {
	token string
	to Recipientstruct
}

var c parser.Config
var config = c.ReadYml()
var ErrH = util.HandleErr

//Options struct we're passing is used to declare a new
//properties for our bot
type Options struct {
	AppSecret   string
	VerifyToken string
	AccessToken string
	WebHookURL  string
	Mux         *http.ServeMux
	Verify      bool
}

//New creates a new bot
func New(mon Options) *Messenger {
	if mon.Mux == nil {
		mon.Mux = http.NewServeMux()
	}

	monty := &Messenger{
		mux:       mon.Mux,
		token:     mon.AccessToken,
		verify:    mon.Verify,
		appsecret: mon.AppSecret,
	}

	if mon.WebHookURL == "" {
		mon.WebHookURL = "/api/webhook"
	}

	monty.verifyhandler = util.WebhookVerify(c.AccessToken)
	monty.mux.HandleFunc(mon.WebHookURL)
}

//this function is being handled by the http Mux we declared in the Messenger struct
func (mes *Messenger) handleFunc(q http.ResponseWriter, t *http.Request) {
	if t.Method == "GET" {
		mes.verifyhandler(q, t)
		//return immediately
		return
	}

	var sc Body

	body, _ := ioutil.ReadAll(t.Body)

	err := json.Unmarshal(body, &sc)
	if err != nil {
		fmt.Println("Error marshalling the response::", err)
		fmt.Fprintf(q, "not parsed")
		//return immediately
		return
	}

	if sc.Object != "page" {
		fmt.Println("The webhook event we received isnt a page subscription", sc.Object)
	}

	//this handles the firing of all the necessary webhook event
	//handlers based on the webhook event we received
	//mes.dispatch(sc)
}

//check https://github.com/OAyomide/goblin/blob/ab4dbe668e4604f1e3b4dca453604ca13c1db747/webhook.go#L128
func (mes *Messenger) fireNecessaryEventHandlers(rec Body) {
	for _, entries := range rec.Entry {
		for _, messaging := range entries.Messaging {
			//here is where we are filter the type of event we're receiving
			eventType := mes.separateEventTypes(messaging, entries)

			if eventType == UnknownEventType {
				ErrH(messaging, "Oops!! Dont know that webhook event type")
				continue
			}

			response := &Response{
				token: mes.token,
				to: Recipientstruct{ID: messaging.Sender.ID}
			}

			switch eventType{
			case TextEvent:
				for _, handl := range mes.messagesHandlers{
					message := *messaging.Message
					message.Sender = messaging.Sender
					message.Recipient = messaging.Recipient
					message.Time = time.Unix(messaging.Timestamp/int64(time.Microsecond), 0)
					handl(message, response)
				}
			}
		}
	}
}

//SendResponse returns the response object
func (mes *Messenger) SendResponse(to int64) *Response {
	return &Response{
		to: Recipientstruct{ID: to},
		token: mes.token
	}
}

func (mes *Messenger) separateEventTypes(ev MessagingStruct, e Entry) Event {
	switch ev {
	case ev.Message != nil:
		return TextEvent
	case ev.Postback != nil:
		return PostBackEvent
	case ev.Read != nil:
		return ReadEvent
	case ev.Delivery != nil:
		return DeliveryEvent
	case ev.OptIn != nil:
		return OptInEvent
	case ev.AccountLinking != nil:
		return AccountLinkingEvent
	case ev.ReferralMessage != nil:
		return ReferralEvent
	default:
		return UnknownEventType
	}
}
