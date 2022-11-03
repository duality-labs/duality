-------------------------------- MODULE Grants ---------------------------------
(******************************************************************************)

(******************************************************************************)
CONSTANTS
    \* @typeAlias: ACCOUNT = Str;
    \* @type: Set(ACCOUNT);
    Accounts,

    \* @typeAlias: VALIDATOR = Str;
    \* @type: Set(VALIDATOR);
    Validators
--------------------------------------------------------------------------------
CONSTANTS 
    \* @typeAlias: COINS = Int;
    \* @type: Set(COINS);
    Coins,
    \* @type: COINS;
    NoMaxCoins

--------------------------------------------------------------------------------
Generic == INSTANCE GenericAuthorization
Send == INSTANCE SendAuthorization
Stake == INSTANCE StakeAuthorization

MsgTypeUrls == 
    Generic!MsgTypeUrls \cup 
    Send!MsgTypeUrls \cup 
    Stake!MsgTypeUrls

Authorization == 
    Generic!Authorization \cup 
    Send!Authorization \cup 
    Stake!Authorization

\* @typeAlias: SDK_MSG_CONTENT = [amount: COINS, fromAddress: ACCOUNT, toAddress: ACCOUNT, delegatorAddress: VALIDATOR, validatorAddress: VALIDATOR, validatorSrcAddress: VALIDATOR, validatorSrcAddress: VALIDATOR, validatorDstAddress: VALIDATOR, typeUrl: MSG_TYPE_URL];
SdkMsgContent == 
    Generic!SdkMsgContent \cup 
    Send!SdkMsgContent \cup 
    Stake!SdkMsgContent

--------------------------------------------------------------------------------
MsgTypeURL(auth) ==
    CASE auth.authorizationType \in Generic!MsgTypeUrls -> 
        Generic!MsgTypeURL(auth)
      [] auth.authorizationType \in Send!MsgTypeUrls -> 
        Send!MsgTypeURL(auth)
      [] auth.authorizationType \in Stake!MsgTypeUrls -> 
        Stake!MsgTypeURL(auth)

Accept(auth, msg) ==
    CASE msg.typeUrl \in Generic!MsgTypeUrls -> 
        Generic!Accept(auth, msg)
      [] msg.typeUrl \in Send!MsgTypeUrls -> 
        Send!Accept(auth, msg)
      [] msg.typeUrl \in Stake!MsgTypeUrls -> 
        Stake!Accept(auth, msg)

--------------------------------------------------------------------------------
(******************************************************************************)
(* A grant is an allowance to execute a Msg by the grantee on behalf of the
granter. Grants are identified by combining granter address (the address 
bytes of the granter), grantee address (the address bytes of the grantee) 
and Authorization type (its type URL). Hence we only allow one grant for 
the (granter, grantee, Authorization) triple. *)
(******************************************************************************)
\* @typeAlias: GRANT_ID = [grantee: ACCOUNT, granter: ACCOUNT, msgTypeUrl: MSG_TYPE_URL];
\* @type: Set(GRANT_ID);
GrantIds == [
    granter: Accounts,
    grantee: Accounts,
    msgTypeUrl: MsgTypeUrls
]

\* @type: (GRANT_ID) => Bool;
IsValid(g) == g.granter # g.grantee

\* @type: Set(GRANT_ID);
ValidGrantIds == { g \in GrantIds: IsValid(g) }

\* Time when the grant will expire with respect to the moment when the related
\* event happens. If "none", then the grant doesn't have an expiration time and 
\* other conditions in the authorization may apply to invalidate it.
\* @typeAlias: EXPIRATION_TIME = Str;
\* @type: Set(EXPIRATION_TIME);
ExpirationTimes == {"past", "future", "none"}

\* Grant gives permissions to execute the provide method with expiration time.
\* https://github.com/cosmos/cosmos-sdk/blob/c1b6ace7d542925b526cf3eef6df38a206eab8d8/x/authz/authz.pb.go#L74
\* @typeAlias: GRANT = [authorization: AUTH, expirationTime: EXPIRATION_TIME];
\* @type: Set(GRANT);
Grants == [
    authorization: Authorization,
    expirationTime: ExpirationTimes
]

\* @type: AUTH;
NoAuthorization == [ type |-> "NoAuthorization" ]

\* @type: GRANT;
NoGrant == [ authorization |-> NoAuthorization, expirationTime |-> "none" ]

================================================================================
Created by Hernán Vanzetto on 10 August 2022
