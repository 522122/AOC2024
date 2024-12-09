const demoInput = `3   4
4   3
2   5
1   3
3   9
3   3`;

const mapC = (c: number[]) => {
  const m = new Map<number, number>();
  for (const n of c) {
    m.set(n, (m.get(n) || 0) + 1);
  }
  return m;
};

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");

  // const input = demoInput

  const c1: number[] = [];
  const c2: number[] = [];
  const distances: number[] = [];
  const similarities: number[] = [];

  for (const line of input.split("\n")) {
    const [a, b] = line.split("   ");
    c1.push(Number(a));
    c2.push(Number(b));
  }

  c1.sort((a, b) => a - b);
  c2.sort((a, b) => a - b);
  const m = mapC(c2);

  for (let i = 0; i < c1.length; i++) {
    distances.push(Math.abs(c1[i] - c2[i]));
    similarities.push((m.get(c1[i]) || 0) * c1[i]);
  }

  const sum1 = distances.reduce((acc, curr) => acc + curr, 0);
  const sum2 = similarities.reduce((acc, curr) => acc + curr, 0);

  console.log("Pt1:", sum1);
  console.log("Pt2:", sum2);
};
