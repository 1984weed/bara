import React from 'react'
import { Grommet, Box, Button } from 'grommet';
import { Card, Value, Form, TextInputField, TextAreaField, SelectField } from 'grommet-controls';
import { useMutation, FetchData } from 'graphql-hooks'
import Layout from '../../components/Layout'
import { CodeLanguage, TestCaseArgType } from '../../graphql/types'

export const createQuestion = `
mutation createQuestion($title: String!, $description: String!, $functionName: String!, $languageID: CodeLanguage!, $argsNum: Int!, $args:  [CodeArg!]!, $testCases: [TestCase!]!) {
    createQuestion(input: {title: $title, description: $description, functionName: $functionName, languageID: $languageID, argsNum: $argsNum, args: $args, testCases: $testCases}) {
      title
  }
}
`

export default () =>  {
    const [createPost, state] = useMutation(createQuestion)
  return (
      <Grommet>
   <Layout title="Admin problems">
     <h1>New problem</h1>
     <p>This is the about page</p>
     <Form onSubmit={(event: any) => handleSubmit(event, createPost)}>
       <div>
         <TextInputField label='Title' name='title'  />
       </div>
       <div>
         <TextAreaField
          rows='5'
          label='Description'
          name='description'
        />
       </div>
       <div>
         <TextInputField label='Function name' name='functionName'  />
       </div>
       <div>
         <div>
          <TextInputField label='Args' name='argsNum'  />
         </div>
         <SelectField
            label='Argument Type:'
            name='argumentType'
            options={['Number', 'String', 'List', 'Node']}
          />
     </div>
     <div>
         <SelectField
            label='Language'
            name='codeLanguage'
            options={['JavaScript']}
          />
     </div>
    <Box pad='small'>
      <Button type='submit' label='Submit' />
    </Box>
   </Form>
   </Layout>
      </Grommet>

  )

}

async function handleSubmit (event: React.FormEvent<HTMLFormElement>, createPost: FetchData<any>) {
    event.preventDefault()
    console.log(createPost)
    const form = event.target as HTMLFormElement
    const formData = new FormData(form)
    const title = formData.get('title')
    const description = formData.get('description')
    const functionName = formData.get('functionName')
    const languageID = CodeLanguage.JavaScript;//formData.get('codeLanguage')
    const argsNum = 1;//formData.get('argsNum')
    const args = [{
      name: "target",
      type: TestCaseArgType.Number
    }]
    const testCases = [{
      input: ["10"],
      output: "10"
    }];
    // form.reset()
    const result = await createPost({
      variables: {
        title,
        description,
        functionName,
        argsNum,
        languageID,
        testCases,
        args
      }
    })
    console.log(result)
  }

// import * as React from 'react'
// import Layout from '../../components/Layout'
// import { useMutation, FetchData } from 'graphql-hooks'
// import { CodeLanguage, TestCaseArgType } from '../../graphql/types'
// import { Form, TextInputField } from 'grommet-controls';

// export const createQuestion = `
// mutation createQuestion($title: String!, $description: String!, $functionName: String!, $languageID: CodeLanguage!, $argsNum: Int!, $args:  [CodeArg!]!, $testCases: [TestCase!]!) {
//     createQuestion(input: {title: $title, description: $description, functionName: $functionName, languageID: $languageID, argsNum: $argsNum, args: $args, testCases: $testCases}) {
//       title
//   }
// }
// `

// type Props = {
//   onSubmission: () => void;
// }

// const AdminQuestionPage: React.FunctionComponent<Props> = ({onSubmission}: Props) => {
//     const [createPost, state] = useMutation(createQuestion)


//     return (
//   <Layout title="Admin problems">
//     <h1>New problem</h1>
//     <p>This is the about page</p>
//     <Form onSubmit={(event: any) => handleSubmit(event, createPost, onSubmission)}>
//       <div>
//         Title: <input type="text" name="title" />
//         <TextInputField label='Text' name='text'  />
//       </div>
//       <div>
//         Description: <textarea name="description"></textarea>
//       </div>
//       <div>
//         Function name: <input type="text" name="functionName" />
//       </div>
//       <div>
//         <div>
//           <label>Args: </label><input type="text" name="argsNum" />
//         </div>
//         <label>Argument Type:</label> 
//         <select name="argumentType">
//             <option value="number">Number</option>
//             <option value="string">String</option>
//             <option value="list">List</option>
//             <option value="node">Node</option>
//         </select>
//     </div>
//     <div>
//       <label>Language: </label>
//       <select name="codeLanguage">
//           <option value="js">JavaScript</option>
//       </select>
//     </div>
//     <button type='submit'>{state.loading ? 'Loading...' : 'Submit'}</button>
//   </Form>
//   <style jsx>{`
//         form {
//           border-bottom: 1px solid #ececec;
//           padding-bottom: 20px;
//           margin-bottom: 20px;
//         }
//         h1 {
//           font-size: 20px;
//         }
//         input {
//           display: block;
//           margin-bottom: 10px;
//         }
//   `}</style>
//   </Layout>
//     )
// }

// export default AdminQuestionPage

// async function handleSubmit (event: React.FormEvent<HTMLFormElement>, createPost: FetchData<any>, onSubmission: () => void) {
//     event.preventDefault()
//     console.log(createPost)
//     console.log(onSubmission)
//     const form = event.target as HTMLFormElement
//     const formData = new FormData(form)
//     const title = formData.get('title')
//     const description = formData.get('description')
//     const functionName = formData.get('functionName')
//     const languageID = CodeLanguage.JavaScript;//formData.get('codeLanguage')
//     const argsNum = 1;//formData.get('argsNum')
//     const args = [{
//       name: "target",
//       type: TestCaseArgType.Number
//     }]
//     const testCases = [{
//       input: ["10"],
//       output: "10"
//     }];
//     // form.reset()
//     const result = await createPost({
//       variables: {
//         title,
//         description,
//         functionName,
//         argsNum,
//         languageID,
//         testCases,
//         args
//       }
//     })
//     console.log(result)
//     onSubmission && onSubmission()
//   }