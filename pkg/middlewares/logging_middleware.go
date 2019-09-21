package middlewares

import negronilogrus "github.com/meatballhat/negroni-logrus"

// LoggingMiddleware ...
func LoggingMiddleware() *negronilogrus.Middleware {
	return negronilogrus.NewMiddleware()
}
