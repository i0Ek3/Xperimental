#!/usr/bin/env python
# coding=utf-8

import numpy as np

import utils as u

def zero_one_loss(y, y_pred):
    if u.validate(y, y_pred):
        return 0
    else:
        return 1

