from utils import *

# SolutionA:    24000   70698
# SolutionB:    45000   206643

def input01(full: bool) -> list[str]:
    return readFile(getPath(1, full))

def day01A():
    full = True 
    maxCalories, current = 0, 0
    for line in input01(full):
        if line == "":
            maxCalories = max(maxCalories, current)
            current = 0 
        else:
            current += int(line)
    maxCalories = max(maxCalories, current)
    print("MaxCalories:", maxCalories)

def day01B():
    full = True 
    top3, current = [0,0,0], 0
    for line in input01(full):
        if line == "":
            top3 = adjustTop3(top3, current)
            current = 0
        else:
            current += int(line)
    top3 = adjustTop3(top3, current)
    print("Total:", sum(top3))

def adjustTop3(top3: list[int], current: int) -> list[int]:
    top3.append(current)
    top3.sort()
    return top3[1:4]

if __name__ == '__main__':
    # day01A()
    day01B()