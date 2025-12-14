#!/usr/bin/env python
# coding=utf-8

import numpy as np

def validate(x, y):
    if np.shape(x) != np.shape(y):
        return False
    else:
        if x.any() == y.any():
            print("----")
            return True
        else:
            print("++++")
            return False
