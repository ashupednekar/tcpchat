package selectors

import (
	"fmt"

	"github.com/ashupednekar/tcpchat/chat"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, IP string) (error, chat.Users) {
	user := chat.Users{}
	r := db.Model(&chat.Users{}).Where("ip = ?", IP).First(&user)
	return r.Error, user
}

func GetIPsFromGroupName(db *gorm.DB, Name string) (error, []string) {
	group := chat.Groups{}
	r := db.Model(&chat.Groups{}).Where("name = ?", Name).First(&group)
	if r.Error != nil {
		return nil, []string{}
	}
	var IPs []string
	var mappings []chat.GroupUserMap
	db.Model(&chat.GroupUserMap{}).Where("group_id = ?", group.ID).Find(&mappings)
	for _, mapping := range mappings {
		fmt.Println("ip: ", mapping.User)
		IPs = append(IPs, mapping.User.IP)
	}
	fmt.Println("IPs: ", IPs)
	return nil, []string{}
}

func GetGroup(db *gorm.DB, Name string) (error, chat.Groups) {
	grp := chat.Groups{}
	r := db.Model(&chat.Groups{}).Where("name = ?", Name).First(&grp)
	return r.Error, grp
}
