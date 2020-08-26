import React from 'react'
import initGraphQL from './init-graphql'
import Head from 'next/head'
import { getInitialState } from 'graphql-hooks-ssr'
import { GraphQLClient } from 'graphql-hooks'
import { NextPageContext } from 'next'
import cookies from 'next-cookies'
import { Session } from '../types/Session'

type Props = {
  graphQLState: object;
  authToken: string;
}

export interface NextPageContextWithGraphql extends NextPageContext {
  client: GraphQLClient;
  session: Session;
}

export default (MyAppI: any) => {
  return class GraphQLHooks extends React.Component {
    private graphQLClient: GraphQLClient;
    static displayName = 'GraphQLHooks(App)'
    static async getInitialProps (app: any) {
      const { Component, router, ctx } = app
      const isSever = typeof window === 'undefined';
      const initState = isSever ? {headers: {Authorization: cookies(ctx)["auth-token"], ...ctx.req.header}} : null
      const graphQLClient = initGraphQL(initState)

      let appProps = {}
      if (MyAppI.getInitialProps != null) {
        appProps = await MyAppI.getInitialProps({client: graphQLClient, ...app})
      }

      let graphQLState = {}
      if (isSever) {
        try {
          // Run all GraphQL queries
          graphQLState = await getInitialState({
            App: (
              <MyAppI
                pageProps={null}
                {...appProps}
                Component={Component}
                router={router}
                graphQLClient={graphQLClient}
              />
            ),
            client: graphQLClient
          })
        } catch (error) {
          console.error('Error while running `getInitialState`', error)
        }

        Head.rewind()
      }

      return {
        ...appProps,
        graphQLState,
       authToken: cookies(ctx)["auth-token"]
      }
    }

    constructor (props: Props) {
      super(props)
      this.graphQLClient = initGraphQL({...props.graphQLState, ...{headers: {Authorization: props.authToken}}})
    }

    render () {
      return <MyAppI {...this.props} graphQLClient={this.graphQLClient} />
    }
  }
}