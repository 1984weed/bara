export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string,
  String: string,
  Boolean: boolean,
  Int: number,
  Float: number,
};

export type CodeArg = {
  type: Scalars['String'],
  name: Scalars['String'],
};

export type CodeArgType = {
   __typename?: 'CodeArgType',
  type: Scalars['String'],
  name: Scalars['String'],
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
  input?: Maybe<Scalars['String']>,
  result: Scalars['String'],
  status: Scalars['String'],
  time: Scalars['Int'],
};

export type CodeSnippet = {
   __typename?: 'CodeSnippet',
  code: Scalars['String'],
  lang: CodeLanguage,
};

export type Contest = {
   __typename?: 'Contest',
  id: Scalars['ID'],
  slug: Scalars['String'],
  title: Scalars['String'],
  startTimestamp: Scalars['String'],
  duration?: Maybe<Scalars['String']>,
  problems?: Maybe<Array<Problem>>,
};

export type ContestProblemsUserResult = {
   __typename?: 'ContestProblemsUserResult',
  done: Scalars['Boolean'],
};

export type Mutation = {
   __typename?: 'Mutation',
  createContest: Contest,
  updateContest: Contest,
  deleteContest?: Maybe<Scalars['Boolean']>,
  startCalculateRanking?: Maybe<Ranking>,
  submitCode: CodeResult,
  testRunCode: CodeResult,
  createProblem: Problem,
  updateProblem: Problem,
  updateUser?: Maybe<User>,
};


export type Mutation_CreateContestArgs = {
  newContest: NewContest
};


export type Mutation_UpdateContestArgs = {
  contestID: Scalars['ID'],
  newContest: NewContest
};


export type Mutation_DeleteContestArgs = {
  slug: Scalars['String']
};


export type Mutation_StartCalculateRankingArgs = {
  slug: Scalars['String']
};


export type Mutation_SubmitCodeArgs = {
  input: SubmitCode
};


export type Mutation_TestRunCodeArgs = {
  inputStr: Scalars['String'],
  input: SubmitCode
};


export type Mutation_CreateProblemArgs = {
  input: NewProblem
};


export type Mutation_UpdateProblemArgs = {
  problemID: Scalars['Int'],
  input: NewProblem
};


export type Mutation_UpdateUserArgs = {
  input: UserInput
};

export type NewContest = {
  slug: Scalars['String'],
  title: Scalars['String'],
  startTimestamp: Scalars['String'],
  duration?: Maybe<Scalars['String']>,
  problemIDs?: Maybe<Array<Scalars['ID']>>,
};

export type NewProblem = {
  title: Scalars['String'],
  description: Scalars['String'],
  functionName: Scalars['String'],
  outputType: Scalars['String'],
  argsNum: Scalars['Int'],
  args: Array<CodeArg>,
  testCases: Array<TestCase>,
};

export type Node = {
  id: Scalars['ID'],
};

export type Problem = {
   __typename?: 'Problem',
  score: Scalars['Int'],
  userResult?: Maybe<ContestProblemsUserResult>,
  id: Scalars['Int'],
  slug: Scalars['String'],
  title: Scalars['String'],
  description: Scalars['String'],
  codeSnippets: Array<CodeSnippet>,
  problemDetailInfo?: Maybe<ProblemDetailInfo>,
  sampleTestCase?: Maybe<Scalars['String']>,
};

export type ProblemDetailInfo = {
   __typename?: 'ProblemDetailInfo',
  functionName: Scalars['String'],
  outputType: Scalars['String'],
  argsNum: Scalars['Int'],
  args: Array<CodeArgType>,
  testCases: Array<TestCaseType>,
};

export type ProblemSubmitResult = {
   __typename?: 'ProblemSubmitResult',
  problem: Problem,
  status: Scalars['String'],
  completedTime: Scalars['String'],
};

export type Query = {
   __typename?: 'Query',
  contests: Array<Contest>,
  contest: Contest,
  problems: Array<Problem>,
  problem: Problem,
  me?: Maybe<User>,
  user?: Maybe<User>,
  testNewProblem: Problem,
  submissionList: Array<Submission>,
};


export type Query_ContestsArgs = {
  limit?: Maybe<Scalars['Int']>,
  offset?: Maybe<Scalars['Int']>
};


export type Query_ContestArgs = {
  slug: Scalars['String']
};


export type Query_ProblemsArgs = {
  limit?: Maybe<Scalars['Int']>,
  offset?: Maybe<Scalars['Int']>
};


export type Query_ProblemArgs = {
  slug?: Maybe<Scalars['String']>
};


export type Query_UserArgs = {
  userName?: Maybe<Scalars['String']>
};


export type Query_TestNewProblemArgs = {
  input: NewProblem
};


export type Query_SubmissionListArgs = {
  problemSlug: Scalars['String'],
  limit?: Maybe<Scalars['Int']>,
  offset?: Maybe<Scalars['Int']>
};

export type Ranking = {
   __typename?: 'Ranking',
  rank: Scalars['Int'],
  user: User,
  problemSubmitResult: ProblemSubmitResult,
};

export type RunCode = {
  typedCode: Scalars['String'],
  lang: Scalars['String'],
  slug: Scalars['String'],
};

export type Submission = {
   __typename?: 'Submission',
  id: Scalars['ID'],
  langSlug: CodeLanguage,
  runtimeMS: Scalars['Int'],
  statusSlug: Scalars['String'],
  url: Scalars['String'],
  timestamp: Scalars['String'],
};

export type SubmitCode = {
  typedCode: Scalars['String'],
  lang: Scalars['String'],
  slug: Scalars['String'],
};

export type TestCase = {
  input?: Maybe<Array<Maybe<Scalars['String']>>>,
  output: Scalars['String'],
};

export type TestCaseType = {
   __typename?: 'TestCaseType',
  input?: Maybe<Array<Maybe<Scalars['String']>>>,
  output: Scalars['String'],
};

export type User = {
   __typename?: 'User',
  id: Scalars['ID'],
  displayName: Scalars['String'],
  userName: Scalars['String'],
  email: Scalars['String'],
  image: Scalars['String'],
  role?: Maybe<UserRole>,
  bio: Scalars['String'],
};

export type UserInput = {
  displayName?: Maybe<Scalars['String']>,
  userName?: Maybe<Scalars['String']>,
  email?: Maybe<Scalars['String']>,
  image?: Maybe<Scalars['String']>,
  bio?: Maybe<Scalars['String']>,
};

export enum UserRole {
  Admin = 'admin',
  Normal = 'normal'
}
