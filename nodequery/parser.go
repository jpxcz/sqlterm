package nodequery

import (
	"database/sql"

	"github.com/charmbracelet/bubbles/table"
)

type QueryResult struct {
	query   string
	columns []string
	rows    [][]string
}

func NewQueryResult(query string, rows *sql.Rows) (*QueryResult, error) {
	queryResult := &QueryResult{
		query:   query,
		columns: nil,
		rows:    nil,
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	queryResult.columns = columns

	rawResult := make([][]byte, len(columns))
	result := make([]string, len(columns))

	dest := make([]interface{}, len(columns))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, err
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\n"
			} else {
				result[i] = string(raw)
			}
		}
		queryResult.rows = append(queryResult.rows, result)
	}

	return queryResult, nil
}

func (q *QueryResult) GetColumnsViewData() []table.Column {
	columns := []table.Column{}
	for _, col := range q.columns {
		columns = append(columns, table.Column{Title: col, Width: 20})
	}

	return columns
}

func (q *QueryResult) GetRowsViewData() []table.Row {
    rows := []table.Row{}
    for _, row := range q.rows {
        r := table.Row(row)
        rows = append(rows, r)
    }

    return rows
}
