package entity

import (
	"gorm.io/gorm"
)

type PatientType struct {
	gorm.Model
	Typename      string
	PatientRights []PatientRight `gorm:"foreignKey:PatientTypeID"`
}

type PatientRight struct {
	gorm.Model

	Name string

	Discount          uint
	PatientRegisterID *uint
	PatientRegister   PatientRegister `gorm:"references:ID"`

	// Createlimit   int16
	PatientTypeID *uint
	PatientType   PatientType

	//Bills		[]Bill		`gorm:"foreignKey:PatientRightID"`
}

// func init() {
// 	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
// 		t := i.(time.Time)
// 		return t.Before(time.Now())
// 	})

// 	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
// 		t := i.(time.Time)
// 		return t.After(time.Now())
// 	})

// 	govalidator.CustomTypeTagMap.Set("Now", func(i interface{}, context interface{}) bool {
// 		t := i.(time.Time)
// 		return t.Equal(time.Now())
// 	})

// 	govalidator.CustomTypeTagMap.Set("DelayNow3Min", func(i interface{}, context interface{}) bool {
// 		t := i.(time.Time)
// 		return t.After(time.Now().Add(3 - time.Minute))
// 	})

// 	govalidator.CustomTypeTagMap.Set("positive", func(i interface{}, context interface{}) bool {
// 		num := i
// 		return num.(int) >= 0
// 	})
// 	govalidator.CustomTypeTagMap.Set("positivenum", func(i interface{}, context interface{}) bool {
// 		num := i
// 		return num.(int) > 0
// 	})
// }
