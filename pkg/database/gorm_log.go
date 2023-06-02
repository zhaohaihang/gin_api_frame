package database

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLogger struct {
	logger                *logrus.Logger
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
}

func NewGormLogger(l *logrus.Logger) *GormLogger {
	return &GormLogger{
		logger:                l,
	}
}

func (gl *GormLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return gl
}

func (gl *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	gl.logger.WithContext(ctx).Infof(s, args)
}

func (gl *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	gl.logger.WithContext(ctx).Warnf(s, args)
}

func (gl *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	gl.logger.WithContext(ctx).Errorf(s, args)
}

func (gl *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if gl.SourceField != "" {
		fields[gl.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && gl.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		gl.logger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if gl.SlowThreshold != 0 && elapsed > gl.SlowThreshold {
		gl.logger.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	if gl.Debug {
		gl.logger.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
	}
}
