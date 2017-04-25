package main

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
)

type SaltyClient struct {
	http.Client
	SaltyState //TODO: maybe make this a private var?
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

func NewSaltyClient(email, pass string) (*SaltyClient, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	client := &SaltyClient{http.Client{Jar: jar}, *new(SaltyState)}

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

func (client *SaltyClient) GetWalletBalance() (int, error) {
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

func (client *SaltyClient) TournamentBalance() (int, error) {
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

func (client *SaltyClient) GetState() (SaltyState, error) {
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

func (client *SaltyClient) PlaceBet(player, ammount int) error {
	data := url.Values{
		"selectedplayer": {"player" + strconv.Itoa(player)},
		"wager":          {strconv.Itoa(ammount)},
	}

	resp, err := client.PostForm("http://www.saltybet.com/ajax_place_bet.php", data)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Bad http response code.")
	}
	return nil
}
