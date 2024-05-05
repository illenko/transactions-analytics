import {createTheme} from "@mui/material";

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

export default theme;