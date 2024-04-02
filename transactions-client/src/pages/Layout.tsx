import {Outlet} from "react-router-dom";
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from "@mui/material/Button";
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';

const Layout = () => {
    const pages = ['Transactions', "Analytic-Groups", "Analytic-Dates"];

    return (
        <Box>
            <AppBar position="static">
                <Toolbar variant="dense">
                    <AccountBalanceIcon></AccountBalanceIcon>
                    <Typography
                        noWrap
                        component="a"
                        href="/"
                        sx={{
                            mr: 2,
                            fontWeight: 100,
                            fontFamily: 'roboto',
                            color: 'white',
                            letterSpacing: '.2rem',
                            textDecoration: 'none',
                        }}
                    >
                        Banking App
                    </Typography>
                    <Box sx={{flexGrow: 1, display: 'flex'}}>
                        {pages.map((page) => (
                            <Button href={"/" + page.toLowerCase()} sx={{my: 3, color: 'white', display: 'block'}}>
                                {page}
                            </Button>
                        ))}
                    </Box>
                </Toolbar>
            </AppBar>

            <div className="page">
                <Outlet/>
            </div>
        </Box>
    )
};

export default Layout;