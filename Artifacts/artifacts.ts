export const Artifacts: { [index: string]: any } = {
    'ARBITER': require('./contracts/Arbiter.json'),
    'BONDAGE': require('./contracts/Bondage.json'),
    'DISPATCH': require('./contracts/Dispatch.json'),
    'REGISTRY': require('./contracts/Registry.json'),
    'CurrentCost': require('./contracts/CurrentCost.json'),
    'ZAP_TOKEN': require('./contracts/ZapToken.json'),
    'Client1': require('./contracts/Client1.json'),
    'Client2': require('./contracts/Client2.json'),
    'Client3': require('./contracts/Client3.json'),
    'Client4': require('./contracts/Client4.json'),
    'ZAPCOORDINATOR': require('./contracts/ZapCoordinator.json'),
    'TOKENDOTFACTORY': require('./contracts/TokenDotFactory.json')
};

console.log(Artifacts['DISPATCH'])
