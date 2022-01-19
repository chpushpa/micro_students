package students

import (
	"github.com/chpushpa/micro_student/datasource/mysql/users_db"
	"github.com/chpushpa/micro_student/domain/httperrors"
)

const (
	queryInsertUser = "INSERT INTO users(id, name, grade, sec, class) VALUES (?, ?, ?, ?, ?)"
)

func (std *Student) Save() *httperrors.HttpError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return httperrors.NewInternalServeError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(std.Id, std.Name, std.Grade, std.Sec, std.Class)
	if saveErr != nil {
		return httperrors.NewInternalServeError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return httperrors.NewInternalServeError("database error")
	}
	std.Id = userId
	return nil
}
