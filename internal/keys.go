package internal

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
)

// PrivateKey is the master key to create a licenses.
// Private key needs to be put in a safe location
type PrivateKey ecdsa.PrivateKey

// PublicKey is used to check the validity of the licenses. You can share it
// freely.
type PublicKey ecdsa.PublicKey

// NewPrivateKey generates a new private key
func NewPrivateKey() (*PrivateKey, error) {
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}

	privateKey := PrivateKey(*key)
	return &privateKey, nil
}

// ToBytes transforms the private key to a  []byte.
func (k PrivateKey) ToBytes() ([]byte, error) {
	ek := k.toEcdsa()
	c := elliptic.Marshal(
		ek.PublicKey.Curve,
		ek.PublicKey.X,
		ek.PublicKey.Y)

	return toBytes(c)
}

// ToBase64String transforms the private key to a base64 string.
func (k PrivateKey) ToBase64String() (string, error) {
	b, err := k.ToBytes()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (k *PrivateKey) toEcdsa() *ecdsa.PrivateKey {
	r := ecdsa.PrivateKey(*k)
	return &r
}

func toBytes(obj interface{}) ([]byte, error) {
	var buffBin bytes.Buffer

	encoderBin := gob.NewEncoder(&buffBin)
	if err := encoderBin.Encode(obj); err != nil {
		return nil, err
	}

	return buffBin.Bytes(), nil
}

// GetPublicKey returns the PublicKey associated with the private key.
func (k PrivateKey) GetPublicKey() *PublicKey {
	pk := PublicKey(k.PublicKey)
	return &pk
}

// ToBase64String transforms the public key to a base64 string.
func (k PublicKey) ToBase64String() string {
	return base64.StdEncoding.EncodeToString(
		k.ToBytes(),
	)
}

// ToBytes transforms the public key to a []byte.
func (k PublicKey) ToBytes() []byte {
	// return toBytes(k)
	ek := k.toEcdsa()
	return elliptic.Marshal(ek.Curve, ek.X, ek.Y)
}

func (k *PublicKey) toEcdsa() *ecdsa.PublicKey {
	r := ecdsa.PublicKey(*k)
	return &r
}
