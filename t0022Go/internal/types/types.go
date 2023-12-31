package types

import "time"

//"github.com/sharin-sushi/0022loginwithJWT/internal/types"
// types.

// var all types.AllData
// var alls []types.AllData
// var st types.Streamer
// var sts []types.Streamer
// var mo types.Movie
// var mos []types.Movie
// var ka types.KaraokeList
// var kas []types.KaraokeList

type StremerMovie struct {
	Streamer
	Movie
}

//GORMなのに何故か `json:"~~"`が無いとスネークにならない

type Streamer struct {
	StreamerId      int     `gorm:"primaryKey"` //`json:"streamer_id"` //
	StreamerName    string  //`json:"streamer_name"`
	NameKana        *string //`json:"name_kana"`
	SelfIntroUrl    *string //`json:"self_intro_url"`
	StreamInputerId *string //`json:"stream_inputer_id"`
}

type Movie struct {
	StreamerId *int    //`json:"streamer_id"`
	MovieId    *int    //`json:"movie_id"`
	MovieUrl   string  `gorm:"primaryKey"` //`json:"movie_url"`
	MovieTitle *string //`json:"movie_title"`
}

type KaraokeList struct {
	MovieUrl      string  `gorm:"primaryKey"` //
	SongId        int     `gorm:"primaryKey"` //`json:"song_id"`
	SingStart     *string //`json:"sing_start"` //nill可にするためのポインタ
	Song          string  //`json:"song"`
	SongInputerId string  //`json:"song_inputer_id"`
}

// コピペ用全カラム
// ーーキャメル
type AllColumns struct {
	StreamerId       int
	StreamerName     string
	NameKana         string
	SelfIntro_url    string
	StreamInputer_id string
	MovieId          int
	MovieUrl         string
	MovieTitle       string
	SongId           int
	SingStart        *string
	Song             string
	SongInputerId    string
}

//ーーースネーク
// streamer_id
// streamer_name
// name_kana
// self_intro_url
// stream_inputer_id
// movie_id
// movie_url
// movie_title
// song_id
// sing_start
// song
// song_inputer_id

//最初、シンボル変更できなかったので、どこかで変更残りがあるかも User→Member
type Member struct { //dbに対してはtable名 小文字かつ複数形に自動変換
	//gorm.Model CreatedAtは機能無し
	MemberId   string `gorm:"primaryKey"`
	MemberName string
	Email      string
	Password   string
	CreatedAt  time.Time
}

type UserInfoFromFront struct { //dbに対してはtable名 小文字かつ複数形に自動変換
	//gorm.Model CreatedAtは機能無し
	MemberId   string
	MemberName string
	Email      string
	Password   string
}
