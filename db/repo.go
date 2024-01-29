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
		CreatedAt:   &now,
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

func (d *MysqlDB) GetAll() ([]model.Task, error) {
	rows, dErr := d.db.Query("select * from todo")

	if dErr != nil {
		return nil, fmt.Errorf("can't execute command: %w", dErr)
	}

	records, dErr := scanRecords(rows)
	if len(records) == 0 {
		return nil, fmt.Errorf("not found: %w", dErr)
	}
	return records, nil
}

func (d *MysqlDB) Update(req model.Task) error {
	_, err := d.db.Exec("update todo set title=?, description=?, done=? where id=?", req.Title, req.Description, req.Done, req.Id)
	if err != nil {
		return fmt.Errorf("can't execute command: %w", err)
	}
	return nil
}

func (d *MysqlDB) DeleteByID(id int64) error {
	_, err := d.db.Exec("delete from todo where id= ?", id)
	if err != nil {
		return fmt.Errorf("can't execute command: %w", err)
	}
	return nil
}

func scanRecord(row *sql.Row) (model.Task, error) {
	var task model.Task

	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Done, &task.CreatedAt)
	return task, err
}

func scanRecords(rows *sql.Rows) ([]model.Task, error) {
	var tasks []model.Task
	var task model.Task

	for rows.Next() {
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Done, &task.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
