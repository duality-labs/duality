# ITF-Diff Markdown

## Step 1 to Step 2

<details open>

<summary>Variables</summary>

<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.coins`|`None`|`<<[ amount \|-> 1, denom \|-> "muon" ], [ amount \|-> 20, denom \|-> "gluon" ], [ amount \|-> 300, denom \|-> "atom" ]>>`|
|`action.receiver`|`None`|`"Alice"`|
|`action.sender`|`None`|`"Bob"`|
|`action.tag`|`init`|`send`|
|`action.balances`|`SetAsFun({<<"Alice", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Bob", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Carol", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Dave", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Eve", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>})`|`None`|

</details>
<details open>

<summary><code>balances</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`balances("Alice")("atom")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`57896044618658097711785492504343953926634992332820282019728792003956564820267`|
|`balances("Alice")("gluon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`57896044618658097711785492504343953926634992332820282019728792003956564819987`|
|`balances("Alice")("muon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`57896044618658097711785492504343953926634992332820282019728792003956564819968`|
|`balances("Bob")("atom")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`57896044618658097711785492504343953926634992332820282019728792003956564819667`|
|`balances("Bob")("gluon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`57896044618658097711785492504343953926634992332820282019728792003956564819947`|
|`balances("Bob")("muon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`57896044618658097711785492504343953926634992332820282019728792003956564819966`|

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
|`action.receiver`|`Alice`|`Bob`|
|`action.coins[0]`|`[ amount \|-> 1, denom \|-> "muon" ]`|`[ amount \|-> 115792089237316195423570985008687907853269984665640564039457584007913129639935, denom \|-> "gluon" ]`|
|`action.coins[1]`|`[ amount \|-> 20, denom \|-> "gluon" ]`|`None`|
|`action.coins[2]`|`[ amount \|-> 300, denom \|-> "atom" ]`|`None`|

</details>
<details open>

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|`SUCCESS`|`INSUFFICIENT_FUNDS`|

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

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|`INSUFFICIENT_FUNDS`|`SUCCESS`|

</details>
<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`2`|`3`|

</details>
<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.coins[0].amount`|`115792089237316195423570985008687907853269984665640564039457584007913129639935`|`1`|

</details>

</details>

