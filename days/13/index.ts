const demo = `
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`;

const price = (n, m) => {
  return n * 3 + m;
};

export default () => {
  //n * A + m * B = X
  //n * A + m * B = Y

  //n = (X - m * B) / A
  //m = (Y - n * A) / B

  // const input = Deno.readTextFileSync(import.meta.dirname + "/input");

  const input = demo;

  const lines = input.trim().split("\n\n");
  const results = lines.map((line) => {
    const [buttonA, buttonB, prize] = line.split("\n");
    const [_1, ax, ay] = buttonA.match(/X\+(\d+), Y\+(\d+)/);
    const [_2, bx, by] = buttonB.match(/X\+(\d+), Y\+(\d+)/);
    const [_3, x, y] = prize.match(/X=(\d+), Y=(\d+)/);

    const Ax = parseInt(ax);
    const Bx = parseInt(bx);
    const Ay = parseInt(ay);
    const By = parseInt(by);
    const X = parseInt(x);
    const Y = parseInt(y);

    const maxA = Math.max(Math.floor(X / Ax), Math.floor(Y / Ay));
    const maxB = Math.max(Math.floor(X / Bx), Math.floor(Y / By));

    const max = Math.max(maxA, maxB);

    let o1;
    let o2;

    for (let i = 0; i <= max; i++) {
      const n = (X - i * Bx) / Ax;
      const m = (Y - n * Ay) / By;

      if (
        Number.isInteger(n) &&
        Number.isInteger(m) &&
        n >= 0 &&
        m >= 0 &&
        n <= 100 &&
        m <= 100
      ) {
        o1 = [n, m];
      }
    }

    for (let i = 0; i <= max; i++) {
      const m = (X - i * Ax) / Bx;
      const n = (Y - m * By) / Ay;

      if (
        Number.isInteger(n) &&
        Number.isInteger(m) &&
        n >= 0 &&
        m >= 0 &&
        n <= 100 &&
        m <= 100
      ) {
        o2 = [n, m];
      }
    }

    let r;
    if (o1 && o2) {
      r = price(o1[0], o1[1]) > price(o2[0], o2[1]) ? o2 : o1;
    } else if (o1) {
      r = o1;
    } else if (o2) {
      r = o2;
    }

    console.log(o1, o2, r);
    return r;
  });

  const sum = results.filter(Boolean).reduce((acc, [n, m]) => {
    // console.log(n, m, price(n, m));
    return acc + price(n, m);
  }, 0);

  console.log(sum);
};
