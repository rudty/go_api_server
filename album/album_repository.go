package album

import (
	"api/utils"

	"github.com/jmoiron/sqlx"
)

// Album db 의 앨범 관련 정보
type Album struct {
	AlbumID int    `db:"album_id"`
	Title   string `db:"title"`
}

// FindAlbumByID 앨범 아이디로 앨범 1개를 가져옵니다.
func FindAlbumByID(albumID int) *Album {
	a := Album{}
	if utils.GetDB().Get(&a, "select album_id, title from album where album_id=?;", albumID) != nil {
		return nil
	}
	return &a
}

// InsertAlbum 앨범을 추가합니다
func InsertAlbum(a *Album) {
	var db *sqlx.DB = utils.GetDB()
	db.MustExec("insert into album (album_id, title) values(?,?);", a.AlbumID, a.Title)
}

// SaveAlbum 앨범을 저장합니다(수정, 삭제)
func SaveAlbum(a *Album) {
	var db *sqlx.DB = utils.GetDB()
	db.NamedExec(`
		insert into album (album_id, title) 
			values(:album_id,:title) 
		on duplicate key 
		update 
			album_id = :album_id, 
			title=:title;`, a)
}

// DeleteAlbumByAlbumID 앨범 아이디로 해당 앨범 1개를 삭제합니다
func DeleteAlbumByAlbumID(albumID int) {
	var db *sqlx.DB = utils.GetDB()
	db.MustExec("delete from album where album_id=?;", albumID)
}
