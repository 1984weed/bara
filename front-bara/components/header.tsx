import Link from 'next/link'
import { withRouter, Router } from 'next/router'
import { Box } from 'grommet';

type Props = {
    router: Router
}

const Header: React.FunctionComponent<Props>  = ({ router: { pathname } }: Props) => (
  <header>
    <Box>
      <Link href='/'>
        <a className={pathname === '/' ? 'is-active' : ''}>Home</a>
      </Link>
      <Link href='/about'>
        <a className={pathname === '/about' ? 'is-active' : ''}>About</a>
      </Link>
    </Box>
    <style jsx>{`
      header {
        margin-bottom: 25px;
      }
      a {
        font-size: 14px;
        margin-right: 15px;
        text-decoration: none;
      }
      .is-active {
        text-decoration: underline;
      }
    `}</style>
  </header>
)

export default withRouter(Header)
