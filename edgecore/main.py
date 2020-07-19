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
    cfg = util.load_config_from_file(name)
    l=[ obj.Struct(**node.parsing(i.strip())) for i in cfg.split('\n')]
    cfg_new = node.gen_graph(l)
    node.show_node(cfg_new,debug=True)
    device = model.DEVICE()

    for i in cfg_new:
        if i.TYPE==ast.AST.hostname:
            device.hostname = i.RGX[1];
        elif i.TYPE==ast.AST.ip_default_gateway:
            device.default_gateway = i.RGX[1]
        elif i.TYPE==ast.AST.svi:
            mng = i.Node[0]
            device.mng_ip, device.mng_mask = mng.RGX[2], mng.RGX[3]
            device.mng_int_vlan = i.RGX[1]
        elif i.TYPE==ast.AST.interface:
            port = device.getport(int(i.RGX[1]))
            switchport_allowed_tagged, switchport_allowed_untagged, switchport_native, description = None, None, None, None
            for j in i.Node:
                if j.TYPE ==ast.AST.description:
                    port.setdescription(j.RGX[1])
                elif j.TYPE==ast.AST.switchport_allowed_tagged:
                    port.addtagged( j.VLAN )
                elif j.TYPE==ast.AST.switchport_allowed_untagged:
                    port.adduntagged( j.VLAN )
                elif j.TYPE==ast.AST.description:
                    port.setdescription( j.RGX[1] )
                elif j.TYPE==ast.AST.shutdown:
                    port.setdown();
        elif i.TYPE==ast.AST.vlan:
            if i.RGX[1] not in ['1','4093']:
                device.addvlan(model.VLAN(i.RGX[4],i.RGX[1]))
        elif i.TYPE==ast.AST.ip_igmp_snooping:
            if i.RGX[4] not in device.igmp_snooping:
                device.igmp_snooping.append(i.RGX[4]);
        elif i.TYPE==ast.AST.ip_dhcp_snooping:
            if i.RGX[4] not in device.dhcp_snooping:
                device.dhcp_snooping.append(i.RGX[4]);
   print("*"*45);
    return device


def main():
    name = "0.0.0.0"
    cfg = create(f"{name}.txt")
    with open(f"{name}.json", "w") as fp:
        json.dump(json.loads(cfg.toJSON()), fp , sort_keys=True, indent=4)

if __name__== "__main__":
    main()

