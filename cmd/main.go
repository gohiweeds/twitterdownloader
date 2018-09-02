package main

import (
	"flag"
	"fmt"
	"strings"

	tw "github.com/gohiweeds/twitterdownloader"
)

func main() {
	//client := &tw.Client{}
	//client.Init()
	// _, err := client.ClientWithSOCKS5("tcp", "127.0.0.1:1080")
	// _, err := client.ClientWithSOCKS5("tcp", "127.0.0.1:1087")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	url := flag.String("url", "", "download provided url")

	flag.Parse()

	if *url == "" {
		fmt.Println("URL not qualified or exist")
		return
	}
	client := &tw.Client{}

	twitter := &tw.Twitter{}
	twitter.SetupClient(client)
	//twitter.DownloadVideo("https://twitter.com/i/status/1035056498307522560")

	if strings.Contains(*url, ".jpg") {
		twitter.DownloadJPG(*url)
	} else {
		//https://pbs.twimg.com/media/Dl_0nDtXoAAZu4x.jpg:large
		twitter.DownloadVideo(*url)
	}
}

func Request(url string) {
	client := &tw.Client{}
	_, err := client.ClientWithSOCKS5("tcp", "127.0.0.1:1080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = client.GetWithProxy(url)
	if err != nil {
		fmt.Println("Exit", err.Error())
	}
}
