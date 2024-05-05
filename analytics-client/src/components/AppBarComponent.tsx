import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from "@mui/material/Button";
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import { Link } from 'react-router-dom';

const AppBarComponent = () => {
    const pages = ['Transactions', "Analytics-Groups", "Analytics-Dates"];

    return (
        <AppBar position="static">
            <Toolbar variant="dense">
                <AccountBalanceIcon />
                <Typography
                    noWrap
                    component={Link}
                    to="/"
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
                <Box sx={{ flexGrow: 1, display: 'flex' }}>
                    {pages.map((page) => (
                        <Button component={Link} to={"/" + page.toLowerCase()} sx={{ my: 3, color: 'white', display: 'block' }}>
                            {page}
                        </Button>
                    ))}
                </Box>
            </Toolbar>
        </AppBar>
    );
};

export default AppBarComponent;