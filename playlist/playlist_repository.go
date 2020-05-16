package playlist

import (
	"api/utils"

	"github.com/jmoiron/sqlx"
)

// PlayList db 플레이리스트 저장 관련정보
type PlayList struct {
	PlayListID   int64
	UserID       int64
	PlayListName string
	SongIDs      []int64
}

// NewPlayList 새로운 플레이리스트를 만들고 아이디를 반환합니다
func NewPlayList(userID int, name string) int64 {
	var db *sqlx.DB = utils.GetDB()
	var newPlayListID int64

	res := db.MustExec(`insert into play_list (user_id, name) values(?,?);`, userID, name)
	newPlayListID, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return newPlayListID
}

// AddSongPlayList a
func AddSongPlayList(userID, playListID, songID int) {
	// var db *sqlx.DB = utils.GetDB()

}
