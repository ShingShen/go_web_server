package userentity

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	mock_sqloperator "server/tests/mocks"
)

func TestCreateUserColumns(t *testing.T) {
	const dbName string = "db_name"
	const user string = "user"
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlResult := mock_sqloperator.NewMockISqlResult(ctrl)
	addColUserAccount := fmt.Sprintf("ALTER TABLE %s.%s ADD user_account VARCHAR(64);", dbName, user)
	addColUserPassword := fmt.Sprintf("ALTER TABLE %s.%s ADD user_password VARCHAR(512);", dbName, user)
	addColEmail := fmt.Sprintf("ALTER TABLE %s.%s ADD email VARCHAR(64);", dbName, user)
	addColPhone := fmt.Sprintf("ALTER TABLE %s.%s ADD phone VARCHAR(64) NOT NULL DEFAULT '';", dbName, user)
	addColUserProfile := fmt.Sprintf("ALTER TABLE %s.%s ADD user_profile VARCHAR(512) NOT NULL DEFAULT '';", dbName, user)
	addColFirstName := fmt.Sprintf("ALTER TABLE %s.%s ADD first_name VARCHAR(64) NOT NULL DEFAULT '';", dbName, user)
	addColLastName := fmt.Sprintf("ALTER TABLE %s.%s ADD last_name VARCHAR(64) NOT NULL DEFAULT '';", dbName, user)
	addColGender := fmt.Sprintf("ALTER TABLE %s.%s ADD gender TINYINT UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	addColBirthday := fmt.Sprintf("ALTER TABLE %s.%s ADD birthday DATE;", dbName, user)
	addColHeight := fmt.Sprintf("ALTER TABLE %s.%s ADD height SMALLINT UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	addColMedCompliance := fmt.Sprintf("ALTER TABLE %s.%s ADD med_compliance TINYINT UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	addColAlert := fmt.Sprintf("ALTER TABLE %s.%s ADD alert TINYINT(1) UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	addColAllergies := fmt.Sprintf("ALTER TABLE %s.%s ADD allergies VARCHAR(64) NOT NULL DEFAULT '';", dbName, user)
	addColRole := fmt.Sprintf("ALTER TABLE %s.%s ADD role TINYINT UNSIGNED NOT NULL DEFAULT 1;", dbName, user)
	addColEnabled := fmt.Sprintf("ALTER TABLE %s.%s ADD enabled TINYINT UNSIGNED NOT NULL DEFAULT 0;", dbName, user)
	addColCreatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;", dbName, user)
	addColUpdatedTime := fmt.Sprintf("ALTER TABLE %s.%s ADD updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;", dbName, user)
	addUniqueUUserAccount := fmt.Sprintf("ALTER TABLE %s.%s ADD UNIQUE u_user_account (user_account);", dbName, user)
	addUniqueUEmail := fmt.Sprintf("ALTER TABLE %s.%s ADD UNIQUE u_email (email);", dbName, user)

	mockDB.EXPECT().Exec(addColUserAccount).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUserPassword).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColEmail).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColPhone).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUserProfile).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColFirstName).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColLastName).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColGender).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColBirthday).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColHeight).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColMedCompliance).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColAlert).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColAllergies).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColRole).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColEnabled).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColCreatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addColUpdatedTime).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addUniqueUUserAccount).Return(mockSqlResult, nil)
	mockDB.EXPECT().Exec(addUniqueUEmail).Return(mockSqlResult, nil)

	CreateUserColumns(mockDB, dbName, user)
}
