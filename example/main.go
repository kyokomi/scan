package main

import (
	"fmt"

	"github.com/kyokomi/scan"
	"os"
	"log"
)

func main() {
	c := scan.CliScan{Scans: []scan.Scan{
		{
			Name:  "hoge",
			Value: "fuga",
			Usage: "input pleese [hoge or fuga]",
			Env: "HOGE",
		},
	}}

	fmt.Println(c.Get("hoge"))

	if err := os.Setenv("HOGE", "test"); err != nil {
		log.Fatalln(err)
	}
	c.Reset("hoge")

	fmt.Println(c.Get("hoge"))
	fmt.Println("scan => ", c.Scan("hoge"))
	fmt.Println(c.Get("hoge"))
}
