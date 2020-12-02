package db

import "database/sql"

var (
	mockDB map[uint64] *Row
)

type Row struct {
	Col1 string

	Col2 string
}

func init()  {
	mockDB = make(map[uint64]*Row)

	mockDB[1] = &Row{Col1: "bob", Col2: "123"}
	mockDB[2] = &Row{Col1: "alice", Col2: "123"}
	mockDB[3] = &Row{Col1: "wang", Col2: "123"}
}

func GetByPrimary(id uint64) (*Row, error)  {
	if v, ok := mockDB[id];ok {
		return v, nil
	}

	//return error
	return nil, sql.ErrNoRows
}
