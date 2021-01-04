package apiHelpers

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type tokenData struct {
	Email string
	Id    uint
}

func RandomKeyGenerator(strSize int, randType string) string {
	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	_, _ = rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func ImageDelete(path string) {
	os.Remove(path)
}

func ImageUpload(c *gin.Context, dirName string, prefix string) (string, bool) {
	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	defer file.Close()

	// to manage maximum size
	if handler.Size > 3690254 {
		return "", true
	}

	extension := filepath.Ext(handler.Filename)

	// assign image(path) to temp variable with prefix product name
	tempFile, err := ioutil.TempFile(dirName, prefix+"-*"+extension)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read image
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write image
	tempFile.Write(fileBytes)
	if len(strings.Split(tempFile.Name(), "/")) > 0 {
		return strings.Split(tempFile.Name(), "/")[3], false
	}
	return "", false
}

func GenerateAuthToken(email string, id uint) string {

	tokenData := &tokenData{
		Email: email,
		Id:    id,
	}

	fmt.Println("helper id", id)
	fmt.Println("before", tokenData)
	token, err := GenerateToken([]byte(os.Getenv("API_AUTH_KEY")), tokenData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("after", tokenData)
	return token
}

func GenerateToken(k []byte, userData interface{}) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["userData"] = userData
	fmt.Println("-=-=-**", userData)
	claims["exp"] = time.Now().Add(time.Hour * 8760).Unix()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(k)
	return tokenString, err
}

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		// Check if toke in correct format
		// ie Bearer: xx03xllasx
		b := "Bearer "
		if !strings.Contains(token, b) {
			c.JSON(403, gin.H{"message": "Your request is not authorized", "status": 403})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(403, gin.H{"message": "An authorization token was not supplied", "status": 403})
			c.Abort()
			return
		}

		// Validate token
		valid, err := ValidateToken(t[1], os.Getenv("API_AUTH_KEY"))
		if err != nil {
			c.JSON(403, gin.H{"message": "Invalid authorization token", "status": 403})
			c.Abort()
			return
		}

		// set userId Variable
		c.Set("userData", valid.Claims.(jwt.MapClaims)["userData"])
		c.Next()
	}
}

func ValidateToken(t string, k string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(k), nil
	})

	return token, err
}

func AuthUser(c *gin.Context) map[string]interface{} {
	userId := c.MustGet("userData")
	data := userId.(map[string]interface{})
	return data
}

func GetPort() string {
	port := os.Getenv("port")
	if port == "" {
		port = "8080" //localhost
	}
	return port
}
