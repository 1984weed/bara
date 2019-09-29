// import App from '../components/app'
// import Header from '../components/header'

// export default () => (
//   <App>
//     <Header />
//     Questions list: 

//   </App>
// )
import * as React from 'react'
import Link from 'next/link'
import Layout from '../components/Layout'

const QuestionPage: React.FunctionComponent = () => (
  <Layout title="About | Next.js + TypeScript Example">
    <h1>About</h1>
    <p>This is the about page</p>
    <p>
      <Link href="/">
        <a>Go home</a>
      </Link>
    </p>
  </Layout>
)

export default QuestionPage
