from pprint import pprint as dump
import json

class Struct:
    def __init__(self, **entries):
        self.__dict__.update(entries)

class Object:
    def  toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
            sort_keys=True, indent=4)

class Node:
    def __init__(self,type_node,data=None):
      self.__type_node = type_node
      self.__data = data
      self.__nextNode = None
      self.__prevNode = None
      self.__parentNode  = None
      self.__childNodes = []
      self.__firstChild = None
      self.__lastChild = None

    def get_childNodes(self):
      return self.__childNodes

    def set_childNodes(self,childNodes):
      if isinstance(child,list)==True:
        old__childNodes = self.__childNodes
        for child in childNodes:
          err = self.appendChild(child)
          if err ==False:
            self.__childNodes = old__childNodes
            return False        
        return True
      else:
        return False

    def appendChild(self,child):
      if child!=None and isinstance(child,Node)==True:
        if self.__firstChild==None:
          self.__firstChild = child
          self.__lastChild = child
          child.setparentNode(self)
          self.__childNodes.append(child)
          return True
        else:
          last_node = self.__lastChild
          child.setprevNode(last_node)
          child.setparentNode(self)
          self.__childNodes.append(child)
          last_node.setnextNode(child)
          self.__lastChild = child
          return True
      else:
        return False

    def removeChild(self,child):
      if child in self.__childNodes and isinstance(child,Node)==True:
        removed_child_id = self.__childNodes.index(child)        
        removed_child = self.__childNodes[removed_child_id]
        next_node_removed_child = removed_child.getnextNode()
        prev_node_removed_child = removed_child.getprevNode()
        del self.__childNodes[removed_child_id]
        if next_node_removed_child!=None:
            next_node_removed_child.setprevNode(prev_node_removed_child)
        if prev_node_removed_child!=None:
            prev_node_removed_child.setnextNode(next_node_removed_child)

        return True
      else:
        return False
    
    def insertBefore(self, newElement, referenceElement):
      """
      добавляет элемент в  список дочерних элементов родителя перед указанным элементом.
      """
      if newElement!=None and referenceElement!=None:
        if isinstance(newElement,Node)==True and isinstance(referenceElement,Node)==True:
          referenceElement_id = self.__childNodes.index(referenceElement)

          referenceElement_node = self.__childNodes[referenceElement_id]
          referenceElement_prev_node = referenceElement_node.getprevNode()

          newElement.setparentNode(self)
          newElement.setprevNode(referenceElement_prev_node)
          newElement.setnextNode(referenceElement_node)

          self.__childNodes.insert(referenceElement_id, newElement)

          if referenceElement_node!=None:
            referenceElement_node.setprevNode(newElement)
          if referenceElement_prev_node!=None:
            referenceElement_prev_node.setnextNode(newElement)
          return True
        else:
          return False
      else:
        return False

    def flash_nodes(self):
      self.__nextNode = None
      self.__prevNode = None
      self.__parentNode = None
    def setnextNode(self, nextNode):
      if isinstance(nextNode,Node)==True:
        self.__nextNode = nextNode
        return True
      else:
        return False
    def setprevNode(self,prevNode):
      if isinstance(prevNode,Node)==True:
        self.__prevNode = prevNode
        return True
      else:
        return False

    def setchildNode(self,childNode):
      if isinstance(childNode,Node)==True:
        self.__childNode = childNode
        return True
      else:
        return False

    def setparentNode(self,parentNode):
      self.__parentNode =parentNode

    def getnextNode(self):
      return self.__nextNode
    def getprevNode(self):
      return self.__prevNode
    def getchildrenNode(self):
      return self.__childNode
    def getparentNode(self):
      return self.__parentNode

    def setdata(self,data):
      self.__data = data
    def getdata(self):
      return self.__data

    def gettype(self):
      return self.__type_node

    def show(self):
      print(self.__data)

n1 = Node("n1","node1")

n1_sub1 = Node("n1_sub1","n1_sub1")
n1_sub2 = Node("n1_sub2","n1_sub2")

n1.appendChild(n1_sub1)
n1.appendChild(n1_sub2)

n2 = Node("n2","node2")
n3 = Node("n3","node3")
n4 = Node("n4","node4")

