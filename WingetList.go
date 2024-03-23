package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type WingetList struct {
	IDs []string
}

func NewWingetList() *WingetList {
	return &WingetList{
		IDs: []string{},
	}
}

func (w *WingetList) FetchFromGist() error {
	// Fetch the list of IDs from the gist
	resp, err := http.Get(Env.GistURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Print body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ids, err := ParseBody(string(body))
	if err != nil {
		return err
	}

	sort.Slice(ids, func(i, j int) bool {
		return strings.Compare(strings.ToLower(ids[i]), strings.ToLower(ids[j])) < 0
	})

	w.IDs = ids

	return nil
}

func (w *WingetList) RunInstall(m CLIModel) error {
	if len(m.GetSelected()) == 0 {
		return fmt.Errorf("no IDs to install")
	}

	// Run the install command
	wingetArguments := []string{}
	wingetArguments = append(wingetArguments, "install")
	wingetArguments = append(wingetArguments, m.GetSelected()...)

	fmt.Printf("Running: winget %s\n", strings.Join(wingetArguments, " "))

	cmd := exec.Command("winget", wingetArguments...)
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func ParseBody(body string) ([]string, error) {
	// Parse the body of the response
	ids := []string{}

	start := strings.Index(body, "install ")
	if start == -1 {
		return ids, fmt.Errorf("could not find 'install' in body")
	}
	start += len("install ")

	end := strings.Index(body[start:], "\n")
	if end == -1 {
		return ids, fmt.Errorf("could not find newline after 'install'")
	}

	ids = strings.Split(body[start:start+end], " ")

	return ids, nil
}
