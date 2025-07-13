package chid

// getIdOption is a function that modifies getIdOptions.
type getIdOption func(*getIdOptions)

// getIdOptions holds configuration for ID extraction.
type getIdOptions struct {
	idKey         string
	readFormValue bool
}

// defaultGetIdOptions returns the default options for ID extraction.
func defaultGetIdOptions() getIdOptions {
	return getIdOptions{
		idKey: "id",
	}
}

// WithIDKey sets a custom key for ID extraction.
func WithIDKey(key string) getIdOption {
	return func(o *getIdOptions) {
		o.idKey = key
	}
}

// WithFormOrQuery configures extraction from form/query values instead of URL parameters.
func WithFormOrQuery() getIdOption {
	return func(o *getIdOptions) {
		o.readFormValue = true
	}
}
