package bot

import "time"

//Postbackstruct is the struct for our postback event sent by the webhook
type Postbackstruct struct {
	Payload   string          `json:"payload"`
	Sender    Senderstruct    `json:"-"`
	Recipient Recipientstruct `json:"-"`
	Time      time.Time       `json:"-"`
	Referral  Referral        `json:"referral"`
}

//Referral struct is the struct for our referral webhook event object
type Referral struct {
	Ref    string `json:"ref"`
	Source string `json:"source"`
	Type   string `json:"type"`
}

//ReferralMessage is the struct for our referral message webhook event object
type ReferralMessage struct {
	Ref       *Referral
	Sender    Sender          `json:"-"`
	Recipient Recipientstruct `json:"-"`
	Time      time.Time       `json:"-"`
}

//AccountLinking is the struct for the Accountlinking webhook event
type AccountLinking struct {
	Sender            Senderstruct    `json:"-"`
	Recipient         Recipientstruct `json:"-"`
	Time              time.Time       `json:"-"`
	Status            string          `json:"status"`
	AuthorizationCode string          `json:"authorization_code"`
}

//Payloadstruct is the struct for our payload object
type Payloadstruct struct {
	URL         string
	Reusable    bool         `json:"is_reusable"`
	Coordinates *Coordinates `json:"coordinates,omitempty`
}

//Coordinates is sent when we have a location webhook event object (I think. Not sure how to put it)
type Coordinates struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

//Attachmentstruct is the struct for the attachment webhook event
type Attachmentstruct struct {
	Type    string
	Payload Payloadstruct `json:"-"`
	Time    time.Time     `json:"-"`
}

//Messagestruct is the struct for a facebook message
type Messagestruct struct {
	Text        string
	Attachments []*Attachmentstruct
}

//Recipientstruct is the struct for the Recipient webhook object
type Recipientstruct struct {
	ID int64 `json:"id,string"`
}

//Senderstruct is the struct for who sent the even
type Senderstruct struct {
	ID int64 `json:"id,string"`
}

//MessagingStruct is the struct for the messages sent
type MessagingStruct struct {
	Timestamp       int64            `json:"timestamp"`
	Sender          Senderstruct     `json:"sender"`
	Recipient       Recipientstruct  `json:"recipient"`
	Message         *Messagestruct   `json:"recipient"`
	Postback        *Postbackstruct  `json:"postback"`
	Read            *Read            `json:"read"`
	OptIn           *OptIn           `json:"optin"`
	ReferralMessage *ReferralMessage `json:"referral"`
	AccountLinking  *AccountLinking  `json:"account_linking"`
}

//OptIn is another event from the webhook when user clicks on the checkbox plugin, for example
type OptIn struct {
	Sender    Senderstruct    `json:"-"`
	Recipient Recipientstruct `json:"-"`
	Time      time.Time       `json:"-"`
	Ref       string          `json:"ref"`
}

//Body is the whole event sent by the webhook
type Body struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        int64 `json:"id,string"`
		Time      int64 `json:"time"`
		Messaging []struct {
			Timestamp int64
			Sender    Senderstruct    `json:"sender"`
			Recipient Recipientstruct `json:"recipient"`
			Message   *Messagestruct  `json:"message"`
			Postback  *Postbackstruct `json:"postback"`
		} `json:"messaging"`
	} `json:"entry"`
}

//Delivery is the event sent by the webhook when a message has been delivered
type Delivery struct {
	Mids      []string `json:"mids"`
	Watermark int64    `json:"watermark"`
	Seq       int      `json:"seq"`
}

//Read is the event sent by the webhook when a message has been read
type Read struct {
	Watermark int64 `json:"watermark"`
	Seq       int   `json:"seq"`
}

//GetStarted is the first button that will be seen on first interaction with our bot
type GetStarted struct {
	Payload string `json:"payload"`
}

type getStartedButton struct {
	GetS GetStarted `json:"get_started"`
}

//Quickreply is for sending quick replies on messenger
type Quickreply struct {
	ContentType string `json:"content_type, omitempty"`
	Title       string `json:"title,omitempty"`
	Payload     string `json:"payload"`
}
