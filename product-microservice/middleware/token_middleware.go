package middleware

import (
	"errors"
	"net/http"
	"product/dto"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func GetUserData(c *gin.Context) dto.UserData {
	var userData dto.UserData
	token := decodeJwtToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	sub := claims["sub"].(map[string]interface{})
	userData.Id = sub["user_id"].(int)
	userData.Role = sub["role"].(string)
	return userData
}

func AuthenticationRequired(c *gin.Context) {
	token := decodeJwtToken(c)
	if token == nil {
		return
	}
	c.Next()
}

func AuthorizationRequired(roles []string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
        token := decodeJwtToken(c)

		if token == nil {
			return
		}

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