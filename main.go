package main

import (
	"os"

	"gotrack/courie/yamato"
	"gotrack/tablewriter"
)

func main() {
	t := tablewriter.NewTableWriter(os.Stdout)
	y := yamato.NewYamato()
	table, err := y.FindShipmentsTable([]string{"397006850170", "397006850170"})
	if err != nil {
		panic(err)
	}
	t.Write(table)
}
