-------------------------------- MODULE Authz ----------------------------------
(******************************************************************************)
(* 
Official documentation: https://docs.cosmos.network/v0.46/modules/authz/
ADR: https://github.com/cosmos/cosmos-sdk/blob/main/docs/architecture/adr-030-authz-module.md
*)
(******************************************************************************)
EXTENDS GrantMessages, Maps, Integers

VARIABLES
    \* @type: GRANT_ID -> GRANT;  
    grantStore, \* Representation of the KV store implemented by the authz 
                \* module in the server, used to store mappings from grant 
                \* triples to authorizations.
    
    \* @type: EVENT;
    event,

    \* @type: RESPONSE_MSG;
    expectedResponse,

    \* @type: Int;
    numRequests

-------------------------------------------------------------------------------
\* @type: (GRANT_ID) => Bool;
HasGrant(g) == g \in DOMAIN grantStore

\* @type: (GRANT_ID) => Bool;
IsExpired(g) == 
    /\ HasGrant(g)
    /\ grantStore[g].expirationTime = "past"

--------------------------------------------------------------------------------
AcceptErrors == {
    "none", 
    "validator-not-allowed",
    "validator-denied",
    "insufficient-amount"
}

NoUpdate == [type |-> "no-update"]

(* AcceptResponse instruments the controller of an authz message if the request
is accepted and if it should be updated or deleted. *)
\* https://github.com/cosmos/cosmos-sdk/blob/f2cea6a137ce19ad8987fa8a0cb99f4b37c4484d/x/authz/authorizations.go#L20
\* @typeAlias: ACCEPT_RESPONSE = [accept: Bool, delete: Bool, updated: AUTH, error: Str];
\* @type: Set(ACCEPT_RESPONSE);
AcceptResponse == [
    \* If Accept=true, the controller can accept and authorization and handle the update.
    accept: BOOLEAN,
    
    \* If Delete=true, the controller must delete the authorization object and release storage resources.
    delete: BOOLEAN,
    
    \* Controller, who is calling Authorization.Accept must check if `Updated !=
    \* nil`. If yes, it must use the updated version and handle the update on the
    \* storage level.
    updated: Authorization \cup {NoUpdate},

    \* This field does not appear in the code; it's here just to simplify the spec.
    error: AcceptErrors
]

--------------------------------------------------------------------------------
(******************************************************************************)
(* Operators that model processing of request messages.                       *)
(******************************************************************************)

\* The interface that includes the three operations below:
\* https://github.com/cosmos/cosmos-sdk/blob/3a1027c74b15ad78270dbe68b777280bde393576/x/authz/tx.pb.go#L331

\* https://github.com/cosmos/cosmos-sdk/blob/afab2f348ab36fe323b791d3fc826292474b678b/x/authz/keeper/msg_server.go#L14
\* @type: (MSG_GRANT) => RESPONSE_GRANT;
CallGrant(msgGrant) == 
    IF msgGrant.granter = msgGrant.grantee THEN 
        [type |-> "response-grant", ok |-> FALSE, error |-> "granter-equal-grantee"]
    ELSE IF msgGrant.grant.expirationTime = "past" THEN 
        [type |-> "response-grant", ok |-> FALSE, error |-> "authorization-expired"]
    ELSE 
        [type |-> "response-grant", ok |-> TRUE, error |-> "none"]

--------------------------------------------------------------------------------
\* https://github.com/cosmos/cosmos-sdk/blob/afab2f348ab36fe323b791d3fc826292474b678b/x/authz/keeper/msg_server.go#L52
CallRevoke(msgRevoke) == 
    CHOOSE response \in MsgRevokeResponses: response.ok

