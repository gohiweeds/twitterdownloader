package twitterdownloader

import (
  "net/http"
  "os"
  "net/url"
  "fmt"
  "golang.org/x/net/proxy"
)

const (
  DIRECT = iota
  SOCKS5
  PROXY
)
///Refer to https://www.cnblogs.com/baicaiyoudu/p/5929138.html
//Set URL Proxy
func (c *Client)ClientWithProxy(proxyUrl string) (*http.Client, error) {
  u := url.URL{}
  urlProxy, err := u.Parse(proxyUrl)

  if err != nil {
    return nil, err
  }

  client := &http.Client{
    Transport: &http.Transport{
      Proxy: http.ProxyURL(urlProxy),
    },
  }

  c.client = client
  c.proxyUrl = proxyUrl
  c.proxyType = PROXY
  return client, nil
}
//Set Environment
func (c Client)SetProxyEnvironment(httpUrl, httpsUrl string){
  os.Setenv("HTTP_PROXY", httpUrl)
  os.Setenv("HTTPS_PROXY", httpsUrl)
}

//Set SOCKS
func (c *Client)ClientWithSOCKS5(proto, ipport string) (*http.Client, error){
  //create a socks5 dialer
  dialer, err := proxy.SOCKS5(proto, ipport, nil, proxy.Direct)
  if err != nil {
    fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
    os.Exit(1)
  }

  //setup a http client
  httpTransport := &http.Transport{}
  httpClient := &http.Client{
    Transport: httpTransport,
  }

  //set our socks5 as the dialer
  httpTransport.Dial = dialer.Dial

  c.client = httpClient
  c.socks5Proto = proto
  c.socks5IpPort = ipport
  c.proxyType = SOCKS5
  return httpClient, nil
}
