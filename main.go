package main

import "gotrack/courie/yamato"

func main() {
	y := yamato.NewYamato()
	y.FindShipments([]string{"397006850170", "397006850170"})
}
