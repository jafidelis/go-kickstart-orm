package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/go-kickstart-orm/model/entity"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func GenerateJWT(user entity.User) string {
	loadKeyfile()
	claims := entity.Claim{
		User: entity.User{
			FirstName: user.FirstName,
			Login:     user.Login,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Login",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Não foi possível assinar o token: ", err)
	}

	return result
}

func OAuthFilter(inner http.Handler) http.Handler {
	loadKeyfile()
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token, err := request.ParseFromRequestWithClaims(req, request.OAuth2Extractor,
			&entity.Claim{}, func(token *jwt.Token) (interface{}, error) {
				return publicKey, nil
			})

		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					fmt.Fprintln(w, "Seu token está expirado")
					return
				case jwt.ValidationErrorSignatureInvalid:
					fmt.Fprintln(w, "A assinatura do token não coincide")
					return
				default:
					fmt.Fprintln(w, "Seu token não é válido")
					return
				}
			default:
				fmt.Fprintln(w, "Seu token não é válido", err)
				return
			}
		}

		if token.Valid {
			// w.WriteHeader(http.StatusAccepted)
			// fmt.Fprintln(w, "Bem Vindo")
			inner.ServeHTTP(w, req)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Token inválido")
		}
	})
}

func loadKeyfile() {
	privateBytes, err := ioutil.ReadFile("./auth/private.rsa")
	if err != nil {
		log.Fatal("Não foi possível ler a chave privada")
	}

	publicBytes, err := ioutil.ReadFile("./auth/public.rsa.pub")
	if err != nil {
		log.Fatal("Não foi possível ler a chave publica")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Não foi possível fazer o parse da chave privada")
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Não foi possível fazer o parse da chave publica")
	}
}
