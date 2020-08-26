import React from "react"
import SwipeableViews from "react-swipeable-views"
import { Submission } from "../../graphql/types"
import { Table, TableRow, TableCell, TableContainer, TableHead, TableBody, Paper, Tabs, Tab, Box, makeStyles } from "@material-ui/core"
import Markdown from "../Markdown"

type Props = {
    title: string
    description: string
    submissionList: Array<Submission>
    clickSubmission: () => void
}

const ListItem = ({ submissionList }: { submissionList: Array<Submission> }) => {
    const listItems = submissionList.map((s, i) => {
        const dateStr = s.timestamp.slice(0, 10)
        const time = s.timestamp.slice(11, 19)
        return (
            <TableRow key={i} style={{ background: i % 2 === 0 ? "white" : "#F2F2F2" }}>
                <TableCell>
                    {dateStr} {time}
                </TableCell>
                <TableCell>{s.statusSlug}</TableCell>
                <TableCell>{s.runtimeMS}ms</TableCell>
            </TableRow>
        )
    })

    return (
        <TableContainer>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell scope="col">Time Submitted</TableCell>
                        <TableCell scope="col">Status</TableCell>
                        <TableCell scope="col">Runtime</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>{listItems}</TableBody>
            </Table>
        </TableContainer>
    )
}

const SideArea: React.FunctionComponent<Props> = ({ title, description, submissionList, clickSubmission }: Props) => {
    const [value, setValue] = React.useState(0)
    return (
        <Box>
            <Box borderBottom={1} borderColor="primary.main">
                <Tabs
                    value={value}
                    indicatorColor="primary"
                    textColor="primary"
                    onChange={(_: React.ChangeEvent<{}>, newValue: number) => {
                        setValue(newValue)
                    }}
                    centered
                >
                    <Tab label="Description" />
                    <Tab label="Submissions" />
                </Tabs>
            </Box>
            <SwipeableViews axis="x" index={value} onChangeIndex={(value: number) => {
                setValue(value)
            }}>
                <TabPanel value={value} index={0} >
                    <Markdown>
                        {description}
                    </Markdown>
                </TabPanel>
                <TabPanel value={value} index={1} >
                    <ListItem submissionList={submissionList} />
                </TabPanel>
            </SwipeableViews>
        </Box>
    )
}
type TabPanelProps = {
    children?: React.ReactNode
    dir?: string
    index: any
    value: any
}

const TabPanel = ({ children, value, index, ...other }: TabPanelProps) => {
    return (
        <div role="tabpanel" hidden={value !== index} id={`full-width-tabpanel-${index}`} aria-labelledby={`full-width-tab-${index}`} {...other}>
            {value === index && <Box p={3}>{children}</Box>}
        </div>
    )
}

export default SideArea
