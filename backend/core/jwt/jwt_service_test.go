package jwt_test

import (
	"backend/roralis/config"
	"backend/roralis/core/jwt"
	"testing"
)

// These are dummy keys - NOT used in production, but for tests
const (
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`
	publicKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA2FU6M1YK7J/NUZSOZ6Gd
MSKVO9M9bSBRMGsnWJ2Gtc8MnAO4ADV03b5RUdlzGPKDVAfzgposqf5rZyo8IocQ
uoQAjYEy8N6wzXXlxqJ34ssomdMgDS0MlEWS5DmkRo4uxwwjdUZZBJ2K/0X/iA/r
gRo6Vo9BSpsU3c2ErCoH5Ezj5jMKePWu7FQ3GK9Mep0JzRQt4qIMAG7V6Zsgtweg
9up5IoX600m6BbxGMSTOfOEGCWL32O1yElrDMn6hDmX3h+77ea96xwdMKxFyWRIt
qMKd0tAf+3yN/XG3ICiDBFeflB+IcKuKXmAQWUniO/zBrsAs+J6e1VOO04Npzk7b
spCIgVaJU0p3mHkvO0/qsqlJL9r1CcBSrY98B22vuwZqCBdufrZtwYGh+IgZFUxs
aAp4PXhSa7PWBeOpGLe1BbHoOK98IBbibDkzncS3et97tepBOuTmRCvq3rh2+D3o
aLDSLO6SF1Ek078y0mtt0wS/R56rIcspGJ/8+bU/mlJ+uw6KNMCdu270JRtZ5bAO
qvEkSMnD3ZeYoJcOWWVWkZ823q8/zHG/g7WyHb4UGhtsG0z5yGfMblAvmOATsT+m
HiEGbTPzuGS2vmsisyeICgU3Hxt3J9Aq1EFjpua+iQuDCB119s4PeKc4kSwrtPAl
9wO5oQVc4jKfBLND5m1rs7kCAwEAAQ==
-----END PUBLIC KEY-----`
)

func TestJWTService(t *testing.T) {
	secret, err := config.LoadRSAKeys(privateKey, publicKey)
	if err != nil {
		t.Errorf("Error creating jwt secrets: %v", err)
	}
	jwtService := jwt.NewJWTService(secret)

	claims := &jwt.JWTClaims{
		ID:       1,
		Name:     "Test",
		Verified: true,
		Role:     10,
	}
	token, err := jwtService.NewJWT(claims)
	if err != nil {
		t.Errorf("Error creating jwt: %v", err)
	}

	rawClaims, err := jwtService.VerifyJWT(&token)
	if err != nil {
		t.Errorf("Error verifying jwt: %v", err)
	}
	if *rawClaims != *claims {
		t.Errorf("Got different claims, wanted %+v, got %+v", claims, rawClaims)
	}
	// Purposely invalidate token
	// Couldn't figure out a shorter way to invalidate
	// Adding characters also errors but I want to check specifically for modifying the jwt
	if token[len(token)-1] >= '0' && token[len(token)-1] <= '9' {
		bytes := []byte(token)
		bytes[len(token)-1] = 'a'
		token = string(bytes)
	} else {
		bytes := []byte(token)
		bytes[len(token)-1] = '0'
		token = string(bytes)
	}

	_, err = jwtService.VerifyJWT(&token)
	if err == nil {
		t.Error("Invalidated JWT still passed...")
	}
}
