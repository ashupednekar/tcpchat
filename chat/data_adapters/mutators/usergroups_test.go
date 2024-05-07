package mutators

import (
	"testing"

	"github.com/ashupednekar/tcpchat/chat"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	db := chat.GetDb()
	CreateUser(db, "ashu", "192.168.0.1:50001")
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

func TestJoinGroup(t *testing.T) {
	db := chat.GetDb()
	JoinGroup(db, "techtalks", "192.168.0.1:50001")
}
