package repository //因為在repository folder中

import (
	"backend/internal/models"
	"database/sql"
)

//即使你在将来改变了数据库实现，只要新的实现符合 DatabaseRepo 接口，你就可以无缝替换它，而无需更改依赖这些方法的代码
type DatabaseRepo interface {
	// 这个接口定义了两种方法
	Connection() *sql.DB                 // 返回一个数据库连接指针
	AllMovies() ([]*models.Movie, error) // 不带参数，返回一个电影的切片和可能的错误
}
