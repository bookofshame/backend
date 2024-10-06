package locale

import (
	"net/http"
)

type Locale interface {
	SetContext(next http.Handler) http.Handler
}
