package main

import (
	controller "github.com/Siriwan38/se-65/controller"
	controller_diagnosisrecord "github.com/Siriwan38/se-65/controller/diagnosisrecord"
	controller_employee "github.com/Siriwan38/se-65/controller/employee"
	controller_historysheet "github.com/Siriwan38/se-65/controller/historysheet"
	controller_medicinerecord "github.com/Siriwan38/se-65/controller/medicinerecord"
	controller_patientregister "github.com/Siriwan38/se-65/controller/patientregister"
	controller_patientright "github.com/Siriwan38/se-65/controller/patientright"
	controller_payment "github.com/Siriwan38/se-65/controller/payment"
	controller_treatmentrecord "github.com/Siriwan38/se-65/controller/treatmentrecord"

	"github.com/Siriwan38/se-65/middlewares"

	"github.com/Siriwan38/se-65/entity"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	// r.GET("/treatmentrecords", controller_treatmentrecord.ListTreatmentRecords)
	// r.GET("/treatmentrecord/:id", controller_treatmentrecord.GetTreatmentRecord)
	// r.POST("/medicinerecords", controller_medicinerecord.CreateMedicineRecord)
	// r.GET("/medicinerecords", controller_medicinerecord.ListMedicineRecords)
	// r.GET("/statusmeds", controller_medicinerecord.ListStatusMeds)
	// r.GET("/statusmed/:id", controller_medicinerecord.GetStatusMed)
	// r.GET("/employees", controller.ListEmployee)
	// r.GET("/employee/:id", controller.GetEmployee)
	// User Routes
	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// PatientRegister Routes
			protected.GET("/patientregisters", controller_patientregister.ListPatientRegisters)
			protected.GET("/patientregister/:id", controller_patientregister.GetPatientRegister)
			protected.POST("/createpatientregister", controller_patientregister.CreatePatientRegister)
			protected.PATCH("/patientregisters", controller_patientregister.UpdatePatientRegister)
			protected.DELETE("/patientregisters/:id", controller_patientregister.DeletePatientRegister)

			// Gender Routes
			protected.GET("/genders", controller_patientregister.ListGenders)
			protected.GET("/gender/:id", controller_patientregister.GetGender)
			protected.POST("/creategender", controller_patientregister.CreateGender)
			protected.PATCH("/genders", controller_patientregister.UpdateGender)
			protected.DELETE("/genders/:id", controller_patientregister.DeleteGender)

			// Prefix Routes
			protected.GET("/prefixes", controller_patientregister.ListPrefixes)
			protected.GET("/prefix/:id", controller_patientregister.GetPrefix)
			protected.POST("/createprefixe", controller_patientregister.CreatePrefix)
			protected.PATCH("/prefixes", controller_patientregister.UpdatePrefix)
			protected.DELETE("/prefixes/:id", controller_patientregister.DeletePrefix)

			// Nationality Routes
			protected.GET("/nationalities", controller_patientregister.ListNationalities)
			protected.GET("/nationality/:id", controller_patientregister.GetNationality)
			protected.POST("/createnationalitie", controller_patientregister.CreateNationality)
			protected.PATCH("/nationalities", controller_patientregister.UpdateNationality)
			protected.DELETE("/nationalities/:id", controller_patientregister.DeleteNationality)

			// Religion Routes
			protected.GET("/religions", controller_patientregister.ListReligions)
			protected.GET("/religion/:id", controller_patientregister.GetReligion)
			protected.POST("/createreligion", controller_patientregister.CreateReligion)
			protected.PATCH("/religions", controller_patientregister.UpdateReligion)
			protected.DELETE("/religions/:id", controller_patientregister.DeleteReligion)

			// BloodType Routes
			protected.GET("/bloodtypes", controller_patientregister.ListBloodTypes)
			protected.GET("/bloodtype/:id", controller_patientregister.GetBloodType)
			protected.POST("/createbloodtype", controller_patientregister.CreateBloodType)
			protected.PATCH("/bloodtypes", controller_patientregister.UpdateBloodType)
			protected.DELETE("/bloodtypes/:id", controller_patientregister.DeleteBloodType)

			// MaritalStatus Routes
			protected.GET("/maritalstatuses", controller_patientregister.ListMaritalStatuses)
			protected.GET("/maritalstatus/:id", controller_patientregister.GetMaritalStatus)
			protected.POST("/createmaritalstatuse", controller_patientregister.CreateMaritalStatus)
			protected.PATCH("/maritalststuses", controller_patientregister.UpdateMaritalStatus)
			protected.DELETE("/maritalstatuses/:id", controller_patientregister.DeleteMaritalStautus)

			// SubDistrict Routes
			protected.GET("/subdistricts", controller_patientregister.ListSubDistricts)
			protected.GET("/subdistrict/:id", controller_patientregister.GetSubDistrict)
			protected.POST("/createsubdistrict", controller_patientregister.CreateSubDistrict)
			protected.PATCH("/subdistricts", controller_patientregister.UpdateSubDistrict)
			protected.DELETE("/subdistricts/:id", controller_patientregister.DeleteSubDistrict)

			protected.GET("/subdistricts/districts/:id", controller_patientregister.ListSubDistrictsByDistricts)

			// District Routes
			protected.GET("/districts", controller_patientregister.ListDistricts)
			protected.GET("/district/:id", controller_patientregister.GetDistrict)
			protected.POST("/createdistrict", controller_patientregister.CreateDistrict)
			protected.PATCH("/districts", controller_patientregister.UpdateDistrict)
			protected.DELETE("/districts/:id", controller_patientregister.DeleteDistrict)

			protected.GET("/districts/provinces/:id", controller_patientregister.ListDistrictsByProvinces)

			//============================================================================
			// Province Routes
			protected.GET("/provinces", controller_patientregister.ListProvinces)
			protected.GET("/province/:id", controller_patientregister.GetProvince)
			protected.POST("/createprovince", controller_patientregister.CreateProvince)
			protected.PATCH("/provinces", controller_patientregister.UpdateProvince)
			protected.DELETE("/provinces/:id", controller_patientregister.DeleteProvince)

			// Ugency Routes
			protected.GET("/drugallergies", controller_historysheet.ListDrugAllergies)
			protected.GET("/urgency/:id", controller_historysheet.GetDrugAllergy)
			protected.POST("/createdrugallergie", controller_historysheet.CreateDrugAllergy)
			protected.PATCH("/drugallergies", controller_historysheet.UpdateDrugAllergy)
			protected.DELETE("/drugallergies/:id", controller_historysheet.DeleteDrugAllergy)

			// Nurse Routes
			protected.GET("/nurses", controller_historysheet.ListNurses)
			protected.GET("/nurse/:id", controller_historysheet.GetNurse)
			protected.POST("/createnurse", controller_historysheet.CreateNurse)
			protected.PATCH("/nurses", controller_historysheet.UpdateNurse)
			protected.DELETE("/nurses/:id", controller_historysheet.DeleteNurse)

			// HistorySheet Routes
			protected.GET("/historysheets", controller_historysheet.ListHistorySheets)
			protected.GET("/historysheet/:id", controller_historysheet.GetHistorySheet)
			protected.POST("/createhistorysheet", controller_historysheet.CreateHistorySheet)
			protected.PATCH("/historysheets", controller_historysheet.UpdateHistorySheet)
			protected.DELETE("/historysheets/:id", controller_historysheet.DeleteHistorySheet)

			//================================================================
			//================================================================
			protected.GET("/patientrights", controller_patientright.ListPatientRights)
			protected.GET("/patientright/:id", controller_patientright.GetPatientRights)
			protected.POST("/createpatientright", controller_patientright.CreatePatientRight)
			protected.PATCH("/patientrights", controller_patientright.UpdatePatientRight)
			protected.DELETE("/patientrights/:id", controller_patientright.DeletePatientRight)
			//================================================================
			protected.GET("/patienttypes", controller_patientright.ListPatientTypes)
			protected.GET("/patienttypes/:id", controller_patientright.GetPatientType)

			//================================================================

			// DiagonsisRecord Routes
			protected.GET("/diagnosisrecords", controller_diagnosisrecord.ListDiagnosisRecords)
			protected.GET("/diagnosisrecord/:id", controller_diagnosisrecord.GetDiagnosisRecord)
			protected.POST("/creatediagnosisrecord", controller_diagnosisrecord.CreateDiagnosisRecord)
			protected.PATCH("/diagnosisrecords", controller_diagnosisrecord.UpdateDiagnosisRecord)
			protected.DELETE("/diagnosisrecords/:id", controller_diagnosisrecord.DeleteDiagnosisRecord)

			// Disease Routes
			protected.GET("/diseases", controller_diagnosisrecord.ListDiagnosisRecords)
			protected.GET("/disease/:id", controller_diagnosisrecord.GetDiagnosisRecord)
			protected.POST("/createdisease", controller_diagnosisrecord.CreateDiagnosisRecord)
			protected.PATCH("/diseases", controller_diagnosisrecord.UpdateDiagnosisRecord)
			protected.DELETE("/diseases/:id", controller_diagnosisrecord.DeleteDiagnosisRecord)

			//================================================================

			// Medicine Routes
			protected.GET("/medicines", controller_diagnosisrecord.ListMedicines)
			protected.GET("/medicine/:id", controller_diagnosisrecord.GetMedicine)
			protected.POST("/createmedicine", controller_diagnosisrecord.CreateMedicine)
			protected.PATCH("/medicines", controller_diagnosisrecord.UpdateMedicine)
			protected.DELETE("/medicines/:id", controller_diagnosisrecord.DeleteMedicine)

			// TreatmentRecord Routes
			protected.POST("/createmtreatmentrecord", controller_treatmentrecord.CreateTreatmentRecord)
			protected.GET("/treatmentrecords", controller_treatmentrecord.ListTreatmentRecords)
			protected.GET("/treatmentrecord/:id", controller_treatmentrecord.GetTreatmentRecord)
			protected.PATCH("/treatmentrecords", controller_treatmentrecord.UpdateTreatmentRecord)
			protected.DELETE("/treatmentrecords/:id", controller_treatmentrecord.DeleteTreatmentRecord)

			//================================================================

			// MedicineRecord Routes
			protected.POST("/createmedicinerecord", controller_medicinerecord.CreateMedicineRecord)
			protected.GET("/medicinerecords", controller_medicinerecord.ListMedicineRecords)
			protected.GET("/medicinerecord/:id", controller_medicinerecord.GetMedicineRecord)
			protected.PATCH("/medicinerecord", controller_medicinerecord.UpdateMedicineRecord)
			protected.DELETE("/medicinerecords/:id", controller_medicinerecord.DeleteMedicineRecord)

			//StatusMed Routes
			protected.GET("/statusmeds", controller_medicinerecord.ListStatusMeds)
			protected.GET("/statusmed/:id", controller_medicinerecord.GetStatusMed)

			//================================================================
			//Employee Routes
			// Employee Routes
			protected.GET("/employees", controller_employee.ListEmployee)
			protected.GET("/employee/:id", controller_employee.GetEmployee)
			protected.GET("/employeerole/:roleid", controller_employee.GetEmployeerole)
			protected.POST("/createemployee", controller_employee.CreateEmployee)
			protected.PATCH("/employees", controller_employee.UpdateEmployee)
			protected.DELETE("/employees/:id", controller_employee.DeleteEmployee)
			// Admin Routes
			protected.GET("/admin/:id", controller_employee.GetAdmin) //ดึง Admin ด้วย id
			protected.GET("/admins", controller_employee.ListAdmin)   //ดึง Admin
			//Title
			protected.GET("/title/:id", controller_employee.GetTitle)
			protected.GET("/titles", controller_employee.ListTitle) //ดึง Title
			//Role
			protected.GET("/role/:id", controller_employee.GetRole)
			protected.GET("/roles", controller_employee.ListRole) //ดึง Role

			//================================================================

			// //Paymenttype Routes
			protected.GET("/paymenttypes", controller_payment.ListPaymentTypes)
			protected.GET("/paymenttype/:id", controller_payment.GetPaymentType)

			//Payment Routes
			protected.POST("/createpayment", controller_payment.CreatePayment)
			protected.GET("/payments", controller_payment.ListPayments)
			protected.GET("/payment/:id", controller_payment.GetPayment)
			protected.PATCH("/payments", controller_payment.UpdatePayment)
			protected.DELETE("/payments/:id", controller_payment.DeletePayment)

			protected.DELETE("/medbypatien/:id", controller_payment.GetMedbyPatient)

		}
	}
	// User Routes
	// r.POST("/employees", controller.CreateEmployee)

	// Authentication Routes
	r.POST("/login", controller.Login)
	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}

}

// func main() {

// 	entity.SetupDatabase()

// 	r := gin.Default()

// 	// User Routes

// 	r.GET("/users", controller.ListUsers)

// 	r.GET("/user/:id", controller.GetUser)

// 	r.POST("/users", controller.CreateUser)

// 	r.PATCH("/users", controller.UpdateUser)

// 	r.DELETE("/users/:id", controller.DeleteUser)

// 	// Run the server

// 	r.Run()

// }
