// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const starExpansionReserved = `-- name: StarExpansionReserved :many
SELECT "group" FROM foo
`

func (q *Queries) StarExpansionReserved(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.db.QueryContext(ctx, starExpansionReserved)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var group sql.NullString
		if err := rows.Scan(&group); err != nil {
			return nil, err
		}
		items = append(items, group)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
