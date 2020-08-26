import React from "react"
import ReactMarkdown from "markdown-to-jsx"
import { withStyles, Theme } from "@material-ui/core/styles"
import Typography from "@material-ui/core/Typography"
import Link from "@material-ui/core/Link"
import { Styles } from "@material-ui/core/styles/withStyles"

const styles: Styles<Theme, any, any> = theme => ({
    listItem: {
        marginTop: theme.spacing(1),
    },
    pre: {
        border: "1px solid #ccc",
    },
    code: {
        display: "block",
        overflowX: "auto",
        padding: "0.5em",
        color: "#333",
        background: "#f8f8f8",
    },
})

const options = {
    overrides: {
        h1: {
            component: Typography,
            props: {
                gutterBottom: true,
                variant: "h5",
            },
        },
        h2: { component: Typography, props: { gutterBottom: true, variant: "h6" } },
        h3: { component: Typography, props: { gutterBottom: true, variant: "subtitle1" } },
        h4: {
            component: Typography,
            props: { gutterBottom: true, variant: "caption", paragraph: true },
        },
        pre: {
            component: withStyles(styles)(({ classes, ...props }: any) => <pre {...props} className={classes.pre} />),
        },
        code: {
            component: withStyles(styles)(({ classes, ...props }: any) => <code {...props} className={classes.code} />),
        },
        p: { component: Typography, props: { paragraph: true } },
        a: { component: Link },
        li: {
            component: withStyles(styles)(({ classes, ...props }: any) => (
                <li className={classes.listItem}>
                    <Typography component="span" {...props} />
                </li>
            )),
        },
    },
}

export default function Markdown(props) {
    return <ReactMarkdown options={options} {...props} />
}
