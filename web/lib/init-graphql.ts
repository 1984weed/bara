import { GraphQLClient } from "graphql-hooks"
import unfetch from "isomorphic-unfetch"
import memCache from "graphql-hooks-memcache"
const graphqlURL = process.env.GRAPHQL_URL

let graphQLClient: null | GraphQLClient = null
import axios from "axios"
import { buildAxiosFetch } from "@lifeomic/axios-fetch"

const gqlAxios = axios.create()
gqlAxios.interceptors.response.use(
    function(response) {
        return response
    },
    async error => {
        // Pass all errors except 401
        if (error.response.status !== 401) {
            return Promise.reject(error)
        }

        const tokenRes = await getIdToken()

        localStorage.setItem("token", tokenRes.token)
        // axios.defaults.headers.common["Authorization"] = `Bearer ${tokenRes.token}`

        error.hasRefreshedToken = true
        return Promise.reject(error)
    }
)

const rootURL = "http://localhost:3000"

export const getIdToken = () => {
    return fetch(`${rootURL}/auth/getToken`)
        .then(req => req.json())
        .then(res => {
            return res
        })
}

function create(this: unknown, state: { headers?: any } = {}, ssr: boolean = false): GraphQLClient {
    return new GraphQLClient({
        ssrMode: ssr,
        url: graphqlURL,
        cache: ssr ? memCache() : undefined,
        headers: {
            ...state.headers,
        },
        fetchOptions: {
            mode: "cors",
            credentials: "include",
        },
        fetch: buildAxiosFetch(gqlAxios), 
    })
}

export default function initGraphQL(initialState?: object) {
    if (typeof window === "undefined") {
        return create(initialState, true)
    }

    if (!graphQLClient) {
        graphQLClient = create(initialState)
    }

    return graphQLClient
}
