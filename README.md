# leveledlog

`leveledlog` is a very simple Go package which enables level-based logging.


## Installation

```
go get -u github.com/ojroques/leveledlog
```


## Usage

```go
package main

import (
	"os"
	"github.com/ojroques/leveledlog"
)

func main() {
	// the logger writes to the standard error output here (os.Stderr)
	logger := leveledlog.DefaultNew(os.Stderr, leveledlog.DEBUG)
	logger.Debug("debug message %d", 1)
	logger.Info("info message")
	logger.Warning("warning message")
	logger.Error("error message")

	logger.SetPrefix(leveledlog.DEBUG, "D ")
	logger.Debug("debug message %d with a different level prefix", 2)

	logger.SetLevel(leveledlog.WARNING)
	logger.Info("this won't be printed")
	logger.Warning("this will be printed")
	logger.Error("this too")

	logger = leveledlog.New(os.Stderr, leveledlog.INFO, leveledlog.Ldate|leveledlog.Ltime)
	logger.Info("a different log entry prefix")

	logger.SetLevel(leveledlog.NONE)
	logger.Error("logging disabled")
}
```

Output:
```sh
2019/11/27 12:40:39 main.go:11 [DEBUG] debug message 1
2019/11/27 12:40:39 main.go:12 [INFO] info message
2019/11/27 12:40:39 main.go:13 [WARNING] warning message
2019/11/27 12:40:39 main.go:14 [ERROR] error message
2019/11/27 12:40:39 main.go:17 D debug message 2 with a different level prefix
2019/11/27 12:40:39 main.go:21 [WARNING] this will be printed
2019/11/27 12:40:39 main.go:22 [ERROR] this too
2019/11/27 12:40:39 [INFO] a different log entry prefix
```


## Documentation

```go
// CONSTANTS
const (
	NONE = iota
	ERROR
	WARNING
	INFO
	DEBUG
)  // levels

const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	LstdFlags     = log.LstdFlags
)  // flags (see https://golang.org/pkg/log/#pkg-constants)

// TYPE
type LeveledLog

// CONSTRUCTORS
func New(out io.Writer, loglevel uint32, flag int) *LeveledLog {}
func DefaultNew(out io.Writer, loglevel uint32) *LeveledLog {}

// METHODS
func (llogger *LeveledLog) Debug(msg string, v ...interface{}) {}
func (llogger *LeveledLog) Info(msg string, v ...interface{}) {}
func (llogger *LeveledLog) Warning(msg string, v ...interface{}) {}
func (llogger *LeveledLog) Error(msg string, v ...interface{}) {}
func (llogger *LeveledLog) SetLevel(level uint32) {}
func (llogger *LeveledLog) SetPrefix(level uint32, prefix string) {}
```
