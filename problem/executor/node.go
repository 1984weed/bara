package executor

const Node = `
var readline = require("readline");

%s

function getResultStr(result) {
  let resultStr = ""
  try {
    resultStr = JSON.stringify(JSON.parse(result));
  } catch {
    if(Array.isArray(result)) {
      resultStr = "[" + result + "]"
    } else {
      resultStr = result
    }
  }

  return resultStr
}

function getExpected(line) {
  let expected = ""
  try {
    expected = JSON.stringify(JSON.parse(line));
  } catch {
    expected = line
  }
  return expected
}

async function main() {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false
  });
  let inputNum = 0;
  let lineCount = 0;
  let inputs = [];
  let result = null;
  let expected = ""

  let successFlag = false;

  for await (const line of rl) {
    if(line === "") {
      break;
    }
    if (lineCount === 0) {
      inputNum = parseInt(line);
    } else if (lineCount - 1 < inputNum) {
      try{
        inputs.push(JSON.parse(line));
      } catch {
        inputs.push(line);
      }

      if (inputs.length === inputNum) {
        result = %s(...inputs);
      }
    } else if(result != null) {
      const resultStr = getResultStr(result)
      expected = getExpected(line)
      if (resultStr !== expected) {
        successFlag = false;
        console.log(
          JSON.stringify({
            status: "fail",
            result: resultStr,
            expected,
          })
        );
        break;
      } 
      successFlag = true;
    }
    lineCount++;
  }
  if (successFlag) {
    console.log(
      JSON.stringify({
        status: "success",
        result: getResultStr(result),
        expected,
      })
    );
  }
}
main();
`
