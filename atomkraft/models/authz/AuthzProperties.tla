--------------------------- MODULE AuthzProperties -----------------------------
EXTENDS Authz, FiniteSets, Integers

(******************************************************************************)

\** Apalache typecheker fails:
\* \* Filter the function `f` by the values in its range. For entries k |-> v, keep
\* \* only the entries where `g(v)` is true.
\* \* @ type: ((a -> b), (b -> Bool)) => (a -> b);
\* FilterRange(f, g(_)) ==
\*     LET dom == {x \in DOMAIN f: g(f[x])} IN
\*     [x \in dom |-> f[x]]

\* \* @type: Int;
\* NumActiveGrants == 
\*     LET activeStore == FilterRange(grantStore, 
\*         LAMBDA grant: grant # NoGrant /\ grant.expirationTime # "past") 
\*     IN
\*     Cardinality(activeStore)

--------------------------------------------------------------------------------

GrantSuccess ==
    /\ event.type = "request-grant"
    /\ event.granter # event.grantee
    /\ expectedResponse.ok = TRUE

NotGrantSuccess == ~ GrantSuccess

GrantFailedSameAccounts ==
    /\ event.type = "request-grant"
    /\ expectedResponse.ok = FALSE
    /\ expectedResponse.error = "granter-equal-grantee"

NotGrantFailedSameAccounts == ~ GrantFailedSameAccounts

GrantFailedAuthExpired ==
    /\ event.type = "request-grant"
    /\ expectedResponse.ok = FALSE
    /\ expectedResponse.error = "authorization-expired"

NotGrantFailedAuthExpired == ~ GrantFailedAuthExpired

RevokeSuccess ==
    /\ event.type = "request-revoke"
    /\ expectedResponse.ok = TRUE

NotRevokeSuccess == ~ RevokeSuccess

--------------------------------------------------------------------------------
\* @typeAlias: TRACE = [grantStore: GRANT_ID -> GRANT, event: EVENT, expectedResponse: RESPONSE_MSG, numRequests: Int];
\* @type: Seq(TRACE) => Bool;
ExpireRevokeFailure(trace) ==
    \E i, j \in DOMAIN trace: j = i + 1 /\
        LET state1 == trace[i] IN 
        LET state2 == trace[j] IN
        /\ state1.event.type = "expire"
        /\ state2.event.type = "request-revoke" 
        /\ state2.expectedResponse.ok = FALSE

NotExpireRevokeFailure(trace) == ~ ExpireRevokeFailure(trace)

\* @type: Seq(TRACE) => Bool;
ExpireRevokeFailureSameGrant(trace) ==
    \E i, j \in DOMAIN trace: j = i + 1 /\
        LET state1 == trace[i] IN 
        LET state2 == trace[j] IN
        /\ state1.event.type = "expire"
        /\ state2.event.type = "request-revoke" 
        /\ state2.expectedResponse.ok = FALSE
        /\ state1.event.grantId = grantIdOfMsgRevoke(state2.event)

NotExpireRevokeFailureSameGrant(trace) == ~ ExpireRevokeFailureSameGrant(trace)

\* @type: Seq(TRACE) => Bool;
GrantExpireExec(trace) ==
    LET 
        state1 == trace[1] 
        g1 == grantIdOfMsgGrant(state1.event)
    IN
    /\ state1.event.type = "request-grant"
    /\ state1.expectedResponse.ok = TRUE
    /\ \E j, k \in DOMAIN trace: 
        /\ 1 < j 
        /\ j < k 
        /\  LET 
                state2 == trace[j]
                state3 == trace[k] 
            IN
            /\ state2.event.type = "request-expire"
            /\ g1 = grantIdOfMsgRevoke(state2.event)
            /\ state3.event.type = "request-execute" 
            /\ LET 
                \* @type: SDK_MSG;
                msg == CHOOSE m \in state3.event.msgs: TRUE IN
                g1.msgTypeUrl = msg.content.typeUrl

NotGrantExpireExec(trace) == ~ GrantExpireExec(trace)

\* \* 3. MIREL Test Case: testing grant failures and then successful creation of grant
\* \* @type: Seq(STATE) => Bool;
\* GrantFailedFollowedBySuccess(trace) ==    
\*     \E i \in DOMAIN trace:
\*         LET state1 == trace[i] IN 
\*         LET state2 == trace[i+1] IN
\*         /\ state1.outcome_status \in GRANT_FAILURE_REASONS
\*         \* /\ state1.outcome_status = GRANT_FAILED
\*         /\ state2.outcome_status = GRANT_SUCCESS            
\*         /\ Len(trace) >= i+1

\* \* 4. MIREL Test Case: testing grant failure, succesful creating grant and then Revoke
\* (* <--- *)
\* \* @type: Seq(STATE) => Bool;
\* GrantFailuresPreconditionsNotMetFollowedBySuccess(trace) ==
\*     \E i \in DOMAIN trace:
\*         LET state1 ==trace[i] IN 
\*         LET state2 == trace [j] IN 
\*         LET state3 == trace [k] IN
\*         LET grant1 == state1.action_taken.grant IN
\*         LET grant2 == state2.action_taken.grant IN
\*         LET grant3 == state3.action_taken.grant IN    
\*         /\ i < j
\*         /\ j < k
\*         /\ state1.outcome_status \in GRANT_FAILURE_REASONS
\*         /\ state2.outcome_status = GRANT_SUCCESS
\*         /\ state3.outcome_status = REVOKE_SUCCESS
\*         /\ grant1 = grant2
\*         /\ grant2 = grant3
\*         /\ Len(trace) >= k

================================================================================
Created by Hernán Vanzetto on 10 August 2022
