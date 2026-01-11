package controllers

import (
	"net/http"
	"os"

	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pocketbase/pocketbase/core"
)

func AuthDocusign(e *core.RequestEvent) error {
	//response con json status ok
	token, err := GenerateDocuSignJWT()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate JWT",
		})
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "Docusign authenticated successfully",
		"token":   token,
	})
}
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func GenerateDocuSignJWT() (string, error) {
	privateKey, err := loadPrivateKey(".privateKey")
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"iss":   os.Getenv("DOCUSIGN_INTEGRATOR_KEY"),
		"sub":   os.Getenv("DOCUSIGN_USER_ID"),
		"aud":   "account-d.docusign.com",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Second * 6000).Unix(),
		"scope": "signature impersonation",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}
