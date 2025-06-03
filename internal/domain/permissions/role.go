package permissions

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Permissions uint64 `json:"permissions"`
}

func (role *Role) AddPermission(permission uint64) {
	role.Permissions |= permission
}

func (role *Role) RemovePermission(permission uint64) {
	role.Permissions &^= permission
}

func (role *Role) HasPermission(permission uint64) bool {
	return (role.Permissions & permission) != 0
}
