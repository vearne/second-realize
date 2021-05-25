package rwlock

type Option func(*Options)

type Options struct {
	maxWaitWriteGoroutine int
	maxReadLocked         int
}

func WithMaxReadLocked(maxReadLocked int) Option {
	return func(o *Options) {
		o.maxReadLocked = maxReadLocked
	}
}

func WithMaxWaitWriteGoroutine(maxWaitWriteGoroutine int) Option {
	return func(o *Options) {
		o.maxWaitWriteGoroutine = maxWaitWriteGoroutine
	}
}
