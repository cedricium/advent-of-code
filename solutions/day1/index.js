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

const computeTotalFuelRequired = (fuel, runningTotal = 0) => {
  const totalFuel = computeFuelPerModule(fuel)
  if (totalFuel > 0) {
    return computeTotalFuelRequired(totalFuel, (runningTotal += totalFuel))
  } else {
    return runningTotal
  }
}

const run = () => {
  const moduleMasses = fs
    .readFileSync(INPUT_PATH, 'utf-8')
    .split('\n')
    .filter(Boolean)
    .map(numStr => Number(numStr))

  const totalFuelForModules = moduleMasses.reduce(
    (acc, curr) => computeFuelPerModule(curr) + acc,
    0,
  )

  const totalFuelRequired = moduleMasses.reduce(
    (acc, curr) => computeTotalFuelRequired(curr) + acc,
    0,
  )

  console.log(`[Part 1] fuel required for all modules: ${totalFuelForModules}`)
  console.log(`[Part 2] total fuel required: ${totalFuelRequired}`)
}
run()
