import Box from "@material-ui/core/Box"
import { FetchData, useMutation } from "graphql-hooks"
import { useRouter } from "next/router"
import React, { useState } from "react"
import { useForm } from "react-hook-form"
import Layout from "../../../components/Layout"
import ErrorNotification from "../../../components/notifications/ErrorNotification"
import ProblemForm from "../../../components/problems/ProblemForm"
import { useRememberState } from "../../../hooks/useRememberState"
import { useProtectAdminPage } from "../../../hooks/useProtectAdminPage"

export const createProblem = `
mutation createProblem($title: String!, $slug: String!, $description: String!, $functionName: String!, $outputType: String!, $argsNum: Int!, $args:  [CodeArg!]!, $testCaseNum: Int!, $testCases: [TestCase!]!) {
  createProblem(input: {title: $title, slug: $slug, description: $description, functionName: $functionName, outputType: $outputType, argsNum: $argsNum, args: $args, testCaseNum: $testCaseNum, testCases: $testCases}) {
    title,
    slug
  }
}
`

export const testQuestion = `
query testNewQuestion($title: String!, $description: String!, $functionName: String!, $outputType: String!, $argsNum: Int!, $args:  [CodeArg!]!, $testCases: [TestCase!]!) {
    testNewQuestion(input: {title: $title, description: $description, functionName: $functionName, outputType: $outputType, argsNum: $argsNum, args: $args, testCases: $testCases}) {
      slug,
      title,
      description, 
      codeSnippets {
        code,
        lang
      }
  }
}
`

const NEW_PROBLEM_KEY = "new-problem-form"

const NewProblem = () => {
    const router = useRouter()

    // Protect this page
    useProtectAdminPage()

    const [createPost, { loading }] = useMutation(createProblem)

    const [submitError, setSubmitError] = useState(false)

    const [formState, setFormState] = useRememberState(
        NEW_PROBLEM_KEY,
        JSON.stringify({
            title: "",
            slug: "",
            description: "",
            functionName: "",
            outputType: "int[]",
            testCaseNum: "0",
            argumentNum: "0",
            argumentNames: [],
            argumentTypes: [],
            testCaseGenState: [],
            testCaseGen: [{ count: 0 }],
        })
    )

    if (loading) {
        return <div>Loading</div>
    }

    const formHooks = useForm({
        defaultValues: JSON.parse(formState),
    })
    const currentState = formHooks.watch({ nest: true })

    setTimeout(() => {
        setFormState(JSON.stringify(currentState))
    })

    return (
        <Layout title="Admin problems">
            <h1>Create a new problem</h1>

            {submitError && <ErrorNotification onClose={() => setSubmitError(false)} />}
            <Box>
                <ProblemForm
                    {...formHooks}
                    onClickSubmit={async values => {
                        const data = await handleSubmit(createPost, values)

                        if (data) {
                            router.push(`/problems/${data.createProblem.slug}`)
                            return
                        }

                        setSubmitError(true)
                        localStorage.setItem(NEW_PROBLEM_KEY, "")
                    }}
                ></ProblemForm>
            </Box>
        </Layout>
    )
}

export function createNewProblemsVariables(formState: any, argsNum: number, testCaseNum: number): any {
    const testCases = []
    for (let i = 0; i < testCaseNum; i++) {
        const inputArray = formState[`inputTestCase-${i}`]
        inputArray.forEach(s => {
            s.replace(/ /g, "")
        })

        testCases.push({
            input: inputArray,
            output: formState["outTestCase"][i].replace(/ /g, ""),
        })
    }

    const args = []
    for (let i = 0; i < argsNum; i++) {
        args.push({
            name: formState["argumentNames"][i],
            type: formState["argumentTypes"][i],
        })
    }
    const { title, slug, description, functionName, codeLanguage, outputType } = formState

    return {
        title,
        slug: slug === "" ? null : slug,
        description,
        functionName,
        argsNum,
        languageID: codeLanguage,
        outputType,
        testCases,
        testCaseNum,
        args,
    }
}
async function handleTest(doTestform: FetchData<any>, formState: any, argsNum: number, testCaseNum: number) {
    const result = await doTestform({
        variables: createNewProblemsVariables(formState, argsNum, testCaseNum),
    })
    return result
}

async function handleSubmit(createPost: FetchData<any>, formState: any): Promise<any> {
    const { data, error } = await createPost({
        variables: createNewProblemsVariables(formState, parseInt(formState.argumentNum) || 0, parseInt(formState.testCaseNum)),
    })
    if (error != null) {
        return null
    }

    return data
}

export default NewProblem
