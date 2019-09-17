package retry

import (
	"fmt"
	"strings"
	"time"
)

type RetryableFunc func() error

func With(retryableFunc RetryableFunc, opts ...Option) error {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	errors := Errors{}
	n := uint(0)
	for n < options.attempts {
		err := retryableFunc()
		if err != nil {
			errors = append(errors, err)
			if n == options.attempts-1 {
				break
			}
			d := options.delayStrategyFunc(n, options)
			time.Sleep(d)
		} else {
			return nil
		}
		n++
	}
	return errors
}

type Errors []error

func (e Errors) Error() string {
	sb := strings.Builder{}
	sb.WriteString("Errors:")
	for i, l := range e {
		sb.WriteString(fmt.Sprintf("\n#%d: %s", i+1, l.Error()))
	}
	return sb.String()
}
