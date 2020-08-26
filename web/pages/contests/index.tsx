import { useQuery } from "graphql-hooks"
import moment from "moment"
import Link from "next/link"
import React from "react"
import Layout from "../../components/Layout"
import { Contest } from "../../graphql/types"
import { Paper, TableContainer, TableHead, TableRow, TableCell, Table, makeStyles, TableBody } from "@material-ui/core"

type Props = {
    session: any
}

const contests = `
query getContests($limit: Int!) {
    contests(limit: $limit) {
        title,
        slug,
        startTimestamp
    }
}
`

const defaultProps = {
    bgcolor: "background.paper",
    m: 1,
    border: 1,
    style: { width: "5rem", height: "5rem" },
}

const useStyles = makeStyles({
    table: {
        minWidth: 650,
    },
    link: {
        textDecoration: "none",
    },
})

const problems: React.FunctionComponent<Props> = ({ session }: Props) => {
    const { data } = useQuery<{ contests: Contest[] }>(contests, {
        variables: { limit: 50 },
    })

    const classes = useStyles()
    return (
        <Layout session={session}>
            <TableContainer>
                <Table className={classes.table} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>ID</TableCell>
                            <TableCell>Title</TableCell>
                            <TableCell>Start time</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {!data ? (
                            <div>Loading</div>
                        ) : (
                            data.contests.map((c, i) => {
                                const linkUrl = `/contests/${c.slug}`
                                return (
                                    <TableRow key={c.slug}>
                                        <TableCell>{i + 1}</TableCell>
                                        <TableCell>
                                            <Link href={linkUrl}>
                                                <a className={classes.link}>{c.title}</a>
                                            </Link>
                                        </TableCell>
                                        <TableCell>{moment(c.startTimestamp).format("YYYY-MM-DD hh:ss")}</TableCell>
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
