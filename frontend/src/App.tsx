import React, { useState, useEffect } from "react";
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from "react-router-dom";
import { styled, createTheme, ThemeProvider } from "@mui/material/styles";
import MuiDrawer from "@mui/material/Drawer";
import MuiAppBar, { AppBarProps as MuiAppBarProps } from "@mui/material/AppBar";

import MedicineRecords from "./components/MedicineRecord/MedicineRecords";
import MedicineRecordCreate from "./components/MedicineRecord/MedicineRecordCreate";
import Payments from "./components/Payment/Payments";
import PaymentCreate from "./components/Payment/PaymentCreate";
import SignIn from "./components/SignIn";
import Navbar from "./components/Navbar";
import Home from "./components/Home";


  export default function App() {
    // const [token, setToken] = React.useState<string | null>();
    const token: string | null = localStorage.getItem("token"); // ตัวแปร token จะดึงค่า token ที่อยู่ใน local storage

    // useEffect(() => {
    //   setToken();
    // }, []);


    // ถ้า local storage ไม่มี token หรือ ตัวแปร token เป็น null จะสามารถเข้าถึง path Home, Log in และ UserCreate ได้
    // if (!token) {
    //   return (
    //     <Router>
    //       <div>
    //         <Navbar />
    //         <Routes>
    //           <Route path="/" element={<Home />} />
    //           <Route path="/login" element={<Login />} />
    //           <Route path="/employees" element={<Employees />} />
    //           <Route path="/createemployee" element={<EmployeeCreate />} />
    //           <Route path="/createworkload" element={<WorkloadCreate />}/>
    //           <Route path="/workloads" element={<Workloads />}/>
    //           {/* <Route path="/create" element={<UserCreate />} /> */}
    //           <Route path="*" element={<Navigate to="/" />} />
    //         </Routes>
    //       </div>
          
    //     </Router>
    //   );
    // }


    // ถ้า local storage มี token หรือ ตัวแปร token ไม่เป็น null จะสามารถเข้าถึง path Home, Users
    return (
      <Router>
        <div>
          <Navbar />
          <Routes>
          <Route path="/" element={<Home />} />
                   <Route path="/payments" element={<Payments />} />
                   <Route path="/payment/create" element={<PaymentCreate />} />
                   <Route path="/medicinerecords" element={<MedicineRecords />} />
                   <Route path="/medicinerecord/create/:id?" element={<MedicineRecordCreate />}/>
                   <Route path="" element={<Navigate to="/" />} />
          </Routes>
        </div>
      </Router>
    );
  }



