package port

type ResetPasswordRepository interface {
	ResetPassword(userid int, Password string) error
}