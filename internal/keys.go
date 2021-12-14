package internal

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/gob"
)

// GenerateKeys generates a private and public key pair
func GenerateKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privkey, err := rsa.GenerateKey(rand.Reader, 368)
	if err != nil {
		return nil, nil, err
	}
	return privkey, &privkey.PublicKey, nil
}

// GenerateBase64Keys generates a private and public key pair and encodes them to base64 strings
func GenerateBase64Keys() (string, string, error) {
	priv, pub, err := GenerateKeys()

	if err != nil {
		return "", "", err
	}

	privB64, err := toB64String(priv)

	if err != nil {
		return "", "", err
	}

	pubB64, err := toB64String(pub)

	if err != nil {
		return "", "", err
	}

	return privB64, pubB64, nil
}

// toB64String converts an object to a base64 encoded string
func toB64String(obj interface{}) (string, error) {
	b, err := toBytes(obj)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func toBytes(obj interface{}) ([]byte, error) {
	var buffBin bytes.Buffer

	encoderBin := gob.NewEncoder(&buffBin)
	if err := encoderBin.Encode(obj); err != nil {
		return nil, err
	}

	return buffBin.Bytes(), nil
}
