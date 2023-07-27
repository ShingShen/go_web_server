package memoentity

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestCreateMemoColumns(t *testing.T) {
	const dbName string = "db_name"
	const memo string = "memo"
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	addColContent := fmt.Sprintf("ALTER TABLE %s.%s ADD content VARCHAR(128) NOT NULL DEFAULT '';", dbName, memo)
	addColHasRead := fmt.Sprintf("ALTER TABLE %s.%s ADD has_read TINYINT UNSIGNED NOT NULL DEFAULT 0;", dbName, memo)
	addColCreatedBy := fmt.Sprintf("ALTER TABLE %s.%s ADD created_by VARCHAR(64);", dbName, memo)
	addColUpdatedBy := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_by VARCHAR(64);", dbName, memo)
	addColUserId := fmt.Sprintf("ALTER TABLE %s.%s ADD user_id BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id);", dbName, memo)
	addColCreatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;", dbName, memo)
	addColUpdatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;", dbName, memo)

	mockDB.EXPECT().Exec(addColContent).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColHasRead).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColCreatedBy).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUpdatedBy).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUserId).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColCreatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUpdatedTime).Return(mockSqlResult, nil)

	CreateMemoColumns(mockDB, dbName, memo)
}
