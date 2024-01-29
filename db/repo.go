package db

import (
	"TodoBackend/model"
	"database/sql"
	"fmt"
	"time"
)

func (d *MysqlDB) Save(req model.Task) (model.Task, error) {
	now := time.Now()

	res, err := d.db.Exec("insert into todo (title,description,done,created_at) values (?,?,?,?)", req.Title, req.Description, false, now)
	if err != nil {
		return model.Task{}, fmt.Errorf("can't execute command: %w", err)
	}

	id, _ := res.LastInsertId()

	return model.Task{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		Done:        false,
		CreatedAt:   now,
	}, nil

}

func (d *MysqlDB) GetByID(id int64) (model.Task, error) {
	row := d.db.QueryRow("select * from todo where id = ?", id)
	task, err := scanRecord(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Task{}, fmt.Errorf("not found: %w", err)
		}

		return model.Task{}, fmt.Errorf("can't execute command: %w", err)
	}
	return task, nil
}

func scanRecord(row *sql.Row) (model.Task, error) {
	var createdAt []uint8
	var task model.Task

	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Done, &createdAt)
	return task, err

}
