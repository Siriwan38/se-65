package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("se-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&Employee{},
		&DrugAllergy{},
		&TreatmentRecord{},
		&Medicine{},
		&HistorySheet{},
		&PatientRight{},
		&PatientRegister{},
		&DiagnosisRecord{},
		&MedicineRecord{},
		&Payment{},
		&PatientType{},
		&PaymentType{},
		&StatusMed{},
		&MedicineOrder{})

	db = database

	// //Role Data
	// db.Model(&Role{}).Create(&nurse)
	// pharmacist := Role{
	// 	Position: "Pharmacist",
	// }

	//Role Data

	pharmacist := Role{
		Name: "Pharmacist",
	}
	db.Model(&Role{}).Create(&pharmacist)
	cashier := Role{
		Name: "Cashier",
	}
	db.Model(&Role{}).Create(&cashier)

	// Employee Data
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	Employee1 := Employee{
		IDCard:      "1410087223152",
		FirstName:   "Siriwan",
		LastName:    "Pratan",
		PhoneNumber: "0999999909",
		Email:       "Siri@gmail.com",
		Password:    string(password),
		Salary:      21000,
		Birthday:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Role:        cashier,
	}
	db.Model(&Employee{}).Create(&Employee1)

	Employee2 := Employee{
		IDCard:      "1410087220152",
		FirstName:   "Park ",
		LastName:    "Chaeyong",
		PhoneNumber: "0829123212",
		Email:       "Park@gmail.com",
		Password:    string(password),
		Salary:      21000,
		Birthday:    time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC),
		Role:        pharmacist,
	}
	db.Model(&Employee{}).Create(&Employee2)
	drugallergy1 := DrugAllergy{
		Name: "แพ้ยา/Drug Allergy",
	}
	db.Model(&DrugAllergy{}).Create(&drugallergy1)

	drugallergy2 := DrugAllergy{
		Name: "ไม่แพ้ยา/Not Allergic",
	}
	db.Model(&DrugAllergy{}).Create(&drugallergy2)
	patientregister1 := PatientRegister{
		FirstName:            "A",
		LastName:             "B",
		IdentificationNumber: "1436287954230",
		Age:                  5,
		// BirthDay:                               dob5,
		Mobile:                                 "0824298341",
		Email:                                  "AB@g.com",
		Occupation:                             "Lawyer",
		Address:                                "SUT",
		EmergencyPersonFirstName:               "G",
		EmergencyPersonLastName:                "H",
		EmergencyPersonMobile:                  "0657954135",
		EmergencyPersonOccupation:              "Police",
		EmergencyPersonRelationshipWithPatient: "Older brother",
		// Prefix:                                 "Mister",
		// Gender:                                 "Male",
		// Nationality:                            "Qutar",
		// Religion:                               "Islam",
		// BloodType:                              "A",
		// MaritalStatus:                          "Single",
		// SubDistrict:                            "Rattanawapi",
		// District:                               "Rattanawapi",
		// Province:                               "Nong Khai",
	}
	db.Model(&PatientRegister{}).Create(&patientregister1)
	//HistorySheet data
	historysheet1 := HistorySheet{
		Weight:                 50.00,
		Height:                 165.00,
		Temperature:            35.40,
		SystolicBloodPressure:  110,
		DiastolicBloodPressure: 85,
		HeartRate:              96,
		RespiratoryRate:        5,
		OxygenSaturation:       89,
		// DrugAllergySymtom:      79,
		// PatientSymtom:          normall,
		PatientRegister: patientregister1,
		// Nurse:                  Employee2,

		// DrugAllergyID *uint
		// DrugAllergy   DrugAllergy `gorm:"references:ID"`

	}
	db.Model(&HistorySheet{}).Create(&historysheet1)

	diagnosis1 := DiagnosisRecord{
		PatientRegister: historysheet1.PatientRegister,
		Doctor:          Employee2,
		HistorySheet:    historysheet1,
		// Disease:         disease10,
		Examination:        "ผู้ป่วยมีอาการ",
		MedicalCertificate: true,
		Date:               time.Now(),
	}
	db.Model(&DiagnosisRecord{}).Create(&diagnosis1)

	statusmed1 := StatusMed{
		Number: 1,
		Status: "การจ่ายยาครั้งที่1",
	}
	db.Model(&StatusMed{}).Create(&statusmed1)
	statusmed2 := StatusMed{
		Number: 2,
		Status: "การจ่ายยาครั้งที่2",
	}
	db.Model(&StatusMed{}).Create(&statusmed2)

	medicine := Medicine{
		Name:        "Paracetamal",
		Description: "รับประทานลังอาหาร",
		Price:       100,
	}
	db.Model(&Medicine{}).Create(&medicine)

	medicine1 := Medicine{
		Name:        "Vitamin C",
		Description: "วันละไม่เกิน 1000 mg.",
		Price:       200,
	}
	db.Model(&Medicine{}).Create(&medicine1)

	payment1 := PaymentType{
		Type: "เงินสด",
	}
	db.Model(&PaymentType{}).Create(&payment1)

	payment2 := PaymentType{
		Type: "Mobile Banking",
	}
	db.Model(&PaymentType{}).Create(&payment2)
	patienttype1 := PatientType{
		Typename: "สุขภาพถ้วนหน้า",
	}
	db.Model(&PatientType{}).Create(&patienttype1)
	patientright1 := PatientRight{
		Name: "ประกันสุขภาพ",

		Discount:        80,
		PatientRegister: patientregister1,

		PatientType: patienttype1,
	}
	db.Model(&PatientRight{}).Create(&patientright1)

	t := true
	order := []MedicineOrder{

		{Medicine: medicine1, OrderAmount: 6},
	}

	treatmentrecord1 := TreatmentRecord{
		// PatientRegister:     diagnosis1.PatientRegister,
		Doctor:          Employee1,
		DiagnosisRecord: diagnosis1,
		Treatment:       "ให้ทานยา และพักผ่อน",
		Note:            "รอดูอาการในการตรวจครั้งหน้า",
		Appointment:     &t,
		MedicineOrders:  order,
		Date:            time.Now(),
	}
	db.Model(&TreatmentRecord{}).Create(&treatmentrecord1)
	medicinerecord1 := MedicineRecord{
		Employee:        Employee2,
		TreatmentRecord: treatmentrecord1,

		StatusMed:  statusmed1,
		Advicetext: "none",

		// MedicinePrice: 100,
		MedTime: time.Now(),
	}
	db.Model(&MedicineRecord{}).Create(&medicinerecord1)

}
