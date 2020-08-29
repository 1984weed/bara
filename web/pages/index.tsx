import { Table, TableContainer, TableHead, TableRow, TableCell, TableBody } from "@material-ui/core"
import { useQuery } from "graphql-hooks"
import Link from "next/link"
import React from "react"
import Layout from "../components/Layout"
import { Problem } from "../graphql/types"
import { useSession } from "../lib/session"

type Props = {
    session: any
}

const problem = `
query getProblems($limit: Int!) {
    problems(limit: $limit) {
        id,
        title,
        slug,
        description,
    }
}
`

const problems: React.FunctionComponent<Props> = ({ session }: Props) => {
    const { data } = useQuery<{ problems: Problem[] }>(problem, {
        variables: { limit: 50 },
    })

    const sessions = useSession()
    console.log("sessions", sessions)
    return (
        <Layout session={session}>
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
                            data.problems.map((p, i) => {
                                const linkUrl = `/problems/${p.slug}`
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
            </TableContainer>
        </Layout>
    )
}

export default problems
