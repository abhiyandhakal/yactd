package handler

import (
	"yactd/internal/protocol"
)

func Dispatch(req protocol.Request) protocol.Response {
	switch req.Action {
	case "run":
		return RunHandler(req)
	case "stop":
		return StopHandler(req)
	case "inspect":
		return InspectHandler(req)
	default:
		return protocol.Response{Status: "error", Error: "unknown action"}
	}
}

func RunHandler(req protocol.Request) protocol.Response {
	// TODO: implement run handler
	return protocol.Response{Status: "error", Error: "RunHandler not implemented"}
}

func StopHandler(req protocol.Request) protocol.Response {
	// TODO: implement stop handler
	return protocol.Response{Status: "error", Error: "StopHandler not implemented"}
}

func InspectHandler(req protocol.Request) protocol.Response {
	// TODO: implement inspect handler
	return protocol.Response{Status: "error", Error: "InspectHandler not implemented"}
}
