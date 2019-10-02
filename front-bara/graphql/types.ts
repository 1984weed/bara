export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string,
  String: string,
  Boolean: boolean,
  Int: number,
  Float: number,
};

export enum CodeLanguage {
  JavaScript = 'JavaScript'
}

export type CodeResult = {
   __typename?: 'CodeResult',
  result: CodeResultDetail,
  stdout: Scalars['String'],
};

export type CodeResultDetail = {
   __typename?: 'CodeResultDetail',
  expected: Scalars['String'],
  result: Scalars['String'],
  status: Scalars['String'],
  time: Scalars['Int'],
};

export type Mutation = {
   __typename?: 'Mutation',
  submitCode: CodeResult,
  createQuestion: Question,
};


export type Mutation_SubmitCodeArgs = {
  input: SubmitCode
};


export type Mutation_CreateQuestionArgs = {
  input: NewQuestion
};

export type NewQuestion = {
  title: Scalars['String'],
  description: Scalars['String'],
  functionName: Scalars['String'],
  languageID: CodeLanguage,
  argsNum: Scalars['Int'],
  argsTypes: Array<TestCaseArgType>,
  testCases: Array<TestCase>,
};

export type Query = {
   __typename?: 'Query',
  Questions: Array<Question>,
};


export type Query_QuestionsArgs = {
  limit?: Maybe<Scalars['Int']>,
  offset?: Maybe<Scalars['Int']>
};

export type Question = {
   __typename?: 'Question',
  slug: Scalars['String'],
  title: Scalars['String'],
  description: Scalars['String'],
};

export type SubmitCode = {
  typedCode: Scalars['String'],
  lang: Scalars['String'],
  slug: Scalars['String'],
};

export type TestCase = {
  input: Array<Scalars['String']>,
  output: Scalars['String'],
};

export enum TestCaseArgType {
  Number = 'NUMBER',
  String = 'STRING'
}
