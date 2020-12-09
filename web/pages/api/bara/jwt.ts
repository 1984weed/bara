import jwt from "next-auth/jwt"
import * as jwtgen from "jsonwebtoken"

const secret = process.env.SECRET || "secret"

export default async (req, res) => {
    const token = await jwt.getToken({ req, secret })
    res.send(JSON.stringify({
        accessToken: getAccessToken(token)
    }, null, 2))
}

const jwtSecret = process.env.JWT_SECRET || "secret"
const jwtOptions: jwt.SignOptions = {
    algorithm: "HS256",
}
function getAccessToken(user) {
    const jwtPayload = {
        sub: user.sub,
        role: "admin",
    }

    return jwtgen.sign(jwtPayload, jwtSecret, jwtOptions)
}