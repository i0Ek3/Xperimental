#!/usr/bin/env python
# coding=utf-8

import numpy as np

def mahalanobis(x, y):
    dist = []
    X = np.vstack([x, y])
    T = X.T
    S = np.cov(X)
    I = np.linalg.inv(S)
    n = T.shape[0]

    for i in range(0, n):
        for j in range(i+1, n):
            delta = T[i] - T[j]
            d = np.sqrt(np.dot(np.dot(delta, I), delta.T))
            dist.append(d)

    return dist
