package jp

type IJPOperator interface {
	TrackShipment(id string) error
}
