package main

import (
	"fmt"

	tw "github.com/gohiweeds/twitterdownloader"
)

func main() {
	//c := &tw.Client{}
	//c.Get("http://www.baidu.com/xxx")

	//Picture
	//err = client.GetWithProxy("https://pbs.twimg.com/media/Dlldrd-U4AAmzxJ.jpg:large")
	//Video
	//err = client.GetWithProxy("https://twitter.com/i/status/1033468719001989121")
	//err = client.GetWithProxy("https://twitter.com/i/status/1033716646911729665")

	//err = client.GetWithProxy("https://twitter.com/i/videos/1033716646911729665?embed_source=facebook")
	//err = client.GetWithProxy("https://api.twitter.com/1.1/videos/tweet/config/1033716646911729665.json")
	//err = client.GetWithProxy("https://abs.twimg.com/web-video-player/TwitterVideoPlayerIframe.9f3fa50c5fbf9f33.js")

	//err := client.Get("https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1534863418726&di=d6ea6c86f5bb6dac00c8e8ad25bd0eea&imgtype=0&src=http%3A%2F%2Ftu.qiumibao.com%2Fuploads%2Fday_180814%2F201808142247577825.jpg")
	//https://twitter.com/i/status/1035056498307522560
	//https://twitter.com/i/videos/tweet/

	//Request("https://twitter.com/i/videos/tweet/1035056498307522560")

	//Need authenticate
	//Request("https://api.twitter.com/1.1/videos/tweet/config/1035056498307522560.json")

	// client := &tw.Client{}
	// _, err := client.ClientWithSOCKS5("tcp", "127.0.0.1:1080")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	//
	// twitter := &tw.Twitter{}
	// twitter.Activate(client)
	// twitter.Activate2(client)
	// twitter.GetVideoJson(client)
	// twitter.GetVideoJson2(client)
	// twitter.Getm3u8List(client)
	//Request("https://video.twimg.com/ext_tw_video/1035056301657583617/pu/pl/540x960/Vyiq7upZ9HjHDruX.m3u8")
	Request("https://video.twimg.com/ext_tw_video/1035056301657583617/pu/vid/0/3000/540x960/DJ0xXa4TpDDejYlV.ts")
	Request("https://video.twimg.com/ext_tw_video/1035056301657583617/pu/vid/3000/6000/540x960/Ij-0cqTubNDJbd1E.ts")
	Request("https://video.twimg.com/ext_tw_video/1035056301657583617/pu/vid/6000/10043/540x960/kljEHBu_sl5nP90X.ts")
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
