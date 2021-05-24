package main

import (
	"flag"
	"os"

	"gotrack/courie/yamato"
	"gotrack/tablewriter"
)

var (
	// Dependency
	tableWriterOP tablewriter.ITableWriterOperator
	yamatoOP      yamato.IYamatoOperator

	flagYamato = flag.String("ymt", "", "track shipment transported by Yamato Transport")
)

func init() {
	// Dependency
	tableWriterOP = tablewriter.NewTableWriterOperator(os.Stdout)
	yamatoOP = yamato.NewYamatoOperator(tableWriterOP)
}

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	if *flagYamato != "" {
		if err := yamatoOP.TrackShipments([]string{*flagYamato}); err != nil {
			return 1
		}
	}
	return 0
}
