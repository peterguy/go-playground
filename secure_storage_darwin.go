//go:build darwin

package main

import (
	"bytes"

	"github.com/keybase/go-keychain"
)

func secure_store(service, account string, secret []byte) error {

	if existing_secret, err := secure_retrieve(service, account); err == nil {
		if bytes.Equal(secret, existing_secret) {
			// the stored secret and supplied secret are the same
			// nothing to do
			return nil
		}
		// found a secret, but it is different from the supplied secret, so update the secret
		return keychain.UpdateItem(buildItem(service, account, nil), buildItem(service, account, secret))
	} else if err != keychain.ErrorItemNotFound {
		// encountered an error checking for the secret; bail now
		return err
	}

	item := buildItem(service, account, secret)
	if err := keychain.AddItem(item); err != nil {
		if err == keychain.ErrorDuplicateItem {
			// silently skip duplicates
			// really shouldn't happen after the duplication detection above,
			// but just in case (this process is not atomic - another running process could have snuck a duplicate in the keychain)
			return nil
		}
		return err
	}
	return nil
}

func secure_retrieve(service, account string) ([]byte, error) {
	query := buildItem(service, account, nil)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return nil, err
	} else if len(results) != 1 {
		return nil, keychain.ErrorItemNotFound
	} else {
		return results[0].Data, nil
	}
}

func buildItem(service, account string, secret []byte) keychain.Item {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	item.SetAccessGroup("go-playground")
	item.SetService(service)
	item.SetAccount(account)
	item.SetData(secret)
	item.SetComment("playground secret")
	return item
}
