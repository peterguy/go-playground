package main

import (
	"fmt"
	"os"

	"github.com/99designs/keyring"
)

func secure_store(service, account string, secret []byte) error {
	ring, err := keyring.Open(keyring.Config{
		ServiceName:                    service,
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: false,
	})
	if err != nil {
		return err
	}
	if err := ring.Set(keyring.Item{
		Key:         account,
		Data:        secret,
		Description: "secret key",
	}); err != nil {
		return err
	}
	return nil
	//return keyring.Set(service, account, string(secret))
}

func secure_retrieve(service, account string) ([]byte, error) {
	ring, err := keyring.Open(keyring.Config{
		ServiceName:                    service,
		KeychainTrustApplication:       true,
		KeychainSynchronizable:         false,
		KeychainAccessibleWhenUnlocked: false,
	})
	if err != nil {
		return nil, err
	}
	item, err := ring.Get(account)
	if err != nil {
		return nil, err
	}
	return item.Data, nil
	// secret, err := keyring.Get(service, account)
	// return []byte(secret), err
}

func main() {
	action := os.Args[1]
	secure_store_service := "example.com"
	secure_store_account := "peterguy"
	switch action {
	case "store":
		secret := os.Args[2]
		err := secure_store(secure_store_service, secure_store_account, []byte(secret))
		if err != nil {
			fmt.Printf("ERROR storing secret: %v\n", err)
		}
	case "retrieve":
		secret, err := secure_retrieve(secure_store_service, secure_store_account)
		if err != nil {
			fmt.Printf("ERROR retrieving secret: %v\n", err)
		} else {
			fmt.Printf("secret: %s\n", string(secret))
		}
	}
}
