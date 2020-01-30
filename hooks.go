package main

import (
	"log"

	"crypto/rand"

	coprocess "github.com/TykTechnologies/tyk-protobuf/bindings/go"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// MyPreHook1 performs a header injection:
func MyPreHook1(object *coprocess.Object) (*coprocess.Object, error) {
	log.Println("MyPreHook1 is called")
	object.Request.SetHeaders = map[string]string{
		"Myheader": "Myvalue",
	}
	var err error
	targetSz := grpcMaxSize - 1000
	log.Println("MyPrehook1, generating random bytes with length=", targetSz)
	object.Request.RawBody, err = GenerateRandomBytes(targetSz)
	if err != nil {
		panic(err)
	}
	log.Println("MyPreHook1 is called, raw body length=", len(object.Request.RawBody))

	return object, nil
}

// MyPreHook2 performs a header injection:
func MyPreHook2(object *coprocess.Object) (*coprocess.Object, error) {
	log.Println("MyPreHook2 is called, raw body length=", len(object.Request.RawBody))
	object.Request.SetHeaders = map[string]string{
		"Myheader": "Myvalue",
	}

	return object, nil
}
