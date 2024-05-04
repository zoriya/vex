package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

type statusMsg int

type errMsg struct{ error }
type missingJwtMsg struct{}

func (e errMsg) Error() string { return e.error.Error() }

type getEntriesSuccessMsg []Entry
type httpErrorMsg error

const serverUrl = "http://localhost:1597"

type loginSuccessMsg struct{ string }
type registerSuccessMsg struct{ string }

func getData(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("err", err)
		return nil, httpErrorMsg(err)
	}
	defer resp.Body.Close() // nolint: errcheck

	data, err := io.ReadAll(resp.Body)
	log.Print("body", resp.Body)
	if err != nil {
		return nil, httpErrorMsg(err)
	}
	return data, nil

}

func login(username string, password string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("%s/login", serverUrl)
		body := struct {
			Name     string `json:"email"`
			Password string `json:"password"`
		}{
			Name: username, Password: password,
		}
		out, err := json.Marshal(body)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(out))
		req.Header.Add("Content-type", "application/json")
		data, err := getData(req)
		if err != nil {
			return err
		}
		var loginResp AuthRes

		err = json.Unmarshal(data, &loginResp)
		if err != nil {
			return httpErrorMsg(err)
		}

		return loginSuccessMsg{loginResp.Token}
	}
}

type AuthRes struct {
	Token string `json:"token"`
}

func register(username string, password string, email string) tea.Cmd {
	return func() tea.Msg {

		url := fmt.Sprintf("%s/register", serverUrl)
		body := struct {
			Name     string `json:"name"`
			Password string `json:"password"`
			Email    string `json:"email"`
		}{
			Name: username, Password: password, Email: email,
		}
		out, err := json.Marshal(body)
		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(out))
		req.Header.Add("Content-type", "application/json")
		data, err := getData(req)
		if err != nil {
			return err
		}
		var registerResp AuthRes
		err = json.Unmarshal(data, &registerResp)
		if err != nil {
			return httpErrorMsg(err)
		}
		return registerSuccessMsg{registerResp.Token}
	}
}

func getEntries(jwt *string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("%s/entries", serverUrl)

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if jwt == nil {
			return missingJwtMsg{}
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *jwt))

		if err != nil {
			return httpErrorMsg(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return httpErrorMsg(err)
		}
		defer resp.Body.Close() // nolint: errcheck

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return httpErrorMsg(err)
		}

		var entries []Entry

		err = json.Unmarshal(data, &entries)
		if err != nil {
			return httpErrorMsg(err)
		}

		return getEntriesSuccessMsg(entries)
	}

}

type getFeedsSuccessMsg []Feed
type getFeedsErrMsg error
