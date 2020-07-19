try:
    from pprint import pprint as dump
    from enum import Enum
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)

class AST(Enum):
    hostname					= "hostname"
    interface_range_ethernet			= "interface range ethernet"
    interface_ethernet				= "interface ethernet"
    vlan 					= "vlan"
    vlan_database				= "vlan database"
    line					= "line"
    description					= "description"
    switchport_mode_trunk                 	= "switchport mode trunk"
    switchport_mode_general               	= "switchport mode general"
    switchport_access                     	= "switchport access"
    switchport_trunk_native_vlan		= "switchport trunk native vlan"
    switchport_trunk_allowed_vlan_add     	= "switchport trunk allowed vlan add"
    switchport_general_allowed_vlan_add   	= "switchport general allowed vlan add"
    switchport_general_pvid			= "switchport general pvid"
    switchport_general                    	= "switchport general"
    interface_vlan                        	= "interface vlan"
    name                                  	= "name"
    ip_address                            	= "ip address"
    ip_default_gateway                    	= "ip default-gateway"
    ip_igmp_snooping                      	= "ip igmp snooping"
    ip_dhcp_snooping_pppoe_snooping       	= "ip dhcp snooping pppoe snooping"
    pppoe_snooping_vlan                   	= "pppoe snooping vlan"
    ip_dhcp_snooping                      	= "ip dhcp snooping"
    exit					= "exit"

AST_MAP = {
    "hostname":					AST.hostname,
    "interface range ethernet":			AST.interface_range_ethernet,
    "interface ethernet":			AST.interface_ethernet,
    "vlan": 					AST.vlan,
    "vlan database":				AST.vlan_database,
    "line":					AST.line,
    "description":				AST.description,
    "switchport mode trunk":                	AST.switchport_mode_trunk,
    "switchport mode general":              	AST.switchport_mode_general,
    "switchport access":                    	AST.switchport_access,
    "switchport trunk native vlan":		AST.switchport_trunk_native_vlan,
    "switchport trunk allowed vlan add":    	AST.switchport_trunk_allowed_vlan_add,
    "switchport general allowed vlan add":  	AST.switchport_general_allowed_vlan_add,
    "switchport general":                   	AST.switchport_general,
    "switchport general pvid":			AST.switchport_general_pvid,
    "interface vlan":                       	AST.interface_vlan,
    "name":                                 	AST.name,
    "ip address":                           	AST.ip_address,
    "ip default-gateway":                   	AST.ip_default_gateway,
    "ip igmp snooping":                     	AST.ip_igmp_snooping,
    "ip dhcp snooping pppoe snooping":      	AST.ip_dhcp_snooping_pppoe_snooping,
    "pppoe snooping vlan":                  	AST.pppoe_snooping_vlan,
    "ip dhcp snooping":                     	AST.ip_dhcp_snooping,
    "exit":					AST.exit
}

AST_NODE		= [ AST.interface_range_ethernet, AST.vlan_database, AST.interface_ethernet, AST.line, AST.interface_vlan  ]
AST_PORT_MAP		= [  ]
AST_NOTYPE		= ""
#default_management 	= "default-management"
#default_gateway		= "default-gateway"
