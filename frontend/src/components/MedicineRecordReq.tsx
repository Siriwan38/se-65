import { Box, Button, Divider, Grid, Paper, Typography } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";

import { Container } from "@mui/material";

export default function Home() {
    const token: string | null = localStorage.getItem("token"); // ตัวแปร token จะดึงค่า token ที่อยู่ใน local storage

    // เมื่อมีการ log out ฟังก์ชันนี้จะทำการ clear token ใน local storage และเปลี่ยน path ไปที่หน้า log in
    const logOut = () => {
        localStorage.clear();
        window.location.href = "/login";
    }
    return (
        <Container maxWidth={"md"}>
            <Paper>
                <Box display="flex" sx={{ marginTop: 10, }}>
                    <Typography component="h2" variant="h4" color="primary" sx={{ marginTop: 2, marginBottom: 2, marginLeft: 39 }}>
                        ระบบบันทึกการจ่ายยาและเวชภัณฑ์
                    </Typography>
                </Box>
                <Divider />
                <Box>
                    <Box sx={{ paddingX: 2, paddingY: 1, }}>
                        <Typography component="h2" variant="h5" color="primary" sx={{ marginTop: 2 }}>
                            Requirement(บันทึกการจ่ายยาและเวชภัณฑ์)
                        </Typography>
                        <Typography component="h2" variant="h6" color="action" sx={{ paddingX: 5 }}>
                            ระบบสมัครสมาชิกการจองใช้ห้องของบริษัท Room Booking เป็นระบบที่ผู้ใช้บริการ
                        </Typography>
                        <Typography component="h2" variant="h6" color="action">
                            สามารถสมัครเป็นสมาชิกของระบบสำหรับการจองใช้ห้องเรียนทฤษฎี ห้องเรียนปฏิบัติการและห้องสมุด โดยการสมัครสมาชิกจะมีพนักงานระบบเป็นผู้ทำการบักทึกข้อมูลของผู้ใช้บริการเช่น คำนำหน้าชื่อ ชื่อ-สกุล วันเดือนปีเกิด เลขประจำตัวประชาชน อีเมล รหัสผ่าน เพศ หมายเลขโทรศัพท์ ที่อยู่ (เลขที่บ้าน, หมู่บ้าน, ตำบล, อำเภอ, จังหวัด) และพนักงานบันทึกข้อมูล เป็นต้น
                        </Typography>
                        <Typography component="h2" variant="h6" color="action" sx={{ paddingX: 5 }}>
                            นอกจากนี้ ผู้ใช้บริการซึ่งเป็นสมาชิกสามารถเลือกใช้บริการระบบการจองใช้ห้องได้ตามที่
                        </Typography>
                        <Typography component="h2" variant="h6" color="action">
                            ต้องการ รวมทั้งการบริการต่างๆที่มีอยู่ในระบบ
                        </Typography>
                        <Typography component="h2" variant="h5" color="primary" sx={{ marginTop: 4 }}>
                            User Story(ระบบสมัครสมาชิก)
                        </Typography>
                        <Typography component="h2" variant="h6" color="action">
                            <b>ในบทบาทของ</b> พนักงานระบบ
                        </Typography>
                        <Typography component="h2" variant="h6" color="action">
                            <b>ฉันต้องการ</b> บันทึกข้อมูลของผู้ใช้บริการ
                        </Typography>
                        <Typography component="h2" variant="h6" color="action">
                            <b>เพื่อ</b> ให้ได้ข้อมูลของผู้ใช้บริการ
                        </Typography>
                    </Box>
                </Box>
                <Divider />


                </Paper>
        </Container>
    )
}