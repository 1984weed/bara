import React from 'react'
import { Grommet, Box } from 'grommet';
import Layout from '../components/Layout';
const { Card } = require('grommet-controls')

export default () => 
  <Layout> 
    <Grommet>
      <Box 
        pad='medium'
        basis='full'
        height="66px"
      >
        Problems
        <Card height='small'>
          <Card.CardTitle>
            Problem 1
          </Card.CardTitle>
          <Card.CardContent align='center'>
            説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明
          </Card.CardContent>
        </Card>
        <Card height='small'>
          <Card.CardTitle>
            Problem 2
          </Card.CardTitle>
          <Card.CardContent align='center'>
            説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明 説明説明
          </Card.CardContent>
        </Card>
      </Box>
    </Grommet>
  </Layout>
