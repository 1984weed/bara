import { GraphQLClient } from "graphql-hooks";
import unfetch from "isomorphic-unfetch";
import memCache from "graphql-hooks-memcache";
const graphqlURL = process.env.GRAPHQL_URL;

let graphQLClient: null | GraphQLClient = null;

function create(
    this: unknown,
    state: { headers?: any } = {},
    ssr: boolean = false
): GraphQLClient {
    return new GraphQLClient({
        ssrMode: ssr,
        url: graphqlURL,
        cache: ssr ? memCache() : undefined,
        headers: {
            ...state.headers
        },
        fetchOptions: {
            mode: "cors",
            credentials: "include"
        },
        fetch: typeof window !== "undefined" ? fetch.bind(this) : unfetch // eslint-disable-line
    });
}

export default function initGraphQL(initialState?: object) {
    if (typeof window === "undefined") {
        return create(initialState, true);
    }

    if (!graphQLClient) {
        graphQLClient = create(initialState);
    }

    return graphQLClient;
}
