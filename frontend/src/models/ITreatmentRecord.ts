import { DiagnosisRecordsInterface } from "./IDiagnosisRecord";
import { PatientRegistersInterface } from "./IPatientRegister";
import { EmployeesInterface } from "./IEmployee";


export interface TreatmentRecordsInterface {
    ID: number;

	DoctorID?: number;
	Doctor?: EmployeesInterface;

    DiagnosisRecordID?: number;
    DiagnosisRecord?:   DiagnosisRecordsInterface;

    MedicineID?: number;
    Medicine?:   MedicinesInterface;

    MedicineOrders?: MedicineOrdersInterface[];
    
   
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

export interface MedicineOrdersInterface {
    ID: number;
    OrderAmount?: number;
    MedicineID?: number;
    Medicine?:   MedicinesInterface;
    TreatmentRecordID?: number;
    TreatmentRecord?: TreatmentRecordsInterface;

}