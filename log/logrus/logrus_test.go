package logrus_test

import (
	"testing"

	"github.com/hadan/gogox/log"
	"github.com/hadan/gogox/log/logrus"
	sirupsen_logrus "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_Logger(t *testing.T) {
	t.Run("should not panic when logger is nil", func(t *testing.T) {
		dummyMd := log.Metadata{"key": "value"}
		logger := logrus.New(nil, log.Metadata{"service": "api"})
		assert.NotPanics(t, func() {
			logger.Trace("test")
			logger.Tracew("test", dummyMd)
			logger.Debug("test")
			logger.Debugw("test", dummyMd)
			logger.Info("test")
			logger.Infow("test", dummyMd)
			logger.Warn("test")
			logger.Warnw("test", dummyMd)
			logger.Error("test")
			logger.Errorw("test", dummyMd)
		})
	})

	t.Run("should not panic when metadata is nil", func(t *testing.T) {
		logger := logrus.New(sirupsen_logrus.New(), nil)
		assert.NotPanics(t, func() {
			logger.Trace("test")
			logger.Tracew("test", nil)
			logger.Debug("test")
			logger.Debugw("test", nil)
			logger.Info("test")
			logger.Infow("test", nil)
			logger.Warn("test")
			logger.Warnw("test", nil)
			logger.Error("test")
			logger.Errorw("test", nil)
		})
	})
}
