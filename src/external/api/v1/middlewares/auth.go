package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hcsouza/fiap-tech-fast-food/src/external/api/infra/config"
	cognitoJwtVerify "github.com/jhosan7/cognito-jwt-verify"
)

func CheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		tokenParts := strings.Split(authorizationHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Malformed token"})
			return
		}

		accessToken := tokenParts[1]
		if !validateAccessToken(accessToken) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			return
		}
		c.Next()
	}
}

func validateAccessToken(accessToken string) bool {
	cognitoConfig := cognitoJwtVerify.Config{
		UserPoolId: config.GetApiCfg().AuthConfig.UserPoolId,
		ClientId:   config.GetApiCfg().AuthConfig.ClientId,
		TokenUse:   config.GetApiCfg().AuthConfig.TokenUse,
	}

	verify, err := cognitoJwtVerify.Create(cognitoConfig)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = verify.Verify(accessToken)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return false
	}
	return true
}
