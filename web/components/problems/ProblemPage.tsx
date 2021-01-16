import { Box, Button, Checkbox, FormControlLabel, Grid, TextareaAutosize, Typography } from "@material-ui/core/"
import ReplayIcon from "@material-ui/icons/Replay"
import Router from "next/router"
import React, { ReactElement, useState } from "react"
import { useFormState } from "react-use-form-state"
import Editor from "../../components/Editor"
import SubmittedResult from "../../components/problems/SubmittedResult"
import { CodeLanguage } from "../../graphql/types"
import { useRememberState } from "../../hooks/useRememberState"
import { Status } from "../../types/Contraints"

type Props = {
    testRunResult: any
    submittedResult?: SubmitCodeResult | null
    targetCodeSnippet: any
    localTypedCodeKey: string
    defaultCode: string
    sampleTestCase: string
    title: string
    language: CodeLanguage
    contestSlug?: string
    onRunCode: (e: RunCodeType) => void
    onSubmitCode: (e: SubmitCodeType) => void
}

export type SubmitCodeResult = {
    status: string
    expected: string
    time: string
    input: string
    result: string
    stdout: string
}

export type RunCodeType = {
    typedCode: string
    language: string 
    testcase: string
}

export type SubmitCodeType = {
    contestSlug?: string
    typedCode: string
    language: string

}

export const EditorArea = ({
    testRunResult,
    submittedResult,
    targetCodeSnippet,
    localTypedCodeKey,
    defaultCode,
    sampleTestCase,
    title,
    language,
    contestSlug,
    onRunCode,
    onSubmitCode
}: Props) => {
    const [editor, setEditor] = useState(null)
    const [typedCode, setTypedCode] = useRememberState(localTypedCodeKey, defaultCode)
    const [testcase, setTestcase] = useState(sampleTestCase)
    const [checkedCustomTestcase, setCheckedCustomTestcase] = useState(false)
    const initialCode = typedCode === "" || typedCode === null ?  defaultCode : typedCode

    return (
        <Grid container direction="column" justify="center" alignItems="stretch" style={{ height: "100%" }}>
            <Grid item>
                <Typography variant="body1">JavaScript</Typography>
                <div title="Reset code" onClick={() => editor.changeDefaultValue(targetCodeSnippet.code)} style={{ cursor: "pointer" }}>
                    <ReplayIcon />
                </div>
            </Grid>
            <Grid item xs>
                <div style={{ display: "flex", flexDirection: "column", height: "100%", alignItems: "stretch" }}>
                    <Editor
                        ref={(editor: Editor) => {
                            setEditor(editor)
                        }}
                        onChange={(changeCode: string) => {
                            setTypedCode(changeCode)
                        }}
                        value={initialCode}
                    />
                    {(() => {
                        // the test run fails 
                        if (testRunResult.error) {
                            return displayCodeResultError()
                        }

                        // the test run data
                        if (testRunResult.data) {
                            return displayCodeResult(
                                title,
                                language,
                                contestSlug,
                                {
                                    status: testRunResult.data.testRunCode.result.status,
                                    expected: testRunResult.data.testRunCode.result.expected,
                                    input: testRunResult.data.testRunCode.result.input,
                                    result: testRunResult.data.testRunCode.result.result,
                                    stdout: testRunResult.data.testRunCode.stdout,
                                },
                                true
                            )
                        }

                        if (submittedResult) {
                            return displayCodeResult(
                                title,
                                language,
                                contestSlug,
                                {
                                    status: submittedResult.status as Status,
                                    expected: submittedResult.expected,
                                    input: submittedResult.input,
                                    result: submittedResult.result,
                                    stdout: submittedResult.stdout,
                                },
                                false
                            )
                        }
                    })()}
                </div>
            </Grid>
            <Grid container direction="row" justify="space-around" alignItems="center">
                <Grid item>
                    <FormControlLabel
                        control={
                            <Checkbox
                                checked={checkedCustomTestcase}
                                onChange={event => setCheckedCustomTestcase(event.target.checked)}
                                color="primary"
                            />
                        }
                        label="Custom Testcase"
                    />
                </Grid>
                <Grid item>
                    <Grid container justify="space-around" style={{ width: "220px" }}>
                        <Grid item>
                            <Button
                                onClick={() => {
                                    submittedResult = null
                                    onRunCode({typedCode, language, testcase})
                                }}
                                variant="contained"
                                color="primary"
                            >
                                Run code
                            </Button>
                        </Grid>
                        <Grid item>
                            <Button
                                type="button"
                                variant="contained"
                                onClick={() => {
                                    testRunResult.data = null

                                    onSubmitCode({
                                        contestSlug,
                                        typedCode,
                                        language
                                    })
                                }}
                                color="primary"
                            >
                                Submit
                            </Button>
                        </Grid>
                    </Grid>
                </Grid>
            </Grid>
            {checkedCustomTestcase && (
                <TextareaAutosize
                    style={{ maxWidth: "100%", maxHeight: "200px", resize: "vertical" }}
                    aria-label="minimum height"
                    rowsMin={3}
                    defaultValue={testcase}
                    onChange={event => setTestcase(event.target.value)}
                />
            )}
        </Grid>
    )
}

function displayCodeResultError() {
    return (
        <Box border={1} borderColor="error.main">
            Error happens
        </Box>
    )
}
function displayCodeResult(
    title: string,
    language: CodeLanguage,
    contestSlug: string,
    result: {
        status: Status
        expected: string
        input: string
        time?: string
        result: string
        stdout: string
    },
    onlyRun: boolean
): ReactElement {
    console.log("result", result)
    return (
        <Box border={1} borderColor="primary.main">
            <SubmittedResult title={title} language={language} {...result} onlyRun={onlyRun} />
            {result.status === "success" && contestSlug != null ? (
                <Button
                    onClick={() => {
                        Router.push(`/contests/${contestSlug}`)
                    }}
                    color="primary"
                >
                    Back contest
                </Button>
            ) : (
                <></>
            )}
        </Box>
    )
}
