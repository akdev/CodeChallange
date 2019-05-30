# Highest Scores

## Problem Text:
Given a data file containing scored records, in your favorite programming language,
write a program to output the N highest record IDs by score, highest score first. The
output should be correctly formatted JSON.
The input data file has a record per line. Each line has the following structure:

`<score>:<jsondictionary>`

The JSON can be any kind of well structured JSON doc, with the obvious exception of
no line breaks. Besides needing to be well­formed without line breaks, the only
requirement is that the dictionary contains a "id" key which uniquely defines that data
row (you can assume uniqueness of ids in the data set). If the JSON part of the record
does not have an "id" key it should be considered an input formatting error and handled
as described below.
```
8795136:{"id":"d2e257c282b54347ac14b2d8d0da814c","x":"foo","y":7,"z":11,
"payload":"somelargeamountofdata"}
5317020:{"id":"619236365add4a0ca6e501fc62cfbaba","type":"purple","payload":"some
smallamountofdata"}
...
```

An example program execution is:
```
$./highest.pyscore_recs.data5
[
{
"score":16774838,
"id":"9ab7247c02044c65936a467016fff6b6"
},
{
"score":16763774,
"id":"c51a310f80604ef68a4cb2b83bffcb7e"
},
{
"score":16761021,
"id":"c1dbd109336242e0a64527ba8cffc0bd"
},
{
"score":16755441,
"id":"57b9ea55db954cbc8f452b34a2ffaaf1"
},
{
"score":16753041,
"id":"e8cafaf8cf2b41639422781fbdffa191"
}
]
```
Upon successful running, the program should exit with exit code 0. Input formatting
errors should exit with code 2 while file not found should exit with code 1. Empty lines
should be ignored rather than treated as improperly formatted.
