
# mysql <a id="table-doc-top"></a>

<div style="text-align: right">
mysql 
</div>




* [columns_priv](#table-columns_priv) Column privileges
* [db](#table-db) Database privileges
* [engine_cost](#table-engine_cost) 
* [event](#table-event) Events
* [func](#table-func) User defined functions
* [general_log](#table-general_log) General log
* [gtid_executed](#table-gtid_executed) 
* [help_category](#table-help_category) help categories
* [help_keyword](#table-help_keyword) help keywords
* [help_relation](#table-help_relation) keyword-topic relation
* [help_topic](#table-help_topic) help topics
* [innodb_index_stats](#table-innodb_index_stats) 
* [innodb_table_stats](#table-innodb_table_stats) 
* [ndb_binlog_index](#table-ndb_binlog_index) 
* [plugin](#table-plugin) MySQL plugins
* [proc](#table-proc) Stored Procedures
* [procs_priv](#table-procs_priv) Procedure privileges
* [proxies_priv](#table-proxies_priv) User proxy privileges
* [server_cost](#table-server_cost) 
* [servers](#table-servers) MySQL Foreign Servers table
* [slave_master_info](#table-slave_master_info) Master Information
* [slave_relay_log_info](#table-slave_relay_log_info) Relay Log Information
* [slave_worker_info](#table-slave_worker_info) Worker Information
* [slow_log](#table-slow_log) Slow log
* [tables_priv](#table-tables_priv) Table privileges
* [time_zone](#table-time_zone) Time zones
* [time_zone_leap_second](#table-time_zone_leap_second) Leap seconds information for time zones
* [time_zone_name](#table-time_zone_name) Time zone names
* [time_zone_transition](#table-time_zone_transition) Time zone transitions
* [time_zone_transition_type](#table-time_zone_transition_type) Time zone transition types
* [user](#table-user) Users and global privileges



## columns_priv <a id="table-columns_priv"></a>

BASE TABLE

> Column privileges 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Host | char(60) |  | NO |  | 
2 | Db | char(64) |  | NO |  | 
3 | User | char(32) |  | NO |  | 
4 | Table_name | char(64) |  | NO |  | 
5 | Column_name | char(64) |  | NO |  | 
6 | Timestamp | timestamp | CURRENT_TIMESTAMP | NO |  | 
7 | Column_priv | set(&#39;Select&#39;,&#39;Insert&#39;,&#39;Update&#39;,&#39;References&#39;) |  | NO |  | 

```sql
CREATE TABLE `columns_priv` (
  `Host` char(60) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Db` char(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `User` char(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Table_name` char(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Column_name` char(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `Column_priv` set('Select','Insert','Update','References') CHARACTER SET utf8 NOT NULL DEFAULT '',
  PRIMARY KEY (`Host`,`Db`,`User`,`Table_name`,`Column_name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='Column privileges'
```

[▲ top](#table-doc-top)

## db <a id="table-db"></a>

BASE TABLE

> Database privileges 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Host | char(60) |  | NO |  | 
2 | Db | char(64) |  | NO |  | 
3 | User | char(32) |  | NO |  | 
4 | Select_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
5 | Insert_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
6 | Update_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
7 | Delete_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
8 | Create_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
9 | Drop_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
10 | Grant_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
11 | References_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
12 | Index_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
13 | Alter_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
14 | Create_tmp_table_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
15 | Lock_tables_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
16 | Create_view_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
17 | Show_view_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
18 | Create_routine_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
19 | Alter_routine_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
20 | Execute_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
21 | Event_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
22 | Trigger_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 

```sql
CREATE TABLE `db` (
  `Host` char(60) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Db` char(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `User` char(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Select_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Insert_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Update_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Delete_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Drop_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Grant_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `References_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Index_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Alter_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_tmp_table_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Lock_tables_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_view_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Show_view_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_routine_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Alter_routine_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Execute_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Event_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Trigger_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  PRIMARY KEY (`Host`,`Db`,`User`),
  KEY `User` (`User`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='Database privileges'
```

[▲ top](#table-doc-top)

## engine_cost <a id="table-engine_cost"></a>

BASE TABLE

>  



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | engine_name | varchar(64) | null | NO |  | 
2 | device_type | int(11) | null | NO |  | 
3 | cost_name | varchar(64) | null | NO |  | 
4 | cost_value | float | null | YES |  | 
5 | last_update | timestamp | CURRENT_TIMESTAMP | NO |  | 
6 | comment | varchar(1024) | null | YES |  | 

```sql
CREATE TABLE `engine_cost` (
  `engine_name` varchar(64) NOT NULL,
  `device_type` int(11) NOT NULL,
  `cost_name` varchar(64) NOT NULL,
  `cost_value` float DEFAULT NULL,
  `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comment` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`cost_name`,`engine_name`,`device_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0
```

[▲ top](#table-doc-top)

## event <a id="table-event"></a>

BASE TABLE

> Events 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | db | char(64) |  | NO |  | 
2 | name | char(64) |  | NO |  | 
3 | body | longblob | null | NO |  | 
4 | definer | char(93) |  | NO |  | 
5 | execute_at | datetime | null | YES |  | 
6 | interval_value | int(11) | null | YES |  | 
7 | interval_field | enum(&#39;YEAR&#39;,&#39;QUARTER&#39;,&#39;MONTH&#39;,&#39;DAY&#39;,&#39;HOUR&#39;,&#39;MINUTE&#39;,&#39;WEEK&#39;,&#39;SECOND&#39;,&#39;MICROSECOND&#39;,&#39;YEAR_MONTH&#39;,&#39;DAY_HOUR&#39;,&#39;DAY_MINUTE&#39;,&#39;DAY_SECOND&#39;,&#39;HOUR_MINUTE&#39;,&#39;HOUR_SECOND&#39;,&#39;MINUTE_SECOND&#39;,&#39;DAY_MICROSECOND&#39;,&#39;HOUR_MICROSECOND&#39;,&#39;MINUTE_MICROSECOND&#39;,&#39;SECOND_MICROSECOND&#39;) | null | YES |  | 
8 | created | timestamp | CURRENT_TIMESTAMP | NO |  | 
9 | modified | timestamp | 0000-00-00 00:00:00 | NO |  | 
10 | last_executed | datetime | null | YES |  | 
11 | starts | datetime | null | YES |  | 
12 | ends | datetime | null | YES |  | 
13 | status | enum(&#39;ENABLED&#39;,&#39;DISABLED&#39;,&#39;SLAVESIDE_DISABLED&#39;) | ENABLED | NO |  | 
14 | on_completion | enum(&#39;DROP&#39;,&#39;PRESERVE&#39;) | DROP | NO |  | 
15 | sql_mode | set(&#39;REAL_AS_FLOAT&#39;,&#39;PIPES_AS_CONCAT&#39;,&#39;ANSI_QUOTES&#39;,&#39;IGNORE_SPACE&#39;,&#39;NOT_USED&#39;,&#39;ONLY_FULL_GROUP_BY&#39;,&#39;NO_UNSIGNED_SUBTRACTION&#39;,&#39;NO_DIR_IN_CREATE&#39;,&#39;POSTGRESQL&#39;,&#39;ORACLE&#39;,&#39;MSSQL&#39;,&#39;DB2&#39;,&#39;MAXDB&#39;,&#39;NO_KEY_OPTIONS&#39;,&#39;NO_TABLE_OPTIONS&#39;,&#39;NO_FIELD_OPTIONS&#39;,&#39;MYSQL323&#39;,&#39;MYSQL40&#39;,&#39;ANSI&#39;,&#39;NO_AUTO_VALUE_ON_ZERO&#39;,&#39;NO_BACKSLASH_ESCAPES&#39;,&#39;STRICT_TRANS_TABLES&#39;,&#39;STRICT_ALL_TABLES&#39;,&#39;NO_ZERO_IN_DATE&#39;,&#39;NO_ZERO_DATE&#39;,&#39;INVALID_DATES&#39;,&#39;ERROR_FOR_DIVISION_BY_ZERO&#39;,&#39;TRADITIONAL&#39;,&#39;NO_AUTO_CREATE_USER&#39;,&#39;HIGH_NOT_PRECEDENCE&#39;,&#39;NO_ENGINE_SUBSTITUTION&#39;,&#39;PAD_CHAR_TO_FULL_LENGTH&#39;) |  | NO |  | 
16 | comment | char(64) |  | NO |  | 
17 | originator | int(10) unsigned | null | NO |  | 
18 | time_zone | char(64) | SYSTEM | NO |  | 
19 | character_set_client | char(32) | null | YES |  | 
20 | collation_connection | char(32) | null | YES |  | 
21 | db_collation | char(32) | null | YES |  | 
22 | body_utf8 | longblob | null | YES |  | 

```sql
CREATE TABLE `event` (
  `db` char(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `name` char(64) NOT NULL DEFAULT '',
  `body` longblob NOT NULL,
  `definer` char(93) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `execute_at` datetime DEFAULT NULL,
  `interval_value` int(11) DEFAULT NULL,
  `interval_field` enum('YEAR','QUARTER','MONTH','DAY','HOUR','MINUTE','WEEK','SECOND','MICROSECOND','YEAR_MONTH','DAY_HOUR','DAY_MINUTE','DAY_SECOND','HOUR_MINUTE','HOUR_SECOND','MINUTE_SECOND','DAY_MICROSECOND','HOUR_MICROSECOND','MINUTE_MICROSECOND','SECOND_MICROSECOND') DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `last_executed` datetime DEFAULT NULL,
  `starts` datetime DEFAULT NULL,
  `ends` datetime DEFAULT NULL,
  `status` enum('ENABLED','DISABLED','SLAVESIDE_DISABLED') NOT NULL DEFAULT 'ENABLED',
  `on_completion` enum('DROP','PRESERVE') NOT NULL DEFAULT 'DROP',
  `sql_mode` set('REAL_AS_FLOAT','PIPES_AS_CONCAT','ANSI_QUOTES','IGNORE_SPACE','NOT_USED','ONLY_FULL_GROUP_BY','NO_UNSIGNED_SUBTRACTION','NO_DIR_IN_CREATE','POSTGRESQL','ORACLE','MSSQL','DB2','MAXDB','NO_KEY_OPTIONS','NO_TABLE_OPTIONS','NO_FIELD_OPTIONS','MYSQL323','MYSQL40','ANSI','NO_AUTO_VALUE_ON_ZERO','NO_BACKSLASH_ESCAPES','STRICT_TRANS_TABLES','STRICT_ALL_TABLES','NO_ZERO_IN_DATE','NO_ZERO_DATE','INVALID_DATES','ERROR_FOR_DIVISION_BY_ZERO','TRADITIONAL','NO_AUTO_CREATE_USER','HIGH_NOT_PRECEDENCE','NO_ENGINE_SUBSTITUTION','PAD_CHAR_TO_FULL_LENGTH') NOT NULL DEFAULT '',
  `comment` char(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `originator` int(10) unsigned NOT NULL,
  `time_zone` char(64) CHARACTER SET latin1 NOT NULL DEFAULT 'SYSTEM',
  `character_set_client` char(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `collation_connection` char(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `db_collation` char(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `body_utf8` longblob,
  PRIMARY KEY (`db`,`name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='Events'
```

[▲ top](#table-doc-top)

## func <a id="table-func"></a>

BASE TABLE

> User defined functions 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | name | char(64) |  | NO |  | 
2 | ret | tinyint(1) | 0 | NO |  | 
3 | dl | char(128) |  | NO |  | 
4 | type | enum(&#39;function&#39;,&#39;aggregate&#39;) | null | NO |  | 

```sql
CREATE TABLE `func` (
  `name` char(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `ret` tinyint(1) NOT NULL DEFAULT '0',
  `dl` char(128) COLLATE utf8_bin NOT NULL DEFAULT '',
  `type` enum('function','aggregate') CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='User defined functions'
```

[▲ top](#table-doc-top)

## general_log <a id="table-general_log"></a>

BASE TABLE

> General log 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | event_time | timestamp(6) | CURRENT_TIMESTAMP(6) | NO |  | 
2 | user_host | mediumtext | null | NO |  | 
3 | thread_id | bigint(21) unsigned | null | NO |  | 
4 | server_id | int(10) unsigned | null | NO |  | 
5 | command_type | varchar(64) | null | NO |  | 
6 | argument | mediumblob | null | NO |  | 

```sql
CREATE TABLE `general_log` (
  `event_time` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `user_host` mediumtext NOT NULL,
  `thread_id` bigint(21) unsigned NOT NULL,
  `server_id` int(10) unsigned NOT NULL,
  `command_type` varchar(64) NOT NULL,
  `argument` mediumblob NOT NULL
) ENGINE=CSV DEFAULT CHARSET=utf8 COMMENT='General log'
```

[▲ top](#table-doc-top)

## gtid_executed <a id="table-gtid_executed"></a>

BASE TABLE

>  



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | source_uuid | char(36) | null | NO | uuid of the source where the transaction was originally executed. | 
2 | interval_start | bigint(20) | null | NO | First number of interval. | 
3 | interval_end | bigint(20) | null | NO | Last number of interval. | 

```sql
CREATE TABLE `gtid_executed` (
  `source_uuid` char(36) NOT NULL COMMENT 'uuid of the source where the transaction was originally executed.',
  `interval_start` bigint(20) NOT NULL COMMENT 'First number of interval.',
  `interval_end` bigint(20) NOT NULL COMMENT 'Last number of interval.',
  PRIMARY KEY (`source_uuid`,`interval_start`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1
```

[▲ top](#table-doc-top)

## help_category <a id="table-help_category"></a>

BASE TABLE

> help categories 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | help_category_id | smallint(5) unsigned | null | NO |  | 
2 | name | char(64) | null | NO |  | 
3 | parent_category_id | smallint(5) unsigned | null | YES |  | 
4 | url | text | null | NO |  | 

```sql
CREATE TABLE `help_category` (
  `help_category_id` smallint(5) unsigned NOT NULL,
  `name` char(64) NOT NULL,
  `parent_category_id` smallint(5) unsigned DEFAULT NULL,
  `url` text NOT NULL,
  PRIMARY KEY (`help_category_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='help categories'
```

[▲ top](#table-doc-top)

## help_keyword <a id="table-help_keyword"></a>

BASE TABLE

> help keywords 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | help_keyword_id | int(10) unsigned | null | NO |  | 
2 | name | char(64) | null | NO |  | 

```sql
CREATE TABLE `help_keyword` (
  `help_keyword_id` int(10) unsigned NOT NULL,
  `name` char(64) NOT NULL,
  PRIMARY KEY (`help_keyword_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='help keywords'
```

[▲ top](#table-doc-top)

## help_relation <a id="table-help_relation"></a>

BASE TABLE

> keyword-topic relation 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | help_topic_id | int(10) unsigned | null | NO |  | 
2 | help_keyword_id | int(10) unsigned | null | NO |  | 

```sql
CREATE TABLE `help_relation` (
  `help_topic_id` int(10) unsigned NOT NULL,
  `help_keyword_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`help_keyword_id`,`help_topic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='keyword-topic relation'
```

[▲ top](#table-doc-top)

## help_topic <a id="table-help_topic"></a>

BASE TABLE

> help topics 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | help_topic_id | int(10) unsigned | null | NO |  | 
2 | name | char(64) | null | NO |  | 
3 | help_category_id | smallint(5) unsigned | null | NO |  | 
4 | description | text | null | NO |  | 
5 | example | text | null | NO |  | 
6 | url | text | null | NO |  | 

```sql
CREATE TABLE `help_topic` (
  `help_topic_id` int(10) unsigned NOT NULL,
  `name` char(64) NOT NULL,
  `help_category_id` smallint(5) unsigned NOT NULL,
  `description` text NOT NULL,
  `example` text NOT NULL,
  `url` text NOT NULL,
  PRIMARY KEY (`help_topic_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='help topics'
```

[▲ top](#table-doc-top)

## innodb_index_stats <a id="table-innodb_index_stats"></a>

BASE TABLE

>  



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | database_name | varchar(64) | null | NO |  | 
2 | table_name | varchar(199) | null | NO |  | 
3 | index_name | varchar(64) | null | NO |  | 
4 | last_update | timestamp | CURRENT_TIMESTAMP | NO |  | 
5 | stat_name | varchar(64) | null | NO |  | 
6 | stat_value | bigint(20) unsigned | null | NO |  | 
7 | sample_size | bigint(20) unsigned | null | YES |  | 
8 | stat_description | varchar(1024) | null | NO |  | 

```sql
CREATE TABLE `innodb_index_stats` (
  `database_name` varchar(64) COLLATE utf8_bin NOT NULL,
  `table_name` varchar(199) COLLATE utf8_bin NOT NULL,
  `index_name` varchar(64) COLLATE utf8_bin NOT NULL,
  `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `stat_name` varchar(64) COLLATE utf8_bin NOT NULL,
  `stat_value` bigint(20) unsigned NOT NULL,
  `sample_size` bigint(20) unsigned DEFAULT NULL,
  `stat_description` varchar(1024) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`database_name`,`table_name`,`index_name`,`stat_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin STATS_PERSISTENT=0
```

[▲ top](#table-doc-top)

## innodb_table_stats <a id="table-innodb_table_stats"></a>

BASE TABLE

>  



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | database_name | varchar(64) | null | NO |  | 
2 | table_name | varchar(199) | null | NO |  | 
3 | last_update | timestamp | CURRENT_TIMESTAMP | NO |  | 
4 | n_rows | bigint(20) unsigned | null | NO |  | 
5 | clustered_index_size | bigint(20) unsigned | null | NO |  | 
6 | sum_of_other_index_sizes | bigint(20) unsigned | null | NO |  | 

```sql
CREATE TABLE `innodb_table_stats` (
  `database_name` varchar(64) COLLATE utf8_bin NOT NULL,
  `table_name` varchar(199) COLLATE utf8_bin NOT NULL,
  `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `n_rows` bigint(20) unsigned NOT NULL,
  `clustered_index_size` bigint(20) unsigned NOT NULL,
  `sum_of_other_index_sizes` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`database_name`,`table_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin STATS_PERSISTENT=0
```

[▲ top](#table-doc-top)

## ndb_binlog_index <a id="table-ndb_binlog_index"></a>

BASE TABLE

>  



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Position | bigint(20) unsigned | null | NO |  | 
2 | File | varchar(255) | null | NO |  | 
3 | epoch | bigint(20) unsigned | null | NO |  | 
4 | inserts | int(10) unsigned | null | NO |  | 
5 | updates | int(10) unsigned | null | NO |  | 
6 | deletes | int(10) unsigned | null | NO |  | 
7 | schemaops | int(10) unsigned | null | NO |  | 
8 | orig_server_id | int(10) unsigned | null | NO |  | 
9 | orig_epoch | bigint(20) unsigned | null | NO |  | 
10 | gci | int(10) unsigned | null | NO |  | 
11 | next_position | bigint(20) unsigned | null | NO |  | 
12 | next_file | varchar(255) | null | NO |  | 

```sql
CREATE TABLE `ndb_binlog_index` (
  `Position` bigint(20) unsigned NOT NULL,
  `File` varchar(255) NOT NULL,
  `epoch` bigint(20) unsigned NOT NULL,
  `inserts` int(10) unsigned NOT NULL,
  `updates` int(10) unsigned NOT NULL,
  `deletes` int(10) unsigned NOT NULL,
  `schemaops` int(10) unsigned NOT NULL,
  `orig_server_id` int(10) unsigned NOT NULL,
  `orig_epoch` bigint(20) unsigned NOT NULL,
  `gci` int(10) unsigned NOT NULL,
  `next_position` bigint(20) unsigned NOT NULL,
  `next_file` varchar(255) NOT NULL,
  PRIMARY KEY (`epoch`,`orig_server_id`,`orig_epoch`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1
```

[▲ top](#table-doc-top)

## plugin <a id="table-plugin"></a>

BASE TABLE

> MySQL plugins 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | name | varchar(64) |  | NO |  | 
2 | dl | varchar(128) |  | NO |  | 

```sql
CREATE TABLE `plugin` (
  `name` varchar(64) NOT NULL DEFAULT '',
  `dl` varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='MySQL plugins'
```

[▲ top](#table-doc-top)

## proc <a id="table-proc"></a>

BASE TABLE

> Stored Procedures 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | db | char(64) |  | NO |  | 
2 | name | char(64) |  | NO |  | 
3 | type | enum(&#39;FUNCTION&#39;,&#39;PROCEDURE&#39;) | null | NO |  | 
4 | specific_name | char(64) |  | NO |  | 
5 | language | enum(&#39;SQL&#39;) | SQL | NO |  | 
6 | sql_data_access | enum(&#39;CONTAINS_SQL&#39;,&#39;NO_SQL&#39;,&#39;READS_SQL_DATA&#39;,&#39;MODIFIES_SQL_DATA&#39;) | CONTAINS_SQL | NO |  | 
7 | is_deterministic | enum(&#39;YES&#39;,&#39;NO&#39;) | NO | NO |  | 
8 | security_type | enum(&#39;INVOKER&#39;,&#39;DEFINER&#39;) | DEFINER | NO |  | 
9 | param_list | blob | null | NO |  | 
10 | returns | longblob | null | NO |  | 
11 | body | longblob | null | NO |  | 
12 | definer | char(93) |  | NO |  | 
13 | created | timestamp | CURRENT_TIMESTAMP | NO |  | 
14 | modified | timestamp | 0000-00-00 00:00:00 | NO |  | 
15 | sql_mode | set(&#39;REAL_AS_FLOAT&#39;,&#39;PIPES_AS_CONCAT&#39;,&#39;ANSI_QUOTES&#39;,&#39;IGNORE_SPACE&#39;,&#39;NOT_USED&#39;,&#39;ONLY_FULL_GROUP_BY&#39;,&#39;NO_UNSIGNED_SUBTRACTION&#39;,&#39;NO_DIR_IN_CREATE&#39;,&#39;POSTGRESQL&#39;,&#39;ORACLE&#39;,&#39;MSSQL&#39;,&#39;DB2&#39;,&#39;MAXDB&#39;,&#39;NO_KEY_OPTIONS&#39;,&#39;NO_TABLE_OPTIONS&#39;,&#39;NO_FIELD_OPTIONS&#39;,&#39;MYSQL323&#39;,&#39;MYSQL40&#39;,&#39;ANSI&#39;,&#39;NO_AUTO_VALUE_ON_ZERO&#39;,&#39;NO_BACKSLASH_ESCAPES&#39;,&#39;STRICT_TRANS_TABLES&#39;,&#39;STRICT_ALL_TABLES&#39;,&#39;NO_ZERO_IN_DATE&#39;,&#39;NO_ZERO_DATE&#39;,&#39;INVALID_DATES&#39;,&#39;ERROR_FOR_DIVISION_BY_ZERO&#39;,&#39;TRADITIONAL&#39;,&#39;NO_AUTO_CREATE_USER&#39;,&#39;HIGH_NOT_PRECEDENCE&#39;,&#39;NO_ENGINE_SUBSTITUTION&#39;,&#39;PAD_CHAR_TO_FULL_LENGTH&#39;) |  | NO |  | 
16 | comment | text | null | NO |  | 
17 | character_set_client | char(32) | null | YES |  | 
18 | collation_connection | char(32) | null | YES |  | 
19 | db_collation | char(32) | null | YES |  | 
20 | body_utf8 | longblob | null | YES |  | 

```sql
CREATE TABLE `proc` (
  `db` char(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `name` char(64) NOT NULL DEFAULT '',
  `type` enum('FUNCTION','PROCEDURE') NOT NULL,
  `specific_name` char(64) NOT NULL DEFAULT '',
  `language` enum('SQL') NOT NULL DEFAULT 'SQL',
  `sql_data_access` enum('CONTAINS_SQL','NO_SQL','READS_SQL_DATA','MODIFIES_SQL_DATA') NOT NULL DEFAULT 'CONTAINS_SQL',
  `is_deterministic` enum('YES','NO') NOT NULL DEFAULT 'NO',
  `security_type` enum('INVOKER','DEFINER') NOT NULL DEFAULT 'DEFINER',
  `param_list` blob NOT NULL,
  `returns` longblob NOT NULL,
  `body` longblob NOT NULL,
  `definer` char(93) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `modified` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `sql_mode` set('REAL_AS_FLOAT','PIPES_AS_CONCAT','ANSI_QUOTES','IGNORE_SPACE','NOT_USED','ONLY_FULL_GROUP_BY','NO_UNSIGNED_SUBTRACTION','NO_DIR_IN_CREATE','POSTGRESQL','ORACLE','MSSQL','DB2','MAXDB','NO_KEY_OPTIONS','NO_TABLE_OPTIONS','NO_FIELD_OPTIONS','MYSQL323','MYSQL40','ANSI','NO_AUTO_VALUE_ON_ZERO','NO_BACKSLASH_ESCAPES','STRICT_TRANS_TABLES','STRICT_ALL_TABLES','NO_ZERO_IN_DATE','NO_ZERO_DATE','INVALID_DATES','ERROR_FOR_DIVISION_BY_ZERO','TRADITIONAL','NO_AUTO_CREATE_USER','HIGH_NOT_PRECEDENCE','NO_ENGINE_SUBSTITUTION','PAD_CHAR_TO_FULL_LENGTH') NOT NULL DEFAULT '',
  `comment` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `character_set_client` char(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `collation_connection` char(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `db_collation` char(32) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `body_utf8` longblob,
  PRIMARY KEY (`db`,`name`,`type`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='Stored Procedures'
```

[▲ top](#table-doc-top)

## procs_priv <a id="table-procs_priv"></a>

BASE TABLE

> Procedure privileges 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Host | char(60) |  | NO |  | 
2 | Db | char(64) |  | NO |  | 
3 | User | char(32) |  | NO |  | 
4 | Routine_name | char(64) |  | NO |  | 
5 | Routine_type | enum(&#39;FUNCTION&#39;,&#39;PROCEDURE&#39;) | null | NO |  | 
6 | Grantor | char(93) |  | NO |  | 
7 | Proc_priv | set(&#39;Execute&#39;,&#39;Alter Routine&#39;,&#39;Grant&#39;) |  | NO |  | 
8 | Timestamp | timestamp | CURRENT_TIMESTAMP | NO |  | 

```sql
CREATE TABLE `procs_priv` (
  `Host` char(60) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Db` char(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `User` char(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Routine_name` char(64) CHARACTER SET utf8 NOT NULL DEFAULT '',
  `Routine_type` enum('FUNCTION','PROCEDURE') COLLATE utf8_bin NOT NULL,
  `Grantor` char(93) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Proc_priv` set('Execute','Alter Routine','Grant') CHARACTER SET utf8 NOT NULL DEFAULT '',
  `Timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`Host`,`Db`,`User`,`Routine_name`,`Routine_type`),
  KEY `Grantor` (`Grantor`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='Procedure privileges'
```

[▲ top](#table-doc-top)

## proxies_priv <a id="table-proxies_priv"></a>

BASE TABLE

> User proxy privileges 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Host | char(60) |  | NO |  | 
2 | User | char(32) |  | NO |  | 
3 | Proxied_host | char(60) |  | NO |  | 
4 | Proxied_user | char(32) |  | NO |  | 
5 | With_grant | tinyint(1) | 0 | NO |  | 
6 | Grantor | char(93) |  | NO |  | 
7 | Timestamp | timestamp | CURRENT_TIMESTAMP | NO |  | 

```sql
CREATE TABLE `proxies_priv` (
  `Host` char(60) COLLATE utf8_bin NOT NULL DEFAULT '',
  `User` char(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Proxied_host` char(60) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Proxied_user` char(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `With_grant` tinyint(1) NOT NULL DEFAULT '0',
  `Grantor` char(93) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`Host`,`User`,`Proxied_host`,`Proxied_user`),
  KEY `Grantor` (`Grantor`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='User proxy privileges'
```

[▲ top](#table-doc-top)

## server_cost <a id="table-server_cost"></a>

BASE TABLE

>  



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | cost_name | varchar(64) | null | NO |  | 
2 | cost_value | float | null | YES |  | 
3 | last_update | timestamp | CURRENT_TIMESTAMP | NO |  | 
4 | comment | varchar(1024) | null | YES |  | 

```sql
CREATE TABLE `server_cost` (
  `cost_name` varchar(64) NOT NULL,
  `cost_value` float DEFAULT NULL,
  `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `comment` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`cost_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0
```

[▲ top](#table-doc-top)

## servers <a id="table-servers"></a>

BASE TABLE

> MySQL Foreign Servers table 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Server_name | char(64) |  | NO |  | 
2 | Host | char(64) |  | NO |  | 
3 | Db | char(64) |  | NO |  | 
4 | Username | char(64) |  | NO |  | 
5 | Password | char(64) |  | NO |  | 
6 | Port | int(4) | 0 | NO |  | 
7 | Socket | char(64) |  | NO |  | 
8 | Wrapper | char(64) |  | NO |  | 
9 | Owner | char(64) |  | NO |  | 

```sql
CREATE TABLE `servers` (
  `Server_name` char(64) NOT NULL DEFAULT '',
  `Host` char(64) NOT NULL DEFAULT '',
  `Db` char(64) NOT NULL DEFAULT '',
  `Username` char(64) NOT NULL DEFAULT '',
  `Password` char(64) NOT NULL DEFAULT '',
  `Port` int(4) NOT NULL DEFAULT '0',
  `Socket` char(64) NOT NULL DEFAULT '',
  `Wrapper` char(64) NOT NULL DEFAULT '',
  `Owner` char(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`Server_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='MySQL Foreign Servers table'
```

[▲ top](#table-doc-top)

## slave_master_info <a id="table-slave_master_info"></a>

BASE TABLE

> Master Information 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Number_of_lines | int(10) unsigned | null | NO | Number of lines in the file. | 
2 | Master_log_name | text | null | NO | The name of the master binary log currently being read from the master. | 
3 | Master_log_pos | bigint(20) unsigned | null | NO | The master log position of the last read event. | 
4 | Host | char(64) | null | YES | The host name of the master. | 
5 | User_name | text | null | YES | The user name used to connect to the master. | 
6 | User_password | text | null | YES | The password used to connect to the master. | 
7 | Port | int(10) unsigned | null | NO | The network port used to connect to the master. | 
8 | Connect_retry | int(10) unsigned | null | NO | The period (in seconds) that the slave will wait before trying to reconnect to the master. | 
9 | Enabled_ssl | tinyint(1) | null | NO | Indicates whether the server supports SSL connections. | 
10 | Ssl_ca | text | null | YES | The file used for the Certificate Authority (CA) certificate. | 
11 | Ssl_capath | text | null | YES | The path to the Certificate Authority (CA) certificates. | 
12 | Ssl_cert | text | null | YES | The name of the SSL certificate file. | 
13 | Ssl_cipher | text | null | YES | The name of the cipher in use for the SSL connection. | 
14 | Ssl_key | text | null | YES | The name of the SSL key file. | 
15 | Ssl_verify_server_cert | tinyint(1) | null | NO | Whether to verify the server certificate. | 
16 | Heartbeat | float | null | NO |  | 
17 | Bind | text | null | YES | Displays which interface is employed when connecting to the MySQL server | 
18 | Ignored_server_ids | text | null | YES | The number of server IDs to be ignored, followed by the actual server IDs | 
19 | Uuid | text | null | YES | The master server uuid. | 
20 | Retry_count | bigint(20) unsigned | null | NO | Number of reconnect attempts, to the master, before giving up. | 
21 | Ssl_crl | text | null | YES | The file used for the Certificate Revocation List (CRL) | 
22 | Ssl_crlpath | text | null | YES | The path used for Certificate Revocation List (CRL) files | 
23 | Enabled_auto_position | tinyint(1) | null | NO | Indicates whether GTIDs will be used to retrieve events from the master. | 
24 | Channel_name | char(64) | null | NO | The channel on which the slave is connected to a source. Used in Multisource Replication | 
25 | Tls_version | text | null | YES | Tls version | 

```sql
CREATE TABLE `slave_master_info` (
  `Number_of_lines` int(10) unsigned NOT NULL COMMENT 'Number of lines in the file.',
  `Master_log_name` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT 'The name of the master binary log currently being read from the master.',
  `Master_log_pos` bigint(20) unsigned NOT NULL COMMENT 'The master log position of the last read event.',
  `Host` char(64) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT 'The host name of the master.',
  `User_name` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The user name used to connect to the master.',
  `User_password` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The password used to connect to the master.',
  `Port` int(10) unsigned NOT NULL COMMENT 'The network port used to connect to the master.',
  `Connect_retry` int(10) unsigned NOT NULL COMMENT 'The period (in seconds) that the slave will wait before trying to reconnect to the master.',
  `Enabled_ssl` tinyint(1) NOT NULL COMMENT 'Indicates whether the server supports SSL connections.',
  `Ssl_ca` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The file used for the Certificate Authority (CA) certificate.',
  `Ssl_capath` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The path to the Certificate Authority (CA) certificates.',
  `Ssl_cert` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The name of the SSL certificate file.',
  `Ssl_cipher` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The name of the cipher in use for the SSL connection.',
  `Ssl_key` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The name of the SSL key file.',
  `Ssl_verify_server_cert` tinyint(1) NOT NULL COMMENT 'Whether to verify the server certificate.',
  `Heartbeat` float NOT NULL,
  `Bind` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'Displays which interface is employed when connecting to the MySQL server',
  `Ignored_server_ids` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The number of server IDs to be ignored, followed by the actual server IDs',
  `Uuid` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The master server uuid.',
  `Retry_count` bigint(20) unsigned NOT NULL COMMENT 'Number of reconnect attempts, to the master, before giving up.',
  `Ssl_crl` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The file used for the Certificate Revocation List (CRL)',
  `Ssl_crlpath` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'The path used for Certificate Revocation List (CRL) files',
  `Enabled_auto_position` tinyint(1) NOT NULL COMMENT 'Indicates whether GTIDs will be used to retrieve events from the master.',
  `Channel_name` char(64) NOT NULL COMMENT 'The channel on which the slave is connected to a source. Used in Multisource Replication',
  `Tls_version` text CHARACTER SET utf8 COLLATE utf8_bin COMMENT 'Tls version',
  PRIMARY KEY (`Channel_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='Master Information'
```

[▲ top](#table-doc-top)

## slave_relay_log_info <a id="table-slave_relay_log_info"></a>

BASE TABLE

> Relay Log Information 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Number_of_lines | int(10) unsigned | null | NO | Number of lines in the file or rows in the table. Used to version table definitions. | 
2 | Relay_log_name | text | null | NO | The name of the current relay log file. | 
3 | Relay_log_pos | bigint(20) unsigned | null | NO | The relay log position of the last executed event. | 
4 | Master_log_name | text | null | NO | The name of the master binary log file from which the events in the relay log file were read. | 
5 | Master_log_pos | bigint(20) unsigned | null | NO | The master log position of the last executed event. | 
6 | Sql_delay | int(11) | null | NO | The number of seconds that the slave must lag behind the master. | 
7 | Number_of_workers | int(10) unsigned | null | NO |  | 
8 | Id | int(10) unsigned | null | NO | Internal Id that uniquely identifies this record. | 
9 | Channel_name | char(64) | null | NO | The channel on which the slave is connected to a source. Used in Multisource Replication | 

```sql
CREATE TABLE `slave_relay_log_info` (
  `Number_of_lines` int(10) unsigned NOT NULL COMMENT 'Number of lines in the file or rows in the table. Used to version table definitions.',
  `Relay_log_name` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT 'The name of the current relay log file.',
  `Relay_log_pos` bigint(20) unsigned NOT NULL COMMENT 'The relay log position of the last executed event.',
  `Master_log_name` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT 'The name of the master binary log file from which the events in the relay log file were read.',
  `Master_log_pos` bigint(20) unsigned NOT NULL COMMENT 'The master log position of the last executed event.',
  `Sql_delay` int(11) NOT NULL COMMENT 'The number of seconds that the slave must lag behind the master.',
  `Number_of_workers` int(10) unsigned NOT NULL,
  `Id` int(10) unsigned NOT NULL COMMENT 'Internal Id that uniquely identifies this record.',
  `Channel_name` char(64) NOT NULL COMMENT 'The channel on which the slave is connected to a source. Used in Multisource Replication',
  PRIMARY KEY (`Channel_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='Relay Log Information'
```

[▲ top](#table-doc-top)

## slave_worker_info <a id="table-slave_worker_info"></a>

BASE TABLE

> Worker Information 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Id | int(10) unsigned | null | NO |  | 
2 | Relay_log_name | text | null | NO |  | 
3 | Relay_log_pos | bigint(20) unsigned | null | NO |  | 
4 | Master_log_name | text | null | NO |  | 
5 | Master_log_pos | bigint(20) unsigned | null | NO |  | 
6 | Checkpoint_relay_log_name | text | null | NO |  | 
7 | Checkpoint_relay_log_pos | bigint(20) unsigned | null | NO |  | 
8 | Checkpoint_master_log_name | text | null | NO |  | 
9 | Checkpoint_master_log_pos | bigint(20) unsigned | null | NO |  | 
10 | Checkpoint_seqno | int(10) unsigned | null | NO |  | 
11 | Checkpoint_group_size | int(10) unsigned | null | NO |  | 
12 | Checkpoint_group_bitmap | blob | null | NO |  | 
13 | Channel_name | char(64) | null | NO | The channel on which the slave is connected to a source. Used in Multisource Replication | 

```sql
CREATE TABLE `slave_worker_info` (
  `Id` int(10) unsigned NOT NULL,
  `Relay_log_name` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `Relay_log_pos` bigint(20) unsigned NOT NULL,
  `Master_log_name` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `Master_log_pos` bigint(20) unsigned NOT NULL,
  `Checkpoint_relay_log_name` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `Checkpoint_relay_log_pos` bigint(20) unsigned NOT NULL,
  `Checkpoint_master_log_name` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `Checkpoint_master_log_pos` bigint(20) unsigned NOT NULL,
  `Checkpoint_seqno` int(10) unsigned NOT NULL,
  `Checkpoint_group_size` int(10) unsigned NOT NULL,
  `Checkpoint_group_bitmap` blob NOT NULL,
  `Channel_name` char(64) NOT NULL COMMENT 'The channel on which the slave is connected to a source. Used in Multisource Replication',
  PRIMARY KEY (`Channel_name`,`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='Worker Information'
```

[▲ top](#table-doc-top)

## slow_log <a id="table-slow_log"></a>

BASE TABLE

> Slow log 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | start_time | timestamp(6) | CURRENT_TIMESTAMP(6) | NO |  | 
2 | user_host | mediumtext | null | NO |  | 
3 | query_time | time(6) | null | NO |  | 
4 | lock_time | time(6) | null | NO |  | 
5 | rows_sent | int(11) | null | NO |  | 
6 | rows_examined | int(11) | null | NO |  | 
7 | db | varchar(512) | null | NO |  | 
8 | last_insert_id | int(11) | null | NO |  | 
9 | insert_id | int(11) | null | NO |  | 
10 | server_id | int(10) unsigned | null | NO |  | 
11 | sql_text | mediumblob | null | NO |  | 
12 | thread_id | bigint(21) unsigned | null | NO |  | 

```sql
CREATE TABLE `slow_log` (
  `start_time` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  `user_host` mediumtext NOT NULL,
  `query_time` time(6) NOT NULL,
  `lock_time` time(6) NOT NULL,
  `rows_sent` int(11) NOT NULL,
  `rows_examined` int(11) NOT NULL,
  `db` varchar(512) NOT NULL,
  `last_insert_id` int(11) NOT NULL,
  `insert_id` int(11) NOT NULL,
  `server_id` int(10) unsigned NOT NULL,
  `sql_text` mediumblob NOT NULL,
  `thread_id` bigint(21) unsigned NOT NULL
) ENGINE=CSV DEFAULT CHARSET=utf8 COMMENT='Slow log'
```

[▲ top](#table-doc-top)

## tables_priv <a id="table-tables_priv"></a>

BASE TABLE

> Table privileges 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Host | char(60) |  | NO |  | 
2 | Db | char(64) |  | NO |  | 
3 | User | char(32) |  | NO |  | 
4 | Table_name | char(64) |  | NO |  | 
5 | Grantor | char(93) |  | NO |  | 
6 | Timestamp | timestamp | CURRENT_TIMESTAMP | NO |  | 
7 | Table_priv | set(&#39;Select&#39;,&#39;Insert&#39;,&#39;Update&#39;,&#39;Delete&#39;,&#39;Create&#39;,&#39;Drop&#39;,&#39;Grant&#39;,&#39;References&#39;,&#39;Index&#39;,&#39;Alter&#39;,&#39;Create View&#39;,&#39;Show view&#39;,&#39;Trigger&#39;) |  | NO |  | 
8 | Column_priv | set(&#39;Select&#39;,&#39;Insert&#39;,&#39;Update&#39;,&#39;References&#39;) |  | NO |  | 

```sql
CREATE TABLE `tables_priv` (
  `Host` char(60) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Db` char(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `User` char(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Table_name` char(64) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Grantor` char(93) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `Table_priv` set('Select','Insert','Update','Delete','Create','Drop','Grant','References','Index','Alter','Create View','Show view','Trigger') CHARACTER SET utf8 NOT NULL DEFAULT '',
  `Column_priv` set('Select','Insert','Update','References') CHARACTER SET utf8 NOT NULL DEFAULT '',
  PRIMARY KEY (`Host`,`Db`,`User`,`Table_name`),
  KEY `Grantor` (`Grantor`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='Table privileges'
```

[▲ top](#table-doc-top)

## time_zone <a id="table-time_zone"></a>

BASE TABLE

> Time zones 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Time_zone_id | int(10) unsigned | null | NO |  | 
2 | Use_leap_seconds | enum(&#39;Y&#39;,&#39;N&#39;) | N | NO |  | 

```sql
CREATE TABLE `time_zone` (
  `Time_zone_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `Use_leap_seconds` enum('Y','N') NOT NULL DEFAULT 'N',
  PRIMARY KEY (`Time_zone_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='Time zones'
```

[▲ top](#table-doc-top)

## time_zone_leap_second <a id="table-time_zone_leap_second"></a>

BASE TABLE

> Leap seconds information for time zones 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Transition_time | bigint(20) | null | NO |  | 
2 | Correction | int(11) | null | NO |  | 

```sql
CREATE TABLE `time_zone_leap_second` (
  `Transition_time` bigint(20) NOT NULL,
  `Correction` int(11) NOT NULL,
  PRIMARY KEY (`Transition_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='Leap seconds information for time zones'
```

[▲ top](#table-doc-top)

## time_zone_name <a id="table-time_zone_name"></a>

BASE TABLE

> Time zone names 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Name | char(64) | null | NO |  | 
2 | Time_zone_id | int(10) unsigned | null | NO |  | 

```sql
CREATE TABLE `time_zone_name` (
  `Name` char(64) NOT NULL,
  `Time_zone_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='Time zone names'
```

[▲ top](#table-doc-top)

## time_zone_transition <a id="table-time_zone_transition"></a>

BASE TABLE

> Time zone transitions 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Time_zone_id | int(10) unsigned | null | NO |  | 
2 | Transition_time | bigint(20) | null | NO |  | 
3 | Transition_type_id | int(10) unsigned | null | NO |  | 

```sql
CREATE TABLE `time_zone_transition` (
  `Time_zone_id` int(10) unsigned NOT NULL,
  `Transition_time` bigint(20) NOT NULL,
  `Transition_type_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`Time_zone_id`,`Transition_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='Time zone transitions'
```

[▲ top](#table-doc-top)

## time_zone_transition_type <a id="table-time_zone_transition_type"></a>

BASE TABLE

> Time zone transition types 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Time_zone_id | int(10) unsigned | null | NO |  | 
2 | Transition_type_id | int(10) unsigned | null | NO |  | 
3 | Offset | int(11) | 0 | NO |  | 
4 | Is_DST | tinyint(3) unsigned | 0 | NO |  | 
5 | Abbreviation | char(8) |  | NO |  | 

```sql
CREATE TABLE `time_zone_transition_type` (
  `Time_zone_id` int(10) unsigned NOT NULL,
  `Transition_type_id` int(10) unsigned NOT NULL,
  `Offset` int(11) NOT NULL DEFAULT '0',
  `Is_DST` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `Abbreviation` char(8) NOT NULL DEFAULT '',
  PRIMARY KEY (`Time_zone_id`,`Transition_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 STATS_PERSISTENT=0 COMMENT='Time zone transition types'
```

[▲ top](#table-doc-top)

## user <a id="table-user"></a>

BASE TABLE

> Users and global privileges 



No. | Name | Type | Default | Nullable | Comment | Memo
--- | ---- | ---- | ------- | -------- | ------- | ---
1 | Host | char(60) |  | NO |  | 
2 | User | char(32) |  | NO |  | 
3 | Select_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
4 | Insert_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
5 | Update_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
6 | Delete_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
7 | Create_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
8 | Drop_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
9 | Reload_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
10 | Shutdown_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
11 | Process_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
12 | File_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
13 | Grant_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
14 | References_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
15 | Index_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
16 | Alter_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
17 | Show_db_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
18 | Super_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
19 | Create_tmp_table_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
20 | Lock_tables_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
21 | Execute_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
22 | Repl_slave_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
23 | Repl_client_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
24 | Create_view_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
25 | Show_view_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
26 | Create_routine_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
27 | Alter_routine_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
28 | Create_user_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
29 | Event_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
30 | Trigger_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
31 | Create_tablespace_priv | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
32 | ssl_type | enum(&#39;&#39;,&#39;ANY&#39;,&#39;X509&#39;,&#39;SPECIFIED&#39;) |  | NO |  | 
33 | ssl_cipher | blob | null | NO |  | 
34 | x509_issuer | blob | null | NO |  | 
35 | x509_subject | blob | null | NO |  | 
36 | max_questions | int(11) unsigned | 0 | NO |  | 
37 | max_updates | int(11) unsigned | 0 | NO |  | 
38 | max_connections | int(11) unsigned | 0 | NO |  | 
39 | max_user_connections | int(11) unsigned | 0 | NO |  | 
40 | plugin | char(64) | mysql_native_password | NO |  | 
41 | authentication_string | text | null | YES |  | 
42 | password_expired | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 
43 | password_last_changed | timestamp | null | YES |  | 
44 | password_lifetime | smallint(5) unsigned | null | YES |  | 
45 | account_locked | enum(&#39;N&#39;,&#39;Y&#39;) | N | NO |  | 

```sql
CREATE TABLE `user` (
  `Host` char(60) COLLATE utf8_bin NOT NULL DEFAULT '',
  `User` char(32) COLLATE utf8_bin NOT NULL DEFAULT '',
  `Select_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Insert_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Update_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Delete_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Drop_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Reload_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Shutdown_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Process_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `File_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Grant_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `References_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Index_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Alter_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Show_db_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Super_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_tmp_table_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Lock_tables_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Execute_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Repl_slave_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Repl_client_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_view_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Show_view_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_routine_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Alter_routine_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_user_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Event_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Trigger_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `Create_tablespace_priv` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `ssl_type` enum('','ANY','X509','SPECIFIED') CHARACTER SET utf8 NOT NULL DEFAULT '',
  `ssl_cipher` blob NOT NULL,
  `x509_issuer` blob NOT NULL,
  `x509_subject` blob NOT NULL,
  `max_questions` int(11) unsigned NOT NULL DEFAULT '0',
  `max_updates` int(11) unsigned NOT NULL DEFAULT '0',
  `max_connections` int(11) unsigned NOT NULL DEFAULT '0',
  `max_user_connections` int(11) unsigned NOT NULL DEFAULT '0',
  `plugin` char(64) COLLATE utf8_bin NOT NULL DEFAULT 'mysql_native_password',
  `authentication_string` text COLLATE utf8_bin,
  `password_expired` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  `password_last_changed` timestamp NULL DEFAULT NULL,
  `password_lifetime` smallint(5) unsigned DEFAULT NULL,
  `account_locked` enum('N','Y') CHARACTER SET utf8 NOT NULL DEFAULT 'N',
  PRIMARY KEY (`Host`,`User`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='Users and global privileges'
```

[▲ top](#table-doc-top)
