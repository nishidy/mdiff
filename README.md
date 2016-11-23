# mdiff
Diff command for multiple lines regardless of its appearance order in a file.

# How is the way different from historical "sort" and "diff"?

```
$ wc -l orig_data 
 1000000 orig_data

$ head -n 10 orig_data 
LRFKQYUQFJ
KXYQVNRTYS
FRZRMZLYGF
VEULQFPDBH
LQDQRRCRWD
NXEUOQQEKL
AITGDPHCSP
IJTHBSFYFV
LADZPBFUDK
KLRWQAOZMI

$ cp orig_data copy_data

$ sed -n '10000p' copy_data 
WKOKGCNPVA

$ sed -i '' '10000d' copy_data 

$ wc -l copy_data 
  999999 copy_data

$ /usr/bin/time -lp sort orig_data > sorted_orig_data
real        14.09
user        14.02
sys          0.05
  59740160  maximum resident set size
         0  average shared memory size
         0  average unshared data size
         0  average unshared stack size
     14593  page reclaims
         1  page faults
         0  swaps
        13  block input operations
         2  block output operations
         0  messages sent
         0  messages received
         0  signals received
        15  voluntary context switches
       387  involuntary context switches

$ /usr/bin/time -lp sort copy_data > sorted_copy_data
real        14.22
user        14.15
sys          0.06
  59736064  maximum resident set size
         0  average shared memory size
         0  average unshared data size
         0  average unshared stack size
     14593  page reclaims
         0  page faults
         0  swaps
         0  block input operations
         3  block output operations
         0  messages sent
         0  messages received
         0  signals received
         0  voluntary context switches
      1738  involuntary context switches

$ /usr/bin/time -lp diff sorted_orig_data sorted_copy_data 
861444d861443
< WKOKGCNPVA
real         0.07
user         0.03
sys          0.02
  34750464  maximum resident set size
         0  average shared memory size
         0  average unshared data size
         0  average unshared stack size
      8654  page reclaims
         1  page faults
         0  swaps
         0  block input operations
         1  block output operations
         0  messages sent
         0  messages received
         0  signals received
         2  voluntary context switches
        12  involuntary context switches

```

```
$ /usr/bin/time -lp ./mdiff orig_data copy_data 
orig_data
L10000: WKOKGCNPVA

copy_data

real         1.62
user         1.66
sys          0.19
 489885696  maximum resident set size
         0  average shared memory size
         0  average unshared data size
         0  average unshared stack size
    119594  page reclaims
         0  page faults
         0  swaps
         0  block input operations
         0  block output operations
         0  messages sent
         0  messages received
         0  signals received
       724  voluntary context switches
      6364  involuntary context switches

$ /usr/bin/time -lp ./mdiff orig_data sorted_copy_data 
orig_data
L10000: WKOKGCNPVA

sorted_copy_data

real         1.60
user         1.68
sys          0.22
 523038720  maximum resident set size
         0  average shared memory size
         0  average unshared data size
         0  average unshared stack size
    127693  page reclaims
         2  page faults
         0  swaps
         0  block input operations
         1  block output operations
         0  messages sent
         0  messages received
         0  signals received
       673  voluntary context switches
      6962  involuntary context switches

```

