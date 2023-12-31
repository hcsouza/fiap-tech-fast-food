package repository

type IDatabaseAdapter interface {
	FindOne(key, value string) (interface{}, error)
	Save(identifier string, data interface{}) (id interface{}, err error)
}
