import React, { useState } from "react";
import { Grommet, Box, Button } from "grommet";
import {
  NumberInput,
  Form,
  TextInputField,
  TextAreaField,
  SelectField
} from "grommet-controls";
import { useMutation, FetchData } from "graphql-hooks";
import Layout from "../../components/Layout";
import { CodeLanguage, TestCaseArgType } from "../../graphql/types";

export const createQuestion = `
mutation createQuestion($title: String!, $description: String!, $functionName: String!, $languageID: CodeLanguage!, $argsNum: Int!, $args:  [CodeArg!]!, $testCases: [TestCase!]!) {
    createQuestion(input: {title: $title, description: $description, functionName: $functionName, languageID: $languageID, argsNum: $argsNum, args: $args, testCases: $testCases}) {
      title
  }
}
`;

export default () => {
  const [createPost, state] = useMutation(createQuestion);
  const [argsNum, setArgsNum] = useState(0);
  const [testCaseNum, setTestCaseNum] = useState(1);

  return (
    <Layout title="Admin problems">
      <h1>New problem</h1>
      <p>This is the about page</p>
      <Box alignContent="center">
        <Form
          onSubmit={(event: any) => handleSubmit(event, createPost)}
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
            Args Count
            <NumberInput
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
              <TextInputField label={`Arg Name ${i + 1}:`} name="argName" />
              <SelectField
                label={`Argument Type ${i + 1}:`}
                name="argumentType"
                options={["Number", "String", "List", "Node"]}
              />
            </Box>
          ))}
          <Box>
            Testcase Count
            <NumberInput
              label="Testcase Count"
              name="testcaseNum"
              value={testCaseNum}
              min={1}
              max={10}
              onChange={({ target: { value } }) => setTestCaseNum(value)}
            />
            {Array.from({ length: testCaseNum }).map((_, testcaseIndex) => (
              <Box 
                key={testcaseIndex}
              >
                {Array.from({ length: argsNum }).map((_, i) => (
                  <TextInputField key={i} label={`Arg Name ${i + 1}:`} name="input" />
                ))}
                <TextInputField label="Output" name="output" />
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
          <Box pad="small">
            <Button type="submit" label="Submit" />
          </Box>
        </Form>
      </Box>
    </Layout>
  );
};

async function handleSubmit(
  event: React.FormEvent<HTMLFormElement>,
  createPost: FetchData<any>
) {
  event.preventDefault();
  console.log(createPost);
  const form = event.target as HTMLFormElement;
  const formData = new FormData(form);
  const title = formData.get("title");
  const description = formData.get("description");
  const functionName = formData.get("functionName");
  const languageID = CodeLanguage.JavaScript; //formData.get('codeLanguage')
  const argsNum = 1; //formData.get('argsNum')
  const args = [
    {
      name: "target",
      type: TestCaseArgType.Number
    }
  ];
  const testCases = [
    {
      input: ["10"],
      output: "10"
    }
  ];
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
