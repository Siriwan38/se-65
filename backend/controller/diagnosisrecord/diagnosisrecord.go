package controller

import (
	"net/http"

	"github.com/Siriwan38/se-65/entity"
	"github.com/gin-gonic/gin"
)

// POST	/diagnosisrecords
func CreateDiagnosisRecord(c *gin.Context) {

	// var patient	entity.PatientRegister
	var employee entity.Employee
	var historysheet entity.HistorySheet
	var disease entity.Disease
	var diagnosisrecord entity.DiagnosisRecord

	if err := c.ShouldBindJSON(&diagnosisrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ค้นหา Doctor ด้วย id
	if tx := entity.DB().Table("employees").Where("id = ?", diagnosisrecord.DoctorID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// //ค้นหา Patient ด้วย id
	// if tx := entity.DB().Table("patient_registers").Where("id = ?", diagnosisrecord.PatientRegisterID).First(&patient); tx.RowsAffected == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
	// 	return
	// }

	//ค้นหา HistorySheet ด้วย id
	if tx := entity.DB().Table("history_sheets").Where("id = ?", diagnosisrecord.HistorySheetID).First(&historysheet); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "historysheeet not found"})
		return
	}

	//ค้นหา Disease ด้วย id
	if tx := entity.DB().Table("diseases").Where("id = ?", diagnosisrecord.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}

	//สร้าง DiagnosisRecord
	dr := entity.DiagnosisRecord{
		// PatientRegister:    patient,
		Doctor:             employee,
		HistorySheet:       historysheet,
		Disease:            disease,
		Examination:        diagnosisrecord.Examination,
		MedicalCertificate: diagnosisrecord.MedicalCertificate,
		Date:               diagnosisrecord.Date,
	}

	//บันทึก
	if err := entity.DB().Create(&dr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dr})

}

// GET /diagnosisrecord/:id
func GetDiagnosisRecord(c *gin.Context) {
	var diagnosisrecord entity.DiagnosisRecord
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM diagnosis_records WHERE id = ?", id).
		// Preload("PatientRegister").
		// Preload("Employee").
		Preload("HistorySheet").
		Preload("Disease").
		Find(&diagnosisrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diagnosisrecord})
}

// GET /diagnosisrecords
func ListDiagnosisRecords(c *gin.Context) {
	var diagnosisrecords []entity.DiagnosisRecord
	if err := entity.DB().Raw("SELECT * FROM diagnosis_records").
		// Preload("PatientRegister").
		// Preload("Doctor").
		Preload("HistorySheet").
		Preload("HistorySheet.PatientRegister").
		Preload("Disease").
		Find(&diagnosisrecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diagnosisrecords})
}

// DELETE /diagnosisrecords/:id
func DeleteDiagnosisRecord(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM diagnosis_records WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diagnosisrecord not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

//PATCH /diagnosisrecords
func UpdateDiagnosisRecord(c *gin.Context) {
	var diagnosisrecord entity.DiagnosisRecord
	var disease entity.Disease

	if err := c.ShouldBindJSON(&diagnosisrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", diagnosisrecord.ID).First(&diagnosisrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "diagnosisrecord not found"})
		return
	}

	diagnosisrecord.Disease = disease

	if err := entity.DB().Save(&diagnosisrecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diagnosisrecord})
}
