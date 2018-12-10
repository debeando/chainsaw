# MySQLDataChunk
MySQL Tool to modify big lot of data on table without blocking

## Install

## Usage

We have two cases to modify data:

### Case 1:


```
./mysql_data_chunk --host=127.0.0.1 \
                   --user=root \
                   --password=admin \
                   --schema=demo \
                   --table=foo \
                   --key=id \
                   --template=examples/template_1.tsql \
                   --chunk=100 \
                   --sleep=0.2
```

### Case 2:

```
./mysql_data_chunk --host=127.0.0.1 \
                   --user=root \
                   --password=admin \
                   --schema=demo \
                   --table=foo \
                   --fileids=examples/ids.txt \
                   --template=examples/template_2.tsql \
                   --chunk=100 \
                   --sleep=0.2
```
