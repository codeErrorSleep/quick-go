package repo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMockGORM(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows)

	// gormDB, err := gorm.Open(mysql.New(mysql.Config{
	// 	SkipInitializeWithVersion: true,
	// 	Conn:                      db,
	// }), &gorm.Config{})
	// if nil != err {
	// 	t.Fatalf("Init DB with sqlmock failed, err %v", err)
	// }
	// spuRepo := NewMysqlSpuRepository(gormDB)
	// ctx, _ := context.WithCancel(context.Background())
	// if _, err = spuRepo.GetSpuDetail(ctx, "aa", "dsfds"); err == nil {
	// 	t.Errorf("was expecting an error, but there was none")
	// }

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
