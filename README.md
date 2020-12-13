# Create signed URLs 

[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/i-pva/signed-url/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/i-pva/signed-url)](https://goreportcard.com/report/github.com/i-pva/signed-url)

This package can create URLs by adding an expiration date and a signature to the URL.
Signed URLs are especially useful for routes that are publicly accessible yet need a layer of protection against URL manipulation.

# Installation

```bash
go get github.com/i-pva/signed-url
```
 
# Example

```go
package main

import (
    "log"
    "net/http"
    "net/url"
    
    surl "github.com/i-pva/signed-url"
)

func main(){

	surl.SecretKey = []byte("your-secret-key") //rewrite secret key


	u := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "path",
	}

	//URL Signing

	signedUrl , err := surl.Signed(u)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(signedUrl.String()) // https://example.com/path?signature=XXX
	

	// URL Validating

	if surl.HasValidURL(&http.Request{URL: u}) {
		log.Println("URL is valid")
		// do something
	}

}
```

# URLs Signing with a limited lifetime  

```go
signedUrl , err := url.TemporarySigned(u, 1 * time.Hour) // will be valid for 1 hour
if err != nil {
     return err
}

log.Println(signedUrl.String()) // https://example.com/path?expires=XXX&signature=XXX
```

# URLs Validating with handler
```go

mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("route under signed url"))
})
 
handler := surl.Handler(mux) // handler for validating signed URLs
http.ListenAndServe(":", handler)

```

# [License](LICENSE)
This package is released under the MIT license.