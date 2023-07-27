package logintokenentity

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestCreateLoginTokenColumn(t *testing.T) {
	const dbName string = "db_name"
	const loginToken string = "login_token"
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	addColClickTimes := fmt.Sprintf("ALTER TABLE %s.%s ADD login_token VARCHAR(512);", dbName, loginToken)
	addColUserId := fmt.Sprintf("ALTER TABLE %s.%s ADD user_id BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id);", dbName, loginToken)
	addColCreatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;", dbName, loginToken)
	addColUpdatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;", dbName, loginToken)
	addUniqueUUserId := fmt.Sprintf("ALTER TABLE %s.%s ADD UNIQUE u_user_id (user_id);", dbName, loginToken)

	mockDB.EXPECT().Exec(addColUserId).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColClickTimes).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColCreatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUpdatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addUniqueUUserId).Return(mockSqlResult, nil)

	CreateLoginTokenColumn(mockDB, dbName, loginToken)
}
