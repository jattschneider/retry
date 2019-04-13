package retry_test

import (
	"errors"
	"testing"

	"github.com/jattschneider/retry"
)

func TestRetrySuccessful(t *testing.T) {
	callCount := 0
	err := retry.With(func(attempts int) error {
		t.Logf("Attempt: %d\n", attempts)
		callCount++
		return nil
	}, 5)
	if err != nil {
		t.Error(err)
	}
	if callCount != 1 {
		t.Fail()
	}
}

func TestRetryFailed(t *testing.T) {
	errWentWrong := errors.New("something went wrong")
	callCount := 0
	err := retry.With(func(attempts int) error {
		t.Logf("Attempt: %d\n", attempts)
		callCount++
		return errWentWrong
	}, 5)
	if err != errWentWrong {
		t.Fail()
	}
	if callCount != 6 {
		t.Fail()
	}
}

func TestMaxRetriesLimit(t *testing.T) {
	errNope := errors.New("nope")
	err := retry.With(func(attempts int) error {
		t.Logf("Attempt: %d\n", attempts)
		return errNope
	}, 3)
	if err == nil {
		t.Fail()
	}
	if err != errNope {
		t.Fail()
	}
}
