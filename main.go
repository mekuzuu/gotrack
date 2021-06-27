package main

import (
	"flag"
	"fmt"
	"os"

	"gotrack/courie/jp"
	"gotrack/courie/sagawa"
	"gotrack/courie/yamato"
	"gotrack/tablewriter"
)

var (
	// Dependencies, registered during Init.
	tableWriterOP tablewriter.ITableWriterOperator
	yamatoOP      yamato.IYamatoOperator
	sagawaOP      sagawa.ISagawaOperator
	jpOP          jp.IJPOperator

	// Flags, registered during Init.
	flagYamato *string
	flagSagawa *string
	flagJP     *string
)

func init() {
	// Dependencies.
	tableWriterOP = tablewriter.NewTableWriterOperator(os.Stdout)
	yamatoOP = yamato.NewYamatoOperator(tableWriterOP)
	sagawaOP = sagawa.NewSagawaOperator(tableWriterOP)
	jpOP = jp.NewJPOperator(tableWriterOP)

	// Flags.
	flagYamato = flag.String("ymt", "", "track shipment transported by Yamato Transport")
	flagSagawa = flag.String("sgw", "", "track shipment transported by Sagawa Express")
	flagJP = flag.String("jp", "", "track shipment transported by Japan Post")
}

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()

	if *flagYamato != "" {
		fmt.Println("track shipment transported by Yamato Transport")
		if err := yamatoOP.TrackShipments(*flagYamato); err != nil {
			fmt.Println(err.Error())
			return 1
		}
	}

	if *flagSagawa != "" {
		fmt.Println(flagSagawa)
		fmt.Println("track shipment transported by Sagawa Express")
		if err := sagawaOP.TrackShipment(*flagSagawa); err != nil {
			fmt.Println(err.Error())
			return 1
		}
	}

	if *flagJP != "" {
		fmt.Println("track shipment transported by Japan Post")
		if err := jpOP.TrackShipment(*flagJP); err != nil {
			fmt.Println(err.Error())
			return 1
		}
	}

	if err := jpOP.TrackShipment("ET700332656VN"); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	return 0
}
