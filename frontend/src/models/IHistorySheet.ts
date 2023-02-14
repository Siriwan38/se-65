export interface HistorySheetsInterface {
    ID: string;
    Weight: number;
    Height: number;
    BMI: number;
    Temperature: number;
    SystolicBloodPressure: number;
    DiastolicBloodPressure: number;
    HeartRate: number;
    RespiratoryRate: number;
    OxygenSaturation: number;
    DrugAllergySymtom: string;
    PatientSymtom: string;

    PatientRegisterID: number;
    PatientRegister: PatientRegistersInterface;

    NurseID: number;
    Nurse: NursesInterface;

    DrugAllergyID: number;
    DrugAllergy: DrugAllergysInterface;
}
export interface NursesInterface {
    ID: string;
    FirstName: string;
    LastName: string;
    IdentificationNumber: string;
    BirthDay: Date | null;
    Mobile: string;
    Email: string;
    Password: string;
    Salary: number;
}
export interface DrugAllergysInterface{
    ID: string;
    Name: string;
}
export interface PatientRegistersInterface {
    ID: string,
    FirstName: string;
    LastName: string;
    IdentificationNumber:   string;
    Age:    number;
    BirthDay: Date | null;
    Mobile: string;
    Email: string;
    Occupation: string;
    Address:    string;
    EmergencyPersonFirstName:           string;
	EmergencyPersonLastName :           string;
	EmergencyPersonMobile :             string;
	EmergencyPersonOccupation :         string;
	EmergencyPersonRelationWithPatient: string;

    Gender: string;
    Prefix: string;
    Nationality: string;
    Religion: string;
    BloodType: string;
    MaritalStatus: string;
    SubDistrict: string;
    District: string;
    Province: string;
}