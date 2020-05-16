package album

import (
	"api/utils"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

func Test_FindById(t *testing.T) {
	var db *sqlx.DB = utils.GetDB()
	db.Exec("insert into album values(9009199, 'test');")
	a := FindAlbumByID(9009199)
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

func Test_InsertAlbum(t *testing.T) {
	a := Album{}
	a.AlbumID = 9999999
	a.Title = "test"
	InsertAlbum(&a)

	g := FindAlbumByID(a.AlbumID)
	if g.Title != a.Title {
		t.Error("must title equal")
	}
}

func Test_SaveAlbum(t *testing.T) {
	a := Album{}
	a.AlbumID = 9999997
	a.Title = "test_save"

	SaveAlbum(&a)

	g := FindAlbumByID(a.AlbumID)
	if g.Title != a.Title {
		t.Error("must title equal")
	}

	a.Title = "test_save2"
	SaveAlbum(&a)
}

func Test_DeleteAlbum(t *testing.T) {
	a := Album{}
	a.AlbumID = 9999998
	a.Title = "test_save"

	SaveAlbum(&a)

	DeleteAlbumByAlbumID(a.AlbumID)

	if FindAlbumByID(a.AlbumID) != nil {
		t.Error("album not deleted")
	}
}

func TestMain(m *testing.M) {
	utils.BeginTransactionForTest()
	ret := m.Run()
	utils.RollbackTransactionForTest()
	os.Exit(ret)
}
