package tablewriter

type ITableWriterOperator interface {
	Write(header []string, data [][]string)
}
