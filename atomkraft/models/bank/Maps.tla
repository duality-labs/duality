--------------------------------- MODULE Maps ----------------------------------
(******************************************************************************)
(* This module provides maps as a representation of partial functions.

In TLA+ functions are total, that is, they are defined over all elements of
their domain S. A map represents a partial function over S, that is, a function
whose domain is a subset of S, or possibly S itself. *)
(******************************************************************************)

IsMap(f, S, T) ==
    /\ f = [x \in DOMAIN f |-> f[x]]
    /\ DOMAIN f \subseteq S
    /\ \A x \in DOMAIN f: f[x] \in T

\* The empty tuple is the only function in TLA+ with an empty domain.
EmptyMap == <<>>

\* @type: (a -> b, a, b) => a -> b;
MapPut(map, key, value) ==
    [k \in (DOMAIN map) \cup {key} |-> IF k = key THEN value ELSE map[k]]

\* @type: (a -> b, a) => a -> b;
MapRemove(map, key) ==
    [k \in DOMAIN map \ {key} |-> map[k]]

================================================================================
Created by Hernán Vanzetto on 10 August 2022
