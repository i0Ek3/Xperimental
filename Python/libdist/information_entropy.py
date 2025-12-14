#!/usr/bin/env python
# coding=utf-8

import numpy as np

def information_entropy(x, p):
    ret = [0] * (len(x)+1)
    for i in range(len(x)):
        ret[i] = -p[i] * np.log2(p[i])
        result = np.sum(ret[i])
    return result
