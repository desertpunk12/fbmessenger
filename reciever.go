package fbmessenger

import (
	"encoding/json"
	"fmt"
	"io"
	"bytes"
)

//Recieved represents a single sent message
type Recieved struct {
	//Page should be equal to "page"
	Page  string
	Entry []*Entry
}

//Entry contains event data
type Entry struct {
	ID        string
	Time      json.Number
	Messaging []Messaging
}

//Messaging contains the Message and Sender and Reciever ID's
type Messaging struct {
	Sender    struct{ ID string }
	Recipient struct{ ID string }
	Timestamp json.Number
	Message   Message
}

//Message contains the actual message
type Message struct {
	Mid        string
	Seq        json.Number
	Text       string
	Attachment []struct {
		Type    string
		Payload struct {
			URL string
		}
	}
}

//Parse the recieved reader into a Recieved type
func Parse(body io.Reader) *Recieved {
	buff := new(bytes.Buffer)
	buff.ReadFrom(body)

	var rec Recieved
	err := json.Unmarshal(buff.Bytes(), &rec)
	if err != nil {
		panic(err)
	}


	fmt.Printf("Message sent by %v: %v", rec.Entry[0].Messaging[0].Sender.ID, rec.Entry[0].Messaging[0].Message.Text)

	return &rec
}


//ParseMin an io.Reader into a minimum Messaging type
func ParseMin(body io.Reader) *Messaging {
	rec := Parse(body)
	return &rec.Entry[0].Messaging[0]
}
