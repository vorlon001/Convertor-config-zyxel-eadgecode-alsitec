import re
import copy
import json
from pprint import pprint as dump



#######################################################################################################################

test_str = "66.666*cos(d)*((3-4^2)/2*(21%4-3))+(((20^2 - 10 ) * (30%5 - 20     ) / 10 +35+66)+ asd.fg*math.cos(33*d,55,\"   eeee\") ) * 2 - 2345 +23456"


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
    def  __get_node(self, id=0):
        id_new = id
        Node = {}
        id_var = 0
        if self.__code[id] == "(":
            id_new += 1
            while id_new < self.__max_code:
                h, id_new = self.__get_node( id_new);
                if h!=None:
                    if isinstance(h,dict):
                        type_var = f"NODE"
                    else:
                        type_var = f"VARS"
                    Node[id_var] = h
                    id_var += 1
                    id_new += 1
                else:
                    break
            return Node, id_new
        elif self.__code[id] == ")":
            return None,id
        elif self.__code[id].isdigit() == True:
            return self.__code[id],id
        elif isinstance(self.__code[id_new],str) == True:
            return self.__code[id],id
        elif self.__code[id] in ["+","-","*","/","^","%"]:
            return self.__code[id],id

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
    def __parsing_step_1(self, d, depth=0):
        for k,v in sorted(d.items(),key=lambda x: x[0]):
            if isinstance(v, dict):
                self.__parsing_step_1(v,depth+1)
            else:
                if '"' in v:
                    c = { "Type":"STRING", "VARS":v }
                else:
                    matches = re.finditer( self.__regex__parsing_step_1, v, re.MULTILINE)
                    c = { k: { "Type": self.__isVars(m.group()), "VARS": m.group()} for k,m in enumerate(matches, start=1)}
                if len(c)>0:
                    d.update({k: c});


    # валидация и поиск функций, формирование в функции аргументов, исправление графа
    def __parsing_step_2(self, d , node = None , idnode = 0, path = [],depth = 0 ):
        for k,v in sorted(d.items(),key=lambda x: x[0]):
            if "Type" in v:
                if v["Type"]=="FUNCTION":
                    v.update({"ARG": node[idnode+1]});
                    del(node[idnode+1])
            elif isinstance(v, dict):
                p = copy.deepcopy(path)
                p.append(k)
                trace_child = self.__parsing_step_2(v, node= d, idnode = k, path = p, depth = depth+1)
            else:
                del(d[k])


    # свертка графа рекурсивно
    def __parsing_step_3(self, graph_distance, code, d , path="0"):
        for k,v in sorted(d.items(),key=lambda x: x[0] ):
            status =False
            for f,j in sorted(v.items(),key=lambda x: x[0] ):
                if "Type" in j:
                    status = True
            if status==True:
                graph_distance.update({ path+str(k): v});
            if status == False:
                _ = self.__parsing_step_3(graph_distance, code,v, path + str(k) )
        return graph_distance

    def __dist_graph(self):

        ## def calc_dist_graph(
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

        self.__z=list(range(self.__min_dist,self.__max_dist+1))
        self.__z.reverse() #реверс массива
        self.__dist = copy.deepcopy(dist)


    #######################################################################################################################
    #  компилятор выражений без скобок
    #
    def __linker_1( self, o ):
        b = copy.deepcopy(o)

        def find_function(f):
          for k,v in f.items():
            if v['Type'] == 'FUNCTION':
              return k
          return None


        def find_prev_id(f,id):
          j=id-1
          while j>=0 and j not in f:
            j -= 1
          if j==0:
            return None
          else:
            return j


        def find_mux(a,o):
            for k,v in a.items():
              if v['Type']=='OPERATOR' and v['VARS'] in o:
                  return k
            return None

        vars_stack_id = 0
        status = True
        while status == True:
          id = find_function(b)
          if id==None or vars_stack_id>999:
            status = False
          else:
            vars_id = f"VARS_stack_{vars_stack_id}"
            print(vars_id,"= CALL ",b[id])
            b[id]={'Type': 'VARS', 'VARS': vars_id}
            vars_stack_id += 1

        for h in [["^","%"],["*","/"],["+","-"]]:
          status = True
          while status == True:
              id = find_mux(b,h)
              if id==None or vars_stack_id>999:
                status = False
              else:
                vars_id = f"VARS_stack_{vars_stack_id}"
                prev_id = find_prev_id(b,id)
                print(vars_id,"=",b[prev_id],b[id],b[id+1])
                b[prev_id]={'Type': 'VARS', 'VARS': vars_id}
                del(b[id])
                del(b[id+1])
                vars_stack_id += 1

        return b[1];


    def decode(self):
        matches = re.finditer(self.__regex, self.__source, re.MULTILINE)
        self.__code = [ m.group() for _,m in enumerate(matches, start=1)]

        i,Node,id_new,id_var = 0,{},0,0

        self.__max_code = len(self.__code);
        while id_new < self.__max_code:
            h, id_new = self.__get_node( id=id_new);
            if h!=None:
                if isinstance(h,dict):
                    type_var = f"NODE"
                else:
                    type_var = f"VARS"
                Node[id_var] = h
                id_new += 1
                id_var += 1
            else:
                break

        self.__Node = Node

        print("!"*153)
        dump(self.__Node)


        self.__parsing_step_1(self.__Node)

        print("!"*153)
        dump(self.__Node)
        print("!"*153)

        self.__parsing_step_2(self.__Node, self.__Node)

        self.__graph_distance = {}
        self.__graph_distance = self.__parsing_step_3( self.__graph_distance, self.__Node, self.__Node)
        self.__dist_graph()

        dump(self.__Node)
        dump(self.__graph_distance)
        dump(self.__dist);

        print("="*153)

        ### def вызов из run функции compiler_1
        for i in self.__z:
            print("="*10,i,"="*140)
            f = {}
            for k in self.__get_dist(self.__dist,i):
                if k[:len(k)-1] not in f:
                    f[k[:len(k)-1]]=[]
                f[k[:len(k)-1]].append(k)
            if i>=self.__min_dist:
                for k in f:
                    if len(f[k])>1:
                        u,y =1,{}
                        for z in f[k]:
                            if z in self.__graph_distance:
                                for h,g in self.__graph_distance[z].items():
                                    y.update({u: g})
                                    u += 1
                            else:
                                y.update({u: {'Type': 'VARS', 'VARS': f'VARS_TMP_{z}' }})
                                u += 1
                        print(f"VARS_TMP_{k}","=",self.__linker_1(y));
                    elif len(f[k])==1:
                        print( f"VARS_TMP_{f[k][0][:len(f[k][0])-1]}", "=", self.__linker_1(self.__graph_distance[f[k][0]]))
                    print("-"*23);
                    if int(k)==0:
                        print(f"FINAL RESULT RETURN",k,f" in VARS_TMP_{k}")
                    elif len(f[k]) == 1:
                        self.__dist[len(k)].append(k)
                        self.__dist[len(k)] = sorted(self.__dist[len(k)])
                    elif len(f[k])>1:
                        self.__dist[len(k)].append(k)
                        self.__dist[len(k)] = sorted(self.__dist[len(k)])
        print("="*153)



g = Graph(test_str)
g.decode();
