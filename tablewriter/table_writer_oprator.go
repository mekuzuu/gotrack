package tablewriter

type ITableWriterOperator interface {
	Write(param *TableWriterParameter)
}
