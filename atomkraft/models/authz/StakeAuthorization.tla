------------------------ MODULE StakeAuthorization -----------------------------
(******************************************************************************)
(* StakeAuthorization implements the Authorization interface for messages in the
staking module. It takes an AuthorizationType to specify whether you want to
authorise delegating, undelegating or redelegating (i.e. these have to be
authorised seperately). It also takes a MaxTokens that keeps track of a limit to
the amount of tokens that can be delegated/undelegated/redelegated. If left
empty, the amount is unlimited. Additionally, this Msg takes an AllowList and a
DenyList, which allows you to select which validators you allow grantees to
stake with. *)
(******************************************************************************)
LOCAL INSTANCE MsgTypes
LOCAL INSTANCE Integers

CONSTANT
    \* @typeAlias: VALIDATOR = Str;
    \* @type: Set(VALIDATOR);
    Validators,
    \* @typeAlias: COINS = Int;
    \* @type: Set(COINS);
    Coins,
    \* @type: COINS;
    NoMaxCoins

\* We want our model of Coins to include some negative number.
ASSUME \E min, max \in Int: 
    /\ min < 0
    /\ max > 0 
    /\ Coins = min .. max

\* @type: COINS;
ASSUME NoMaxCoins \in Int /\ NoMaxCoins \notin Coins

\* @typeAlias: MSG_TYPE_URL = Str;
\* @typeAlias: SDK_MSG_CONTENT = [amount: COINS, delegatorAddress: VALIDATOR, validatorAddress: VALIDATOR, validatorSrcAddress: VALIDATOR, validatorSrcAddress: VALIDATOR, validatorDstAddress: VALIDATOR, typeUrl: MSG_TYPE_URL];

\* MsgDelegate defines a SDK message for performing a delegation of coins from a
\* delegator to a validator.
\* https://github.com/cosmos/cosmos-sdk/blob/f848e4300a8a6036a4dbfb628c7a9e7874a8e6db/x/staking/types/tx.pb.go#L205
\* @type: Set(SDK_MSG_CONTENT);
MsgDelegate == [
    typeUrl: { DELEGATE_TYPE_URL },
    delegatorAddress: Validators,
    validatorAddress: Validators,
    amount: Coins
]

\* MsgUndelegate defines a SDK message for performing an undelegation from a
\* delegate and a validator.
\* https://github.com/cosmos/cosmos-sdk/blob/f848e4300a8a6036a4dbfb628c7a9e7874a8e6db/x/staking/types/tx.pb.go#L370
\* @type: Set(SDK_MSG_CONTENT);
MsgUndelegate == [
    typeUrl: { UNDELEGATE_TYPE_URL},
    delegatorAddress: Validators,
    validatorAddress: Validators,
    amount: Coins
]

\* MsgBeginRedelegate defines a SDK message for performing a redelegation of
\* coins from a delegator and source validator to a destination validator.
\* https://github.com/cosmos/cosmos-sdk/blob/f848e4300a8a6036a4dbfb628c7a9e7874a8e6db/x/staking/types/tx.pb.go#L283
\* @type: Set(SDK_MSG_CONTENT);
MsgBeginRedelegate == [
    typeUrl: { BEGIN_REDELEGATE_TYPE_URL },
    delegatorAddress: Validators,
    validatorSrcAddress: Validators,
    validatorDstAddress: Validators,
    amount: Coins
]

\* @type: Set(SDK_MSG_CONTENT);
SdkMsgContent == MsgDelegate \cup MsgUndelegate \cup MsgBeginRedelegate

\* Types of messages allowed to be granted permission
\* @type: Set(MSG_TYPE_URL);
MsgTypeUrls == { m.typeUrl: m \in SdkMsgContent }

