package tablewriter

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

type tableWriter struct {
	table *tablewriter.Table
}

func NewTableWriter(w io.Writer) *tableWriter {
	return &tableWriter{
		table: tablewriter.NewWriter(w),
	}
}

func (t *tableWriter) Write(header []string, data [][]string) {
	t.table.SetHeader(header)
	for _, v := range data {
		t.table.Append(v)
	}
	t.table.Render()
}
