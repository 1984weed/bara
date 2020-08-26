FROM node:12-slim

# Setting working directory. All the path will be relative to WORKDIR
WORKDIR /usr/src/app

# Installing dependencies
COPY package*.json ./

RUN npm install

# Copying source files
COPY . .

ARG GRAPHQL_URL 

ENV GRAPHQL_URL=${GRAPHQL_URL}

# Building app
RUN npm run build

EXPOSE 3000

CMD [ "npm", "run", "start" ]
