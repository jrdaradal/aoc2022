from utils import * 

# SolutionA:    157     7553
# SolutionB:    

def input03(full: bool) -> list[str]:
    return readFile(getPath(3, full))

def day03A():
    full = True 
    total = 0 
    for line in input03(full):
        total += getPriority(findCommon(line))
    print('Total:', total)

def getPriority(char: str) -> int:
    v = ord(char)
    if 97 <= v and v <= 122:
        return v - 96
    elif 65 <= v and v <= 90:
        return v - 38

def findCommon(line: str) -> str:
    mid = len(line) // 2 
    chars = set()
    for i, char in enumerate(line):
        if i < mid:
            chars.add(char)
        elif char in chars:
            return char 


if __name__ == '__main__':
    day03A()