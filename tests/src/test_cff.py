def compute_sum(n):
    total = 0
    for i in range(n):
        total += i
    return total

def main():
    print("Starting computation")
    res = compute_sum(10)
    print("Result:", res)
    if res == 45:
        print("Success")
    else:
        print("Failure")

if __name__ == '__main__':
    main()
