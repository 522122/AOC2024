const demoInput = `2333133121414131402`;

const compactDisk = (diskMap: string) => {
  const lengths = diskMap.split("").map(Number);

  const blocks: Array<number | "."> = [];
  const blocks2: Array<number | "."> = [];
  let fileId = 0;

  for (let i = 0; i < lengths.length; i++) {
    const length = lengths[i];
    const isFile = i % 2 === 0;

    if (isFile) {
      blocks.push(...Array(length).fill(fileId));
      blocks2.push(...Array(length).fill(fileId));
      fileId++;
    } else {
      blocks.push(...Array(length).fill("."));
      blocks2.push(...Array(length).fill("."));
    }
  }

  let moved = true;
  while (moved) {
    moved = false;

    // Find the leftmost free space
    const freeSpaceIndex: number = blocks.indexOf(".");
    if (freeSpaceIndex === -1) break;

    // Scan from right to left for a block to move
    for (let i = blocks.length - 1; i > freeSpaceIndex; i--) {
      if (blocks[i] !== "." && blocks[freeSpaceIndex] === ".") {
        blocks[freeSpaceIndex] = blocks[i];
        blocks[i] = ".";
        moved = true;
        break;
      }
    }
  }

  moved = true;
  let highestFileId = Math.max(...blocks2.filter((x) => x !== "."));

  const reorder = () => {
    let movedFiles = 0;
    while (true) {
      if (highestFileId <= 0) break;

      const fi = blocks2.indexOf(highestFileId);
      const li = blocks2.lastIndexOf(highestFileId);

      const fileSize = li - fi + 1;

      let lastFreeSpaceIndex = 0;
      while (true) {
        const fs = blocks2.indexOf(".", lastFreeSpaceIndex);
        if (fs === -1) break;
        let ls = fs;

        while (true) {
          const next = blocks2[ls + 1];
          if (next === ".") {
            ls++;
          } else {
            break;
          }
        }

        const freeSpaceSize = ls - fs + 1;

        if (freeSpaceSize >= fileSize && fs < fi) {
          blocks2.splice(fs, fileSize, ...Array(fileSize).fill(highestFileId));
          blocks2.splice(fi, fileSize, ...Array(fileSize).fill("."));

          movedFiles++;
          // console.log(`${highestFileId} moved to ${fs}`)
          break;
        }

        lastFreeSpaceIndex = ls + 1;
      }

      highestFileId--;
    }
    return movedFiles;
  };

  reorder();

  // console.log(blocks2.join(''))

  // Calculate checksum
  const cs1 = blocks.reduce<number>((sum, block, index) => {
    return block !== "." ? sum + index * block : sum;
  }, 0);

  const cs2 = blocks2.reduce<number>((sum, block, index) => {
    return block !== "." ? sum + index * block : sum;
  }, 0);

  return [cs1, cs2];
};

export default () => {
  // const input = Deno.readTextFileSync(import.meta.dirname + '/input')
  const input = demoInput;

  console.log(compactDisk(input));
};
