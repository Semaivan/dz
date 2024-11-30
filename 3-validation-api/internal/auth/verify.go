package auth

import (
	configmail "MyDz/3-validation-api/configs"
	res "MyDz/3-validation-api/pkg"
	"MyDz/3-validation-api/pkg/req"
	sendmail "MyDz/3-validation-api/pkg/send"

	"fmt"

	"encoding/base64"
	//"encoding/json"

	//"io"
	//"encoding/json"
	//"fmt"
	"log"
	"net/http"

	//"net/mail"

	//"reflect"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandlerDeps struct {
	*configmail.Config
}

type AuthHandler struct {
	*configmail.Config
}

var MailToken string
var MailTo string
var TokenPass string

// GenerateToken returns a unique token based on the provided email string
func GenerateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Hash to store:", string(hash))
	return base64.StdEncoding.EncodeToString(hash)
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	strGet := "GET /auth/verify/" + MailToken
	router.HandleFunc("POST /auth/send", handler.Send())
	router.HandleFunc(strGet, handler.Verify())
	router.HandleFunc("GET /auth/register", handler.Register())
}

func (handler *AuthHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data := VerifyResponse{Http: "/auth/verify" + MailToken}
		res.Json(w, data, 200)
	}

}
func (handler *AuthHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		MailTo = body.Email
		TokenPass = body.Password
		MailToken = GenerateToken(body.Email)
		//b, err := json.Marshal(body)
		//fmt.Println(MailToken)
		sendmail.СreateMail(MailTo, MailToken)

	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
	}
}
