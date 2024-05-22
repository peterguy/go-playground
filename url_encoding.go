package main

import (
	"fmt"
	"net/url"
)

func url_encoding() {
	path := "some-path-here"
	secret := "foo bar"
	u, err := url.Parse("https://example.com/")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("query escaped secret: %s\n", url.QueryEscape(secret))
	u.Path = ".api/" + path
	q := u.Query()
	q.Set("secret", secret)
	fmt.Printf("encoded query: %s\n", q.Encode())
	u.RawQuery = q.Encode()
	fmt.Printf("url: %s\n", u.String())
	fmt.Printf("secret: %s\n", u.Query().Get("secret"))
}
