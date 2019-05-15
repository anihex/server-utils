package tools

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// RSA holds the basic informations about basic RSA values.
type RSA struct {
	Priv *rsa.PrivateKey
	Pub  *rsa.PublicKey
	Name string
}

// NewRSA generates a new instance of an RSA. It takes a name and creates a
// public and private key for it.
func NewRSA(name string) *RSA {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil
	}

	result := RSA{
		Name: name,
		Pub:  &priv.PublicKey,
		Priv: priv,
	}

	return &result
}

// Sign signs a message using the private key of the RSA.
func (c *RSA) Sign(msg string) string {
	var options rsa.PSSOptions
	options.SaltLength = rsa.PSSSaltLengthAuto

	hasher := sha256.New()
	hasher.Write([]byte(msg))
	hash := hasher.Sum(nil)

	result, err := rsa.SignPSS(
		rand.Reader,
		c.Priv,
		crypto.SHA256,
		hash,
		&options,
	)

	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("%x", result)
}

// Verify checks if a given message is signed correctly.
func (c *RSA) Verify(msg string, signature string) error {
	var options rsa.PSSOptions
	options.SaltLength = rsa.PSSSaltLengthAuto

	hasher := sha256.New()
	hasher.Write([]byte(msg))
	hash := hasher.Sum(nil)

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return err
	}

	err = rsa.VerifyPSS(
		c.Pub,
		crypto.SHA256,
		hash,
		sig,
		&options,
	)

	return err
}

// ExportPublicKey exports a RSA Public-Key to a given filename.
func (c *RSA) ExportPublicKey(filename string) error {
	pemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(c.Pub),
	}

	data := pem.EncodeToMemory(pemBlock)
	err := ioutil.WriteFile(filename, data, os.ModePerm)

	return err
}

// ExportPrivateKey exports a RSA Private-Key to a given filename.
func (c *RSA) ExportPrivateKey(filename string) error {
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(c.Priv),
	}

	data := pem.EncodeToMemory(pemBlock)
	err := ioutil.WriteFile(filename, data, os.ModePerm)

	return err
}

// ImportPublicKey imports a RSA Public-Key from a given filename.
func (c *RSA) ImportPublicKey(filename string) error {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data, _ := pem.Decode(buffer)
	c.Pub, err = x509.ParsePKCS1PublicKey(data.Bytes)

	return err
}

// ImportPrivateKey imports a RSA Private-Key from a given filename.
func (c *RSA) ImportPrivateKey(filename string) error {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	data, _ := pem.Decode(buffer)
	c.Priv, err = x509.ParsePKCS1PrivateKey(data.Bytes)

	return err
}
