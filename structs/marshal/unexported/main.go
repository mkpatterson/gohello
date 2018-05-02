package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string `json:"-"`
	Age int `json:"Wisdom Score"`
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

}