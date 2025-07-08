package state

type ContainerState struct {
	ID  string
	PID int
	// other fields as needed
}

var state = map[string]*ContainerState{}

func Register(id string, pid int) {
	state[id] = &ContainerState{ID: id, PID: pid}
}

func Stop(id string) error {
	// TODO: implement stop logic
	delete(state, id)
	return nil
}

func Inspect(id string) *ContainerState {
	return state[id]
}
