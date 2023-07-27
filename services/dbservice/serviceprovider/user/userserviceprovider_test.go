package userserviceprovider

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"server/dto"
	mock_sqloperator "server/tests/mocks"
)

func TestGetUserServiceFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	userServiceFuncFactory := &UserServiceFuncFactory{}
	const name string = "userServiceFunc"
	const errName string = "errName"

	userServiceFunc, err := userServiceFuncFactory.GetUserServiceFunc(mockDB, name)
	if err == nil {
		t.Logf("userServiceFunc passed: %v, %v", userServiceFunc, err)
	} else {
		t.Errorf("userServiceFunc failed: %v, %v", userServiceFunc, err)
	}

	userServiceFuncNameErr, err := userServiceFuncFactory.GetUserServiceFunc(mockDB, errName)
	if err != nil {
		t.Logf("userServiceFuncNameErr passed: %v, %v", userServiceFuncNameErr, err)
	} else {
		t.Errorf("userServiceFuncNameErr failed: %v, %v", userServiceFuncNameErr, err)
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `INSERT INTO mysql_db.user(user_account, user_password, first_name, last_name, email, phone, gender, birthday, role) values(?,?,?,?,?,?,?,?,?);`
	var user dto.User
	userServiceFunc := &UserServiceFunc{mockDB}
	encodeUserPassword := "encodeUserPassword"

	// createUser Res Err
	mockDB.EXPECT().Exec(
		sql,
		user.UserAccount,
		encodeUserPassword,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		user.Role,
	).Return(nil, fmt.Errorf("no result"))
	createUserResErr := userServiceFunc.CreateUser(encodeUserPassword, user)
	if createUserResErr != nil {
		t.Logf("createUserResErr passed: %v", createUserResErr)
	} else {
		t.Errorf("createUserResErr failed: %v", createUserResErr)
	}

	// createUser Insert ID Err
	mockDB.EXPECT().Exec(
		sql,
		user.UserAccount,
		encodeUserPassword,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		user.Role,
	).Return(mockResult, nil)
	mockResult.EXPECT().LastInsertId().Return(int64(0), fmt.Errorf("Failed to insert ID"))
	createUserInsertIDErr := userServiceFunc.CreateUser(encodeUserPassword, user)
	if createUserInsertIDErr != nil {
		t.Logf("createUserInsertIDErr passed: %v", createUserInsertIDErr)
	} else {
		t.Errorf("createUserInsertIDErr failed: %v", createUserInsertIDErr)
	}

	// createUser Rows Affected Err
	mockDB.EXPECT().Exec(
		sql,
		user.UserAccount,
		encodeUserPassword,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		user.Role,
	).Return(mockResult, nil)
	mockResult.EXPECT().LastInsertId().Return(int64(10), nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	createUserRowsAffectedErr := userServiceFunc.CreateUser(encodeUserPassword, user)
	if createUserRowsAffectedErr != nil {
		t.Logf("createUserRowsAffectedErr passed: %v", createUserRowsAffectedErr)
	} else {
		t.Errorf("createUserRowsAffectedErr failed: %v", createUserRowsAffectedErr)
	}

	// createUser
	mockDB.EXPECT().Exec(
		sql,
		user.UserAccount,
		encodeUserPassword,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		user.Role,
	).Return(mockResult, nil)
	mockResult.EXPECT().LastInsertId().Return(int64(10), nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	createUser := userServiceFunc.CreateUser(encodeUserPassword, user)
	if createUser == nil {
		t.Logf("createUser passed: %v", createUser)
	} else {
		t.Errorf("createUser failed: %v", createUser)
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `UPDATE mysql_db.user SET first_name=?, last_name=?, email=?, phone=?, gender=?, birthday=? WHERE user_id=?;`
	var user dto.User
	userServiceFunc := &UserServiceFunc{mockDB}
	const userId uint64 = 1

	// updateUser Res Err
	mockDB.EXPECT().Exec(
		sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		userId,
	).Return(nil, fmt.Errorf("no result"))
	updateUserResErr := userServiceFunc.UpdateUser(userId, user)
	if updateUserResErr != nil {
		t.Logf("updateUserResErr passed: %v", updateUserResErr)
	} else {
		t.Errorf("updateUserResErr failed: %v", updateUserResErr)
	}

	// updateUser Rows Affected Err
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
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	updateUserRowsAffectedErr := userServiceFunc.UpdateUser(userId, user)
	if updateUserRowsAffectedErr != nil {
		t.Logf("updateUserRowsAffectedErr passed: %v", updateUserRowsAffectedErr)
	} else {
		t.Errorf("updateUserRowsAffectedErr failed: %v", updateUserRowsAffectedErr)
	}

	// updateUser
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
	updateUser := userServiceFunc.UpdateUser(userId, user)
	if updateUser == nil {
		t.Logf("updateUser passed: %v", updateUser)
	} else {
		t.Errorf("updateUser failed: %v", updateUser)
	}
}

func TestUpdateUserProfileUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `UPDATE mysql_db.user SET user_profile=? WHERE user_id=?;`
	const userId uint64 = 1
	const userProfileUrl string = "userProfileUrl"
	userServiceFunc := &UserServiceFunc{mockDB}

	// updateUserProfileUrl Res Err
	mockDB.EXPECT().Exec(sql, userProfileUrl, userId).Return(nil, fmt.Errorf("no result"))
	updateUserProfileUrlResErr := userServiceFunc.UpdateUserProfileUrl(userId, userProfileUrl)
	if updateUserProfileUrlResErr != nil {
		t.Logf("updateUserProfileUrlResErr passed: %v", updateUserProfileUrlResErr)
	} else {
		t.Errorf("updateUserProfileUrlResErr failed: %v", updateUserProfileUrlResErr)
	}

	// updateUserProfileUrl Rows Affected Err
	mockDB.EXPECT().Exec(sql, userProfileUrl, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	updateUserProfileUrlRowsAffectedErr := userServiceFunc.UpdateUserProfileUrl(userId, userProfileUrl)
	if updateUserProfileUrlRowsAffectedErr != nil {
		t.Logf("updateUserProfileUrlRowsAffectedErr passed: %v", updateUserProfileUrlRowsAffectedErr)
	} else {
		t.Errorf("updateUserProfileUrlRowsAffectedErr failed: %v", updateUserProfileUrlRowsAffectedErr)
	}

	// updateUserProfileUrl
	mockDB.EXPECT().Exec(sql, userProfileUrl, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	updateUserProfileUrl := userServiceFunc.UpdateUserProfileUrl(userId, userProfileUrl)
	if updateUserProfileUrl == nil {
		t.Logf("updateUserProfileUrl passed: %v", updateUserProfileUrl)
	} else {
		t.Errorf("updateUserProfileUrl failed: %v", updateUserProfileUrl)
	}
}

func TestResetUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `UPDATE mysql_db.user SET user_password=? WHERE user_id=?;`
	userServiceFunc := &UserServiceFunc{mockDB}
	const userId uint64 = 1
	const newUserPassword string = "newUserPassword"

	// resetUserPassword Res Err
	mockDB.EXPECT().Exec(sql, newUserPassword, userId).Return(nil, fmt.Errorf("no result"))
	resetUserPasswordResErr := userServiceFunc.ResetUserPassword(userId, newUserPassword)
	if resetUserPasswordResErr != nil {
		t.Logf("resetUserPasswordResErr passed: %v", resetUserPasswordResErr)
	} else {
		t.Errorf("resetUserPasswordResErr failed: %v", resetUserPasswordResErr)
	}

	// resetUserPassword Rows Affected Err
	mockDB.EXPECT().Exec(sql, newUserPassword, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	resetUserPasswordRowsAffectedErr := userServiceFunc.ResetUserPassword(userId, newUserPassword)
	if resetUserPasswordRowsAffectedErr != nil {
		t.Logf("resetUserPasswordRowsAffectedErr passed: %v", resetUserPasswordRowsAffectedErr)
	} else {
		t.Errorf("resetUserPasswordRowsAffectedErr failed: %v", resetUserPasswordRowsAffectedErr)
	}

	// resetUserPassword
	mockDB.EXPECT().Exec(sql, newUserPassword, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	resetUserPassword := userServiceFunc.ResetUserPassword(userId, newUserPassword)
	if resetUserPassword == nil {
		t.Logf("resetUserPassword passed: %v", resetUserPassword)
	} else {
		t.Errorf("resetUserPassword failed: %v", resetUserPassword)
	}
}

func TestEnableUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	const sql string = `UPDATE mysql_db.user SET enabled=? WHERE user_id=?;`
	userServiceFunc := &UserServiceFunc{mockDB}
	const userId uint64 = 1

	// enableUser Res Err
	mockDB.EXPECT().Exec(sql, 1, userId).Return(nil, fmt.Errorf("no result"))
	enableUserResErr := userServiceFunc.EnableUser(userId)
	if enableUserResErr != nil {
		t.Logf("enableUserResErr passed: %v", enableUserResErr)
	} else {
		t.Errorf("enableUserResErr failed: %v", enableUserResErr)
	}

	// enableUser Rows Affected Err
	mockDB.EXPECT().Exec(sql, 1, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("no rows affected"))
	enableUserRowsAffectedErr := userServiceFunc.EnableUser(userId)
	if enableUserRowsAffectedErr != nil {
		t.Logf("enableUserRowsAffectedErr passed: %v", enableUserRowsAffectedErr)
	} else {
		t.Errorf("enableUserRowsAffectedErr failed: %v", enableUserRowsAffectedErr)
	}

	// enableUser
	mockDB.EXPECT().Exec(sql, 1, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	enableUser := userServiceFunc.EnableUser(userId)
	if enableUser == nil {
		t.Logf("enableUser passed: %v", enableUser)
	} else {
		t.Errorf("enableUser failed: %v", enableUser)
	}
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	sql := `SELECT user_id, user_account, first_name, last_name, email, phone, user_profile, gender, birthday, height, allergies, med_compliance FROM mysql_db.user WHERE user_id=?;`
	var user dto.User
	const userId uint64 = 1
	userServiceFunc := &UserServiceFunc{mockDB}

	// getUser Scan Err
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
	).Return(fmt.Errorf("Scan error"))
	getUserScanErr, err := userServiceFunc.GetUser(userId)
	if err != nil {
		t.Logf("getUserScanErr passed: %v, %v", getUserScanErr, err)
	} else {
		t.Errorf("getUserScanErr failed: %v, %v", getUserScanErr, err)
	}

	// getUser
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
	getUser, err := userServiceFunc.GetUser(userId)
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
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
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
	userServiceFunc := &UserServiceFunc{mockDB}

	// getUserByAccount Scan Err
	mockDB.EXPECT().QueryRow(sql, userAccount).Return(mockRow)
	mockRow.EXPECT().Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.UserProfile,
		&user.Gender,
	).Return(fmt.Errorf("Scan error"))
	getUserByAccountScanErr, err := userServiceFunc.GetUserByAccount(userAccount)
	if err != nil {
		t.Logf("getUserByAccountScanErr passed: %v, %v", getUserByAccountScanErr, err)
	} else {
		t.Errorf("getUserByAccountScanErr failed: %v, %v", getUserByAccountScanErr, err)
	}

	// getUserByAccount
	mockDB.EXPECT().QueryRow(sql, userAccount).Return(mockRow)
	mockRow.EXPECT().Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.UserProfile,
		&user.Gender,
	).Return(nil)
	getUserByAccount, err := userServiceFunc.GetUserByAccount(userAccount)
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
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	sql := `SELECT * FROM mysql_db.user WHERE user_id!=? ORDER BY created_time DESC;`
	userServiceFunc := &UserServiceFunc{mockDB}

	// GetAllUsers Query Err
	mockDB.EXPECT().Query(sql, 0).Return(nil, fmt.Errorf("Query Error"))
	getAllUsersQueryErr, err := userServiceFunc.GetAllUsers()
	if err != nil {
		t.Logf("getAllUsersQueryErr passed: %v, %v", getAllUsersQueryErr, err)
	} else {
		t.Errorf("getAllUsersQueryErr failed: %v, %v", getAllUsersQueryErr, err)
	}

	// GetAllUsers Col Err
	mockDB.EXPECT().Query(sql, 0).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return(nil, fmt.Errorf("Columns error"))
	getAllUsersColErr, err := userServiceFunc.GetAllUsers()
	if err != nil {
		t.Logf("getAllUsersColErr passed: %v, %v", getAllUsersColErr, err)
	} else {
		t.Errorf("getAllUsersColErr failed: %v, %v", getAllUsersColErr, err)
	}

	// GetAllUsers
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
	getAllUsers, err := userServiceFunc.GetAllUsers()
	if err == nil {
		t.Logf("getAllUsers passed: %v, %v", getAllUsers, err)
	} else {
		t.Errorf("getAllUsers failed: %v, %v", getAllUsers, err)
	}
}

func TestGetSpecificRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRows := mock_sqloperator.NewMockISqlRows(ctrl)
	sql := `SELECT * FROM mysql_db.user WHERE role=?;`
	const role uint8 = 1
	userServiceFunc := &UserServiceFunc{mockDB}

	// getSpecificRoles Query Err
	mockDB.EXPECT().Query(sql, role).Return(nil, fmt.Errorf("Query Error"))
	getSpecificRolesQueryErr, err := userServiceFunc.GetSpecificRoles(role)
	if err != nil {
		t.Logf("getSpecificRolesQueryErr passed: %v, %v", getSpecificRolesQueryErr, err)
	} else {
		t.Errorf("getSpecificRolesQueryErr failed: %v, %v", getSpecificRolesQueryErr, err)
	}

	// getSpecificRoles Col Err
	mockDB.EXPECT().Query(sql, role).Return(mockRows, nil)
	mockRows.EXPECT().Columns().Return(nil, fmt.Errorf("Columns error"))
	getSpecificRolesColErr, err := userServiceFunc.GetSpecificRoles(role)
	if err != nil {
		t.Logf("getSpecificRolesColErr passed: %v, %v", getSpecificRolesColErr, err)
	} else {
		t.Errorf("getSpecificRolesColErr failed: %v, %v", getSpecificRolesColErr, err)
	}

	// getSpecificRoles
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
	getSpecificRoles, err := userServiceFunc.GetSpecificRoles(role)
	if err == nil {
		t.Logf("getSpecificRoles passed: %v, %v", getSpecificRoles, err)
	} else {
		t.Errorf("getSpecificRoles failed: %v, %v", getSpecificRoles, err)
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockResult := mock_sqloperator.NewMockISqlResult(ctrl)
	sql := `DELETE FROM mysql_db.user WHERE user_id=?;`
	const userId uint64 = 1
	userServiceFunc := &UserServiceFunc{mockDB}

	// deleteUser res err
	mockDB.EXPECT().Exec(sql, userId).Return(nil, fmt.Errorf("res err"))
	deleteUserResErr := userServiceFunc.DeleteUser(userId)
	if deleteUserResErr != nil {
		t.Logf("deleteUserResErr passed: %v", deleteUserResErr)
	} else {
		t.Errorf("deleteUserResErr failed: %v", deleteUserResErr)
	}

	// deleteUser rows affected err
	mockDB.EXPECT().Exec(sql, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(0), fmt.Errorf("rows affected error"))
	deleteUserRowsAffectedErr := userServiceFunc.DeleteUser(userId)
	if deleteUserRowsAffectedErr != nil {
		t.Logf("deleteUserRowsAffectedErr passed: %v", deleteUserRowsAffectedErr)
	} else {
		t.Errorf("deleteUserRowsAffectedErr failed: %v", deleteUserRowsAffectedErr)
	}

	// deleteUser
	mockDB.EXPECT().Exec(sql, userId).Return(mockResult, nil)
	mockResult.EXPECT().RowsAffected().Return(int64(1), nil)
	deleteUser := userServiceFunc.DeleteUser(userId)
	if deleteUser == nil {
		t.Logf("deleteUser passed: %v", deleteUser)
	} else {
		t.Errorf("deleteUser failed: %v", deleteUser)
	}
}

func TestGetUserIdByAccountAndEncodePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mock_sqloperator.NewMockISqlDB(ctrl)
	mockRow := mock_sqloperator.NewMockISqlRow(ctrl)
	sql := `SELECT user_id FROM mysql_db.user WHERE user_account=? AND user_password=?;`
	var user dto.User
	const authAccount string = "authAccount"
	const encodeAuthPassword string = "encodeAuthPassword"
	userServiceFunc := &UserServiceFunc{mockDB}

	// getUserIdByAccountAndEncodePassword Scan Err
	mockDB.EXPECT().QueryRow(sql, authAccount, encodeAuthPassword).Return(mockRow)
	mockRow.EXPECT().Scan(&user.UserId).Return(fmt.Errorf("Scan error"))
	getUserIdByAccountAndEncodePasswordScanErr, err := userServiceFunc.GetUserIdByAccountAndEncodePassword(authAccount, encodeAuthPassword)
	if err != nil {
		t.Logf("getUserIdByAccountAndEncodePasswordScanErr passed: %v, %v", getUserIdByAccountAndEncodePasswordScanErr, err)
	} else {
		t.Errorf("getUserIdByAccountAndEncodePasswordScanErr failed: %v, %v", getUserIdByAccountAndEncodePasswordScanErr, err)
	}

	// getUserIdByAccountAndEncodePassword
	mockDB.EXPECT().QueryRow(sql, authAccount, encodeAuthPassword).Return(mockRow)
	mockRow.EXPECT().Scan(&user.UserId).Return(nil)
	getUserIdByAccountAndEncodePassword, err := userServiceFunc.GetUserIdByAccountAndEncodePassword(authAccount, encodeAuthPassword)
	if err == nil {
		t.Logf("getUserIdByAccountAndEncodePassword passed: %v, %v", getUserIdByAccountAndEncodePassword, err)
	}
	if err != nil {
		t.Errorf("getUserIdByAccountAndEncodePassword failed: %v, %v", getUserIdByAccountAndEncodePassword, err)
	}
}
