import sys

def main():
    lines = sys.stdin.read().split()
    if not lines:
        return
    n = int(lines[0])
    for s in lines[1:n+1]:
        if len(s) > 10:
            print(f"{s[0]}{len(s)-2}{s[-1]}")
        else:
            print(s)

if __name__ == '__main__':
    main()
