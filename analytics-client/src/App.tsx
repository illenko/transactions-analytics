import {BrowserRouter, Route, Routes} from "react-router-dom";
import Layout from "./pages/Layout.tsx";
import Transactions from "./pages/Transactions";
import {FC} from "react";
import {createTheme, ThemeProvider} from "@mui/material";
import AnalyticsGroups from "./pages/AnalyticsGroups.tsx";
import AnalyticsDates from "./pages/AnalyticsDates.tsx";

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
                        <Route path="/analytics-groups" element={<AnalyticsGroups/>}/>
                        <Route path="/analytics-dates" element={<AnalyticsDates/>}/>
                    </Route>
                </Routes>
            </BrowserRouter>
        </ThemeProvider>
    );
};

export default App;