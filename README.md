# Twitter Downloader

Support jpeg and video download in Twitter


* Image - Get the URL and save it to disk is fine.

* Video - Get the Video Clip in Twitter website

### How we know where is the Video ?

Request: issue by xhr
```
OPTIONS https://api.twitter.com/1.1/guest/activate.json
Accept:*/*
Accept-Encoding:gzip,deflate,br
Access-Control-Request-Headers:authorization,x-csrf-token
Access-Control-Request-Method:POST
Origin: https://twitter.com
User-Agent:Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/65.0.3325.181 Chrome/65.0.3325.181 Safari/537.36
```
Response
```
access-control-allow-credentials: true
access-control-allow-headers: Authorization, Cache-Control, Content-Type, Dtab-Local, If-Modified-Since, LivePipeline-Session, Server, X-Act-As-User-Id, X-B3-Flags, X-CSRF-Token, X-Contribute-To-User-Id, X-Guest-Token, X-TD-Iff-Mtime, X-TD-Mtime-Check, X-Twitter-Active-User, X-Twitter-Auth-Type, X-Twitter-Client, X-Twitter-Client-Language, X-Twitter-Client-Version, X-Twitter-Polling, X-Twitter-UTCOffset
access-control-allow-methods: GET, POST, HEAD, PUT, DELETE
access-control-allow-origin: https://twitter.com
access-control-max-age: 1728000
content-length: 0
date: Fri, 31 Aug 2018 06:06:16 GMT
server: tsa_a
status: 200
vary: Origin
x-connection-hash: 5c8af79361490fa357defce397a007d3
```

Request- activate.json
```
:authority: api.twitter.com
:method: POST
:path: /1.1/guest/activate.json
:scheme: https
accept: */*
accept-encoding: gzip, deflate, br
accept-language: zh-CN,zh;q=0.9,ar;q=0.8
authorization: Bearer AAAAAAAAAAAAAAAAAAAAAIK1zgAAAAAA2tUWuhGZ2JceoId5GwYWU5GspY4%3DUq7gzFoCZs1QfwGoVdvSac3IniczZEYXIcDyumCauIXpcAPorE
content-length: 0
cookie: personalization_id="v1_pid1UUVchOmH31FJFT2ZLA=="; guest_id=v1%3A153569557385510737
origin: https://twitter.com
referer: https://twitter.com/i/videos/tweet/1035056498307522560
user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/65.0.3325.181 Chrome/65.0.3325.181 Safari/537.36
x-csrf-token: undefined
```

Response -activate.json
```
access-control-allow-credentials: true
access-control-allow-origin: https://twitter.com
cache-control: no-cache, no-store, must-revalidate, pre-check=0, post-check=0
content-disposition: attachment; filename=json.json
content-encoding: gzip
content-length: 63
content-type: application/json;charset=utf-8
date: Fri, 31 Aug 2018 06:06:18 GMT
expires: Tue, 31 Mar 1981 05:00:00 GMT
last-modified: Fri, 31 Aug 2018 06:06:18 GMT
pragma: no-cache
server: tsa_a
status: 200 OK
status: 200
strict-transport-security: max-age=631138519
x-access-level: read
x-connection-hash: df213f7d0fdbde0c19ef76e16fb50e71
x-content-type-options: nosniff
x-frame-options: SAMEORIGIN
x-response-time: 13
x-transaction: 00c5dcb00038ee42
x-twitter-response-tags: BouncerCompliant
x-xss-protection: 1; mode=block; report=https://twitter.com/i/xss_report
```
