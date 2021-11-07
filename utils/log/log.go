package log

import (
	"os"
	"strings"

	logging "github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultLogLevel = logging.INFO
)

var (
	// Logger instance for quick declarative logging levels
	Logger = logging.MustGetLogger("eth-swap-ap")

	// log levels that are available
	levels = map[string]logging.Level{
		"CRITICAL": logging.CRITICAL,
		"ERROR":    logging.ERROR,
		"WARNING":  logging.WARNING,
		"NOTICE":   logging.NOTICE,
		"INFO":     logging.INFO,
		"DEBUG":    logging.DEBUG,
	}
)

type LogsConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"maxFileSizeInMB"`
	MaxBackupsOfLogFiles         int    `json:"maxBackupsOfLogFiles"`
	MaxAgeToRetainLogFilesInDays int    `json:"maxAgeToRetainLogFilesInDays"`
	UseConsoleLogger             bool   `json:"useConsoleLogger"`
	UseFileLogger                bool   `json:"useFileLogger"`
	Compress                     bool   `json:"compress"`
}

// InitLogger initialises the logger.
func InitLogger(config LogsConfig) {
	backends := make([]logging.Backend, 0)
	if config.UseConsoleLogger {
		consoleFormat := logging.MustStringFormatter(`%{level} %{color}%{time:2006-01-02T15:04:05.000} %{shortfunc}%{color:reset} %{message}`)
		consoleLogger := logging.NewLogBackend(os.Stderr, "", 0)
		consoleFormatter := logging.NewBackendFormatter(consoleLogger, consoleFormat)
		consoleLoggerLeveled := logging.AddModuleLevel(consoleFormatter)
		consoleLoggerLeveled.SetLevel(levels[config.Level], "")
		backends = append(backends, consoleLoggerLeveled)

	}

	if config.UseFileLogger {
		fileLogger := logging.NewLogBackend(&lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxFileSizeInMB,              // MaxSize is the maximum size in megabytes of the log file
			MaxBackups: config.MaxBackupsOfLogFiles,         // MaxBackups is the maximum number of old log files to retain
			MaxAge:     config.MaxAgeToRetainLogFilesInDays, // MaxAge is the maximum number of days to retain old log files
			Compress:   config.Compress,
		}, "", 0)
		fileFormat := logging.MustStringFormatter(`%{level} %{time:2006-01-02T15:04:05.000} %{shortfunc} %{message}`)
		fileFormatter := logging.NewBackendFormatter(fileLogger, fileFormat)
		fileLoggerLeveled := logging.AddModuleLevel(fileFormatter)
		fileLoggerLeveled.SetLevel(levels[config.Level], "")
		backends = append(backends, fileLoggerLeveled)
	}

	logging.SetBackend(backends...)
}

func GetLogger(level logging.Level) func(string, ...interface{}) {
	if !Logger.IsEnabledFor(level) {
		return GetLogger(defaultLogLevel)
	}
	switch level {
	case levels["CRITICAL"]:
		return Logger.Criticalf
	case levels["ERROR"]:
		return Logger.Errorf
	case levels["WARNING"]:
		return Logger.Warningf
	case levels["NOTICE"]:
		return Logger.Noticef
	case levels["DEBUG"]:
		return Logger.Debugf
	default:
		return Logger.Infof
	}
}

func ParsePrefixedLogString(str string) (level logging.Level, unPrefixed string) {
	level = LevelForPrefixedLogString(str)
	unPrefixed = str
	if level != defaultLogLevel {
		split := strings.SplitN(str, " ", 2)
		if len(split) == 2 {
			unPrefixed = split[1]
		}
	}
	return level, unPrefixed
}

func LevelForPrefixedLogString(str string) (level logging.Level) {
	level = defaultLogLevel
	split := strings.SplitN(str, " ", 2)
	if len(split) < 2 {
		return level
	}
	for key, lvl := range levels {
		if key != split[0] {
			continue
		}
		level = lvl
		break
	}
	return level
}
