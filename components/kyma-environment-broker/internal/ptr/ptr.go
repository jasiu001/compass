package ptr

func Bool(in bool) *bool {
	return &in
}

func String(str string) *string {
	return &str
}

func Integer(in int) *int {
	return &in
}
