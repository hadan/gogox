package errorx_test

import (
	"fmt"
	"testing"

	"github.com/hadan/gogox/errorx"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	e := errorx.New("some_code", "some_string")

	assert.Equal(t, "some_code", e.Code)
	assert.Equal(t, "some_string", e.Error())
	assert.Equal(t, fmt.Sprintf("[%s] %s", e.Code, e.Error()), e.LogError())
}

func Test_Newf(t *testing.T) {
	e := errorx.Newf("some_code", "%d %d", 1, 2)

	assert.Equal(t, "some_code", e.Code)
	assert.Equal(t, "1 2", e.Error())
	assert.Equal(t, fmt.Sprintf("[%s] %s", e.Code, e.Error()), e.LogError())
}

func Test_NewWithLog(t *testing.T) {
	e := errorx.NewWithLog(errorx.CodeInternal, "some_string", "some custom log message with more information")

	assert.Equal(t, errorx.CodeInternal, e.Code)
	assert.Equal(t, "some custom log message with more information", e.LogError())
}

func Test_NewfWithLog(t *testing.T) {
	e := errorx.NewfWithLog(errorx.CodeInternal, "%d %d", "some custom log message with more information", 1, 2)

	assert.Equal(t, errorx.CodeInternal, e.Code)
	assert.Equal(t, "1 2", e.Error())
	assert.Equal(t, "some custom log message with more information", e.LogError())
}

func Test_Parse(t *testing.T) {
	notStandardError := fmt.Errorf("some error")
	e1, ok1 := errorx.Parse(notStandardError)
	assert.Nil(t, e1)
	assert.False(t, ok1)

	standardError := errorx.New(errorx.CodeInternal, "unauthorized")
	e2, ok2 := errorx.Parse(standardError)
	assert.Equal(t, errorx.CodeInternal, e2.Code)
	assert.Equal(t, "unauthorized", e2.Error())
	assert.True(t, ok2)
}

func Test_Wrapf(t *testing.T) {
	cause := fmt.Errorf("original error")

	err := errorx.Wrapf(cause, "upper error", "upper error message %d", 1)
	assert.Equal(t, "[upper error] upper error message 1: original error", err.LogError())
}

func Test_WrapfWithLog(t *testing.T) {
	cause := fmt.Errorf("original error")

	err := errorx.WrapfWithLog(cause, "upper error", "upper error message %d", "some log", 1)
	assert.Equal(t, "upper error message 1", err.Error())
	assert.Equal(t, "some log: original error", err.LogError())
}

func Test_LogError(t *testing.T) {
	t.Run("without cause", func(t *testing.T) {
		err := errorx.New("some_code", "some_string")
		assert.Equal(t, "[some_code] some_string", err.LogError())
	})

	t.Run("with cause: go standard error", func(t *testing.T) {
		cause := fmt.Errorf("original error")

		err := errorx.Wrap(cause, "upper error", "upper error message")
		assert.Equal(t, "[upper error] upper error message: original error", err.LogError())
	})

	t.Run("with cause: rapor standard error", func(t *testing.T) {
		cause := errorx.New("some_code", "some_string")

		err := errorx.Wrap(cause, "upper error", "upper error message")
		assert.Equal(t, "[upper error] upper error message: [some_code] some_string", err.LogError())
	})

	t.Run("with cause using WrapWithLog: go standard error", func(t *testing.T) {
		cause := errorx.New("some_code", "some_string")

		err := errorx.WrapWithLog(cause, "upper error", "upper error message", "some_log_message")
		assert.Equal(t, "some_log_message: [some_code] some_string", err.LogError())
	})

	t.Run("with cause using WrapWithLog: rapor standard error", func(t *testing.T) {
		cause := errorx.New("some_code", "some_string")

		err := errorx.WrapWithLog(cause, "upper error", "upper error message", "some_log_message")
		assert.Equal(t, "some_log_message: [some_code] some_string", err.LogError())
	})
}

func Test_AddDetails(t *testing.T) {
	err := errorx.New(errorx.CodeInternal, "some error")
	err.AddDetails(&errorx.Details{
		Field:   "name",
		Message: "Name is empty",
	})
	assert.Equal(t, []*errorx.Details{
		{
			Field:   "name",
			Message: "Name is empty",
		},
	}, err.Details)
}

func Test_PrintStackTrace(t *testing.T) {
	assert.NotPanics(t, func() {
		errorx.New("some_code", "some_string").PrintStackTrace()
	})
}

func Test_ParseAndWrap(t *testing.T) {
	t.Run("parse fail, return wrapped error", func(t *testing.T) {
		err := fmt.Errorf("some error")

		gErr := errorx.ParseAndWrap(err, "This is some default error message")
		assert.Equal(t, errorx.CodeInternal, gErr.Code)
		assert.Equal(t, "This is some default error message", gErr.Message)
		assert.Equal(t, "[common.internal] This is some default error message: some error", gErr.LogError())
	})

	t.Run("parse success, return orig error", func(t *testing.T) {
		err := errorx.New("some.code", "some error message")

		gErr := errorx.ParseAndWrap(err, "This is some default error message")
		assert.Equal(t, "some.code", gErr.Code)
		assert.Equal(t, "some error message", gErr.Message)
		assert.Equal(t, err.LogError(), gErr.LogError())
	})
}
