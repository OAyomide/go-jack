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

func (mes *Messenger) separateEventTypes(ev MessagingStruct)
