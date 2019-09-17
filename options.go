package retry

import (
	"time"
)

type Options struct {
	attempts          uint
	delay             time.Duration
	delayStrategyFunc DelayStrategyFunc
}

func defaultOptions() *Options {
	return &Options{
		attempts:          3,
		delayStrategyFunc: BackOff,
		delay:             100 * time.Millisecond,
	}
}

type Option func(*Options)

type DelayStrategyFunc func(n uint, o *Options) time.Duration

func BackOff(n uint, o *Options) time.Duration {
	return o.delay * (1 << n)
}

func Fixed(_ uint, o *Options) time.Duration {
	return o.delay
}

func DelayStrategy(delayStrategyFunc DelayStrategyFunc) Option {
	return func(o *Options) {
		o.delayStrategyFunc = delayStrategyFunc
	}
}

func Delay(delay time.Duration) Option {
	return func(o *Options) {
		o.delay = delay
	}
}

func Attempts(attempts uint) Option {
	return func(o *Options) {
		o.attempts = attempts
	}
}
