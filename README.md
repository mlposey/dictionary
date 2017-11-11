# dictionary
[![Build Status](https://travis-ci.org/mlposey/dictionary.svg?branch=master)](https://travis-ci.org/mlposey/dictionary)  
A bucketized cuckoo table with 3-independent string hashing

Wikipedia offers a [detailed look](https://en.wikipedia.org/wiki/Cuckoo_hashing) at the benefits of
cuckoo hashing, but in short, this library offers three important things:
* Constant-time retrieval
* A [3-independent](https://en.wikipedia.org/wiki/Tabulation_hashing) set of hash functions for strings
* A high load factor of 0.976, as per [the paper](http://ieeexplore.ieee.org/abstract/document/4221787/) by Kenneth Ross