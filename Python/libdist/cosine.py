#!/usr/bin/env python
# coding=utf-8

import numpy as np

def cosine(a, b):
    mol = 0.0
    der_a, der_b = 0, 0
    for i in range(len(a)):
        mol += a[i] * b[i]
        der_a += a[i] ** 2
        der_b += b[i] ** 2

    return mol / (np.sqrt(der_a) * np.sqrt(der_b))
