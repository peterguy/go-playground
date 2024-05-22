//go:build !darwin

package main

import "errors"

func secure_store(service, account string, secret []byte) error {
	return errors.New("not implemented")
}

func secure_retrieve(service, account string) ([]byte, error) {
	return nil, errors.New("not implemented")
}
