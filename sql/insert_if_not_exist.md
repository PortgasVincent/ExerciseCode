https://stackoverflow.com/questions/3164505/mysql-insert-record-if-not-exists-in-table
https://stackoverflow.com/questions/1361340/how-can-i-do-insert-if-not-exists-in-mysql

// ignore 忽视出现的错误
INSERT IGNORE INTO `mytable`
SET `field0` = '2',
`field1` = 12345,
`field2` = 12678;

INSERT INTO Devices(unique_id, time) 
VALUES('$device_id', '$current_time') 
ON DUPLICATE KEY UPDATE time = '$current_time';

// 如果发现要插入的记录已经存在，则该语句将先删除该记录，再插入新的记录。这种方法的缺点是如果发生重复键错误，旧记录会被完全删除
REPLACE INTO `table` VALUES (5, 'John', 'Doe', SHA1('password'));


INSERT INTO table_listnames (name, address, tele)
SELECT * FROM (SELECT 'Rupert', 'Somewhere', '022') AS tmp
WHERE NOT EXISTS (
    SELECT name FROM table_listnames WHERE name = 'Rupert'
) LIMIT 1;


delete / update if exist

delete from table where exist (
  select...
)

update tableName set colName = 'a' where exists {
  select...
}

select * from table as a join (select * from table where fan_id = 1) as b on a.fan_id = 2 and a.follower_id = b.follower_id