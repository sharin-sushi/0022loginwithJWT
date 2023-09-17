package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/sharin-sushi/0022loginwithJWT/t0022Go/internal/controller/crypto"

	validation "github.com/go-ozzo/ozzo-validation"
)

//最初、シンボル変更できなかったので、どこかで変更残りがあるかも User→Member
type Member struct { //dbに対してはtable名 小文字かつ複数形に自動変換
	//gorm.Model CreatedAtは機能無し
	MemberId   string `gorm:"primaryKey"`
	MemberName string
	Email      string
	Password   string
	CreatedAt  time.Time
}

type EntryMember struct {
	//gorm.Model CreatedAtは機能無し
	// MemberId   string
	MemberName string
	Email      string
	Password   string
}

// 自前のbycrypt関数があるので、それを使うためコメントアウト化。いずれ消す。
// →crypto.PasswordEncryptNoBackErr
// func Encrypt(char string) string {
// 	encryptText := fmt.Sprintf("%x", sha256.Sum256([]byte(char)))
// 	return encryptText
// }

func (m *Member) CreateMember(db *gorm.DB) (Member, error) { //Member構造体の型で新規発行したIDと共にユーザー情報を返す
	fmt.Printf("CreateMemberで使用されるm= %v \n", m)

	user := Member{
		MemberName: m.MemberName,
		Email:      m.Email,
		Password:   crypto.PasswordEncryptNoBackErr(m.Password),
	}
	// ここまで動作確認

	newId, err := CreateNewUserId(db) //最新ユーザーのidから新規ユーザーidを発行
	if err != nil {
		fmt.Printf("Failed create a new id")
		return user, err
	}
	fmt.Println(2.3)
	user.MemberId = newId
	fmt.Printf("新規id込みでuser= %v \n", user)
	result := db.Create(&user)
	if result != nil {
		return user, result.Error
	}
	fmt.Println(2.4)
	user, _ = FindUserByEmail(db, m.Email) //user情報取得
	return user, result.Error
}

// 最低文字数の制限は今のところここでしかやってない、元々1, 255。8, 255だった。
// import "validation "github.com/go-ozzo/ozzo-validation"
func (m *Member) Validate() error {
	err := validation.ValidateStruct(m,
		validation.Field(&m.MemberName,
			validation.Required.Error("Name is requred"),
			validation.Length(2, 20).Error("Name needs 2~20 cahrs"),
		),
		validation.Field(&m.Password,
			validation.Required.Error("Password is required"),
			validation.Length(4, 20).Error("Password needs 4 ~ 20 chars"),
		),
		validation.Field(&m.Email,
			validation.Required.Error("Email is required"),
			validation.Length(10, 100).Error("Email needs 4 ~ 20 chars"), //メアドは現状、これ以外の制限はしてない
		),
	)
	return err
}

// members
// +-------------+--------------+------+-----+-------------------+-------------------+
// | Field       | Type         | Null | Key | Default           | Extra             |
// +-------------+--------------+------+-----+-------------------+-------------------+
// | member_id   | varchar(4)   | NO   | PRI | NULL              |                   |
// | member_name | varchar(20)  | NO   | MUL | NULL              |                   |
// | email       | varchar(100) | NO   | UNI | NULL              |                   |
// | password    | varchar(100) | NO   |     | NULL              |                   |
// | created_at  | datetime     | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
// +-------------+--------------+------+-----+-------------------+-------------------+

func FindUserByEmail(db *gorm.DB, email string) (Member, error) {
	var user Member
	result := db.Where("email = ?", email).First(&user)
	fmt.Printf("%v \n", user)
	return user, result.Error
}

func CreateNewUserId(db *gorm.DB) (string, error) { //最新ユーザーのIDを取得し、+1して返す
	var lastUser Member
	result := db.Select("member_id ").Last(&lastUser)
	// SELECT member_id From members  ORDER BY member_id DESC LIMIT 1;
	if result.Error != nil {
		fmt.Println("最新ユーザーのid取得に失敗しました。error:", result.Error)
	}
	fmt.Printf("lastUser= %v", lastUser)

	fmt.Printf("lastUser.MemberId= %v", lastUser.MemberId)

	parts := strings.Split(lastUser.MemberId, "L") //　"", "1"に分ける(Lは消える)
	fmt.Printf("parts= %v", parts)

	if len(parts) != 2 { // Lで分割し\、要素数が2でなければエラー
		return "", errors.New("invalid MemberId format")
	}
	fmt.Println(3.2)

	lastUserIdNum := parts[1]
	fmt.Println(3.3)
	s, _ := strconv.Atoi(lastUserIdNum)
	s++
	fmt.Printf("newIdNum= %v \n", s)
	i := strconv.Itoa(s)
	fmt.Println(3.4)
	newId := "L" + i
	fmt.Printf("newId= %v \n", newId)

	return newId, nil
}
