package main

import (
	"encoding/json"
	"errors"
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
	if resp.StatusCode != 200 && resp.StatusCode != 302 {
		return nil, errors.New("Bad http response code.")
	}

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

func TournamentBalance(client *http.Client) (int, error) {
	resp, err := client.Get("http://saltybet.com/ajax_tournament_start.php")
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

type SaltyState struct {
	P1name    string
	P2name    string
	P1total   string
	P2total   string
	Status    string
	Alert     string
	X         int
	Remaining string
}

func GetState(client *http.Client) (SaltyState, error) {
	var state SaltyState
	resp, err := client.Get("http://saltybet.com/state.json")
	if err != nil {
		return state, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal(b, &state)
	if err != nil {
		return state, err
	}
	return state, nil
}
