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

func (handler *Handler) SignUpHandler(ctx *gin.Context) {
	var signUpInput types.UserInfoFromFront
	err := ctx.ShouldBind(&signUpInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	member := &types.Member{
		MemberName: signUpInput.MemberName,
		Password:   signUpInput.Password,
	}
	err = member.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	newUser, err := member.Create(handler.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	// 追加部分
	token, err := utils.GenerateToken(newUser.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to sign up",
		})
		return
	}

	// Cookieにトークンをセット
	cookieMaxAge := 60 * 60
	ctx.SetCookie("token", token, cookieMaxAge, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"user_id": newUser.ID,
		"message": "Successfully created user",
	})
}