--------------------------------------------------------------------------------
\* https://github.com/cosmos/cosmos-sdk/blob/afab2f348ab36fe323b791d3fc826292474b678b/x/authz/keeper/keeper.go#L90
\* @type: (ACCOUNT, SDK_MSG) => ACCEPT_RESPONSE;
DispatchActionsOneMsg(grantee, msg) == 
    LET 
        granter == msg.signer \* Actually, granter is the first of the list of signers.
        grantId == [granter |-> granter, grantee |-> grantee, msgTypeUrl |-> msg.content.typeUrl]
    IN
    IF granter = grantee THEN 
        [accept |-> TRUE, delete |-> FALSE, updated |-> NoUpdate, error |-> "none"]
    ELSE IF ~ HasGrant(grantId) THEN 
        [accept |-> FALSE, delete |-> FALSE, updated |-> NoUpdate, error |-> "grant-not-found"]
    ELSE IF grantStore[grantId].expirationTime = "past" THEN 
        [accept |-> FALSE, delete |-> FALSE, updated |-> NoUpdate, error |-> "authorization-expired"]
    ELSE 
        Accept(grantStore[grantId].authorization, msg.content)

\* https://github.com/cosmos/cosmos-sdk/blob/afab2f348ab36fe323b791d3fc826292474b678b/x/authz/keeper/msg_server.go#L72
\* @type: (MSG_EXEC) => <<RESPONSE_EXEC, ACCEPT_RESPONSE>>;
CallExec(msgExec) == 
    LET 
        acceptResponse == DispatchActionsOneMsg(msgExec.grantee, msgExec.msg)
        execResponse == [
            type |-> "response-execute", 
            ok |-> acceptResponse.accept, 
            error |-> acceptResponse.error
        ] 
    IN 
    <<execResponse, acceptResponse>>

--------------------------------------------------------------------------------
\* Only for the initial state.
NoEvent == [type |-> "no-event"]

\* EmptyMap is not accepted by Apalache's typechecker.
\* @type: GRANT_ID -> GRANT;
EmptyStore == [x \in {} |-> NoGrant]

TypeOK == 
    /\ IsMap(grantStore, ValidGrantIds, Grants \cup {NoGrant})
    /\ event \in RequestMessages \cup ExpireEvents \cup {NoEvent}
    /\ expectedResponse \in Responses \cup {NoResponse}
    /\ numRequests \in Nat

Init ==
    /\ grantStore = EmptyStore
    /\ event = NoEvent
    /\ expectedResponse = NoResponse
    /\ numRequests = 0

--------------------------------------------------------------------------------
\* @type: (MSG_GRANT) => GRANT_ID;
grantIdOfMsgGrant(msg) == [
    grantee |-> msg.grantee,
    granter |-> msg.granter,
    msgTypeUrl |-> MsgTypeURL(msg.grant.authorization)
]

(******************************************************************************)
(* Request to give a grant from a granter to a grantee.                       *)
(*                                                                            *)
(* From the code: "An authorization grant is created using the MsgGrant message.
If there is already a grant for the (granter, grantee, Authorization) triple,
then the new grant overwrites the previous one. To update or extend an existing
grant, a new grant with the same (granter, grantee, Authorization) triple should
be created."

The message handling should fail if:
- both granter and grantee have the same address.
- provided Expiration time is less than current unix timestamp.
- provided Grant.Authorization is not implemented.
- Authorization.MsgTypeURL() is not defined in the router (there is no defined
handler in the app router to handle that Msg types). *)
(******************************************************************************)
\* @type: (ACCOUNT, ACCOUNT, GRANT) => Bool;
RequestGrant(granter, grantee, grant) ==
    LET 
        msg == [type |-> "request-grant", granter |-> granter, grantee |-> grantee, grant |-> grant]
        g == grantIdOfMsgGrant(msg)
    IN
    /\ IsValid(g)
    /\ ~ HasGrant(g) \/ IsExpired(g)
    /\ LET response == CallGrant(msg) IN
        /\ IF response.ok THEN
                grantStore' = MapPut(grantStore, g, grant)
            ELSE
                UNCHANGED grantStore
        /\ expectedResponse' = response
    /\ event' = msg
    /\ numRequests' = numRequests + 1

