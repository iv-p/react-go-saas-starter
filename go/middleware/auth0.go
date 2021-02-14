package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// Auth0Authenticator provides middleware that auths using auth0 / authorization header JWT
type Auth0Authenticator struct {
	Domain string
}

// Jwks = json web key set
type Jwks struct {
	Keys []JSONWebKey `json:"keys"`
}

// JSONWebKey = one key
type JSONWebKey struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// NewAuth0Authenticator with the provided params. Domain needs to include
// `https://` and a trailing slash
func NewAuth0Authenticator(domain string) Auth0Authenticator {
	return Auth0Authenticator{Domain: domain}
}

func (a Auth0Authenticator) getPemCert(token *jwt.Token) (interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s.well-known/jwks.json", a.Domain))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return nil, err
	}

	var cert []byte
	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = []byte("-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----")
			break
		}
	}
	if len(cert) == 0 {
		return nil, errors.New("unable to find appropriate key")
	}

	return jwt.ParseRSAPublicKeyFromPEM(cert)
}

// Protect from unauthorized invocations
func (a Auth0Authenticator) Protect(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, a.getPemCert)
		if err != nil {
			http.Error(w, "invalid token", 401)
			return
		}

		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(a.Domain, false)
		if !checkIss {
			http.Error(w, "invalid issuer", 401)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
