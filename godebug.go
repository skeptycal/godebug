package godebug

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var (
	defaultConfig Session
	config        Session
)

// init initializes the session
func init() {
	defaultConfig.Start("anansi", true, false, DEBUG, LogDEBUG)

	if !config.isActive() {
		config = defaultConfig
	}
	if config.IsLogger() {
		LogFormatter := new(log.TextFormatter)
		LogFormatter.TimestampFormat = "02-01-2006 15:04:05"
		LogFormatter.FullTimestamp = true
		log.SetFormatter(LogFormatter)

		log.Info("logrus initialized")
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
