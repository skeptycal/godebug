package godebug

import (
	"log"
	"os"
	"time"
)

func init() {
	log.Println("init in debug.go")
}

// VerboseLevel constants describe the level of output and logging.
// Output will be every category where ("verbose setting variable" >= VerboseLevel)
//
//      TRACE VerboseLevel = 5
//      DEBUG VerboseLevel = 10
//      INFO VerboseLevel = 20
//      SUCCESS VerboseLevel = 25
//      WARNING VerboseLevel = 30
//      ERROR VerboseLevel = 40
//      CRITICAL VerboseLevel = 50
//
// e.g.
//      DEBUG will output all debug info, successes, warnings, errors
//      ERROR will output only error info
//      SUCCESS will output only success info and above (warnings,etc.)
type VerboseLevel int8

const (
	// TRACE - Output every dam thing
	TRACE VerboseLevel = 5
	// DEBUG - Output all including debug info
	DEBUG VerboseLevel = 10
	// INFO - Output standard information
	INFO VerboseLevel = 20
	// SUCCESS - Output successful task and errors
	SUCCESS VerboseLevel = 25
	// WARNING - Output all nonfatal warnings and errors
	WARNING VerboseLevel = 30
	// ERROR - Output only Fatal errors
	ERROR VerboseLevel = 40
	// CRITICAL - Output only Panic errors
	CRITICAL VerboseLevel = 50
)

type devMode bool

// devMode returns true if the build is in dev mode
const (
	production  devMode = false
	development devMode = true
)

// Session defines session information for DEV or PRODUCTION modes
type Session struct {
	Name         string
	IsDevMode    devMode // DEV or PRODUCTION
	UseLogger    bool    // use Log?
	Verbose      VerboseLevel
	sessionStart time.Time
	sessionEnd   time.Time
	userID       int
}

func (s *Session) init(name string, mode devMode, useLogger bool, verbose VerboseLevel) {
	s.Name = name
	s.IsDevMode = mode
	s.UseLogger = useLogger
	s.Verbose = verbose
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
var defaultValues Session

func init() {
	defaultConfig.init("anansi", true, false, verboseAll)
}
