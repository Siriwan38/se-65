import React from "react";

import Container from "@mui/material/Container";
import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import Box from "@mui/material/Box";

import PaymentReq from "./PaymentReq";
import MedicineRecordReq from "./MedicineRecordReq";

function Home() {
  const [value, setValue] = React.useState("1");

  const handleChange = (event: React.ChangeEvent<{}>, newValue: string) => {
    setValue(newValue);
  }

  return (
    <div>
      <Container sx={{ marginTop: 2}} maxWidth="md">
        
        <TabContext value={value}>
          <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <TabList onChange={handleChange} variant="scrollable" scrollButtons={true} allowScrollButtonsMobile aria-label="lab API tabs example">
              <Tab label="ระบบบันทึกการจองใช้ห้อง" value="1" />
              <Tab label="ระบบบันทึกการชำระเงิน" value="2" />
              
            </TabList>
          </Box>
          <TabPanel value="1">
            <MedicineRecordReq />
          </TabPanel>
          <TabPanel value="2">
            <PaymentReq />
          </TabPanel>
        
        </TabContext>
      </Container>
      
    </div>
  );
}
export default Home;