package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = os.Get("JWT_Token") then add JWT_Token to env var
var mySigningKey = []byte("secretPhaseYouWant")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Alpha Bravo"
	// Token will be expired in 30 minutes
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	} else {
		return tokenString, nil
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		// Print valid token:
		// fmt.Fprintf(w, validToken)

		// Create new request to server
		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)

		// Add token to header
		req.Header.Set("Token", validToken)
		res, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err.Error())
		}
		fmt.Fprintf(w, string(body))
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Simple client")
	// tokenString, err := GenerateJWT()
	// if err != nil {
	// 	fmt.Println("Can't generate token")
	// } else {
	// 	fmt.Println(tokenString)
	// }
	handleRequests()
}
