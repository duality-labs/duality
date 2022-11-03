---------------------------- MODULE BankSendTests ------------------------------
EXTENDS BankSend
(******************************************************************************)
\* Tests for Cosmos SDK Bank module coins send functionality between accounts
--------------------------------------------------------------------------------
\* TLA+ test operators below describe the shape of a test trace you want to see
(******************************************************************************)

\* To generate a trace from e.g. operator ThreeSteps run
\*   atomkraft model sample --model-path models/bank/BankSendTests.tla --examples ThreeSteps --traces-dir traces/bank

\* To generate multiple traces from a single assertion you can add e.g. the following parameters:
\*   --max_error=10 --view=OutcomeView
\* This should generate 10 traces, and the traces should differ according to the OutcomeView operator below,
\* which says that expected outcomes should be different. 
\* You can use other views for focusing on other interesting behaviors.

OutcomeView == outcome
ActionView == action
StepView == step

(******************************************************************************)
\* The simplest test assertion: a trace with 3 steps
TestThreeSteps == step = 3

\* The following command generates 10 arbitrary test traces in ./traces/bank/TestThreeSteps directory
\* each trace cosisting of 3 steps, and the traces should differ with each other in expected outcomes:
\*
\*   atomkraft model sample --model-path models/bank/BankSendTests.tla --examples TestThreeSteps --traces-dir traces/bank --max_error=10 --view=OutcomeView

(******************************************************************************)
\* The assertions below require to have a trace with a particular outcome

TestSuccess      == outcome = "SUCCESS"
TestDuplicate    == outcome = "DUPLICATE_DENOM"
TestNotPositive  == outcome = "AMOUNT_NOT_POSITIVE"
TestInsufficient == outcome = "INSUFFICIENT_FUNDS"

\* If you inspect the traces generated via TestSuccess, you'll notice 
\* that all succeed with the sender and receiver being the same.
\* To generate more interesting traces we add an additional constraint.
TestSuccessNotSelf == 
  /\ outcome = "SUCCESS" 
  /\ action.sender /= action.receiver


(******************************************************************************)
\* How about writing a test for two consecutive execution steps?
\* In TLA+ this can be done easily: you refer to the next step using primes (') notation;
\* e.g. balances' refers to the wallet balances in the next execution step.

\* In the BankSend model there are these state variables:
\* - balances: a map from wallet names to their balances; each balance is a map from denomination to amount
\* - action: action taken; for "send" action it contains sender, receiver, and coins that are being sent
\* - outcome: the expected outcome of the action taken (see above)
\* - step: the number of the current execution step

\* Two consecutive steps are connected as follows:
\* balances are transformed to balances', by executing action', and the expected outcome is outcome'

\* There is a wallet from which all funds are drained in one step
TestDrainAllFunds ==
  \E wallet \in DOMAIN balances:
    /\ \E denom \in DOMAIN balances[wallet]: balances[wallet][denom] > 0
    /\ \A denom \in DOMAIN balances'[wallet]: balances'[wallet][denom] = 0

\* For the same sender, receiver, and coins, 
\* transfer first succeeded, but then failed
TestSuccessFailure ==
  /\ action = action'
  /\ outcome = "SUCCESS"
  /\ outcome' /= "SUCCESS"

\* For the same sender and receiver, 
\* transfer first failed, but then succeeded
TestFailureSuccess ==
  /\ action.sender = action'.sender
  /\ action.receiver = action'.receiver
  /\ outcome /= "SUCCESS"
  /\ outcome' = "SUCCESS"


\* For the same sender, receiver and denom, 
\* transfer first failed with insufficient funds, but then succeeded
TestInsufficientSuccess ==
  /\ action.sender = action'.sender
  /\ action.receiver = action'.receiver
  /\ \E c1 \in DOMAIN action.coins:
     \E c2 \in DOMAIN action'.coins:
      action.coins[c1].denom = action'.coins[c2].denom
  /\ outcome = "INSUFFICIENT_FUNDS"
  /\ outcome' = "SUCCESS"


(******************************************************************************)
\* Having generated the test traces from the above test assertions,
\* you can execute them against your blockchain.
\* Assuming the blockchain parameters are configured, 
\* if you want to execute the tests generated from e.g. TestSuccess assertion,
\* the command below will generate and execute Pytest scripts:
\*
\* atomkraft test trace --path traces/bank/TestSuccess --reactor reactors/bank.py --keypath action.tag

================================================================================
Created by Andrey Kuprianov on 8 September 2022
