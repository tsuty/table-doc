# table-doc
[![Build Status](https://travis-ci.com/tsuty/table-doc.svg?branch=master)](https://travis-ci.com/tsuty/table-doc)

Create SQL table documents.

## Usage
```
Usage:
  table-doc [options] dir

Application Options:
      --user=name                           User name (default: root (MySQL) | postgres (PostgreSQL))
      --password=password                   User password (If not given, it's asked from the prompt)
      --host=host-name                      Host name (default: 127.0.0.1)
      --port=port                           Port number (default: 3306 (MySQL) | 5432 (PostgreSQL))
      --socket=socket                       Socket file
      --mysql                               MySQL data source
      --postgres                            PostgreSQL data source
  -h, --help                                Show this help
  -s, --schema=schema-name                  Schema name
  -f, --format=[html|bootstrap|markdown]    Output format (default: bootstrap)
```

### Options

#### --password

Override `MYSQL_PWD` or `PGPASSWORD` environment.

If not given and set any environment, it's asked from the prompt.

#### -s, --schema

##### multiple schemas

`--schema=foofoo --schema=hogehoge`

##### alias

`--schema=foofoo_test@foofoo` "foofoo" is alias name of "foofoo_test"

#### --format

* `html` very simple html
* `bootstrap` a little rich html
* `markdown` at first build _bootstrap_ and convert to markdown

## Example

execute following command 

```bash
table-doc --mysql -s mysql -f markdown ./example
```

- [bootstrap](example/mysql.html)
- [markdown](example/mysql.md)

## Memo 

After `table-doc` command, edit HTML of generated for memo. that memo keeps after regenerate HTML. (the markdown always converts from HTML. editing not apply)

- The schema memo is `data-table-doc="schema-memo"` attribute.
- The table memo is `data-table-doc="table-memo"` attribute.
- The column memo is `data-table-doc="column-memo"` attribute.
