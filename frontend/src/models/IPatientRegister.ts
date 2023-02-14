import { EmployeesInterface } from "./IEmployee";
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

    EmployeeID: number;
    Employee: EmployeesInterface;

    GenderID: number;
    Gender: GendersInterface;

    PrefixID: number;
    Prefix: PrefixsInterface;

    NationalityID: number;
    Nationality: NationalitysInterface;

    ReligionID: number;
    Religion: ReligionsInterface;

    BloodTypeID: number;
    BloodType: BloodTypesInterface;

    MaritalStatusID: number;
    MaritalStatus: MaritalStatussInterface;

    SubDistrictID: number;
    SubDistrict: SubDistrictsInterface;

    DistrictID: number;
    District: DistrictsInterface;

    ProvinceID: number;
    Province: ProvincesInterface;
}
// export interface EmployeesInterface {
//     ID: string;
//     FirstName: string;
//     LastName: string;
//     IdentificationNumber: string;
// 	BirthDay:             Date | null;
// 	Mobile:               string;
// 	Email:                string;
// 	Password:             string;
// 	Salary:               number;
// }
export interface GendersInterface {
    ID: string;
    Name: string;
}
export interface PrefixsInterface {
    ID: string;
    Name: string;
}
export interface NationalitysInterface {
    ID: string;
    Name: string;
}
export interface ReligionsInterface {
    ID: string;
    Name: string;
}
export interface BloodTypesInterface {
    ID: string;
    Name: string;
}
export interface MaritalStatussInterface {
    ID: string;
    Name: string;
}
export interface SubDistrictsInterface {
    ID: string;
    Name: string;
    PostCode: string;
    District: string;
}
export interface DistrictsInterface {
    ID: string;
    Name: string;
    Province: string;
}
export interface ProvincesInterface {
    ID: string;
    Name: string;
}