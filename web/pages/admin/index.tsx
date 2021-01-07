import { useQuery } from "graphql-hooks"
import React from "react"
import Layout from "../../components/Layout"
import { Problem, Contest } from "../../graphql/types"
import Link from "next/link"
import { Session } from "../../types/Session"
import { TableContainer, Table, TableHead, TableCell, TableRow, TableBody, Box } from "@material-ui/core"

type Props = {
    session: Session
}

const problem = `
query getProblems($limit: Int!) {
    problems(limit: $limit) {
        id,
        title,
        slug,
        description,
    }
    contests(limit: $limit) {
      id,
      title,
      startTimestamp,
      slug
    }
}
`

const problems: React.FunctionComponent<Props> = ({ session }: Props) => {
    const { data } = useQuery<{ problems: Problem[]; contests: Contest[] }>(problem, {
        variables: { limit: 50 },
    })
    return (
        <Layout session={session}>
            <Box>
                Contests
                <Box>
                    <Link href="/admin/contests/new">Create a new contest</Link>
                </Box>
                <TableContainer>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell scope="col">ID</TableCell>
                                <TableCell scope="col">Title</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {!data ? (
                                <div>Loading</div>
                            ) : (
                                data.contests.map((c, i) => {
                                    const linkUrl = `/admin/contests/${c.slug}/edit`
                                    return (
                                        <TableRow>
                                            <TableCell>{c.id}</TableCell>
                                            <TableCell scope="row">
                                                <Link href={linkUrl}>
                                                    <a>{c.title}</a>
                                                </Link>
                                            </TableCell>
                                        </TableRow>
                                    )
                                })
                            )}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Box>
            <Box>
                Problems
                <Box>
                    <Link href="/admin/problems/new">Create a new problem</Link>
                </Box>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell scope="col">ID</TableCell>
                            <TableCell scope="col">Title</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {!data ? (
                            <div>Loading</div>
                        ) : (
                            data.problems.map((p, i) => {
                                const linkUrl = `/admin/problems/edit/${p.slug}`
                                return (
                                    <TableRow>
                                        <TableCell>{p.id}</TableCell>
                                        <TableCell scope="row">
                                            <Link href={linkUrl}>
                                                <a>{p.title}</a>
                                            </Link>
                                        </TableCell>
                                    </TableRow>
                                )
                            })
                        )}
                    </TableBody>
                </Table>
            </Box>
        </Layout>
    )
}

export default problems
