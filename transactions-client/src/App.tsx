import {BrowserRouter, Route, Routes} from "react-router-dom";
import Layout from "./pages/Layout.tsx";
import Transactions from "./pages/Transactions";
import {FC} from "react";
import {createTheme, ThemeProvider} from "@mui/material";
import AnalyticGroups from "./pages/AnalyticGroups.tsx";
import AnalyticDates from "./pages/AnalyticDates.tsx";

const theme = createTheme({
    components: {
        MuiToolbar: {
            styleOverrides: {
                dense: {
                    height: 48,
                    minHeight: 32
                }
            }
        }
    },
})

const App: FC = () => {
    return (
        <ThemeProvider theme={theme}>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Layout/>}>
                        <Route path="/transactions" element={<Transactions/>}/>
                        <Route path="/analytic-groups" element={<AnalyticGroups/>}/>
                        <Route path="/analytic-dates" element={<AnalyticDates/>}/>
                    </Route>
                </Routes>
            </BrowserRouter>
        </ThemeProvider>
    );
};

export default App;