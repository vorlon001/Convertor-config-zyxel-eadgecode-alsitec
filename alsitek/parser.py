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

def parsing_interface(test_str):
  if test_str=="all":
      return obj.Struct(**{"CONFIG":test_str,"PORT": [i for i in range (1,29) ] });
  regex = r"(e)?\(?([0-9\-\,]+)?\)?,?(g)?\(?([0-9\-\,]+)?\)?"
  matches = re.findall(regex, test_str)
  matches=matches.pop(0) if len(matches)>1 else ('','','','','','','')

  def get(port):
    _,_,_,_,_= (h:=[ [ int(i) if len(i)>0 else 0  for i in x.split('-')] for x in port.split(',')]), \
              (e:=[]), \
              (f:=[]), \
              {e.append([ i for i in range(k[0],k[1]+1 if len(k)==2 else k[0]+1) ]) for k in h}, \
              { f.extend(i) for i in e}
    return f

  _,_,_,_,_ =   (e:=[]), e.extend(get(matches[1])) if matches[0]=='e' and len(matches[1])>0 else 0, \
                {e.extend([ i+24 for i in get(matches[3])]) if matches[2]=='g' and len(matches[3])>0 else 0}, \
                (e_n:=[]),{ i:e_n.append(i) if i>0 else i for i in  e }

  return obj.Struct(**{"CONFIG":test_str,"PORT":e_n });

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
  return obj.Struct(**{"CONFIG":test_str,"VLAN":get(matches)})


rgx = [
        obj.Struct(**{ 'RGX': '^(hostname)\s+?(.+)'  			}),
        obj.Struct(**{ 'RGX': '^(interface\srange\sethernet)\s+(.+)'	}),
        obj.Struct(**{ 'RGX': '^(interface\sethernet)\s+(e|g)(.+)'	}),
        obj.Struct(**{ 'RGX': '^(description)\s+(.+)'			}),
	obj.Struct(**{ 'RGX': '^(line)\s+(.+)'				}),
	obj.Struct(**{ 'RGX': '^(switchport\smode\strunk)'		}),
        obj.Struct(**{ 'RGX': '^(switchport\smode\sgeneral)'		}),
	obj.Struct(**{ 'RGX': '^(switchport\saccess)\s+(vlan)\s+(\d+)'	}),
        obj.Struct(**{ 'RGX': '^(switchport\strunk\snative\svlan)\s+(\d+)'  			}),
	obj.Struct(**{ 'RGX': '^(switchport\strunk\sallowed\svlan\sadd)\s+(\d+)'		}),
        obj.Struct(**{ 'RGX': '^(switchport general pvid)\s+(\d+)'                		}),
	obj.Struct(**{ 'RGX': '^(switchport\sgeneral\sallowed\svlan\sadd)\s+(\d+)\s+(untagged)'	}),
        obj.Struct(**{ 'RGX': '^(switchport\sgeneral\sallowed\svlan\sadd)\s+(\d+)'		}),
        obj.Struct(**{ 'RGX': '^(switchport\sgeneral)\s+(pvid)\s+(\d+)'				}),
	obj.Struct(**{ 'RGX': '^(interface\svlan)\s+(\d+)'			}),
	obj.Struct(**{ 'RGX': '^(name)\s+(.+)'					}),
	obj.Struct(**{ 'RGX': '^(vlan database)'				}),
	obj.Struct(**{ 'RGX': '^(vlan)\s+([0-9\,\-]+)'				}),
	obj.Struct(**{ 'RGX': '^(ip\saddress)\s+([0-9\.]+)\s+([0-9\.]+)'	}),
	obj.Struct(**{ 'RGX': '^(ip\sdefault-gateway)\s+([0-9\.]+)'		}),
	obj.Struct(**{ 'RGX': '^(ip\sigmp\ssnooping)'				}),
	obj.Struct(**{ 'RGX': '^(ip\sigmp\ssnooping)\s+(leave-time-out immediate-leave)'		}),
	obj.Struct(**{ 'RGX': '^(ip\sigmp\ssnooping)\s+(forbidden mrouter ports add ethernet)\s+(.+)'	}),
	obj.Struct(**{ 'RGX': '^(ip\sigmp\ssnooping)\s+(mrouter ports add ethernet)\s+(.+)'		}),
	obj.Struct(**{ 'RGX': '^(ip\sdhcp\ssnooping\spppoe\ssnooping)\s+(vlan)\s+(\d+)'			}),
	obj.Struct(**{ 'RGX': '^(pppoe\ssnooping\svlan)\s+(\d+)'					}),
	obj.Struct(**{ 'RGX': '^(ip\sdhcp\ssnooping)\s+(vlan)\s+(\d+)'					}),
        obj.Struct(**{ 'RGX': '^(exit)(\s)?' 								})
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
            elif cfg_node.TYPE!=None:
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
