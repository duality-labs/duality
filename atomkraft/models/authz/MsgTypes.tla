------------------------------ MODULE MsgTypes --------------------------------
(******************************************************************************)
(* SDK message type URLs                                                      *)
(******************************************************************************)

\* @typeAlias: MSG_TYPE_URL = Str;
\* @type: MSG_TYPE_URL;
SEND_TYPE_URL == "send"

\* @type: MSG_TYPE_URL;
DELEGATE_TYPE_URL == "delegate"

\* @type: MSG_TYPE_URL;
UNDELEGATE_TYPE_URL == "undelegate"

\* @type: MSG_TYPE_URL;
BEGIN_REDELEGATE_TYPE_URL == "redelegate"

MsgTypes == {
    SEND_TYPE_URL, 
    DELEGATE_TYPE_URL, 
    UNDELEGATE_TYPE_URL, 
    BEGIN_REDELEGATE_TYPE_URL
}

================================================================================
Created by Hernán Vanzetto on 6 September 2022
