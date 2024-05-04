package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

type statusMsg int

type errMsg struct{ error }
type missingJwtMsg struct{}

func (e errMsg) Error() string { return e.error.Error() }

type getEntriesSuccessMsg []Entry
type getEntriesErrMsg error

const serverUrl = "localhost:3000"

type loginSuccessMsg struct{ string }

func login(username string, password string) tea.Cmd {
	return func() tea.Msg {
		_ = username
		_ = password
		return loginSuccessMsg{"dawdaw"}
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
			return getEntriesErrMsg(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return getEntriesErrMsg(err)
		}
		defer resp.Body.Close() // nolint: errcheck

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return getEntriesErrMsg(err)
		}

		var entries []Entry

		err = json.Unmarshal(data, &entries)
		if err != nil {
			return getEntriesErrMsg(err)
		}

		return getEntriesSuccessMsg(entries)
	}

}

type getFeedsSuccessMsg []Feed
type getFeedsErrMsg error
