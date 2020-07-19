try:
    import re
    from pprint import pprint as dump
    from enum import Enum
    import AST as ast
    import struct as obj
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)

def parsing_port(test_str):
  regex = r"([0-9\-\,]+)?"
  matches = re.findall(regex, test_str)
  matches=matches.pop(0) if len(matches)>1 else ('')
  def get(port):
    _,_,_,_,_= (h:=[ [int(i) for i in x.split('-')] for x in port.split(',')]), \
              (e:=[]), \
              (f:=[]), \
              {e.append([ i for i in range(k[0],k[1]+1 if len(k)==2 else k[0]+1) ]) for k in h}, \
              { f.extend(i) for i in e}
    if 1 in f:
        del f[f.index(1)];
    if 4093 in f:
        del f[f.index(4093)]
    return { i:i for i in f }
  return obj.Struct(**{"CONFIG":test_str,"VLAN":get(matches)})


rgx = [
        obj.Struct(**{ 'RGX': '^(vlan)\s+(\d+)\s+((name)\s+(.+)\s+)?(media)\s+(ethernet)\s(state)\s+(active)$'  }),
        obj.Struct(**{ 'RGX': '^(ip\sdhcp\ssnooping)(\s+)(vlan)(\s+)([0-9\,\-]+)'  }),
        obj.Struct(**{ 'RGX': '^(ip\sigmp\ssnooping)(\s+)(vlan)(\s+)([0-9\,\-]+)\s+(mrouter)'  }),
        obj.Struct(**{ 'RGX': '^(hostname)\s+?(.+)'  }),
        obj.Struct(**{ 'RGX': '^(interface\svlan)\s+([0-9]+)'}),
        obj.Struct(**{ 'RGX': '^(interface\sethernet)\s+1\/(\d+)'}),
        obj.Struct(**{ 'RGX': '^(ip)\s+(address)\s+([0-9\.]+)\s+([0-9\.]+)$' }),
        obj.Struct(**{ 'RGX': '^(description)\s+(.+)'  }),
        obj.Struct(**{ 'RGX': '^(switchport)\s+(allowed)\s+(vlan)\s+(add)\s+(\d+)\s+(untagged)'  }),
        obj.Struct(**{ 'RGX': '^(switchport)\s+(allowed)\s+(vlan)\s+(add)\s+([0-9\,\-]+)\s+(tagged)'  }),
        obj.Struct(**{ 'RGX': '^(switchport)\s+(native)\s(vlan)(\s+)(\d+)(\s+)?'  }),
        obj.Struct(**{ 'RGX': '^(shutdown)'  }),
        obj.Struct(**{ 'RGX': '^(ip\sdefault-gateway)\s+(\d+\.\d+\.\d+\.\d+)'  }),
        obj.Struct(**{ 'RGX': '(!)(\s)?' })
    ]
def parsing(cmd):
    for i in rgx:
        r = re.findall(i.RGX,cmd)
        if len(r)==1:
            if isinstance(r[0], str):
                r = r.pop()
                return {"CMD": cmd, "RGX": r, "TYPE": ast.AST_MAP[r] if r in ast.AST_MAP else r}
            else:
                r = r.pop()
                return {"CMD": cmd, "RGX": [i for i in r] , "TYPE": ast.AST_MAP[r[0]] if r[0] in ast.AST_MAP else r[0]}
    return {"CMD": cmd, "RGX": [], "TYPE": None }


def gen_graph(l):
    iterator_cfg = obj.generic(-1);
    cfg_new = []
    for i in iterator_cfg:
        if i<len(l):
            cfg_node = l[i]
            cfg_node.Node = None
            if cfg_node.TYPE == ast.AST.svi:
                if l[i+1].TYPE == ast.AST.ip:
                    cfg_node.Node = [ l[ iterator_cfg.next() ] ]
                    cfg_new.append(cfg_node);
            elif cfg_node.TYPE in ast.AST_NODE:
                cfg_tmp = []
                for j in iterator_cfg:
                    if j<len(l):
                        if l[j].TYPE==ast.AST.exit:
                            break
                        elif l[j].TYPE!=ast.AST_NOTYPE:
                            if l[j].TYPE in ast.AST_PORT_MAP:
                                cfg = parsing_port(l[j].RGX[4])
                                l[j].VLAN=cfg.VLAN
                            rgx = l[j].RGX
                            if l[j].TYPE == ast.AST.switchport:
                                l[j].TYPE=ast.AST_MAP[f"{rgx[0]}_{rgx[1]}_{rgx[5]}"];
                            cfg_tmp.append(l[j])
                    else:
                        break
                cfg_node.Node = cfg_tmp
                cfg_new.append(cfg_node)
            else:
                if cfg_node.TYPE!=ast.AST_NOTYPE and cfg_node.TYPE!=ast.AST.exit:
                    cfg_new.append(cfg_node)
        else:
            break
    return cfg_new



def show_node(cfg,debug=False):
    print("*"*45);
    for i in cfg:
        if i.Node!=None:
            print(i.CMD,i.RGX,i.TYPE);
            for j in i.Node:
                if j.TYPE!=None:
                     print("   ",j.__dict__)
                elif debug==True:
                     print("   ",j.__dict__)
        elif i.TYPE!=None:
            print(i.CMD,i.RGX,i.TYPE)
        else:
            if debug==True:
                print("WILL REMOVED", i.__dict__)
    print("*"*45);
