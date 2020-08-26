const webpack = require("webpack");
const dotenv = require('dotenv')
const fs = require('fs')

const envFileMapping = {
  'production': '.env.production',
}
const file = envFileMapping[process.env.NODE_ENV] || '.env'
try{
  const envConfig = dotenv.parse(fs.readFileSync(file))

  Object.keys(envConfig).forEach((key) => {
    if (!process.env.hasOwnProperty(key)) {
      process.env[key] = envConfig[key];
    }
  });
} catch {}
const withImages = require('next-images')

module.exports = withImages({
  webpack: config => {
    const env = Object.keys(process.env).reduce((acc, curr) => {
      acc[`process.env.${curr}`] = JSON.stringify(process.env[curr]);
      return acc;
    }, {});

    config.node = {
      fs: 'empty'
    }

    config.plugins.push(new webpack.DefinePlugin(env));

    return config
  }
})
