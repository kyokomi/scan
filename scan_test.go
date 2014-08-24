package scan

import "testing"

func TestGet(t *testing.T) {
	c := CliScan{Scans: []Scan{
		{
			Name:  "hoge",
			Value: "fuga",
			Usage: "input pleese [hoge or fuga]",
		},
	}}

	if c.Get("hoge") != "fuga" {
		t.Error("get is not fuga")
	}
}
