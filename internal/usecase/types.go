package usecase

type DAO interface {
	FindOneOrNone(model interface{}, filter map[string]interface{}) (interface{}, error)
	FindAll(model interface{}, filter map[string]interface{}) (interface{}, error)
	Add(model interface{}) error
	Delete(model interface{}, filter map[string]interface{}) error
	Update(model interface{}, id uint, values map[string]interface{}) error
}
