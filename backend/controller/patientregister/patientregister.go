package controller

import (
	"net/http"

	"github.com/Siriwan38/se-65/entity"
	"github.com/gin-gonic/gin"
)

// POST /users
// 	func CreateUser(c *gin.Context) {
// 		var user entity.User
// 		if err := c.ShouldBindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if err := entity.DB().Create(&user).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"data": user})
// 	}

// POST /patientregisters
func CreatePatientRegister(c *gin.Context) {
	var patientregister entity.PatientRegister
	var employee entity.Employee
	var prefix entity.Prefix
	var gender entity.Gender
	var nationality entity.Nationality
	var religion entity.Religion
	var bloodtype entity.BloodType
	var maritalstatus entity.MaritalStatus
	var province entity.Province
	var district entity.District
	var subdistrict entity.SubDistrict

	if err := c.ShouldBindJSON(&patientregister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Table("employees").Where("id = ?", patientregister.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Table("prefixes").Where("id = ?", patientregister.PrefixID).First(&prefix); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prefix not found"})
		return
	}
	if tx := entity.DB().Table("genders").Where("id = ?", patientregister.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}
	if tx := entity.DB().Table("nationalities").Where("id = ?", patientregister.NationalityID).First(&nationality); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nationality not found"})
		return
	}
	if tx := entity.DB().Table("religions").Where("id = ?", patientregister.ReligionID).First(&religion); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "religion not found"})
		return
	}
	if tx := entity.DB().Table("blood_types").Where("id = ?", patientregister.BloodTypeID).First(&bloodtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bloodtype not found"})
		return
	}
	if tx := entity.DB().Table("marital_statuses").Where("id = ?", patientregister.MaritalStatusID).First(&maritalstatus); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "maritalstatus not found"})
		return
	}
	if tx := entity.DB().Table("provinces").Where("id = ?", patientregister.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}
	if tx := entity.DB().Table("districts").Where("id = ?", patientregister.DistrictID).First(&district); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "district not found"})
		return
	}
	if tx := entity.DB().Table("sub_districts").Where("id = ?", patientregister.SubDistrictID).First(&subdistrict); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subdistrict not found"})
		return
	}
	p := entity.PatientRegister{
		FirstName:                              patientregister.FirstName,
		LastName:                               patientregister.LastName,
		IdentificationNumber:                   patientregister.IdentificationNumber,
		Age:                                    patientregister.Age,
		BirthDay:                               patientregister.BirthDay,
		Mobile:                                 patientregister.Mobile,
		Email:                                  patientregister.Email,
		Occupation:                             patientregister.Occupation,
		Address:                                patientregister.Address,
		EmergencyPersonFirstName:               patientregister.EmergencyPersonFirstName,
		EmergencyPersonLastName:                patientregister.EmergencyPersonLastName,
		EmergencyPersonMobile:                  patientregister.EmergencyPersonMobile,
		EmergencyPersonOccupation:              patientregister.EmergencyPersonOccupation,
		EmergencyPersonRelationshipWithPatient: patientregister.EmergencyPersonRelationshipWithPatient,
		Employee:                               employee,
		Prefix:                                 prefix,
		Gender:                                 gender,
		Nationality:                            nationality,
		Religion:                               religion,
		BloodType:                              bloodtype,
		MaritalStatus:                          maritalstatus,
		SubDistrict:                            subdistrict,
		District:                               district,
		Province:                               province,
	}
	if err := entity.DB().Create(&p).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": p})
}

// *******************************************************************************************************

// GET /user/:id
// func GetUser(c *gin.Context) {
// 	var user entity.User
// 	id := c.Param("id")
// 	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user).Error; err != nil {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		   return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

// GET /patientregister/:id
func GetPatientRegister(c *gin.Context) {
	var patientregister entity.PatientRegister
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patient_registers WHERE id = ?", id).
		Preload("Employee").
		Preload("Gender").
		Preload("Prefix").
		Preload("Nationality").
		Preload("Religion").
		Preload("BloodType").
		Preload("MaritalStatus").
		Preload("Province").
		Preload("District").
		Preload("SubDistrict").
		Find(&patientregister).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patientregister})
}

// *******************************************************************************************************

// GET /users
// func ListUsers(c *gin.Context) {
// 	var users []entity.User
// 	if err := entity.DB().Raw("SELECT * FROM users").Scan(&users).Error; err != nil {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		   return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": users})
// }

// GET /patientregisters
func ListPatientRegisters(c *gin.Context) {
	var patientregisters []entity.PatientRegister
	if err := entity.DB().Raw("SELECT * FROM patient_registers").
		Preload("Employee").
		Preload("Gender").
		Preload("Prefix").
		Preload("Nationality").
		Preload("Religion").
		Preload("BloodType").
		Preload("MaritalStatus").
		Preload("Province").
		Preload("District").
		Preload("SubDistrict").
		Find(&patientregisters).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patientregisters})
}

// *******************************************************************************************************

// DELETE /users/:id
// func DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		   return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }

// DELETE /patientregisters/:id
func DeletePatientRegister(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM patient_registers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patientregister not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// *******************************************************************************************************

// PATCH /users
// func UpdateUser(c *gin.Context) {
// 	var user entity.User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		   return
// 	}
// 	if tx := entity.DB().Where("id = ?", user.ID).First(&user); tx.RowsAffected == 0 {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		   return
// 	}
// 	if err := entity.DB().Save(&user).Error; err != nil {
// 		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		   return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

// PATCH /patientregisters
func UpdatePatientRegister(c *gin.Context) {
	var patientregister entity.PatientRegister
	if err := c.ShouldBindJSON(&patientregister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", patientregister.ID).First(&patientregister); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patientregister not found"})
		return
	}
	if err := entity.DB().Save(&patientregister).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patientregister})
}
