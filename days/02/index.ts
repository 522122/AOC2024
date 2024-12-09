const demoInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`;

const distance = (a: number, b: number) => {
  return Math.abs(a - b);
};

const isSafe = (distance: number) => {
  return distance <= 3 && distance > 0;
};

const direction = (a: number, b: number) => {
  return a < b ? 1 : -1;
};

const unsafeIndex = (numbers: number[]) => {
  let firstDirection: number | null = null;
  let index = -1;

  for (let i = 0; i < numbers.length - 1; i++) {
    const dir = direction(numbers[i], numbers[i + 1]);
    const dist = distance(numbers[i], numbers[i + 1]);

    if (!isSafe(dist)) {
      index = i;
      break;
    }

    if (firstDirection === null) {
      firstDirection = dir;
    } else {
      if (dir !== firstDirection) {
        index = i;
        break;
      }
    }
  }

  return index;
};

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");

  // const input = demoInput

  let safePt1 = 0;
  let safePt2 = 0;

  for (const line of input.split("\n")) {
    const numbers = line.split(" ").map(Number);
    const numbersCopy = [...numbers];
    const numbersCopy2 = [...numbers];
    const numbersCopy3 = [...numbers];

    const index = unsafeIndex(numbers);
    if (index === -1) safePt1++;

    numbersCopy.splice(index, 1);
    numbersCopy2.splice(index + 1, 1);
    numbersCopy3.splice(index - 1, 1);

    const index2 = unsafeIndex(numbersCopy);
    const index3 = unsafeIndex(numbersCopy2);
    const index4 = unsafeIndex(numbersCopy3);
    if (index2 === -1 || index3 === -1 || index4 === -1) safePt2++;
  }

  console.log("Pt1:", safePt1);
  console.log("Pt2:", safePt2);
};
