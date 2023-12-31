package lib

type Envelope struct {
	// Id of this particular message.
	// Note: it is _not_ unique, since multiple rivers can deliver the same message.
	// The combination of (river, messageId) is unique.
	MessageId string `json:"messageId"`
	// Id shared by all messages in the current trace, ie. stemming from the same
	// origin.
	TraceId   string `json:"traceId"`
	// (Optional) Id corresponding to a specific originator. This id is rotated
	// occasionally, but in the short term it is unique and consistent. Same
	// sessionId implies the trace originated from the same device.
	SessionId string `json:"sessionId"`}
