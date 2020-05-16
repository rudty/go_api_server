package song

import (
	"api/utils"
	"fmt"
	"os"
	"testing"
)

func Test_FindSongById(t *testing.T) {
	s := FindSongByID(1)
	if s == nil {
		panic("must exists")
	}
	fmt.Println(s)
}
func TestMain(m *testing.M) {
	utils.BeginTransactionForTest()
	ret := m.Run()
	utils.RollbackTransactionForTest()
	os.Exit(ret)
}
