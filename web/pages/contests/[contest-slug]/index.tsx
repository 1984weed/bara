import { Box, Container, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography } from "@material-ui/core"
import { makeStyles } from "@material-ui/core/styles"
import CheckIcon from "@material-ui/icons/Check"
import { useQuery } from "graphql-hooks"
import { NextPage } from "next"
import Link from "next/link"
import { useRouter } from "next/router"
import React from "react"
import Layout from "../../../components/Layout"
import { Contest } from "../../../graphql/types"

type Props = {}

const getContestProblemsQuery = `
query getContestProblems($slug: String!) {
    contest(slug: $slug) {
        title,
        startTimestamp,
        slug,
        problems {
            title,
            slug, 
            userResult {
                done
            }
        },
        duration
    }
}
`

const useStyles = makeStyles(theme => ({
    problemDoneIcon: {
        marginRight: "5px",
    },
}))

const ContestComponent: NextPage<Props> = () => {
    const classes = useStyles()
    const router = useRouter()

    const { data } = useQuery<{ contest: Contest }>(getContestProblemsQuery, {
        variables: { slug: router.query["contest-slug"] },
    })

    if(!data) {
        return <>Loading</>
    }

    const { contest } = data

    return (
        <Layout>
            <Container maxWidth="md" component="main">
                <Box>
                    <Typography variant="h2" color="textPrimary">
                        {contest.title}
                    </Typography>
                    <Typography variant="subtitle1">Welcome to the {contest.title} contest</Typography>
                </Box>
                <Box>
                    <TableContainer>
                        <Table>
                            <TableHead>
                                <TableRow>
                                    <TableCell>ID</TableCell>
                                    <TableCell>Problem Titles</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {contest.problems.map((p, i) => {
                                    const linkUrl = `/contests/${contest.slug}/problems/${p.slug}`
                                    return (
                                        <TableRow key={p.slug}>
                                            <TableCell>{i + 1}</TableCell>
                                            <TableCell>
                                                <Box display="flex" alignItems="center" justifyContent="space-between">
                                                    <Link href={linkUrl}>
                                                        <a className="problem-link">{p.title}</a>
                                                    </Link>
                                                    {p.userResult?.done ? <CheckIcon className={classes.problemDoneIcon} /> : <></>}
                                                </Box>
                                            </TableCell>
                                        </TableRow>
                                    )
                                })}
                            </TableBody>
                        </Table>
                    </TableContainer>
                </Box>
            </Container>
        </Layout>
    )
}

export default ContestComponent
