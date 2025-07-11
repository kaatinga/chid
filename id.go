package shared

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

// getIDType is a type constraint for supported integer ID types.
type getIDType interface {
	int16 | int32 | int64
}

// ErrUnsupportedIDType is returned when an unsupported ID type is used.
var ErrUnsupportedIDType = errors.New("unsupported id type")

// GetID extracts and validates an integer ID from an HTTP request.
// It supports int16, int32, and int64 types, and can read from URL params or form values.
// Options can be provided to customize the key, bit size, and source.
func GetID[T getIDType](r *http.Request, optFns ...getIdOption) (T, error) {
	var opts = defaultGetIdOptions()
	for _, o := range optFns {
		o(&opts)
	}

	var bitSize int
	var zero T
	switch any(zero).(type) {
	case int64:
		bitSize = 64
	case int32:
		bitSize = 32
	case int16:
		bitSize = 16
	default:
		return 0, ErrUnsupportedIDType
	}

	var id int64
	var err error
	if opts.readFormValue {
		id, err = strconv.ParseInt(strings.TrimSpace(r.FormValue(opts.idKey)), 10, bitSize)
	} else {
		id, err = strconv.ParseInt(strings.TrimSpace(chi.URLParam(r, opts.idKey)), 10, bitSize)
	}
	if err != nil {
		source := "url parameter"
		if opts.readFormValue {
			source = "form value"
		}
		return 0, fmt.Errorf("unable to parse '%s' %s: %w", opts.idKey, source, err)
	}
	if id == 0 {
		return 0, errors.New("'" + opts.idKey + "' == 0")
	}
	if id <= 0 {
		return 0, errors.New("'" + opts.idKey + "' <= 0")
	}

	return T(id), nil
}
