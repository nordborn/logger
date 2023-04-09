// Package logger
// Probably, the simplest and idiomatic Go logger with different outputs ;)
// Provides Debug, Info, Err to different system outputs with std log semantics.
//   - USE .SetPrefix(..), .SetFlags(..), .SetOutput(..) to modify sub-loggers
//     similar to standard Go log
//   - TO DISCARD output: Debug.SetOutput(ioutil.Discard) (yes, the same as for std log pkg)
//   - DEFAULT msg fmt: `[ERR] 2023/04/08 20:31:13 main.go:19: YOUR MESSAGE`
package logger

import (
	"log"
	"os"
)

const flag = log.Ldate | log.Ltime | log.LUTC | log.Lshortfile

var (
	Debug = log.New(os.Stdout, "[DBG] ", flag)
	Info  = log.New(os.Stdout, "[INF] ", flag)
	Err   = log.New(os.Stderr, "[ERR] ", flag)
)
