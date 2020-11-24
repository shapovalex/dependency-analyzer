
![Build status](https://travis-ci.com/shapovalex/dependency-analyzer.svg?branch=develop "Build status")

Parameters

-d - Dependency manager

-f - Input file/files, comma separated

-r - Output file. result.txt by default

-o - Operation to perform

Supported combinations of -d -o flags:

pypi compare - ./main -d pip -f req1.txt,req2.txt -o compare -r diff.txt
pypi compare - ./main -d pip -f req1.txt -o license
 
