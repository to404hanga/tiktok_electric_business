package dao

type MerchandiseDAO interface {
}

type Merchandise struct {
	Id          int64 `gorm:"primaryKey,autoIncrement"`
	Name        string
	Description string
	ImageUrl    string
	Price       int64 `gorm:"index"` // 单位：分
	InventoryId int64 // 库存 id
	CategoryId  int64 // 类目 id
	CreatedAt   int64
	UpdatedAt   int64 `gorm:"index"`
}

type Inventory struct {
	Id        int64 `gorm:"primaryKey,autoIncrement"`
	Residual  int64 `gorm:"index"` // 剩余库存
	UpdatedAt int64 `gorm:"index"`
}
