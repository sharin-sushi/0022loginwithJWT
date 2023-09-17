package utility

//"github.com/sharin-sushi/0022loginwithJWT/internal/utility"

import (
	// "database/sql"
	"fmt"
	"net/http"
	"os"

	// "time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sharin-sushi/0022loginwithJWT/t0022Go/internal/controller/crypto"
	"github.com/sharin-sushi/0022loginwithJWT/t0022Go/internal/types"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DbGo *sql.DB
var Db *gorm.DB

type Handler struct {
	DB *gorm.DB
}

// init packageがimportされたときに１度だけ自動で呼び出される
func init() {
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DATABASE")

	var path string = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
	// sqlへ接続するための文字列の生成
	var err error

	// fmt.Printf("%s\n%s\n", path, err)

	Db, err = gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	if Db == nil {
		panic("failed to connect database")

	} //このif Db文消したい意味的に重複してる

	fmt.Printf("path=%s\n, err=%s\n", path, err)
	// checkConnect(1)
	// defer Db.Close()
}

// 会員登録
// func CalltoSignUpHandler(){
// 	h := Handler()
// 	a := h.Handler(c)
// }

// func CalltoSignUpHandler(r *gin.RouterGroup, h *controllers.Handler) {
//     auth := r.group("/auth")
//     {
//         auth.POST("/signup", h.SignUpHandler)
//     }
// }
func CalltoSignUpHandler(c *gin.Context) {
	h := Handler{DB: Db}
	h.SignUpHandler(c)
}
func (h *Handler) SignUpHandler(c *gin.Context) {
	var signUpInput types.EntryMember
	err := c.ShouldBind(&signUpInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	fmt.Printf("bindしたsignUpInput = %v \n", signUpInput)

	existingUser, _ := types.FindUserByEmail(h.DB, signUpInput.Email) //メアドが未使用ならnil
	if existingUser.MemberId != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "the E-mail address already in use",
		})
		return
	}

	member := &types.Member{
		MemberName: signUpInput.MemberName,
		Password:   signUpInput.Password,
		Email:      signUpInput.Email,
	}

	err = member.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newMember, err := member.CreateMember(h.DB) //Member構造体の型で新規発行したIDと共にユーザー情報を返す
	if err != nil {
		fmt.Printf("新規idのerr= %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user or find user after it",
		})
		return
	}
	fmt.Printf("新規id=%v \n", newMember)
	//ここまで動作確認ok

	// Token発行　＝　JWTでいいのかな？
	token, err := GenerateToken(newMember.MemberId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign up",
		})
		return
	}

	// Cookieにトークンをセット
	cookieMaxAge := 60 * 60
	c.SetCookie("token", token, cookieMaxAge, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"memberId":   newMember.MemberId,
		"memberName": newMember.MemberName,
		"message":    "Successfully created user, and logined",
	})
}

//ログイン
func CalltoLogInHandler(c *gin.Context) {
	h := Handler{DB: Db}
	h.LoginHandler(c)
}
func (h *Handler) LoginHandler(c *gin.Context) {
	var loginInput types.Member
	if err := c.ShouldBind(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	user, err := types.FindUserByEmail(h.DB, loginInput.Email) //メアドが未登録なら err = !nil
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "the E-mail address id NOTalready in use",
		})
		return
	}

	CheckPassErr := crypto.CompareHashAndPassword(user.Password, loginInput.Password) //pass認証
	if CheckPassErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is invalid",
		})
		return
	}
	fmt.Printf("ChechkPassErr=%v \n", CheckPassErr)

	// Token発行　＝　JWTでいいのかな？
	token, err := GenerateToken(user.MemberId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign up",
		})
		return
	}

	// Cookieにトークンをセット
	cookieMaxAge := 60 * 60
	c.SetCookie("token", token, cookieMaxAge, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Successfully logged in",
		"memberId":   user.MemberId,
		"memberName": user.MemberName,
	})
}
