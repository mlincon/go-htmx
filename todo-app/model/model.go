package model

type Todo struct {
	ID   uint64 `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func CreateTodo(todo string) error {
	stmt := `INSERT INTO todo.todos(todo, done) VALUES($1, $2);`
	_, err := db.Query(stmt, todo, false)
	return err
}

func GetAllTodos() ([]Todo, error) {
	todos := make([]Todo, 0)

	stmt := `SELECT id, todo, done FROM todo.todos;`

	// execute the query and return results
	rows, err := db.Query(stmt)
	if err != nil {
		return todos, err
	}
	defer rows.Close()

	// a cursor is returned, so iterate over and fetch the values
	for rows.Next() {
		var (
			id    uint64
			title string
			done  bool
		)

		if err := rows.Scan(&id, &title, &done); err != nil {
			return todos, err
		}
		todo := Todo{
			ID:   id,
			Todo: title,
			Done: done,
		}
		todos = append(todos, todo)
	}
	return todos, err
}

func GetTodo(id uint64) (Todo, error) {
	stmt := `SELECT todo, done FROM todo.todos WHERE id=$1;`

	todo := Todo{}

	row, err := db.Query(stmt, id)
	if err != nil {
		return todo, err
	}
	defer row.Close()

	for row.Next() {
		var (
			title string
			done  bool
		)

		if err := row.Scan(&title, &done); err != nil {
			return todo, err
		}

		todo = Todo{
			ID:   id,
			Todo: title,
			Done: done,
		}
	}
	return todo, err
}

func ChangeDoneStatus(id uint64) error {
	todo, err := GetTodo(id)
	if err != nil {
		return err
	}

	stmt := `UPDATE todo.todos SET done=$2 WHERE id=$1;`
	_, err = db.Query(stmt, id, !todo.Done)

	return err
}

func Delete(id uint64) error {
	stmt := `DELETE FROM todo.todos WHERE id=$1;`

	_, err := db.Exec(stmt, id)
	return err
}
