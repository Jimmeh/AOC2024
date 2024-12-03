package helpers

type TestTextReader struct {
	Data string
}

func (r TestTextReader) Text() (string, error) {
	return r.Data, nil
}
