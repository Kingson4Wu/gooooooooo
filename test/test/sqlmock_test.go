package test

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

/**
————————————————
版权声明：本文为CSDN博主「翔云123456」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/lanyang123456/article/details/123303324
*/

type OrangeProcess struct {
	Id        int64
	Hostname  string
	Port      int
	StartTime string
	UID       string
}

func WriteOrangeProcess(db *sql.DB, orangeProcess *OrangeProcess) error {
	sqlResult, err := db.Exec(`
			insert ignore
				into orange_process (
					hostname,
					port,
					start_time,
					uid,
				) values (
					?,
					?,
					NOW(),
					?
				)
			`,
		orangeProcess.Hostname, orangeProcess.Port,
		orangeProcess.UID,
	)
	if err != nil {
		return fmt.Errorf("insert ignore into orange_process failed:%s, orangeProcess:%+v", err, orangeProcess)
	}
	rows, err := sqlResult.RowsAffected()
	if err != nil {
		return fmt.Errorf("get sql RowsAffected failed:%s, orangeProcess:%+v", err, orangeProcess)
	}
	if rows == 0 {
		return fmt.Errorf("create orange_process record failed, orangeProcess:%+v", orangeProcess)
	}

	return nil
}

func TestWriteOrangeProcess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	orangeProcess := &OrangeProcess{Hostname: "number-1", Port: 3306, UID: "AA-BB-CC"}

	sqlInsertIgnoreIntoSql := "insert ignore into orange_process"

	// 模拟 insert ignore into报错
	mock.ExpectExec(sqlInsertIgnoreIntoSql).WillReturnError(errors.New("insert error"))
	err = WriteOrangeProcess(db, orangeProcess)
	if !strings.Contains(err.Error(), "insert ignore into orange_process failed:insert error") {
		t.Fatalf("unexpected error:%s", err)
	}

	// 模拟 insert ignore into 返回Result 中存在错误
	mock.ExpectExec(sqlInsertIgnoreIntoSql).WillReturnResult(sqlmock.NewErrorResult(errors.New("result error")))
	err = WriteOrangeProcess(db, orangeProcess)
	if !strings.Contains(err.Error(), "get sql RowsAffected failed:result error") {
		t.Fatalf("unexpected error:%s", err)
	}

	// 模拟 insert ignore into 影响行数为0
	mock.ExpectExec(sqlInsertIgnoreIntoSql).WillReturnResult(sqlmock.NewResult(0, 0))
	err = WriteOrangeProcess(db, orangeProcess)
	if !strings.Contains(err.Error(), "create orange_process record failed") {
		t.Fatalf("unexpected error:%s", err)
	}

	// 模拟 insert ignore into 正常
	mock.ExpectExec(sqlInsertIgnoreIntoSql).WillReturnResult(sqlmock.NewResult(0, 1))
	err = WriteOrangeProcess(db, orangeProcess)
	if err != nil {
		t.Fatalf("unexpected error:%s", err)
	}
}
