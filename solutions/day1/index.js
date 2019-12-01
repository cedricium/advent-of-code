const fs = require('fs')
const INPUT_PATH = 'input.txt'

/**
 * Algorithm used to calculate the fuel required for the given module:
 * `divide the mass by 3, round down, then subtract 2`
 * @param {Number} mass mass of the current module
 * @returns {Number} computed fuel required for the given module's mass
 */
const computeFuelPerModule = mass => {
  return Math.floor(mass / 3) - 2
}

const run = () => {
  const moduleMasses = fs
    .readFileSync(INPUT_PATH, 'utf-8')
    .split('\n')
    .filter(Boolean)
    .map(numStr => Number(numStr))

  const totalFuel = moduleMasses.reduce(
    (acc, curr) => computeFuelPerModule(curr) + acc,
    0,
  )

  console.log(totalFuel)
}
run()
