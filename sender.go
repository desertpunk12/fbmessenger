package fbmessenger

import (
	"strings"
	"net/http"
)

//QuickReply contains the quickreply data
type QuickReply struct {
	ContentType string
	Title       string
	Payload     string
	ImageURL    string
}

//SendMessage sends the message via your facebook page
func SendMessage(message, recipientid, pageaccesstoken string) {
	text := `{
  			"recipient":{
				"id":"` + recipientid + `"
			},
			"message":{
				"text":"` + message + `"
			}
		}`

	fbURL := "https://graph.facebook.com/v2.6/me/messages?access_token=" + pageaccesstoken
	_, err := http.Post(fbURL, "application/json", strings.NewReader(text))
	if err != nil {
		panic(err) }

}