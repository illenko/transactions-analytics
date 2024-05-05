import {BrowserRouter} from "react-router-dom";
import {FC} from "react";
import {ThemeProvider} from "@mui/material";
import AppRoutes from "./components/Routes";
import theme from "./theme";

const App: FC = () => {
    return (
        <ThemeProvider theme={theme}>
            <BrowserRouter>
                <AppRoutes/>
            </BrowserRouter>
        </ThemeProvider>
    );
};

export default App;