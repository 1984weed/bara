import React from "react"

const footers = [
    {
        title: "Company",
        description: ["Team", "History", "Contact us", "Locations"],
    },
    {
        title: "Features",
        description: ["Cool stuff", "Random feature", "Team feature", "Developer stuff", "Another one"],
    },
    {
        title: "Resources",
        description: ["Resource", "Resource name", "Another resource", "Final resource"],
    },
    {
        title: "Legal",
        description: ["Privacy policy", "Terms of use"],
    },
]

import Icon from "./icons/WhiteIcon"
import { makeStyles, Container, Grid, Typography, Box, Link } from "@material-ui/core"
import LinkNext from "next/link"
import { MaterialLink } from "./atoms/MaterialLInk"

const useStyles = makeStyles(theme => ({
    footer: {
        borderTop: `1px solid ${theme.palette.divider}`,
        marginTop: theme.spacing(8),
        paddingTop: theme.spacing(3),
        paddingBottom: theme.spacing(3),
        [theme.breakpoints.up("sm")]: {
            paddingTop: theme.spacing(6),
            paddingBottom: theme.spacing(6),
        },
        backgroundColor: theme.palette.primary.light,
    },
    footerText: {
      color: theme.palette.primary.contrastText,
      verticalAlign: "middle"
    },
    copyRight: {
      fontWeight: "bold",
      color: theme.palette.primary.contrastText
    }
}))

export default () => {
    const classes = useStyles()

    return (
        <Container maxWidth="md" component="footer" className={classes.footer}>
            <Grid container spacing={4} direction="column" justify="space-evenly">
                <Grid item>
                    <Box display="flex" justifyContent="center" alignContent="center">
                      <Icon style={{ "--color_fill": "#fff" }} />
                      <Typography variant="h3" className={classes.footerText} component="span">
                          Baracode
                      </Typography>
                    </Box>
                </Grid>
                <Grid item>
                    <Box mt={2}>
                        <Copyright />
                    </Box>
                </Grid>
            </Grid>
        </Container>
    )
}

const Copyright = () => {
    const classes = useStyles()

    return (
        <Typography variant="body2" color="textSecondary" align="center" className={classes.copyRight}>
            {"Copyright © "}
            Baracode
            {new Date().getFullYear()}
            {"."}
        </Typography>
    )
}