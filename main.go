package main

import (
	"flag"
	"fmt"
	"os"

	"gotrack/courie/sagawa"
	"gotrack/courie/yamato"
	"gotrack/tablewriter"
)

var (
	// Dependency
	tableWriterOP tablewriter.ITableWriterOperator
	yamatoOP      yamato.IYamatoOperator
	sagawaOP      sagawa.ISagawaOperator

	flagYamato = flag.String("ymt", "", "track shipment transported by Yamato Transport")
	flagSagawa = flag.String("sgw", "", "track shipment transported by Sagawa Express")
)

func init() {
	// Dependency
	tableWriterOP = tablewriter.NewTableWriterOperator(os.Stdout)
	yamatoOP = yamato.NewYamatoOperator(tableWriterOP)
	sagawaOP = sagawa.NewSagawaOperator(tableWriterOP)
}

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	if *flagYamato != "" {
		if err := yamatoOP.TrackShipments([]string{*flagYamato}); err != nil {
			fmt.Println(err.Error())
			return 1
		}
	}

	if *flagSagawa != "" {
		if err := sagawaOP.TrackShipment(*flagSagawa); err != nil {
			fmt.Println(err.Error())
			return 1
		}
	}

	return 0
}
