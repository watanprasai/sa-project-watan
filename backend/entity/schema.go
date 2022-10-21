package entity

import(
	"gorm.io/gorm"
	"time"
)

type User struct{
	gorm.Model				
	User_Name				string			`gorm:"uniqueIndex"`
	Password				string
	User_Type				string
	// 1 User สามารถตรวจได้หลายคน
	Symptoms				[]Symptom		`gorm:"foreignKey:CheckID"`
}

type Level struct{
	gorm.Model
	Level_name				string
	Symptoms				[]Symptom		`gorm:"foreignKey:LevelID"`
}

type Map_Bed struct{
	gorm.Model
	Trigae_ID				uint
	Admit_time				time.Time
	Bed_ID					uint
	Mapb_comment			string
	Symptoms				[]Symptom		`gorm:"foreignKey:MapbID"`
}

type Symptom struct{
	gorm.Model
	Check_date				time.Time
	Temperature				uint
	Pressure				uint
	Heart_rate				uint
	Comment					string
	Medicine				string
	
	// CheckID ทำหน้าที่เป็น FK
	CheckID					*uint
	// เป็นข้อมูล User เมื่อ join ตาราง
	Check					User

	LevelID					*uint
	Level					Level
	
	MapbID					*uint
	Mapb					Map_Bed
}

