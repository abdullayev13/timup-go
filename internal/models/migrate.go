package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&SmsCode{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&WorkCategory{})
	db.AutoMigrate(&BusinessProfile{})
	db.AutoMigrate(&Booking{})
}
