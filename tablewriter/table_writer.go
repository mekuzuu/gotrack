package tablewriter

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

type tableWriterOperator struct {
	table *tablewriter.Table
}

func NewTableWriterOperator(os io.Writer) *tableWriterOperator {
	return &tableWriterOperator{
		table: tablewriter.NewWriter(os),
	}
}

func (t *tableWriterOperator) Write(header []string, data [][]string) {
	t.table.SetHeader(header)
	for _, v := range data {
		t.table.Append(v)
	}

	t.table.SetAutoMergeCells(true)
	t.table.Render()
}
