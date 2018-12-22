package bot

type Event int

const (
	UnknownEventType Event = iota - 1

	TextEvent
	DeliveryEvent
	PostBackEvent
	ReadEvent
	OptInEvent
	ReferralEvent
	AccountLinkingEvent
)
