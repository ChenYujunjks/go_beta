PostgreSQL 的 `psql` 和 MySQL 的 CLI 在操作方式上有一些相似之处，但也有显著的不同。以下是一些关键区别和常见命令的对比：

### 1. 基本命令对比

以下是一些 PostgreSQL (`psql`) 和 MySQL 的常用命令对比：

| **操作**             | **PostgreSQL (psql)**        | **MySQL**              |
| -------------------- | ---------------------------- | ---------------------- |
| 列出所有数据库       | `\l`                         | `SHOW DATABASES;`      |
| 切换数据库           | `\c database_name`           | `USE database_name;`   |
| 列出所有表           | `\dt`                        | `SHOW TABLES;`         |
| 查看表结构           | `\d table_name`              | `DESCRIBE table_name;` |
| 退出 CLI             | `\q`                         | `exit;`                |
| 显示当前连接的用户   | `SELECT current_user;`       | `SELECT USER();`       |
| 显示当前使用的数据库 | `SELECT current_database();` | `SELECT DATABASE();`   |
| 执行 SQL 脚本        | `\i filename.sql`            | `source filename.sql;` |

### 2. 特殊命令

- **PostgreSQL (psql)**: 大多数特殊命令以反斜杠 (`\`) 开头，例如 `\l` 列出数据库，`\dt` 列出表等。这些命令是 `psql` 专用的，并且不需要以分号结尾。
- **MySQL**: MySQL CLI 的命令通常都是 SQL 标准命令，需要以分号 (`;`) 结尾。MySQL 也有一些独特的命令，比如 `SHOW` 系列命令。

### 3. SQL 语法的区别

- 虽然两者的 SQL 语法在基础部分是相似的，但在高级功能上可能会有不同。例如，PostgreSQL 支持一些高级的 SQL 特性（如窗口函数、CTE、丰富的类型系统）和扩展，而 MySQL 的 SQL 方言可能在某些方面更加简单或不同。

### 4. 结果输出

- **PostgreSQL**: 你可以使用 `\x` 命令来切换结果的显示格式，通常用于更好地查看宽表的数据。
- **MySQL**: MySQL 默认的结果输出比较紧凑，但可以通过 `\G` 来改变结果的显示方式。

### 5. 在线帮助

- **PostgreSQL**: 你可以在 `psql` 中输入 `\?` 来查看所有命令的帮助。
- **MySQL**: 输入 `help` 或者 `\h` 可以查看命令帮助。

### 总结

虽然 `psql` 和 MySQL CLI 都用于与数据库交互，但命令语法和使用方式上有些不同。熟悉 `psql` 中的命令后，你会发现它非常强大，特别是在处理复杂查询和数据库管理任务时。

如果你需要帮助适应 `psql` 的命令或想要了解更多 PostgreSQL 的特性，可以随时问我。
