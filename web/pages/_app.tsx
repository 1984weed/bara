import { ClientContext, GraphQLClient } from "graphql-hooks"
import { AppContext } from "next/app"
import { AppPropsType, NextComponentType } from "next/dist/next-server/lib/utils"
import React, { useEffect } from "react"
// import withGraphQLClient, { NextPageContextWithGraphql } from "../lib/with-graphql-client"
import { Session } from "../types/Session"
import { useGraphQLClient } from "../lib/graphql-client"
import { useSession, Provider } from "../lib/session"
import { ThemeProvider } from "@material-ui/core/styles"
import CssBaseline from "@material-ui/core/CssBaseline"
import { theme } from "../lib/theme"

interface Props extends AppPropsType {
    graphQLClient: GraphQLClient
    session: any
}

export default function App({ Component, pageProps }) {
    const graphQLClient = useGraphQLClient(pageProps.initState)

    return (
        <Provider session={pageProps.session}>
            <ClientContext.Provider value={graphQLClient}>
                <ThemeProvider theme={theme}>
                    <CssBaseline />
                    <Component {...pageProps} />
                </ThemeProvider>
            </ClientContext.Provider>
        </Provider>
    )
}
