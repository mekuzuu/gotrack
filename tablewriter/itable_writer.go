package tablewriter

type ITableWriter interface {
	Write(table *TableWriter)
}