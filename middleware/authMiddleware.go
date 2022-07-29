package middleware

import (
	"github.com/gin-goinc/gin"
)

func Authentication() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := cRequest.header.get("token")
		if clientToken == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Auth header provided")})
			c.Abort()
			return
		}

		claims, err = helper.ValidateToken(clientToken)

		if er != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Lirst_name)
		c.Set("uid", claims.Uid)

		c.Next()
	}
}