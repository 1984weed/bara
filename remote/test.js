// function main(input) {
//     console.log(input);
// }
// main(require('fs').readFileSync('/dev/stdin', 'utf8'));
var readline = require('readline');

var twoSum = function(nums, target) {
    const map = {}
    for(let i = 0; i < nums.length; i++) {
        map[nums[i]]= i
    }
    for(let i = 0; i < nums.length; i++) {
        if(map[target - nums[i]] != null && i != map[target - nums[i]]) {
            return [i, map[target - nums[i]]]
        }
    }
    
    return null
};

function main() {
    var rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout,
        terminal: false
    });
    let testCaseNum = 0
    let inputNum = 0
    let lineCount = 0
    let inputs = []
    let result = null

    rl.on("line", function(line){
        if(lineCount === 0) {
            testCaseNum = parseInt(line)
        } else if(lineCount === 1) {
            inputNum = parseInt(line)
        } else if(inputNum + 2 === lineCount) {
            if(JSON.stringify(result) === JSON.stringify(JSON.parse(line))) {
                console.log("Success")
            } else {
                console.log("Error")
            }
        } else if(line != "\n") {
            inputs.push(JSON.parse(line))

            if(inputs.length === inputNum) {
                result = twoSum(...inputs)
            }
        }
        lineCount++
    })
}

main()
