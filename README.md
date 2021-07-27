# Optimize "Naming in the USA Database" with Trie

## About
  Inspired by Hussein Nasser's video, I created my first mini-project in Go to optimzie regex-like query. 
  
  I'm using package [go-pq](https://github.com/go-pg/pg) to create table and query on Postgresql.
  
## Performance
  With the use of Trie, I achieved 4X performance of 2.5 ms with Trie:
 
![Query with trie](./query_wtrie.png)

  Compares to 10 ms query from database.
  
![Query with database](./query_pg.png)



  
  
## Misc
Checkout the video [Database Indexing Explained (with PostgreSQL)](https://www.youtube.com/watch?v=-qNSXK7s7_w) - 

    and his channel [Hussein Nasser](https://www.youtube.com/c/HusseinNasser-software-engineering)

Checkout project's diary at [doc.txt](./doc.txt)
