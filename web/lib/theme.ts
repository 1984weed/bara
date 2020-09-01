import { createMuiTheme, responsiveFontSizes } from "@material-ui/core"

export const theme = responsiveFontSizes(
    createMuiTheme({
        palette: {
            background: {
                default: "#fff",
            },
        },
    })
)