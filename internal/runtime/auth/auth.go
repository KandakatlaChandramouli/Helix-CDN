package auth

func Validate(
	token string,
) bool {
	return token != ""
}
