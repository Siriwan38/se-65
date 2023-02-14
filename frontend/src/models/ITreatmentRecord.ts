import { DiagnosisRecordsInterface } from "./IDiagnosisRecord";
import { PatientRegistersInterface } from "./IPatientRegister";
import { EmployeesInterface } from "./IEmployee";


export interface TreatmentRecordsInterface {
    ID: number;

    PatientRegisterID?: number;
	PatientRegister?: PatientRegistersInterface;

	DoctorID?: number;
	Doctor?: EmployeesInterface;

    DiagnosisRecordID?: number;
    DiagnosisRecord?:   DiagnosisRecordsInterface;

    MedicineID?: number;
    Medicine?:   MedicinesInterface;
    
    MedicineQuantity?: number;
    Treatment?: string;
    Note?: string;
    Appointment?: boolean;//int;
    Date?: Date | null;
}

export interface MedicinesInterface {
    ID: number;

    Name?: string;
    Description?: string;
    Price?: number;
}