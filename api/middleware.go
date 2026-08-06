package api

import (
	"net/http"
	"simplebank/token"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationType       = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader(authorizationHeaderKey)
		if authHeader == "" || len(authHeader) == 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is not provided"})
			return
		}

		field := strings.Fields(authHeader)
		if len(field) < 2 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		authType := strings.ToLower(field[0])
		if authorizationType != authType {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unsupported authorization type"})
			return
		}

		accessToken := field[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		context.Set(authorizationPayloadKey, payload)

		context.Next()
	}
}
