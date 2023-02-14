package entity

import (
	"time"

	"gorm.io/gorm"
)

type HistorySheet struct {
	gorm.Model
	Weight                 float32
	Height                 float32
	Temperature            float32
	SystolicBloodPressure  uint16
	DiastolicBloodPressure uint16
	HeartRate              uint8
	RespiratoryRate        uint8
	OxygenSaturation       uint8
	DrugAllergySymtom      string
	PatientSymtom          string

	PatientRegisterID *uint
	PatientRegister   PatientRegister `gorm:"references:ID"`

	NurseID *uint
	Nurse   Nurse `gorm:"references:ID"`

	DrugAllergyID *uint
	DrugAllergy   DrugAllergy `gorm:"references:ID"`

	DiagnosisRecords []DiagnosisRecord `gorm:"foreignKey:HistorySheetID"`
}

type Nurse struct {
	gorm.Model
	Name                 string
	IdentificationNumber string
	BirthDay             time.Time
	Mobile               string
	Email                string
	Password             string
	Salary               uint16
	HistorySheets        []HistorySheet `gorm:"foreignKey:NurseID"`
}

type DrugAllergy struct {
	gorm.Model
	Name          string
	HistorySheets []HistorySheet `gorm:"foreignKey:DrugAllergyID"`
}
