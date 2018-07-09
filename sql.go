package main

const columDefSQL = `
SELECT
  ORDINAL_POSITION, COLUMN_NAME, COLUMN_COMMENT,
  DATA_TYPE, 
  case when IS_NULLABLE = 'YES' then true else false end as NULLABLE,
  case when COLUMN_KEY = 'PRI' then true else false end as IS_PRIMARY_KEY,
  COLUMN_TYPE
FROM INFORMATION_SCHEMA.COLUMNS
WHERE TABLE_SCHEMA = ?
and TABLE_NAME = ?
`

const tableDefSQL = `
select TABLE_NAME, TABLE_COMMENT from information_schema.tables
where TABLE_SCHEMA = ?
`

const fkDefSQL = `
select k.COLUMN_NAME, k.REFERENCED_TABLE_NAME, k.REFERENCED_COLUMN_NAME,
  k.CONSTRAINT_NAME,
  case when c1.COLUMN_KEY = 'PRI' then true else false end as IS_SOURCE_COL_PRIMARY_KEY,
  case when c2.COLUMN_KEY = 'PRI' then true else false end as IS_TARGET_COL_PRIMARY_KEY
from information_schema.KEY_COLUMN_USAGE k
  join INFORMATION_SCHEMA.COLUMNS c1 -- for source
    on k.COLUMN_NAME = c1.COLUMN_NAME
       and k.TABLE_NAME = c1.TABLE_NAME
       and k.TABLE_SCHEMA = c1.TABLE_SCHEMA
  join INFORMATION_SCHEMA.COLUMNS c2 -- for target
    on k.REFERENCED_COLUMN_NAME = c2.COLUMN_NAME
       and k.REFERENCED_TABLE_NAME = c2.TABLE_NAME
       and k.REFERENCED_TABLE_SCHEMA = c2.TABLE_SCHEMA
where REFERENCED_TABLE_SCHEMA is not null -- foreign key only
and k.CONSTRAINT_SCHEMA = ?
and k.TABLE_NAME = ?
`
