package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "ゲスト"
    }
    length := len([]rune(name))
    score := length * 10
    fmt.Fprintf(w, "こんにちは、%s さん！<br>名前の文字数 × 10 = %d", name, score)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("サーバー起動: http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
