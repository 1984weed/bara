import NextAuth from "next-auth"
import Providers from "next-auth/providers"
import * as jwt from "jsonwebtoken"
import { Pool } from "pg"
import { hash, compare } from "bcrypt"

const saltRounds = 10
const ADMIN_PASSWORD = process.env.ADMIN_PASSWORD || "admin"
const ADMIN_EMAIL = process.env.ADMIN_PASSWORD || "adminatmail.com"
const ADMIN_IMAGE = process.env.ADMIN_IMAGE || "/admin/profile"

const postgresURI = "postgres://postgres:postgres@localhost:5432/bara"
const store = new Pool({ connectionString: postgresURI })

initSetupAdmin();

// Admin user always exists, this method can create Admin ser
async function initSetupAdmin() {
    const res = await store.query({
        text: `SELECT * FROM users WHERE role = 'admin'`
    })

    // If an admin user already exists on database, skip it 
    if (res.rows.length === 1) {
        return Promise.resolve()
    }

    const hashPass = await getHashPassword(ADMIN_PASSWORD)

    try{
        await store.query({
            text: `INSERT INTO users(name, unique_name, role, password, email, image, updated_at, created_at) VALUES ('admin', 'admin', 'admin', $1, $2, $3, current_timestamp, current_timestamp)`,
            values: [hashPass, ADMIN_EMAIL, ADMIN_IMAGE],
        })
    }catch(e) {
        console.log(e)
        return Promise.reject("Error: While it inserts the admin user, something happened.")
    } 

    return Promise.resolve()

}

function getHashPassword(password: string): Promise<string> {
    return new Promise<string>((resolve, reject) => {
        hash(password, saltRounds, function(err, hash) {
            resolve(hash)
        })
    })
}

async function comparePassword(plain: string, hash: string): Promise<boolean> {
    return compare(plain, hash)
}

// For more information on each option (and a full list of options) go to
// https://next-auth.js.org/configuration/options
const options = {
    // https://next-auth.js.org/configuration/providers
    providers: [
        Providers.Google({
            clientId: process.env.GOOGLE_ID,
            clientSecret: process.env.GOOGLE_SECRET,
        }),
        Providers.Credentials({
            id: "credentials",
            name: "Addmin Account",
            authorize: async (credentials) => {
                const adminSignInURL = "/admin/signin"
                let res
                try{
                    res = await store.query({
                        text: `SELECT id, name, email, email_verified, password, image, created_at, updated_at FROM users WHERE unique_name = $1`,
                        values: [credentials.username],
                    })
                } catch (e){
                    console.log(e)
                }

                if(res.rows.length === 0) {
                    return Promise.reject(adminSignInURL)
                }
                const [adminUser, ] = res.rows
                const equalPass = await comparePassword(credentials.password, adminUser.password)

                if(!equalPass) {
                    return Promise.reject(adminSignInURL)
                }
                // Remove password from user
                adminUser.password = null
                // Role admin
                adminUser.role = "admin" 
                delete adminUser["password"]

                return Promise.resolve(adminUser)
            },
            credentials: {
                username: { label: "Username", type: "text ", placeholder: "admin" },
                password: { label: "Password", type: "password" },
            },
        }),
    ],
    database: "postgres://postgres:postgres@localhost:5432/bara",

    // The secret should be set to a reasonably long random string.
    // It is used to sign cookies and to sign and encrypt JSON Web Tokens, unless
    // a seperate secret is defined explicitly for encrypting the JWT.
    secret: process.env.SECRET,

    session: {
        // Use JSON Web Tokens for session instead of database sessions.
        // This option can be used with or without a database for users/accounts.
        // Note: `jwt` is automatically set to `true` if no database is specified.
        jwt: true,

        // Seconds - How long until an idle session expires and is no longer valid.
        // maxAge: 30 * 24 * 60 * 60, // 30 days

        // Seconds - Throttle how frequently to write to database to extend a session.
        // Use it to limit write operations. Set to 0 to always update the database.
        // Note: This option is ignored if using JSON Web Tokens
        // updateAge: 24 * 60 * 60, // 24 hours
    },

    // JSON Web tokens are only used for sessions if the `jwt: true` session
    // option is set - or by default if no database is specified.
    // https://next-auth.js.org/configuration/options#jwt
    jwt: {
        // A secret to use for key generation (you should set this explicitly)
        secret: process.env.JWT_SECRET || "secret", //'INp8IvdIyeMcoGAgFGoA61DdBglwwSqnXJZkgz8PSnw',
        // raw: true,
        // cookieName: "bara",
        // Set to true to use encryption (default: false)
        // encryption: true,
        // You can define your own encode/decode functions for signing and encryption
        // if you want to override the default behaviour.
        // encode: async ({ secret, token, maxAge }) => {},
        // decode: async ({ secret, token, maxAge }) => {},
    },

    // You can define custom pages to override the built-in pages.
    // The routes shown here are the default URLs that will be used when a custom
    // pages is not specified for that route.
    // https://next-auth.js.org/configuration/pages
    pages: {
        signIn: "/signin", // Displays signin buttons
        // signOut: '/api/auth/signout', // Displays form with sign out button
        // error: '/api/auth/error', // Error code passed in query string as ?error=
        // verifyRequest: '/api/auth/verify-request', // Used for check email page
        // newUser: null // If set, new users will be directed here on first sign in
    },

    // Callbacks are asynchronous functions you can use to control what happens
    // when an action is performed.
    // https://next-auth.js.org/configuration/callbacks
    callbacks: {
        signIn: async (user, account, profile) => {
            user.accessToken = getAccessToken(user)

            return Promise.resolve(true)
        },
        redirect: async (url, baseUrl) => {
            // It always redirects to top page
            return Promise.resolve(baseUrl)
        },
        jwt: async (token, user, account, profile, isNewUser) => {
            if (user) {
                token.sub = `${user.id}`
                token.role = user.role || "normal"
                token.accessToken = getAccessToken(token)
            }


            return Promise.resolve(token)
        },
        session: async (session, token) => {
            session.accessToken = token.accessToken
            // session user includes their role
            session.user.role = token.role

            return Promise.resolve(session)
        },
    },

    // Events are useful for logging
    // https://next-auth.js.org/configuration/events
    events: {},

    // Enable debug messages in the console if you are having problems
    debug: false,
}

const jwtSecret = process.env.JWT_SECRET || "secret"
const jwtOptions: jwt.SignOptions = {
    algorithm: "HS256",
}

function getAccessToken(user) {
    const jwtPayload = {
        sub: user.sub,
        role: user.role || "normal",
    }

    return jwt.sign(jwtPayload, jwtSecret, jwtOptions)
}

export default (req, res) => NextAuth(req, res, options)
