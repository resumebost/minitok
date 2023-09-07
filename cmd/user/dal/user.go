package dal

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"minitok/internal/dal"
	"time"
)

var GormDB *gorm.DB

type User struct {
	//gorm.Model
	ID              int64 `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `gorm:"index"`
	Username        string     `binding:"required"`
	Password        string     `binding:"required"`
	Avatar          string     `json:"omitempty"`
	BackgroundImage string     `gorm:"default:default_background.jpg"`
	Signature       string
}

func SetUserDB() {
	GormDB = dal.InitGorm()
}

func (User) TableName() string {
	return "users"
}

func GetUserByNameByRegister(ctx context.Context, userName string) (*User, error) {
	res := new(User)

	db := GormDB.WithContext(ctx)

	err := db.Model(&User{}).Select("username").Where("username = ?", userName).Find(&res).Error
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func CreateUser(ctx context.Context, user *User) error {
	err := GormDB.Clauses(dbresolver.Write).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func GetUserByNameByLogin(ctx context.Context, userName string) (*User, error) {
	res := new(User)

	db := GormDB.WithContext(ctx)

	err := db.Model(&User{}).Select("id", "password").Where("username = ?", userName).Find(&res).Error
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return res, nil

}

// GetUserByID 根据id获取用户数据
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)

	db := GormDB.WithContext(ctx)

	err := db.Model(&User{}).Where("id = ?", userID).Find(&res).Error
	fmt.Println(res.ID, res.Username)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetUsersByIDList(ctx context.Context, userIDList []int64) ([]*User, error) {
	var users = make([]*User, len(userIDList))

	for i, id := range userIDList {
		result := GormDB.WithContext(ctx).First(&users[i], id)
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return users, nil
}
