import { DBConfig, User, ProviderProfile, OAuthUser } from "./model"

import { Pool } from "pg"
import { hash } from "bcrypt"

const saltRounds = 10
export interface AuthStore {
    findByID(id: number): Promise<User | null>
    findByEmail(email: string): Promise<User | null>
    createUser(email: string, password: string): Promise<number>
    createOAuthUser(user: User, oauth: OAuthUser): Promise<User>
    findProviderLinkedUser(profileOauth: Partial<ProviderProfile>): Promise<User | null>
    updateOAuthUser(oauth: OAuthUser): Promise<User>
    linkedExistingUser(userId: number, oauth: OAuthUser): Promise<User>
}

export function createDBStore(config: DBConfig): AuthStore {
    return new AuthDBStore(new Pool(config))
}

class AuthDBStore implements AuthStore {
    constructor(private store: Pool) {}

    async findByID(id: number): Promise<User | null> {
        const res = await this.store.query({
            text: `SELECT id, user_name, email FROM users WHERE id = $1`,
            values: [id],
        })
        if (res.rows.length === 0) {
            return null
        }
        return res.rows[0]
    }

    async findByEmail(email: string): Promise<User | null> {
        const res = await this.store.query({
            text: `SELECT * FROM users WHERE email = $1`,
            values: [email],
        })
        if (res.rows.length === 0) {
            return null
        }
        return res.rows[0]
    }

    async createUser(email: string, password: string): Promise<null> {
        const hashPass = await new Promise((resolve, reject) => {
            hash(password, saltRounds, function(err, hash) {
                resolve(hash)
            })
        })
        const res = await this.store.query({
            text: `INSERT INTO users(user_name, password, email, updated_at, created_at) VALUES ($1, $2, $3, current_timestamp, current_timestamp) RETURNING id`,
            values: [email, hashPass, email],
        })
        const { id } = res.rows[0]
        return id
    }

    async findOrCreateUser(user: User): Promise<number> {
        const u = await this.findByEmail(user.email)

        if (u) {
            return u.id
        }

        const res = await this.store.query({
            text: `INSERT INTO users(user_name, display_name, email, provider, image, updated_at, created_at) VALUES ($1, $2, $3 ,$4, $5, current_timestamp, current_timestamp) RETURNING id`,
            values: [user.userName, user.displayName, user.email, user.provider, user.imageUrl],
        })

        return res.id
    }

    async createOAuthUser(user: User, oauth: OAuthUser): Promise<User> {
        let resId: number;
        try {
            await this.store.query("BEGIN")

            const resUsers = await this.store.query({
                text: `INSERT INTO users(user_name, display_name, email, provider, image, updated_at, created_at) VALUES ($1, $2, $3 ,$4, $5, current_timestamp, current_timestamp) RETURNING id`,
                values: [user.userName, user.displayName, user.email, user.provider, user.imageUrl],
            })

            resId = resUsers.rows[0].id

            await this.store.query({
                text: `INSERT INTO oauth_users(user_id, access_token, refresh_token, provider, provider_id) VALUES($1, $2, $3, $4, $5)`,
                values: [resId,  oauth.accessToken, oauth.refreshToken, oauth.providerName, oauth.id],
            })

            await this.store.query("COMMIT")
        } catch(err) {
            await this.store.query("ROLLBACK")
            throw err
        }

        return await this.findByID(resId)
    }

    async findProviderLinkedUser(profileOauth: Partial<ProviderProfile>): Promise<User | null> {
        const res = await this.store.query({
            text: `SELECT users.* FROM users, oauth_users WHERE users.id = oauth_users.user_id AND oauth_users.provider = $1 AND oauth_users.provider_id = $2`,
            values: [profileOauth.name, profileOauth.id],
        })

        if (res.rows.length === 0) {
            return null
        }
        return res.rows[0]
    }

    async updateOAuthUser(oauth: OAuthUser): Promise<User> {
        await this.store.query({
            text: `UPDATE oauth_users SET provider_id=$1, access_token=$2, refresh_token=$3 WHERE provider = $4 AND provider_id = $5`,
            values: [oauth.id, oauth.accessToken, oauth.refreshToken, oauth.providerName, oauth.id],
        })

        return await this.findProviderLinkedUser({name: oauth.providerName, id: oauth.id})
    }

    async linkedExistingUser(userId: number, oauth: OAuthUser): Promise<User> {
        await this.store.query({
            text: `INSERT INTO oauth_users(user_id, access_token, refresh_token, provider, provider_id) VALUES($1, $2, $3, $4, $5)`,
            values: [userId,  oauth.accessToken, oauth.refreshToken, oauth.providerName, oauth.id],
        })

        return await this.findProviderLinkedUser({name: oauth.providerName, id: oauth.id})
    }
}
