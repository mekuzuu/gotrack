package yamato

type IYamatoOperator interface {
	TrackShipments(id string) error
}
