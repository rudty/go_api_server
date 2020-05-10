package album

import (
	"api/utils"
	"testing"

	"github.com/jmoiron/sqlx"
)

func Test_FindById(t *testing.T) {
	var db *sqlx.DB = utils.GetDB()
	db.Exec("insert into album values(1, 'test');")
	a := FindAlbumByID(1)
	if a == nil {
		t.Error("must not nil")
	}

}

func Test_FindById_NotExists(t *testing.T) {
	a := FindAlbumByID(9009999)
	if a != nil {
		t.Error("must nil")
	}
}
