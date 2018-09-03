# Twitter Downloader

Support jpeg and video download in Twitter


* Image - Get the URL and save it to disk is fine.

* Video - Get the Video Clip in Twitter website

# Usage

  Install dependency:
  ```
  go get github.com/grafov/m3u8
  ```

  Run Test command:
  ```
  go run cmd/main.go -url=https://twitter.com/i/status/1035056498307522560
  ```

### How we know where is the Video ?
  * provide tweet url which contains video, like `https://twitter.com/i/status/1035056498307522560`
  * request activate.json twice get guest token
  * use guest token request tweet_id.json twice get video config
  * request video url to get mp4 file, or m3u8 to get sub-m3u8 or video clips
