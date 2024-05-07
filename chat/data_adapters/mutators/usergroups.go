package mutators

import (
	"fmt"

	"github.com/ashupednekar/tcpchat/chat"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, Name string, IP string) error {
	user := chat.Users{
		Name:   Name,
		IP:     IP,
		Online: true,
	}
	r1 := db.Create(&user)
	if r1.Error != nil {
		return r1.Error
	}
	individual_grp := []int{user.ID}
	CreateGroup(db, Name, individual_grp)
	return nil
}

func CreateGroup(db *gorm.DB, Name string, Users []int) error {
	r2 := db.Create(&chat.Groups{
		Name:  Name,
		Users: Users,
	})
	if r2.Error != nil {
		return r2.Error
	}
	return nil
}

func MarkOnline(db *gorm.DB, IP string) error {
	r := db.Model(&chat.Users{}).Where("ip = ?", IP).Update("online", true)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != 1 {
		return fmt.Errorf("no rows changed")
	}
	return nil
}

func MarkOffline(db *gorm.DB, IP string) error {
	r := db.Model(&chat.Users{}).Where("ip = ?", IP).Update("online", false)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected != 1 {
		return fmt.Errorf("no rows changed")
	}
	fmt.Println(r.RowsAffected)
	return nil
}

func JoinGroup(db *gorm.DB, Name string, IP string) error {
	user := chat.Users{}
	r := db.Model(&chat.Users{}).Where("ip = ?", IP).First(&user)
	if r.Error != nil {
		return r.Error
	}
	group := chat.Groups{}
	grp_query := db.Model(&chat.Groups{}).Where("name = ?", Name)
	r1 := grp_query.First(&group)
	if r1.Error != nil {
		fmt.Println("err: ", r1.Error)
		if r1.Error.Error() == "record not found" {
			fmt.Println("group not present, creating...")
			individual_grp := []int{user.ID}
			fmt.Println("new grp name: ", Name)
			return CreateGroup(db, Name, individual_grp)
		} else {
			return r1.Error
		}
	}
	group.Users = append(group.Users, user.ID)
	r2 := grp_query.Update("users", group.Users)
	if r2.Error != nil {
		return r2.Error
	}
	if r2.RowsAffected != 1 {
		return fmt.Errorf("no rows changed")
	}
	return nil
}
