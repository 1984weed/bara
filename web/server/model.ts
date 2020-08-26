import * as session from "express-session"
import { AuthStore } from "./auth-store"

export type AuthSeverConfig = {
    port: number
    store: AuthStore
    cookieName: string
    sessionMaxAge: number
    sessionCookie: string
    sessionSecret: string
    domain: string
    pathPrefix: string
    providers: Provider[]
    sessionRevalidateAge?: number
}

export type Provider = {
    name: string
    options?: string[]
    strategy: any
    strategyOptions: {
        consumerKey: string
        consumerSecret: string
        profileFields: string[]
        includeEmail?: boolean,
        callbackURL?: string
        passReqToCallback?: boolean
    }
    getProfile(profile: any): ProviderProfile
}

export type ProviderProfile = {
    id: string
    name: string
    displayName: string
    imageUrl: string
    email: string
}

export type OAuthUser = {
    id: string
    accessToken: string
    refreshToken: string
    providerName: string
}


export type DBConfig = {
    user: string
    host?: string
    database: string
    password: string
    port?: string
}

export type User = {
    id?: number
    provider?: string
    userName?: string
    displayName?: string
    email?: string
    password?: string
    imageUrl?: string
}
