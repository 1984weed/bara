import Button from "@material-ui/core/Button"
import CssBaseline from "@material-ui/core/CssBaseline"
import { makeStyles } from "@material-ui/core/styles"
import Typography from "@material-ui/core/Typography"
import { useMutation } from "graphql-hooks"
import React from "react"
import { useForm } from "react-hook-form"
import { ContestForm } from "../../../components/contests/ContestForm"
import Layout from "../../../components/Layout"
import { Problem } from "../../../graphql/types"
// import { NextPageContextWithGraphql } from "../../../lib/with-graphql-client"

export const allProblemsQuery = `
query problems($limit: Int) {
  problems(limit: $limit) {
    id,
    title,
    slug,
    description,
  }
}
`
export const createContest = `
mutation createContest($title: String!, $slug: String!, $startTimestamp: String!, $duration: String!, $problemIDs: [ID!]!) {
  createContest(newContest: {title: $title, slug: $slug, startTimestamp: $startTimestamp, duration: $duration, problemIDs: $problemIDs}) {
    slug,
    title
  }
}
`

const useStyles = makeStyles(theme => ({
    paper: {
        marginTop: theme.spacing(8),
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
    },
    form: {
        width: "100%", // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
}))

const NewContest = ({ session, problems }) => {
    const form = useForm({
        defaultValues: {
            title: "",
            slug: "",
            problemIDs: [],
            startTimestamp: new Date().toISOString()
        },
    })
    const classes = useStyles()
    const [createPost] = useMutation(createContest)

    const onSubmit = async formData => {
        formData.duration = "90"
        const { data, error } = await createPost({
            variables: formData,
        })

        if (error != null) {
            return
        }
    }

    return (
        <Layout title="Admin problems" session={session}>
            <Typography component="h1" variant="h5">
                New contest
            </Typography>
            <CssBaseline />
            <div className={classes.paper}>
                <form className={classes.form} onSubmit={form.handleSubmit(onSubmit)}>
                    <ContestForm 
                        {...form}
                        problems={problems}
                    />
                    <Button type="submit" fullWidth variant="contained" color="primary" className={classes.submit}>
                        Create contest
                    </Button>
                </form>
            </div>
        </Layout>
    )
}

// NewContest.getInitialProps = async ({ query, client }: NextPageContextWithGraphql) => {
//     const result = await client.request(
//         {
//             query: allProblemsQuery
//         }
//     )

//     if (result.data == null) {
//         return Promise.resolve({
//             contest: null,
//             problems: [],
//             session: "",
//             pathname: "",
//         })
//     }

//     const { problems } = result.data as {
//         problems: Problem[]
//     }

//     return Promise.resolve({
//         problems,
//         session: "",
//         pathname: "",
//     })
// }

export default NewContest
