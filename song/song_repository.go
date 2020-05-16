package song

import (
	"api/utils"

	"github.com/jmoiron/sqlx"
)

// Song 앨범에 들어가는 노래
type Song struct {
	SongID  int    `db:"idx" json:"song_id"`
	Length  int    `db:"length" json:"length"`
	Title   string `db:"title" json:"title"`
	Track   int    `db:"track" json:"track"`
	AlbumID int    `db:"album_id" json:"album_id"`
}

// FindSongByID ID 로 Song 을 가져옵니다
func FindSongByID(songID int) *Song {
	var db *sqlx.DB = utils.GetDB()
	s := Song{}
	if err := db.Get(&s, `select 
		song_id,
		length,
		title,
		track,
		album_id 
	from song where idx=? limit 1;`, songID); err != nil {
		panic(err)
	}
	return &s
}
