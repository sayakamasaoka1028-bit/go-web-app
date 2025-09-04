package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("使い方: go run calc_app2.go [add|sub|mul|div] 数値1 数値2 ...")
        return
    }

    op := os.Args[1]
    nums := os.Args[2:]

    result := 0
    for i, arg := range nums {
        num, err := strconv.Atoi(arg)
        if err != nil {
            fmt.Printf("数値を入力してください: %s\n", arg)
            return
        }
        if i == 0 {
            result = num
            continue
        }
        switch op {
        case "add":
            result += num
        case "sub":
            result -= num
        case "mul":
            result *= num
        case "div":
            result /= num
        default:
            fmt.Println("演算は add, sub, mul, div を指定してください")
            return
        }
    }

    fmt.Printf("結果: %d\n", result)
}