n_addon = Node("node_addon","node_addon")
n_addon_end = Node("node_addon_end","node_addon_end")
n_addon_start = Node("node_addon_start","node_addon_start")

body = Node("BODY","BODY")
body.appendChild(n1)
body.appendChild(n2)
body.appendChild(n3)
body.appendChild(n4)

dump( body.__dict__ );
print("*"*30)
dump(n1.__dict__)
print("*"*30)
dump(n2.__dict__)
print("*"*30)
dump(n3.__dict__)
print("*"*30)
dump(n4.__dict__)
print("*"*30)

print("="*43)
print("INIT")
print("="*43)
n1.show()
n1.getnextNode().show()
print("="*43)
n2.getprevNode().show()
n2.show()
n2.getnextNode().show()
print("="*43)
n3.getprevNode().show()
n3.show()
n3.getnextNode().show()
print("="*43)
n4.getprevNode().show()
n4.show()
print("="*43)

body.insertBefore(n_addon,n2);
print("="*43)
print("body.insertBefore(n_addon,n2);")
print("="*43)
n1.show()
n1.getnextNode().show()
print("="*43)
n_addon.getprevNode().show()
n_addon.show()
n_addon.getnextNode().show()
print("="*43)
n2.getprevNode().show()
n2.show()
n2.getnextNode().show()
print("="*43)
n3.getprevNode().show()
n3.show()
n3.getnextNode().show()
print("="*43)
n4.getprevNode().show()
n4.show()
print("="*43)


body.removeChild(n2);
print("="*43)
print("body.removeChild(n2);")
print("="*43)
n1.show()
n1.getnextNode().show()
print("="*43)
n_addon.getprevNode().show()
n_addon.show()
n_addon.getnextNode().show()
print("="*43)
n3.getprevNode().show()
n3.show()
n3.getnextNode().show()
print("="*43)
n4.getprevNode().show()
n4.show()
print("="*43)

body.insertBefore(n_addon_end,n4);
print("="*43)
print("body.insertBefore(n_addon_end,n4);")
print("="*43)
n1.show()
n1.getnextNode().show()
print("="*43)
n_addon.getprevNode().show()
n_addon.show()
n_addon.getnextNode().show()
print("="*43)
n3.getprevNode().show()
n3.show()
n3.getnextNode().show()
print("="*43)
n_addon_end.getprevNode().show()
n_addon_end.show()
n_addon_end.getnextNode().show()
print("="*43)
n4.getprevNode().show()
n4.show()
print("="*43)


body.insertBefore(n_addon_start,n1);

print("="*43)
print("body.insertBefore(n_addon_start,n1);")
print("="*43)
n_addon_start.show()
n_addon_start.getnextNode().show()
print("="*43)
n1.getprevNode().show()
n1.show()
n1.getnextNode().show()
print("="*43)
n_addon.getprevNode().show()
n_addon.show()
n_addon.getnextNode().show()
print("="*43)
n3.getprevNode().show()
n3.show()
n3.getnextNode().show()
print("="*43)
n_addon_end.getprevNode().show()
n_addon_end.show()
n_addon_end.getnextNode().show()
print("="*43)
n4.getprevNode().show()
n4.show()
print("="*43)

print("="*43)
childNodes = body.get_childNodes();
for i in childNodes:
  dump(i.__dict__);
print("="*43)

print(isinstance(n1,Node))
print("="*43)
def show(node,i=0):
  print("\t"*(i),node.getdata(),node.gettype());
  childNodes = node.get_childNodes()
  if node.get_childNodes()!=[]:
    for j in node.get_childNodes():
      show(j,i+1)

def get_dom(dom,node,i=0):
  dom_current = {'TYPE': node.gettype(),  'DATA': node.getdata(), 'CHILD': None }
  dom.append(dom_current);
  childNodes = node.get_childNodes()
  if node.get_childNodes()!=[]:
    child=[]
    for j in node.get_childNodes():
      child.append(get_dom(child,j,i+1))
    dom_current.update({ 'CHILD': child })
  if i>0:
    return dom_current
  else:
    return True
dom = []
show(body)
result = get_dom(dom,body)
dump(result);
dump(dom);
print("="*43)
