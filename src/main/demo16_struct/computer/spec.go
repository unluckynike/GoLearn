package computer

/*
导出结构体和字段
如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。同样，如果结构体里的字段首字母大写，它也能被其他包访问到。
*/
type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}
