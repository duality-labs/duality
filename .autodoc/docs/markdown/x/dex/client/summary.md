[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x/dex/client)

The `.autodoc/docs/json/x/dex/client` folder contains the `cli` package, which provides a set of command-line interface (CLI) commands for interacting with the Duality decentralized exchange (DEX) module. These commands enable users to perform various actions such as querying the DEX state, placing and canceling limit orders, depositing and withdrawing tokens, and performing token swaps.

For instance, the `CmdPlaceLimitOrder` command allows users to place a limit order on the DEX by specifying the tokens involved, the tick index, the amount of input tokens, the order type, and the expiration time. Users can execute this command as follows:

```bash
dualitycli place-limit-order bob tokenA tokenB -10 1000 GOOD_TIL_CANCELLED '01/01/2022 12:00:00' --from alice
```

The `CmdCancelLimitOrder` command allows users to cancel a limit order by providing the unique identifier of the limit order (tranche-key). Users can execute this command as follows:

```bash
dualitycli cancel-limit-order TRANCHEKEY123 --from alice
```

The `CmdDeposit` and `CmdWithdrawal` commands allow users to deposit and withdraw tokens from the DEX, respectively. Users can execute these commands as follows:

```bash
dualitycli deposit alice tokenA tokenB 100,50 [-10,5] 1,1 --from alice
dualitycli withdrawal alice tokenA tokenB 100,50 [-10,5] 1,1 --from alice
```

The `CmdSwap` and `CmdMultiHopSwap` commands allow users to perform token swaps on the DEX, either directly or through multiple hops. Users can execute these commands as follows:

```bash
dualitycli swap alice 100 tokenA tokenB --from alice
dualitycli multi-hop-swap alice "tokenA/tokenB,tokenB/tokenC" 100 --from alice
```

The `query` subcommands allow users to retrieve information about the DEX state, such as user positions, limit orders, and pool reserves. Users can execute these commands as follows:

```bash
dualitycli query dex show-user-positions alice
dualitycli query dex list-limit-order-tranche tokenA<>tokenB tokenA --page=1 --limit=10
dualitycli query dex show-pool-reserves tokenA<>tokenB [-5] tokenA 1
```

In summary, the `cli` package in the `.autodoc/docs/json/x/dex/client` folder offers a comprehensive set of CLI commands for interacting with the Duality DEX module. These commands allow users to manage their assets, execute trades, and query the DEX state. Developers can use these commands in conjunction with other parts of the Duality project to build more complex trading strategies and applications on top of the DEX module.
