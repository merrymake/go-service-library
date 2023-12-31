package merrymake

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"github.com/merrymake/go-service-library/lib"
)

type Merrymake struct {
	action       string
	envelope     lib.Envelope
	payloadBytes []byte
}

// This is the root call for a Merrymake service.
// # Arguments
// * `args` -- the arguments from the main method
// # Returns
// A Merrymake builder to make further calls on
func (m Merrymake) Service(args []string) lib.IMerrymake {

	m.action = args[len(args)-2]

	err := json.Unmarshal([]byte(os.Args[len(os.Args)-1]), &m.envelope)
	if err != nil {
		fmt.Println(err)
	}

	m.payloadBytes, err = io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}

	return m
}

func (m Merrymake) Handle(action string, handler func([]byte, lib.Envelope)) lib.IMerrymake {

	if m.action == action {
		handler(m.payloadBytes, m.envelope)
		return new(lib.NullMerrymake)
	}

	return m
}

func (m Merrymake) Initialize(f func()) {
	f()
}

// Post a reply back to the originator of the trace, with a payload and its
// content type.
// # Arguments
// * `body` --        the payload
// * `contentType` -- the content type of the payload
func ReplyStringToOrigin(body string, contentType lib.MimeType) {
	PostToRapids("$reply", []byte(body), contentType)
}

// Post a reply back to the originator of the trace, with a payload and its
// content type.
// # Arguments
// * `body` --        the payload
// * `contentType` -- the content type of the payload
func ReplyBytesToOrigin(body []byte, contentType lib.MimeType) {
	PostToRapids("$reply", body, contentType)
}

// Send a file back to the originator of the trace.
// # Arguments
// * `path` --        the path to the file starting from main/resources
// * `contentType` -- the content type of the file
func ReplyFileWithMimetypeToOrigin(filepath string, contentType lib.MimeType) {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	PostToRapids("$reply", fileBytes, contentType)
}
// Send a file back to the originator of the trace.
// # Arguments
// * `path` -- the path to the file starting from main/resources
func ReplyFileToOrigin(filepath string) {
	extension := path.Ext(filepath)[1:] // skip the dot
	mimetype := lib.GetMimeType(strings.ToLower(extension))
	ReplyFileWithMimetypeToOrigin(filepath, mimetype)
}

// Post an event to the central message queue (Rapids), with a payload and its
// content type.
// # Arguments
// * `event` --       the event to post
// * `body` --        the payload
// * `contentType` -- the content type of the payload
func PostStringToRapids(event string, body string, contentType lib.MimeType) {
	PostToRapids(event, []byte(body), contentType)
}

// Post an event to the central message queue (Rapids), with a payload and its
// content type.
// # Arguments
// * `event` --       the event to post
// * `body` --        the payload
// * `contentType` -- the content type of the payload
func PostToRapids(event string, body []byte, contentType lib.MimeType) {

	url := fmt.Sprintf("%s/%s", os.Getenv("RAPIDS"), event)

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", contentType.ToString())

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

}

// Subscribe to a channel, so events will stream back messages broadcast to that
// channel. You can join multiple channels. You stay in the channel until the
// request is terminated.
//
// Note: The origin-event has to be set as "streaming: true" in the
// event-catalogue.
// # Arguments
// * `channel` -- the channel to join
func JoinChannel(channel string) {
	PostStringToRapids("$join", channel, lib.GetMimeType("txt"))
}

type BroadcastObject struct {
	To string `json:"to"`
	Event string `json:"event"`
	Payload string `json:"payload"`
}

// Broadcast a message (event and payload) to all listeners in a channel.
// # Arguments
// * `to` -- the channel to broadcast to
// * `event` -- the event-type of the message
// * `payload` -- the payload of the message
func BroadcastToChannel(to string, event string, payload string) {
	obj := &BroadcastObject{ To: to, Event: event, Payload: payload }
	b, _ := json.Marshal(obj)
	PostStringToRapids("$broadcast", string(b), lib.GetMimeType("json"))
}