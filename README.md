# chid

A small Go package for extracting and validating integer IDs from HTTP requests, designed for use with the [chi/v5](https://github.com/go-chi/chi) router, but can be adapted for other frameworks.

## Features
- Extracts integer IDs from HTTP requests (URL params or form values)
- Supports int16, int32, and int64 types
- Customizable key, bit size, and source via options
- Clear error handling for invalid or missing IDs

## Installation

```
go get github.com/kaatinga/chid
```

## Usage

```go
import (
    "net/http"
    "github.com/kaatinga/chid/shared"
)

// Example handler using chi router
func handler(w http.ResponseWriter, r *http.Request) {
    id, err := shared.GetID[int64](r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Use id (int64)
}
```

### Customization
You can customize how the ID is extracted using options:

- `shared.WithIDKey(key string)`: Use a custom key instead of the default "id"
- `shared.WithBitSize(size int)`: Specify bit size (16, 32, or 64)
- `shared.WithFormValue()`: Extract from form values instead of URL params

**Example:**
```go
id, err := shared.GetID[int32](r, shared.WithIDKey("user_id"), shared.WithBitSize(32), shared.WithFormValue())
```

## Error Handling

`GetID` returns descriptive errors for:
- Unsupported ID type
- Unable to parse the ID value
- ID is zero or below zero

## License
MIT