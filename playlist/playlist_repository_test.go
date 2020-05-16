package playlist

import (
	"api/utils"
	"os"
	"testing"
)

func Test_NewPlayList(t *testing.T) {
	playListID := NewPlayList(1, "aa")
	if playListID == 0 {
		t.Error("play list > 0")
	}
}

func TestMain(m *testing.M) {
	utils.BeginTransactionForTest()
	ret := m.Run()
	utils.RollbackTransactionForTest()
	os.Exit(ret)
}
