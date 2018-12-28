package bot

type AttachmentType string
type MessagingType string
type TopElementStyle string
type ImageAspectRatio string

//SendTextWithQuickReplies allows us reply the user with a text and optionally attach a quick reply/quick replies
func (mes *Messenger) SendTextWithQuickReplies(to Recipientstruct, message string, replies []Quickreply, msgType MessagingType, tags ...string) error {
	response := &Response{
		token: mes.token,
		to: to,
	}
}


func (mes *Messenger) SendText(to Recipientstruct, message string, msgType MessagingType, tags ...string) {
	//essentially we're calling sendtextwithquiclreplies but since the quick replies
	//are actually optional, we can simply pass nil and thus send just text to the user

	//tags... essentially works the same way spread parameters work in javascript
	return mes.SendTextWithQuickReplies(to, message, nil, MessagingType, tags...)
}