--------------------------------------------------------------------------------
\* The authorization for delegate/undelegate/redelegate.
\* Issue for bug when deny list is not empty: https://github.com/cosmos/cosmos-sdk/issues/11391
\* https://github.com/cosmos/cosmos-sdk/blob/55054282d2df794d9a5fe2599ea25473379ebc3d/x/staking/types/authz.go#L16
\* @typeAlias: AUTH = [maxTokens: COINS, validators: Set(VALIDATOR), allow: Bool, authorizationType: MSG_TYPE_URL];
\* @type: Set(AUTH);
Authorization == [  
    \* Specifies the maximum amount of tokens can be delegate to a validator. If
    \* it is empty, there is no spend limit and any amount of coins can be
    \* delegated.
    maxTokens: Coins \cup {NoMaxCoins},

    \* A set of validator addresses to whom delegation of tokens is either
    \* allowed or denied.
    validators: SUBSET Validators,

    \* Extra field not present in the code.
    \* If TRUE, validators is a list of allowed addresses. 
    \* If FALSE, validators is a list of address that should not be granted delegation.
    allow: BOOLEAN,

    \* Specifies one of three authorization types.
    authorizationType: MsgTypeUrls
]

\* Apalache does not like the expression:
\*      [auth EXCEPT !.maxTokens = auth.maxTokens - amount]
\* Error message:         
\*     The specification is malformed: An updated record has more fields than its
\*     declared type: A record with the inferred type `[allow: Bool,
\*     authorizationType: Str, maxTokens: Int, type: Str, validators: Set(Str)]`
\*     has been updated with the key `validators` in an `EXCEPT` expression and the
\*     updated record has more fields than are specified in its type annotation.
\*     For details see
\*     https://apalache.informal.systems/docs/apalache/known-issues.html#updating-records-with-excess-fields
\* @type: (AUTH, COINS) => AUTH;
UpdateMaxTokens(auth, maxTokens) == [
    maxTokens |-> maxTokens, 
    validators |-> auth.validators, 
    allow |-> auth.allow, 
    authorizationType |-> auth.authorizationType
]

--------------------------------------------------------------------------------
\* https://github.com/cosmos/cosmos-sdk/blob/55054282d2df794d9a5fe2599ea25473379ebc3d/x/staking/types/authz.go#L38
\* @type: (AUTH) => MSG_TYPE_URL;
MsgTypeURL(auth) ==
    auth.authorizationType

\* @type: (SDK_MSG_CONTENT) => VALIDATOR;
ValidatorAddressOf(msg) ==
    CASE msg.typeUrl = DELEGATE_TYPE_URL -> msg.validatorAddress 
      [] msg.typeUrl = UNDELEGATE_TYPE_URL -> msg.validatorAddress 
      [] msg.typeUrl = BEGIN_REDELEGATE_TYPE_URL -> msg.validatorDstAddress 

\* https://github.com/cosmos/cosmos-sdk/blob/55054282d2df794d9a5fe2599ea25473379ebc3d/x/staking/types/authz.go#L58
\* @typeAlias: ACCEPT_RESPONSE = [accept: Bool, delete: Bool, updated: AUTH, error: Str];
\* @type: (AUTH, SDK_MSG_CONTENT) => ACCEPT_RESPONSE;
Accept(auth, msg) == 
    LET 
        \* @type: COINS;
        amount == msg.amount
        \* @type: VALIDATOR;
        validatorAddress == ValidatorAddressOf(msg)
    IN
    IF auth.allow /\ validatorAddress \notin auth.validators THEN
        [accept |-> FALSE, delete |-> FALSE, updated |-> auth, error |-> "validator-not-allowed"]
    ELSE IF ~ auth.allow /\ validatorAddress \in auth.validators THEN
        [accept |-> FALSE, delete |-> FALSE, updated |-> auth, error |-> "validator-denied"]
    ELSE IF auth.maxTokens = NoMaxCoins THEN 
        [ accept |-> TRUE, delete |-> FALSE, updated |-> auth, error |-> "none" ]
    ELSE [ 
        accept |-> amount <= auth.maxTokens, 
        delete |-> amount = auth.maxTokens, 
        updated |-> IF amount < auth.maxTokens
            THEN UpdateMaxTokens(auth, auth.maxTokens - amount)
            ELSE auth,
        error |-> IF amount <= auth.maxTokens THEN "none" ELSE "insufficient-amount"
    ]

================================================================================
Created by Hernán Vanzetto on 10 August 2022
