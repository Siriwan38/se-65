package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	PaymentTime time.Time `valid:"present~Payment incorrect"`
	Total       int       `valid:"Total~The value must be in range 1-9999"`

	PatientRightID *uint
	PatientRight   PatientRight `gorm:"references:ID" valid:"-"`

	PaymentTypeID *uint
	PaymentType   PaymentType `gorm:"references:ID" valid:"-"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:ID"`

	//MedicineRecordID ทำหน้าที่เป็น ForeignKey
	MedicineRecordID *uint
	MedicineRecord   MedicineRecord `gorm:"foreignKey: ID" valid:"-"`
}

type PaymentType struct {
	gorm.Model

	Type string

	Payments []Payment `gorm:"foreignKey:PaymentTypeID" valid:"-"`
}
