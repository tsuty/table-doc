package main

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/howeyc/gopass"
	"github.com/jessevdk/go-flags"
	"github.com/tsuty/table-doc"
	"log"
	"net/url"
	"os"
	"sort"
)

const (
	defaultHost           = "127.0.0.1"
	defaultMySQLUser      = "root"
	defaultMySQLPort      = 3306
	defaultPostgreSQLUser = "postgres"
	defaultPostgreSQLPort = 5432
)

type Options struct {
	User       string   `long:"user" description:"User name (default: root (MySQL) | postgres (PostgreSQL))" value-name:"name"`
	Password   string   `long:"password" description:"User password (If not given, it's asked from the prompt)" value-name:"password"`
	Host       string   `long:"host" description:"Host name" value-name:"host-name" default:"127.0.0.1"`
	Port       uint     `long:"port" description:"Port number (default: 3306 (MySQL) | 5432 (PostgreSQL))" value-name:"port"`
	Socket     string   `long:"socket" description:"Socket file" value-name:"socket"`
	MySQL      bool     `long:"mysql" description:"MySQL data source"`
	PostgreSQL bool     `long:"postgres" description:"PostgreSQL data source"`
	Help       bool     `long:"help" short:"h" description:"Show this help"`
	Schemas    []string `long:"schema" short:"s" description:"Schema name" value-name:"schema-name"`
	Format     string   `long:"format" short:"f" choice:"html" choice:"bootstrap" choice:"markdown" description:"Output format" default:"bootstrap"`
}

func (o *Options) dataSource() string {
	if o.MySQL {
		return table_doc.DataSourceMySQL
	} else if o.PostgreSQL {
		return table_doc.DataSourcePostgreSQL
	} else {
		return ""
	}
}

func (o *Options) hasPassword() bool {
	if o.Password != "" {
		return true
	}
	return false
}

func (o *Options) buildMySQLConfig() *table_doc.MySQLConfig {
	c := mysql.NewConfig()
	if o.User != "" {
		c.User = o.User
	} else {
		c.User = defaultMySQLUser
	}

	if o.Password != "" {
		c.Passwd = o.Password
	} else {
		password, ok := os.LookupEnv("MYSQL_PWD")
		if ok {
			c.Passwd = password
		} else {
			// like docker
			// https://hub.docker.com/_/mysql
			_, ok := os.LookupEnv("MYSQL_ALLOW_EMPTY_PASSWORD")
			if !ok {
				fmt.Printf("Enter Password: ")
				pass, err := gopass.GetPasswd()
				if err != nil {
					log.Fatal(err)
				}
				c.Passwd = string(pass)
			}
		}
	}

	if o.Socket != "" {
		c.Net = "unix"
		c.Addr = o.Socket
	} else {
		c.Net = "tcp"
		if o.Host != "" {
			c.Addr = o.Host
		} else {
			c.Addr = defaultHost
		}

		if o.Port > 0 {
			c.Addr = fmt.Sprintf("%s:%d", c.Addr, o.Port)
		} else {
			c.Addr = fmt.Sprintf("%s:%d", c.Addr, defaultMySQLPort)
		}
	}
	m := table_doc.MySQLConfig{}
	m.Config = c
	m.Schemas = o.Schemas
	return &m
}

func (o *Options) buildPostgreSQLConfig() *table_doc.PostgresConfig {
	u := url.URL{
		Scheme: "postgres",
	}

	if o.User != "" {
		if o.hasPassword() {
			u.User = url.UserPassword(o.User, o.Password)
		} else {
			password, ok := os.LookupEnv("PGPASSWORD")
			if ok {
				u.User = url.UserPassword(o.User, password)
			} else {
				fmt.Printf("Enter Password: ")
				pass, err := gopass.GetPasswd()
				if err != nil {
					log.Fatal(err)
				}
				u.User = url.UserPassword(o.User, string(pass))
			}
		}
	} else {
		if o.hasPassword() {
			u.User = url.UserPassword(defaultPostgreSQLUser, o.Password)
		} else {
			u.User = url.User(defaultPostgreSQLUser)
		}
	}

	if o.Socket != "" {
		u.Host = o.Socket
	} else {

		if o.Port > 0 {
			u.Host = fmt.Sprintf("%s:%d", o.Host, o.Port)
		} else {
			u.Host = fmt.Sprintf("%s:%d", o.Host, defaultPostgreSQLPort)
		}
	}
	v := url.Values{}
	// PGSSLMODE
	pgsslmode, ok := os.LookupEnv("PGSSLMODE")
	if ok {
		v.Add("sslmode", pgsslmode)
	} else {
		v.Add("sslmode", "disable")
	}
	u.RawQuery = v.Encode()

	p := table_doc.PostgresConfig{}
	p.URL = &u
	p.Schemas = o.Schemas
	return &p
}

func (o *Options) makeConfig() (table_doc.DatabaseConfig, error) {
	if o.dataSource() == table_doc.DataSourceMySQL {
		return o.buildMySQLConfig(), nil
	} else if o.dataSource() == table_doc.DataSourcePostgreSQL {
		return o.buildPostgreSQLConfig(), nil
	} else {
		return nil, errors.New("Unknown data source")
	}
}

func main() {
	opts := Options{}
	parser := flags.NewParser(&opts, flags.None)
	parser.Usage = "[options] dir"
	parser.Name = "table-doc"
	args, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Help {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	if opts.dataSource() == "" {
		fmt.Fprintln(os.Stderr, "Unknown data source.")
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	config, err := opts.makeConfig()
	if err != nil {
		log.Fatal(err)
	}

	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		fmt.Fprintln(os.Stderr, "Require dir argument.")
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	sort.Strings(opts.Schemas)

	option := table_doc.Option{
		SchemaNames: opts.Schemas,
		Path:        path,
		Format:      opts.Format,
		DataSource:  opts.dataSource(),
	}

	if err := table_doc.MakeDoc(config, option); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
