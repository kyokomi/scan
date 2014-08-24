package main

import (
	"fmt"

	"github.com/kyokomi/scan"
)

func main() {
	c := scan.CliScan{Scans: []scan.Scan{
		{
			Name:  "hoge",
			Value: "fuga",
			Usage: "input pleese [hoge or fuga]",
		},
	}}

	fmt.Println(c.Get("hoge"))
	fmt.Println(c.Scan("hoge"))
	fmt.Println(c.Get("hoge"))
}
