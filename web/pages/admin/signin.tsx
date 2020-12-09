import React from "react"

import { csrfToken } from "next-auth/client"
import Layout from "../../components/Layout"
import { Typography, TextField, Button, makeStyles } from "@material-ui/core"

export default function SignIn({ csrfToken }) {
    return (
        <Layout title="Login">
            <Typography component="h1" variant="h5">
                Sign in
            </Typography>
            <form method="post" action="/api/auth/callback/credentials">
                <input name="csrfToken" type="hidden" defaultValue={csrfToken} />
                <TextField
                    variant="outlined"
                    margin="normal"
                    required
                    fullWidth
                    id="username"
                    label="Admin Username"
                    name="username"
                    autoComplete="username"
                    autoFocus
                />
                <TextField
                    variant="outlined"
                    margin="normal"
                    required
                    fullWidth
                    name="password"
                    label="Password"
                    type="password"
                    id="password"
                    autoComplete="current-password"
                />
                <Button type="submit" fullWidth variant="contained" color="primary">
                    Sign in as Admin
                </Button>
            </form>

            {/* <div className={classes.signinButtons}>
                <GoogleButton
                    onClick={() => {
                      signIn(providerOptions["google"].id)
                    }}
                />
            </div> */}
        </Layout>
    )
}

SignIn.getInitialProps = async context => {
    return {
        csrfToken: await csrfToken(context),
    }
}
