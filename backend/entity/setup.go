package entity

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"time"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&User{},
		&Level{},
		&Map_Bed{},
		&Symptom{},
	)

	// =========================================== เพิ่ม data =================================================================== //
	db = database

	db.Model(&User{}).Create(&User{
		User_Name: "Watan",
		Password: "123456",
		User_Type: "แพทย์",
	})

	db.Model(&User{}).Create(&User{
		User_Name: "Thananya",
		Password: "123456789",
		User_Type: "แพทย์",
	})

	db.Model(&User{}).Create(&User{
		User_Name: "armarm",
		Password: "1234",
		User_Type: "พยาบาล",
	})

	db.Model(&User{}).Create(&User{
		User_Name: "Name",
		Password: "5678",
		User_Type: "พยาบาล",
	})

	var watan User
	var thananya User
	var armarm User
	var Name User
	db.Raw("Select * FROM users WHERE user_name = ?","Watan").Scan(&watan)
	db.Raw("Select * FROM users WHERE user_name = ?","Thananya").Scan(&thananya)
	db.Raw("Select * FROM users WHERE user_name = ?","armarm").Scan(&armarm)
	db.Raw("Select * FROM users WHERE user_name = ?","Name").Scan(&Name)

	Better := Level{
		Level_name: "Better",
	}
	db.Model(&Level{}).Create(&Better)

	Stable := Level{
		Level_name: "Stable",
	}
	db.Model(&Level{}).Create(&Stable)

	Worse := Level{
		Level_name: "Worse",
	}
	db.Model(&Level{}).Create(&Worse)

	mapBed01 := Map_Bed{
		Trigae_ID: 1,
		Admit_time: time.Now(),
		Bed_ID: 1,
		Mapb_comment: "Don't eat anythings",
	}
	db.Model(&Map_Bed{}).Create(&mapBed01)

	mapBed02 := Map_Bed{
		Trigae_ID: 2,
		Admit_time: time.Now(),
		Bed_ID: 2,
		Mapb_comment: "Don't eat anythings",
	}
	db.Model(&Map_Bed{}).Create(&mapBed02)

	mapBed03 := Map_Bed{
		Trigae_ID: 3,
		Admit_time: time.Now(),
		Bed_ID: 3,
		Mapb_comment: "NULL",
	}
	db.Model(&Map_Bed{}).Create(&mapBed03)

	db.Model(&Symptom{}).Create(&Symptom{
		Check_date: time.Now(),
		Temperature: 36,
		Pressure: 101,
		Heart_rate: 77,
		Comment: "มีผื่นใสๆ ขึ้นตามตัว",
		Mapb: mapBed01,
		Check: watan,
		Level: Better,
		Medicine: "HALOPERIDOL 5 MG.TAB",
	})
	db.Model(&Symptom{}).Create(&Symptom{
		Check_date: time.Now(),
		Temperature: 37,
		Pressure: 115,
		Heart_rate: 98,
		Comment: "ไอ มีน้ำมูก",
		Mapb: mapBed03,
		Check: thananya,
		Level: Stable,
		Medicine: "AMOXY + CLAVUL[ER][AMK]1gm.TAB",
	})

	// =========================================== เพิ่ม data =================================================================== //


	// =========================================== check list ================================================================= //

	var target User
	db.Model(&User{}).Find(&target, db.Where("user_id = ?","Watan"))

	var symp_lists []*Symptom
	db.Model(&Symptom{}).
		Joins("Map_Bed").
		Joins("Level").
		Find(&symp_lists, db.Where("check_id = ?",target.ID))

	// =========================================== check list ================================================================= //

}
