from utils import * 

# SolutionA: 13     6339
# SolutionB: 36     2541

coords = list[int]

def input09(full: bool) -> list[coords]:
    def create(line: str) -> coords:
        p = line.split()
        d = int(p[1])
        if p[0] == 'U':
            return [d, 0]
        elif p[0] == 'D':
            return [-d, 0]
        elif p[0] == 'L':
            return [0, -d]
        elif p[0] == 'R':
            return [0, d]
    return [create(line) for line in readFile(getPath(9, full))]

def day09A():
    full = True 
    head, tail = [0,0], [0,0]
    visited = set([tuple(tail)])
    for delta in input09(full):
        head, tail = moveRope(head, tail, delta, visited)
    print('Visited:', len(visited))

def day09B():
    full = True 
    tail = 9 
    pos = [[0,0] for i in range(tail+1)]
    visited = set()
    visited.add(tuple(pos[tail]))
    for delta in input09(full):
        pos = moveChain(pos, delta, visited)
    print('Visited:', len(visited))

def moveRope(head: coords, tail: coords, delta: coords, visited: set) -> tuple[coords, coords]:
    steps, idx, factor = unpackDelta(delta)
    for i in range(steps):
        head[idx] += 1 * factor 
        if not isAdjacent(head, tail):
            tail = follow(head, tail)
            visited.add(tuple(tail))
    return head, tail

def moveChain(pos: list[coords], delta: coords, visited: set) -> list[coords]:
    tail = len(pos)-1
    steps, idx, factor = unpackDelta(delta) 
    for n in range(steps):
        pos[0][idx] += 1 * factor 
        for i in range(1,tail+1):
            if not isAdjacent(pos[i-1], pos[i]):
                pos[i] = follow(pos[i-1], pos[i])
        visited.add(tuple(pos[tail]))
    return pos

def unpackDelta(delta: coords) -> tuple[int, int, int]:
    dy, dx = delta 
    if dy == 0 and dx > 0:
        return (dx, 1, 1)
    elif dy == 0 and dx < 0:
        return (-dx, 1, -1)
    elif dx == 0 and dy > 0:
        return (dy, 0, 1)
    elif dx == 0 and dy < 0:
        return (-dy, 0, -1)

def isAdjacent(c1: coords, c2: coords) -> bool:
    dy = abs(c1[0]-c2[0])
    dx = abs(c1[1]-c2[1])
    return dy <= 1 and dx <= 1

def follow(c1: coords, c2: coords) -> coords:
    dy = c1[0]-c2[0]
    dx = c1[1]-c2[1]
    if dy > 0:
        c2[0] += 1
    elif dy < 0:
        c2[0] -= 1
    if dx > 0:
        c2[1] += 1
    elif dx < 0:
        c2[1] -= 1
    return c2

if __name__ == '__main__':
    # day09A()
    day09B()