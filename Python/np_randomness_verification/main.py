#!/usr/bin/env python3
# coding=utf-8

import sys

from numpy import random as rd
import numpy as np

import rich as r


def verficate(level=10, times=100):
    cnt = [0] * level

    n = 0
    while n < times:
        num = rd.randint(level)
        cnt[num] += 1
        n += 1

    for i in range(level):
        cnt[i] = np.divide(cnt[i], times)
        
        ret = "P({0}) = {1}\n"
        print(ret.format(i, cnt[i]))
    
    return cnt

def output():
    p = verficate()
    for i in range(len(p)):
        r.print(p[i])

def main(argv):
    if len(argv) == 1:
        verficate()
    elif len(argv) == 2:
        verficate(int(argv[1]))
    else:
        verficate(int(argv[1]), int(argv[2]))

if __name__ == "__main__":
    main(sys.argv)
