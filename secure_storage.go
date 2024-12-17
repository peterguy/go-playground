package main

import (
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
