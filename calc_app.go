package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("使い方: go run calc_app.go [数値1] [数値2] ...")
        return
    }

    sum := 0
    for _, arg := range os.Args[1:] {
        num, err := strconv.Atoi(arg)
        if err != nil {
            fmt.Printf("数値を入力してください: %s\n", arg)
            return
        }
        sum += num
    }

    fmt.Printf("合計: %d\n", sum)
}
