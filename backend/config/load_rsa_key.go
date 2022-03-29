package config

import (
	"backend/roralis/auth"

	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Loads RSA keys from viper.
// This function is mostly copied from the internet because working with crypto keys is golang is a pain
func loadRSAKeys(jwtPrivate, jwtPublic string) (*auth.JWTSecret, error) {

	var err error

	// start decoding

	privateBytes := []byte(jwtPrivate)
	privatePem, _ := pem.Decode(privateBytes)

	if privatePem.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("Private key is not the expected type")
	}

	privPemBytes := privatePem.Bytes

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		return nil, errors.New("Unable to parse RSA private key: " + err.Error())
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("Unable to parse RSA private key")
	}

	// start decoding public key

	publicBytes := []byte(jwtPublic)
	publicPem, _ := pem.Decode(publicBytes)

	if publicPem.Type != "PUBLIC KEY" {
		return nil, errors.New("Public key is not the expected type: " + publicPem.Type)
	}

	publicPemBytes := publicPem.Bytes

	if parsedKey, err = x509.ParsePKIXPublicKey(publicPemBytes); err != nil {
		return nil, errors.New("Unable to parse RSA public key: " + err.Error())
	}

	var publicKey *rsa.PublicKey
	publicKey, ok = parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Unable to parse RSA public key")
	}

	privateKey.PublicKey = *publicKey

	return &auth.JWTSecret{
		SignKey:   privateKey,
		VerifyKey: publicKey,
	}, nil

}
