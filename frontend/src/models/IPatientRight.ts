import { EmployeesInterface } from "./IEmployee";
import  {PatientRegistersInterface} from "./IPatientRegister";

export interface PatientTypesInterface {
    ID: number,
    Typename: string,
}

export interface PatientRightsInterface {
    ID: number,
    Name: string,
    PatientRegisterID: number,
    PatientRegister: PatientRegistersInterface,
    PatientTypeID: number,
    PatientType: PatientTypesInterface,
    EmployeeID: number,
    Employee: EmployeesInterface,
    Discount: number,
    // Role: RoleInterface,
}