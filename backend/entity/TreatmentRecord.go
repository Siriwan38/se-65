package entity

import (
	"time"

	"gorm.io/gorm"
)

type TreatmentRecord struct {
	gorm.Model

	// //PatientID เป็น FK
	PatientRegisterID *uint
	PatientRegister   PatientRegister `gorm:"references:id" valid:"-"`

	//DoctorID เป็น FK
	DoctorID *uint
	Doctor   Employee `gorm:"references:id" valid:"-"`

	//DiagnosisRecord เป็น FK
	DiagnosisRecordID *uint
	DiagnosisRecord   DiagnosisRecord `gorm:"references:id" valid:"-"`

	Treatment   string
	Note        string
	Appointment *bool

	//MedicineID เป็น FK
	MedicineID       *uint
	Medicine         Medicine `gorm:"references:id" valid:"-"`
	MedicineQuantity int

	Date time.Time
	// PatientRegisters []PatientRegister `gorm:"foreignKey:TreatmentRecordID"`
}

type Medicine struct {
	gorm.Model

	Name        string
	Description string
	Price       int

	TreatmentRecords []TreatmentRecord `gorm:"foreignKey:MedicineID"`
}

// type MedicineItem struct {
// 	gorm.Model

// 	//MedicineID เป็น FK
// 	MedicineID *uint    `gorm:"uniquelndex"`
// 	Medicine   Medicine `gorm:"references:id" valid:"-"`

// 	TreatmentRecordID *uint
// 	TreatmentRecord   TreatmentRecord `gorm:"references:id" valid:"-"`

// }
