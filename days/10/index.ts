const demoInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`;

const directions = [
  [-1, 0], // up
  [0, 1], // right
  [1, 0], // down
  [0, -1], // left
];

const parseInput = (input: string) => {
  return input
    .trim()
    .split("\n")
    .map((row) => row.split("").map(Number));
};

function isValidPosition(r: number, c: number, grid: number[][]) {
  return r >= 0 && r < grid.length && c >= 0 && c < grid[0].length;
}

const findStarts = (grid: number[][]) => {
  const starts = [];
  for (let r = 0; r < grid.length; r++) {
    for (let c = 0; c < grid[0].length; c++) {
      if (grid[r][c] === 0) {
        starts.push([r, c]);
      }
    }
  }
  return starts;
};

const findNext = (r: number, c: number, grid: number[][]) => {
  const current = grid[r][c];
  const shouldGo = current + 1;

  const next = [];

  for (const [dr, dc] of directions) {
    const newR = r + dr;
    const newC = c + dc;

    if (!isValidPosition(newR, newC, grid)) continue;

    if (grid[newR][newC] === shouldGo) {
      next.push([newR, newC]);
    }
  }

  return next;
};

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");
  // const input = demoInput;

  const grid = parseInput(input);

  const starts = findStarts(grid);

  const trails: Record<string, Array<[number, number]>> = {};

  for (const s of starts) {
    // const next = findNext(s[0], s[1], grid);
    // console.log(next);

    const currentTrail: Array<[number, number]> = [];

    let queue = findNext(s[0], s[1], grid);

    while (queue.length > 0) {
      const [r, c] = queue.shift() as [number, number];
      if (grid[r][c] === 9) {
        currentTrail.push([r, c]);
      } else {
        // console.log(r, c, findNext(r, c, grid));
        queue.push(...findNext(r, c, grid));
        // part 2 =)
        // make queue unique
        // queue = queue.filter(
        //   (v, i, a) => a.findIndex((t) => t[0] === v[0] && t[1] === v[1]) === i
        // );
      }
    }
    trails[JSON.stringify(s)] = currentTrail;
  }

  // console.log(trails);

  console.log(
    Object.values(trails).reduce((acc, curr) => acc + curr.length, 0)
  );
};
