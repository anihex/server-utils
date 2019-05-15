package tools

import (
	"log"
	"os"
)

// DefaultLogger is a regular logger. It's defined the same way as the
// standard logger of the log package.
// This means it outputs the messages to StdErr.
var DefaultLogger = log.New(os.Stderr, "", log.LstdFlags)

// DummyLogger is a dummy logger. It uses a DummyWriter to skip all logs.
var DummyLogger = log.New(&DummyWriter{}, "", log.LstdFlags)
