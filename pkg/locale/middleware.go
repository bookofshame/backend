package locale

import (
	"github.com/invopop/ctxi18n"
	"net/http"
)

// SetContext Middleware to set the locale in the context.
func SetContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Accept-Language")

		// supported language en and bn
		if lang != "en" && lang != "bn" {
			lang = "en"
		}

		ctx, _ := ctxi18n.WithLocale(r.Context(), lang)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
