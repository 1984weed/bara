import { Box, Container, CssBaseline } from "@material-ui/core"
import Head from "next/head"
import { useSession } from "../lib/session"
import Footer from "./Footer"
import Header from "./Header"

type Props = {
    title?: string
}

const Layout: React.FunctionComponent<Props> = ({ title, children }) => {
    const [, load] = useSession()

    if(load) {
        return <></>
    }

    return (<>
        <CssBaseline />
        <Head>
            <title>{title || "default"}</title>
        </Head>
        <Header />
        <main>
            <Container maxWidth="md">
                <Box pt={3}>{children}</Box>
            </Container>
        </main>
        <Footer />
    </>)
}

export default Layout
