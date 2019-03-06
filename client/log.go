package client

import (
	"github.com/btcsuite/btclog"
	"os"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var (
	backendLog     = btclog.NewBackend(logWriter{})
	logger         = backendLog.Logger("CLIENT")
	servicesLogger = backendLog.Logger("SERVICES")
)

// logWriter implements an io.Writer that outputs to both standard output and
// the write-end pipe of an initialized log rotator.
type logWriter struct{}

func (logWriter) Write(p []byte) (n int, err error) {
	os.Stdout.Write(p)
	return len(p), nil
}
