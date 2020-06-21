package domain

// IsAdmin ...
func IsAdmin(username string) bool {
	return username == "admin"
}

// UserForUpdate ...
type UserForUpdate struct {
	UserName    *string
	DisplayName *string
	Email       *string
	Bio         *string
	Image       *string
	ImageURL    *string
	ImageUUID   *string
}
