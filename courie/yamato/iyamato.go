package yamato

import "gotrack/shipment"

type IYamato interface {
	FindShipments(ids []string) ([]shipment.Shipment, error)
}
