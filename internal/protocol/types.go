package protocol

type Request struct {
	Action string   `json:"action"`
	Image  string   `json:"image"`
	Cmd    []string `json:"cmd"`
	ID     string   `json:"id,omitempty"`
}

type Response struct {
	Status string `json:"status"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}
