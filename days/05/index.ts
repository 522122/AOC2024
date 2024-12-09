const demoInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`;

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");
  // const input = demoInput

  const [rules, pages] = input.split("\n\n");

  const rulesArr = rules.split("\n").map((rule) => rule.split("|").map(Number));
  const rulesMap: Record<number, number[]> = {};

  const pagesArr = pages.split("\n").map((page) => page.split(",").map(Number));

  for (const [n, after] of rulesArr) {
    rulesMap[n] = rulesMap[n] ?? [];
    rulesMap[n].push(after);
  }

  // console.log(rulesMap, pagesArr)

  const correctLines = [];
  const incorrectLines = [];

  for (const line of pagesArr) {
    let hasCorrectOrder = true;
    const rulesForLine = rulesArr.filter((r) =>
      line.includes(r[0]) && line.includes(r[1])
    );
    for (let i = 0; i < line.length; i++) {
      const [n, ...rest] = line.slice(i);
      // const prev = line.slice(0, i)

      const rulesForN = rulesForLine.filter((r) => r.includes(n));

      for (const n2 of rest) {
        if (rulesForN.some((r) => r[1] === n2)) {
          // console.log('correct order')
        } else {
          hasCorrectOrder = false;
          break;
        }
      }
    }
    if (hasCorrectOrder) {
      correctLines.push(line);
    } else {
      incorrectLines.push(line);
    }
  }

  const middles = correctLines.map((line) => {
    const middle = line[Math.floor(line.length / 2)];
    return middle;
  });

  const sorted = [];

  for (const line of incorrectLines) {
    line.sort((a, b) => {
      const aRules = rulesMap[a] ?? [];
      const bRules = rulesMap[b] ?? [];

      if (aRules.includes(b)) {
        return -1;
      } else if (bRules.includes(a)) {
        return 1;
      } else {
        return 0;
      }
    });
    sorted.push(line);
  }

  const middles2 = sorted.map((line) => {
    const middle = line[Math.floor(line.length / 2)];
    return middle;
  });

  console.log(middles.reduce((acc, curr) => acc + curr, 0));
  console.log(middles2.reduce((acc, curr) => acc + curr, 0));
};
