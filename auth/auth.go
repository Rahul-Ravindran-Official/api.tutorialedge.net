package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dgrijalva/jwt-go"
)

// Jwks struct for all keys
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys struct for Web Keys
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// TokenInfo struct contains all
// the payload information from the JWT
type TokenInfo struct {
	Iss   string   `json:"iss"`
	Sub   string   `json:"sub"`
	Iat   string   `json:"iat"`
	Exp   string   `json:"exp"`
	Azp   string   `json:"azp"`
	Aud   []string `json:"aud"`
	Scope string   `json:"scope"`
}

// Authenticate -
// Takes in a request and returns a true or false as to whether or not
// the incoming request is authenticated
func Authenticate(request events.APIGatewayProxyRequest) (bool, TokenInfo) {
	fmt.Println("Attempting to Authenticate Incoming Request...")

	var tokenInfo TokenInfo

	header := request.Headers["Authorization"]
	if header == "" {
		return false, TokenInfo{}
	}

	tokenString := strings.Split(string(header), " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Verify 'aud' claim
		aud := os.Getenv("API_AUDIENCE_ID")
		checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		if !checkAud {
			return token, errors.New("Invalid audience")
		}

		iss := "https://tutorialedge.eu.auth0.com/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer")
		}

		cert, err := getPemCert(token)
		if err != nil {
			panic(err.Error())
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		if err != nil {
			panic(err.Error())
		}

		return verifyKey, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return false, TokenInfo{}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenInfo.Sub = claims["sub"].(string)
		tokenInfo.Iss = claims["iss"].(string)
	} else {
		fmt.Println("Failed to parse token information properly...")
		return false, TokenInfo{}
	}

	return token.Valid, tokenInfo
}

func Auth(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Attempting to Authenticate Incoming Request...")
		var tokenInfo TokenInfo

		header := r.Header["Authorization"]
		if header == "" {
			fmt.Fprintf(w, "Not Authorized")
			return
		}

		tokenString := strings.Split(string(header), " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// Verify 'aud' claim
			aud := os.Getenv("API_AUDIENCE_ID")
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return token, errors.New("Invalid audience")
			}

			iss := "https://tutorialedge.eu.auth0.com/"
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer")
			}

			cert, err := getPemCert(token)
			if err != nil {
				panic(err.Error())
			}

			verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				panic(err.Error())
			}

			return verifyKey, nil
		})

		if err != nil {
			fmt.Println(err.Error())
			fmt.Fprintf(w, err.Error())
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			tokenInfo.Sub = claims["sub"].(string)
			tokenInfo.Iss = claims["iss"].(string)
		} else {
			fmt.Println("Failed to parse token information properly...")
			fmt.Fprintf(w, "Claims are not valid")
		}

		if token.Valid {
			endpoint(w, r)
		}

	})
}

// getPemCert retrieves the most up-to-date JWK.json file for
// tutorialedge so that we can validate the JWT using the correct certs
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://tutorialedge.eu.auth0.com/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}

// UnauthorizedResponse returns a pre-defined
// unauthorized APIGatewayProxyResponse
func UnauthorizedResponse() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "Not Authorized",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 503,
	}
}
