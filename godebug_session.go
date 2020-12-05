package godebug

import (
	"os"
	"time"
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
	LogLevel() LogLevel
	SetVerbose(verbose VerboseLevel)
	SetLogLevel(logLevel LogLevel)

	// private methods are managed by the godebug package
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

// Name returns the name of the session.
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

// Verbose returns the verbosity level.
func (s *session) Verbose() VerboseLevel {
	return s.verbose
}

// SetVerbose sets the verbosity level.
func (s *session) SetVerbose(verbose VerboseLevel) {
	s.verbose = verbose
}

// LogLevel returns the log level.
func (s *session) LogLevel() LogLevel {
	return s.logLevel
}

// SetLogLevel sets the log level.
func (s *session) SetLogLevel(logLevel LogLevel) {
	s.logLevel = logLevel
}

// whoami returns the userID of the session.
func (s *session) whoami() int {
	return s.userID
}

// isActive returns true if the session is active.
func (s *session) isActive() bool {
	return s.active
}
