import { useRouter } from "next/router";
import Layout from "../../components/Layout";
import Editor from "../../components/Editor";
import React from "react";
import { useQuery, useMutation, FetchData } from "graphql-hooks";
import { Question, CodeLanguage } from "../../graphql/types";
import { Grid, Button, Box } from "grommet";
import { useFormState } from "react-use-form-state";
import SubmittedResult from "../../components/problems/SubmittedResult";
import useLocalStorage from "../../hooks/useRememberState";

type Props = {};

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
`;

const submitCodeMutation = `
mutation submitCode($typedCode: String!, $lang: String!, $slug: String!) {
    submitCode(input: {typedCode: $typedCode, lang: $lang, slug: $slug}) {
    result {
      status,
      expected,
      time,
      result
    }
  }
}
`;

const getLocalStorageKey = (slug: string) => {
  return `${slug}-typed-code`;
};

const Problem: React.FunctionComponent<Props> = ({  }: Props) => {
  const router = useRouter();
  const { slug } = router.query;

  const [formState, { raw }] = useFormState({ code: null });
  const [typedCode, setTypedCode] = useLocalStorage(
    getLocalStorageKey(slug as string),
    ""
  );

  const { error, data } = useQuery<{ Question: Question }>(problem, {
    variables: { slug }
  });
  const [submitCode, submittedResult] = useMutation(submitCodeMutation);

  if (error) return <span>Error</span>;
  if (!data) return <div>Loading</div>;

  const { Question } = data;
  const language = CodeLanguage.JavaScript;
  const targetCodeSnippet = Question.codeSnippets.find(
    a => a.lang === "JavaScript"
  ) || { code: "" };

  const defaultCode =
    typedCode === "" || typedCode === null ? targetCodeSnippet.code : typedCode;

  return (
    <Layout title="">
      <Grid
        rows={["flex", "60px", "auto"]}
        columns={["flex", "5px", "flex"]}
        gap="1px"
        areas={[
          { name: "description", start: [0, 0], end: [0, 0] },
          { name: "partition", start: [1, 0], end: [1, 0] },
          { name: "editor", start: [2, 0], end: [2, 0] },
          { name: "controls", start: [0, 1], end: [2, 1] },
          { name: "result", start: [0, 2], end: [2, 2] }
        ]}
      >
        <Box gridArea="description">
          <h1>{Question.title}</h1>
          <Box>{Question.description}</Box>
        </Box>
        <div></div>
        <Box gridArea="editor">
          <Editor
            {...raw({
              name: "code"
            })}
            onChange={(typedCode: string) => {
              formState.setField("code", typedCode);
              setTypedCode(typedCode);
            }}
            value={defaultCode}
          />
        </Box>
        <Box gridArea="controls">
          <Box direction="row" justify="end" margin={{ top: "medium" }}>
            <Button
              label="Reset code"
              onClick={() => {
                setTypedCode(targetCodeSnippet.code);
              }}
            />
            <Button label="Cancel" />
            <Button
              type="button"
              label="Submit"
              onClick={() => {
                const typedCode =
                  formState.values.code == null
                    ? targetCodeSnippet.code
                    : formState.values.code;
                return handleSubmit(
                  submitCode,
                  typedCode,
                  language,
                  slug as string
                );
              }}
              primary
            />
          </Box>
        </Box>
        {submittedResult.data != null && (
          <Box gridArea="result">
            <SubmittedResult
              title={Question.title}
              language={language}
              {...submittedResult.data.submitCode.result}
            />
          </Box>
        )}
      </Grid>
    </Layout>
  );
};

async function handleSubmit(
  submitCode: FetchData<any>,
  typedCode: string,
  lang: string,
  slug: string
) {
  const result = await submitCode({
    variables: {
      typedCode,
      lang,
      slug
    }
  });
  console.log(result);
}

export default Problem;
