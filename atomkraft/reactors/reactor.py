import time
from datetime import datetime, timedelta

import pytest
from modelator.pytest.decorators import step
from munch import Munch
from terra_proto.cosmos.staking.v1beta1 import AuthorizationType
from terra_sdk.client.lcd import LCDClient
from terra_sdk.client.lcd.api.tx import CreateTxOptions
from terra_sdk.core import Coin
from terra_sdk.core.authz import (
    AuthorizationGrant,
    GenericAuthorization,
    MsgExecAuthorized,
    MsgGrantAuthorization,
    MsgRevokeAuthorization,
    SendAuthorization,
    StakeAuthorization,
)
from terra_sdk.core.authz.data import StakeAuthorizationValidators
from terra_sdk.core.bank import MsgSend
from terra_sdk.core.fee import Fee
from terra_sdk.core.staking import MsgBeginRedelegate, MsgDelegate, MsgUndelegate
from terra_sdk.key.mnemonic import MnemonicKey
from terra_sdk.util.converter import to_isoformat

MSG_TYPE = {
    "msg_delegate": MsgDelegate.type_url,
    "msg_send": MsgSend.type_url,
    "msg_undelegate": MsgUndelegate.type_url,
    "msg_redelegate": MsgBeginRedelegate.type_url,
    "msg_alpha": "msg_alpha",
}


STAKE_AUTH_TYPE = {
    "msg_delegate": AuthorizationType.AUTHORIZATION_TYPE_DELEGATE,
    "msg_undelegate": AuthorizationType.AUTHORIZATION_TYPE_UNDELEGATE,
    "msg_redelegate": AuthorizationType.AUTHORIZATION_TYPE_REDELEGATE,
    "msg_alpha": AuthorizationType.AUTHORIZATION_TYPE_UNSPECIFIED,
}

AUTHZ_FAILURE_OUTCOME = {
    "inappropriate_auth",  # INAPPROPRIATE_AUTH
    "inappropriate_auth_generic",  # INAPPROPRIATE_AUTH_GENERIC
    "inappropriate_auth_send",  # INAPPROPRIATE_AUTH_SEND
    "inappropriate_auth_stake_not_allow",  # INAPPROPRIATE_AUTH_STAKE_NOT_ALLOW
    "inappropriate_auth_stake_deny",  # INAPPROPRIATE_AUTH_STAKE_DENY
    "message_not_supported_by_the_authorization",  # INAPPROPRIATE_AUTH_FOR_MESSAGE
    "insufficient_auth_exec",  # INSUFFICIENT_GRANT_EXEC
    "non_existent_auth",  # NONEXISTENT_GRANT_EXEC
    # "expired_grant", #EXPIRED_GRANT
    "tried to execute an expired grant",  # EXPIRED_AUTH_EXEC
    "grant_failed",  #
    "give grant failed: granter and grantee are the same",  # INVALID_GRANTEE_AND_GRANTER
    "give grant failed: invalid spending limit",  # INVALID_SPENDING_LIMIT
    "give grant failed: invalid validator lists",  # INVALID_VALIDATOR_LISTS
    "give grant failed: expiration time is in the past",  # INVALID_EXPIRATION_TIME_IN_PAST
    "revoke_failed",  # REVOKE_FAILED
    # , #
}


@pytest.fixture
def state():
    return Munch()


def status(ok):
    "OK" if ok else "FAIL"


def create_give_grant_auth(
    testnet,
    state,
    grant_message_type,
    authorization_logic,
    spend_limit,
    no_limits,
    allow_list_tla,
    deny_list_tla,
    exp_time,
):
    match authorization_logic:
        case "send":
            authorization = SendAuthorization(spend_limit)
        case "generic":
            msg_type = MSG_TYPE[grant_message_type]
            authorization = GenericAuthorization(msg=msg_type)
        case "stake":
            authorization_type = STAKE_AUTH_TYPE[grant_message_type]
            # INFINITY check for staking type
            if no_limits == "inf":
                max_tokens = None
            else:
                max_tokens = Coin.from_str(spend_limit)

            if allow_list_tla:
                address = [
                    testnet.validators[state.validators.index(tla_id)].address(
                        testnet.prefix
                    )
                    for tla_id in allow_list_tla
                ]
                allow_list = StakeAuthorizationValidators(address)
            else:
                allow_list = None
            if deny_list_tla:
                address = [
                    testnet.validators[state.validators.index(tla_id)].address(
                        testnet.prefix
                    )
                    for tla_id in deny_list_tla
                ]
                deny_list = StakeAuthorizationValidators(address)
            else:
                deny_list = None
            authorization = StakeAuthorization(
                authorization_type, max_tokens, allow_list, deny_list
            )

    match exp_time:
        case "past":
            expiration = datetime.now() - timedelta(seconds=100)
        case "future":
            expiration = datetime.now() + timedelta(seconds=100)
        case "EXPIRE_SOON":
            expiration = datetime.now() + timedelta(seconds=10)
        case "infinite":
            # terra.py client requires a concrete timestamp
            expiration = datetime.now() + timedelta(seconds=10000)

    return AuthorizationGrant(authorization, to_isoformat(expiration))


