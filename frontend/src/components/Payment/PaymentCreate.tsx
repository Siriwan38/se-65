import React, { useEffect } from "react";
import { Link as RouterLink ,useParams} from "react-router-dom";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Divider from "@mui/material/Divider";
import FormControl from "@mui/material/FormControl";
import Grid from "@mui/material/Grid";
import MenuItem from "@mui/material/MenuItem";
import Paper from "@mui/material/Paper";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import Snackbar from "@mui/material/Snackbar";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import { PaymentsInterface, PaymentTypesInterface } from "../../models/IPayment";
import { MedicineRecordsInterface } from "../../models/IMedicineRecord";
import { EmployeesInterface } from "../../models/IEmployee";
import { PatientRightsInterface } from "../../models/IPatientRight";

import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import Button from "@mui/material/Button";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
// import moment from "moment";
import { GetMedicineRecord } from "../../services/HttpClientServiceMedicineRecord";
import { GetTreatmentRecordforMed } from "../../services/HttpClientServiceTreatmentRecord";
import { GetPatientRight } from "../../services/HttpClientServicePatientRight";
import { GetPaymentType, GetEmployee, GetPayment,GetPaymentById } from "../../services/HttpClientServicePayment";
// import { GetPayment } from "../../services/HttpClientServicePayment";
import {  MedicineOrdersInterface } from "../../models/ITreatmentRecord";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

