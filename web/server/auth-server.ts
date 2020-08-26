import { compare, hash } from "bcrypt"
import BodyParser from "body-parser"
import Express from "express"
import ExpressSession from "express-session"
import Server from "next/dist/next-server/server/next-server"
import passport from "passport"
import passportLocal from "passport-local"
import { AuthSeverConfig, User, OAuthUser } from "./model"
import { logger } from "./logger"
import * as jwt from 'jsonwebtoken';

const LocalStrategy = passportLocal.Strategy

export default (
    nextApp: Server,
    {
        port,
        sessionSecret,
        cookieName,
        store,
        sessionCookie,
        sessionMaxAge = 60000 * 60 * 24 * 7,
        domain,
        sessionRevalidateAge = 60000,
        pathPrefix,
        providers,
    }: AuthSeverConfig
): Promise<void> => {
    const expressApp = Express()

    expressApp.all("/_next/*", (req, res) => {
        let nextRequestHandler = nextApp.getRequestHandler()
        return nextRequestHandler(req, res)
    })
    expressApp.use(BodyParser.json())
    expressApp.use(BodyParser.urlencoded())
    expressApp.use(
        ExpressSession({
            secret: sessionSecret,
            resave: true,
            rolling: true,
            saveUninitialized: true,
            name: cookieName,
            cookie: {
                name: sessionCookie,
                httpOnly: true,
                secure: "auto",
                maxAge: sessionMaxAge,
                domain: domain,
            },
        })
    )
    expressApp.use(passport.initialize())
    expressApp.use(passport.session())

    passport.serializeUser((user, next) => {
        if (!user) {
            next(null, false)
            return
        }
        next(null, user.id)
    })

    passport.deserializeUser((id: number, next) => {
        store
            .findByID(id)
            .then(user => {
                if (!user) {
                    return null
                }
                return {
                    id: user.id,
                    name: user.userName,
                    email: user.email,
                }
            })
            .then(user => {
                if (!user) {
                    return next(null, false)
                }
                return next(null, user)
            })
            .catch(err => {
                next(err, false)
            })
    })

    passport.use(
        new LocalStrategy(
            {
                usernameField: "email",
                passwordField: "password",
            },
            function(username, password, done) {
                store
                    .findByEmail(username)
                    .then(user => {
                        if (user == null) {
                            return done(null, false, { message: "Incorrect username." })
                        }

                        return new Promise((resolve, reject) => {
                            compare(password, user.password, function(err) {
                                if (err) {
                                    return done(null, false, { message: "Incorrect password." })
                                }

                                done(null, user)
                            })
                        })
                    })
                    .catch(err => {
                        done(err)
                    })
            }
        )
    )

    /**
     * Sign in a user
     */
    expressApp.post(`${pathPrefix}/signin`, passport.authenticate("local", { successRedirect: "/", failureRedirect: "/signin" }))

    /**
     * Generate token
     */

     expressApp.get(`${pathPrefix}/getToken`, (req, res) => {
        if(req.user) {

const jwtPayload = {
    email: 'user1@example.com',
    name: 'JWT Taro',
};
const jwtSecret = 'secret_key_goes_here';
const jwtOptions: jwt.SignOptions = {
    algorithm: 'HS256',
    expiresIn: '3s',
};

const token = jwt.sign(jwtPayload, jwtSecret, jwtOptions);
console.log("==================token===============", token)
            return res.json({
                result: "ok",
                token: ""
            })
        }

        return res.json({
            error: "Session timeout"
        })
     })

    /**
     * Sign out a user
     */
    expressApp.post(`${pathPrefix}/signout`, (req, res) => {
        req.logout()
        req.session.destroy(() => {
            return res.redirect("/")
        })
    })

    expressApp.post(`${pathPrefix}/signup`, async (req, res) => {
        const { email, password } = req.body

        return store
            .createUser(email, password)
            .then(id => {
                return {
                    id,
                    email,
                }
            })
            .then(user => {
                req.logIn(user, err => {
                    if (err) return res.redirect(`${pathPrefix}/error?action=signup&type=credentials`)
                    return res.redirect("/")
                })
            }).catch((err) => {
                logger.error("Fail sign up", err)
            })
    })

    expressApp.get(`${pathPrefix}/session`, (req, res) => {
        const session = {
            maxAge: sessionMaxAge,
            revalidateAge: sessionRevalidateAge,
            user: null
        }

        if (req.user) {
            session.user = req.user
        }

        return res.json(session)
    })

    // Prepare each provider's setting
    providers.forEach(({ name, strategy, strategyOptions, getProfile }) => {
        strategyOptions.callbackURL = `${pathPrefix}/oauth/${name.toLowerCase()}/callback`
        strategyOptions.passReqToCallback = true

        passport.use(
            name,
            new strategy(strategyOptions, async function(
                req,
                token: string,
                tokenSecret: string,
                profileFromRemote: any,
                done: (err: any, user?: User) => void
            ) {
                const profile = getProfile(profileFromRemote)
                const newUser = {
                    provider: name,
                    userName: profile.name,
                    displayName: profile.displayName,
                    email: profile.email,
                    imageUrl: profile.imageUrl,
                }

                const oauth: OAuthUser = {
                    id: profile.id,
                    accessToken: token,
                    refreshToken: tokenSecret,
                    providerName: profile.name,
                }
                const linkedUser = await store.findProviderLinkedUser(profile)

                // It means the user has already signed in
                if (req.user) {
                    if (linkedUser) {
                        // This code handles the user is already linked to a Provider
                        let user = req.user
                        if (req.user.id === linkedUser.id) {
                            try {
                                user = await store.updateOAuthUser(oauth)
                            } catch (err) {
                                logger.info("Failed to update oauth user", err)
                            }
                        }
                        return done(null, user)
                    }
                }

                // Below code is for not sign in users
                if (linkedUser) {
                    // Just update tokens and id
                    const user = await store.updateOAuthUser(oauth)
                    return done(null, user)
                }

                const existingUser = await store.findByEmail(newUser.email)

                // Linked existing account to oauth
                if (existingUser) {
                    await store.linkedExistingUser(existingUser.id, oauth)
                    return done(null, existingUser)
                }

                // SignUp
                const user = await store.createOAuthUser(newUser, oauth)
                done(null, user)
            })
        )
    })

    providers.forEach(({ name, options }) => {
        // Route to start sign in
        expressApp.get(`${pathPrefix}/oauth/${name.toLowerCase()}`, passport.authenticate(name, options))

        // Route to call back to after signing in
        expressApp.get(
            `${pathPrefix}/oauth/${name.toLowerCase()}/callback`,
            passport.authenticate(name, {
                successRedirect: `/`,
                failureRedirect: `/signin`,
            })
        )
    })

    return new Promise((resolve, reject) => {
        expressApp.all("*", (req, res) => {
            let nextRequestHandler = nextApp.getRequestHandler()
            return nextRequestHandler(req, res)
        })

        return expressApp.listen(port, err => {
            if (err) reject(err)
            return resolve()
        })
    })
}
