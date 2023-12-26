package connections

import (
	"fmt"
	"sync"
)

// GetSchema returns the schema of the database.
func (p *connection) getFullDbSchema() (map[string]interface{}, error) {
	tables, err := p.getTables()
	if err != nil {
		return nil, err
	}
	
	var wg sync.WaitGroup
	var data = make(chan map[string]interface{})
	for _, table := range tables {
		wg.Add(1)
		go func(tableName string) {
			defer wg.Done()
			columns, err := p.getColumns(tableName)
			if err != nil {
				return
			}
			constraints, err := p.getConstraints(tableName)
			if err != nil {
				return
			}

			data <- map[string]interface{}{
				tableName: map[string]interface{}{
					"columns":     columns,
					"constraints": constraints,
				},
			}
		}(table)
	}

	go func() {
		wg.Wait()
		close(data)
	}()

	schema := make(map[string]interface{})
	for table := range data {
		for tableName, tableData := range table {
			schema[tableName] = tableData
		}
	}
	return schema, nil
}

// getTables returns a list of tables in the database.
func (p *connection) getTables() ([]string, error) {
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'"
	rows, err := p.cnx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tables, nil
}

// getColumns returns a map of columns in the given table.
func (p *connection) getColumns(tableName string) (map[string]string, error) {
	query := fmt.Sprintf("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = '%s'", tableName)
	rows, err := p.cnx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := make(map[string]string)
	for rows.Next() {
		var columnName string
		var dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			return nil, err
		}
		columns[columnName] = dataType
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return columns, nil
}

// getConstraints returns a map of constraints in the given table.
func (p *connection) getConstraints(tableName string) (map[string]string, error) {
	query := fmt.Sprintf("SELECT constraint_name, constraint_type FROM information_schema.table_constraints WHERE table_name = '%s'", tableName)
	rows, err := p.cnx.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	constraints := make(map[string]string)
	for rows.Next() {
		var constraintName string
		var constraintType string
		if err := rows.Scan(&constraintName, &constraintType); err != nil {
			return nil, err
		}
		constraints[constraintName] = constraintType
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return constraints, nil
}
