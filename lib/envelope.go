package lib

type Envelope struct {
	MessageId string `json:"messageId"`
	TraceId   string `json: traceId`
}
