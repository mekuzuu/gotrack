package tablewriter

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

type TableWriterParameter struct {
	Header                  []string
	Data                    [][]string
	EnableMergeCells        bool
	MergeCellsByColumnIndex []int
}

type tableWriterOperator struct {
	table *tablewriter.Table
}

func NewTableWriterOperator(os io.Writer) *tableWriterOperator {
	return &tableWriterOperator{
		table: tablewriter.NewWriter(os),
	}
}

func (t *tableWriterOperator) Write(param *TableWriterParameter) {
	t.table.SetHeader(param.Header)
	for _, v := range param.Data {
		t.table.Append(v)
	}
	t.setOptions(param)
	t.table.Render()
}

func (t *tableWriterOperator) setOptions(param *TableWriterParameter) {
	// 何も指定しなくともtrueを設定されてしまうため、指定がある場合のみオプションを設定するようにしている
	if len(param.MergeCellsByColumnIndex) > 0 {
		t.table.SetAutoMergeCellsByColumnIndex(param.MergeCellsByColumnIndex)
	}
}
