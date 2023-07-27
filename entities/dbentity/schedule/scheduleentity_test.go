package scheduleentity

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestCreateScheduleColumns(t *testing.T) {
	const dbName string = "db_name"
	const schedule string = "schedule"
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	addColTitle := fmt.Sprintf("ALTER TABLE %s.%s ADD title VARCHAR(128) NOT NULL DEFAULT '';", dbName, schedule)
	addColNote := fmt.Sprintf("ALTER TABLE %s.%s ADD note VARCHAR(128) NOT NULL DEFAULT '';", dbName, schedule)
	addColStartTime := fmt.Sprintf("ALTER TABLE %s.%s ADD start_time DATETIME NOT NULL;", dbName, schedule)
	addColEndTime := fmt.Sprintf("ALTER TABLE %s.%s ADD end_time DATETIME NOT NULL;", dbName, schedule)
	addColUserId := fmt.Sprintf("ALTER TABLE %s.%s ADD user_id BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id);", dbName, schedule)
	addColCreatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;", dbName, schedule)
	addColUpdatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;", dbName, schedule)

	mockDB.EXPECT().Exec(addColTitle).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColNote).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColStartTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColEndTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUserId).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColCreatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUpdatedTime).Return(mockSqlResult, nil)

	CreateScheduleColumns(mockDB, dbName, schedule)
}
