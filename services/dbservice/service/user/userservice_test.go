package userservice

import (
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	logintokenserviceprovider "server/services/dbservice/serviceprovider/logintoken"
	userserviceprovider "server/services/dbservice/serviceprovider/user"
	mock_sqloperator "server/tests/mocks"
)

func TestGetUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userServiceFactory := &UserServiceFactory{}
	const name string = "user"
	const errName string = "errName"

	userService, err := userServiceFactory.GetUserService(mockDB, name)
	if err == nil {
		t.Logf("userService passed: %v, %v", userService, err)
	} else {
		t.Errorf("userService failed: %v, %v", userService, err)
	}

	userServiceNameErr, err := userServiceFactory.GetUserService(mockDB, errName)
	if err != nil {
		t.Logf("userServiceNameErr passed: %v, %v", userServiceNameErr, err)
	} else {
		t.Errorf("userServiceNameErr failed: %v, %v", userServiceNameErr, err)
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	const sql string = `UPDATE mysql_db.user SET first_name=?, last_name=?, email=?, phone=?, gender=?, birthday=? WHERE user_id=?;`
	var user dto.User
	const userId uint64 = 1
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(
		sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		userId,
	).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)

	userService := &UserService{
		mockDB,
		&userserviceprovider.UserServiceFuncFactory{},
		&logintokenserviceprovider.LoginTokenServiceFuncFactory{},
	}
	updateUser := userService.UpdateUser(userId, user)
	if updateUser == nil {
		t.Logf("updateUser passed: %v", updateUser)
	} else {
		t.Errorf("updateUser failed: %v", updateUser)
	}
}

func TestEnableUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	const sql string = `UPDATE mysql_db.user SET enabled=? WHERE user_id=?;`
	const userId uint64 = 1
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(sql, 1, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)

	userService := &UserService{
		mockDB,
		&userserviceprovider.UserServiceFuncFactory{},
		&logintokenserviceprovider.LoginTokenServiceFuncFactory{},
	}
	enableUser := userService.EnableUser(userId)
	if enableUser == nil {
		t.Logf("enableUser passed: %v", enableUser)
	} else {
		t.Errorf("enableUser failed: %v", enableUser)
	}
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	sql := `SELECT user_id, user_account, first_name, last_name, email, phone, user_profile, gender, birthday, height, allergies, med_compliance FROM mysql_db.user WHERE user_id=?;`
	var user dto.User
	const userId uint64 = 1
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	mockDB.EXPECT().QueryRow(sql, userId).Return(mockRow)
	mockRow.EXPECT().Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.UserProfile,
		&user.Gender,
		&user.Birthday,
		&user.Height,
		&user.Allergies,
		&user.MedCompliance,
	).Return(nil)
	userService := &UserService{
		mockDB,
		&userserviceprovider.UserServiceFuncFactory{},
		&logintokenserviceprovider.LoginTokenServiceFuncFactory{},
	}
	getUser, err := userService.GetUser(userId)
	if getUser != nil {
		t.Logf("getUser passed: %v, %v", getUser, err)
	}
	if err != nil {
		t.Errorf("getUser failed: %v, %v", getUser, err)
	}
}

func TestGetUserByAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	sql := `SELECT 
	user_id,
	user_account,
	first_name,
	last_name,
	user_profile,
	gender
	FROM mysql_db.user WHERE user_account=?;`
	var user dto.User
	const userAccount string = "userAccount"
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	mockDB.EXPECT().QueryRow(sql, userAccount).Return(mockRow)
	mockRow.EXPECT().Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.UserProfile,
		&user.Gender,
	).Return(nil)
	userService := &UserService{
		mockDB,
		&userserviceprovider.UserServiceFuncFactory{},
		&logintokenserviceprovider.LoginTokenServiceFuncFactory{},
	}
	getUserByAccount, err := userService.GetUserByAccount(userAccount)
	if getUserByAccount != nil {
		t.Logf("getUserByAccount passed: %v, %v", getUserByAccount, err)
	}
	if err != nil {
		t.Errorf("getUserByAccount failed: %v, %v", getUserByAccount, err)
	}
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	sql := `SELECT * FROM mysql_db.user WHERE user_id!=? ORDER BY created_time DESC;`

	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	mockDB.EXPECT().Query(sql, 0).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return([]string{
		"user_id",
		"user_account",
		"user_password",
		"phone",
		"first_name",
		"last_name",
		"birthday",
		"created_time",
		"updated_time",
		"email",
		"allergies",
		"enabled",
		"role",
		"gender",
		"height",
		"user_profile",
		"med_compliance",
		"alert",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	userService := &UserService{
		mockDB,
		&userserviceprovider.UserServiceFuncFactory{},
		&logintokenserviceprovider.LoginTokenServiceFuncFactory{},
	}
	getAllUsers, err := userService.GetAllUsers()
	if getAllUsers != nil {
		t.Logf("getAllUsers passed: %v, %v", getAllUsers, err)
	}
	if err != nil {
		t.Errorf("getAllUsers failed: %v, %v", getAllUsers, err)
	}
}

func TestGetSpecificRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	sql := `SELECT * FROM mysql_db.user WHERE role=?;`
	const role uint8 = 1

	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	mockDB.EXPECT().Query(sql, role).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return([]string{
		"user_id",
		"user_account",
		"user_password",
		"phone",
		"first_name",
		"last_name",
		"birthday",
		"created_time",
		"updated_time",
		"email",
		"allergies",
		"enabled",
		"role",
		"gender",
		"height",
		"user_profile",
		"med_compliance",
		"alert",
	}, nil)
	mockRows.EXPECT().Next().Return(false)
	userService := &UserService{
		mockDB,
		&userserviceprovider.UserServiceFuncFactory{},
		&logintokenserviceprovider.LoginTokenServiceFuncFactory{},
	}
	getSpecificRoles, err := userService.GetSpecificRoles(role)
	if getSpecificRoles != nil {
		t.Logf("getSpecificRoles passed: %v, %v", getSpecificRoles, err)
	}
	if err != nil {
		t.Errorf("getSpecificRoles failed: %v, %v", getSpecificRoles, err)
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockSqlTx := mock_sqloperator.NewMockISqlTx(ctrl)
	mockDB.EXPECT().Begin().Return(mockSqlTx, nil)
	mockSqlTx.EXPECT().Rollback().Return(nil)
	sql := `DELETE FROM mysql_db.user WHERE user_id=?;`
	const userId uint64 = 1

	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	mockDB.EXPECT().Exec(sql, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	userService := &UserService{
		mockDB,
		&userserviceprovider.UserServiceFuncFactory{},
		&logintokenserviceprovider.LoginTokenServiceFuncFactory{},
	}
	deleteUser := userService.DeleteUser(userId)
	if deleteUser == nil {
		t.Logf("deleteUser passed: %v", deleteUser)
	} else {
		t.Errorf("deleteUser failed: %v", deleteUser)
	}
}
