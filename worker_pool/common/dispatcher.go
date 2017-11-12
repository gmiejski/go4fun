package common


type JobDoer interface {
	Do(payload Payload)
}

