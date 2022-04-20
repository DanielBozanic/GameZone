package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SECRET_KEY = "7kRh7cgjun9S83Hu06HhqhB8sGYkZKHrZ7CRkpQJHfOzXTllQPcIWCSn3IcUccq"

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("Token is missing")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("Incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Bad signed method received!")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, errors.New("Bad JWT token!")
	}

	return token, nil
}

func decodeJwtToken(c *gin.Context) (*jwt.Token) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return nil
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return nil
	}
	return token
}

func AuthenticationRequired(c *gin.Context) {
	decodeJwtToken(c)
	c.Next()
}

func AuthorizationRequired(roles []string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
        token := decodeJwtToken(c)

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Unable to parse claims!"})
			return
		}

		role := claims["sub"].(map[string]interface{})["role"].(string)

		valid := false
		for _, r := range roles {
			if r == role {
				valid = true
				break;
			}
		}

		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized!"})
			return
		}
		c.Next()
    }

    return gin.HandlerFunc(fn)
}