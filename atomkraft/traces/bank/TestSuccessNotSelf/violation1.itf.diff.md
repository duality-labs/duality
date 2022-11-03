# ITF-Diff Markdown

## Step 1 to Step 2

<details open>

<summary>Variables</summary>

<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.coins`|`None`|`<<[ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819966, denom \|-> "gluon" ], [ amount \|-> 100, denom \|-> "muon" ]>>`|
|`action.receiver`|`None`|`"Eve"`|
|`action.sender`|`None`|`"Bob"`|
|`action.tag`|`init`|`send`|
|`action.balances`|`SetAsFun({<<"Alice", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Bob", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Carol", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Dave", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Eve", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>})`|`None`|

</details>
<details open>

<summary><code>balances</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`balances("Bob")("muon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`57896044618658097711785492504343953926634992332820282019728792003956564819867`|
|`balances("Eve")("muon")`|`0`|`100`|
|`balances("Bob")("gluon")`|`57896044618658097711785492504343953926634992332820282019728792003956564819967`|`1`|
|`balances("Eve")("gluon")`|`0`|`57896044618658097711785492504343953926634992332820282019728792003956564819966`|

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

