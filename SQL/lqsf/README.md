# lqsf

> Please read it in reverse.

Use SQL easily after learning it.

## Install & Run SQL

- Download MySQL and install it, download link: [https://dev.mysql.com/downloads/mysql/](https://dev.mysql.com/downloads/mysql/)
- Or use command: `brew install mysql` to install mysql on macOS
- And then run command: `brew services start mysql` and `sudo mysql -u root -p` to start mysql
> when terminal show you message to enter the password, just hit enter to pass it.

## Usage

```
$ mysql -uroot -p < init-test-data.sql

mysql> show databases;
mysql> show tables from test;

-- use following SQL commands to get results
```

## YSK

- Do not use business related fields as primary key
    - Use BIGINT or GUID type
- Use `ADD INDEX index_name (name);` to create index for colunm name
    - Use `ADD UNIQUE INDEX uni_name (name);` or `ADD CONSTRAINT uni_name UNIQUE (name);` 
- Query tips
    - `SELECT * FROM table_name;`
    - `SELECT * FROM table_name WHERE condition;`
    - `SELECT a, b, c FROM table_name WHERE condition;`
    - `SELECT a, b, c FROM table_name WHERE condition ORDER BY field DESC;`
    - `SELECT a, b, c FROM table_name WHERE condition ORDER BY field DESC LIMIT 3 OFFSET 0;` -- query first page which contain 3 items
    - `SELECT count(*) alias_name FROM table_name;`
    - `SELECT * FROM table_name GROUP BY field;`
    - `SELECT * FROM table1_name, table2_name;` -- be careful to use multi table query
    - `SELECT field FROM table1 INNER JOIN table2 ON condition;` -- INNER JOIN return data which exist in two tables, RIGHT OUTER JOIN and LEFT OUTER JOIN just returns coresponding data which exist in RIGHT or LEFT table, FULL OUTER JOIN return all data from tables
- Alter tips
    - `INSERT [IGNORE] INTO table_name (field1, ...) VALUES (value1, ...);`
    - `UPDATE table_name SET field1=value1, ... [WHERE condition];` -- [] means this is optional item 
    - `DELETE FROM table_name [WHERE condition];`
    - `CREATE DATABASE/TABLE db_name/table_name;`
    - `DROP DATABASE/TABLE db_name/table_name;`
    - `USE db_name;`
- Transaction
    - `BEGIN; SQL commands; COMMIT/ROLLBACK;`

## Credit

[https://www.liaoxuefeng.com](https://www.liaoxuefeng.com).
