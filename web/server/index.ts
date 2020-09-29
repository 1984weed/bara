import expressSession, * as session from "express-session"
import next from "next"
import redis from "redis"
import { createProviders, ProviderKeys } from "./auth-providers"
import authServer from "./auth-server"
import { createDBStore } from "./auth-store"
const RedisStore = require("connect-redis")(expressSession)

// Load environment variables from .env
require("dotenv").load()

// Initialize Next.js
const nextApp = next({
    dir: ".",
    dev: process.env.NODE_ENV === "development",
})

nextApp
    .prepare()
    .then(async () => {
        const port = ((process.env.PORT as unknown) as number) || 3000

        try {
            await authServer(nextApp, {
                port: port,
                sessionSecret: process.env.SESSION_SECRET || "bara",
                sessionMaxAge: 60000 * 60 * 24 * 7,
                cookieName: process.env.COOKIE_NAME || "auth-token",
                pathPrefix: "/auth",
                sessionCookie: "connect.sid",
                sessionStore: initCacheStore(process.env.REDIS_URI),
                jwtSecret: process.env.JWT_SECRET || "secret",
                jwtOptions: {
                    algorithm: "HS256",
                    expiresIn: "60m",
                },
                domain: process.env.SUB_DOMAIN || null,
                store: createDBStore(initDBConfig()),
                providers: createProviders({
                    twitter: initProviderKeys(process.env.TWITTER_CONSUMER_KEY, process.env.TWITTER_CONSUMER_SECRET),
                    github: initProviderKeys(process.env.GITHUB_CONSUMER_KEY, process.env.GITHUB_CONSUMER_SECRET),
                }),
            })
        } catch (e) {
            throw e
        }

        console.log(`Ready on http://localhost:${port}`)
    })
    .catch(err => {
        console.log("An error occurred, unable to start the server")
        console.log(err)
    })

function initDBConfig() {
    if (process.env.PGHOST) {
        return {
            user: process.env.DB_USER,
            host: process.env.PGHOST,
            database: process.env.DB_NAME,
            password: process.env.DB_PASS,
            port: process.env.DB_PORT,
        }
    }

    const dbSocketPath = process.env.DB_SOCKET_PATH || "/cloudsql"

    return {
        user: process.env.DB_USER,
        host: `${dbSocketPath}/${process.env.INSTANCE_CONNECTION_NAME}`,
        database: process.env.DB_NAME,
        password: process.env.DB_PASS,
    }
}

function initCacheStore(redisUri: string): session.Store {
    return new RedisStore({
        client: redis.createClient({
            url: redisUri,
        }),
    })
}

function initProviderKeys(consumerKey: string, consumerSecret: string): ProviderKeys {
    return {
        consumerKey,
        consumerSecret,
    }
}
