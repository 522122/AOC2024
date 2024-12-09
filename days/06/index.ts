const demoInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`;

function rotateVector90Degrees(vector: [number, number]): [number, number] {
  const [dx, dy] = vector;
  return [-dy, dx];
}

const play = (matrix: string[][], startX: number, startY: number) => {
  const maxLoopSize = matrix.length * matrix[0].length;
  const guard = {
    x: startX,
    y: startY,
  };
  const direction: [number, number] = [0, -1];
  const positions = new Set();
  let steps = 0;
  while (true) {
    const [dx, dy] = direction;

    const nx = guard.x + dx;
    const ny = guard.y + dy;

    switch (matrix[ny]?.[nx]) {
      case "O":
      case "#": {
        const newDirection = rotateVector90Degrees(direction);
        direction[0] = newDirection[0];
        direction[1] = newDirection[1];
        break;
      }
      case "^":
      case ".":
        if (!positions.has(`${guard.x},${guard.y}`)) {
          positions.add(`${guard.x},${guard.y}`);
        }
        steps++;
        guard.x = nx;
        guard.y = ny;
        break;
      case undefined:
        return [positions.size + 1, steps, false];
    }

    if (steps > maxLoopSize) {
      return [positions.size + 1, steps, true];
    }
  }
};

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");
  // const input = demoInput

  const lines = input.split("\n");

  const matrix = lines.map((line) => line.split(""));

  const [startX, startY] = matrix.reduce((acc, row, y) => {
    const x = row.findIndex((c) => c === "^");
    return x !== -1 ? [x, y] : acc;
  }, [-1, -1]);

  let infinite = 0;

  for (let line = 0; line < matrix.length; line++) {
    for (let col = 0; col < matrix[line].length; col++) {
      if (matrix[line][col] === "." || matrix[line][col] === "^") {
        matrix[line][col] = "O";
      }
      if (matrix[line][col] === "#") {
        continue;
      }
      const result = play(matrix, startX, startY);

      if (result[2]) {
        infinite++;
      }

      matrix[line][col] = ".";
    }
  }

  console.log(play(matrix, startX, startY));
  console.log(infinite);
};
