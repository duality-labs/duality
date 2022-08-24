
module.exports = function(cosmosGenesisPath, igniteGenesisPath) {
    const cosmosGenesis = require(cosmosGenesisPath);
    const igniteGenesis = require(igniteGenesisPath);
    require('fs').writeFileSync(cosmosGenesisPath, JSON.stringify(
        Object.assign({}, cosmosGenesis, {
            app_state: Object.assign({}, cosmosGenesis.app_state, {
                dex: igniteGenesis.app_state.dex,
                ibc: igniteGenesis.app_state.ibc,
                monitoringp: igniteGenesis.app_state.monitoringp,
                router: igniteGenesis.app_state.router,
                transfer: igniteGenesis.app_state.transfer,
            }),
        }),
        null,
        2
    ));
}
