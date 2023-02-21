package entity

import (
	"time"

	"gorm.io/gorm"
)

type TreatmentRecord struct {
	gorm.Model

	//DoctorID เป็น FK
	DoctorID *uint
	Doctor   Employee `gorm:"references:id" valid:"-"`

	//DiagnosisRecord เป็น FK
	DiagnosisRecordID *uint
	DiagnosisRecord   DiagnosisRecord `gorm:"references:id" valid:"-"`

	Treatment   string `valid:"required~Treatment cannot be blank"`
	Note        string `valid:"-"`
	Appointment *bool

	MedicineOrders []MedicineOrder `gorm:"foreignKey:TreatmentRecordID;  constraint:OnDelete:CASCADE"`

	Date time.Time `valid:"present~Date must be present"`
}

type Medicine struct {
	gorm.Model

	Name        string
	Description string
	Price       float32

	MedicineOrder []MedicineOrder `gorm:"foreignKey:MedicineID"`
}

type MedicineOrder struct {
	gorm.Model

	OrderAmount int `valid:"int, range(0|100)~Order Amount must not be negative"`

	MedicineID *uint
	Medicine   Medicine `gorm:"refernces:ID" valid:"-"`

	TreatmentRecordID *uint
	TreatmentRecord   TreatmentRecord `gorm:"references:ID" valid:"-"`
}
