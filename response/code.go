package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ok sends a 200 OK success response with the provided data.
// This is the standard response for successful GET, PUT, and PATCH requests
// that return data to the client.
func Ok[T any](ctx *gin.Context, data T) {
	Success(ctx, http.StatusOK, data)
}

// Created sends a 201 Created success response with the provided data.
// This is the standard response for successful POST requests that create
// new resources and return the created resource data.
func Created[T any](ctx *gin.Context, data T) {
	Success(ctx, http.StatusCreated, data)
}

// NoContent sends a 204 No Content success response.
// This is the standard response for successful DELETE requests or other
// operations that complete successfully but return no data.
func NoContent(ctx *gin.Context) {
	ctx.String(http.StatusNoContent, "")
}

// BadRequest sends a 400 Bad Request error response.
// This is used when the client sends invalid or malformed request data
// that cannot be processed by the server.
func BadRequest(ctx *gin.Context, data error) {
	Error(ctx, http.StatusBadRequest, data)
}

// Unauthorized sends a 401 Unauthorized error response.
// This is used when the client lacks valid authentication credentials
// required to access the requested resource.
func Unauthorized(ctx *gin.Context, data error) {
	Error(ctx, http.StatusUnauthorized, data)
}

// Forbidden sends a 403 Forbidden error response.
// This is used when the client has valid credentials but lacks permission
// to access the requested resource.
func Forbidden(ctx *gin.Context, data error) {
	Error(ctx, http.StatusForbidden, data)
}

// NotFound sends a 404 Not Found error response.
// This is used when the requested resource does not exist or cannot be
// found on the server.
func NotFound(ctx *gin.Context, data error) {
	Error(ctx, http.StatusNotFound, data)
}

// MethodNotAllowed sends a 405 Method Not Allowed error response.
// This is used when the HTTP method used in the request is not allowed
// for the specified resource.
func MethodNotAllowed(ctx *gin.Context, data error) {
	Error(ctx, http.StatusMethodNotAllowed, data)
}

// InternalServerError sends a 500 Internal Server Error response.
// This is used when the server encounters an unexpected error that
// prevents it from fulfilling the request.
func InternalServerError(ctx *gin.Context, data error) {
	Error(ctx, http.StatusInternalServerError, data)
}

// BadGateway sends a 502 Bad Gateway error response.
// This is used when the server, acting as a gateway or proxy, receives
// an invalid response from the upstream server.
func BadGateway(ctx *gin.Context, data error) {
	Error(ctx, http.StatusBadGateway, data)
}

// ServiceUnavailable sends a 503 Service Unavailable error response.
// This is used when the server is temporarily unable to handle the request
// due to maintenance or overload conditions.
func ServiceUnavailable(ctx *gin.Context, data error) {
	Error(ctx, http.StatusServiceUnavailable, data)
}
