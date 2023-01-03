package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var Secret = []byte(`JwtSecret`)

var User map[string]string = map[string]string{
	"test1111@gmail.com": "qaz123wsx!@#",
	"test2222@gmail.com": "zse!@#qsc!@#",
}

func main() {
	h := new(Handler)
	r := gin.Default()
	api := r.Group("/api/v1")
	api.POST("/login", h.login)
	api.GET("/ping", middlewareVerifyJWT(), h.ping)

	r.Run(":8000")
}

type Handler struct{}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) login(c *gin.Context) {
	var l Login
	if err := c.BindJSON(&l); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "missing email or password"})
		return
	}

	// verify user
	pwd, exist := User[l.Email]
	if !exist {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if pwd != l.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "password was wrong"})
		return
	}

	jwt, err := generateJWTToken(l.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": jwt})
}

func (h *Handler) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

type JWTClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func generateJWTToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Minute)),
		},
	})
	return token.SignedString(Secret)
}

func middlewareVerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing authorization"})
			return
		}

		bearer := strings.Split(auth, " ")
		if bearer[0] != "Bearer" || len(bearer) != 2 || len(bearer[1]) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid authorization"})
			return
		}

		tokenDecode, err := jwt.ParseWithClaims(bearer[1], &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return Secret, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		_, ok := tokenDecode.Claims.(*JWTClaims)
		if !ok || !tokenDecode.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		c.Next()
	}
}
