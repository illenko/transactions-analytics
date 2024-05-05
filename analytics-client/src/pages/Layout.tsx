import {Outlet} from "react-router-dom";
import Box from '@mui/material/Box';
import AppBarComponent from '../components/AppBarComponent';

const Layout = () => {
    return (
        <Box>
            <AppBarComponent />
            <div className="page">
                <Outlet/>
            </div>
        </Box>
    )
};

export default Layout;