package internal

import (
	"api-rest/internal/business"
	"api-rest/internal/platform"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Email             string
	Password          string
	Pass_confirmation string
	Name              string
	Last_name         string
}

type loginRequest struct {
	Email    string
	Password string
}

type Handlers struct {
}

func (h *Handlers) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "holis",
	})

}

func (h *Handlers) RegisterHandler(c *gin.Context) {
	var user registerRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db, _ := platform.DbConnection()

	if user.Password != user.Pass_confirmation {
		return
	}

	var err error
	user.Password, err = business.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong with the hash",
		})
	}

	mUser := business.User{
		Email:     user.Email,
		Password:  user.Password,
		Name:      user.Name,
		Last_name: user.Last_name,
	}

	result := db.Create(&mUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func verifyPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func (h *Handlers) LoginHandler(c *gin.Context) {
	var input loginRequest

	//binding json
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//connect DB
	db, _ := platform.DbConnection()

	var user business.User

	//query Email
	result := db.Model(user).Where("email = ?", input.Email).Take(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong with the Email",
		})
		return
	}
	fmt.Printf("%v", user)

	//password
	err := verifyPassword(user.Password, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong with the password",
		})
		return
	}

	fmt.Println(GenerateToken(user.User_id))

	c.JSON(http.StatusOK, input)

}

func (h *Handlers) GetPostHandler(c *gin.Context) {
	var post []business.Post

	db, _ := platform.DbConnection()

	result := db.Preload("User").Find(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	for i := range post {
		db.Model(&post[i]).Association("User").Find(&post[i].User)
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handlers) GetPostByIdHandler(c *gin.Context) {
	id := c.Param("id")

	var post []business.Post
	db, _ := platform.DbConnection()

	result := db.Where("post_id = ?", id).Preload("User").Find(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handlers) NewPostHandler(c *gin.Context) {
	var post business.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get user_id from header
	user_id, _ := c.Get("user_id")
	userId := user_id.(float64)
	post.User_id = int(userId)

	db, _ := platform.DbConnection()
	result := db.Preload("user").Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *Handlers) DeletePostHandler(c *gin.Context) {
	id := c.Param("id")
	var post []business.Post

	db, _ := platform.DbConnection()

	delete := db.Where("post_id = ?", id).Delete(&post)

	if delete.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handlers) GetCommentsHandler(c *gin.Context) {
	id := c.Param("id")

	var comments []business.Comments

	db, _ := platform.DbConnection()

	result := db.Where("post_id = ?", id).Preload("User").Find(&comments)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *Handlers) NewCommentHandler(c *gin.Context) {
	id := c.Param("id")

	var comments business.Comments

	if err := c.ShouldBindJSON(&comments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user_id, _ := c.Get("user_id")
	userId := user_id.(float64)
	comments.User_id = int(userId)

	db, _ := platform.DbConnection()

	// Check if the post exists
	var post business.Post

	res := db.First(&post, id)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Post not found",
		})
		return
	}

	comments.Post_id = post.Post_id

	result := db.Create(&comments)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}
