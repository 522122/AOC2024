const demoInput = `125 17`;

const splitNumber = (num) => {
  const str = num.toString();
  const mid = Math.floor(str.length / 2);
  return [
    parseInt(str.slice(0, mid), 10) || 0,
    parseInt(str.slice(mid), 10) || 0,
  ];
};

const blinkStones = (initialStones, blinks) => {
  const cache = new Map();

  const processStone = (stone) => {
    if (cache.has(stone)) {
      return cache.get(stone);
    }

    let result;
    if (stone === 0) {
      result = [1];
    } else if (stone.toString().length % 2 === 0) {
      const [left, right] = splitNumber(stone);
      result = [left || 0, right || 0];
    } else {
      result = [stone * 2024];
    }

    cache.set(stone, result);
    return result;
  };

  let freqMap = new Map();
  for (const stone of initialStones) {
    freqMap.set(stone, (freqMap.get(stone) || 0) + 1);
  }

  for (let i = 0; i < blinks; i++) {
    const newFreqMap = new Map();

    for (const [stone, count] of freqMap) {
      const transformedStones = processStone(stone);
      for (const newStone of transformedStones) {
        newFreqMap.set(newStone, (newFreqMap.get(newStone) || 0) + count);
      }
    }

    freqMap = newFreqMap;
  }

  let totalStones = 0;
  for (const count of freqMap.values()) {
    totalStones += count;
  }

  return totalStones;
};

export default () => {
  const input = Deno.readTextFileSync(import.meta.dirname + "/input");
  // const input = demoInput;

  const blinks = 75;

  const totalStones = blinkStones(input.split(" ").map(Number), blinks);

  console.log(blinks, totalStones);
};
