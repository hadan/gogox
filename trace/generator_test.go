package trace_test

import (
	"testing"

	"github.com/hadan/gogox/trace"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	assert.NotNil(t, trace.New())
}
