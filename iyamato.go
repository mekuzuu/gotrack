package main

type IYamato interface {
	FindShipments(ids []string) ([]Shipment, error)
}
