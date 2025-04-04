from utils import *

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

if __name__ == '__main__':
    day01A()