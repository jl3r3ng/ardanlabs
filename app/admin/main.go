package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"gitlab.com/FireH24d/business/data/schema"
	"gitlab.com/FireH24d/foundation/database"

	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func main() {
	//keygen()
	//tokengen()
	migrate()
}
func migrate() {
	cfg := database.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       "0.0.0.0",
		Name:       "postgres",
		DisableTLS: true,
	}
	db, err := database.Open(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	if err := schema.Migrate(db); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("migrations complete")
	if err := schema.Seed(db); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("seed data complete")

}
func tokengen() {
	privatePEM, err := ioutil.ReadFile("/Users/ASUS/GolandProjects/class/private.pem")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		log.Fatal(err)
	}

	claims := struct {
		jwt.StandardClaims
		Roles []string `json:"roles"`
	}{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "service project",
			Subject:   "123456789",
			ExpiresAt: jwt.At(time.Now().Add(8760 * time.Hour)),
			IssuedAt:  jwt.Now(),
		},
		Roles: []string{"ADMIN"},
	}

	method := jwt.GetSigningMethod("RS256")
	tkn := jwt.NewWithClaims(method, claims)
	tkn.Header["kid"] = "54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"
	str, err := tkn.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("-----BEGIN TOKEN-----\n%s\n-----END TOKEN-----\n", str)
}

func keygen() {
	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		log.Fatal(err)
	}
	defer privateKeyFile.Close()

	privateBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(privateKeyFile, &privateBlock)
	if err != nil {
		log.Fatal(err)
	}

	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		log.Fatal(err)
	}
	defer publicKeyFile.Close()

	asn1Bytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
	publicBlock := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	err = pem.Encode(publicKeyFile, &publicBlock)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE")
}
