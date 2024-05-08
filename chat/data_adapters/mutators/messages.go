package mutators

import (
	"database/sql"

	"github.com/ashupednekar/tcpchat/chat"
	"github.com/ashupednekar/tcpchat/chat/data_adapters/selectors"
	"gorm.io/gorm"
)

func SaveTextMessage(db *gorm.DB, FromIP string, ToName string, Text string) error {
	err, FromUser := selectors.GetUser(db, FromIP)
	if err != nil {
		return err
	}
	err1, ToGroup := selectors.GetGroup(db, ToName)
	if err1 != nil {
		return err
	}
	message := chat.Messages{
		From: FromUser,
		To:   ToGroup,
		Text: sql.NullString{
			String: Text,
			Valid:  true,
		},
	}
	r := db.Create(&message)
	return r.Error
}

func SaveAttachmentMessage(db *gorm.DB, FromIP string, ToName string, File byte) error { // TODO: later
	err, FromUser := selectors.GetUser(db, FromIP)
	if err != nil {
		return err
	}
	err1, ToGroup := selectors.GetGroup(db, ToName)
	if err1 != nil {
		return err
	}
	message := chat.Messages{
		From: FromUser,
		To:   ToGroup,
		File: sql.NullByte{
			Byte:  File,
			Valid: true,
		},
	}
	r := db.Create(&message)
	return r.Error
}
