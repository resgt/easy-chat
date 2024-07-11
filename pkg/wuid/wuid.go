/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

/*
CREATE TABLE `wuid` (
  `h` int(10) NOT NULL AUTO_INCREMENT,
  `x` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`h`),
  UNIQUE KEY `h` (`h`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
CREATE TABLE `wuid` (   `h` int(10) NOT NULL AUTO_INCREMENT,   `x` tinyint(4) NOT NULL DEFAULT '0',   PRIMARY KEY (`x`),   UNIQUE KEY `h` (`h`) ) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;
*/
// 根据数据库生成唯一的uuid
package wuid

import (
	"database/sql"
	"fmt"
	"github.com/edwingeng/wuid/mysql/wuid"
)

var w *wuid.WUID

func Init(dsn string) {

	newDB := func() (*sql.DB, bool, error) {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, false, err
		}
		return db, true, nil
	}

	w = wuid.NewWUID("default", nil)
	_ = w.LoadH28FromMysql(newDB, "wuid")
}

func GenUid(dsn string) string {
	if w == nil {
		Init(dsn)
	}

	return fmt.Sprintf("%#016x", w.Next())
}
