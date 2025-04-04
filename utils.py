def getPath(day: int, full: bool) -> str:
    suffix = 'test' if full else 'sample'
    return 'data/%.2d_%s.txt' % (day, suffix)

def readFile(path: str) -> list[str]:
    f = open(path, 'r')
    lines = [x.strip() for x in f.readlines()]
    f.close()
    return lines