import { useMemo } from "react"
import { GraphQLClient } from "graphql-hooks"
import memCache from "graphql-hooks-memcache"

const graphqlURL = process.env.GRAPHQL_URL
import axios from "axios"
import { buildAxiosFetch } from "@lifeomic/axios-fetch"



const gqlAxios = axios.create()
gqlAxios.interceptors.request.use(async function(config) {
    let token = ""
    console.log("gqlAxios.interceptors.request")
    if(window) {
        token = window.localStorage.getItem("api-token")
    }

    // If local doesn't have api-token, it needs to get new one
    if (!token) {
        const tokenRes = await getIdToken()
        token = tokenRes.token
        window && window.localStorage.setItem("api-token", token)
    }
    config.headers.Authorization = `Bearer ${token}`

    return config
})

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

        localStorage.setItem("api-token", tokenRes.token)
        axios.defaults.headers.common["Authorization"] = `Bearer ${tokenRes.token}`

        error.hasRefreshedToken = true
        return Promise.reject(error)
    }
)

export const getIdToken = () => {
    return fetch(`/api/bara/jwt`)
        .then(req => req.json())
        .then(res => {
            return res
        })
}

let graphQLClient

function createClient(initialState) {
    return new GraphQLClient({
        ssrMode: false,//typeof window === "undefined",
        url: graphqlURL,
        cache: memCache({ initialState }),
        fetchOptions: {
            mode: "cors",
            // credentials: "include",
        },
        fetch: buildAxiosFetch(gqlAxios),
    })
}

export function initializeGraphQL(initialState = null) {
    const _graphQLClient = graphQLClient ?? createClient(initialState)

    // After navigating to a page with an initial GraphQL state, create a new cache with the
    // current state merged with the incoming state and set it to the GraphQL client.
    // This is necessary because the initial state of `memCache` can only be set once
    if (initialState && graphQLClient) {
        graphQLClient.cache = memCache({
            initialState: Object.assign(graphQLClient.cache.getInitialState(), initialState),
        })
    }
    // For SSG and SSR always create a new GraphQL Client
    if (typeof window === "undefined") return _graphQLClient
    // Create the GraphQL Client once in the client
    if (!graphQLClient) graphQLClient = _graphQLClient

    return _graphQLClient
}

export function useGraphQLClient(initialState) {
    const store = useMemo(() => initializeGraphQL(initialState), [initialState])
    return store
}
