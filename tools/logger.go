package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"path/filepath"
	"time"
)

const fileNameFormat = "06-01-02"

var (
	logger *zap.Logger
)

func InitLogger(prod bool) {
	// production config
	cfg := zap.NewProductionConfig()
	cfg.Development = !prod
	cfg.Sampling = &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	}
	if prod {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		cfg.Encoding = "json"
	} else {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		cfg.Encoding = "console"
	}

	// encoder
	encodeCfg := zap.NewProductionEncoderConfig()
	encodeCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	if prod {
		encodeCfg.LevelKey = ""
	} else {
		encodeCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	cfg.EncoderConfig = encodeCfg
	var encoder zapcore.Encoder
	if prod {
		encoder = zapcore.NewJSONEncoder(encodeCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encodeCfg)
	}

	// writer
	infoDir := path.Join("logs", "info")
	warnDir := path.Join("logs", "warn")
	errDir := path.Join("logs", "err")
	pacDir := path.Join("logs", "pac")
	fatDir := path.Join("logs", "fat")
	err := CreateDir(infoDir)
	if err != nil {
		panic(err)
	}
	err = CreateDir(warnDir)
	if err != nil {
		panic(err)
	}
	err = CreateDir(errDir)
	if err != nil {
		panic(err)
	}
	err = CreateDir(pacDir)
	if err != nil {
		panic(err)
	}
	err = CreateDir(fatDir)
	if err != nil {
		panic(err)
	}
	infoWriteSyncer := &dateWriteSyncer{outPath: infoDir}
	warnWriteSyncer := &dateWriteSyncer{outPath: warnDir}
	errWriteSyncer := &dateWriteSyncer{outPath: errDir}
	pacWriteSyncer := &dateWriteSyncer{outPath: pacDir}
	fatWriteSyncer := &dateWriteSyncer{outPath: fatDir}

	// core
	core := zapcore.NewTee(
		// info
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriteSyncer), zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
			return lv < zapcore.WarnLevel
		})),
		// warn
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriteSyncer), zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
			return (lv >= zapcore.WarnLevel) && (lv < zapcore.ErrorLevel)
		})),
		// error
		zapcore.NewCore(encoder, zapcore.AddSync(errWriteSyncer), zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
			return (lv >= zapcore.ErrorLevel) && (lv < zapcore.DPanicLevel)
		})),
		// panic
		zapcore.NewCore(encoder, zapcore.AddSync(pacWriteSyncer), zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
			return (lv >= zapcore.DPanicLevel) && (lv < zapcore.FatalLevel)
		})),
		// fatal
		zapcore.NewCore(encoder, zapcore.AddSync(fatWriteSyncer), zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
			return lv >= zapcore.FatalLevel
		})),
	)

	// logger
	logger = zap.New(core)
	//var err error
	//logger, err = cfg.Build()
	//if err != nil {
	//	panic(err)
	//}
}

type dateWriteSyncer struct {
	file    *os.File
	outPath string
}

func (d *dateWriteSyncer) Write(p []byte) (n int, err error) {
	fileName := filepath.Join(d.outPath, time.Now().Format(fileNameFormat)+".log")
	if d.file == nil || d.file.Name() != fileName {
		if d.file != nil {
			_ = d.file.Close()
		}
		d.file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return 0, err
		}
	}
	return d.file.Write(p)
}

func (d *dateWriteSyncer) Sync() error {
	if d.file != nil {
		return d.file.Sync()
	}
	return nil
}

func LogExit() {
	_ = logger.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
