
def generic(s: int):
    def init(self,s): self.s = s;
    def next(self): self.s += 1; return self.s;
    return type("gen", (), { "__init__": init, "__iter__": lambda self: self, "__next__": next, "next": next })(s)

r = generic(55);
print(r.next())
#for i in r:
#    print(i);

