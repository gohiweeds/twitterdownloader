package twitterdownloader

import (
	"errors"
	"fmt"
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
func (c *Client) Init() {
	c.client = &http.Client{}
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

func (c *Client) GetWithProxy(url string) error {
	if c.client == nil {
		return errors.New("Client is nil, should init with proxy")
	}
	resp, err := c.client.Get(url)

	if err != nil {
		fmt.Println("GetWithProxy:", err.Error())
		return err
	}
	defer resp.Body.Close()

	//TODO:check url whether it is a jpeg or video

	// if !strings.Contains(url, ".jpeg") &&
	// 	!strings.HasSuffix(url, ".js") {
	// 	//Parse response body
	// 	videoUrl := c.ParseVideoUrl(resp.Body)
	// 	if videoUrl != "" {
	// 		//This is video
	// 		//return errors.New("video url parsed failed")
	// 		fmt.Println("Video Link:", videoUrl)
	// 		fmt.Println("Client:", c)
	// 		return downloader(videoUrl, c)
	// 	}
	// }

	fileName := extractFilename(url)
	fileName = strings.Split(fileName, ":")[0]
	body, err := ioutil.ReadAll(resp.Body)

	if err = ioutil.WriteFile(fileName, body, os.ModePerm); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
func extractFilename(url string) string {
	strs := strings.Split(url, "/")
	fileName := strs[len(strs)-1]
	// for k, v := range strs {
	// 	fmt.Println(k, v)
	// }

	params := strings.Split(fileName, "&")
	param := params[len(params)-1]
	return param
}

func (c *Client) RetriveVideoConfig(videoUrl string) error {
	if c.client == nil {
		return errors.New("Client is nil, should init with proxy")
	}
	resp, err := c.client.Get(videoUrl)

	if err != nil {
		fmt.Println("GetWithProxy:", err.Error())
		return err
	}
	defer resp.Body.Close()

	c.ParseVideoJS(resp.Body)
	//Parse Json
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Retrive video config failed")
		return err
	}

	fmt.Println("video config:", string(body))
	return nil
}
