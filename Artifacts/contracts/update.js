let args = process.argv.slice(2);
let copy_from = args[0];
let network_id = args[1];

let path = require('path');
let fs = require('fs');
let package_path = __dirname;
let contracts_path = __dirname + '/' + copy_from;
let package_art = fs.readdirSync(package_path);
for (let n of package_art){
    let n_parts = n.split('.');
    if (n_parts.length > 1 && n_parts[1] == 'json'){
        let c_art = require(path.join(contracts_path, n));
        let p_art = require(path.join(package_path, n));
        if (c_art.hasOwnProperty('networks') &&
			c_art.networks.hasOwnProperty(network_id) &&
			p_art.hasOwnProperty('networks') &&
			p_art.networks.hasOwnProperty(network_id)
        )
        {
            let address_42 = 	c_art.networks['42'].address;
            p_art.networks['42']['address'] = address_42;
            if (!fs.existsSync(__dirname + '/../contracts_temp')) {
                fs.mkdirSync(__dirname + '/../contracts_temp');
            }
            fs.writeFileSync(__dirname + '/../contracts_temp/' + n, JSON.stringify(p_art, null, 2));

        }

    }
}
