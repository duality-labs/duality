# ITF-Diff Markdown

## Step 1 to Step 2

<details open>

<summary>Variables</summary>

<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.coins`|`None`|`<<[ amount \|-> 0, denom \|-> "gluon" ], [ amount \|-> 115792089237316195423570985008687907853269984665640564039457584007913129639935, denom \|-> "gluon" ]>>`|
|`action.receiver`|`None`|`"Alice"`|
|`action.sender`|`None`|`"Eve"`|
|`action.tag`|`init`|`send`|
|`action.balances`|`SetAsFun({<<"Alice", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Bob", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Carol", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Dave", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Eve", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>})`|`None`|

</details>
<details open>

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|``|`DUPLICATE_DENOM`|

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
|`action.receiver`|`Alice`|`Carol`|
|`action.sender`|`Eve`|`Bob`|
|`action.coins[0]`|`[ amount \|-> 0, denom \|-> "gluon" ]`|`[ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819967, denom \|-> "gluon" ]`|
|`action.coins[1]`|`[ amount \|-> 115792089237316195423570985008687907853269984665640564039457584007913129639935, denom \|-> "gluon" ]`|`[ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819967, denom \|-> "muon" ]`|
|`action.coins[2]`|`None`|`[ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819967, denom \|-> "atom" ]`|

</details>
<details open>

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|`DUPLICATE_DENOM`|`SUCCESS`|

</details>
<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`1`|`2`|

</details>
<details open>

<summary><code>balances</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`balances("Bob")("atom")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`0`|
|`balances("Bob")("gluon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`0`|
|`balances("Bob")("muon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`0`|
|`balances("Carol")("atom")`|`0`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|
|`balances("Carol")("gluon")`|`0`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|
|`balances("Carol")("muon")`|`0`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|

</details>

</details>

