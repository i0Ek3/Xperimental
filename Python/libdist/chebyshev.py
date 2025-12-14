#!/usr/bin/env python
# coding=utf-8

import numpy as np

def chebyshev(x, y):
    diff = np.abs(x-y)
    dist = np.max(diff)
    return dist
