import { useRouter } from 'next/router'
import Layout from '../../components/Layout'
import Editor from '../../components/Editor'
import React from 'react'
import { useQuery } from 'graphql-hooks'
import {Question} from '../../graphql/types'

type Props = {}

const problem = `
query getQuestion($slug: String!) {
    Question(slug: $slug) {
        title,
        slug,
        codeSnippets {
            lang,
            code
        }
    }
}
`

const Problem: React.FunctionComponent<Props> = ({}: Props) => {
    const router = useRouter()
    // const [skip, setSkip] = useState(0)
    const { slug } = router.query
    const { error, data } = useQuery<{Question: Question}>(problem, {
        variables: { slug },
        updateData: (_, result) => ({
          ...result
        //   allPosts: [...prevResult.allPosts, ...result.allPosts]
        })
      })

    if (error) return <span>Error</span>
    if (!data) return <div>Loading</div>

    console.log(data)
    const { Question } = data
    const targetCodeSnippet = Question.codeSnippets.find(a => a.lang === "JavaScript") || {code: ""};
    console.log(targetCodeSnippet)
    
    return (
        <Layout title="">
            <h1>{slug}</h1>
            <Editor
                value={targetCodeSnippet.code}
            />
        </Layout>
    )
}

export default Problem