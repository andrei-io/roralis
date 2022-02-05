package dic

import (
	"country/domain/entity"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/sarulabs/di"
	"github.com/spf13/viper"
)

func loadRSAKeys(ctn di.Container) (interface{}, error) {
	var err error

	privateBytes := []byte(viper.GetString("JWT_PRIVATE"))
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

	// *------

	publicBytes := []byte(viper.GetString("JWT_PUBLIC"))
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

	return entity.JWTSecret{
		SignKey:   privateKey,
		VerifyKey: publicKey,
	}, nil

}
