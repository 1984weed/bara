import React from "react"
import { Contest } from "../../../graphql/types"
import { NextPage } from "next"
import Layout from "../../../components/Layout"
import Link from "next/link"
// import { NextPageContextWithGraphql } from "../../../lib/with-graphql-client"
import { Container, Typography, Box, TableContainer, Table, TableBody, TableHead, TableRow, TableCell } from "@material-ui/core"
import { makeStyles } from "@material-ui/core/styles"
import { grey } from "@material-ui/core/colors"
import CheckIcon from "@material-ui/icons/Check"

type Props = {
    session: any
    contest: Contest
}

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

const ContestComponent: NextPage<Props> = ({ session, contest }: Props) => {
    const classes = useStyles()

    return (
        <Layout session={session}>
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

// ContestComponent.getInitialProps = async ({ query, client }: NextPageContextWithGraphql) => {
//     const result = await client.request(
//         {
//             query: getContestProblemsQuery,
//             variables: { slug: query["contest-slug"] },
//         },
//         {}
//     )

//     const { contest } = result.data as {
//         contest: Contest
//     }

//     return Promise.resolve({
//         session: "",
//         contest,
//     })
// }

export default ContestComponent
