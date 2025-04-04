from utils import * 

# SolutionA:    2   485
# SolutionB:    

setrange = tuple[int,int]
rangepair = tuple[setrange,setrange]

def input04(full: bool) -> list[rangepair]:
    lines = readFile(getPath(4, full))
    return [parseRangePair(x) for x in lines]

def day04A():
    full = True
    count = 0 
    for pair in input04(full):
        if isSupersetPair(pair):
            count += 1 
    print('Count:', count)

def parseRangePair(text: str) -> rangepair:
    p = [x.strip() for x in text.split(',')]
    return (parseRange(p[0]), parseRange(p[1]))

def parseRange(text: str) -> setrange:
    p = [int(x.strip()) for x in text.split('-')]
    return (p[0], p[1])

def isSupersetPair(p: rangepair) -> bool:
    return isSuperset(p[0], p[1]) or isSuperset(p[1], p[0])

def isSuperset(r1: setrange, r2: setrange) -> bool:
    return r1[0] <= r2[0] and r2[1] <= r1[1]


if __name__ == '__main__':
    day04A()