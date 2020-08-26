import Button from "@material-ui/core/Button"
import { makeStyles } from "@material-ui/core/styles"
import TextField from "@material-ui/core/TextField"
import Typography from "@material-ui/core/Typography"
import React from "react"
import GithubSignInButton from "../components/atoms/GithubSignInButton"
import TwitterSignInButton from "../components/atoms/TwitterSignInButton"
import Layout from "../components/Layout"


const useStyles = makeStyles(theme => ({
    paper: {
        marginTop: theme.spacing(8),
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: "100%", // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
    authProvidersCtr: {
        marginTop: "5px",
        height: "100px",
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-around"
        
    },
}))

export default ({ session }) => {
    const classes = useStyles()
    return (
        <Layout title="Login" session={session}>
            <Typography component="h1" variant="h5">
                Sign in
            </Typography>
            <form className={classes.form} noValidate action="/auth/signin" method="post">
                <TextField
                    variant="outlined"
                    margin="normal"
                    required
                    fullWidth
                    id="email"
                    label="Email Address"
                    name="email"
                    autoComplete="email"
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
                <Button type="submit" fullWidth variant="contained" color="primary" className={classes.submit}>
                    Sign In
                </Button>
            </form>
            <div className={classes.authProvidersCtr}>
                <TwitterSignInButton label="Sign in with Twitter" />
                <GithubSignInButton label="Sign in with Github" />
            </div>
        </Layout>
    )
}
