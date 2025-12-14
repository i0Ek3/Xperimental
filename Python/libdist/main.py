#!/usr/bin/env python
# coding=utf-8

from numpy import random as rd

import chebyshev as ch
import hamming as hm
import minkowski as mk
import mahalanobis as ma
import information_entropy as ie
import jaccard as jc
import cosine as cs
import utils as u

def main():
    print(ch.chebyshev(rd.random(10), rd.random(10)))
    
    print(hm.hamming(rd.random(10), rd.random(10)))
    
    print(mk.L_p(rd.random(10), rd.random(10), 0))
    print(mk.L_p(rd.random(10), rd.random(10), 1))
    print(mk.L_p(rd.random(10), rd.random(10), 2))
    
    print(ma.mahalanobis(rd.random(10), rd.random(10)))

    print(ie.information_entropy((1, 2, 3, 4), (0.1, 0.2, 0.5, 0.7)))

    print(jc.jaccard_distance((1, 3, 4), (2, 4, 6)))
    print(jc.jaccard_similarity_coefficient((1, 3, 4), (1, 4, 6)))

    print(cs.cosine((1, 3, 4, 5), (2, 3, 4, 5)))

if __name__ == '__main__':
    main()
