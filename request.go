package twitterdownloader

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/osrtss/rtss/m3u8"
)

type Twitter struct {
	guest_token string
	playbackUrl string
}

// t = `
// OPTIONS https://api.twitter.com/1.1/guest/activate.json
// Accept:*/*
// Accept-Encoding:gzip,deflate,br
// Access-Control-Request-Headers:authorization,x-csrf-token
// Access-Control-Request-Method:POST
// Origin: https://twitter.com
// User-Agent:Mozilla/5.0 (X11; Linux x86_64)
// `

func (t *Twitter) Activate(c *Client) error {
	if c.client == nil {
		return errors.New("Client is nil, should init with proxy")
	}

	activateUrl := "https://api.twitter.com/1.1/guest/activate.json"
	req, err := http.NewRequest("OPTIONS", activateUrl, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip,deflate,br")
	req.Header.Add("Access-Control-Request-Headers", "authorization,x-csrf-token")
	req.Header.Add("Access-Control-Request-Method", "POST")
	req.Header.Add("Origin", "https://twitter.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")

	if err != nil {
		fmt.Println("GetWithProxy:", err.Error())
		return err
	}

	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	//fmt.Println("Activate Response: ", resp)
	return nil
}

// tt = `
// :authority: api.twitter.com
// :method: POST
// :path: /1.1/guest/activate.json
// :scheme: https
// accept: */*
// accept-encoding: gzip, deflate, br
// accept-language: zh-CN,zh;q=0.9,ar;q=0.8
// authorization: Bearer AAAAAAAAAAAAAAAAAAAAAIK1zgAAAAAA2tUWuhGZ2JceoId5GwYWU5GspY4%3DUq7gzFoCZs1QfwGoVdvSac3IniczZEYXIcDyumCauIXpcAPorE
// content-length: 0
// cookie: personalization_id="v1_pid1UUVchOmH31FJFT2ZLA=="; guest_id=v1%3A153569557385510737
// origin: https://twitter.com
// referer: https://twitter.com/i/videos/tweet/1035056498307522560
// user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/65.0.3325.181 Chrome/65.0.3325.181 Safari/537.36
// x-csrf-token: undefined
// `

type GuestToken struct {
	GuestToken string `json:"guest_token"`
}

func (t *Twitter) Activate2(c *Client) error {
	if c.client == nil {
		return errors.New("Client is nil, should init with proxy")
	}
	activateUrl := "https://api.twitter.com/1.1/guest/activate.json"
	req, err := http.NewRequest("POST", activateUrl, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip,deflate,br")
	req.Header.Add("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAIK1zgAAAAAA2tUWuhGZ2JceoId5GwYWU5GspY4%3DUq7gzFoCZs1QfwGoVdvSac3IniczZEYXIcDyumCauIXpcAPorE")
	req.Header.Add("cookie", `personalization_id="v1_pid1UUVchOmH31FJFT2ZLA=="; guest_id=v1%3A153569557385510737`)
	req.Header.Add("content-length", "0")
	req.Header.Add("Origin", "https://twitter.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")

	if err != nil {
		fmt.Println("GetWithProxy:", err.Error())
		return err
	}

	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}
	body, err := ioutil.ReadAll(reader)

	var guest GuestToken
	err = json.Unmarshal(body, &guest)
	fmt.Println("respJons:", guest.GuestToken, err)
	t.guest_token = guest.GuestToken
	return nil
}

// tt=`
// :authority: api.twitter.com
// :method: OPTIONS
// :path: /1.1/videos/tweet/config/1035056498307522560.json
// :scheme: https
// accept: */*
// accept-encoding: gzip, deflate, br
// accept-language: zh-CN,zh;q=0.9,ar;q=0.8
// access-control-request-headers: authorization,x-csrf-token,x-guest-token
// access-control-request-method: GET
// origin: https://twitter.com
// user-agent: Mozilla/5.0 (X11; Linux x86_64)
// `
func (t *Twitter) GetVideoJson(c *Client) error {
	if c.client == nil {
		return errors.New("Client is nil, should init with proxy")
	}

	jsonUrl := "https://api.twitter.com/1.1/videos/tweet/config/1035056498307522560.json"
	req, err := http.NewRequest("OPTIONS", jsonUrl, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip,deflate,br")
	req.Header.Add("Access-Control-Request-Headers", "authorization,x-csrf-token")
	req.Header.Add("Access-Control-Request-Method", "GET")
	req.Header.Add("Origin", "https://twitter.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")

	if err != nil {
		fmt.Println("GetWithProxy:", err.Error())
		return err
	}

	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	//fmt.Println("JsonConf Response: ", resp)
	return nil
}

// tt=`
// :authority: api.twitter.com
// :method: GET
// :path: /1.1/videos/tweet/config/1035056498307522560.json
// :scheme: https
// accept: */*
// accept-encoding: gzip, deflate, br
// accept-language: zh-CN,zh;q=0.9,ar;q=0.8
// authorization: Bearer AAAAAAAAAAAAAAAAAAAAAIK1zgAAAAAA2tUWuhGZ2JceoId5GwYWU5GspY4%3DUq7gzFoCZs1QfwGoVdvSac3IniczZEYXIcDyumCauIXpcAPorE
// cookie: personalization_id="v1_pid1UUVchOmH31FJFT2ZLA=="; guest_id=v1%3A153569557385510737; gt=1035408416074657792
// origin: https://twitter.com
// referer: https://twitter.com/i/videos/tweet/1035056498307522560
// user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/65.0.3325.181 Chrome/65.0.3325.181 Safari/537.36
// x-csrf-token: undefined
// x-guest-token: 1035408416074657792
// `

type VideoConfig struct {
	Track Track `json:"track"`
}
type Track struct {
	ContentId    string `json:"contentId"`
	ContentType  string `json:"contentType"`
	Duration     int    `json:"durationMs"`
	ExpandedUrl  string `json:"expandedUrl"`
	PlaybackType string `json:"playbackType"`
	PlaybackUrl  string `json:"playbackUrl"`
}

func (t *Twitter) GetVideoJson2(c *Client) error {
	if c.client == nil {
		return errors.New("Client is nil, should init with proxy")
	}
	jsonUrl := "https://api.twitter.com/1.1/videos/tweet/config/1035056498307522560.json"
	req, err := http.NewRequest("GET", jsonUrl, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip,deflate,br")
	req.Header.Add("authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAIK1zgAAAAAA2tUWuhGZ2JceoId5GwYWU5GspY4%3DUq7gzFoCZs1QfwGoVdvSac3IniczZEYXIcDyumCauIXpcAPorE")
	req.Header.Add("cookie", `personalization_id="v1_pid1UUVchOmH31FJFT2ZLA=="; guest_id=v1%3A153569557385510737`)
	req.Header.Add("content-length", "0")
	req.Header.Add("Origin", "https://twitter.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	req.Header.Add("x-guest-token", t.guest_token)

	if err != nil {
		fmt.Println("GetWithProxy:", err.Error())
		return err
	}

	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}
	body, err := ioutil.ReadAll(reader)

	var vc VideoConfig
	err = json.Unmarshal(body, &vc)
	//fmt.Println("respJons:", vc, err)
	t.playbackUrl = vc.Track.PlaybackUrl
	return nil
}

func (t *Twitter) Getm3u8List(c *Client) error {
	if c.client == nil {
		return errors.New("Client is nil, should init with proxy")
	}

	req, err := http.NewRequest("GET", t.playbackUrl, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip,deflate,br")
	// req.Header.Add("Access-Control-Request-Headers", "authorization,x-csrf-token")
	// req.Header.Add("Access-Control-Request-Method", "GET")
	req.Header.Add("Origin", "https://twitter.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")

	if err != nil {
		fmt.Println("GetWithProxy:", err.Error())
		return err
	}

	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}
	body, err := ioutil.ReadAll(reader)

	// fileName := extractFilename(t.playbackUrl)
	fmt.Println("m3u8 list:\n", string(body))
	// if err = ioutil.WriteFile(fileName, body, os.ModePerm); err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	playlist, listType, err := m3u8.DecodeFrom(resp.Body, false)
	if err != nil {
		log.Fatalf("M3U8 decode failed, err: %v", err)
	}

	fmt.Println("palylist:", playlist)
	fmt.Println("palyType:", listType)
	return nil
}
