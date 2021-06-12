package sagawa

type ISagawaOperator interface {
	TrackShipment(id string) error
}