(******************************************************************************)
(* Request to revoke a grant.                                                 *)
(******************************************************************************)
\* @type: (ACCOUNT, ACCOUNT, MSG_TYPE_URL) => Bool;
RequestRevoke(granter, grantee, msgTypeUrl) == 
    LET g == [granter |-> granter, grantee |-> grantee, msgTypeUrl |-> msgTypeUrl] IN
    /\ IsValid(g)
    /\ HasGrant(g)
    /\ LET msg == [type |-> "request-revoke", granter |-> granter, grantee |-> grantee, msgTypeUrl |-> msgTypeUrl] IN
        /\ event' = msg
        /\ LET response == CallRevoke(msg) IN
            /\ IF response.ok THEN
                    grantStore' = MapRemove(grantStore, g)
                ELSE 
                    UNCHANGED grantStore
            /\ expectedResponse' = response
    /\ numRequests' = numRequests + 1

\* https://github.com/cosmos/cosmos-sdk/blob/4eec00f9899fef9a2ea3f937ac960ee97b2d7b18/x/authz/keeper/keeper.go#L99
\* @type: (ACCOUNT, SDK_MSG, ACCEPT_RESPONSE) => Bool;
PostProcessExec(grantee, msg, acceptResponse) == 
    LET 
        \* @type: GRANT_ID;
        g == [granter |-> msg.signer, grantee |-> grantee, msgTypeUrl |-> msg.content.typeUrl] 
    IN
    IF acceptResponse.delete THEN
        grantStore' = MapRemove(grantStore, g)
    ELSE IF acceptResponse.updated # NoUpdate THEN
        grantStore' = [grantStore EXCEPT ![g].authorization = acceptResponse.updated]
    ELSE
        UNCHANGED <<grantStore>>

(******************************************************************************)
(* Request to execute a message on behalf of a grantee.                       *)
(******************************************************************************)
\* @type: (ACCOUNT, SDK_MSG) => Bool;
RequestExec(grantee, msg) ==
    LET 
        request == [type |-> "request-execute", grantee |-> grantee, msg |-> msg]
        responses == CallExec(request) 
    IN
    /\ event' = request
    /\ expectedResponse' = responses[1]
    /\ PostProcessExec(grantee, msg, responses[2])
    /\ numRequests' = numRequests + 1

(******************************************************************************)
(* Timeout being reached is abstracted away by the action Expire. What is
lost then is the relation between different grants and their expirations (in
real life, there could be a dependency: if A expires, then definitely B and C
have to expire). This action is a pure environment action and thus it does not
remove the grant from a set of active grants (the system will do so once it
notices that the grant expired, upon running the Execute function). *)
(******************************************************************************)
Expire(grantId) ==
    /\ HasGrant(grantId)
    /\ ~ IsExpired(grantId)
    /\ grantStore[grantId].expirationTime # "none"
    /\ grantStore' = [grantStore EXCEPT ![grantId].expirationTime = "past"]
    /\ event' = [type |-> "expire", grantId |-> grantId]
    /\ expectedResponse' = NoResponse
    /\ UNCHANGED numRequests

--------------------------------------------------------------------------------
\* We keep action RequestExec separated from the other actions to be able to 
\* check properties on grants without executing them. 
NextA == 
    \/ \E granter, grantee \in Accounts, grant \in Grants: 
        RequestGrant(granter, grantee, grant)
    \/ \E grantId \in GrantIds: 
        RequestRevoke(grantId.granter, grantId.grantee, grantId.msgTypeUrl)
    \/ \E grantId \in ValidGrantIds: 
        Expire(grantId)

Next == 
    \/ NextA
    \* NB: The implementation allows to send more than one message in an Exec
    \* request. We model execution requests of only one message per call.
    \/ \E grantee \in Accounts, msg \in SdkMsgs: 
        RequestExec(grantee, msg)

================================================================================
Created by Hernán Vanzetto on 10 August 2022
