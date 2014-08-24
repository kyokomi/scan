// Package scan terminal input scan golang package
package scan

import "fmt"

// CliScan scans.
type CliScan struct {
	Scans []Scan
}

// Scan data.
type Scan struct {
	Name  string
	Value string
	Usage string
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
	return s.Value
}

func (c *CliScan) findScan(name string) *Scan {
	var s Scan
	for _, s = range c.Scans {
		if s.Name == name {
			break
		}
	}

	if s.Name != name {
		return nil
	}

	return &s
}
