package twitterdownloader

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	client       *http.Client
	proxyType    int
	proxyUrl     string
	socks5Proto  string
	socks5IpPort string
}

/// Init client structure
func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c Client) Get(url string) error {
	if c.client == nil {
		c.client = &http.Client{}
	}

	resp, err := c.client.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileName := extractFilename(url)

	body, err := ioutil.ReadAll(resp.Body)

	if err = ioutil.WriteFile(fileName, body, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetWithProxy(url string) (string, error) {
	if c.client == nil {
		return "", errors.New("Client is nil, should init with proxy")
	}
	resp, err := c.client.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fileName := extractFilename(url)
	fileName = strings.Split(fileName, ":")[0]
	body, err := ioutil.ReadAll(resp.Body)

	if err = ioutil.WriteFile(fileName, body, os.ModePerm); err != nil {
		return "", err
	}
	return fileName, nil
}
func extractFilename(url string) string {
	strs := strings.Split(url, "/")
	fileName := strs[len(strs)-1]

	params := strings.Split(fileName, "&")
	param := params[len(params)-1]
	return param
}
