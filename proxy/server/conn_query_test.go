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
		sql := "use `aa`; alter table `aab`.`b` add column c char(255)"
		sql = conn.rewriteSql(sql)
		fmt.Println(sql)
	}

	{
		sql := "use `aa`; alter table `aab`.`aba_1ax_12A-12` add column c char(255)"
		sql = conn.rewriteSql(sql)
		fmt.Println(sql)
	}

}

