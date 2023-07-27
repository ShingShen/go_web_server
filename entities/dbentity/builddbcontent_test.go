package dbentity

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestGetBuildDbContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	buildDbContentFactory := &buildDbContentFactory{}
	const dbName string = "db_name"
	const name string = "buildDbContent"
	const errName string = "errName"

	getBuildDbContent, err := buildDbContentFactory.getBuildDbContent(mockDB, dbName, name)
	if err == nil {
		t.Logf("getBuildDbContent passed: %v, %v", getBuildDbContent, err)
	} else {
		t.Errorf("getBuildDbContent failed: %v, %v", getBuildDbContent, err)
	}

	getBuildDbContentNameErr, err := buildDbContentFactory.getBuildDbContent(mockDB, dbName, errName)
	if err != nil {
		t.Logf("getBuildDbContentNameErr passed: %v, %v", getBuildDbContentNameErr, err)
	} else {
		t.Errorf("getBuildDbContentNameErr failed: %v, %v", getBuildDbContentNameErr, err)
	}
}

func TestBuildTables(t *testing.T) {
	const dbName string = "db_name"
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	createTableLoginToken := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.login_token (login_token_id BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL);", dbName)
	createTableMemo := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.memo (memo_id BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL);", dbName)
	createTableSchedule := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.schedule (schedule_id BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL);", dbName)
	createTableUser := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.user (user_id BIGINT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL);", dbName)

	mockDB.EXPECT().Exec(createTableLoginToken).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(createTableMemo).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(createTableSchedule).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(createTableUser).Return(mockSqlResult, nil)

	buildDbContent := &buildDbContent{mockDB, dbName}
	buildDbContent.buildTables()
}

