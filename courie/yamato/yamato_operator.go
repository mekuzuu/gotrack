package yamato

type IYamatoOperator interface {
	TrackShipments(ids []string) error
}
