# 如何使用
```cgo
type Student struct {
	Id       string     `json:"id"`
	List     []string   `json:"list"`
	Students []*Student `json:"students"`
}

err := copier.CopyWithOptional(&target, source, 
    copier.WithDeepCopyOption(), // 深拷贝选项
    copier.WithInitNilSlice(),  // 初始化数组选项
    copier.WithInitNilPointer()) // 初始化nil对象
    if err != nil {
        t.Error(err)
        return
    }
```