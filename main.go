/**
 * @Author: csq
 * @Date: 2022/4/12 2:58 下午
 */
package main

import (
	"database/sql"
	"db2struct/Init"
	"db2struct/config"
	"db2struct/tool"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
)

func main() {
	//初始化
	Init.Init()
	// 连接mysql
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.SysConfig.DBUserName, config.SysConfig.DBPassword, config.SysConfig.DBIp,
		config.SysConfig.DBPort, config.SysConfig.DBName)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println("连接数据库失败,失败原因是:", err)
		return
	}
	//获取所有表
	var sql = "select table_name from information_schema.tables where table_schema=?"
	// 指定生成某表
	if config.SysConfig.Table != "all_in" {
		sql = fmt.Sprintf(" select table_name from  information_schema.columns where table_schema =? and table_name = '%s' ", config.SysConfig.Table)
	}
	rows, err2 := db.Query(sql, config.SysConfig.DBName)
	if err2 != nil {
		fmt.Println("查询数据库失败,失败原因是:", err2)
		return
	}

	defer func() {
		if rows != nil {
			//防止过多连接没有释放 导致内存泄露
			rows.Close()
		}
	}()
	tableStrut := config.DbTable{}
	for rows.Next() {
		err := rows.Scan(&tableStrut.Name)
		if err != nil {
			fmt.Printf("生成表失败,失败原因是:err:%v", err)
			return
		}
		//获取单个表所有字段
		sqlStr := fmt.Sprintf("select column_name columnName, data_type dataType, column_comment columnComment, column_key columnKey, extra from information_schema.columns where table_name = '%s' and table_schema =(select database()) order by ordinal_position ", tableStrut.Name)
		fieldConn, err3 := db.Query(sqlStr)
		if err3 != nil {
			fmt.Println("mysql 获取单个表所有字段 err:", err3)
		}
		defer func() {
			if fieldConn != nil {
				fieldConn.Close()
			}
		}()

		// ----- 拼接生成的struct  start--------
		structStr := fmt.Sprintf("type %s struct { \n", tool.InitialToCapital(tableStrut.Name))
		column := config.Column{}
		for fieldConn.Next() {
			err := fieldConn.Scan(&column.ColumnName, &column.DataType, &column.ColumnComment, &column.ColumnKey, &column.Extra)
			if err != nil {
				fmt.Printf("获取失败,失败原因是:err:%v", err)
				return
			}
			structStr += "    " + tool.InitialToCapital(column.ColumnName)
			if column.DataType == "int" || column.DataType == "tinyint" {
				structStr += " int "
			} else if column.DataType == "decimal" {
				structStr += " float64 "
			} else {
				structStr += " string "
			}

			if column.Extra != "auto_increment" {
				structStr += fmt.Sprintf("`gorm:\"comment('%s')\" json:\"%s\"` \n",
					column.ColumnComment, column.ColumnName)
			} else {
				structStr += fmt.Sprintf("`gorm:\"not null comment('%s') INT(11)\" json:\"%s\"` \n",
					column.ColumnComment, column.ColumnName)
			}

		}
		structStr += "}"
		modelHead := "package model \n\n"
		// ----- 拼接生成的struct end --------
		if tool.PathProcessing(config.SysConfig.Path) {
			fmt.Println("路径不匹配")
			return
		}
		//导出文件并创建文件夹
		body := modelHead + structStr
		filename := fmt.Sprintf("%s/%s.go", config.SysConfig.Path, tableStrut.Name)
		error2 := os.MkdirAll(config.SysConfig.Path, os.ModePerm)
		if error2 != nil {
			fmt.Println("创建文件夹 err:", error2)
		}
		err4 := ioutil.WriteFile(filename, []byte(body), 0666)
		if err4 != nil {
			fmt.Println("写入文件错误:", err4)
		}
	}
	fmt.Println("End")
}
