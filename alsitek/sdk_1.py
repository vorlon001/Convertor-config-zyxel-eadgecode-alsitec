import re
import struct as obj
from pprint import  pprint as dump


#test_str = "e(1-10,22),g(1-4)"
test_str = "e1,g2"

def parsing(test_str):
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

  _,_,_,_,_ = 	(e:=[]), e.extend(get(matches[1])) if matches[0]=='e' and len(matches[1])>0 else 0, \
  		{e.extend([ i+24 for i in get(matches[3])]) if matches[2]=='g' and len(matches[3])>0 else 0}, \
 		(e_n:=[]),{ i:e_n.append(i) if i>0 else i for i in  e }

  return obj.Struct(**{"CONFIG":test_str,"PORT":e_n });
test_str = "e(1-10,22),g(1-4)"
dump(parsing(test_str).__dict__)
test_str = "e1,g2"
dump(parsing(test_str).__dict__)
test_str = "all"
dump(parsing(test_str).__dict__)




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

test_str = "1-22,43,55,45"
dump(parsing_port(test_str).__dict__)
test_str = "3"
dump(parsing_port(test_str).__dict__)

