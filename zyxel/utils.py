try:
    import os
except Exception as e:
    print("INIT type error: " + str(e),__file__)
    import sys
    sys.exit(1)

def load_config_from_file(name):
    if len(name)>0:
        if os.path.exists(name):
            f = open(name, 'r+')
            cfg = ''.join(f.readlines())
            f.close()
            return cfg
        else:
            print(f"File name not found {name} in load_config_from_file()");
            return None
    else:
        print("File name not set in var name in load_config_from_file()");
        return None

