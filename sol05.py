from utils import *

# SolutionA:    CMZ     HNSNMTLHQ
# SolutionB:    MCD     RNLFDJMCT

def input05(full: bool) -> tuple[list,list]:
    stacks, moves = [], []
    stackMode = True
    for line in readFile(getPath(5, full), strip=False):
        clean = line.strip() 
        if clean == '':
            stackMode = False 
            continue 
        if stackMode:
            if not clean.startswith('['):
                continue 
            if len(stacks) == 0:
                count = len(line) // 4 
                for i in range(count):
                    stacks.append([])
            for i,char in enumerate(line):
                if i%4 != 1 or char == ' ':
                    continue 
                idx = i // 4
                stacks[idx].append(char)
        else:
            moves.append(parseMove(line)) 
    return (stacks, moves) 

def day05A():
    full = True 
    processMoves(input05(full), True)

def day05B():
    full = True 
    processMoves(input05(full), False)

def parseMove(line: str) -> tuple[int,int,int]:
    p = [x.strip() for x in line.split('from')]
    count = int(p[0].split()[1])
    i = [int(x.strip()) for x in p[1].split('to')]
    return (count,i[0]-1,i[1]-1)

def processMoves(data: tuple[list,list], reverse: bool):
    transferFn = transferReverse if reverse else transferAsIs
    stack, moves = data 
    for move in moves:
        count, idx1, idx2 = move 
        stack[idx1], stack[idx2] = transferFn(count, stack[idx1], stack[idx2])
        
    top = [s[0] for s in stack]
    print('Top:', ''.join(top))

def transferReverse(count: int, s1: list, s2:list) -> tuple[list, list]:
    move = s1[:count][::-1]
    move.extend(s2)
    return (s1[count:], move)

def transferAsIs(count: int, s1: list, s2: list) -> tuple[list, list]:
    move = s1[:count][:]
    move.extend(s2)
    return (s1[count:], move)

if __name__ == '__main__':
    # day05A()
    day05B()