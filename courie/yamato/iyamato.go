package yamato

import "gotrack/tablewriter"

type IYamato interface {
	FindShipmentsTable(ids []string) (*tablewriter.TableWriterModel, error)
}
