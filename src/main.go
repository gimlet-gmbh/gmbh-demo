package main

import (
	"errors"
	"os"
	"time"

	"github.com/abedick/gmbh"
	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	cabal := gmbh.NewComsModule()
	cabal.SetClient()
	cabal.SetServer()
	cabal.Route("test", handleTest)
	cabal.Route("two", handleTest2)
	cabal.Route("tkn", handleTkn)
	cabal.Start("demo")
}

func handleTest(req gmbh.Request, resp *gmbh.Responder) {
	resp.Result = "Hello from _cabal-generic, we received: " + req.Data1
}
func handleTest2(req gmbh.Request, resp *gmbh.Responder) {
	resp.Result = "Hello from _cabal-generic TEST2, we received: " + req.Data1
}

func handleTkn(req gmbh.Request, resp *gmbh.Responder) {

	tkn, err := generateToken(req.Data1, "guest")
	if err != nil {
		panic(err)
	}

	resp.Result = tkn

}

func generateToken(username string, admin string) (string, error) {
	tokenSecret := "gmbH"
	key := []byte(tokenSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	header := token.Header
	claims := token.Claims.(jwt.MapClaims)

	header["kid"] = os.Getenv("USERKID")
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	claims["id"] = username
	claims["usergroup"] = admin
	claims["iat"] = time.Now().Unix()
	claims["iss"] = "_cabal-authentication-module"

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("Could not generate valid token")
	}
	return tokenString, nil
}
