package chid

// getIdOption is a function that modifies getIdOptions.
type getIdOption func(*getIdOptions)

// getIdOptions holds configuration for ID extraction.
type getIdOptions struct {
	idKey         string
	bitSize       int
	readFormValue bool
}

// defaultGetIdOptions returns the default options for ID extraction.
func defaultGetIdOptions() getIdOptions {
	return getIdOptions{
		idKey:   "id",
		bitSize: 64,
	}
}

// WithIDKey sets a custom key for ID extraction.
func WithIDKey(key string) getIdOption {
	return func(o *getIdOptions) {
		o.idKey = key
	}
}

// WithBitSize sets a custom bit size for ID parsing.
func WithBitSize(size int) getIdOption {
	return func(o *getIdOptions) {
		o.bitSize = size
	}
}

// WithFormValue configures extraction from form values instead of URL params.
func WithFormValue() getIdOption {
	return func(o *getIdOptions) {
		o.readFormValue = true
	}
}
