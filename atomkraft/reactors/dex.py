import logging

from atomkraft.chain import Testnet
from modelator.pytest.decorators import step
from terra_sdk.core.msg import Msg

class DepositMsg(Msg):
    pass

@step("deposit")
def deposit(testnet: Testnet, action):
    logging.info("Step: Deposit")

    coins = action.coins
    creator_id = action.creator
    receiver_id = action.receiver
    token_a = action.token_a
    token_b = action.token_b
    amounts_a = action.amounts_a
    amounts_b = action.amounts_b
    tick_indexes = action.tick_indexes
    fee_indexes = action.fee_indexes

    creator_addr = testnet.acc_addr(creator_id)
    receiver_addr = testnet.acc_addr(receiver_id)

    coins_str = ",".join("{amount}{denom}".format(**e) for e in coins)

    src = [{"address": creator_addr, "coins": coins_str}]
    dst = [{"address": receiver_addr, "coins": coins_str}]

    msg = DepositMsg(inputs=src, outputs=dst)
    # msg = MsgMultiSend(inputs=src, outputs=dst)

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


@step("withdraw")
def withdraw(testnet: Testnet, action):
    logging.info("Step: Withdraw")
    pass

@step("swap")
def swap(testnet: Testnet, action):
    logging.info("Step: Swap")
    pass

@step("place limit order")
def place_limit_order(testnet: Testnet, action):
    logging.info("Step: Place limit order")
    pass

@step("cancel limit order")
def cancel_limit_order(testnet: Testnet, action):
    logging.info("Step: Cancel limit order")
    pass