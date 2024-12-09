const demoInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`;

class AntennaMap {
  map: string[][];
  rows: number;
  cols: number;
  antennas: Record<string, Array<[number, number]>>;
  antinodes: Set<string>;

  constructor(mapInput: string) {
    this.map = mapInput.split("\n").map((row) => row.split(""));
    this.rows = this.map.length;
    this.cols = this.map[0].length;
    this.antennas = this.findAntennas();
    this.antinodes = new Set();
  }

  findAntennas() {
    const antennas: Record<string, Array<[number, number]>> = {};
    for (let r = 0; r < this.rows; r++) {
      for (let c = 0; c < this.cols; c++) {
        const freq = this.map[r][c];
        if (freq !== ".") {
          if (!antennas[freq]) antennas[freq] = [];
          antennas[freq].push([r, c]);
        }
      }
    }
    return antennas;
  }

  isValidCoord(r: number, c: number) {
    return r >= 0 && r < this.rows && c >= 0 && c < this.cols;
  }

  findAntinodes() {
    for (const freq in this.antennas) {
      const frequencyAntennas = this.antennas[freq];

      for (let i = 0; i < frequencyAntennas.length; i++) {
        for (let j = i + 1; j < frequencyAntennas.length; j++) {
          const [r1, c1] = frequencyAntennas[i];
          const [r2, c2] = frequencyAntennas[j];

          const deltaR = r2 - r1;
          const deltaC = c2 - c1;

          for (let k = 1; k <= Math.max(this.rows, this.cols); k++) {
            const antinodeR1 = r1 + k * deltaR;
            const antinodeC1 = c1 + k * deltaC;
            const antinodeR2 = r1 - k * deltaR;
            const antinodeC2 = c1 - k * deltaC;

            if (this.isValidCoord(antinodeR1, antinodeC1)) {
              this.antinodes.add(`${antinodeR1},${antinodeC1}`);
            }
            if (this.isValidCoord(antinodeR2, antinodeC2)) {
              this.antinodes.add(`${antinodeR2},${antinodeC2}`);
            }
          }

          this.antinodes.add(`${r1},${c1}`);
          this.antinodes.add(`${r2},${c2}`);
        }
      }
    }
  }

  calculateAntinodeImpact() {
    this.findAntinodes();
    return this.antinodes.size;
  }
}

export default () => {
  // const input = Deno.readTextFileSync(import.meta.dirname + '/input')
  const input = demoInput;

  const antennaMap = new AntennaMap(input);

  console.log(antennaMap.calculateAntinodeImpact());
};
