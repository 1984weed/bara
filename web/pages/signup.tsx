import Layout from "../components/Layout"
import { Typography, TextField, Button, makeStyles } from "@material-ui/core"

const useStyles = makeStyles(theme => ({
    form: {
        width: "100%",
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
}))

export default ({ session }) => {
    const classes = useStyles()
    return (
        <Layout title="Login" session={session}>
            <Typography component="h1" variant="h5">
                Sign up
            </Typography>
            <form className={classes.form} noValidate action="/auth/signup" method="post">
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
                    Sign Up
                </Button>
            </form>
        </Layout>
    )
}
