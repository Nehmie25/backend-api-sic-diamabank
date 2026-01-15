package port

type TokenGenerator interface {
	Generate(userID string, role string) (string, error)
	Validate(token string) (map[string]interface{}, error)
}
