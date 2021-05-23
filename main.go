package main

import (
	"os"

	"gotrack/courie/yamato"
	"gotrack/tablewriter"
)

func main() {
	tableWriterOP := tablewriter.NewTableWriterOperator(os.Stdout)
	yamatoOP := yamato.NewYamatoOperator(tableWriterOP)
	err := yamatoOP.TrackShipments([]string{})
	if err != nil {
		panic(err)
	}
}
