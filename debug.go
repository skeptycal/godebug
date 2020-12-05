package godebug

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	defaultConfig Session
	config        Session

	// LogFormatter - configures the log formatter
	LogFormatter = new(log.TextFormatter)
)

// session defines session information for DEV or PRODUCTION modes
// private fields are automatically initialized and managed
type session struct {
	name         string
	isDevMode    bool
	isLogger     bool
	verbose      VerboseLevel
	logLevel     LogLevel
	sessionStart time.Time
	sessionEnd   time.Time
	userID       int
	active       bool
}

type Session interface {
	Start(name string, devMode bool, useLogger bool, verbose VerboseLevel, logLevel LogLevel)
	Stop()
	Name() string
	IsDevMode() bool
	IsLogger() bool
	Verbose() VerboseLevel
	SetLogLevel() LogLevel
	SetVerbose(verbose VerboseLevel)
	SetDebugLevel(logLevel LogLevel)

	// private methods are managed by the gologLevel package
	isActive() bool
	whoami() int
}

// Start starts a new session with output and logging specified.
func (s session) Start(name string, devMode bool, useLogger bool, verbose VerboseLevel, logLevel LogLevel) {
	s.name = name
	s.isDevMode = devMode
	s.isLogger = useLogger
	s.verbose = verbose
	s.logLevel = logLevel
	s.sessionStart = time.Now()
	s.userID = os.Getuid()
	s.active = true
}

// Stop stops the session.
func (s *session) Stop() {
	s.sessionEnd = time.Now()
	s.active = false
}

func (s *session) Name() string {
	return s.name
}

// IsLogger returns true if the session is using a logger.
func (s *session) IsLogger() bool {
	return s.isLogger
}

// IsDevMode returns true if the session is in DEV mode.
func (s *session) IsDevMode() bool {
	return s.isDevMode
}

func (s *session) Verbose() VerboseLevel {
	return s.verbose
}

func (s *session) SetVerbose(verbose VerboseLevel) {
	s.verbose = verbose
}

// whoami returns the userID of the session.
func (s *session) whoami() int {
	return s.userID
}

// isActive returns true if the session is active.
func (s *session) isActive() bool {
	return s.active
}

// init initializes the session
func init() {
	if !config.isActive() {
		config = defaultConfig
		defaultConfig.Start("anansi", true, false, DEBUG, LogDEBUG)
	}
	if config.IsLogger() {
		LogFormatter.TimestampFormat = "02-01-2006 15:04:05"
		LogFormatter.FullTimestamp = true
		log.SetFormatter(LogFormatter)
		log.Info("logrus initialized")

		log.Info("Example info: Some info. Earth is not flat.")
		log.Warning("Example warning: This is a warning")
		log.Error("Example error: Not fatal. An error. Won't stop execution")
		log.Fatal("Example fatal: MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
		log.Panic("Example panic: Do not panic")
	}
}

// VerboseLevel constants describe the level of output and logging.
/* Output will be every category where ("verbose setting variable" >= VerboseLevel)

        TRACE VerboseLevel = 5
        DEBUG VerboseLevel = 10
        INFO VerboseLevel = 20
        SUCCESS VerboseLevel = 25
        WARNING VerboseLevel = 30
        ERROR VerboseLevel = 40
        CRITICAL VerboseLevel = 50

   e.g.
        DEBUG will output all logLevel info, successes, warnings, errors
        ERROR will output only error info
        SUCCESS will output only success info and above (warnings,etc.)
*/
type VerboseLevel int8

const (
	// TRACE - Output every dam thing
	TRACE VerboseLevel = 5
	// DEBUG - Output all including logLevel info
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

type LogLevel int8

const (
	// TRACE - Output every dam thing
	LogTRACE LogLevel = 5
	// DEBUG - Output all including logLevel info
	LogDEBUG LogLevel = 10
	// INFO - Output standard information
	LogINFO LogLevel = 20
	// SUCCESS - Output successful task and errors
	LogSUCCESS LogLevel = 25
	// WARNING - Output all nonfatal warnings and errors
	LogWARNING LogLevel = 30
	// ERROR - Output only Fatal errors
	LogERROR LogLevel = 40
	// CRITICAL - Output only Panic errors
	LogCRITICAL LogLevel = 50
)

// LogPrintln respects the VerboseLevel setting in the session configuration
// func LogPrintln(v ...interface{}) {
// 	if config.Verbose <= DEBUG {
// 		log.Println("----------")
// 		defer log.Println("----------")
// 		log.Println(v...)
// 	}
// }

// Println prints while respecting session configuration
func Println(v ...interface{}) error {
	if !config.isActive() {
		return fmt.Errorf("cannot print when logLevel session is not active")
	}
	if config.IsLogger() {
		log.Info(v...)
	}
	fmt.Println(v...)
	return nil
}
