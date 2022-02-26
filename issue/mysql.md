# Record MySQL Issue

## MySQL Container 無法正常顯示

### 原因

character 預設是 `latin1` 。  

```
mysql> show variables like '%char%';
+--------------------------+----------------------------+
| Variable_name            | Value                      |
+--------------------------+----------------------------+
| character_set_client     | latin1                     |
| character_set_connection | latin1                     |
| character_set_database   | utf8mb4                    |
| character_set_filesystem | binary                     |
| character_set_results    | latin1                     |
| character_set_server     | utf8mb4                    |
| character_set_system     | utf8                       |
| character_sets_dir       | /usr/share/mysql/charsets/ |
+--------------------------+----------------------------+
8 rows in set (0.07 sec)
```  

### 解法

在 container 查看支援的編碼格式。  
```
root@d95f5c1497ff:/# locale -a
C
C.UTF-8
POSIX
```  

在 container 新增環境變數。  
```
export LANG=[locale -a 其中一種編碼格式]
```  

或在 docker-compose 新增變數。  
```
environment:
  - LANG=[locale -a 其中一種編碼格式]
```  
