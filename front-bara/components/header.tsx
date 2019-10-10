import Link from 'next/link'
import { withRouter, Router } from 'next/router'
import { Box } from 'grommet';

type Props = {
    router: Router
}

const Header: React.FunctionComponent<Props>  = ({ router: { pathname } }: Props) => (
  <Box
    tag='header'
    background=''
    height='66px'
    border={{
      "color": "border",
      "size": "xsmall",
      "style": "solid",
      "side": "bottom"
    }}
  >
    <Box 
      direction="row"
      pad="medium"
      margin="auto"
      alignSelf="center"
      width="1110px"
      gap="small"
    >
        <Box>
          <Link href='/'>
            <a className={pathname === '/' ? 'is-active' : ''}>Top</a>
          </Link>
        </Box>
        <Box>
          <Link href='/problems'>
            <a className={pathname === '/problems' ? 'is-active' : ''}>Problems</a>
          </Link>
        </Box>
        <Box>
          <Link href='/admin'>
            <a className={pathname === '/admin' ? 'is-active' : ''}>Admin</a>
          </Link>
        </Box>
      </Box>
    <style jsx>{`
      .is-active {
        text-decoration: underline;
      }
    `}</style>
  </Box>
)

export default withRouter(Header)
