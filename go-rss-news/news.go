package main

import (
    "fmt"
    "log"

    "github.com/mmcdole/gofeed"
)

func main() {
    rssURL := "https://www3.nhk.or.jp/rss/news/cat0.xml"

    fp := gofeed.NewParser()
    feed, err := fp.ParseURL(rssURL)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("ニュース見出し一覧:")
    for i, item := range feed.Items {
        if i >= 5 { // 上位5件だけ表示
            break
        }
        fmt.Printf("%d: %s\n   URL: %s\n\n", i+1, item.Title, item.Link)
    }
}
