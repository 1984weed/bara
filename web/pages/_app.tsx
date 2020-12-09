import CssBaseline from "@material-ui/core/CssBaseline"
import { ThemeProvider } from "@material-ui/core/styles"
import { ClientContext, GraphQLClient } from "graphql-hooks"
import { Provider } from 'next-auth/client'
import { AppPropsType } from "next/dist/next-server/lib/utils"
import React from "react"
import { useGraphQLClient } from "../lib/graphql-client"
import { theme } from "../lib/theme"

interface Props extends AppPropsType {
    graphQLClient: GraphQLClient
    session: any
}

export default function App({ Component, pageProps }) {
    const graphQLClient = useGraphQLClient(pageProps.initState)

    return (
        <Provider 
        options={{
            clientMaxAge: 0,
            keepAlive: 0
          }}
        session={pageProps.session}>
            <ClientContext.Provider value={graphQLClient}>
                <ThemeProvider theme={theme}>
                    <CssBaseline />
                    <Component {...pageProps} />
                </ThemeProvider>
            </ClientContext.Provider>
        </Provider>
    )
}
