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
	gorm.Model
	UserID          int64   `gorm:"unique;not null" json:"user_id"`
	UserName        string  `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"name,omitempty"`
	Password        string  `gorm:"type:varchar(256);not null" json:"password,omitempty"`
	FavoriteVideos  []Video `gorm:"many2many:user_favorite_videos" json:"favorite_videos,omitempty"`
	FollowingCount  uint    `gorm:"default:0;not null" json:"follow_count,omitempty"`                                                           // 关注总数
	FollowerCount   uint    `gorm:"default:0;not null" json:"follower_count,omitempty"`                                                         // 粉丝总数
	Avatar          string  `gorm:"type:varchar(256)" json:"avatar,omitempty"`                                                                  // 用户头像
	BackgroundImage string  `gorm:"column:background_image;type:varchar(256);default:default_background.jpg" json:"background_image,omitempty"` // 用户个人页顶部大图
	WorkCount       uint    `gorm:"default:0;not null" json:"work_count,omitempty"`                                                             // 作品数
	FavoriteCount   uint    `gorm:"default:0;not null" json:"favorite_count,omitempty"`                                                         // 喜欢数
	TotalFavorited  uint    `gorm:"default:0;not null" json:"total_favorited,omitempty"`                                                        // 获赞总量
	Signature       string  `gorm:"type:varchar(256)" json:"signature,omitempty"`                                                               // 个人简介
}

type Video struct {
	ID            uint      `gorm:"primarykey"`
	CreatedAt     time.Time `gorm:"not null;index:idx_create" json:"created_at,omitempty"`
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Author        User           `gorm:"foreignkey:AuthorID" json:"author,omitempty"`
	AuthorID      uint           `gorm:"index:idx_authorid;not null" json:"author_id,omitempty"`
	PlayUrl       string         `gorm:"type:varchar(255);not null" json:"play_url,omitempty"`
	CoverUrl      string         `gorm:"type:varchar(255)" json:"cover_url,omitempty"`
	FavoriteCount uint           `gorm:"default:0;not null" json:"favorite_count,omitempty"`
	CommentCount  uint           `gorm:"default:0;not null" json:"comment_count,omitempty"`
	Title         string         `gorm:"type:varchar(50);not null" json:"title,omitempty"`
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

	err := db.Model(&User{}).Select("user_name").Where("user_name = ?", userName).Find(&res).Error
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

	err := db.Model(&User{}).Select("user_id", "password").Where("user_name = ?", userName).Find(&res).Error
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

	err := db.Model(&User{}).Where("user_id = ?", userID).Find(&res).Error
	fmt.Println(res.UserID, res.UserName)
	if err != nil {
		return nil, err
	}

	return res, nil
}
