#!/usr/bin/env python
# coding=utf-8

import numpy as np

def hinge_loss(t, y):
    # t = 1 or -1, y is between (-1, 1)
    # Loss = np.max(0, 1 - t * y)
    return np.max(0, 1 - np.multiply(t, y))
