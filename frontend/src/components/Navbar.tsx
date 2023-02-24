import React, { useEffect, useState} from 'react';
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import Tooltip from '@mui/material/Tooltip';
import MenuIcon from "@mui/icons-material/Menu";
import { Chip } from '@mui/material';
import AssignmentIcon from "@mui/icons-material/Assignment";
import LogoutIcon from '@mui/icons-material/Logout';
import { Link as RouterLink } from "react-router-dom";
import Drawer from "@mui/material/Drawer";
import { EmployeesInterface } from '../models/IEmployee';
import List from "@mui/material/List"
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import HomeIcon from "@mui/icons-material/Home";
import CreateIcon from '@mui/icons-material/Create';
import SignIn from './SignIn';
import FaceIcon from '@mui/icons-material/Face';
import { GetPaymentType, GetEmployee, GetPayment,GetPaymentById } from "../services/HttpClientServicePayment";


export default function Navbar() {

  const [token, settoken] = React.useState<String>("");
  const [role, setrole] = React.useState<String>("");
  const menu_admin = [

    { name: "บันทึกการจ่ายยาและเวชภัณฑ์", icon: <CreateIcon />, path: "/medicinerecord/create" },
    { name: "ข้อมูลบันทึกการจ่ายยาและเวชภัณฑ์", icon: <AssignmentIcon />, path: "/medicinerecords" },
    { name: "บันทึกการชำระเงิน", icon: <CreateIcon />, path: "/payment/create" },
    { name: "ข้อมูลบันทึกการชำระเงิน", icon: <AssignmentIcon />, path: "/payments" },
  ]
  // const menu_employee = [

  //   { name: "บันทึกการจองใช้ห้อง", icon: <BedroomChildIcon />, path: "/bookingcreate" },
  //   { name: "รายการจองใช้ห้อง", icon: <MenuBookIcon />, path: "/bookinghistory" },
  // ]
  const [employee, setEmployee] = React.useState<EmployeesInterface>();
  const [openDrawer, setOpenDrawer] = useState(false);

  const getEmployee = async (ID: string | null) => {
    let res = await GetEmployee()
    if (res) {
      setEmployee(res);
    }
  }


  const toggleDrawer = (state: boolean) => (event: any) => {
    if (event.type === "keydown" && (event.key === "Tab" || event.key === "Shift")) {
      return;
    }
    setOpenDrawer(state);
  }


  const Signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };

  useEffect(() => {
    getEmployee(localStorage.getItem("id"));
   
    const token = localStorage.getItem("token")
    const role = localStorage.getItem("role")
    if (token) {
      settoken(token);
    }
    if (role) {
      setrole(role);
    }
  }, []);

  if (role === 'employee') {
    return (
      <div style={{ flexGrow: 1 }}>
        <AppBar position="static" style={{ background: '#4db6ac' }}>
          <Toolbar>
            <Tooltip title="เมนู">
              <IconButton
                onClick={toggleDrawer(true)}
                edge="start"
                sx={{ color: '#e0f2f1',marginRight: 2 }}
                color="inherit"
                aria-label="menu"
              >
                <MenuIcon />
              </IconButton>
            </Tooltip>
            
            <Drawer open={openDrawer} onClose={toggleDrawer(false)}>
              <List
                onClick={toggleDrawer(false)}
                onKeyDown={toggleDrawer(false)}
              >
                {menu_admin.map((item, index) => (
                  <ListItem key={index} button component={RouterLink} to={item.path}>
                    <ListItemIcon>{item.icon}</ListItemIcon>
                    <ListItemText>{item.name}</ListItemText>
                  </ListItem>
                ))}
              </List>
            </Drawer>
            
            <Tooltip title="หน้าแรก">
              <IconButton
                sx={{color: '#e0f2f1', marginRight:1.5}}
                edge="start"
                component={RouterLink} to="/"
              >
                <HomeIcon />
              </IconButton>  
            </Tooltip>

            <Typography variant="h6" sx={{ flexGrow: 1}}>
              ระบบจัดการคนไข้นอก
            </Typography>

            {/* <Tooltip title="ผู้ดูแลระบบที่เข้าใช้งาน">
            <Chip 
              variant="outlined" 
              icon={<FaceIcon style={{ color: '#00796b' }} />}
              style={{ backgroundColor: '#e0f2f1', fontSize: '1rem', color: '#00796b'}}
              label={employee?.FirstName+" "+employee?.LastName+" (ผู้ดูแลระบบ)"}
            />
            </Tooltip> */}

            <Tooltip title="ออกจากระบบ">
              <IconButton
                onClick={Signout}
                edge="end"
                sx={{ marginLeft: 2 ,color: '#e0f2f1'}}
              >
                <LogoutIcon />
              </IconButton>  
            </Tooltip>
          </Toolbar>
        </AppBar>
      </div>
    );
  }

  // if (role === 'user') {


  //   return (
  //     <div style={{ flexGrow: 1 }}>
  //       <AppBar position="static">
  //         <Toolbar>
  //           <IconButton
  //             onClick={toggleDrawer(true)}
  //             edge="start"
  //             sx={{ marginRight: 2 }}
  //             color="inherit"
  //             aria-label="menu"
  //           >
  //             <MenuIcon />
  //           </IconButton>
  //           <Drawer open={openDrawer} onClose={toggleDrawer(false)}>
  //             <List
  //               onClick={toggleDrawer(false)}
  //               onKeyDown={toggleDrawer(false)}
  //             >
  //               <ListItem button component={RouterLink} to="/">
  //                 <ListItemIcon><HomeIcon /></ListItemIcon>
  //                 <ListItemText>หน้าแรก</ListItemText>
  //               </ListItem>
  //               <Divider />
  //               {menu_user.map((item, index) => (
  //                 <ListItem key={index} button component={RouterLink} to={item.path}>
  //                   <ListItemIcon>{item.icon}</ListItemIcon>
  //                   <ListItemText>{item.name}</ListItemText>
  //                 </ListItem>
  //               ))}

  //               <ListItem button onClick={Signout}>
  //                 <ListItemIcon> <ExitToAppIcon /></ListItemIcon>
  //                 <ListItemText>Sign out</ListItemText>
  //               </ListItem>

  //             </List>
  //           </Drawer>

  //           <Link style={{ color: "white", textDecoration: "none", marginRight: "auto" }} to="/">
  //             <Typography variant="h6" sx={{ flexGrow: 1 }}>
  //               ระบบการจองใช้ห้อง
  //             </Typography>
  //           </Link>

  //         </Toolbar>
  //       </AppBar>
  //     </div>
  //   );
  // }
  
  return <SignIn/>;
}
