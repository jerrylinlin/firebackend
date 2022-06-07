package controllers

import (
	"firebackend/common"
	"firebackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ItList(cox *gin.Context) {
	var books []model.Book
	db := common.MysqlInit()
	db.Find(&books)
	cox.JSON(200, gin.H{
		"code": 20000,
		"msg":  "获取成功",
		"data": gin.H{
			"Total": 6,
			"List":  books,
		},
	})
}

func ItSave(cox *gin.Context) {
	var frontarg model.Book
	cox.ShouldBind(&frontarg)
	id := frontarg.ID
	var book model.Book
	db := common.MysqlInit()
	db.AutoMigrate(&book)
	db.Where("id = ?", id).First(&book)
	fmt.Println("book修改前", book)
	if id == book.ID {
		db.Model(&book).Updates(frontarg)
		fmt.Println("更新后的book", book)
		cox.JSON(200, gin.H{
			"code": 20000,
			"msg":  "修改成功",
			"data": gin.H{
				"Total": 4,
				"List":  book,
			},
		})
	} else if id == 0 {
		cox.JSON(200, gin.H{
			"code": 20001,
			"msg":  "id不能为0",
		})
	} else {
		db.Create(&book)
		cox.JSON(200, gin.H{
			"code": 20000,
			"msg":  "创建成功",
			"data": gin.H{
				"Total": 4,
				"List":  book,
			},
		})
	}

}

func ItEdit(cox *gin.Context) {
	var book model.Book
	cox.ShouldBind(&book)
	id := book.ID
	db := common.MysqlInit()
	db.Model(&book).Where("id = ?", id).Updates(book)
	cox.JSON(200, gin.H{
		"code": 20000,
		"msg":  "成功",
		"data": gin.H{
			"Total": 4,
			"List":  book,
		},
	})
}

func ItDel(cox *gin.Context) {
	var book model.Book
	cox.ShouldBind(&book)
	id := book.ID
	db := common.MysqlInit()
	if id != 0 && id == book.ID {
		db.Delete(&book)
		cox.JSON(200, gin.H{
			"code": 20000,
			"msg":  "删除成功",
			"data": gin.H{
				"Total": 4,
			},
		})
	} else {
		cox.JSON(200, gin.H{
			"code": 20000,
			"msg":  "删除失败",
			"data": gin.H{
				"Total": 4,
			},
		})
	}
}
