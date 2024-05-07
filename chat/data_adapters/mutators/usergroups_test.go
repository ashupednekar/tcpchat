package mutators

import (
	"testing"

	"github.com/ashupednekar/tcpchat/chat"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	db := chat.GetDb()
	CreateUser(db, "ashu", "192.168.0.1:50001")
	// assert.EqualErrorf(t, err, "no rows changed", "Error should be: %v, got: %v", "", err)
}

func TestMarkOnline(t *testing.T) {
	db := chat.GetDb()
	err := MarkOnline(db, "192.168.0.1:50001")
	if err != nil {
		t.Fatal(err)
	}
}

func TestMarkOnlineMissingIP(t *testing.T) {
	db := chat.GetDb()
	err := MarkOnline(db, "192.168.0.2:50001")
	assert.EqualErrorf(t, err, "no rows changed", "Error should be: %v, got: %v", "", err)
}

func TestMarkOffline(t *testing.T) {
	db := chat.GetDb()
	err := MarkOffline(db, "192.168.0.1:50001")
	if err != nil {
		t.Fatal(err)
	}
}

func TestMarkOfflineMissingIP(t *testing.T) {
	db := chat.GetDb()
	err := MarkOnline(db, "192.168.0.2:50001")
	assert.EqualErrorf(t, err, "no rows changed", "Error should be: %v, got: %v", "", err)
}

func TestJoinNewGroup(t *testing.T) {
	db := chat.GetDb()
	err := JoinGroup(db, "techtalks", "192.168.0.1:50001")
	if err != nil {
		t.Fatal(err)
	}
}

func TestJoinExistingGroup(t *testing.T) {
	db := chat.GetDb()
	err := JoinGroup(db, "techtalks", "192.168.0.1:50001")
	if err != nil {
		t.Fatal(err)
	}
	err1 := CreateUser(db, "jane", "192.168.0.2:50002")
	if err1 != nil {
		t.Fatal(err1)
	}
	err2 := JoinGroup(db, "techtalks", "192.168.0.2:50002") // TODO: fix, currently failing with sql: Scan error on column index 2, name "users": unsupported Scan, storing driver.Value type []uint8 into type *[]int
	if err2 != nil {
		t.Fatal(err2)
	}
}