func TestBuildColumns(t *testing.T) {
	const dbName string = "db_name"
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)

	const loginToken string = "login_token"
	loginTokenAddColClickTimes := fmt.Sprintf("ALTER TABLE %s.%s ADD login_token VARCHAR(512);", dbName, loginToken)
	loginTokenAddColUserId := fmt.Sprintf("ALTER TABLE %s.%s ADD user_id BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id);", dbName, loginToken)
	loginTokenAddColCreatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;", dbName, loginToken)
	loginTokenAddColUpdatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;", dbName, loginToken)
	loginTokenAddUniqueUUserId := fmt.Sprintf("ALTER TABLE %s.%s ADD UNIQUE u_user_id (user_id);", dbName, loginToken)
	mockDB.EXPECT().Exec(loginTokenAddColUserId).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(loginTokenAddColClickTimes).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(loginTokenAddColCreatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(loginTokenAddColUpdatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(loginTokenAddUniqueUUserId).Return(mockSqlResult, nil)

	const memo string = "memo"
	memoAddColContent := fmt.Sprintf("ALTER TABLE %s.%s ADD content VARCHAR(128) NOT NULL DEFAULT '';", dbName, memo)
	memoAddColHasRead := fmt.Sprintf("ALTER TABLE %s.%s ADD has_read TINYINT UNSIGNED NOT NULL DEFAULT 0;", dbName, memo)
	memoAddColCreatedBy := fmt.Sprintf("ALTER TABLE %s.%s ADD created_by VARCHAR(64);", dbName, memo)
	memoAddColUpdatedBy := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_by VARCHAR(64);", dbName, memo)
	memoAddColUserId := fmt.Sprintf("ALTER TABLE %s.%s ADD user_id BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id);", dbName, memo)
	memoAddColCreatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;", dbName, memo)
	memoAddColUpdatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;", dbName, memo)
	mockDB.EXPECT().Exec(memoAddColContent).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(memoAddColHasRead).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(memoAddColCreatedBy).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(memoAddColUpdatedBy).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(memoAddColUserId).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(memoAddColCreatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(memoAddColUpdatedTime).Return(mockSqlResult, nil)

	const schedule string = "schedule"
	scheduleAddColTitle := fmt.Sprintf("ALTER TABLE %s.%s ADD title VARCHAR(128) NOT NULL DEFAULT '';", dbName, schedule)
	scheduleAddColNote := fmt.Sprintf("ALTER TABLE %s.%s ADD note VARCHAR(128) NOT NULL DEFAULT '';", dbName, schedule)
	scheduleAddColStartTime := fmt.Sprintf("ALTER TABLE %s.%s ADD start_time DATETIME NOT NULL;", dbName, schedule)
	scheduleAddColEndTime := fmt.Sprintf("ALTER TABLE %s.%s ADD end_time DATETIME NOT NULL;", dbName, schedule)
	scheduleAddColUserId := fmt.Sprintf("ALTER TABLE %s.%s ADD user_id BIGINT(20) DEFAULT 0, ADD FOREIGN KEY(user_id) REFERENCES mysql_db.user(user_id);", dbName, schedule)
	scheduleAddColCreatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;", dbName, schedule)
	scheduleAddColUpdatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;", dbName, schedule)
	mockDB.EXPECT().Exec(scheduleAddColTitle).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(scheduleAddColNote).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(scheduleAddColStartTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(scheduleAddColEndTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(scheduleAddColUserId).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(scheduleAddColCreatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(scheduleAddColUpdatedTime).Return(mockSqlResult, nil)

	const user string = "user"
	userAddColUserAccount := fmt.Sprintf("ALTER TABLE %s.%s ADD user_account VARCHAR(64);", dbName, user)
	userAddColUserPassword := fmt.Sprintf("ALTER TABLE %s.%s ADD user_password VARCHAR(512);", dbName, user)
	userAddColEmail := fmt.Sprintf("ALTER TABLE %s.%s ADD email VARCHAR(64);", dbName, user)
	userAddColPhone := fmt.Sprintf("ALTER TABLE %s.%s ADD phone VARCHAR(64) NOT NULL DEFAULT '';", dbName, user)
	userAddColUserProfile := fmt.Sprintf("ALTER TABLE %s.%s ADD user_profile VARCHAR(512) NOT NULL DEFAULT '';", dbName, user)
	userAddColFirstName := fmt.Sprintf("ALTER TABLE %s.%s ADD first_name VARCHAR(64) NOT NULL DEFAULT '';", dbName, user)
	userAddColLastName := fmt.Sprintf("ALTER TABLE %s.%s ADD last_name VARCHAR(64) NOT NULL DEFAULT '';", dbName, user)
	userAddColGender := fmt.Sprintf("ALTER TABLE %s.%s ADD gender TINYINT UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	userAddColBirthday := fmt.Sprintf("ALTER TABLE %s.%s ADD birthday DATE;", dbName, user)
	userAddColHeight := fmt.Sprintf("ALTER TABLE %s.%s ADD height SMALLINT UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	userAddColMedCompliance := fmt.Sprintf("ALTER TABLE %s.%s ADD med_compliance TINYINT UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	userAddColAlert := fmt.Sprintf("ALTER TABLE %s.%s ADD alert TINYINT(1) UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	userAddColAllergies := fmt.Sprintf("ALTER TABLE %s.%s ADD allergies VARCHAR(64) NOT NULL DEFAULT '';", dbName, user)
	userAddColRole := fmt.Sprintf("ALTER TABLE %s.%s ADD role TINYINT UNSIGNED NOT NULL DEFAULT 1;", dbName, user)
	userAddColEnabled := fmt.Sprintf("ALTER TABLE %s.%s ADD enabled TINYINT UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	userAddColCreatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;", dbName, user)
	userAddColUpdatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;", dbName, user)
	userAddUniqueUUserAccount := fmt.Sprintf("ALTER TABLE %s.%s ADD UNIQUE u_user_account (user_account);", dbName, user)
	userAddUniqueUEmail := fmt.Sprintf("ALTER TABLE %s.%s ADD UNIQUE u_email (email);", dbName, user)
	mockDB.EXPECT().Exec(userAddColUserAccount).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColUserPassword).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColEmail).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColPhone).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColUserProfile).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColFirstName).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColLastName).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColGender).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColBirthday).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColHeight).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColMedCompliance).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColAlert).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColAllergies).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColRole).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColEnabled).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColCreatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddColUpdatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddUniqueUUserAccount).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(userAddUniqueUEmail).Return(mockSqlResult, nil)

	buildDbContent := &buildDbContent{mockDB, dbName}
	buildDbContent.buildColumns()
}
