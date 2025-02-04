package main

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/goingtharn/aoc2023/aoc"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "list":
		fmt.Println("\nAvailable solutions (aka keys):")
		for _, k := range aoc.ListSolutions() {
			fmt.Printf("  * %s\n", k)
		}
	case "run":
		run(args[1])
	case "submit":
		res := run(args[1])
		if res != nil {
			fmt.Println("===\nSubmitting...\n===")
			fmt.Println(submit(args[1], res.answer))
		}
	case "all":
		for _, k := range aoc.ListSolutions() {
			fmt.Printf("\n%s\n", k)
			run(k)
		}
	case "stubs":
		stubs(args[1])
	default:
		panic("first arg must be 'run' or 'start'")
	}
}

const defaultTimeout = "300" // 5 minutes

type result struct {
	answer   string
	duration time.Duration
}

func run(key string) *result {
	attempt := aoc.SolutionFor(key)
	if attempt == nil {
		fmt.Println("no solution available for key")
		os.Exit(1)
	}

	year, day := parseKey(key)
	input := fetchInput(year, day)

	timeout := os.Getenv("AOC_TIMEOUT")
	if timeout == "" {
		timeout = defaultTimeout
	}
	wait, err := strconv.Atoi(timeout)
	if err != nil {
		panic(err)
	}

	rc := make(chan result)
	go func() {
		start := time.Now()
		answer := attempt(input)
		rc <- result{answer, time.Since(start)}
	}()

	select {
	case res := <-rc:
		fmt.Printf("\nAnswer in %v\n---\n%s\n", res.duration, res.answer)
		return &res
	case <-time.After(time.Duration(wait) * time.Second):
		fmt.Println("Too slow!")
		return nil
	}
}

func parseKey(key string) (string, string) {
	parts := strings.Split(key, ":")
	if len(parts) < 2 {
		panic("expected at least <year>:<day> in key")
	}
	return parts[0], parts[1]
}

func fetchInput(year, day string) string {
	token := strings.TrimSpace(os.Getenv("AOC_SESSION"))

	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

	var client http.Client
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", "session="+token)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("input fetch returned unexpected status: %d", resp.StatusCode))
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(bodyBytes)
}

func submit(key, answer string) string {
	year, day := parseKey(key)
	level := strings.Split(key, ":")[2]
	token := strings.TrimSpace(os.Getenv("AOC_SESSION"))

	form := url.Values{}
	form.Add("answer", answer)
	form.Add("level", level)

	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", year, day)

	var client http.Client
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", "session="+token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("submit post returned unexpected status: %d", resp.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	return doc.Find("main article").First().Text()
}

//go:embed templates/*
var templateFS embed.FS

func stubs(key string) {
	year, day := parseKey(key)

	// bail if source/test files already exist

	srcName := fmt.Sprintf("y%sd%s.go", year, day)
	srcPath := filepath.Join("aoc", srcName)
	_, err := os.Stat(srcPath)
	if !errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("%s exists", srcPath))
	}

	tstName := fmt.Sprintf("y%sd%s_test.go", year, day)
	tstPath := filepath.Join("aoc", tstName)
	_, err = os.Stat(srcPath)
	if !errors.Is(err, os.ErrNotExist) {
		panic(fmt.Sprintf("%s exists", tstPath))
	}

	tmp, err := template.ParseFS(templateFS, "*/*.tmpl")
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"Year":     year,
		"Day":      day,
		"FuncName": fmt.Sprintf("y%sd%spart", year, day),
	}

	srcOut, err := os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	err = tmp.ExecuteTemplate(srcOut, "solution.go.tmpl", data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("created %s\n", srcPath)

	tstOut, err := os.OpenFile(tstPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	err = tmp.ExecuteTemplate(tstOut, "solution_test.go.tmpl", data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("created %s\n", tstPath)
}
