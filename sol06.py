from utils import *

# SolutionA: 7,5,6,10,11    1140
# SolutionB: 19,23,23,29,26	3495

def input06(full: bool) -> list[str]:
    return readFile(getPath(6, full))

def day06A():
    full = True 
    for line in input06(full):
        print("Marker:", findMarker(line, 4))

def day06B():
    full = True 
    for line in input06(full):
        print("Marker:", findMarker(line, 14))

def findMarker(line: str, numUnique: int) -> int:
    count = len(line)
    for n in range(numUnique,count+1):
        if allUnique(line[n-numUnique:n]):
            return n 

def allUnique(text: str) -> bool:
    return len(text) == len(set(text))

if __name__ == '__main__':
    # day06A()
    day06B()