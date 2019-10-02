import React from 'react'
import { Grommet, Box } from 'grommet';
// import { Card, Value } from 'grommet-controls';
const { Card } = require('grommet-controls')

export default () => 
      <Grommet>
        <Box basis='full'>
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

