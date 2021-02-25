{} let path = require('path');
let fs = require('fs');
let Artifacts = fs.readdirSync(__dirname);
for (let n of Artifacts){
    if (n == 'index.ts') continue;
    let art = require(path.join(__dirname, n));
    let artifact = {
        contract_name: art.contract_name,
        abi: art.abi,
        networks: art.networks
    };
    if (!fs.existsSync(__dirname + '_temp'))
        fs.mkdirSync(__dirname + '_temp');
    fs.writeFileSync(__dirname + '_temp/' + n, JSON.stringify(artifact, null, 2));
}
