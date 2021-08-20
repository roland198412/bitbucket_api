package bitbucket_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type BitBucketClient struct {
	Url      string
	userName string
	passWord string
	client   http.Client
}

/**
Repository Struct
 */
type repositories struct {
	PageLen int `json:"pagelen"`
	Values  []values `json:"values"`
	Next string `json:"next"`
}

type values struct {
	Links links `json:"links"`
}

type links struct {
	Commits linksHRef `json:"commits"`
}

type linksHRef struct {
	HRef string `json:"href"`
}

/**
Commit Struct
*/
type commit struct {
	PageLen int `json:"pagelen"`
	Values  []commitValues `json:"values"`
	Next string `json:"next"`
}

type commitValues struct {
	Author author `json:"author"`
	Date string `json:"date"`
}

type author struct {
	Raw string `json:"raw"`
}

func NewBitBucketClient(url, userName, passWord string, timeout time.Duration) *BitBucketClient {
	b := BitBucketClient{
		Url:      url,
		userName: userName,
		passWord: passWord,
		client: http.Client{
			Timeout: timeout,
		},
	}

	return &b
}

func (b *BitBucketClient) GetRepos() (repo repositories, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s", b.Url), nil)
	req.SetBasicAuth(b.userName,b.passWord)

	if err != nil {
		return repo, err
	}

	response, err := b.client.Do(req)
	if err != nil {
		return repo, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return repo, err
	}

	err = json.Unmarshal(body, &repo)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
			log.Printf(string(body))
		}
		return repo, err
	}

	return
}

func (b *BitBucketClient) GetCommitDetail(url string) (commit commit, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s", url), nil)
	req.SetBasicAuth(b.userName,b.passWord)

	if err != nil {
		return commit, err
	}

	response, err := b.client.Do(req)
	if err != nil {
		return commit, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return commit, err
	}

	err = json.Unmarshal(body, &commit)
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
			log.Printf(string(body))
		}
		return commit, err
	}

	return
}