Identifiers = ["A", "B", "C"]
Validators = ["X", "Y", "Z"]


# Mirel's traces use empty string for init step
@step("")
def init(testnet, state):
    print("Step: init")

    state.accounts = list(Identifiers)
    state.validators = list(Validators)

    # testnet is configured at this step
    testnet.n_account = len(Identifiers)
    testnet.n_validator = len(Validators)
    testnet.verbose = True

    # variables in the initial TLA state can be used to dynamically configure a testnet

    # testnet is started from this point
    testnet.oneshot()
    time.sleep(10)


@step("give grant")
def give_grant(testnet, state, action_taken, outcome_status):
    print("Step: give grant")

    grantee_id = state.accounts.index(action_taken.grant.grantee)
    granter_id = state.accounts.index(action_taken.grant.granter)

    # prepare transaction

    grantee = testnet.accounts[grantee_id].address(testnet.prefix)
    granter = testnet.accounts[granter_id].address(testnet.prefix)

    grant = create_give_grant_auth(
        testnet,
        state,
        action_taken.grant.sdk_message_type,
        action_taken.grant_payload.authorization_logic,
        f"{int(action_taken.grant_payload.limit)}{testnet.denom}",
        action_taken.grant_payload.special_value,
        action_taken.grant_payload.allow_list,
        action_taken.grant_payload.deny_list,
        action_taken.grant.expiration_time,
    )

    msg = MsgGrantAuthorization(granter, grantee, grant)

    rest_endpoint = testnet.get_validator_port(0, "lcd")
    lcdclient = LCDClient(
        url=rest_endpoint,
        chain_id=testnet.chain_id,
        gas_prices=f"10{testnet.denom}",
        gas_adjustment=0.1,
    )

    granter_wallet = lcdclient.wallet(
        MnemonicKey(
            mnemonic=testnet.accounts[granter_id].mnemonic,
            coin_type=testnet.coin_type,
            prefix=testnet.prefix,
        )
    )

    tx = granter_wallet.create_and_sign_tx(
        CreateTxOptions(msgs=[msg], fee=Fee(20000000, f"2000000{testnet.denom}"))
    )

    result = lcdclient.tx.broadcast(tx)

    print("[MSG]", msg)
    print("[RES]", result)
    print(
        "[CMP] Expected {} ({}), Got {}".format(
            status(outcome_status not in AUTHZ_FAILURE_OUTCOME),
            outcome_status,
            status(result.code == 0),
        )
    )

    time.sleep(5)


@step("expire grant")
def expire_grant(testnet, state, action_taken, active_grants, outcome_status):
    print("Step: expire grant")

    grantee_id = state.accounts.index(action_taken.grant.grantee)
    granter_id = state.accounts.index(action_taken.grant.granter)

    if len(active_grants) > 0:
        for active_grant, active_grant_detail in active_grants.items():
            if (
                active_grant.grantee == action_taken.grant.grantee
                and active_grant.granter == action_taken.grant.granter
                and active_grant.sdk_message_type == action_taken.grant.sdk_message_type
            ):
                exp_grant_authorization_logic = active_grant_detail.authorization_logic
                exp_grant_allow_list = active_grant_detail.allow_list
                exp_grant_deny_list = active_grant_detail.deny_list
                exp_grant_spend_limit = (
                    f"{int(active_grant_detail.limit)}{testnet.denom}"
                )
                exp_grant_no_limits = active_grants[active_grant].special_value
                break

        # prepare transaction

        grantee = testnet.accounts[grantee_id].address(testnet.prefix)
        granter = testnet.accounts[granter_id].address(testnet.prefix)

        grant = create_give_grant_auth(
            testnet,
            state,
            action_taken.grant.sdk_message_type,
            exp_grant_authorization_logic,
            exp_grant_spend_limit,
            exp_grant_no_limits,
            exp_grant_allow_list,
            exp_grant_deny_list,
            "EXPIRE_SOON",
        )

        msg = MsgGrantAuthorization(granter, grantee, grant)

        rest_endpoint = testnet.get_validator_port(0, "lcd")
        lcdclient = LCDClient(
            url=rest_endpoint,
            chain_id=testnet.chain_id,
            gas_prices=f"10{testnet.denom}",
            gas_adjustment=0.1,
        )

        granter_wallet = lcdclient.wallet(
            MnemonicKey(
                mnemonic=testnet.accounts[granter_id].mnemonic,
                coin_type=testnet.coin_type,
                prefix=testnet.prefix,
            )
        )

        tx = granter_wallet.create_and_sign_tx(
            CreateTxOptions(msgs=[msg], fee=Fee(20000000, f"2000000{testnet.denom}"))
        )

        result = lcdclient.tx.broadcast(tx)

        print("[MSG]", msg)
        print("[RES]", result)
        print(
            "[CMP] Expected {} ({}), Got {}".format(
                status(outcome_status not in AUTHZ_FAILURE_OUTCOME),
                outcome_status,
                status(result.code == 0),
            )
        )

        time.sleep(5)

        # this sleep is to wait for the grant expiration
        time.sleep(10)


