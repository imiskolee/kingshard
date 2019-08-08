package server

import (
	"fmt"
	"testing"
)

func TestRewriteSQL(t *testing.T) {

	conn := &ClientConn{}

	{
		sql := "use `a`; alter table `a`.`b` add column c char(255)"
		sql = conn.rewriteSql(sql)
		fmt.Println(sql)
	}

	{
		sql := "use `a`; alter table `b` add column c char(255)"
		sql = conn.rewriteSql(sql)
		fmt.Println(sql)
	}

	{
		sql := "    use `aa`; alter table `aab`.`b` add column c char(255)"
		sql = conn.rewriteSql(sql)
		fmt.Println(sql)
	}

	{
		sql := "use `aa`; alter table `aab`.`aba_1ax_12A-12` add column c char(255)"
		sql = conn.rewriteSql(sql)
		fmt.Println(sql)
	}

	{
		sql := "USE `bindo_gateway`; CREATE UNIQUE INDEX `index_name` ON `t1` (`uuid`);"
		sql = conn.rewriteSql(sql)
		fmt.Println(sql)
	}

	{
		sql := "USE `bindo_gateway`; CREATE INDEX `index_name2` ON `t2` (`uuid`);"
		sql = conn.rewriteSql(sql)
		fmt.Println(sql)
	}

}

