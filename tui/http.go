package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

type getEntriesSuccessMsg []Entry
type getEntriesErrMsg error

const serverUrl = "localhost:3000"

func getEntries() tea.Msg {
	url := fmt.Sprintf("%s/entries", serverUrl)

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

type getFeedsSuccessMsg []Feed
type getFeedsErrMsg error
