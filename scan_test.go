package scan

import (
	"testing"
	"os"
)

func TestGet(t *testing.T) {
	c := CliScan{Scans: []Scan{
		{
			Name:  "hoge",
			Value: "fuga",
			Usage: "input pleese [hoge or fuga]",
			Env: "",
		},
	}}

	if c.Get("hoge") != "fuga" {
		t.Error("get is not fuga")
	}
}

func TestEnvGet(t *testing.T) {
	c := CliScan{Scans: []Scan{
		{
			Name:  "hoge",
			Value: "",
			Usage: "input pleese [hoge or fuga]",
			Env: "HOGE",
		},
	}}

	if err := os.Setenv("HOGE", "test"); err != nil {
		t.Error(err)
	}

	if c.Get("hoge") != "test" {
		t.Error("get is not test", c.Get("test"))
	}
}
