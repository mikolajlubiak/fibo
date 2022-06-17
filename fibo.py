#!/usr/bin/env python

from sys import argv

def fibo(x, y, z):
    fibo = [y, z]
    print(f"1: {fibo[0]}\n2: {fibo[1]}")
    for i in range(2, x):
        fibo.append(fibo[i-1] + fibo[i-2])
        print(f"{i+1}: {fibo[i]}")

if __name__ == "__main__":
    try:
        fibo(int(argv[3]), int(argv[1]), int(argv[2]))
    except:
        print("You have to provide 3 arguments - first number, secound number, how many numbers to calculate.\nExample:\t./fibo 0 1 10\n")
        yield ValueError("Not enough arguments!")