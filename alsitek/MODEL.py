try:
    import json
    from pprint import pprint as dump
    from enum import Enum
    import AST as ast
    import struct as obj
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)

class Object:
    def  toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, 
            sort_keys=True, indent=4)

class PORT(Object):
    def __init__(self):
        self.description = None
        self.native_vlan = None
        self.status = 'no shutdown'

        self.general_untag = []
        self.general_tag   = []
        self.general_pid   = []

        self.tagged = []
        self.untagged = []
        self.mvr = obj.Struct(**{"source_port": None, "receiver_port": None, "tag": None });
    def setdescription(self,description):
        self.description = description
    def setdown(self):
        self.status = 'shutdown'

    def addgeneral_untag(self,tag):
        self.general_untag.append(tag)
    def addgeneral_tag(self,tag):
        self.general_tag.append(tag)
    def addgeneral_pid(self,tag):
        self.general_pid.append(tag)


    def addtagged(self,tag):
        self.tagged.append(tag)
    def adduntagged(self,tag):
        self.native_vlan = tag
        self.untagged = [tag]

class VLAN(Object):
    def __init__(self,name, tag):
        self.name = name
        self.tag  = tag

class MVR(Object):
    def __init__(self,name,tag):
        self.name = name
        self.tag  = tag

class DEVICE(Object):
    def __init__(self):
        self.default_gateway = None
        self.dhcp_snooping = []
        self.hostname = None
        self.igmp_snooping = []
        self.mng_int_vlan = ''
        self.mng_ip = ''
        self.mng_mask = ''
        self.port = { v:PORT() for v in range (1,29) }
        self.vlan = {}
        self.mvr = None
    def getport(self,id):
        return self.port[id];
    def getvlan(self,id):
        return self.vlan[id];
    def setvlanname(self,id,name):
        self.vlan[int(id)].name = name
    def getmvr(self):
        return self.mvr;
    def setmvr(self,mvr):
        self.mvr = mvr
    def addvlan(self,vlan):
        self.vlan.update({ vlan.tag: vlan })
    def toJSONs(self):
        r = self.toJSON()
        return r
