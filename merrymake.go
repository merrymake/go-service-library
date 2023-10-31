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

func ReplyStringToOrigin(body string, contentType lib.MimeType) {
	PostToRapids("$reply", []byte(body), contentType)
}

func ReplyBytesToOrigin(body []byte, contentType lib.MimeType) {
	PostToRapids("$reply", body, contentType)
}

func ReplyFileWithMimetypeToOrigin(filepath string, contentType lib.MimeType) {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	PostToRapids("$reply", fileBytes, contentType)
}

func ReplyFileToOrigin(filepath string) {
	extension := path.Ext(filepath)[1:] // skip the dot
	mimetype := lib.GetMimeType(strings.ToLower(extension))
	ReplyFileWithMimetypeToOrigin(filepath, mimetype)
}

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
