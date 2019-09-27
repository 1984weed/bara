package remote

const nodeJsTemplate = `
var readline = require("readline");

%s

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

  var start = new Date();
  var successFlag = false;
  var countTestCase = 1;

  for await (const line of rl) {
    if (lineCount === 0) {
      testCaseNum = parseInt(line);
    } else if (lineCount === 1) {
      inputNum = parseInt(line);
    } else if ((inputNum + 1) * countTestCase + 1 === lineCount) {
      const expected = JSON.stringify(JSON.parse(line));
      const resultStr = JSON.stringify(result);
      if (resultStr !== expected) {
        successFlag = false;
        console.log(
          JSON.stringify({
            status: "fail",
            result: resultStr,
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
      inputs.push(JSON.parse(line));

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
		result: "",
		expected: "",
        time: new Date() - start
      })
    );
  }
}
main();
`
