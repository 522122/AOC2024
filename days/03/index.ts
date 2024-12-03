const demoInput = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

export default () => {
    const input = Deno.readTextFileSync(import.meta.dirname + '/input')
    // const input = demoInput

    const regex = /mul\((\d+),(\d+)\)|do\(\)|don't\(\)/g

    const code = [...input.matchAll(regex)]

    code.sort((a, b) => {
        return a.index - b.index
    })
    
    let ex = true
    let sum2 = 0

    for (const inst of code) {
        if (inst[0] === 'do()') {
            ex = true
        } else if (inst[0] === 'don\'t()') {
            ex = false
        } else if (ex) {
            const a = parseInt(inst[1])
            const b = parseInt(inst[2])
            const result = a * b
            sum2 += result
        }
    }

    let sum = 0

    for (const match of code) {
        if (match[0].startsWith('mul')) {
            const a = parseInt(match[1])
            const b = parseInt(match[2])
            const result = a * b
            sum += result
        }
    }

    console.log('PT1:', sum)
    console.log('PT2:', sum2)
}