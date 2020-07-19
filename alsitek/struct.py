try:
    from pprint import pprint as dump
    #import AST as ast
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)

class Struct:
    def __init__(self, **entries):
        self.__dict__.update(entries)
def generic(s: int):
    def init(self,s): self.s = s;
    def next(self): self.s += 1; return self.s;
    return type("gen", (), { "__init__": init, "__iter__": lambda self: self, "__next__": next, "next": next })(s)
