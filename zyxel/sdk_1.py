import re
from pprint import  pprint as dump

test_str = "1-5,7-9,11,13-15,17,19-27"
def parsing(test_str):
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
  return {"CONFIG":test_str,"PORT":get(matches)}
dump(parsing(test_str))
