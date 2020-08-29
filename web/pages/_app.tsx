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
    // useEffect(() => {
    //     // Remove the server-side injected CSS.
    //     const jssStyles = document.querySelector('#jss-server-side');
    //     if (jssStyles) {
    //       jssStyles.parentElement.removeChild(jssStyles);
    //     }
    //   }, []);


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

// export class MyApp extends App<Props> {
//     // static async getInitialProps({ Component, ctx, client }: AppContextWithGraphql) {
//     //     let pageProps = {}
//     //     const { req } = ctx
//     //     const session = await getSession(req)

//     //     if (Component.getInitialProps) {
//     //         pageProps = await Component.getInitialProps({ session, client, ...ctx })
//     //     }

//     //     return { pageProps, session }
//     // }
//     render() {
//         const { Component, pageProps, graphQLClient, session } = this.props

//         return (
//             <ClientContext.Provider value={graphQLClient}>
//                 <Component {...pageProps} session={session} />
//             </ClientContext.Provider>
//         )
//     }
// }

// export default withGraphQLClient(MyApp)
// export default MyApp
