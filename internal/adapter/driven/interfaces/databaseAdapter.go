package interfaces

type IDatabaseAdapter interface {
	FindOne(value string) (interface{}, error)
	Save(identifier string, data interface{}) error
}
