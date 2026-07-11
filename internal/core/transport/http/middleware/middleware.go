package core_http_middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(
	h http.Handler,
	m ...Middleware,
) http.Handler {
	if len(m) == 0 {
		return h
	}

	// в цикле идем с конца, так как последним будет Trace и накинется он самым первым в CreateUser => порядок вызово middlewares не нарушен
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}
