# Optimize "Naming in the USA Database" with Trie

## About
  Inspired by Hussein Nasser's video, I created my first mini-project in Go to optimzie regex-like query. 
  
  I'm using package ![go-pq](https://github.com/go-pg/pg) to create and query Postgresql.
  
  With the use of Trie, I achieved 4X performance of 2.5 ms compares to 10 ms.
  
  Checkout project's diary at [doc.txt](./doc.txt)
  
![Query with trie](./query_wtrie.png)

![Query with database](./query_pg.png)

## Inspired by 
![Database Indexing Explained (with PostgreSQL)](https://www.youtube.com/watch?v=-qNSXK7s7_w)
