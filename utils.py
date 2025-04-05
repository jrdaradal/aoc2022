def getPath(day: int, full: bool) -> str:
    suffix = 'test' if full else 'sample'
    return 'data/%.2d_%s.txt' % (day, suffix)

def readFile(path: str, strip: bool = True) -> list[str]:
    f = open(path, 'r')
    if strip:
        lines = [x.strip() for x in f.readlines()]
    else:
        lines = [x for x in f.readlines()]
    f.close()
    return lines