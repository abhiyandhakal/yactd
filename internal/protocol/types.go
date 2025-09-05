package protocol

type Request struct {
	Action  string            `json:"action"`  // The command action (e.g., "pull", "run", "list")
	Payload map[string]string `json:"payload"` // Flexible key-value pairs for command arguments
}

type Response struct {
	Status string `json:"status"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}
