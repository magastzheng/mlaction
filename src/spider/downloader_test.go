package spider_test

import (
    "spider"
    "testing"
    "strings"
    "fmt"
    "regexp"
)

func Test_Spider_Get(t *testing.T) {
    content,err := spider.Get("http://www.baidu.com")
    if err != nil {
        fmt.Println("Cannot download")
    }
    //fmt.Println(content)
    
    pattern := "<[a|A]\\s+href=([^>]*\\s*>)"
    re := regexp.MustCompile(pattern)
    urls := re.FindAllString(content, -1)
    fmt.Println(urls)
    Parse_Spider_ParseUrl(urls)
}


func Parse_Spider_ParseUrl(urls []string) {
    prefix := "href="
    suffix := "\""
    for _, url := range urls {
        pos := strings.Index(url, prefix)
        if pos < 0 {
            continue
        }
        
        url = string(url[pos+len(prefix)+1:])
        fmt.Println("Filter prefix: ", url)
        pos = strings.Index(url, suffix)
        url = string(url[:pos])
        fmt.Println("Filter suffix: ", url)
    }
}
