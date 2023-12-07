package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// リクエスト処理ライフサイクルのレンダリングレスポンスフェーズと
// 復元ビューフェーズでアクティビティを独自に処理できるようにするためのプラグ可能メカニズム
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

//ファイルの読み込みは可能だが　エラー処理ができていない
// func loadPage(title string) *Page {
//     filename := title + ".txt"
//     body, _ := os.ReadFile(filename)
//     return &Page{Title: title, Body: body}
// }

// 　エラー処理
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
	// func ReadFile(name string) ([]byte, error)　\
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

