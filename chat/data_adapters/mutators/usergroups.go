package mutators

import (
	"fmt"

	"github.com/ashupednekar/tcpchat/chat"
	"github.com/ashupednekar/tcpchat/chat/data_adapters/selectors"
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
	err, grp := CreateGroup(db, Name)
	err1 := GroupAdd(db, grp, user)
	if err1 != nil {
		return err1
	}
	return err
}

func GroupAdd(db *gorm.DB, Group chat.Groups, User chat.Users) error {
	fmt.Printf("Adding user %s to group %s\n", User.Name, Group.Name)
	r := db.Create(&chat.GroupUserMap{
		Group: Group,
		User:  User,
	})
	return r.Error
}

func CreateGroup(db *gorm.DB, Name string) (error, chat.Groups) {
	grp := chat.Groups{Name: Name}
	r := db.Create(&grp)
	return r.Error, grp
}

func MarkOnline(db *gorm.DB, IP string) error {
	r := db.Model(&chat.Users{}).Where("ip = ?", IP).Update("online", true)
	if r.RowsAffected != 1 {
		return fmt.Errorf("no rows changed")
	}
	return r.Error
}

func MarkOffline(db *gorm.DB, IP string) error {
	r := db.Model(&chat.Users{}).Where("ip = ?", IP).Update("online", false)
	if r.RowsAffected != 1 {
		return fmt.Errorf("no rows changed")
	}
	fmt.Println(r.RowsAffected)
	return r.Error
}

func JoinGroup(db *gorm.DB, Name string, IP string) error {
	err, user := selectors.GetUser(db, IP)
	if err != nil {
		return err
	}
	group := chat.Groups{}
	grp_query := db.Model(&chat.Groups{}).Where("name = ?", Name)
	r1 := grp_query.First(&group)
	if r1.Error != nil {
		fmt.Println("err: ", r1.Error)
		if r1.Error.Error() == "record not found" {
			fmt.Println("group not present, creating...")
			fmt.Println("new group name: ", Name)
			err, grp := CreateGroup(db, Name)
			err1 := GroupAdd(db, grp, user)
			if err1 != nil {
				return err1
			}
			return err
		} else {
			return r1.Error
		}
	}
	return nil
}
