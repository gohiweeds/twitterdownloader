package twitterdownloader

import "fmt"

func downloader(video string, c *Client) error {
	client := c
	fmt.Println("client:", client)
	client.RetriveVideoConfig(video)
	return nil
}
