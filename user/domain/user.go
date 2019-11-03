package domain

// IsAdmin ...
func IsAdmin(username string) bool {
	return username == "admin"
}
