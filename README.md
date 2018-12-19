# silk flows
https://tools.netsa.cert.org/silk/docs.html

## Description
This package makes it easy to read common silk files without using C Go.

## What is silk
Source https://tools.netsa.cert.org/silk/faq.html#what-silk 
>SiLK is a suite of network traffic collection and analysis tools developed and maintained by the CERT Network Situational >Awareness Team (CERT NetSA) at Carnegie Mellon University to facilitate security analysis of large networks. The SiLK tool >suite supports the efficient collection, storage, and analysis of network flow data, enabling network security analysts to >rapidly query large historical traffic data sets.

## Supported File Formats
| Record Size   | Record Version           | Compression   | Supported     |
| ------------- | -------------            | ------------- | ------------- |
| 88            | RWIPV6ROUTING VERSION 1  | None (0)      | YES           |
| 88            | RWIPV6ROUTING VERSION 1  | Zlib (1)      | YES           |
| 88            | RWIPV6ROUTING VERSION 1  | Lzo (2)       | YES           |
| 88            | RWIPV6ROUTING VERSION 1  | Snappy (3)    | YES           |
| 68            | RWIPV6 VERSION 1         | None (0)      | YES           |
| 68            | RWIPV6 VERSION 1         | Zlib (1)      | YES           |
| 68            | RWIPV6 VERSION 1         | Lzo (2)       | YES           |
| 68            | RWIPV6 VERSION 1         | Snappy (3)    | YES           |
| 56            | RWIPV6 VERSION 2         | None (0)      | YES           |
| 56            | RWIPV6 VERSION 2         | Zlib (1)      | YES           |
| 56            | RWIPV6 VERSION 2         | Lzo (2)       | YES           |
| 56            | RWIPV6 VERSION 2         | Snappy (3)    | YES           |
