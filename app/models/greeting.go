package models

import (
	"fmt"

	"github.com/Longneko/lamp/app/lib/database"
)

const (
	defaultName = "anonymous"
)

type Greeting struct {
	BaseModel
	Name string `form:"name" gorm:"column:name"`
}

func (Greeting) TableName() string {
	return "greetings"
}

type GreetingRepository struct {
	db database.Conn
}

func NewGreetingRepository(db database.Conn) (repo *GreetingRepository, err error) {
	if !db.IsInit() {
		err = fmt.Errorf("got uninitialized db")
		return
	}

	repo = &GreetingRepository{db}
	return
}

func NewDefaultDbGreetingRepo() (repo *GreetingRepository, err error) {
	db, err := database.GetDb()
	if err != nil {
		err = fmt.Errorf("error while getting default db for GreetingRepository: `%s`", err)
		return
	}

	repo, err = NewGreetingRepository(db)
	if err != nil {
		err = fmt.Errorf("error while constructing new GreetingRepository with default db: `%s`", err)
		return
	}
	return
}

func (r *GreetingRepository) CreateTable() error {
	if r.db.HasTable(&Greeting{}) {
		// TODO: consider returning an easily checkable error
		return nil
	}
	return r.db.CreateTable(&Greeting{}).Error
}

func (r *GreetingRepository) Store(g Greeting) error {
	if g.Name == "" {
		g.Name = defaultName
	}
	return r.db.Create(&g).Error
}

func (r *GreetingRepository) GetAll() (greetings []Greeting, err error) {
	err = r.db.Order("id DESC").Find(&greetings).Error
	return
}
