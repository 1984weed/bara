package domain

// IsAdmin ...
func IsAdmin(username string) bool {
	return username == "admin"
}

// UserForUpdate ...
type UserForUpdate struct {
	UserName *string
	RealName *string
	Email    *string
	Bio      *string
	Image    *string
}
