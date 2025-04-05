from utils import *

# SolutionA: 7,5,6,10,11    1140

def input06(full: bool) -> list[str]:
    return readFile(getPath(6, full))

def day06A():
    full = True 
    for line in input06(full):
        print("Marker:", findMarker(line))

def findMarker(line: str) -> int:
    count = len(line)
    for n in range(4,count+1):
        if allUnique(line[n-4:n]):
            return n 

def allUnique(text: str) -> bool:
    return len(text) == len(set(text))

if __name__ == '__main__':
    day06A()