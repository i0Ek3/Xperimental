#!/usr/bin/env python
# coding=utf-8

import numpy as np

import utils as u

def L_p(xi, xj, p):
    u.p_msg(p)
    
    diff = np.abs(xi-xj)
    nsum = np.sum(np.power(diff, p)) 
    if p != 0:
        dist = np.power(nsum, np.divide(1, p))
        return dist
