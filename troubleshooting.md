
# Troubleshooting

## "Empty" Msg responses

If you get an "empty" response from `dualityd tx [module] [Msg]` commands like this:
```
code: 0
codespace: ""
data: ""
events: []
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 398D0FFA3B27E59DEEBF48DF7FA2535114002264EC3950E92A5CFC8C94EF3216
```
look up the resulting txhash using `dualityd query tx [txhash]`.

The txhash query result will often contain the real tx failure `code` and more details in `raw_log`.
