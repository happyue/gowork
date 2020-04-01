package utils

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

// IsDomain checks the specified string is domain name.
func IsDomain(s string) bool {
	return !IsIP(s) && "localhost" != s
}

// IsIP checks the specified string is IP.
func IsIP(s string) bool {
	return nil != net.ParseIP(s)
}

// GetRemoteAddr returns remote address of the context.
func GetRemoteAddr(c *gin.Context) string {
	ret := c.GetHeader("X-forwarded-for")
	ret = strings.TrimSpace(ret)
	if "" == ret {
		ret = c.GetHeader("X-Real-IP")
	}
	ret = strings.TrimSpace(ret)
	if "" == ret {
		return c.Request.RemoteAddr
	}

	return strings.Split(ret, ",")[0]
}
