import { useRouter } from 'next/router'
import Layout from '../../components/Layout'
import Editor from '../../components/Editor'

type Props = {}

const Problems: React.FunctionComponent<Props> = ({}: Props) => {
    const router = useRouter()
    const { slug } = router.query

    return (
        <Layout title="">
            {slug}
            <Editor></Editor>
        </Layout>
    )
}

export default Problems