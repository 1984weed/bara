import { GraphQLClient } from 'graphql-hooks'
// import memCache from 'graphql-hooks-memcache'
import unfetch from 'isomorphic-unfetch'

let graphQLClient: null | GraphQLClient = null

function create (this: unknown, _: object = {}): GraphQLClient {
  return new GraphQLClient({
    ssrMode: false,//typeof window === 'undefined',
	url: 'https://api.graph.cool/simple/v1/cixmkt2ul01q00122mksg82pn',
	cache: undefined,
    fetch: typeof window !== 'undefined' ? fetch.bind(this) : unfetch, // eslint-disable-line
  })
}

export default function initGraphQL (initialState?: object ) {
  // Make sure to create a new client for every server-side request so that data
  // isn't shared between connections (which would be bad)
  if (typeof window === 'undefined') {
    return create(initialState)
  }

  // Reuse client on the client-side
  if (!graphQLClient) {
    graphQLClient = create(initialState)
  }

  return graphQLClient
}