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
    cfg_new = node.gen_graph(l)
    node.show_node(cfg_new)

    for i in cfg_new:
        if i.TYPE==ast.AST.hostname:
            device.hostname = i.RGX[1];
        elif i.TYPE==ast.AST.interface:
            for j in i.Node:
                if j.TYPE == ast.AST.name:
                     port = device.getport(int(i.RGX[2]))
                     port.setdescription(j.RGX[1])
                if j.TYPE == ast.AST.inactive:
                     port = device.getport(int(i.RGX[2]))
                     port.setdown()
        elif i.TYPE==ast.AST.mvr:
            source_port, receiver_port, name, tagged = None, None, None, None
            vlan = model.VLAN(name,i.RGX[1])
            device.addvlan(vlan)
            for j in i.Node:
                if j.TYPE==ast.AST.source_port:
                    source_port = j.PORT
                elif j.TYPE==ast.AST.receiver_port:
                    receiver_port = j.PORT
                elif j.TYPE==ast.AST.name:
                    name = j.RGX[1]
                elif j.TYPE==ast.AST.tagged:
                    tagged = j.PORT
            device.setmvr(model.MVR(i.RGX[1],  name));
            for v in source_port:
                port = device.getport(v)
                port.mvr.tag = i.RGX[1]
                port.mvr.receiver_port = True
            for v in receiver_port:
                port = device.getport(v)
                port.mvr.tag = i.RGX[1]
                port.mvr.source_port = True
            for v in tagged:
                port = device.getport(v)
                port.addtagged(i.RGX[1])
        elif i.TYPE==ast.AST.vlan:
            fixed, forbidden, untagged, name  = None, None, None, None
            for j in i.Node:
                if j.TYPE==ast.AST.fixed:
                    fixed=j.PORT
                elif j.TYPE==ast.AST.forbidden:
                    forbidden=j.PORT
                elif j.TYPE==ast.AST.untagged:
                    untagged=j.PORT
                elif j.TYPE==ast.AST.name:
                    name = j.RGX[1]
                elif j.TYPE==ast.AST.ip:
                    if ast.default_management in j.RGX:
                        device.mng_ip, device.mng_mask = j.RGX[3], j.RGX[4]
                        device.mng_int_vlan = i.RGX[1]
                    elif ast.default_gateway in j.RGX:
                        device.default_gateway = j.RGX[3]
            device.addvlan(model.VLAN(name,i.RGX[1]))
            for k,_ in fixed.items():
                port = device.getport(k);
                if k in untagged:
                    port.adduntagged(i.RGX[1])
                else:
                    port.addtagged(i.RGX[1])
    print("*"*45);
    return device

def main():

    name = "0.0.0.0"
    cfg = create(f"{name}")
    with open(f"{name}.json", "w") as fp:
        json.dump(json.loads(cfg.toJSON()), fp , sort_keys=True, indent=4)

if __name__== "__main__":
    main()
