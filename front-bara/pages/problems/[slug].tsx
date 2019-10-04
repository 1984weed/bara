import { useRouter } from 'next/router'
import Layout from '../../components/Layout'
import Editor from '../../components/Editor'
import React from 'react'
import { useQuery } from 'graphql-hooks'
import {Question} from '../../graphql/types'
import { Grid, Button, Box } from 'grommet';

type Props = {}

const problem = `
query getQuestion($slug: String!) {
    Question(slug: $slug) {
        title,
        description,
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
        })
      })

    if (error) return <span>Error</span>
    if (!data) return <div>Loading</div>

    const { Question } = data
    const targetCodeSnippet = Question.codeSnippets.find(a => a.lang === "JavaScript") || {code: ""};

    return (
        <Layout title="">
            <Grid
            rows={['flex', '60px']}
            columns={['flex', '5px', 'flex']}
            gap='1px'
            areas={[
            { name: 'description', start: [0, 0], end: [0, 0] },
            { name: 'partition', start: [1, 0], end: [1, 0] },
            { name: 'editor', start: [2, 0], end: [2, 0] },
            { name: 'controls', start: [0, 1], end: [2, 1] },
            ]}
        >
                <Box gridArea='description'>
                    <h1>{Question.title}</h1>
                    <Box>
                        {Question.description}
                    </Box>
                </Box>
                <div></div>
                <Box gridArea='editor'>
                    <Editor
                        value={targetCodeSnippet.code}
                    />
                </Box>
                <Box gridArea='controls'>
                    <Box direction="row" justify="end" margin={{ top: "medium" }}>
                        <Button label="Cancel" />
                        <Button type="submit" label="Submit" primary />
                    </Box>
                </Box>
            </Grid>
        </Layout>
    )
}

export default Problem