package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func DeleteMany(db *sql.DB, ctx context.Context, tableName string, ids []string) error {
	query := fmt.Sprintf("DELETE FROM %s where id in (", tableName)
	idStrings := make([]string, 0, len(ids))
	for i := range ids {
		idStrings = append(idStrings, fmt.Sprintf("$%d", i+1))
	}
	query += strings.Join(idStrings, ",")
	query += ")"
	fmt.Printf("query: \n%s\n", query)
	// Pass the ids slice directly as variadic arguments
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = interface{}(id)
	}
	_, err := db.ExecContext(ctx, query, args...)
	return err
}
