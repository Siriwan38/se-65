import { EmployeesInterface } from "./IEmployee";
import { HistorySheetsInterface } from "./IHistorySheet";
import { PatientRegistersInterface } from "./IPatientRegister";

export interface DiagnosisRecordsInterface {
	No?: number;

	// PatientRegisterID?: number;
	// PatientRegister?: PatientRegisterInterface;

	DoctorID?: number;
	Doctor?: EmployeesInterface;

	HistorySheetID?: number;
	HistorySheet?: HistorySheetsInterface;

	DiseaseID?: number;
	Disease?: DiseasesInterface;
	
	Examination?: string;
	Medicalcertification?: boolean;//number;
	Date?: Date | null;
}

export interface DiseasesInterface {
	ID: number;
	Name: string;
}