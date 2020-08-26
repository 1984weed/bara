import { ClientContext, GraphQLClient } from "graphql-hooks"
import App, { AppContext } from "next/app"
import { AppPropsType, NextComponentType } from "next/dist/next-server/lib/utils"
import React from "react"
import withGraphQLClient, { NextPageContextWithGraphql } from "../lib/with-graphql-client"
import { Session } from "../types/Session"

interface Props extends AppPropsType {
    graphQLClient: GraphQLClient
    session: any
}

interface AppContextWithGraphql extends AppContext {
    client: GraphQLClient
    Component: NextComponentType<NextPageContextWithGraphql>
}

function getSession(req: any): Promise<Session> {
    return new Promise((resolve, reject) => {
        let session: Session

        if (typeof window === "undefined") {
            session = (req as any).session?.passport
            resolve(session)
            return
        } else {
            const baseUrl = req ? `${(req as any).protocol}://${(req as any).get("Host")}` : ""

            fetch(`${baseUrl}/auth/session`, {
                credentials: "same-origin",
            }).then(res => {
                if (res.ok) {
                    return res.json()
                } else {
                    reject()
                }
            }).then((res) => 
                resolve(res)
            )

        }

    })

}

export class MyApp extends App<Props> {
    static async getInitialProps({ Component, ctx, client }: AppContextWithGraphql) {
        let pageProps = {}
        const { req } = ctx
        const session = await getSession(req)

        if (Component.getInitialProps) {
            pageProps = await Component.getInitialProps({ session, client, ...ctx })
        }

        return { pageProps, session }
    }
    render() {
        const { Component, pageProps, graphQLClient, session } = this.props

        return (
            <ClientContext.Provider value={graphQLClient}>
                <Component {...pageProps} session={session} />
            </ClientContext.Provider>
        )
    }
}

export default withGraphQLClient(MyApp)
