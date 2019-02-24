package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _ Logger = new(ZapLogger)

type ZapLogger struct {
	path  string
	rlog  *lumberjack.Logger
	log   *zap.Logger
	sugar *zap.SugaredLogger

	level zapcore.Level

	rolling    bool
	lastRotate time.Time
}

func NewLogger(path string, level zapcore.Level) (*ZapLogger, error) {
	out := new(ZapLogger)
	out.rlog = new(lumberjack.Logger)

	out.path = path
	out.lastRotate = time.Now()
	out.level = level

	// config lumberjack
	out.rlog.Filename = path
	out.rlog.MaxSize = 0x1000 * 2 // automatic rolling file on it increment than 2GB
	out.rlog.LocalTime = true
	out.rlog.Compress = true
	out.rlog.MaxBackups = 60 // reserve last 60 day logs

	// config encoder config
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	ec.EncodeTime = zapcore.ISO8601TimeEncoder

	// config core
	c := zapcore.AddSync(out.rlog)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(ec), c, out.level)
	out.log = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)

	out.sugar = out.log.Sugar()
	return out, nil
}

func (z *ZapLogger) Debug(v ...interface{}) {
	if !z.level.Enabled(zap.DebugLevel) {
		return
	}
	defer z.log.Sync()
	z.sugar.Debug(v)
}

func (z *ZapLogger) Debugf(format string, v ...interface{}) {
	if !z.level.Enabled(zap.DebugLevel) {
		return
	}
	defer z.log.Sync()
	z.sugar.Debugf(format, v...)
}

func (z *ZapLogger) Info(v ...interface{}) {
	if !z.level.Enabled(zap.InfoLevel) {
		return
	}
	defer z.log.Sync()
	z.sugar.Info(v)
}

func (z *ZapLogger) Infof(format string, v ...interface{}) {
	if !z.level.Enabled(zap.InfoLevel) {
		return
	}
	defer z.log.Sync()
	z.sugar.Infof(format, v...)
}

func (z *ZapLogger) Error(v ...interface{}) {
	if !z.level.Enabled(zap.ErrorLevel) {
		return
	}
	defer z.log.Sync()
	z.sugar.Error(v)
}

func (z *ZapLogger) Errorf(format string, v ...interface{}) {
	if !z.level.Enabled(zap.ErrorLevel) {
		return
	}
	defer z.log.Sync()
	z.sugar.Errorf(format, v...)
}
