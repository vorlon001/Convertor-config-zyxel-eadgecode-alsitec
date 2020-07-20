try:
    import re
    import copy
    import json
    from pprint import pprint as dump
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)

#######################################################################################################################
class Struct:
    def __init__(self, **entries):
        self.__dict__.update(entries)

#######################################################################################################################

class Object:
    def  toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
            sort_keys=True, indent=4)

#######################################################################################################################

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


    def show_tree(self,node=None,i=0):
      node = self if node==None else node;
      print("\t"*(i),node.getdata(),node.gettype());
      childNodes = node.getchildNodes()
      if childNodes!=[]:
        for j in childNodes:
          self.show_tree(j,i+1)

    def get_dom(self,dom,node=None,i=0):

      node = self if node==None else node;

      dom_current = {'TYPE': node.gettype(),  'DATA': node.getdata(), 'CHILD': None }
      if node.gettype()=="FUNCTION":
          dom_current.update({ 'ARG': [ i.__dict__ for i in node.ARG.getchildNodes()] if 'ARG' in node.__dict__ else None })

      dom.append(dom_current);
      childNodes = node.getchildNodes()
      if childNodes!=[]:
        child=[]
        for j in childNodes:
          self.get_dom(child,j,i+1)
        dom_current.update({ 'CHILD': child })
      elif i>0:
        return dom_current
      else:
        return True

    def getchildNodes(self):
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
    def getparentNode(self):
      return self.__parentNode

    def setdata(self,data):
      self.__data = data
    def getdata(self):
      return self.__data

    def settype(self, type):
      self.__type_node = type
    def gettype(self):
      return self.__type_node

    def show(self):
      print(self.__data)

#######################################################################################################################

test_str = "66.666*cos(d)*((3-4^2)/2*(21%4-3))/cos(40)+(((20^2*cos(2) - 10 ) * (30%5 - 20     ) / 10 +35+66)+ asd.fg*math.cos(33*d,55,\"   eeee\") ) * 2 - 2345 +23456"


