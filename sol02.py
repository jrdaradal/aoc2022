from utils import * 

# SolutionA:    15  12740
# SolutionB:    

rps = tuple[int,int]

R, P, S = 1, 2, 3 
L, D, W = 0, 3, 6

winsOver = {R: S, P: R, S: P}
losesTo  = {S: R, R: P, P: S}

def input02(full: bool, mask: dict[str,int]) -> list[rps]:
    def createRPS(line: str) -> rps:
        p = line.split()
        return (mask[p[0]], mask[p[1]])
    return [createRPS(x) for x in readFile(getPath(2, full))]

def day02A():
    full = True 
    mask = {'A': R, 'B': P, 'C': S, 'X': R, 'Y': P, 'Z': S}
    total = 0
    for game in input02(full, mask):
        total += computeGameScore(game)
    print("Total:", total)

def computeGameScore(game: rps) -> int:
    opp, you = game
    score = you 
    if opp == you:
        score += D 
    elif winsOver[you] == opp:
        score += W
    return score

if __name__ == '__main__':
    day02A()