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

type TableWriter struct {
	Header []string
	Data   [][]string
}

func (t *tableWriter) Write(table *TableWriter) {
	t.table.SetHeader(table.Header)
	for _, v := range table.Data {
		t.table.Append(v)
	}
	t.table.SetAutoMergeCells(true)
	t.table.Render()
}
