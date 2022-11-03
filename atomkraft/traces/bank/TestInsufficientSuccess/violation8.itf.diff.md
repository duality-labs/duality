# ITF-Diff Markdown

## Step 1 to Step 2

<details open>

<summary>Variables</summary>

<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.coins`|`None`|`<<[ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819965, denom \|-> "muon" ], [ amount \|-> 2, denom \|-> "atom" ], [ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819965, denom \|-> "gluon" ]>>`|
|`action.receiver`|`None`|`"Eve"`|
|`action.sender`|`None`|`"Alice"`|
|`action.tag`|`init`|`send`|
|`action.balances`|`SetAsFun({<<"Alice", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Bob", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Carol", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Dave", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Eve", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>})`|`None`|

</details>
<details open>

<summary><code>balances</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`balances("Alice")("atom")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`57896044618658097711785492504343953926634992332820282019728792003956564819965`|
|`balances("Eve")("atom")`|`0`|`2`|
|`balances("Alice")("gluon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`2`|
|`balances("Alice")("muon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`2`|
|`balances("Eve")("gluon")`|`0`|`57896044618658097711785492504343953926634992332820282019728792003956564819965`|
|`balances("Eve")("muon")`|`0`|`57896044618658097711785492504343953926634992332820282019728792003956564819965`|

</details>
<details open>

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|``|`SUCCESS`|

</details>
<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`0`|`1`|

</details>

</details>

## Step 2 to Step 3

<details open>

<summary>Variables</summary>

<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.receiver`|`Eve`|`Bob`|
|`action.sender`|`Alice`|`Carol`|
|`action.coins[0]`|`[ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819965, denom \|-> "muon" ]`|`[ amount \|-> 1, denom \|-> "gluon" ]`|
|`action.coins[1]`|`[ amount \|-> 2, denom \|-> "atom" ]`|`[ amount \|-> 0, denom \|-> "muon" ]`|
|`action.coins[2]`|`[ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819965, denom \|-> "gluon" ]`|`None`|

</details>
<details open>

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|`SUCCESS`|`AMOUNT_NOT_POSITIVE`|

</details>
<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`1`|`2`|

</details>

</details>

## Step 3 to Step 4

<details open>

<summary>Variables</summary>

<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.receiver`|`Bob`|`Eve`|
|`action.sender`|`Carol`|`Bob`|
|`action.coins[0]`|`[ amount \|-> 1, denom \|-> "gluon" ]`|`[ amount \|-> 115792089237316195423570985008687907853269984665640564039457584007913129639937, denom \|-> "muon" ]`|
|`action.coins[1]`|`[ amount \|-> 0, denom \|-> "muon" ]`|`[ amount \|-> 3, denom \|-> "atom" ]`|

</details>
<details open>

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|`AMOUNT_NOT_POSITIVE`|`INSUFFICIENT_FUNDS`|

</details>
<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`2`|`3`|

</details>

</details>

## Step 4 to Step 5

<details open>

<summary>Variables</summary>

<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.receiver`|`Eve`|`Bob`|
|`action.coins[0]`|`[ amount \|-> 115792089237316195423570985008687907853269984665640564039457584007913129639937, denom \|-> "muon" ]`|`[ amount \|-> 1, denom \|-> "gluon" ]`|
|`action.coins[1]`|`[ amount \|-> 3, denom \|-> "atom" ]`|`[ amount \|-> 1, denom \|-> "atom" ]`|

</details>
<details open>

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|`INSUFFICIENT_FUNDS`|`SUCCESS`|

</details>
<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`3`|`4`|

</details>

</details>

