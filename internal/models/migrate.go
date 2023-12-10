package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	go db.AutoMigrate(&SmsCode{})
	go db.AutoMigrate(&User{})
	go db.AutoMigrate(&WorkCategory{})
	go db.AutoMigrate(&BusinessProfile{})
	go db.AutoMigrate(&Booking{})
	go db.AutoMigrate(&Following{})
	go db.AutoMigrate(&Post{})
	go db.AutoMigrate(&BookingCategory{})
	go db.AutoMigrate(&PostViewed{})
}
