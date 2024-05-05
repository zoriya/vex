package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/zoryia/vex/tui/models"
)

type statusMsg int

type errMsg struct{ error }
type missingJwtMsg struct{}
type noJwtMsg struct{}
type invalidJwtMsg struct{}
type httpErrorMsg error

func (e errMsg) Error() string { return e.error.Error() }

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
	if err != nil {
		return nil, httpErrorMsg(err)
	}
	return data, nil

}

func checkJwt(jwt *string) tea.Cmd {

	return func() tea.Msg {
		if jwt == nil || *jwt == "" {
			return noJwtMsg{}
		}
		url := fmt.Sprintf("%s/me", serverUrl)
		req, _ := http.NewRequest(http.MethodPost, url, nil)
		req.Header.Add("Content-type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *jwt))
		data, err := getData(req)
		if err != nil {
			return invalidJwtMsg{}
		}
		var resp struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}
		err = json.Unmarshal(data, &resp)
		if err != nil {
			return invalidJwtMsg{}
		}
		return nil
	}

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

type getEntriesSuccessMsg []models.Entry

func getEntries(jwt *string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("%s/entries", serverUrl)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		if jwt == nil {
			return missingJwtMsg{}
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *jwt))
		req.Header.Add("Content-type", "application/json")
		data, err := getData(req)
		if err != nil {
			return httpErrorMsg(err)
		}
		var entries []models.Entry
		err = json.Unmarshal(data, &entries)
		if err != nil {
			return httpErrorMsg(err)
		}
		return getEntriesSuccessMsg(entries)
	}
}

type getFeedsSuccessMsg []models.Feed

func getFeeds(jwt *string) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("%s/feeds", serverUrl)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		if jwt == nil {
			return missingJwtMsg{}
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *jwt))
		data, err := getData(req)
		if err != nil {
			return httpErrorMsg(err)
		}
		var feeds []models.Feed
		err = json.Unmarshal(data, &feeds)
		if err != nil {
			return httpErrorMsg(err)
		}
		return getFeedsSuccessMsg(feeds)
	}
}

type ignorePostSuccessMsg uuid.UUID

func ignorePost(jwt *string, e models.Entry) tea.Cmd {
	return func() tea.Msg {
		url := fmt.Sprintf("%s/ignore/%s", serverUrl, e.Id.String())
		req, _ := http.NewRequest(http.MethodPut, url, nil)
		if jwt == nil {
			return missingJwtMsg{}
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *jwt))
		data, err := getData(req)
		if err != nil {
			return httpErrorMsg(err)
		}
		var feeds []models.Feed
		err = json.Unmarshal(data, &feeds)
		if err != nil {
			return httpErrorMsg(err)
		}
		return getFeedsSuccessMsg(feeds)
	}
}

type toggleReadSuccessMsg uuid.UUID

func toggleRead(jwt *string, e models.Entry) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

type toggleReadLaterSuccessMsg uuid.UUID

func toggleReadLater(jwt *string, e models.Entry) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

type toggleBookmarkSuccessMsg uuid.UUID

func toggleBookmark(jwt *string, e models.Entry) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}
