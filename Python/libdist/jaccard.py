#!/usr/bin/env python
# coding=utf-8

# TODO: class Jaccard
    
def jaccard_distance(a, b):
    set_a = set(a)
    set_b = set(b)
    dis = float(len((set_a | set_b) - (set_a & set_b))) / len(set_a | set_b)
    return dis

def jaccard_similarity_coefficient(a, b):
    set_a = set(a)
    set_b = set(b)
    co = float(len(set_a & set_b)) / len(set_a | set_b)
    return co
