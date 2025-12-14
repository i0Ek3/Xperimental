#!/usr/bin/env python
# coding=utf-8

import numpy as np

def hamming(x, y):
    nsum = np.sum(x != y)
    dist = nsum / len(x)
    return dist
