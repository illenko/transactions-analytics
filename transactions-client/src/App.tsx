import {BrowserRouter, Route, Routes} from "react-router-dom";
import Layout from "./pages/Layout.tsx";
import Transactions from "./pages/Transactions";
import {FC} from "react";
import {AppProps} from "./App.types";
import Statistics from "./pages/Statistics";
import {createTheme, ThemeProvider} from "@mui/material";
import Transaction from "./pages/Transaction";

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

const App: FC<AppProps> = ({title}) => {
    return (
        <ThemeProvider theme={theme}>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Layout/>}>
                        <Route path="/transactions" element={<Transactions title={title}/>}/>
                        <Route path="/transactions/:id" element={<Transaction title={title}/>}/>
                        <Route path="/statistics" element={<Statistics title={"Statistics"}/>}/>
                    </Route>
                </Routes>
            </BrowserRouter>
        </ThemeProvider>
    );
};

export default App;