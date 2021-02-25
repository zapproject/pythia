import { join } from 'path';
export class Utils {
    /**
     * @ignore
     * @param {string} buildDir
     * @returns {any}
     */
    static getArtifacts(buildDir: string) {
        let artifacts: any = {};
        artifacts = {
            ARBITER: require(join(buildDir, 'Arbiter.json')),
            BONDAGE: require(join(buildDir, 'Bondage.json')),
            DISPATCH: require(join(buildDir, 'Dispatch.json')),
            REGISTRY: require(join(buildDir, 'Registry.json')),
            CurrentCost: require(join(buildDir, 'CurrentCost.json')),
            PiecewiseLogic: require(join(buildDir, 'PiecewiseLogic.json')),
            ZAP_TOKEN: require(join(buildDir, 'ZapToken.json')),
            Client1: require(join(buildDir, 'Client1.json')),
            Client2: require(join(buildDir, 'Client2.json')),
            Client3: require(join(buildDir, 'Client3.json')),
            Client4: require(join(buildDir, 'Client4.json')),
            ZAPCOORDINATOR: require(join(buildDir, 'ZapCoordinator.json')),
            TOKENDOTFACTORY: require(join(buildDir, 'TokenDotFactory.json'))
        };
        return artifacts;

    }
}
