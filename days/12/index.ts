function parseMap(input: string): string[][] {
  return input
    .trim()
    .split("\n")
    .map((line) => line.split(""));
}

class MapPosition {
  constructor(public row: number, public col: number) {}

  stringify() {
    return `${this.row},${this.col}`;
  }
}

const plantTypeAt = (plant: MapPosition, map: string[][]) => {
  return map[plant.row]?.[plant.col];
};

class Region {
  members: MapPosition[] = [];
  constructor(public plantType: string) {}

  get size() {
    return this.members.length;
  }

  hasMember(p: MapPosition) {
    return this.members.some((m) => m.stringify() === p.stringify());
  }

  getMember(p: MapPosition) {
    return this.members.find((m) => m.stringify() === p.stringify());
  }

  get price2() {
    console.log("calculating price for", this.plantType);
    return this.size * this.totalFencePieces;
  }
  get price1() {
    return (
      this.size * Object.values(this.edges).reduce((acc, f) => acc + f.size, 0)
    );
  }

  countFence(
    edges: {
      top: Set<MapPosition>;
      bottom: Set<MapPosition>;
      left: Set<MapPosition>;
      right: Set<MapPosition>;
    },
    direction: "top" | "bottom" | "left" | "right"
  ) {
    const arr = Array.from(edges[direction]);

    let fenceCount = 0;

    while (arr.length > 0) {
      const p = arr.pop();
      if (p == null) {
        break;
      }
      const neighbor = arr.findIndex((a) => {
        if (direction === "left") {
          return a.col === p.col && Math.abs(a.row - p.row) === 1;
        } else if (direction === "right") {
          return a.col === p.col && Math.abs(a.row - p.row) === 1;
        } else if (direction === "top") {
          return a.row === p.row && Math.abs(a.col - p.col) === 1;
        } else {
          return a.row === p.row && Math.abs(a.col - p.col) === 1;
        }
      });

      if (neighbor === -1) {
        fenceCount += 1;
      }
    }

    return fenceCount;
  }
  get totalFencePieces() {
    return Object.values(this.fence).reduce((acc, f) => acc + f, 0);
  }

  get fence() {
    const edges = this.edges;
    return {
      top: this.countFence(edges, "top"),
      bottom: this.countFence(edges, "bottom"),
      left: this.countFence(edges, "left"),
      right: this.countFence(edges, "right"),
    };
  }

  get edges() {
    const rows = this.members.map((p) => p.row);
    const cols = this.members.map((p) => p.col);

    const edges: {
      top: Set<MapPosition>;
      bottom: Set<MapPosition>;
      left: Set<MapPosition>;
      right: Set<MapPosition>;
    } = {
      top: new Set(),
      bottom: new Set(),
      left: new Set(),
      right: new Set(),
    };

    for (const r of rows) {
      for (const c of cols) {
        const member = this.getMember(new MapPosition(Number(r), Number(c)));

        if (member == null) {
          continue;
        }

        // top
        if (!this.hasMember(new MapPosition(Number(r) - 1, Number(c)))) {
          edges.top.add(member);
        }
        // bottom
        if (!this.hasMember(new MapPosition(Number(r) + 1, Number(c)))) {
          edges.bottom.add(member);
        }
        // left
        if (!this.hasMember(new MapPosition(Number(r), Number(c) - 1))) {
          edges.left.add(member);
        }
        // right
        if (!this.hasMember(new MapPosition(Number(r), Number(c) + 1))) {
          edges.right.add(member);
        }
      }
    }

    return edges;
  }

  normalize() {
    const minRow = Math.min(...this.members.map((p) => p.row));
    const minCol = Math.min(...this.members.map((p) => p.col));
    this.members = this.members.map((p) => {
      return new MapPosition(p.row - minRow, p.col - minCol);
    });

    return this;
  }
}

// Example usage
const demoInput = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`;

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");

  const map = parseMap(input);

  const seen = new Set();

  const regions = [];

  for (const r in map) {
    for (const c in map[r]) {
      const plant = new MapPosition(parseInt(r), parseInt(c));

      const type = plantTypeAt(plant, map);

      const region = new Region(type);

      const queue = [plant];

      while (queue.length > 0) {
        const current = queue.shift() as MapPosition;
        if (seen.has(current.stringify())) continue;
        seen.add(current.stringify());

        region.members.push(current);

        const candidates = [
          new MapPosition(current.row - 1, current.col),
          new MapPosition(current.row + 1, current.col),
          new MapPosition(current.row, current.col - 1),
          new MapPosition(current.row, current.col + 1),
        ].filter((p) => {
          return (
            plantTypeAt(p, map) === type &&
            !seen.has(p.stringify()) &&
            !queue.some((q) => q.stringify() === p.stringify())
          );
        });

        queue.push(...candidates);
      }

      if (region.members.length > 0) {
        regions.push(region.normalize());
      }
    }
  }

  const totalPrice = regions.reduce((acc, r) => acc + r.price2, 0);

  console.log(totalPrice);
};
