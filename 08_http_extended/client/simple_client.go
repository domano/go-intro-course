package main

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)


type User struct {
	Name string

}
func main() {

	http.Get("http://tarent.de/")

	body, _:= json.Marshal(&User{"Dino"})
	http.Post("http://localhost:10000", "application/json", bytes.NewReader(body))
}
