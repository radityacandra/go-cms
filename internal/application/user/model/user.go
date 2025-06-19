package model

type User struct {
	Id        string  `db:"id"`
	Username  string  `db:"username"`
	Password  string  `db:"password"`
	CreatedAt int64   `db:"created_at"`
	CreatedBy string  `db:"created_by"`
	UpdatedAt *int64  `db:"updated_at"`
	UpdatedBy *string `db:"updated_by"`

	UserRoles []UserRole `db:"-"`
}

type UserRole struct {
	Id        string  `db:"id"`
	RoleId    string  `db:"role_id"`
	UserId    string  `db:"user_id"`
	CreatedAt int64   `db:"created_at"`
	CreatedBy string  `db:"created_by"`
	UpdatedAt *int64  `db:"updated_at"`
	UpdatedBy *string `db:"updated_by"`

	RoleAcls []RoleAcl `db:"-"`
	Role     Role      `db:"-"`
}

type Role struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type RoleAcl struct {
	Id     string `db:"id"`
	RoleId string `db:"role_id"`
	Access string `db:"access"`
}

func NewUser(id, username, password string, createdAt int64, createdBy string) *User {
	return &User{
		Id:        id,
		Username:  username,
		Password:  password,
		CreatedAt: createdAt,
		CreatedBy: createdBy,
	}
}

func (u *User) CollectAllAccess() []string {
	acl := []string{}
	for _, userRole := range u.UserRoles {
		for _, roleAcl := range userRole.RoleAcls {
			acl = append(acl, roleAcl.Access)
		}
	}

	return acl
}
