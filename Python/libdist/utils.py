#!/usr/bin/env python
# coding=utf-8

def p_msg(p):
    if p == 0:
        print("P cannot be zero!")
    elif p == 1:
        print("You'll going with Manhatten distance.")
    elif p == 2:
        print("You'll going with Euclidean distance.")
    else:
        print("You'll going with Minkowski(L_p) distance.")

# TODO: def show(x)
