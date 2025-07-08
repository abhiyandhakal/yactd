package runtime

type Runtime struct {
	Image string
	Idle  bool
	// other fields as needed
}

func FindOrCreateRuntime(image string) (*Runtime, error) {
	// TODO: implement pool lookup/creation
	return &Runtime{Image: image, Idle: true}, nil
}

func AttachContainerToRuntime(runtime *Runtime, cmd []string) (string, error) {
	// TODO: implement attaching logic
	return "attached", nil
}
