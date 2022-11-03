import asyncio
import logging
from contextlib import closing

from atomkraft.chain import Testnet
from atomkraft.chain.utils import TmEventSubscribe
from modelator.pytest.decorators import step
from terra_proto.cosmos.bank.v1beta1 import QueryStub
from terra_sdk.core.bank import MsgMultiSend


@step("init")
def init(testnet: Testnet, action):
    logging.info("Step: Init")
    testnet.set_accounts(list(action.balances.keys()))
    testnet.set_account_balances(action.balances)
    testnet.verbose = True
    testnet.oneshot()
    with TmEventSubscribe({"tm.event": "NewBlock"}):
        logging.info("\tBlockhain is producing blocks")


@step("send")
def transfer(testnet: Testnet, action, balances, outcome):
    logging.info("Step: Transfer")
    coins = action.coins
    sender_id = action.sender
    receiver_id = action.receiver

    sender_addr = testnet.acc_addr(sender_id)
    receiver_addr = testnet.acc_addr(receiver_id)

    coins_str = ",".join("{amount}{denom}".format(**e) for e in coins)

    src = [{"address": sender_addr, "coins": coins_str}]
    dst = [{"address": receiver_addr, "coins": coins_str}]

    msg = MsgMultiSend(inputs=src, outputs=dst)

    logging.info(f"\tSender:    {sender_id} ({sender_addr})")
    logging.info(f"\tReceiver:  {receiver_id} ({receiver_addr})")
    logging.info(f"\tAmount:    {coins_str}")
    logging.info(f"\tExpected:  {outcome}")

    try:
        result = testnet.broadcast_transaction(sender_id, msg)
        if result.code == 0:
            result = "SUCCESS"
        else:
            result = f"ERROR {result.code}: {result.raw_log}"
    except Exception as e:
        result = f"EXCEPTION: {e}"

    logging.info(f"\tStatus:    {result}")

    balances_mismatch = ""
    with closing(testnet.get_grpc_channel()) as channel:
        stub = QueryStub(channel)
        for e_acc in balances:
            query_result = asyncio.run(
                stub.all_balances(address=testnet.acc_addr(e_acc))
            )
            observed = {e.denom: int(e.amount) for e in query_result.balances}
            for e_denom in balances[e_acc]:
                bal = balances[e_acc][e_denom]
                obs_bal = observed.get(e_denom, 0)
                if bal != obs_bal:
                    balances_mismatch += (
                        f"\n\texpected {e_acc}[{e_denom}] = {bal}, got {obs_bal}"
                    )

    if (outcome == "SUCCESS" and result != "SUCCESS") or (
        outcome != "SUCCESS" and result == "SUCCESS"
    ):
        logging.error(f"\tExpected {outcome}, got {result}{balances_mismatch}")
        raise RuntimeError(f"Expected {outcome}, got {result}{balances_mismatch}")
