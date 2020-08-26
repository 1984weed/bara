export type ArrayAutoTestCase = { min: number; max: number; count: number; isUnique: boolean }

export const getArrayAutoTestCaseDef = (): ArrayAutoTestCase => {
    return { min: 0, max: 100, count: 1, isUnique: false }
}
