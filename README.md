# Simple Go logger

Probably, the simplest and idiomatic Go logger with different outputs ;)
Provides Debug, Info, Err to different system outputs with std `log` semantics.

   - USE `.SetPrefix(..), .SetFlags(..), .SetOutput(..)` to modify sub-loggers
     similar to standard Go log
   - TO DISCARD output: `Debug.SetOutput(ioutil.Discard)` (yes, the same as for std log pkg)
   - DEFAULT msg fmt: `[ERR] 2023/04/08 20:31:13 main.go:19: YOUR MESSAGE`