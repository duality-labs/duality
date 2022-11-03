# Test scenarios for Authz module

*Note 1*: all of these scenarios may contain a condition requiring that some number of active grants (e.g., 3) existed. 
This is to make test scenarios a bit more involved in an arbitrary fashion. 
For brevity, I am omitting that condition in the list of scenarios.

*Note 2*: most of these scenarios can be run for all authorization logics together or, when we want to make sure that a particular authorization logic is tested, separated (i.e., allowing only a single authorization logic).

## A) Granting
1. A grant is successfully given.
1. Giving a grant fails (for whatever reason).
1. Giving a grant fails because:
    - provided message type is not defined 
    - granter equals grantee
    - provided authorization is not implemented
    - [banking specific] allowance is not specified
1. There was a failed grant, followed by a successful one.
1. A grant expired.
1. A grant expired, and then the same grant was given.
1. A grant is given and is subsequently overwritten by a grant with a different expiration time (where times could include `nil`, meaning no expiration).
1. A grant is given and is subsequently overwritten by a grant with a different authorization parameters.
1. Grant was successfully revoked.
1. Revoking a grant fails (for whatever reason).
1. Revoking a grant fails because:
    - granter equals grantee
    - provided message type is empty
    - there is no such grant in the list of active grants
1. A grant with no expiration time is created.
1. A grant is given with expiration time in the past (*note*: depending on modeling the expiration, this may or may not be possible to test).


## B) Execution
1. A message is successfully executed based on a grant.
1. A message is successfully executed based on a staking grant with infinite limit.
1. A message is successfuly executed and it fully spent the allowance of the corresponding grant (applicable both for `bank` and `stake`).
1. A message is successfully executed and it fully spent the allowance of the corresponding grant (applicable both for `bank` and `stake`). This was followed by another (unsuccessful) attempt to use the same grant.

1. A message was not executed successfully: there was no relevant grant.
1. A message was not executed successfully: there was a grant, but the authorization did not go through (for any reason).
1. A message was not executed successfully: ther was a grant, but the authorization failed because of insufficient allowance (applicable both for `bank` and `stake`).
1. A message was not executed successfully: the relevant grant has expired.
1. [staking specific] A message was executed successfully, and the grant had a non-empty allow list.
1. [staking specific] A message was executed successfully, and the grant had a non-empty deny list.
1. [staking specific] A message was not executed successfully: there was a grant, but the authorization did not go through because the validator was in the deny list.
1. [staking specific] A message was not executed successfully: there was a grant, but the authorization did not go through because there was an allow list, and validator was not in it.
1. [staking specific] A message was not executed successfully: the staking action did not match one from the grant (e.g., a `delegate` was required and grant allowed for `undelegate`).
1. [staking specific] A grant with no allowance (--> infinite allowance) is given. Then it is overwritten with the same grant, but with finite allowance `x`. Finally, an execution is attempted asking for `y`, where `y > x` (it should fail).


