package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// iss (issuer): Issuer of the JWT
// sub (subject): Subject of the JWT (the user)
// aud (audience): Recipient for which the JWT is intended
// exp (expiration time): Time after which the JWT expires
// nbf (not before time): Time before which the JWT must not be accepted for processing
// iat (issued at time): Time at which the JWT was issued; can be used to determine age of the JWT
// jti (JWT ID): Unique identifier; can be used to prevent the JWT from being replayed (allows a token to be used only once)

// JWT token given as string
var tkn = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY5NzYyNDMyNiwiaWF0IjoxNjk3NjIxMzI2fQ.ZOYjfCChpN3NGmeCeHgTTw32_hwHtwHGbLaVICNjUdtSUc2v5qNaGnLK63UxcoAp-32kjo6SMumQ80Spvj2qpxQkg47xLvjqFhBUP9Fz1URag9TzDuyNT5zrIi74gzDDvAur-hzG0LANMgm0e1k4vchIwuckHE8uEvT4zsa_My1JYPvnjKex-pOU-M6TNryC8eWVfzZ83rcVghLaa6o5pxBfw2pOj0anIyl_cA-BjGykXZCYB0q6jPJRKcGGffIQ0DpQgwBvpWxrcZXurIWN13hPI3inxaZxC7ombEmYEsk_I5wg467vxZriuvO0iWlxGyQ2XHAZm51zRWUlOhRIPA`

func main() {
	// CreateToken()
	ValidateToken()

}

func ValidateToken() {
	// Reads the public key from pubkey.pem file
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		// If there's an error reading the file, print an error message and stop execution
		log.Fatalln("not able to read pem file")
	}

	// Parse the read public key to RSA public key format
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	if err != nil {
		// If there's an error parsing the public key, log the error and stop execution
		log.Fatalln(err)
	}
	var c jwt.RegisteredClaims
	// Parsing the JWT token with the claims
	token, err := jwt.ParseWithClaims(tkn, &c, func(token *jwt.Token) (interface{}, error) {
		// Provides the public key for validating the JWT token
		return publicKey, nil
	})

	if err != nil {
		// If error while parsing the token, print the error and exit
		log.Println("parsing token", err)
		return

	}
	if !token.Valid {
		// If the token is not valid, log the error and exit
		log.Println("invalid token")
		return
	}

	// Print the claims from the token
	fmt.Printf("%+v", c)

}

func CreateToken() {
	privatePem, err := os.ReadFile("private.pem")
	if err != nil {
		log.Fatalln(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePem)
	if err != nil {
		log.Fatalln(err)
	}

	c := jwt.RegisteredClaims{
		Issuer:    "api project",
		Subject:   "101",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(50 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	//NewWithClaims creates a new Token with the specified signing method and claims
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	str, err := tkn.SignedString(privateKey)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(str)

	//eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY5NzYyMzA3MSwiaWF0IjoxNjk3NjIwMDcxfQ.B5TONrhy8jx7zlC2v4EwoiV7TVKVKan_pvJdvW1Xom-NZqo8yynZ3WBQe7MGblPTJ7BUQRaNwjs8E9QOa12Fj57iWovPWzKUN3r58ypgLR8T9jd7xYi_b5kLMXnhaFqfYGzLYz8dTUW2h7ZMcvlOsjWDz62_s0UT0XHlmPS9pLPLEgjDsb6n3SEFuWN2v9OtCV69hlTp_0MOyheNhcK_8k5ljZueFSXwTXAH4PH5p0k9h72gAXlXx5tg8FkGGSFoj2WDqFQV-YXaNkxcxpF84tQYtMqW6rp3B1lECRQmyCbRo7ato22jXBEqqfhmjYI_Xd1GET8GraIkVX5NLvNF8
}
