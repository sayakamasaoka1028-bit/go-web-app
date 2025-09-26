# go-rss-news

NHKニュースのRSSを取得して表示するGoアプリです。

## 機能
- NHKニュースのRSSから最新のニュース見出しを取得
- 各ニュースのリンクも表示
- 表示件数を制限可能（例：最新5件のみ）

## 使い方
1. Go がインストールされていることを確認
2. リポジトリをクローン
   ```bash
   git clone https://github.com/sayakamasaoka1028-bit/go-rss-news.git
   cd go-rss-news
   ```
 
   ```bash
   go get github.com/mmcdole/gofeed
   ```
   ```bash
   go run news.go
   ```

```gitignone
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# IDE/editor config files
.idea/
.vscode/
*.swp
```
