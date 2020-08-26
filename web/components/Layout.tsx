import Footer from "./Footer"
import Header from "./Header"
import { Session } from "../types/Session"
import { Container, CssBaseline, ThemeProvider, createMuiTheme, responsiveFontSizes, Box } from "@material-ui/core"
import Head from "next/head"

type Props = {
    title?: string
    session: Session
}

const theme = responsiveFontSizes(createMuiTheme({
    palette: {
        background:{
            default: "#fff"
        }
    }
}));

const Layout: React.FunctionComponent<Props> = ({ title, children, session }) => (
    <ThemeProvider theme={theme}>
        <CssBaseline />
        <Head>
            <title>{title || "default"}</title>
        </Head>
        <Header session={session} />
        <main>
            <Container maxWidth="md"><Box pt={3}>{children}</Box></Container>
        </main>
        <Footer />
    </ThemeProvider>
)

export default Layout
