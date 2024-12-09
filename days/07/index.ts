const demoInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`;

const OPERATORS = ["+", "*", "||"];

const generateOperators = (n: number) => {
  const op: string[][] = [];
  const generate = (i: number, arr: string[]) => {
    if (i === n) {
      op.push(arr);
      return;
    }
    for (const operator of OPERATORS) {
      generate(i + 1, [...arr, operator]);
    }
  };
  generate(0, []);
  return op;
};

const evaluate = (
  testValue: number,
  numbers: number[],
  operators: string[],
) => {
  let result = numbers[0];

  for (const n of numbers.slice(1)) {
    const operator = operators.shift();
    if (operator === "+") {
      result += n;
    } else if (operator === "*") {
      result *= n;
    } else {
      result = parseInt(`${result}${n}`);
    }
  }

  // console.log(result, testValue)
  return result === testValue;
};

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");
  // const input = demoInput

  const lines = input.trim().split("\n");
  let sum = 0;

  for (const line of lines) {
    const [testValue, ...numbers] = line.split(/:\s*|\s+/).map(Number);
    // console.log(testValue, numbers)

    for (const op of generateOperators(numbers.length)) {
      if (evaluate(testValue, numbers, op)) {
        sum += testValue;
        break;
      }
    }
  }

  console.log(sum);
};
