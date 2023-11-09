package mysql

import (
	"context"
	"project/server/shared/consts"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"

	"gorm.io/gorm"
)

const (
	Active = "active"
)

// User 用户模型
type User struct {
	gorm.Model
	UserId      int64      `json:"user_id" gorm:"index;unique;not null;comment:用户user_id"`
	LikeNum     int64      `json:"like_num"`
	Birthday    *time.Time `json:"birthday"`
	Gender      int8       `json:"gender"   gorm:"size:1"`
	Type        int8       `json:"type"   gorm:"size:5;comment:是否wx登录"`
	Enable      int8       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	CommentNum  uint32     `json:"comment_num"`
	ArticleNum  uint32     `json:"article_num"`
	Status      string     `json:"status"   gorm:"size:10"`
	Company     string     `json:"company"   gorm:"size:500"`
	WxOpenid    string     `json:"wx_openid"   gorm:"size:500"`
	RealName    string     `json:"real_name" gorm:"size:120"`
	NickName    string     `json:"nick_name" gorm:"size:120"`
	UserName    string     `json:"user_name"  gorm:"size:120"`
	Password    string     `json:"-"  gorm:"size:120"`
	Mobile      string     `json:"mobile"  gorm:"size:11"`
	Email       string     `json:"email" gorm:"size:120"`
	Blog        string     `json:"facebook"   gorm:"size:3000"`
	Avatar      string     `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	Description string     `json:"description"  gorm:"default:Ta很懒，还没有添加简介"`
	Location    string     `json:"location"   gorm:"size:500"`
	School      string     `json:"school"   gorm:"size:500"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	if u.UserId != 0 {
		return nil
	}

	sf, err := snowflake.NewNode(consts.UserSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	u.UserId = sf.Generate().Int64()

	return nil
}

type UserManager struct {
	salt string
	db   *gorm.DB
}

func (m *UserManager) ExistOrNotByMobile(ctx context.Context, mobile string) (user *User, exist bool, err error) {
	//TODO implement me
	var count int64
	err = m.db.WithContext(ctx).Model(&User{}).Where("mobile = ?", mobile).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.db.WithContext(ctx).Model(&User{}).Where("mobile = ?", mobile).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func (m *UserManager) CreateUser(ctx context.Context, user *User) error {
	return m.db.WithContext(ctx).Model(&User{}).Create(&user).Error
}

// NewUserManager creates a mysql manager.
func NewUserManager(db *gorm.DB, salt string) *UserManager {
	m := db.Migrator()
	if !m.HasTable(&User{}) {
		if err := m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	return &UserManager{
		db:   db,
		salt: salt,
	}
}
