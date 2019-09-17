package retry_test

import (
	"errors"
	"testing"
	"time"

	"github.com/jattschneider/retry"
	"github.com/stretchr/testify/assert"
)

func TestRetryWithFixedDelay(t *testing.T) {
	attempts := uint(0)
	err := retry.With(
		func() error {
			attempts++
			return errors.New("error while attempting.")
		},
		retry.Delay(time.Nanosecond),
		retry.DelayStrategy(retry.Fixed),
	)
	assert.Error(t, err)

	expectedErrorFormat := `Errors:
#1: error while attempting.
#2: error while attempting.
#3: error while attempting.`
	assert.Equal(t, expectedErrorFormat, err.Error(), "retry error format")
	assert.Equal(t, uint(3), attempts, "right count of retry")
}

func TestRetryWithBackoffDelay(t *testing.T) {
	attempts := uint(0)
	err := retry.With(
		func() error {
			attempts++
			return errors.New("error while attempting.")
		},
		retry.Delay(time.Nanosecond),
		retry.DelayStrategy(retry.BackOff),
		retry.Attempts(5),
	)
	assert.Error(t, err)

	expectedErrorFormat := `Errors:
#1: error while attempting.
#2: error while attempting.
#3: error while attempting.
#4: error while attempting.
#5: error while attempting.`
	assert.Equal(t, expectedErrorFormat, err.Error(), "retry error format")
	assert.Equal(t, uint(5), attempts, "right count of retry")
}
