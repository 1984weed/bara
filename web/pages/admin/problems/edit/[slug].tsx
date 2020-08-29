import { FetchData, useMutation, useQuery } from "graphql-hooks"
import { useRouter } from "next/router"
import React, { useEffect, useState } from "react"
import Layout from "../../../../components/Layout"
import ProblemForm from "../../../../components/problems/ProblemForm"
import { Problem } from "../../../../graphql/types"
import { Session } from "../../../../types/Session"
import { getArrayAutoTestCaseDef } from "../../../../types/TestAutoType"
import ErrorNotification from "../../../../components/notifications/ErrorNotification"
import SuccessNotification from "../../../../components/notifications/SuccessNotification"
import { NextPage } from "next"
// import { NextPageContextWithGraphql } from "../../../../lib/with-graphql-client"
import { useForm } from "react-hook-form"
import { createNewProblemsVariables } from "../new"
import { Box } from "@material-ui/core"

type Props = {
    session: Session
    problem: Problem
}

const problemQuery = `
query getProblemForEdit($slug: String!) {
    problem(slug: $slug) {
        id,
        title,
        description,
        slug,
        problemDetailInfo {
            functionName,
            outputType,
            argsNum,
            args {
                type,
                name
            },
            testCases {
                input,
                output
            }
        }
    }
}
`

const updateProblemMutation = `
mutation updateProblem($problemID: Int!, $slug: String!, $title: String!, $description: String!, $functionName: String!, $outputType: String!, $argsNum: Int!, $args:  [CodeArg!]!, $testCases: [TestCase!]!, $testCaseNum: Int!) {
  updateProblem(problemID: $problemID, input: {title: $title, slug: $slug, description: $description, functionName: $functionName, outputType: $outputType, argsNum: $argsNum, args: $args, testCases: $testCases, testCaseNum: $testCaseNum}) {
    title,
    slug
  }
}
`

const ProblemEditComponent: NextPage<Props> = ({ problem, session }: Props) => {
    const [updateProblem] = useMutation(updateProblemMutation)

    const [submitError, setSubmitError] = useState(false)
    const [updated, setUpdated] = useState(false)
    const testCaseGenState = Array.from({ length: problem.problemDetailInfo.testCases.length }).map(() =>
        Array.from({ length: problem.problemDetailInfo.argsNum }).map(() => getArrayAutoTestCaseDef())
    )

    const router = useRouter()

    const testCases = {}
    for (let i = 0; i < problem.problemDetailInfo.testCases.length; i++) {
        testCases[`inputTestCase-${i}`] = problem.problemDetailInfo.testCases[i].input
        const output = problem.problemDetailInfo.testCases[i].output
        if (testCases["outTestCase"]) {
            testCases["outTestCase"].push(output)
        } else {
            testCases["outTestCase"] = [output]
        }
    }
    const formHooks = useForm({
        defaultValues: {
            title: problem.title,
            description: problem.description,
            slug: problem.slug,
            functionName: problem.problemDetailInfo.functionName,
            outputType: problem.problemDetailInfo.outputType,
            argumentNum: problem.problemDetailInfo.argsNum,
            ...testCases,
            argumentNames: problem.problemDetailInfo.args.map(a => a.name),
            argumentTypes: problem.problemDetailInfo.args.map(a => a.type),
            testCaseNum: problem.problemDetailInfo.testCases.length,
            testCases: problem.problemDetailInfo.testCases,
            testCaseGenState,
        },
    })

    return (
        <Layout title={problem.title} session={session}>
            <ErrorNotification open={submitError} onClose={() => setSubmitError(false)} />
            <SuccessNotification open={updated} label="Success to update this problem" onClose={() => setUpdated(false)} />
            <Box display="flex" >
                <ProblemForm
                    {...formHooks}
                    onClickSubmit={async values => {
                        const result = await handleSubmit(
                            updateProblem,
                            problem.id,
                            createNewProblemsVariables(values, parseInt(values.argumentNum) || 0, parseInt(values.testCaseNum))
                        )

                        if (result == null) {
                            setSubmitError(true)
                            return
                        }
                        setUpdated(true)
                        if(problem.slug !== result.updateProblem.slug) {
                            router.push(`/admin/problems/edit/${result.updateProblem.slug}`)
                        }
                    }}
                ></ProblemForm>
            </Box>
        </Layout>
    )
}

async function handleSubmit(createPost: FetchData<any>, problemID: number, values: any): Promise<any> {
    const { data, error } = await createPost({
        variables: { problemID, ...values },
    })
    if (error != null) {
        return null
    }
    return data
}

// ProblemEditComponent.getInitialProps = async ({ query, client, res }: NextPageContextWithGraphql) => {
//     const result = await client.request(
//         {
//             query: problemQuery,
//             variables: { slug: query.slug },
//         },
//         {}
//     )

//     if (result.data == null) {
//         if (res) {
//             res.statusCode = 404
//             res.end("Not found")
//             return
//         }
//     }

//     return Promise.resolve({
//         problem: ((result.data as unknown) as any).problem as Problem,
//         session: null,
//     })
// }

export default ProblemEditComponent
