package selectors

import (
	"fmt"
	"testing"

	"github.com/ashupednekar/tcpchat/chat"
)

func TestGetIPFromGroupName(t *testing.T) {
	db := chat.GetDb()
	err, IPs := GetIPsFromGroupName(db, "techtalks")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("IPs: ", IPs)
}
