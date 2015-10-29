package spider

import (
    //"net"
    "net/http"
    "io/ioutil"
    "bytes"
)

type Downloader struct {
    
}

func Get(url string) (content string, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    content = string(body)
    return
}

func Post(uri, postdata string) (content string, err error) {
    buf := []byte(postdata)
    resp, err := http.Post(uri, "application/json", bytes.NewReader(buf))
    if err != nil {
        return "", err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    content = string(body)
    return
}

//Regexp FindAll(b []byte, n int) [][]byte
//Regexp FindAllString(s string, n int)[]string


