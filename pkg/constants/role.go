package constants

type UserRole int

const (
	UserRoleUnassigned UserRole = iota
	UserRoleRoot
	end
)

func (ur UserRole) Valid() bool {
	return ur < end
}
