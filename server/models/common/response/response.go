package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	CreatedMsg    = "Registro creado correctamente."
	BadRequestMsg = "Solicitud inválida. Verifique, los parámetros y valores."
	SUCCESS       = "Operación exitosa"
	ERROR         = "Operación fallida"
)

type ValidationError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ValidationErrors map[string]ValidationError

// Response is the structure for the API response
type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// Result sends the API response with the specified code, data, and message
func Result(code int, data interface{}, msg string, c *gin.Context) {
	// Start time
	c.JSON(code, Response{
		data,
		msg,
	})
}

// OK sends a success response with default messages
func OK(c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, SUCCESS, c)
}

// OkWithMessage sends a success response with a custom message
func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, message, c)
}

// OkWithData sends a success response with specified data and a default message
func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, SUCCESS, c)
}

// OkWithDetailed sends a success response with specified data and message
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, data, message, c)
}

// Created sends a success response with specified data and a default message
func Created(c *gin.Context) {
	Result(http.StatusCreated, map[string]interface{}{}, CreatedMsg, c)
}

// CreatedWithMessage sends a success response with specified data and a message
func CreatedWithMessage(msg string, c *gin.Context) {
	Result(http.StatusCreated, map[string]interface{}{}, msg, c)
}

// CreatedWithDetail sends a success response with specified data and a custom message
func CreatedWithDetail(data interface{}, message string, c *gin.Context) {
	Result(http.StatusCreated, data, message, c)
}

// FailWithValidationErrors sends a error response with inputs validations and message
func FailWithValidationErrors(data interface{}, c *gin.Context) {
	Result(http.StatusUnprocessableEntity, data, ERROR, c)
}

// FailClientError sends a failure response with default messages
func FailClientError(message string, c *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, message, c)
}

// FailWithMessage sends a failure response with a custom message
func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusInternalServerError, map[string]interface{}{}, message, c)
}

// FailWithDetailed sends a failure response with specified data and message
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusInternalServerError, data, message, c)
}

// FailBound sends a failure response with specified data and message
func FailBound(msg string, c *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, msg, c)
}
