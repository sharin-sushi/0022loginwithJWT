package postrequest

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/sharin-sushi/0022loginwithJWT/t0022Go/internal/controller/model"
	"github.com/sharin-sushi/0022loginwithJWT/t0022Go/internal/types"
)

// /signup
func PostSignup(c *gin.Context) {
	// var userinfo types.User
	var json types.UserInfoFromFront
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// GPTより、
	// func SomeHandler(c *gin.Context) {
	// 	loggedIn, err := c.Cookie("loggedIn")
	// 	if err != nil {
	// 		// クッキーが存在しない場合の処理
	// 		return
	// 	}

	// 	if loggedIn == "true" {
	// 		// ログインしている場合の処理
	// 	} else {
	// 		// ログインしていない場合の処理
	// 	}
	// }
	var form types.Member

	//loginのやつコピペしたので、要修正
	if c.ShouldBind(&form) == nil {
		if form.MemberName == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"userinfo":    data,
	// })

	//要対応ーーーーーーーーーーーーーー
	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	acn := c.PostForm("accname")
	eml := c.PostForm("email")

	// id := "L2"
	// pw := ""
	// acn := "name"
	// eml :=

	// fmt.Printf("id=%v, pw=%v \n", id, pw)
	user, err := model.Signup(id, pw, acn, eml)
	fmt.Printf("user=%v, err=%v \n", user, err)

	if err != nil {
		c.Redirect(301, "/signup")
		return
	}

	//要対応ーーーーーーーーーーーーーーーーー
	c.HTML(http.StatusOK, "home.html", gin.H{"user": user})
}

// "/login"

func PostLogin(c *gin.Context) {
	type test struct {
		csrfToken string
		types.Member
	}

	var form test
	if c.ShouldBind(&form) != nil {
		c.JSON(401, gin.H{"status": "ログイン情報の送信に失敗しました"})
		// if form.MemberName == "user" && form.Password == "password" { //書き方間違えてない？正しくはその時のセッションから取得して比較する？目的がわからん。
		// 	c.JSON(200, gin.H{"status": "ログイン済み"})
		// } else {
		// 	// c.JSON(401, gin.H{"status": "unauthorized"}) //動作確認ok
		// }
	}
	// fmt.Printf("bitしたform=%v \n", form)

	fmt.Printf("bind内容:crsfToken=%v \n", form.csrfToken)

	fmt.Printf("bind内容:ID=%v, name=%v, mail=%v, pass=%v, crat%v \n", form.MemberId, form.MemberName, form.Email, form.Password, form.CreatedAt)
	//ここまでは処理確認ok
	member, err := model.InquireIntoMember(form.MemberName, form.Password)
	if err != nil {
		c.Redirect(301, "/login")
		return
	}
	fmt.Printf("DBから取得した情報=%+v \n ", member) //%vでも%sでも L1[GIN]

	jwtToken, err := GenerateToken(member.MemberId)
	if err != nil {
		fmt.Print("トークンの生成に失敗しました。")
		c.Redirect(301, "/login")
		return
	}

	// sessionID := generateSessionID() →代わりにJWTtokenを使用
	// sessionManage(c, *member)

	c.SetCookie("loginToken", jwtToken, 3600, "/", "localhost", false, true) //最後のとこ　HttpOnly trueならフロントで読み取れない
	// POSTMANでCookieにセットされていること確認ok

	// フロント側の仕様
	//サーバーがJSON形式のレスポンスを返し、その中に
	///authenticated というフィールドが含まれていることを仮定している？
	// true であれば認証成功

	responseData := gin.H{
		"memberId":   member.MemberId,
		"memberName": member.MemberName,
		"message":    "Login successful",
	}
	fmt.Println(responseData)
	c.JSON(200, responseData)
}

func PostLogout(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("セッション取得")
	session.Clear()
	log.Println("クリア処理")
	session.Save()
}

// JWTtoken使用により不要になった関数？
//SessionManager.go　　セッション生成(alive, ID, 名前)
func sessionManage(g *gin.Context, user types.Member) {
	session := sessions.Default(g) //現在のセッションを取得
	session.Set("alive", true)     //取得したセッションに値を設定
	session.Set("memberID", user.MemberId)
	session.Set("memberName", user.MemberName)
	err := session.Save() //セッションの変更を保存

	if err != nil {
		log.Printf("Error saving session: %v", err)
	}
}

// JWTtoken使用により不要になった関数？
func generateSessionID() string {
	b := make([]byte, 16) // 16バイトのランダムデータを生成
	rand.Read(b)

	return hex.EncodeToString(b)

}
