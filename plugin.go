package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-zglob"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
)

const (
	RequestURI = "https://codecov.io/upload/v2?%s"
)

type (
	Repo struct {
		Fullname string
	}

	Build struct {
		Number      int
		Link        string
		PullRequest int
	}

	Commit struct {
		Sha    string
		Branch string
		Tag    string
	}

	Config struct {
		Token   string
		Name    string
		Pattern string
		Files   []string
		Flags   []string
	}

	Internal struct {
		Matches []string
		Merged  []byte
	}

	Plugin struct {
		Repo     Repo
		Build    Build
		Commit   Commit
		Config   Config
		Internal Internal
	}
)

func (p *Plugin) Exec() error {
	if p.Config.Token == "" {
		return errors.New("you must provide a token")
	}

	if p.Commit.Sha == "" {
		return errors.New("you must provide a commit")
	}

	if err := p.match(); err != nil {
		return err
	}

	if err := p.merge(); err != nil {
		return err
	}

	if err := p.submit(); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) match() error {
	if len(p.Config.Files) > 0 {
		log.Printf("using coverage files: %s", strings.Join(p.Config.Files, ", "))
		matches := make([]string, 0)

		for _, f := range p.Config.Files {
			if _, err := os.Stat(f); err == nil {
				matches = append(matches, f)
			}
		}

		if len(matches) > 0 {
			log.Printf("found coverage files: %s", strings.Join(matches, ", "))
		} else {
			log.Printf("no coverage files found")
		}

		p.Internal.Matches = matches
	} else {
		log.Printf("searching coverage files: %s", p.Config.Pattern)
		matches, err := zglob.Glob(p.Config.Pattern)

		if err != nil {
			return errors.Wrap(err, "failed to match files")
		}

		if len(matches) > 0 {
			log.Printf("found coverage files: %s", strings.Join(matches, ", "))
		} else {
			log.Printf("no coverage files found")
		}

		p.Internal.Matches = matches
	}

	return nil
}

func (p *Plugin) merge() error {
	if len(p.Internal.Matches) == 0 {
		return nil
	}

	buf := bytes.NewBufferString("")

	for _, f := range p.Internal.Matches {
		content, err := ioutil.ReadFile(f)

		if err != nil {
			return errors.Wrap(err, "failed to read file")
		}

		buf.WriteString(string(content))
		buf.WriteString("<<<<<< EOF")
	}

	p.Internal.Merged = buf.Bytes()

	return nil
}

func (p *Plugin) submit() error {
	if len(p.Internal.Matches) == 0 {
		return nil
	}

	if err := os.MkdirAll("/tmp", os.ModePerm); err != nil {
		return errors.Wrap(err, "failed to create tempdir")
	}

	tmpfile, err := ioutil.TempFile("", "codecov-")

	if err != nil {
		return errors.Wrap(err, "failed to create tempfile")
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(p.Internal.Merged); err != nil {
		return errors.Wrap(err, "failed to write tempfile")
	}

	v := url.Values{}
	v.Set("service", "drone.io")
	v.Set("commit", p.Commit.Sha)
	v.Set("token", p.Config.Token)

	if p.Commit.Branch != "" {
		v.Set("branch", p.Commit.Branch)
	}

	if p.Commit.Tag != "" {
		v.Set("tag", p.Commit.Tag)
	}

	if p.Build.Number != 0 {
		v.Set("build", strconv.Itoa(p.Build.Number))
	}

	if p.Build.Link != "" {
		v.Set("build_url", p.Build.Link)
	}

	if p.Config.Name != "" {
		v.Set("name", p.Config.Name)
	}

	if p.Repo.Fullname != "" {
		v.Set("slug", p.Repo.Fullname)
	}

	if p.Repo.Fullname != "" {
		v.Set("slug", p.Repo.Fullname)
	}

	if p.Build.PullRequest != 0 {
		v.Set("pr", strconv.Itoa(p.Build.PullRequest))
	}

	if len(p.Config.Flags) != 0 {
		v.Set("flags", strings.Join(p.Config.Flags, ","))
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(RequestURI, v.Encode()),
		bytes.NewBuffer(p.Internal.Merged),
	)

	if err != nil {
		return errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return errors.Wrap(err, "failed to process request")
	}

	defer res.Body.Close()

	if res.Status == "200 OK" {
		log.Printf("successfully uploaded coverage report")
	} else {
		reason := "Unknown"
		doc, err := html.Parse(res.Body)

		if err != nil {
			return errors.Wrap(err, "failed to parse body")
		}

		var f func(*html.Node)

		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "h1" {
				reason = strings.TrimSpace(n.FirstChild.Data)
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}

		f(doc)

		return errors.Errorf("failed to submit request: %s", reason)
	}

	return nil
}
