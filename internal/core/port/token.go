package port

type TokenGenerator interface {
	Generate(userID string, role string, email string, nom string) (string, error)
	Validate(token string) (map[string]interface{}, error)
}
