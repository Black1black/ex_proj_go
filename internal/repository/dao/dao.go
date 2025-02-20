package dao

import (
	"gorm.io/gorm"
)

type BaseDAO struct {
	db *gorm.DB
}

// Конструктор для создания нового экземпляра BaseDAO
func NewBaseDAO(db *gorm.DB) *BaseDAO {
	return &BaseDAO{
		db: db,
	}
}

// FindOneOrNone ищет одну запись или возвращает nil
func (dao *BaseDAO) FindOneOrNone(model interface{}, filter map[string]interface{}) (interface{}, error) {
	result := dao.db.Where(filter).Take(model)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Возвращаем nil, чтобы имитировать работников Python "one_or_none"
		}
		return nil, result.Error
	}
	return model, nil
}

// FindAll возвращает все записи, которые соответствуют фильтру
func (dao *BaseDAO) FindAll(model interface{}, filter map[string]interface{}) (interface{}, error) {
	result := dao.db.Where(filter).Find(model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}

// Add добавляет новую запись в базу данных
func (dao *BaseDAO) Add(model interface{}) error {
	result := dao.db.Create(model)
	return result.Error
}

// Delete удаляет записи, которые соответствуют фильтру
func (dao *BaseDAO) Delete(model interface{}, filter map[string]interface{}) error {
	result := dao.db.Where(filter).Delete(model)
	return result.Error
}

// Update обновляет записи по ID
func (dao *BaseDAO) Update(model interface{}, id uint, values map[string]interface{}) error {
	result := dao.db.Model(model).Where("id = ?", id).Updates(values)
	return result.Error
}



////////////////////////////////
///////////////////////////////


package main

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "myapp/dao"
    "myapp/models"
)

func main() {
    // Подключение к базе данных
    dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Ошибка подключения к БД: %s", err)
    }

    // Создаём экземпляр DAO
    dao := dao.NewBaseDAO(db)

    // Пример использования с таблицей User
    user := models.User{
        Name:     "John Doe",
        Email:    "johndoe@example.com",
        Password: "hashedpassword123",
    }

    // Добавление пользователя
    if err := dao.Add(&user); err != nil {
        log.Fatalf("Ошибка добавления пользователя: %s", err)
    }
    fmt.Printf("Пользователь успешно добавлен: %+v\n", user)

    // Поиск пользователя
    filter := map[string]interface{}{"email": "johndoe@example.com"}
    foundUser := &models.User{}
    result, err := dao.FindOneOrNone(foundUser, filter)
    if err != nil {
        log.Fatalf("Ошибка поиска пользователя: %s", err)
    }
    if result == nil {
        fmt.Println("Пользователь не найден")
    } else {
        fmt.Printf("Найден пользователь: %+v\n", result)
    }

    // Удаление пользователя
    if err := dao.Delete(&models.User{}, filter); err != nil {
        log.Fatalf("Ошибка удаления: %s", err)
    }
    fmt.Println("Пользователь успешно удалён")
}
