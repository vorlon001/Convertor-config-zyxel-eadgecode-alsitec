try:
    from pprint import pprint as dump
    from enum import Enum
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)

class AST(Enum):
    hostname			= "hostname"
    interface			= "interface ethernet"
    ip_igmp_snooping		= "ip igmp snooping"
    ip_dhcp_snooping		= "ip dhcp snooping"
    switchport			= "switchport"
    switchport_allowed_tagged	= "switchport_allowed_tagged"
    switchport_allowed_untagged	= "switchport_allowed_untagged"
    switchport_native		= "switchport_native_"
    svi				= "interface vlan"
    shutdown			= "shutdown"
    vlan 			= "vlan"
    description			= "description"
    ip 				= "ip"
    ip_default_gateway		= "ip default-gateway"
    exit			= "!"

AST_MAP = {
    "hostname":				AST.hostname,
    "interface ethernet":		AST.interface,
    "ip dhcp snooping":			AST.ip_dhcp_snooping,
    "ip igmp snooping":			AST.ip_igmp_snooping,
    "switchport":			AST.switchport,
    "switchport_allowed_tagged":   	AST.switchport_allowed_tagged,
    "switchport_allowed_untagged": 	AST.switchport_allowed_untagged,
    "switchport_native_":		AST.switchport_native,
    "interface vlan":			AST.svi,
    "ip default-gateway":		AST.ip_default_gateway,
    "shutdown":				AST.shutdown,
    "vlan": 				AST.vlan,
    "description": 			AST.description,
    "ip":				AST.ip,
    "!":				AST.exit
}

AST_NODE		= [ AST.interface ]
AST_PORT_MAP		= [ AST.switchport, AST.ip_dhcp_snooping, AST.ip_igmp_snooping ]
AST_NOTYPE		= ""
#default_management 	= "default-management"
#default_gateway		= "default-gateway"
