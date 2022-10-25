package boltdb

import (
	"fmt"
	"github.com/boltdb/bolt"
	"testing"
)

// TestOpen 测试打开数据库
func TestBoltdb(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	// 使用 Update 方法进行读写事务
	err = db.Update(func(tx *bolt.Tx) error {
		// CreateBucket 创建一张表
		block, err := tx.CreateBucket([]byte("testBlock"))
		if err != nil {
			t.Fatal(err)
			return err
		}
		// 向 Block 表插入 key/value 数据，都是使用 []byte 类型
		if block != nil {
			block.Put([]byte("name"), []byte("xlao"))
		}
		return nil
	})
}

// TestReadBoltdb 测试读取表数据
func TestReadBoltdb(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	// 利用 View 只读数据
	err = db.View(func(tx *bolt.Tx) error {
		// 指定读取的表
		b := tx.Bucket([]byte("testBlock"))
		// Get 方法通过 key 获取 value
		v := b.Get([]byte("name"))
		fmt.Printf("The name value is: %s\n", v)
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

}
