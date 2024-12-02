package helpers

type TestLineReader struct {
	Data []string
}

func (r TestLineReader) Lines() ([]string, error) {
	return r.Data, nil
}
