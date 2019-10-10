import * as React from 'react'
import { Box, Grommet } from "grommet";
import Footer from './Footer'
import Header from './Header'

type Props = {
  title?: string
}

const Layout: React.FunctionComponent<Props> = ({
  children,
  title = 'This is the default title',
}) => (
  <Grommet>
    {title}
    <Header />
    <Box
      pad="medium"
      alignContent="center"
      style={{maxWidth: "1100px", margin: "0 auto"}}
    >
      {children}
    </Box>
    <Footer />
  </Grommet>
)

export default Layout
