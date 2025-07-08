package container

import (
	"os/exec"
	"yactd/internal/state"
)

func RunContainer(image string, cmd []string) (string, error) {
	id := generateContainerID()
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Dir = setupChroot(image)
	output, err := command.CombinedOutput()
	// Register container state (stub)
	state.Register(id, 0) // TODO: use real PID
	return string(output), err
}

func generateContainerID() string {
	// TODO: implement real ID generation
	return "container-123"
}

func setupChroot(image string) string {
	// TODO: implement real chroot setup
	return "/tmp/fake-chroot/" + image
}
