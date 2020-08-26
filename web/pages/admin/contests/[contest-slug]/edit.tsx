import { Box, Checkbox, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, FormControlLabel, Grid } from "@material-ui/core"
import Button from "@material-ui/core/Button"
import CssBaseline from "@material-ui/core/CssBaseline"
import { makeStyles } from "@material-ui/core/styles"
import TextField from "@material-ui/core/TextField"
import Typography from "@material-ui/core/Typography"
import { useMutation } from "graphql-hooks"
import React, { useState } from "react"
import { Controller, useForm } from "react-hook-form"
import { ContestDate } from "../../../../components/contests/ContestDate"
import Layout from "../../../../components/Layout"
import ErrorNotification from "../../../../components/notifications/ErrorNotification"
import SuccessNotification from "../../../../components/notifications/SuccessNotification"
import { Contest, Problem } from "../../../../graphql/types"
import { NextPageContextWithGraphql } from "../../../../lib/with-graphql-client"
import { Session } from "../../../../types/Session"
import { ContestForm } from "../../../../components/contests/ContestForm"

export const contestQuery = `
query contest($slug: String!, $limit: Int) {
  contest(slug: $slug) {
    id,
    slug,
    title,
    startTimestamp,
    duration,
    problems {
        id
    }
  }
  problems(limit: $limit) {
    id,
    title,
    slug,
    description,
  }
}
`

export const updateContestMutation = `
mutation updateContest($id: ID!, $slug: String!, $title: String!, $startTimestamp: String!, $duration: String, $problemIDs: [ID!]) {
  updateContest(contestID: $id, newContest: {title: $title, slug: $slug, startTimestamp: $startTimestamp, duration: $duration, problemIDs: $problemIDs}) {
    id
  }
}
`

type Props = {
    session: Session
    contest: Contest
    problems: Problem[]
}

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

const EditContest = ({ session, contest, problems }: Props) => {
    const [updateContest] = useMutation(updateContestMutation)

    const [submitError, setSubmitError] = useState(false)
    const [updated, setUpdated] = useState(false)
    const [openProblemCheckBox, setOpenProblemCheckBox] = useState(false)
    const handleClickOpen = () => {
        setOpenProblemCheckBox(true)
    }

    const handleClose = () => {
        setOpenProblemCheckBox(false)
    }

    const onSubmit = async formData => {
        const { data, error } = await updateContest({
            variables: { id: contest.id, ...formData },
        })

        if (error != null) {
            setSubmitError(true)
            return
        }
        setUpdated(true)
    }

    const form = useForm({
        defaultValues: {
            problemCount: contest.problems.length,
            title: contest.title,
            slug: contest.slug,
            problemIDs: contest.problems.map(a => a.id),
            startTimestamp: contest.startTimestamp,
            duration: contest.duration,
        },
    })

    const classes = useStyles()

    return (
        <Layout title="Admin problems" session={session}>
            <Typography component="h1" variant="h5">
                Edit
            </Typography>
            <CssBaseline />
            <Grid container spacing={3}>
                <ErrorNotification open={submitError} onClose={() => setSubmitError(false)} />
                <SuccessNotification open={updated} label="Success to update this problem" onClose={() => setUpdated(false)} />
                <form className={classes.form} onSubmit={form.handleSubmit(onSubmit)}>
                    <ContestForm 
                        {...form}
                        problems={problems}
                    />

                    <Button type="submit" fullWidth variant="contained" color="primary" className={classes.submit}>
                        Update contest
                    </Button>
                </form>
            </Grid>
        </Layout>
    )
}

EditContest.getInitialProps = async ({ query, client }: NextPageContextWithGraphql) => {
    const result = await client.request(
        {
            query: contestQuery,
            variables: { slug: query["contest-slug"] },
        },
        {}
    )

    if (result.data == null) {
        return Promise.resolve({
            contest: null,
            session: "",
            pathname: "",
        })
    }

    const { contest, problems } = result.data as {
        contest: Contest
        problems: Problem[]
    }

    return Promise.resolve({
        contest,
        problems,
        session: "",
        pathname: "",
    })
}

export default EditContest
