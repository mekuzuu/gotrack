package main

import (
	"os"

	"gotrack/courie/yamato"
	"gotrack/tablewriter"
)

func main() {
	t := tablewriter.NewTableWriter(os.Stdout)
	y := yamato.NewYamato()
	table, err := y.FindShipmentsTable([]string{})
	if err != nil {
		panic(err)
	}
	t.Write(table)
}
