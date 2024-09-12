import { buildModule } from '@nomicfoundation/hardhat-ignition/modules'

export default buildModule('Storage', (m) => {
  const storage = m.contract('Storage')

  return { storage }
})
