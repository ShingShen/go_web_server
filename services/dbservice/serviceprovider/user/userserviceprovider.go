package userserviceprovider

import (
	"encoding/json"
	"fmt"
	"server/dto"
	"server/utils/operator"

	sqlOperator "server/utils/sqloperator"
)

type UserServiceFuncFactory struct{}

func (u *UserServiceFuncFactory) GetUserServiceFunc(db sqlOperator.ISqlDB, name string) (IUserServiceFunc, error) {
	if name == "userServiceFunc" {
		return &UserServiceFunc{db}, nil
	}
	return nil, fmt.Errorf("wrong user service func type passed")
}

type UserServiceFunc struct {
	db sqlOperator.ISqlDB
}

func (u *UserServiceFunc) CreateUser(encodeUserPassword string, user dto.User) error {
	createUserSql := `INSERT INTO mysql_db.user(user_account, user_password, first_name, last_name, email, phone, gender, birthday, role) values(?,?,?,?,?,?,?,?,?);`
	createUserRes, err := u.db.Exec(
		createUserSql,
		user.UserAccount,
		encodeUserPassword,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		user.Role,
	)
	if err != nil {
		fmt.Printf("Failed to insert user data, err: %v\n", err)
		return err
	}
	lastInsertId, err := createUserRes.LastInsertId()
	if err != nil {
		fmt.Printf("Failed to get inserted user id, err: %v\n", err)
		return err
	}
	fmt.Println("Inserted user id: ", lastInsertId)
	createUserRowsAffected, err := createUserRes.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get createUser affected rows, err: %v\n", err)
		return err
	}

	fmt.Println("CreateUser affected rows: ", createUserRowsAffected)
	return nil
}

func (u *UserServiceFunc) UpdateUser(userId uint64, user dto.User) error {
	sql := `UPDATE mysql_db.user SET first_name=?, last_name=?, email=?, phone=?, gender=?, birthday=? WHERE user_id=?;`
	res, err := u.db.Exec(
		sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Gender,
		user.Birthday,
		userId,
	)
	if err != nil {
		fmt.Printf("Update user data failed, err: %v\n", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v\n", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (u *UserServiceFunc) UpdateUserProfileUrl(userId uint64, userProfileUrl string) error {
	sql := `UPDATE mysql_db.user SET user_profile=? WHERE user_id=?;`
	res, err := u.db.Exec(sql, userProfileUrl, userId)
	if err != nil {
		fmt.Printf("Update user profile failed, err: %v\n", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v\n", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (u *UserServiceFunc) ResetUserPassword(userId uint64, newUserPassword string) error {
	sql := `UPDATE mysql_db.user SET user_password=? WHERE user_id=?;`
	res, err := u.db.Exec(sql, newUserPassword, userId)
	if err != nil {
		fmt.Printf("Reset user password failed, err: %v\n", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v\n", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (u *UserServiceFunc) EnableUser(userId uint64) error {
	sql := `UPDATE mysql_db.user SET enabled=? WHERE user_id=?;`
	res, err := u.db.Exec(sql, 1, userId)
	if err != nil {
		fmt.Printf("Update user data failed, err: %v\n", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v\n", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (u *UserServiceFunc) GetUser(userId uint64) ([]byte, error) {
	var user dto.User
	sql := `SELECT user_id, user_account, first_name, last_name, email, phone, user_profile, gender, birthday, height, allergies, med_compliance FROM mysql_db.user WHERE user_id=?;`
	res := u.db.QueryRow(sql, userId)
	err := res.Scan(
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
	)
	if err != nil {
		return nil, err
	} else {
		userJsonData, _ := json.Marshal(map[string]interface{}{
			"user_id":        user.UserId,
			"user_account":   user.UserAccount,
			"first_name":     user.FirstName,
			"last_name":      user.LastName,
			"email":          user.Email,
			"phone":          user.Phone,
			"user_profile":   user.UserProfile,
			"gender":         user.Gender,
			"birthday":       user.Birthday,
			"height":         user.Height,
			"allergies":      user.Allergies,
			"med_compliance": user.MedCompliance,
		})
		return userJsonData, nil
	}
}

func (u *UserServiceFunc) GetUserByAccount(userAccount string) ([]byte, error) {
	var user dto.User
	sql := `SELECT 
	user_id,
	user_account,
	first_name,
	last_name,
	user_profile,
	gender
	FROM mysql_db.user WHERE user_account=?;`
	res := u.db.QueryRow(sql, userAccount)
	err := res.Scan(
		&user.UserId,
		&user.UserAccount,
		&user.FirstName,
		&user.LastName,
		&user.UserProfile,
		&user.Gender,
	)
	if err != nil {
		return nil, err
	} else {
		userJsonData, _ := json.Marshal(map[string]interface{}{
			"user_id":      user.UserId,
			"user_account": user.UserAccount,
			"first_name":   user.FirstName,
			"last_name":    user.LastName,
			"user_profile": user.UserProfile,
			"gender":       user.Gender,
		})
		return userJsonData, nil
	}
}

func (u *UserServiceFunc) GetAllUsers() ([]byte, error) {
	sql := `SELECT * FROM mysql_db.user WHERE user_id!=? ORDER BY created_time DESC;`
	res, err := u.db.Query(sql, 0)
	if err != nil {
		fmt.Printf("Failed to query users, err: %v\n", err)
		return nil, err
	}

	allUsersList, err := operator.CreatingDataList(res)
	if err != nil {
		return nil, err
	} else {
		return allUsersList, nil
	}
}

func (u *UserServiceFunc) GetSpecificRoles(role uint8) ([]byte, error) {
	sql := `SELECT * FROM mysql_db.user WHERE role=?;`
	res, err := u.db.Query(sql, role)
	if err != nil {
		fmt.Printf("Failed to query users, err: %v\n", err)
		return nil, err
	}

	allUsersList, err := operator.CreatingDataList(res)
	if err != nil {
		return nil, err
	} else {
		return allUsersList, nil
	}
}

func (u *UserServiceFunc) DeleteUser(userId uint64) error {
	// SET FOREIGN_KEY_CHECKS=0;
	// SET FOREIGN_KEY_CHECKS=1;
	sql := `DELETE FROM mysql_db.user WHERE user_id=?;`
	res, err := u.db.Exec(sql, userId)
	if err != nil {
		fmt.Printf("Failed to delete user data, err: %v\n", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Failed to get affected rows, err: %v\n", err)
		return err
	}
	fmt.Println("Affected rows: ", rowsAffected)
	return nil
}

func (u *UserServiceFunc) GetUserIdByAccountAndEncodePassword(authAccount string, encodeAuthPassword string) (uint64, error) {
	var user dto.User
	sql := `SELECT user_id FROM mysql_db.user WHERE user_account=? AND user_password=?;`
	getUserIdRes := u.db.QueryRow(sql, authAccount, encodeAuthPassword)
	getUserIdResScan := getUserIdRes.Scan(&user.UserId)
	if getUserIdResScan != nil {
		fmt.Println("getUserId, Error: ", getUserIdResScan)
		return 0, getUserIdResScan
	}
	return user.UserId, nil
}
