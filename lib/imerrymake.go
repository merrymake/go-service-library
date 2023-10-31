package lib

type IMerrymake interface {
	Handle(string, func([]byte, Envelope)) IMerrymake
	Initialize(func())
}
