package debug

import (
	"log"
	"os"
	"time"
)

func init() {
	log.Println("init in debug.go")
}

// Session defines session information for DEV or PRODUCTION modes
type Session struct {
	Name         string
	IsDevMode    bool
	UseLogger    bool
	Verbose      VerboseLevel
	sessionStart time.Time
	sessionEnd   time.Time
	userID       int
	defaults     defaultValues
}

func (s *Session) init(name string, mode bool, useLogger bool, useVerbose VerboseLevel) {
	s.Name = name
	s.IsDevMode = mode
	s.UseLogger = useLogger
	s.Verbose = useVerbose
	s.sessionStart = time.Now()
	s.userID = os.Getuid()
}

func logPrint(v ...interface{}) {
	if config.Verbose >= verboseDebug {
		log.Println("----------")
		defer log.Println("----------")
		log.Println(v...)
	}
}

var defaultConfig Session

func init() {
	defaultConfig.init("anansi", true, false, verboseAll)
}
