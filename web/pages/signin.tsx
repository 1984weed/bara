import { makeStyles } from "@material-ui/core/styles"
import Typography from "@material-ui/core/Typography"
import { providers, signIn } from "next-auth/client"
import React from "react"
import GoogleButton from "react-google-button"
import Layout from "../components/Layout"

const useStyles = makeStyles(theme => ({
    signinButtons: {
        width: "100%", 
        marginTop: theme.spacing(1),
        display: "flex",
        justifyContent: "center",
    },
}))

const SignIn = ({ providerOptions }: { providerOptions: any[] }) => {
  const classes = useStyles();
    return (
        <Layout title="Login">
            <Typography component="h1" variant="h5">
                Sign in
            </Typography>
            <div className={classes.signinButtons}>
                <GoogleButton
                    onClick={() => {
                      signIn(providerOptions["google"].id)
                    }}
                />
            </div>
        </Layout>
    )
}

SignIn.getInitialProps = async context => {
    return {
        providerOptions: await providers(context),
    }
}

export default SignIn
