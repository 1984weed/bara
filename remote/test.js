var readline = require("readline");

function twoSum(nums, target) {
  console.log("==========aaa")
  console.log(process.env);
  return 9

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

  var start = new Date();
  var successFlag = false;
  var countTestCase = 1;

  debugger;
  for await (const line of rl) {
    // Each line in the readline input will be successively available here as
    // `line`.
    if (lineCount === 0) {
      testCaseNum = parseInt(line);
    } else if (lineCount === 1) {
      inputNum = parseInt(line);
    } else if ((inputNum + 1) * countTestCase + 1 === lineCount) {
        debugger;
      const expected = JSON.stringify(JSON.parse(line));
      const resultStr = JSON.stringify(result);
      if (resultStr !== expected) {
        successFlag = false;
        debugger;
        console.log(
          JSON.stringify({
            status: "fail",
            result: resultStr,
            input: inputs.join("\n"),
            expected
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
        result = twoSum(...inputs);
      }
    }
    lineCount++;
  }
  if (successFlag) {
    console.log(
      JSON.stringify({
        status: "success",
        time: new Date() - start
      })
    );
  }
}
main();