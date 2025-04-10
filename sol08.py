from utils import *

# SolutionA: 21     1785
# SolutionB: 8      345168

def input08(full: bool) -> list[list[int]]:
    lines = readFile(getPath(8, full))
    return [[int(x) for x in line] for line in lines]

def day08A():
    full = True 
    countVisible(input08(full))

def day08B():
    full = True 
    findBestScore(input08(full))

def countVisible(grid: list[list[int]]):
    nRows, nCols = len(grid), len(grid[0])
    visible = set() 

    for row in range(1, nRows-1):
        for col in range(1, nCols-1):
            checkRowVisible(grid, visible, row, col)
            checkColVisible(grid, visible, row, col)
    
    numEdges = (2 * nCols) + (2 * (nRows-2))
    print('Visible:', numEdges + len(visible))

def checkRowVisible(grid: list[list[int]], visible: set, row:int, col:int):
    if (row,col) in visible:
        return 
    isValid = lambda x: x < grid[row][col]
    ok = all(isValid(x) for x in grid[row][:col])
    if ok:
        visible.add((row,col))
        return 
    ok = all(isValid(x) for x in grid[row][col+1:])
    if ok:
        visible.add((row,col))
    
def checkColVisible(grid: list[list[int]], visible: set, row:int, col:int):
    if (row,col) in visible:
        return  
    isValid = lambda x: x < grid[row][col]
    above = [grid[r][col] for r in range(0,row)]
    ok = all(isValid(x) for x in above)
    if ok:
        visible.add((row,col))
        return 
    below = [grid[r][col] for r in range(row+1,len(grid))]
    ok = all(isValid(x) for x in below)
    if ok:
        visible.add((row,col))

def findBestScore(grid: list[list[int]]):
    nRows, nCols = len(grid), len(grid[0])
    best = 0 
    for row in range(1, nRows-1):
        for col in range(1, nCols-1):
            best = max(best, computeScore(grid, row, col))
    print('Best:', best)

def computeScore(grid: list[list[int]], row: int, col:int) -> int:
    nRows, nCols = len(grid), len(grid[0])
    value = grid[row][col]
    # Up 
    n = 0 
    for r in range(row-1,-1,-1):
        n += 1 
        if grid[r][col] >= value: break
    # Down 
    s = 0
    for r in range(row+1, nRows):
        s += 1 
        if grid[r][col] >= value: break
    # Left 
    w = 0 
    for c in range(col-1,-1,-1):
        w += 1 
        if grid[row][c] >= value: break
    # Right
    e = 0 
    for c in range(col+1, nCols):
        e += 1 
        if grid[row][c] >= value: break 
    
    return n*e*w*s

if __name__ == '__main__':
    # day08A()
    day08B()