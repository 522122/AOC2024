const demoInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`;

const REF = "XMAS";

const directions: Array<[number, number]> = [
  [0, -1],
  [0, 1],
  [-1, 0],
  [1, 0],
  [-1, -1],
  [-1, 1],
  [1, -1],
  [1, 1],
];

const checkX = (matrix: string[][], x: number, y: number) => {
  /*
        M . S
        . A .
        M . S
    */
  const M = matrix[x - 1]?.[y - 1];
  const S = matrix[x - 1]?.[y + 1];
  const M2 = matrix[x + 1]?.[y - 1];
  const S2 = matrix[x + 1]?.[y + 1];

  if (M === "M" && S === "S" && M2 === "M" && S2 === "S") {
    return true;
  }

  if (M === "S" && S === "M" && M2 === "S" && S2 === "M") {
    return true;
  }

  if (M === "S" && S === "S" && M2 === "M" && S2 === "M") {
    return true;
  }

  if (M === "M" && S === "M" && M2 === "S" && S2 === "S") {
    return true;
  }

  return false;
};

const justGo = (
  matrix: string[][],
  x: number,
  y: number,
  d: [number, number],
) => {
  const cur = matrix[x][y];
  const next = REF[REF.indexOf(cur) + 1];

  if (cur === "S") return "DONE";

  if (matrix[x + d[0]]?.[y + d[1]] === next) {
    return justGo(matrix, x + d[0], y + d[1], d);
  }

  return null;
};

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");
  // const input = demoInput

  const lines = input.split("\n");

  const matrix = lines.map((line) => line.split(""));

  let sum1 = 0;
  let sum2 = 0;

  for (let i = 0; i < matrix.length; i++) {
    const line = matrix[i];
    for (let j = 0; j < line.length; j++) {
      const cell = line[j];
      if (cell === "X") {
        for (const d of directions) {
          const res = justGo(matrix, i, j, d);
          if (res === "DONE") {
            sum1 += 1;
          }
        }
      }

      if (cell === "A") {
        const res = checkX(matrix, i, j);
        if (res) {
          sum2 += 1;
        }
      }
    }
  }

  console.log(sum1);
  console.log(sum2);
};
