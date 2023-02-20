package tracer

import (
	"context"
	"testing"

	logger "bitbucket.org/ayopop/ct-logger/logger"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var tLogger *logger.Logger

	tracerImpl := New(false, "test", "test", tLogger)
	assert.NotNil(t, tracerImpl)

}

func TestShutdown(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				Shutdown()
			},
		)
	}
}

func TestStartSpan(t *testing.T) {
	tCtx := context.Background()
	ctx, span := StartSpan(tCtx, "test")
	assert.NotNil(t, ctx)
	assert.NotNil(t, span)
}
