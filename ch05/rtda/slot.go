package rtda

// Slot 代表一个槽位，通常用于存储局部变量或操作数栈中的值
type Slot struct {
	num int32   // 存储的数值
	ref *Object // 存储的对象引用
}
