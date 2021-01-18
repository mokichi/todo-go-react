package usecase

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"todo/internal/entity"
	"todo/internal/orm"
)

func TestFindByCompleted(t *testing.T) {
	db, mock := orm.NewDBMock()
	u := TasksUsecase{DB: db}

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "tasks" WHERE completed = $1 ORDER BY updated_at DESC`)).
		WithArgs(false)
	u.FindByCompleted(false)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFirst(t *testing.T) {
	db, mock := orm.NewDBMock()
	u := TasksUsecase{DB: db}

	id := 1
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "tasks" WHERE "tasks"."id" = $1 ORDER BY "tasks"."id" LIMIT 1`)).
		WithArgs(id)
	u.First(id)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestCreate(t *testing.T) {
	db, mock := orm.NewDBMock()
	u := TasksUsecase{DB: db}

	task := entity.Task{Title: "test"}
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "tasks" ("title","completed","created_at","updated_at") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(task.Title, task.Completed, orm.AnyMockArg(), orm.AnyMockArg())
	u.Create(&task)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	db, mock := orm.NewDBMock()
	u := TasksUsecase{DB: db}

	task := entity.Task{ID: 1, Title: "test"}
	mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "tasks" SET "title"=$1,"completed"=$2,"created_at"=$3,"updated_at"=$4 WHERE "id" = $5`)).
		WithArgs(task.Title, task.Completed, orm.AnyMockArg(), orm.AnyMockArg(), task.ID)
	u.Update(&task)

	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	db, mock := orm.NewDBMock()
	u := TasksUsecase{DB: db}

	task := entity.Task{ID: 1}
	mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM "tasks" WHERE "tasks"."id" = $1`)).
		WithArgs(task.ID)
	u.Delete(&task)

	assert.Nil(t, mock.ExpectationsWereMet())
}
