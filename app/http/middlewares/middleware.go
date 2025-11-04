package middlewares

import "net/http"

// HTTPHandlerFunc 简写
type HTTPHandlerFunc func(http.ResponseWriter, *http.Request)