export default function PaymentCreate() {

  const params = useParams();
  const [cashiers, setEmployee] = React.useState<EmployeesInterface>();
  const [paymenttypes, setPaymentType] = React.useState<PaymentTypesInterface[]>([]);
  const [patientrights, setPatientRight] = React.useState<PatientRightsInterface[]>([]);
  const [medicinerecords, setMedicineRecord] = React.useState<MedicineRecordsInterface[]>([]);
  const [medicineorder, setMedicineOrder] = React.useState<MedicineOrdersInterface[]>([]);

  const [selectedPatientRight, setSelectedPatientRight] = React.useState<PatientRightsInterface>();
  const [payments, setPayment] = React.useState<Partial<PaymentsInterface>>(
    { MedicineRecordID:0, PatientRightID: 0, PaymentTypeID: 0, PaymentTime: new Date(), Total: 0, });
  const [success, setSuccess] = React.useState(false);
  const [error, setError] = React.useState(false);
  const [errorMessage, setErrorMessage] = React.useState("");

  console.log(payments)

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };
  // Select Data From formcontrol
  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof payments;
    console.log("handleChage")
    console.log(event.target.value)

   
    setPayment({
      ...payments,
      [name]: event.target.value,
    });
    if (name === "MedicineRecordID") {
      let TreatmentRecordID = medicinerecords.find(m => m.ID === Number(event.target.value))?.TreatmentRecordID;
      getTreatmentRecordforMed(TreatmentRecordID!);
    }
    if (name === "PatientRightID") {
      let patientRight = patientrights.find(p => p.ID === Number(event.target.value));
      setSelectedPatientRight(patientRight!);
    }
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof payments;
    const { value } = event.target;
    setPayment({ ...payments, [id]: value });
  };

  const getTreatmentRecordforMed = async (TreatmentRecordID: number) => {
    let res = await GetTreatmentRecordforMed(TreatmentRecordID)
    if (res) {
      setMedicineOrder(res);
      // console.log(statusmed)
    }
  }
  const getPaymentType = async () => {
    let res = await GetPaymentType()
    if (res) {
      setPaymentType(res)
      console.log(res)
    }
  }
  const getPatientRight = async () => {
    let res = await GetPatientRight()
    if (res) {
      setPatientRight(res)
      console.log(res)
    }
  }
  const getMedicineRecord = async () => {
    let res = await GetMedicineRecord()
    if (res) {
      setMedicineRecord(res)
      console.log(res, "jawooo")
    }
  }


  const getEmployee = async (ID: string | null) => {
    let res = await GetEmployee()
    if (res) {
      setEmployee(res);
    }
  }
  const getPayment = async (id: string) => {
    let res = await GetPaymentById(id)
    if (res) {
      setPayment(res)
      // setMedicineOrder(res.medicinerecords.TreatmentRecord.medicineorder);
      // console.log(medicineorder)
      console.log(res)
    }
  }

  
  

  const convertType = (data: string | number | undefined | null | Date) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  function submit() {
    if (payments.MedicineRecordID == 0) {
      setError(true)
      setErrorMessage("?????????????????????????????????????????? - ?????????????????????")
    }
    else if (payments.PatientRightID == 0) {
      setError(true)
      setErrorMessage("?????????????????????????????????????????????????????????????????????")
    }
    else if (payments.PaymentTypeID == 0) {
      setError(true)
      setErrorMessage("?????????????????????????????????????????????????????????????????????????????????")
    }
    else {
      setError(false)
      let data: any = {
        MedicineRecordID: convertType(payments.MedicineRecordID),
        PatientRightID: convertType(payments.PatientRightID),
        EmployeeID: convertType(cashiers?.ID),
        PaymentTypeID: convertType(payments.PaymentTypeID),
        PaymentTime: payments.PaymentTime,
        Total: payments.Total,
      };
      console.log(data);
      
      let apiUrl : any
      if (params.id){
        data["ID"] = parseInt(params.id);
        apiUrl = "http://localhost:8080/payments"
      }
      else{
        apiUrl = "http://localhost:8080/createpayment"
      }

      // const apiUrl = `http://localhost:8080/createpayment`;
      const requestOptions = {
        method: params.id ? "PATCH" : "POST",
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      };

      fetch(apiUrl, requestOptions)
        .then((response) => response.json())
        .then((res) => {
          console.log(res);
          if (res.data) {
            setSuccess(true);
          } else {
            setError(true);
            
            if (res.error == "The data recorder should be a Cashier") {
              setErrorMessage("???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????")
            }
            else if (res.error == "MedTime must be in the present") {
              setErrorMessage("????????????????????????????????????????????????????????????????????????????????????")
            }

            else {
              setErrorMessage(res.error);
            }
            console.log(res.error);
          }
        });
    }

  }
  // const handleSelectChange = (event: SelectChangeEvent<number>) => {
  //   const name = event.target.name as keyof typeof payments;
  //   setPayment({ ...payments, [name]: event.target.value });
  // };

  console.log("Hello")
  console.log(medicinerecords)

  useEffect(() => {
    getPaymentType();
    getPatientRight();
    getMedicineRecord();
    getEmployee(localStorage.getItem("id"));

    if (params.id){
      getPayment(params.id)
    }
  }, []);
  useEffect(() => {
    setPayment (payment => {
      let MedPrice = medicineorder.reduce((a, b) => a + Number(b!.OrderAmount! * b!.Medicine!.Price!) , 0) ?? 0;
      let PatientRightPrice = selectedPatientRight?.Discount ?? 0;
      return { ...payment, Total: MedPrice - PatientRightPrice }
    });
      
  }, [selectedPatientRight, medicineorder]);

  console.log(payments);
  
  return (
    <Container sx={{ marginTop: 2 }}>
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          ??????????????????????????????????????????????????????
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          ???????????????????????????????????????????????????????????????
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 4,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 2 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ???????????????????????????????????????????????????
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>???????????? - ?????????????????????</p>
              <Select
                native
                value={payments?.MedicineRecordID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "MedicineRecordID",
                }}
              >
                <option aria-label="None" value="">
                  ?????????????????????????????????????????? - ?????????????????????
                </option>
                {medicinerecords.map((item: MedicineRecordsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.TreatmentRecord?.DiagnosisRecord?.HistorySheet?.PatientRegister?.FirstName} {item.TreatmentRecord?.DiagnosisRecord?.HistorySheet?.PatientRegister?.LastName}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>???????????????????????????????????????</p>
              <Select
                native
                value={payments.PatientRightID + ""}
                onChange={handleChange}

                inputProps={{
                  name: "PatientRightID",
                }}
              >
                <option aria-label="None" value="">
                  ?????????????????????????????????????????????????????????????????????
                </option>
                {patientrights.map((item: PatientRightsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>???????????????????????????????????????????????????</p>
              <Select
                native
                value={payments.PaymentTypeID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "PaymentTypeID",
                }}
              >
                <option aria-label="None" value="">
                  ?????????????????????????????????????????????????????????????????????????????????
                </option>
                {paymenttypes.map((item: PaymentTypesInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Type}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <p>??????????????????????????????????????????????????????</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="FirstName"
                disabled
                variant="outlined"
                type="string"
                size="medium"
                value={cashiers?.FirstName + " " + cashiers?.LastName}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={8}>
            <p>??????????????????????????????????????????</p>
            <TableContainer component={Paper} sx={{ marginTop: 4 }}>
              <Table sx={{ miinWidth: 650 }} aria-label="simple table">
                <TableHead>
                  <TableRow>
                    <TableCell align="center" width="10%">??????????????????????????????????????????</TableCell>
                    <TableCell align="center" width="10%">????????????</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {medicineorder.map((row: MedicineOrdersInterface, index) => {
                    // sumTotalPrice()
                    if (payments?.MedicineRecordID == row.ID)
                      return (
                      <TableRow key={index}>
                        <TableCell  align="center">{row.Medicine?.Name}</TableCell>
                        <TableCell align="center">{row!.Medicine!.Price! * row!.OrderAmount!}</TableCell>
                          {/* <TableCell width="5%"><IconButton size="small" onClick={() => removeFromItem(index)}><DeleteIcon /></IconButton></TableCell> */}

                        </TableRow>
                      )
                  })}

                </TableBody>
              </Table>
            </TableContainer>
          </Grid>




          <Grid item xs={4}>
            <FormControl fullWidth variant="outlined">
              <p>???????????????????????????????????????</p>
              <LocalizationProvider dateAdapter={AdapterDateFns}>
                <DatePicker
                  value={payments.PaymentTime}
                  onChange={(newValue) => {
                    setPayment({
                      ...payments,
                      PaymentTime: newValue ? newValue : new Date(),
                    });
                  }}

                  renderInput={(params) => <TextField {...params} />}
                />
              </LocalizationProvider>
            </FormControl>
          </Grid>

          <Grid item xs={3}>
            <TextField disabled id="Total" value={payments.Total} fullWidth />
          </Grid>

          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/payments"
              variant="contained"
              color="inherit"
            >
              ????????????
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              ??????????????????
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}