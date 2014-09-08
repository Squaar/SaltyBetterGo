package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type cookieJar struct {
	jar map[string][]*http.Cookie
}

func (p *cookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	p.jar[u.Host] = cookies
}

func (p *cookieJar) Cookies(u *url.URL) []*http.Cookie {
	return p.jar[u.Host]
}

func AuthClient(email, pass string) (*http.Client, error) {
	jar := &cookieJar{jar: make(map[string][]*http.Cookie)}
	client := &http.Client{Jar: jar}

	data := url.Values{
		"email":        []string{email},
		"pword":        []string{pass},
		"authenticate": []string{"signin"},
	}
	resp, err := client.PostForm("http://www.saltybet.com/authenticate?signin=1", data)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Status)
	return client, nil
}

func WalletBalance(client *http.Client) (int, error) {
	resp, err := client.Get("http://saltybet.com/ajax_tournament_end.php")
	if err != nil {
		return 0, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	balance, _ := strconv.Atoi(string(b))
	return balance, nil
}
