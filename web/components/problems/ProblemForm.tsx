import { Button, Checkbox, FormHelperText, Grid, NativeSelect } from "@material-ui/core"
import Box from "@material-ui/core/Box"
import InputLabel from "@material-ui/core/InputLabel"
import { makeStyles } from "@material-ui/core/styles"
import TextField from "@material-ui/core/TextField"
import Chance from "chance"
import React from "react"
import { Controller } from "react-hook-form"
const chance = new Chance()

const typesArray = ["", "string[]", "string", "int[]", "int", "double[]", "double", "boolean"]

type Props = {
    state?: any
    setState?: any
    onClickSubmit: (data: any) => void
    register: any
    control: any
    handleSubmit: any
    watch: any
    errors: any
    setValue: any
    getValues: any
}

const useStyles = makeStyles(theme => ({
    form: {
        width: "100%", // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
}))

export default ({ state = null, setState = null, register, control, handleSubmit, onClickSubmit, watch, errors, setValue, getValues }: Props) => {
    const classes = useStyles()
    const watchArgsNum = watch("argumentNum")
    const testCaseNum = watch("testCaseNum")
    const argumentTypes = watch("argumentTypes")

    const require = false

    return (
        <form onSubmit={handleSubmit(onClickSubmit)}>
            <Grid container spacing={3}>
                <Grid item xs={12}>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        id="title"
                        label="Problem Title"
                        name="title"
                        autoComplete="Title"
                        fullWidth
                        autoFocus
                        inputRef={register({ required: require })}
                    />
                    {errors.title && <span>This field is required</span>}
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        id="slug"
                        label="Problem URL"
                        name="slug"
                        autoComplete="Slug"
                        fullWidth
                        inputRef={register()}
                    />
                    {errors.title && <span>This field is required</span>}
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        variant="outlined"
                        label="Description"
                        autoComplete="Description"
                        margin="normal"
                        multiline
                        fullWidth
                        size="medium"
                        name="description"
                        inputRef={register({ required: require })}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        variant="outlined"
                        label="Function name"
                        autoComplete="Function name"
                        fullWidth
                        name="functionName"
                        inputRef={register({ required: require })}
                    />
                    <FormHelperText>{errors.functionName && errors.functionName.message}</FormHelperText>
                </Grid>
                <Grid item xs={12}>
                    <Grid item xs={12}>
                        <Controller
                            name="outputType"
                            as={
                                <NativeSelect variant="outlined" fullWidth>
                                    {typesArray.map(t => (
                                        <option key={t} value={t}>
                                            {t}
                                        </option>
                                    ))}
                                </NativeSelect>
                            }
                            inputRef={register({ required: true })}
                            control={control}
                        />
                    </Grid>
                    <FormHelperText>{errors.outputType && errors.outputType.message}</FormHelperText>
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        label="Input Args Count"
                        type="number"
                        name="argumentNum"
                        inputProps={{ min: "0", max: "10", step: "1" }}
                        InputLabelProps={{
                            shrink: true,
                        }}
                        fullWidth
                        inputRef={register({ required: require })}
                        variant="outlined"
                    />
                </Grid>
                {Array.from({ length: watchArgsNum }).map((_, i) => (
                    <Grid item key={i} xs={12}>
                        <TextField
                            label={`Input Arg Name ${i + 1}:`}
                            name={`argumentNames[${i}]`}
                            key={`${i}-argname`}
                            fullWidth
                            inputRef={register({ required: require })}
                        />
                        <InputLabel key={i} id={`input-type-label-${i + 1}`}>
                            Input arg's type: {i + 1}
                        </InputLabel>
                        <Controller
                            key={`${i}-argtype`}
                            name={`argumentTypes[${i}]`}
                            as={
                                <NativeSelect variant="outlined" fullWidth>
                                    {typesArray.map(t => (
                                        <option key={t} value={t}>
                                            {t}
                                        </option>
                                    ))}
                                </NativeSelect>
                            }
                            control={control}
                        />
                        <FormHelperText>{errors.outputType && errors.outputType.message}</FormHelperText>
                    </Grid>
                ))}
                <Grid item xs={12}>
                    <TextField
                        label="Test case Count"
                        type="number"
                        name="testCaseNum"
                        fullWidth
                        inputProps={{ min: "0", max: "10", step: "1" }}
                        inputRef={register({ required: require })}
                        variant="outlined"
                    />
                    {Array.from({ length: parseInt(testCaseNum) }).map((_, testCaseIndex) => (
                        <Grid item xs={12} key={testCaseIndex}>
                            {Array.from({ length: watchArgsNum }).map((_, i) => (
                                <Grid item xs={12} key={`${testCaseIndex}-${i}-con`}>
                                    <TextField
                                        key={`${i}-${testCaseIndex}`}
                                        label={`Input Test Case ${i + 1}:`}
                                        fullWidth
                                        name={`inputTestCase-${testCaseIndex}[${i}]`}
                                        InputLabelProps={{ shrink: true }}
                                        inputRef={register({ required: require })}
                                        variant="outlined"
                                    />
                                    <Box>
                                        {argumentTypes[i] === "int[]" && (
                                            <Box>
                                                <Grid item>
                                                    <TextField
                                                        key={i}
                                                        fullWidth
                                                        name={`testCaseGen-${testCaseIndex}[${i}].count`}
                                                        label="Sample item count"
                                                        inputProps={{ min: "1", max: "1000", step: "1" }}
                                                        type="number"
                                                        inputRef={register}
                                                    />
                                                </Grid>
                                                <Grid item>
                                                    <TextField
                                                        key={i}
                                                        fullWidth
                                                        type="number"
                                                        name={`testCaseGen-${testCaseIndex}[${i}].min`}
                                                        inputProps={{ min: "-100", max: "1000", step: "1" }}
                                                        label="Sample item min"
                                                        inputRef={register}
                                                    />
                                                </Grid>
                                                <Grid item>
                                                    <TextField
                                                        key={i}
                                                        fullWidth
                                                        type="number"
                                                        name={`testCaseGen-${testCaseIndex}[${i}].max`}
                                                        inputProps={{ min: "-100", max: "1000", step: "1" }}
                                                        label="Sample item max"
                                                        inputRef={register}
                                                    />
                                                </Grid>
                                                <Grid item>
                                                    <Controller
                                                        as={<Checkbox />}
                                                        name={`testCaseGen-${testCaseIndex}[${i}].isUnique`}
                                                        value="test"
                                                        control={control}
                                                        defaultValue={false}
                                                        inputRef={register}
                                                    />
                                                    <span>Is unique?</span>
                                                </Grid>
                                                <Button
                                                    type="button"
                                                    color="primary"
                                                    variant="contained"
                                                    onClick={e => {
                                                        const inputTestCaseKey = `inputTestCase-${testCaseIndex}`
                                                        const testCaseGenKey = `testCaseGen-${testCaseIndex}`

                                                        const allState = getValues({ nest: true })
                                                        const inputTestCase = allState[inputTestCaseKey] || []
                                                        const targetCaseGen = allState[testCaseGenKey] ? allState[testCaseGenKey][i] : {}

                                                        const sampleNum = parseInt(targetCaseGen.count) || 1
                                                        const min = parseInt(targetCaseGen.min) || 1
                                                        const max = parseInt(targetCaseGen.max) || 1
                                                        const isUnique = !!targetCaseGen.isUnique
                                                        const sample = []
                                                        for (let i = 0; i < sampleNum; i++) {
                                                            sample.push(
                                                                chance.integer({
                                                                    min,
                                                                    max,
                                                                })
                                                            )
                                                        }
                                                        inputTestCase[i] = `[${sample.join(",")}]`
                                                        setValue(`inputTestCase-${testCaseIndex}`, inputTestCase)
                                                    }}
                                                >
                                                    Create sample input
                                                </Button>
                                            </Box>
                                        )}
                                    </Box>
                                </Grid>
                            ))}
                            <Box>
                                <TextField
                                    label="Output"
                                    variant="outlined"
                                    fullWidth
                                    name={`outTestCase[${testCaseIndex}]`}
                                    inputRef={register({ required: require })}
                                />
                            </Box>
                        </Grid>
                    ))}
                </Grid>
                <Grid item xs={12}>
                    <Button type="submit" fullWidth variant="contained" color="primary" className={classes.submit}>
                        Submit
                    </Button>
                </Grid>
            </Grid>
        </form>
    )
}
