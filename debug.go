package godebug

import (
	"fmt"
	"log"
	"os"
	"time"
)

var defaultConfig Session
var config = defaultConfig

func init() {
	defaultConfig.Start("anansi", true, false, DEBUG)
	LogPrintln("init in debug.go")
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
	IsDevMode    devMode
	UseLogger    bool
	Verbose      VerboseLevel
	sessionStart time.Time
	sessionEnd   time.Time
	userID       int
	active       bool
}

// Start starts a new session with output and logging specified.
func (s *Session) Start(name string, mode devMode, useLogger bool, verbose VerboseLevel) {
	s.Name = name
	s.IsDevMode = mode
	s.UseLogger = useLogger
	s.Verbose = verbose
	s.sessionStart = time.Now()
	s.userID = os.Getuid()
	s.active = true
}

// Stop stops the session.
func (s *Session) Stop() {
	s.sessionEnd = time.Now()
	s.active = false
}

// LogPrintln respects the VerboseLevel setting in the session configuration
func LogPrintln(v ...interface{}) {
	if config.Verbose <= DEBUG {
		log.Println("----------")
		defer log.Println("----------")
		log.Println(v...)
	}
}

// Println prints while respecting session configuration
func Println(v ...interface{}) error {
	if !config.active {
		return fmt.Errorf("cannot print when debug session is not active")
	}
	if config.UseLogger {
		log.Println(v...)
	}
	fmt.Println(v...)
	return nil
}
