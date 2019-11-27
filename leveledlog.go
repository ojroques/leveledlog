package leveledlog

import (
	"fmt"
	"io"
	"log"
	"runtime"
)

const (
	NONE = iota
	ERROR
	WARNING
	INFO
	DEBUG
)

const (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LUTC          = log.LUTC
	LstdFlags     = log.LstdFlags
)

type LeveledLog struct {
	logger   *log.Logger
	loglevel uint32
	prefixes [4]string
	flags    int
}

func New(out io.Writer, loglevel uint32, flag int) *LeveledLog {
	var llogger LeveledLog
	llogger.logger = log.New(out, "", flag&^Lshortfile&^Llongfile)
	llogger.loglevel = loglevel
	llogger.prefixes = [4]string{"[ERROR] ", "[WARNING] ", "[INFO] ", "[DEBUG] "}
	llogger.flags = flag
	return &llogger
}

func DefaultNew(out io.Writer, loglevel uint32) *LeveledLog {
	return New(out, loglevel, Ldate|Ltime|Lmicroseconds|Lshortfile)
}

func (llogger *LeveledLog) formatFile(file string, line int) string {
	if llogger.flags&(Lshortfile|Llongfile) == 0 {
		return ""
	}

	if llogger.flags&Lshortfile != 0 {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
	}

	file = fmt.Sprintf("%s:%d ", file, line)
	return file
}

func (llogger *LeveledLog) printf(entrylevel uint32, msg string, v ...interface{}) {
	if (entrylevel < ERROR) || (DEBUG < entrylevel) {
		return
	}

	if llogger.loglevel < entrylevel {
		return
	}

	file := ""
	if _, f, l, ok := runtime.Caller(2); ok {
		file = llogger.formatFile(f, l)
	}
	fmsg := fmt.Sprintf(msg, v...)
	llogger.logger.Printf("%s%s%s\n", file, llogger.prefixes[entrylevel-1], fmsg)
}

func (llogger *LeveledLog) Debug(msg string, v ...interface{}) {
	llogger.printf(DEBUG, msg, v...)
}

func (llogger *LeveledLog) Info(msg string, v ...interface{}) {
	llogger.printf(INFO, msg, v...)
}

func (llogger *LeveledLog) Warning(msg string, v ...interface{}) {
	llogger.printf(WARNING, msg, v...)
}

func (llogger *LeveledLog) Error(msg string, v ...interface{}) {
	llogger.printf(ERROR, msg, v...)
}

func (llogger *LeveledLog) SetLevel(level uint32) {
	llogger.loglevel = level
}

func (llogger *LeveledLog) SetPrefix(level uint32, prefix string) {
	if (level < ERROR) || (DEBUG < level) {
		return
	}
	llogger.prefixes[level-1] = prefix
}
