const demoInput = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`;

export default () => {
  // const input = Deno.readTextFileSync(import.meta.dirname + "/input");
  const input = demoInput;

  const x = {};
  const allComputers = new Set();

  for (const line of input.trim().split("\n")) {
    const [pc1, pc2] = line.split("-");

    if (pc1 in x) {
      x[pc1].push(pc2);
    } else {
      x[pc1] = [pc2];
    }

    if (pc2 in x) {
      x[pc2].push(pc1);
    } else {
      x[pc2] = [pc1];
    }
  }

  const out = [];

  const go = (
    pc: string,
    out = new Set(),
    visited: Set<string> = new Set()
  ) => {
    visited.add(pc);
    out.add(pc);

    for (const pc2 of x[pc]) {
      for (const pc3 of x[pc2]) {
        if (x[pc3].includes(pc)) {
          out.add(pc3);
        }
        if (!visited.has(pc3)) {
          go(pc3, out, visited);
        }
      }

      // if (x[pc2].includes(pc)) {
      //   out += "-" + pc2;
      // }
      // if (!visited.has(pc2)) {
      //   go(pc2, out, visited);
      // }
    }

    return out;
  };

  for (const pc in x) {
    const path = go(pc);
    out.push(Array.from(path).sort((a, b) => a.localeCompare(b)));
  }

  // pt1
  // for (const pc in x) {
  //   for (const pc2 of x[pc]) {
  //     for (const pc3 of x[pc2]) {
  //       if (x[pc3].includes(pc)) {
  //         out.add([pc, pc2, pc3].sort((a, b) => a.localeCompare(b)).join("-"));
  //       }
  //     }
  //   }
  // }

  // let sum = 0;

  // for (const line of Array.from(out)) {
  //   const [pc1, pc2, pc3] = line.split("-");
  //   if (pc1.at(0) === "t" || pc2.at(0) === "t" || pc3.at(0) === "t") {
  //     sum++;
  //   }
  // }

  console.log(out.map((a) => a.join("-")));
};
