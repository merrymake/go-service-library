package lib

type NullMerrymake struct {
	action       string
	envelope     Envelope
	payloadBytes []byte
}

func (m NullMerrymake) Handle(action string, handler func([]byte, Envelope)) IMerrymake { return m }
func (m NullMerrymake) Initialize(f func())                                             {}
