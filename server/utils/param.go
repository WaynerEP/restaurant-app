package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	ErrorTokenWasNotFound = errors.New("no autorizado: no se encontró el token")
)

// GetIdFromParam extracts id of the c.Param
func GetIdFromParam(c *gin.Context) (uint, error) {
	uuid, err := strconv.Atoi(c.Param("id"))
	return uint(uuid), err
}

func ExtractTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		return "", ErrorTokenWasNotFound
	}
	return token, nil
}
