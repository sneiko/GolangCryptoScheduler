package db

type DbTable struct {
	Table string
}

func NewDbTable(table string) DbTable {
	return DbTable{
		Table: table,
	}
}
