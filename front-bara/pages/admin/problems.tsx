import React, { useState } from "react";
import { Grommet, Box, Button } from "grommet";
import {
  NumberInputField,
  Form,
  TextInputField,
  TextAreaField,
  SelectField
} from "grommet-controls";
import { useClientRequest, useMutation, FetchData } from "graphql-hooks";
import Layout from "../../components/Layout";
import { CodeLanguage, TestCaseArgType } from "../../graphql/types";

export const createQuestion = `
mutation createQuestion($title: String!, $description: String!, $functionName: String!, $languageID: CodeLanguage!, $argsNum: Int!, $args:  [CodeArg!]!, $testCases: [TestCase!]!) {
    createQuestion(input: {title: $title, description: $description, functionName: $functionName, languageID: $languageID, argsNum: $argsNum, args: $args, testCases: $testCases}) {
      title
  }
}
`;

export const testQuestion = `
query testNewQuestion($title: String!, $description: String!, $functionName: String!, $languageID: CodeLanguage!, $argsNum: Int!, $args:  [CodeArg!]!, $testCases: [TestCase!]!) {
    testNewQuestion(input: {title: $title, description: $description, functionName: $functionName, languageID: $languageID, argsNum: $argsNum, args: $args, testCases: $testCases}) {
      slug,
      title,
      description, 
      codeSnippets {
        code,
        lang
      }
  }
}
`;

export default () => {
  const [createPost, state] = useMutation(createQuestion);
  const [doTest ] = useClientRequest(testQuestion);
  const [formState, setFormState] = useState({})
  const [argsNum, setArgsNum] = useState(0);
  const [testCaseNum, setTestCaseNum] = useState(1);
  let form: any = {}

  return (
    <Layout title="Admin problems">
      <h1>New problem</h1>
      <p>This is the about page</p>
      <Box alignContent="center">
          <Form
          ref={(instance: any) => form = instance}
          onSubmit={(values) => 
            handleSubmit(createPost, values, argsNum, testCaseNum)
          }
          basis="full"
        >
          <Box>
            <TextInputField label="Title" name="title" />
          </Box>
          <Box>
            <TextAreaField rows="5" label="Description" name="description" />
          </Box>
          <Box>
            <TextInputField label="Function name" name="functionName" />
          </Box>
          <Box>
            <NumberInputField
              label="Args Count"
              name="argsNum"
              value={argsNum}
              min={0}
              max={10}
              onChange={({ target: { value } }) => setArgsNum(value)}
            />
          </Box>
          {Array.from({ length: argsNum }).map((_, i) => (
            <Box key={i}>
              <TextInputField label={`Arg Name ${i + 1}:`} name={`argName[${i}]`} />
              <SelectField
                label={`Argument Type ${i + 1}:`}
                name={`argumentType[${i}]`}
                options={["NUMBER", "STRING", "LIST", "NODE"]}
              />
            </Box>
          ))}
          <Box>
            <NumberInputField
              label="Testcase Count"
              name="testcaseNum"
              value={testCaseNum}
              min={1}
              max={10}
              onChange={({ target: { value } }) => setTestCaseNum(value)}
            />
            {Array.from({ length: testCaseNum }).map((_, testCaseIndex) => (
              <Box 
                key={testCaseIndex}
              >
                {Array.from({ length: argsNum }).map((_, i) => (
                  <TextInputField key={i} label={`Arg Name ${i + 1}:`} name={`inputTestCase[${testCaseIndex}][${i}]`} />
                ))}
                <TextInputField label="Output" name={`outTestCase[${testCaseIndex}]`} />
              </Box>
            ))}
          </Box>
          <Box>
            <SelectField
              label="Language"
              name="codeLanguage"
              options={["JavaScript"]}
            />
          </Box>
          <Box 
            pad="small" 
            alignContent="center"
            direction="row"
            justify="end"
          >
            <Button label="Test" onClick={() => {
              handleTest(doTest, form.state.data, argsNum, testCaseNum)
            }} />
            <Button type="submit" label="Submit" />
          </Box>
        </Form>
      </Box>
    </Layout>
  );
};

function createNewProblemsVariables(formState: any, argsNum: number, testCaseNum: number): any {
  const testCases = []
  for(let i = 0; i < testCaseNum; i++) {
    const inputArray = new Array(argsNum).fill("").map((_, argIndex) => 
          formState[`inputTestCase[${argIndex}][${i}]`]
        )
    testCases.push(
      {
        input: inputArray,
        output: formState[`outTestCase[${i}]`]
      }
    )
  }
  
  const args = []
  for(let i = 0; i < argsNum; i++) {
    args.push({
      name: formState[`argName[${i}]`],
      type: formState[`argumentType[${i}]`]
    })
  }
  const {title, description, functionName, codeLanguage} = formState;
  return {
    title,
    description,
    functionName,
    argsNum,
    languageID: codeLanguage,
    testCases,
    args
  }
}
async function handleTest(doTestform: FetchData<any>, formState: any, argsNum: number, testCaseNum: number){
  const result = await doTestform({
    variables: createNewProblemsVariables(formState, argsNum, testCaseNum)
  });
  return result
}
async function handleSubmit(
  createPost: FetchData<any>,
 formState: any, argsNum: number, testCaseNum: number
) {
  const result = await createPost({
    variables: createNewProblemsVariables(formState, argsNum, testCaseNum)
  });
  console.log(result);
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
