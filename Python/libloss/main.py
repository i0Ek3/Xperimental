#!/usr/bin/env python
# coding=utf-8

import numpy as np
from numpy import random as rd

import zero_one_loss as zo
import hinge_loss as hl


def main():
    y = rd.random(10)
    y_pred = rd.random(10)
    
    zo.zero_one_loss(y, y_pred)
    

if __name__ == '__main__':
    main()