@step("revoke grant")
def revoke_grant(testnet, state, action_taken, outcome_status):
    print("Step: revoke grant")

    grantee_id = state.accounts.index(action_taken.grant.grantee)
    granter_id = state.accounts.index(action_taken.grant.granter)

    # prepare transaction

    grantee = testnet.accounts[grantee_id].address(testnet.prefix)
    granter = testnet.accounts[granter_id].address(testnet.prefix)

    msg_type_url = MSG_TYPE[action_taken.grant.sdk_message_type]
    msg = MsgRevokeAuthorization(granter, grantee, msg_type_url)

    rest_endpoint = testnet.get_validator_port(0, "lcd")
    lcdclient = LCDClient(
        url=rest_endpoint,
        chain_id=testnet.chain_id,
        gas_prices=f"10{testnet.denom}",
        gas_adjustment=0.1,
    )

    granter_wallet = lcdclient.wallet(
        MnemonicKey(
            mnemonic=testnet.accounts[granter_id].mnemonic,
            coin_type=testnet.coin_type,
            prefix=testnet.prefix,
        )
    )

    tx = granter_wallet.create_and_sign_tx(
        CreateTxOptions(msgs=[msg], fee=Fee(20000000, f"2000000{testnet.denom}"))
    )

    result = lcdclient.tx.broadcast(tx)

    print("[MSG]", msg)
    print("[RES]", result)
    print(
        "[CMP] Expected {} ({}), Got {}".format(
            status(outcome_status not in AUTHZ_FAILURE_OUTCOME),
            outcome_status,
            status(result.code == 0),
        )
    )

    time.sleep(5)


@step("execute grant")
def execute_grant(testnet, state, action_taken, outcome_status):
    print("Step: execute grant")

    grantee_id = state.accounts.index(action_taken.grant.grantee)
    granter_id = state.accounts.index(action_taken.grant.granter)
    for e in range(len(state.accounts)):
        if e not in [grantee_id, granter_id]:
            receiver_id = e

    validator_id = 0
    for e in range(len(state.validators)):
        if e not in [validator_id]:
            dst_validator_id = e

    # prepare transaction

    grantee = testnet.accounts[grantee_id].address(testnet.prefix)
    granter = testnet.accounts[granter_id].address(testnet.prefix)
    receiver = testnet.accounts[receiver_id].address(testnet.prefix)

    validator = testnet.validators[validator_id].validator_address(testnet.prefix)
    dst_validator = testnet.validators[dst_validator_id].validator_address(
        testnet.prefix
    )

    amount = f"{int(action_taken.exec_message.amount)}{testnet.denom}"

    match action_taken.exec_message.message_type:
        case "msg_delegate":
            exec_msg = MsgDelegate(granter, validator, amount)
        case "msg_send":
            exec_msg = MsgSend(granter, receiver, amount)
        case "msg_undelegate":
            exec_msg = MsgUndelegate(granter, validator, amount)
        case "msg_redelegate":
            exec_msg = MsgBeginRedelegate(granter, validator, dst_validator, amount)
        case "msg_alpha":
            pass

    msg = MsgExecAuthorized(grantee, [exec_msg])

    rest_endpoint = testnet.get_validator_port(0, "lcd")
    lcdclient = LCDClient(
        url=rest_endpoint,
        chain_id=testnet.chain_id,
        gas_prices=f"10{testnet.denom}",
        gas_adjustment=0.1,
    )

    grantee_wallet = lcdclient.wallet(
        MnemonicKey(
            mnemonic=testnet.accounts[grantee_id].mnemonic,
            coin_type=testnet.coin_type,
            prefix=testnet.prefix,
        )
    )

    tx = grantee_wallet.create_and_sign_tx(
        CreateTxOptions(msgs=[msg], fee=Fee(20000000, f"2000000{testnet.denom}"))
    )

    result = lcdclient.tx.broadcast(tx)

    print("[MSG]", msg)
    print("[RES]", result)
    print(
        "[CMP] Expected {} ({}), Got {}".format(
            status(outcome_status not in AUTHZ_FAILURE_OUTCOME),
            outcome_status,
            status(result.code == 0),
        )
    )

    time.sleep(5)
