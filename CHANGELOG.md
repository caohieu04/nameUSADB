v.1.2.0

==========================================
v.1.1.0

```diff 
+ Added ```:
  Added `Siege` class, now can send and receive bulk of the requests (3.000 queries per seconds)
  Added `Route` class, a simple APIs created by `mux` package
``` diff 
! Altered ```:
  Moved all excepts main.go to package utils
  Moved ./name.csv to ./data/name.csv

==========================================
v.1.0.0 

```diff 
+ Added ```:
  Created database by bulk importing from name.csv and created indexes on `name` columns
  Built `Trie` for fast retrieving