class Graph:

    __regex = r"[()]|[\=\s0-9a-zA-Z_\"\.\-\+\/\*\^\%]+"
    __regex__parsing_step_1 = r"[\%\^+\/*()-]|[0-9a-zA-Z_\.\"\.]+"

    def __init__( self, source ):

        self.__source = source
        self.__var  = [ "d", "asd.fg"          ]
        self.__vars_info = { "d": { "Type": "FLOAT", "VARS": 2345.2342 } , "asd.fd": { "Type": "INT", "VARS": 234523 }}
        self.__func = [ "math.cos", "cos"      ]

    # сборщик графа или компилятор байт кода из свернутого графа
    def __get_dist(self,dist,i):
        return dist[i];

    # возвращяет тип ноды нода или переменная/оператор
    def  __get_node( self, id = 0 ):
        id_new = id
        Nodes = {}
        nodes = Node( f"NODE", f"NODE" )
        id_var = 0
        if self.__code[id] == "(":
            id_new += 1
            while id_new < self.__max_code:
                id_new, n = self.__get_node( id_new);
                if n!=None:
                    nodes.appendChild(n)
                    id_var += 1
                    id_new += 1
                else:
                    break
            return id_new, nodes
        elif self.__code[id] == ")":
            return id, None
        elif self.__code[id].isdigit() == True:
            return id, Node( f"VARS", self.__code[id] )
        elif isinstance( self.__code[id_new], str ) == True:
            return id, Node( f"VARS", self.__code[id] )
        elif self.__code[id] in ["+","-","*","/","^","%"]:
            return id, Node( f"VARS", self.__code[id] )


    # функция определения тип переменная/оператор/функция
    def __isVars(self,v):
        def is_float(value):
            try:
                float(value)
                return True
            except:
                return False
        def is_int(value):
            try:
                int(value)
                return True
            except:
                return False

        if  is_int(v):
            return "INT"
        elif is_float(v):
            return "FLOAT"
        elif v in ["="]:
            return "EQUAL"
        elif v in ["+","-","*","/","^","%"]:
            return "OPERATOR"
        elif isinstance(v,str) == True and '"' not in v:
            if v in self.__var:
                return "VARS"
            elif v in self.__func:
                return "FUNCTION"
            return "VARS not YET SET"
        elif isinstance(v,str) == True:
             return "STRING"
        return None



    # парсинг элементов строк в переменные, разбираем строки в графе на элементарные операции и данные
    # рекурсивно идем по графу
    def __parsing_step_1(self, node, depth=0):
        c = []
        for v in node.getchildNodes():
            if v.gettype()=="VARS":
                if '"' in v.getdata():
                    c = [Node("STRING", v.getdata() )]
                else:
                    matches = re.finditer( self.__regex__parsing_step_1, v.getdata(), re.MULTILINE)
                    c = [ Node( self.__isVars(m.group()), m.group())  for k,m in enumerate(matches, start=1) ]
                    parentNode = node.getparentNode()
                for i in c:
                    node.insertBefore( i , v );
                node.removeChild( v )
            elif v.gettype()=="NODE":
                self.__parsing_step_1(v,depth+1)

    # валидация и поиск функций, формирование в функции аргументов, исправление графа
    def __parsing_step_2(self, d , node = None , idnode = 0, path = [],depth = 0 ):
        for v in d.getchildNodes():
            if v.gettype() == "FUNCTION":
                next = v.getnextNode()
                v.ARG = copy.deepcopy(next)
                next.settype("DELETE")
            elif v.gettype() == "NODE":
                self.__parsing_step_2(v, node= d, idnode = 0, path = 0, depth = depth+1)
    # удаление нод после свертки функций
    def __parsing_step_2_1(self, d , node = None , idnode = 0, path = [],depth = 0 ):
        for v in d.getchildNodes():
            if v.gettype() == "DELETE":
                d.removeChild( v )
            elif v.gettype() == "NODE":
                self.__parsing_step_2_1(v, node= d, idnode = 0, path = 0, depth = depth+1)



    # свертка графа рекурсивно
    def __parsing_step_3(self, graph_distance, code, d , path="0"):
        for k,v in enumerate(d.getchildNodes()):
             delta = f"{0 if k<10 else ''}{str(k)}"
             if v.gettype()!="NODE":
                 graph_distance.update({ path+delta: v});
             else:
                 _ = self.__parsing_step_3(graph_distance, code,v, path + delta )
        return graph_distance

    def __dist_graph(self):
        # МАТРИЦА - расчетов, вычиление длины всех плеч графа
        dist,min_dist,max_dist = { },9999,0
        for k,v in self.__graph_distance.items():
            d = len(k)
            if d>max_dist:
                max_dist = d
            if d<min_dist:
                min_dist = d
            if d not in dist:
                dist[d] = []
            dist[d].append(k)

        self.__min_dist = min_dist
        self.__max_dist = max_dist
        dist = { k:dist[k] for k in  sorted(dist)};

        self.__z= sorted(dist.keys())
        self.__z.reverse()
        self.__dist = copy.deepcopy(dist)



    def decode(self):

        matches = re.finditer(self.__regex, self.__source, re.MULTILINE)
        self.__code = [ m.group() for _,m in enumerate(matches, start=1)]

        i,Nodes,id_new,id_var = 0,{},0,0
        self.__max_code = len(self.__code);

        html = Node("HTML","HTML")
        body = Node("BODY","BODY")
        html.appendChild(body);

        while id_new < self.__max_code:
            id_new, n = self.__get_node( id=id_new);
            if n!=None:
                body.appendChild(n);
                id_new += 1
                id_var += 1
            else:
                break

        self.__parsing_step_1(body)
        self.__parsing_step_2(body, body)
        self.__parsing_step_2_1(body, body)
        self.__parsing_step_2_1(body, body)

        self.__graph_distance = {}
        self.__graph_distance = self.__parsing_step_3( self.__graph_distance, body, body)
        self.__dist_graph()

        print("="*153)
        print(self.__source)
        print("="*153)

        dom = []
        result = body.get_dom(dom)
        body.show_tree();

        ### def вызов из run функции compiler_1
        for i in self.__z:
            print("="*10,i,"="*140)
            f = {}
            for k in self.__get_dist(self.__dist,i):
                if k[:len(k)-2] not in f:
                    f[k[:len(k)-2]]=[]
                f[k[:len(k)-2]].append(k)
            if i>=self.__min_dist:
                for j in f:
                    vars_stack_id = 0
                    # find and run function
                    for l in f[j]:
                        if l in self.__graph_distance:
                            type = self.__graph_distance[l].gettype()
                            if type=="FUNCTION":
                                vars_id = f"VARS_stack_{vars_stack_id}"
                                print(vars_id,"= CALL ", self.__graph_distance[l].gettype(),self.__graph_distance[l].getdata(),self.__graph_distance[l].ARG )
                                self.__graph_distance[l].settype("VARS")
                                self.__graph_distance[l].setdata(vars_id)
                                vars_stack_id += 1
                    node__body_tmp = Node("BODY",None)
                    for l in f[j]:
                        node_tmp = copy.deepcopy(self.__graph_distance[l] )

                        node_tmp.flash_nodes()

                        node__body_tmp.appendChild( node_tmp )
                    child_tmp = node__body_tmp.getchildNodes()
                    # compiller

                    def find_mux_node(node,o):
                        child = node.getchildNodes()
                        for v in child:
                            if v.gettype()=='OPERATOR' and v.getdata() in o:
                                return v
                        return None

                    for h in [["^","%"],["*","/"],["+","-"]]:
                        status = True
                        while status == True:
                            id = find_mux_node(node__body_tmp,h)
                            if id==None or vars_stack_id>999:
                                status = False
                            else:
                                vars_id = f"VARS_stack_{vars_stack_id}"
                                next_id = id.getnextNode()
                                prev_id = id.getprevNode()
                                print(vars_id,"=","!!",prev_id.gettype(),prev_id.getdata(),"!!",id.gettype(),id.getdata(),"!!",next_id.gettype(),next_id.getdata(),"!!")
                                parentNode = id.getparentNode()
                                node_vars_id = Node("VARS",vars_id)
                                parentNode.insertBefore( node_vars_id , prev_id )
                                parentNode.removeChild(prev_id)
                                parentNode.removeChild(id)
                                parentNode.removeChild(next_id)
                                vars_stack_id += 1
                    print(f"VAR_TMP_{j}"," = ",f"VARS_stack_{vars_stack_id-1}")


                    if int(j)==0:
                        print(f"FINAL RESULT RETURN",j,f" in VARS_TMP_{j}")
                    elif len(f[j]) == 2:
                        self.__dist[len(j)].append(j)
                        self.__dist[len(j)] = sorted(self.__dist[len(j)])
                    elif len(f[j])>2:
                        self.__dist[len(j)].append(j)
                        self.__dist[len(j)] = sorted(self.__dist[len(j)])
                    vars_tmp = Node(j,f"VARS_{j}")
                    self.__graph_distance[j] = vars_tmp

g = Graph(test_str)
g.decode();
