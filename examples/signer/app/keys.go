package main

import (
	"crypto/ed25519"
	"crypto/sha512"
	"strconv"

	"filippo.io/edwards25519"
)

// generateKeys generates a new public/private key pair.
// based on code from https://cs.opensource.google/go/go/+/refs/tags/go1.23.5:src/crypto/ed25519/ed25519.go;l=158
func generateKeys() {
	seed := CDI()
	h := sha512.Sum512(seed)
	s, err := edwards25519.NewScalar().SetBytesWithClamping(h[:32])
	if err != nil {
		panic("ed25519: internal error: setting scalar failed")
	}

	p := edwards25519.NewGeneratorPoint()
	a := p.ScalarMultSlow(s, p)
	publicKey = make([]byte, ed25519.PublicKeySize)
	copy(publicKey, a.Bytes())

	privateKey = make([]byte, ed25519.PrivateKeySize)
	copy(privateKey, seed)
	copy(privateKey[32:], publicKey)
}

// signMessage signs the message with privateKey and returns a signature. It will
// panic if len(privateKey) is not [PrivateKeySize].
// based on code from https://cs.opensource.google/go/go/+/refs/tags/go1.23.5:src/crypto/ed25519/ed25519.go;l=189
func signMessage(privateKey ed25519.PrivateKey, message []byte) []byte {
	// Outline the function body so that the returned signature can be
	// stack-allocated.
	signature := make([]byte, ed25519.SignatureSize)
	sign(signature, privateKey, message, domPrefixPure, "")
	return signature
}

// Domain separation prefixes used to disambiguate Ed25519/Ed25519ph/Ed25519ctx.
// See RFC 8032, Section 2 and Section 5.1.
const (
	// domPrefixPure is empty for pure Ed25519.
	domPrefixPure = ""
	// domPrefixPh is dom2(phflag=1) for Ed25519ph. It must be followed by the
	// uint8-length prefixed context.
	domPrefixPh = "SigEd25519 no Ed25519 collisions\x01"
	// domPrefixCtx is dom2(phflag=0) for Ed25519ctx. It must be followed by the
	// uint8-length prefixed context.
	domPrefixCtx = "SigEd25519 no Ed25519 collisions\x00"
)

// sign signs the message with privateKey and writes the signature to signature.
// based on code from https://cs.opensource.google/go/go/+/refs/tags/go1.23.5:src/crypto/ed25519/ed25519.go;l=210
func sign(signature, privateKey, message []byte, domPrefix, context string) {
	if l := len(privateKey); l != ed25519.PrivateKeySize {
		panic("ed25519: bad private key length: " + strconv.Itoa(l))
	}
	seed, publicKey := privateKey[:ed25519.SeedSize], privateKey[ed25519.SeedSize:]

	h := sha512.Sum512(seed)
	s, err := edwards25519.NewScalar().SetBytesWithClamping(h[:32])
	if err != nil {
		panic("ed25519: internal error: setting scalar failed")
	}
	prefix := h[32:]

	mh := sha512.New()
	if domPrefix != domPrefixPure {
		mh.Write([]byte(domPrefix))
		mh.Write([]byte{byte(len(context))})
		mh.Write([]byte(context))
	}
	mh.Write(prefix)
	mh.Write(message)
	messageDigest := make([]byte, 0, sha512.Size)
	messageDigest = mh.Sum(messageDigest)
	r, err := edwards25519.NewScalar().SetUniformBytes(messageDigest)
	if err != nil {
		panic("ed25519: internal error: setting scalar failed")
	}

	p := edwards25519.NewGeneratorPoint()
	R := p.ScalarMultSlow(r, p)

	kh := sha512.New()
	if domPrefix != domPrefixPure {
		kh.Write([]byte(domPrefix))
		kh.Write([]byte{byte(len(context))})
		kh.Write([]byte(context))
	}
	kh.Write(R.Bytes())
	kh.Write(publicKey)
	kh.Write(message)
	hramDigest := make([]byte, 0, sha512.Size)
	hramDigest = kh.Sum(hramDigest)
	k, err := edwards25519.NewScalar().SetUniformBytes(hramDigest)
	if err != nil {
		panic("ed25519: internal error: setting scalar failed")
	}

	S := edwards25519.NewScalar().MultiplyAdd(k, s, r)

	copy(signature[:32], R.Bytes())
	copy(signature[32:], S.Bytes())
}
