package tools

// DummyWriter is a simple writer that doesn't log anything.
type DummyWriter struct{}

func (dw *DummyWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// H ist ein Hilfsdatentyp um nicht immer map[string]interface{} schreiben zu
// müssen.
type H map[string]interface{}

// TimeType wird zum Benchmarken von Abfragen benötigt.
type TimeType int
