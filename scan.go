// Package scan terminal input scan golang package
package scan

import (
	"fmt"
	"os"
)

// CliScan scans.
type CliScan struct {
	Scans []Scan
}

// Scan data.
type Scan struct {
	Name  string
	Value string
	Usage string
	Env   string
}

// Scan terminal scan.
func (c *CliScan) Scan(name string) string {
	s := c.findScan(name)

	fmt.Print(s.Usage, ": ")
	fmt.Scanln(&s.Value)
	return s.Value
}

// Get result scan value.
func (c *CliScan) Get(name string) string {
	s := c.findScan(name)
	if s.Value == "" {
		s.Value = os.Getenv(s.Env)
	}
	return s.Value
}

func (c *CliScan) Reset(name string) {
	c.findScan(name).Value = ""
}

func (c *CliScan) findScan(name string) *Scan {
	for i := 0; i < len(c.Scans); i ++ {
		if name != c.Scans[i].Name {
			continue
		}
		return &c.Scans[i]
	}
	return nil
}
