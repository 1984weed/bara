import App from 'next/app'
import React from 'react'
import withGraphQLClient from '../lib/with-graphql-client'
import { ClientContext, GraphQLClient } from 'graphql-hooks'
import { NextComponentType, NextPageContext } from 'next'
import { Router } from 'next/dist/client/router'

type Props = {
    graphQLClient: GraphQLClient;
    Component: NextComponentType<NextPageContext>;
    router: Router;
    pageProps: any;
}

export class MyApp extends App<Props> {
  render () {
    const { Component, pageProps, graphQLClient } = this.props
    return (
      <ClientContext.Provider value={graphQLClient}>
        <Component {...pageProps} />
      </ClientContext.Provider>
    )
  }
}

export default withGraphQLClient(MyApp)
