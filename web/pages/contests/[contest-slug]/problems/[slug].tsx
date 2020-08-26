import { Box, Grid } from "@material-ui/core/"
import { makeStyles } from "@material-ui/core/styles"
import { FetchData, useMutation, useQuery } from "graphql-hooks"
import { NextPage } from "next"
import { useRouter } from "next/router"
import React from "react"
import Layout from "../../../../components/Layout"
import { EditorArea, RunCodeType, SubmitCodeType } from "../../../../components/problems/ProblemPage"
import SideArea from "../../../../components/problems/SideArea"
import { CodeLanguage, Problem } from "../../../../graphql/types"
import { useRememberState } from "../../../../hooks/useRememberState"
import { NextPageContextWithGraphql } from "../../../../lib/with-graphql-client"

type Props = {
    session: any
    problem: Problem | null
}

const problemQuery = `
query getProblem($slug: String!) {
    problem(slug: $slug) {
        title,
        description,
        slug,
        sampleTestCase,
        codeSnippets {
            lang,
            code
        }
    }
}
`

const getSubmissionList = `
    query submissionList($slug: String!, $submissionLimit: Int, $submissionOffset: Int) {
        submissionList(problemSlug: $slug, limit: $submissionLimit, offset: $submissionOffset) {
            id,
            langSlug,
            runtimeMS,
            statusSlug,
            url,
            timestamp
        }
    }
`

const submitCodeMutation = `
mutation submitContestCode($contestSlug: String!, $typedCode: String!, $lang: String!, $slug: String!) {
    submitContestCode(contestSlug: $contestSlug, input: {typedCode: $typedCode, lang: $lang, slug: $slug}) {
    result {
      status,
      expected,
      time,
      result,
      input
    },
    stdout
  }
}
`

const testRunCodeMutation = `
mutation testRunCodeMutation($typedCode: String!, $lang: String!, $slug: String!, $input: String!) {
    testRunCode(inputStr: $input, input: {typedCode: $typedCode, lang: $lang, slug: $slug}) {
    result {
      status,
      expected,
      result,
      input
    },
    stdout
  }
}
`

const getLocalStorageKey = (slug: string) => {
    return `${slug}-typed-code`
}

const useStyles = makeStyles(theme => ({
    mainGrid: {
        flexGrow: 1,
        minHeight: "600px",
    },
}))

const ProblemComponent: NextPage<Props> = ({ session, problem }: Props) => {
    const router = useRouter()

    if (router === null) {
        return <></>
    }
    if (problem == null) {
        return <span>Error</span>
    }
    const { slug } = router.query
    const contestSlug = router.query["contest-slug"] as string

    const { data, refetch } = useQuery(getSubmissionList, {
        variables: { slug, offset: 0, limit: 50 },
    })

    const [submitCode, submittedResult] = useMutation(submitCodeMutation)
    const [testRunCode, testRunResult] = useMutation(testRunCodeMutation)

    const language = CodeLanguage.JavaScript
    const targetCodeSnippet = problem.codeSnippets.find(a => a.lang === "JavaScript") || { code: "" }

    const defaultCode = targetCodeSnippet.code
    const classes = useStyles()

    return (
        <Layout title={problem.title} session={session}>
            <Grid container className={classes.mainGrid}>
                <Grid item xs={12} md={6}>
                    <SideArea
                        title={problem.title}
                        description={problem.description}
                        submissionList={data ? data.submissionList : []}
                        clickSubmission={refetch}
                    />
                </Grid>
                <Box style={{ width: "5px" }}></Box>
                <Grid item xs>
                    <EditorArea
                        testRunResult={testRunResult}
                        submittedResult={
                            submittedResult.data == null
                                ? null
                                : {
                                      status: submittedResult.data.submitContestCode.result.status,
                                      expected: submittedResult.data.submitContestCode.result.expected,
                                      time: submittedResult.data.submitContestCode.result.time,
                                      input: submittedResult.data.submitContestCode.result.input,
                                      result: submittedResult.data.submitContestCode.result.result,
                                      stdout: submittedResult.data.submitContestCode.result.stdout,
                                  }
                        }
                        localTypedCodeKey={getLocalStorageKey(slug as string)}
                        targetCodeSnippet={targetCodeSnippet}
                        defaultCode={defaultCode}
                        sampleTestCase={problem.sampleTestCase}
                        title={problem.title}
                        language={language}
                        contestSlug={contestSlug}
                        onRunCode={(e: RunCodeType) => {
                            handleSubmit(testRunCode, {
                                typedCode: e.typedCode,
                                lang: e.language,
                                slug: slug as string,
                                input: e.testcase,
                            })
                        }}
                        onSubmitCode={(e: SubmitCodeType) => {
                            if (session.user == null) {
                                alert("Please login")
                                return
                            }
                            testRunResult.data = null
                            handleSubmit(submitCode, {
                                contestSlug,
                                typedCode: e.typedCode,
                                lang: e.language,
                                slug: slug as string,
                            })
                        }}
                    />
                </Grid>
            </Grid>
        </Layout>
    )
}

function handleSubmit(
    submitCode: FetchData<any>,
    option: {
        contestSlug?: string
        typedCode: string
        lang: string
        slug: string
        input?: string
    }
) {
    return submitCode({
        variables: option,
    })
}

ProblemComponent.getInitialProps = async ({ query, client }: NextPageContextWithGraphql) => {
    const result = await client.request(
        {
            query: problemQuery,
            variables: { slug: query.slug },
        },
        {}
    )

    if (result.data == null) {
        return Promise.resolve({
            problem: null,
            session: "",
            pathname: "",
        })
    }

    const { problem } = result.data as {
        problem: Problem
    }

    return Promise.resolve({
        problem,
        session: "",
        pathname: "",
    })
}

export default ProblemComponent
