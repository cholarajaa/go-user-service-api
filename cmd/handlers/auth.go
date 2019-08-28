package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type jwtToken struct {
	Token  string `json:"token"`
	Expiry int    `json:"expiry"`
}

const expiryMinutes = 20

// Auth takes input for Handlers
// such as logger and other required packages
// it is like injecting dependency for API Handler
type Auth struct {
	log *log.Logger
}

// Token generates auth token
func (a *Auth) Token(w http.ResponseWriter, r *http.Request) {
	defer a.log.Println("Auth token generated")
	var sig string
	var err error
	// generate jwt token with expiry date
	method := jwt.GetSigningMethod("RS256")
	keyContents, err := ioutil.ReadFile("./private.pem")
	if err != nil {
		a.log.Println("Error reading file")
	}
	privatekey, err := jwt.ParseRSAPrivateKeyFromPEM(keyContents)
	if err != nil {
		a.log.Println("Error creating private key")
	}
	if sig, err = method.Sign("somethistokenyes", privatekey); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	a.log.Println(sig)
	token := jwtToken{
		sig,
		expiryMinutes,
	}
	json.NewEncoder(w).Encode(token)
}
