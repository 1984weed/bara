import { Box, Checkbox, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, FormControlLabel } from "@material-ui/core"
import Button from "@material-ui/core/Button"
import { makeStyles } from "@material-ui/core/styles"
import TextField from "@material-ui/core/TextField"
import React, { useState } from "react"
import { Controller } from "react-hook-form"
import { Problem } from "../../graphql/types"
import { ContestDate } from "./ContestDate"

type Props = {
    control: any
    getValues: any
    register: any
    errors: any
    setValue: any
    watch: any
    problems: Problem[]
}

const useStyles = makeStyles(theme => ({
    button: {
        margin: theme.spacing(3),
    },
}))

export const ContestForm = ({ problems, control, getValues, setValue, register, errors, watch }: Props) => {
    const [openProblemCheckBox, setOpenProblemCheckBox] = useState(false)
    const handleClickOpen = () => {
        setOpenProblemCheckBox(true)
    }

    const handleClose = () => {
        setOpenProblemCheckBox(false)
        const currentState = getValues({ nest: true })
        setValue("problemIDs", Array.from(new Set([...currentState.problemIDs])))
    }

    register("problemIDs")

    const handleOK = (checkedProblems: Set<number>) => {
        const currentState = getValues({ nest: true })
        setValue("problemIDs", Array.from(new Set([...currentState.problemIDs, ...Array.from(checkedProblems)])))
        setOpenProblemCheckBox(false)
    }

    const classes = useStyles()
    const problemIDs = watch("problemIDs") as number[]
    const startTimeStamp = watch("startTimestamp") as string

    return (
        <Box>
            <TextField
                variant="outlined"
                margin="normal"
                id="title"
                label="Contest Title"
                name="title"
                autoComplete="Title"
                autoFocus
                fullWidth
                inputRef={register({ required: true })}
            />
            {errors.title && <span>This field is required</span>}
            <TextField
                variant="outlined"
                margin="normal"
                id="slug"
                label="Contest Slug"
                name="slug"
                autoComplete="Contest Slug"
                autoFocus
                fullWidth
                inputRef={register({ required: true })}
            />
            {errors.slug && <span>This field is required</span>}
            <Button variant="outlined" color="primary" className={classes.button} onClick={handleClickOpen}>
                Select problems
            </Button>
            <ProblemDialog open={openProblemCheckBox} handleOK={handleOK} handleClose={handleClose} problems={problems} problemIDs={problemIDs} />
            {problemIDs.map((id) => (
                <Box>Selected ID: {id}</Box>
            ))}
            <Controller
                as={<ContestDate startTimeStamp={startTimeStamp} onChange={(timeStr: string) => setValue("startTimestamp", timeStr)} />}
                control={control}
                name="startTimestamp"
                inputRef={register}
            />
        </Box>
    )
}

const ProblemDialog = ({
    open,
    handleClose,
    handleOK,
    problems,
    problemIDs,
}: {
    open: boolean
    handleClose: () => void
    handleOK: (checkedIDs: Set<number>) => void
    problems: Problem[]
    problemIDs: number[]
}) => {
    const [problemIDsSet, setProblemIDsSet] = useState(new Set(problemIDs))

    return (
        <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
            <DialogTitle id="form-dialog-title">All Problems</DialogTitle>
            <DialogContent>
                <DialogContentText>Check your proper problems for the contest</DialogContentText>
                <Box display="flex" flexWrap="nowrap" flexDirection="column" style={{ height: "200px", overflowY: "scroll" }}>
                    {problems.map((p, i) => (
                        <FormControlLabel
                            control={<Checkbox checked={problemIDsSet.has(p.id)} name={`${p.id}`} color="primary" />}
                            label={`id: ${p.id} title: ${p.title}`}
                            name={`${p.id}`}
                            onChange={e => {
                                const target = e.target as any

                                if (target.checked) {
                                    problemIDsSet.add(p.id)
                                } else if (problemIDsSet.has(p.id)) {
                                    problemIDsSet.delete(p.id)
                                }
                                setProblemIDsSet(new Set(problemIDsSet))
                            }}
                        />
                    ))}
                </Box>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose} color="primary">
                    Cancel
                </Button>
                <Button
                    onClick={() => {
                        handleOK(problemIDsSet)
                    }}
                    color="primary"
                >
                    OK
                </Button>
            </DialogActions>
        </Dialog>
    )
}