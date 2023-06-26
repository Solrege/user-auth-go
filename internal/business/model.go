package business

import "golang.org/x/crypto/bcrypt"

/*type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}*/

func (User) TableName() string {
	return "User"
}

type User struct {
	User_id   int    `gorm:"primaryKey"`
	Email     string `json:"-"`
	Password  string `json:"-"`
	Name      string
	Last_name string
	//Birth             CustomTime
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (Post) TableName() string {
	return "Post"
}

type Post struct {
	Post_id int `gorm:"primaryKey"`
	User_id int
	Text    string
	User    User `gorm:"foreignKey:User_id"`
}

func (Comments) TableName() string {
	return "Comments"
}

type Comments struct {
	Comment_id int `gorm:"primaryKey"`
	Post_id    int
	User_id    int
	Comment    string
	User       User `gorm:"foreignKey:User_id"`
	Post       Post `gorm:"foreignKey:Post_id"`
}

type Likes struct {
	Like_id int
	User_id int
	Post_id int
}

type Relationships struct {
	Id_relationships int
	Follower_user_id int
	Followed_user_id int
}
