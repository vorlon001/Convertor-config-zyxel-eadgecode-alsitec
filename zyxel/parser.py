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
    return { i:i for i in f }
  return obj.Struct(**{"CONFIG":test_str,"PORT":get(matches)})


rgx = [
	obj.Struct(**{ 'RGX': '^(hostname)\s+(.+)'  }),
	obj.Struct(**{ 'RGX': '^(interface)\s+(port-channel)\s+(\d+)' }),
	obj.Struct(**{ 'RGX': '^(vlan)\s+(\d+)' }),
	obj.Struct(**{ 'RGX': '^(mvr)\s+(\d+)' }),
	obj.Struct(**{ 'RGX': '^(fixed)\s+([0-9\,\-]+)'}),
	obj.Struct(**{ 'RGX': '^(ip)\s+(name-server)\s+([0-9\.]+)' }),
	obj.Struct(**{ 'RGX': '^(dhcp snooping)\s+(vlan)\s+([0-9]+)' }),
	obj.Struct(**{ 'RGX': '^(snmp-server)\s+(trap-destination)\s+([0-9\.]+)\s+(enable.+)' }),
	obj.Struct(**{ 'RGX': '^(snmp-server)\s+(trap-destination)\s+([0-9\.]+)' }),
	obj.Struct(**{ 'RGX': '^(port-security)\s+(\d+)\s+(address-limit)\s+(\d+)' }),
	obj.Struct(**{ 'RGX': '^(pvid)\s+(\d+)' }),
	obj.Struct(**{ 'RGX': '^(ip)\s+(address)\s+(default-management)\s+([0-9\.]+)\s+([0-9\.]+)' }),
	obj.Struct(**{ 'RGX': '^(ip)\s+(address)\s+(default-gateway)\s+([0-9\.]+)' }),
	obj.Struct(**{ 'RGX': '^(forbidden)\s+([0-9\,\-]+)'}),
	obj.Struct(**{ 'RGX': '^(untagged)\s+([0-9\,\-]+)'}),
	obj.Struct(**{ 'RGX': '^(source-port)\s+([0-9\,\-]+)'}),
	obj.Struct(**{ 'RGX': '^(receiver-port)\s+([0-9\,\-]+)'}),
	obj.Struct(**{ 'RGX': '^(tagged)\s+([0-9\,\-]+)'}),
	obj.Struct(**{ 'RGX': '^(frame-type)\s+(untagged)'}),
	obj.Struct(**{ 'RGX': '^(frame-type)\s+(tagged)'}),
	obj.Struct(**{ 'RGX': '^(inactive)' }),
	obj.Struct(**{ 'RGX': '^(syslog)\s+(type)\s(.+)'}),
	obj.Struct(**{ 'RGX': '^(syslog)\s+(server)\s+([0-9\.]+)\s+(level)\s+(\d)'}),
	obj.Struct(**{ 'RGX': '^(port-security)\s+(\d+)'}),
	obj.Struct(**{ 'RGX': '^(name)\s+(.+)'}),
        obj.Struct(**{ 'RGX': '^(exit)?(\s)?'})
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
            if cfg_node.TYPE in ast.AST_NODE:
                cfg_tmp = []
                for j in iterator_cfg:
                    if j<len(l):
                        if l[j].TYPE==ast.AST.exit:
                            break
                        elif l[j].TYPE!=ast.AST_NOTYPE:
                            if l[j].TYPE in ast.AST_PORT_MAP:
                                cfg = parsing_port(l[j].RGX[1])
                                l[j].PORT=cfg.PORT
                            cfg_tmp.append(l[j])
                    else:
                        break
                cfg_node.Node = cfg_tmp
                cfg_new.append(cfg_node)
            else:
                if l[i].TYPE!=ast.AST_NOTYPE:
                    cfg_new.append(cfg_node)
        else:
            break
    return cfg_new



def show_node(cfg):
    print("*"*45);
    for i in cfg:
        if i.Node!=None:
            print(i.CMD,i.RGX);
            for j in i.Node:
                print("   ",j.__dict__)
        else:
            print(i.CMD,i.RGX)
    print("*"*45);
