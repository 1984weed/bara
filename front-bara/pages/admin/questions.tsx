import * as React from 'react'
import Layout from '../../components/Layout'
import { useMutation, FetchData } from 'graphql-hooks'
import { FormEvent } from 'react'

export const createQuestion = `
mutation createQuestion($title: String!, $description: String!, $functionName: String!, $languageID: CodeLanguage!, $argsNum: Int!, $argsTypes:  [TestCaseArgType!]!, $testCases: [TestCase!]!) {
    createQuestion(input: {title: $title, description: $description, functionName: $functionName, languageID: $languageID, argsNum: $argsNum, argsTypes: $argsTypes, testCases: $testCases}) {
      title
  }
}
`

type Props = {
  onSubmission: () => void;
}

const AdminQuestionPage: React.FunctionComponent<Props> = ({onSubmission}: Props) => {
    const [createPost, state] = useMutation(createQuestion)

    return (
  <Layout title="About | Next.js + TypeScript Example">
    <h1>new Question</h1>
    <p>This is the about page</p>
    <form onSubmit={event => handleSubmit(event, createPost, onSubmission)}>
      <div>
        Title: <input type="text" name="title" />
      </div>
      <div>
          Description: <textarea name="description"></textarea>
      </div>
      <div>
          Description: <textarea name="description"></textarea>
      </div>
      <div>
        Function name: <input type="text" name="functionName" />
      </div>

      <div>
        <div>
          <label>Args: </label><input type="text" name="argsNum" />
        </div>
        <label>Argument Type:</label> 
        <select name="argumentType">
            <option value="number">Number</option>
            <option value="string">String</option>
            <option value="list">List</option>
            <option value="node">Node</option>
        </select>
    </div>
    <div>
      <label>Languege: </label>
      <select name="codeLanguage">
          <option value="js">JavaScript</option>
      </select>
    </div>
    <button type='submit'>{state.loading ? 'Loading...' : 'Submit'}</button>
  </form>
  <style jsx>{`
        form {
          border-bottom: 1px solid #ececec;
          padding-bottom: 20px;
          margin-bottom: 20px;
        }
        h1 {
          font-size: 20px;
        }
        input {
          display: block;
          margin-bottom: 10px;
        }
  `}</style>
  </Layout>
    )
}

export default AdminQuestionPage

async function handleSubmit (event: FormEvent<HTMLFormElement>, createPost: FetchData<any>, onSubmission: () => void) {
    event.preventDefault()
    const form = event.target as HTMLFormElement
    const formData = new FormData(form)
    const title = formData.get('title')
    const description = formData.get('description')
    const functionName = formData.get('functionName')
    const languageID = formData.get('codeLanguage')
    const argsNum = formData.get('argsNum')
    const testCases = [{}]
    const argsTypes = [{}]
    form.reset()
    const result = await createPost({
      variables: {
        title,
        description,
        functionName,
        argsNum,
        languageID,
        testCases,
        argsTypes
      }
    })
    console.log(result)
    onSubmission && onSubmission()
  }