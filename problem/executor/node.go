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
  var rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
    terminal: false
  });
  let testCaseNum = 0;
  let inputNum = 0;
  let lineCount = 0;
  let inputs = [];
  let result = null;
  let expected = ""

  var start = new Date();
  var successFlag = false;
  var countTestCase = 1;

  for await (const line of rl) {
    if(line === "") {
      break;
    }
    if (lineCount === 0) {
      testCaseNum = parseInt(line);
    } else if (lineCount === 1) {
      inputNum = parseInt(line);
    } else if ((inputNum + 1) * countTestCase + 1 === lineCount) {
      const resultStr = getResultStr(result)
      expected = getExpected(line)
      if (resultStr !== expected) {
        successFlag = false;
        console.log(
          JSON.stringify({
            status: "fail",
            result: resultStr,
            input: inputs.join("\n"),
            expected,
            time: 0
          })
        );
        break;
      }

      successFlag = true;
      inputs = [];
      countTestCase++;
    } else {
      try{
        inputs.push(JSON.parse(line));
      } catch {
        inputs.push(line);
      }

      if (inputs.length === inputNum) {
        result = %s(...inputs);
      }
    }
    lineCount++;
  }
  if (successFlag) {
    console.log(
      JSON.stringify({
        status: "success",
        result: getResultStr(result),
        input: inputs.join("\n"), 
        expected,
        time: new Date() - start
      })
    );
  }
}
main();
`
