package tablewriter

type ITableWriter interface {
	Write(heder []string, data [][]string)
}
