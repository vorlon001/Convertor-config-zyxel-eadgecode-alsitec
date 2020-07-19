try:
    import re, os, json
    from pprint import pprint as dump
    from enum import Enum
    import AST as ast
    import struct as obj
    import parser as node
    import utils as util
    import MODEL as model
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)


def create(name):

    device = model.DEVICE()

    cfg = util.load_config_from_file(name)
    l=[ obj.Struct(**node.parsing(i.strip())) for i in cfg.split('\n')]
    cfg = node.gen_graph(l)

#    node.show_node(cfg,debug=True);

    for i in cfg:
        if i.TYPE==ast.AST.hostname:
            device.hostname = i.RGX[1];
        elif i.TYPE==ast.AST.interface_range_ethernet:
            ports = node.parsing_interface(i.RGX[1]).PORT
            if len(i.Node)>0:
                for j in i.Node:
                    if j.TYPE!=None:
                        if j.TYPE == ast.AST.description:
                            for h in ports:
                                port = device.getport(int( h  ))
                                port.setdescription( j.RGX[1] )
                        elif j.TYPE == ast.AST.switchport_trunk_allowed_vlan_add:
                            for h in ports:
                                port = device.getport( int ( h ) );
                                port.addtagged( j.RGX[1])
                        elif j.TYPE == ast.AST.switchport_access:
                            for h in ports:
                                port = device.getport( int ( h ) );
                                port.adduntagged( j.RGX[2])
                        elif j.TYPE == ast.AST.switchport_trunk_native_vlan:
                            for h in ports:
                                port = device.getport( int ( h ) );
                                port.adduntagged( j.RGX[1])
                        elif j.TYPE == ast.AST.switchport_general_pvid:
                            for h in ports:
#                                print("   port",h," switchport_general_pvid",j.RGX);
                                port = device.getport( int ( h ) );
                                port.adduntagged( j.RGX[1])
                                port.addgeneral_pid(  j.RGX[1] )
                        elif j.TYPE == ast.AST.switchport_general_allowed_vlan_add:
                            for h in ports:
#                                print("   port",h," ast.AST.switchport_general",j.RGX);
                                port = device.getport( int ( h ) );
                                if len(j.RGX)==2:
                                    port.addtagged( j.RGX[1])
                                    port.addgeneral_tag(  j.RGX[1] )
                                elif len(j.RGX)==3:
                                    port.adduntagged( j.RGX[1])
                                    port.addgeneral_untag(  j.RGX[1] )
                                else:
                                    print("   port range", h ," ast.AST.switchport_general_allowed_vlan_add",j.RGX);
        elif i.TYPE==ast.AST.interface_ethernet:
            if len(i.Node)>0:
                for j in i.Node:
                    if j.TYPE != None:
                        port = device.getport( int (i.RGX[2] ) );
                        if j.TYPE == ast.AST.description:
                            port.setdescription( j.RGX[1] )
                        elif j.TYPE == ast.AST.switchport_general_allowed_vlan_add:
                            if len(j.RGX)==2:
                                port = device.getport( int ( i.RGX[2] ) );
                                port.addtagged( j.RGX[1])
                            else:
                                port = device.getport( int ( i.RGX[2] ) );
                                port.addtagged( j.RGX[2])
                        elif j.TYPE == ast.AST.switchport_general_pvid:
                            port.adduntagged( j.RGX[1])
                            port.addgeneral_pid(  j.RGX[1] )
#                            print("   port",i.RGX[2]," switchport_general_pvid",j.RGX);
                        elif j.TYPE == ast.AST.switchport_general:
                            port = device.getport( int ( i.RGX[2] ) );
#                            print("   port",i.RGX[2]," ast.AST.switchport_general",j.RGX);
                            if len(j.RGX)==2:
                                port = device.getport( int ( i.RGX[2] ) );
                                port.addtagged( j.RGX[1])
                                port.addgeneral_tag(  j.RGX[1] )
                            elif len(j.RGX)==3:
                                port.adduntagged( j.RGX[1])
                                port.addgeneral_tag(  j.RGX[1] )
                            else:
                                print("   port",i.RGX[2]," ast.AST.switchport_general",j.RGX);
        elif i.TYPE==ast.AST.vlan_database:
            if len(i.Node)>0:
                for j in i.Node:
                    if j.TYPE==ast.AST.vlan:
                        p = node.parsing_port(j.RGX[1]).VLAN
                        for _,i in p.items():
                            device.addvlan(model.VLAN("",i))
        elif i.TYPE==ast.AST.interface_vlan:
            if len(i.Node)>0:
                for j in i.Node:
                    if j.TYPE!=None:
                        if j.TYPE==ast.AST.name:
                            device.setvlanname( i.RGX[1], j.RGX[1] );
                        elif j.TYPE==ast.AST.ip_address:
                            device.mng_ip, device.mng_mask = j.RGX[1],j.RGX[2]
                            device.mng_int_vlan = i.RGX[1]
                        elif j.TYPE==ast.AST.ip_igmp_snooping:
                            if i.RGX[1] not  in device.igmp_snooping:
                                device.igmp_snooping.append( i.RGX[1] );
                        else:
                            print("   ",j.__dict__);
        elif i.TYPE==ast.AST.ip_dhcp_snooping:
            if i.RGX[2] not  in device.dhcp_snooping:
                device.dhcp_snooping.append( i.RGX[2] );
        elif i.TYPE==ast.AST.ip_default_gateway:
            device.default_gateway = i.RGX[1]

    return device

def main():

    name = "0.0.0.0"
    cfg = create(f"{name}.txt")
    with open(f"{name}.json", "w") as fp:
        json.dump(json.loads(cfg.toJSON()), fp , sort_keys=True, indent=4)


if __name__== "__main__":
    main()

