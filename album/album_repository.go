package album

import (
	"api/utils"
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
