package song

import (
	"api/utils"

	"github.com/jmoiron/sqlx"
)

// Song 앨범에 들어가는 노래
type Song struct {
	SongID  int    `db:"idx"`
	Length  int    `db:"length"`
	Title   string `db:"title"`
	Track   int    `db:"track"`
	AlbumID int    `db:"album_id"`
}

// FindSongByID ID 로 Song 을 가져옵니다
func FindSongByID(songID int) *Song {
	var db *sqlx.DB = utils.GetDB()
	s := Song{}
	if err := db.Get(&s, "select * from song where idx=? limit 1;", songID); err != nil {
		panic(err)
	}
	return &s
}
