package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // Level
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // Log prefix
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // Output format
	Director      string `mapstructure:"director" json:"director" yaml:"director"`                   // Log folder
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // Encoding level
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // Stack name

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // Log retention time
	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // Show line numbers
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // Log to console
}

// ZapEncodeLevel returns zapcore.LevelEncoder based on EncodeLevel
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // Lowercase encoder (default)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // Lowercase encoder with color
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // Uppercase encoder
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // Uppercase encoder with color
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel converts a string to zapcore.Level
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
