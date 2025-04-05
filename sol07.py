from utils import *

# SolutionA: 95437  1182909

class Item:
    def __init__(self, name: str, path: str, size: int):
        self.name = name 
        self.path = path 
        self.size = size
        self.parent: Item|None = None 
        self.children: list[Item]|None = None

    @property 
    def isDir(self) -> bool:
        return self.children is not None
    
    def computeSize(self):
        if self.children is None: return
        size = 0
        for child in self.children:
            size += child.size
        self.size = size
        
fs: dict[str,Item] = {}
cmdCD = "$ cd"
cmdLS = "$ ls"
glue = "/"

def input07(full: bool) -> list[str]:
    return readFile(getPath(7, full))

def day07A():
    full = True 
    buildFS(input07(full))

    limit = 100_000 
    total = 0 
    for _, item in fs.items():
        if item.isDir and item.size <= limit:
            total += item.size 
    print('Total:', total)

def buildFS(lines: list[str]):
    cwd = None 
    for line in lines:
        if line.startswith(cmdCD):
            name = line.split()[2]
            if name == "..":
                cwd = cwd.parent
            else:
                cwd, _ = getDir(name, cwd)
        elif line == cmdLS:
            continue 
        else:
            p = line.split()
            if p[0] == 'dir':
                item, isNew = getDir(p[1], cwd)
                if isNew:
                    cwd.children.append(item)
            else:
                item, isNew = getFile(p[1], int(p[0]), cwd)
                if isNew:
                    cwd.children.append(item)
    
    dirPaths = []
    for path, item in fs.items():
        if item.isDir:
            dirPaths.append(path)
    dirPaths.sort(key=lambda p: len(p.split(glue)), reverse=True)
    for path in dirPaths:
        fs[path].computeSize()

def getDir(name: str, parent: Item|None) -> tuple[Item, bool]:
    path = name if parent is None else parent.path + name + glue 
    item = fs.get(path)
    if item is not None:
        return item, False
    item = Item(name, path, 0)
    item.parent = parent 
    item.children = []
    fs[path] = item 
    return item, True

def getFile(name: str, size: int, parent: Item|None) -> tuple[Item, bool]:
    path = parent.path + name 
    item = fs.get(path)
    if item is not None:
        return item, False 
    item = Item(name, path, size)
    item.parent = parent 
    fs[path] = item 
    return item, True

if __name__ == '__main__':
    day07A()