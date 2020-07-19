try:
    from pprint import pprint as dump
    from enum import Enum
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)

class AST(Enum):
    hostname		= "hostname"
    mvr			= "mvr"
    interface		= "interface"
    inactive		= "inactive"
    vlan 		= "vlan"
    fixed 		= "fixed"
    forbidden 		= "forbidden"
    untagged		= "untagged"
    name		= "name"
    ip 			= "ip"
    default_management 	= "default-management"
    default_gateway 	= "default-gateway"
    exit		= "exit"
    source_port         = "source-port"
    receiver_port       = "receiver-port"
    tagged              = "tagged"

AST_MAP = {
    "hostname":			AST.hostname,
    "interface":		AST.interface,
    "inactive":			AST.inactive,
    "vlan": 			AST.vlan,
    "mvr":			AST.mvr,
    "fixed": 			AST.fixed,
    "forbidden": 		AST.forbidden,
    "untagged":			AST.untagged,
    "name": 			AST.name,
    "ip":			AST.ip,
    "default-management":	AST.default_management,
    "default-gateway":		AST.default_gateway,
    "source-port":		AST.source_port,
    "receiver-port":		AST.receiver_port,
    "tagged":			AST.tagged,
    "exit":			AST.exit
}

AST_NODE		= [ AST.interface, AST.vlan, AST.mvr ]
AST_PORT_MAP		= [ AST.fixed, AST.forbidden, AST.untagged, AST.source_port, AST.receiver_port, AST.tagged]
AST_NOTYPE		= ""
default_management 	= "default-management"
default_gateway		= "default-gateway"
