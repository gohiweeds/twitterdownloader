package main

import (
	"flag"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	tw "github.com/gohiweeds/twitterdownloader"
)

func main() {
	url := flag.String("url", "", "download provided url")

	flag.Parse()

	if *url == "" {
		log.Errorf("URL not qualified or exist")
		return
	}
	client := tw.NewClient()

	//Setup SOCK5
	_, err := client.ClientWithSOCKS5("tcp", "127.0.0.1:1080")
	if err != nil {
		log.Errorf("SOCK5 Init failed: %s", err.Error())
		return
	}

	twitter := &tw.Twitter{}
	twitter.SetupClient(client)
	//twitter.DownloadVideo("https://twitter.com/i/status/1035056498307522560")
	if strings.Contains(*url, ".jpg") {
		fileName, err := twitter.DownloadJPG(*url)
		if err != nil {
			log.Errorf("Download JPG failed: %s", err.Error())
			return
		}
		log.Printf("Download JPG Success: %s", fileName)
	} else {
		//https://pbs.twimg.com/media/Dl_0nDtXoAAZu4x.jpg:large
		fileName, err := twitter.DownloadVideo(*url)
		if err != nil {
			log.Errorf("Download Video failed: %s", err.Error())
			return
		}
		log.Printf("Download Video Success: %s", fileName)
	}
}

//Using SOCK5 to do request
func Request(url string) {
	client := tw.NewClient()
	_, err := client.ClientWithSOCKS5("tcp", "127.0.0.1:1080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = client.GetWithProxy(url)
	if err != nil {
		fmt.Println("Exit", err.Error())
	}
}
