import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddAdmin } from "./types/adminmodule/tx";
import { MsgSubmitProposal } from "./types/adminmodule/tx";
import { MsgDeleteAdmin } from "./types/adminmodule/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/cosmos.adminmodule.adminmodule.MsgAddAdmin", MsgAddAdmin],
    ["/cosmos.adminmodule.adminmodule.MsgSubmitProposal", MsgSubmitProposal],
    ["/cosmos.adminmodule.adminmodule.MsgDeleteAdmin", MsgDeleteAdmin],
    
];

export { msgTypes }