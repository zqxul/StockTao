package dao

// SelectUser ==> Select User by Condition
func SelectUser(u *User) *User {
	db := Connect()
	user := &User{}
	db.First(user, Condition(u))
	return &user
}

// InsertUser ==> Insert User into user table
func InsertUser(u *User) uint64 {
	db := Connect()
	db.Create(u)
	return u.ID
}

// UpdateUser ==> update user
func UpdateUser(u *User) {
	db := Connect()
	db.Update(u)
}

// DeleteUser ==> delete user
func DeleteUser(u *User) {
	db := Connect()
	db.Delete(u)
}
