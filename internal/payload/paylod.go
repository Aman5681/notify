package payload

type Payload struct {
	Action string `json:"action"`
	Data   any    `json:"data"`
}
