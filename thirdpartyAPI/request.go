package thirdapi

import (
	"net/http"
)

// HTTPRequest 请求
type HTTPRequest struct {
	*http.Request
}
