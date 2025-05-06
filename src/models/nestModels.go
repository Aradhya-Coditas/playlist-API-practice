package models

import "crypto/rsa"

// GetInitialKeyResponse represents the response structure when fetching public key 1.
type GetInitialKeyResponse struct {
	Status      string `json:"stat"`
	TomcatCount string `json:"tomcatCount"`
	PublicKey   string `json:"publicKey"`
}

// GetPreAuthKeyResponse represents the response structure when fetching public key 3.
type GetPreAuthKeyResponse struct {
	Status      string `json:"stat"`
	TomcatCount string `json:"tomcatCount"`
	PublicKey3  string `json:"publicKey3"`
}

// NestKeyPair represents a pair of public and private keys.
type NestKeyPair struct {
	PrivateKey       *rsa.PrivateKey
	PublicKey        *rsa.PublicKey
	PublicHashedKey  string
	PrivateHashedKey string
}
