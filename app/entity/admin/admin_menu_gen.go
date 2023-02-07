package admin

import (
	sql "database/sql"
	home "github.com/go-home-admin/home/app"
	providers "github.com/go-home-admin/home/bootstrap/providers"
	database "github.com/go-home-admin/home/bootstrap/services/database"
	logrus "github.com/sirupsen/logrus"
	gorm "gorm.io/gorm"
	"strings"
)

type AdminMenu struct {
	Id        uint32         `gorm:"column:id;autoIncrement;type:int unsigned;primaryKey;comment:'ID'" json:"id"` // ID
	ParentId  uint32         `gorm:"column:parent_id;type:int unsigned;default:0;comment:'父级'" json:"parent_id"`  // 父级
	Name      string         `gorm:"column:name;type:varchar(50);comment:'组件名称'" json:"name"`                     // 组件名称
	Component string         `gorm:"column:component;type:varchar(50);comment:'组件'" json:"component"`             // 组件
	Path      *string        `gorm:"column:path;type:varchar(255);comment:'地址'" json:"path"`                      // 地址
	Redirect  *string        `gorm:"column:redirect;type:varchar(255);comment:'重定向'" json:"redirect"`             // 重定向
	Meta      database.JSON  `gorm:"column:meta;type:json;comment:'元数据'" json:"meta"`                             // 元数据
	Sort      uint32         `gorm:"column:sort;type:int unsigned;default:0;comment:'排序'" json:"sort"`            // 排序
	ApiList   *database.JSON `gorm:"column:api_list;type:json;comment:'api'" json:"api_list"`                     // api
	CreatedAt *database.Time `gorm:"column:created_at;type:timestamp;comment:'created_at'" json:"created_at"`     // created_at
	UpdatedAt *database.Time `gorm:"column:updated_at;type:timestamp;comment:'updated_at'" json:"updated_at"`     // updated_at
}

func (receiver *AdminMenu) TableName() string {
	return "admin_menu"
}

type OrmAdminMenu struct {
	db *gorm.DB
}

func (orm *OrmAdminMenu) TableName() string {
	return "admin_menu"
}

func NewOrmAdminMenu(txs ...*gorm.DB) *OrmAdminMenu {
	var tx *gorm.DB
	if len(txs) == 0 {
		tx = providers.GetBean("mysql, admin").(*gorm.DB).Model(&AdminMenu{})
	} else {
		tx = txs[0].Model(&AdminMenu{})
	}
	return &OrmAdminMenu{db: tx}
}

func (orm *OrmAdminMenu) GetDB() *gorm.DB {
	return orm.db
}

func (orm *OrmAdminMenu) GetTableInfo() interface{} {
	return &AdminMenu{}
}

// Create insert the value into database
func (orm *OrmAdminMenu) Create(value interface{}) *gorm.DB {
	return orm.db.Create(value)
}

// CreateInBatches insert the value in batches into database
func (orm *OrmAdminMenu) CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return orm.db.CreateInBatches(value, batchSize)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (orm *OrmAdminMenu) Save(value interface{}) *gorm.DB {
	return orm.db.Save(value)
}

func (orm *OrmAdminMenu) Row() *sql.Row {
	return orm.db.Row()
}

func (orm *OrmAdminMenu) Rows() (*sql.Rows, error) {
	return orm.db.Rows()
}

// Scan scan value to a struct
func (orm *OrmAdminMenu) Scan(dest interface{}) *gorm.DB {
	return orm.db.Scan(dest)
}

func (orm *OrmAdminMenu) ScanRows(rows *sql.Rows, dest interface{}) error {
	return orm.db.ScanRows(rows, dest)
}

// Connection  use a db conn to execute Multiple commands,this conn will put conn pool after it is executed.
func (orm *OrmAdminMenu) Connection(fc func(tx *gorm.DB) error) (err error) {
	return orm.db.Connection(fc)
}

// Transaction start a transaction as a block, return error will rollback, otherwise to commit.
func (orm *OrmAdminMenu) Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return orm.db.Transaction(fc, opts...)
}

// Begin begins a transaction
func (orm *OrmAdminMenu) Begin(opts ...*sql.TxOptions) *gorm.DB {
	return orm.db.Begin(opts...)
}

// Commit commit a transaction
func (orm *OrmAdminMenu) Commit() *gorm.DB {
	return orm.db.Commit()
}

// Rollback rollback a transaction
func (orm *OrmAdminMenu) Rollback() *gorm.DB {
	return orm.db.Rollback()
}

func (orm *OrmAdminMenu) SavePoint(name string) *gorm.DB {
	return orm.db.SavePoint(name)
}

func (orm *OrmAdminMenu) RollbackTo(name string) *gorm.DB {
	return orm.db.RollbackTo(name)
}

// Exec execute raw sql
func (orm *OrmAdminMenu) Exec(sql string, values ...interface{}) *gorm.DB {
	return orm.db.Exec(sql, values...)
}

// Exists 检索对象是否存在
func (orm *OrmAdminMenu) Exists() (bool, error) {
	dest := &struct {
		H int `json:"h"`
	}{}
	db := orm.db.Select("1 as h").Limit(1).Find(dest)
	return dest.H == 1, db.Error
}

// ------------ 以下是单表独有的函数, 便捷字段条件, Laravel风格操作 ---------

func (orm *OrmAdminMenu) Insert(row *AdminMenu) error {
	return orm.db.Create(row).Error
}

func (orm *OrmAdminMenu) Inserts(rows []*AdminMenu) *gorm.DB {
	return orm.db.Create(rows)
}

func (orm *OrmAdminMenu) Order(value interface{}) *OrmAdminMenu {
	orm.db.Order(value)
	return orm
}

func (orm *OrmAdminMenu) Group(name string) *OrmAdminMenu {
	orm.db.Group(name)
	return orm
}

func (orm *OrmAdminMenu) Limit(limit int) *OrmAdminMenu {
	orm.db.Limit(limit)
	return orm
}

func (orm *OrmAdminMenu) Offset(offset int) *OrmAdminMenu {
	orm.db.Offset(offset)
	return orm
}

// 直接查询列表, 如果需要条数, 使用Find()
func (orm *OrmAdminMenu) Get() AdminMenuList {
	got, _ := orm.Find()
	return got
}

// Pluck used to query single column from a model as a map
//
//	var ages []int64
//	db.Model(&users).Pluck("age", &ages)
func (orm *OrmAdminMenu) Pluck(column string, dest interface{}) *gorm.DB {
	return orm.db.Pluck(column, dest)
}

// Delete 有条件删除
func (orm *OrmAdminMenu) Delete(conds ...interface{}) *gorm.DB {
	return orm.db.Delete(&AdminMenu{}, conds...)
}

// DeleteAll 删除所有
func (orm *OrmAdminMenu) DeleteAll() *gorm.DB {
	return orm.db.Exec("DELETE FROM admin_menu")
}

func (orm *OrmAdminMenu) Count() int64 {
	var count int64
	orm.db.Count(&count)
	return count
}

// First 检索单个对象
func (orm *OrmAdminMenu) First(conds ...interface{}) (*AdminMenu, bool) {
	dest := &AdminMenu{}
	db := orm.db.Limit(1).Find(dest, conds...)
	return dest, db.RowsAffected == 1
}

// Take return a record that match given conditions, the order will depend on the database implementation
func (orm *OrmAdminMenu) Take(conds ...interface{}) (*AdminMenu, int64) {
	dest := &AdminMenu{}
	db := orm.db.Take(dest, conds...)
	return dest, db.RowsAffected
}

// Last find last record that match given conditions, order by primary key
func (orm *OrmAdminMenu) Last(conds ...interface{}) (*AdminMenu, int64) {
	dest := &AdminMenu{}
	db := orm.db.Last(dest, conds...)
	return dest, db.RowsAffected
}

func (orm *OrmAdminMenu) Find(conds ...interface{}) (AdminMenuList, int64) {
	list := make([]*AdminMenu, 0)
	tx := orm.db.Find(&list, conds...)
	if tx.Error != nil {
		logrus.Error(tx.Error)
	}
	return list, tx.RowsAffected
}

// Paginate 分页
func (orm *OrmAdminMenu) Paginate(page int, limit int) (AdminMenuList, int64) {
	var total int64
	list := make([]*AdminMenu, 0)
	orm.db.Count(&total)
	if total > 0 {
		if page == 0 {
			page = 1
		}

		offset := (page - 1) * limit
		tx := orm.db.Offset(offset).Limit(limit).Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
	}

	return list, total
}

// FindInBatches find records in batches
func (orm *OrmAdminMenu) FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB {
	return orm.db.FindInBatches(dest, batchSize, fc)
}

// FirstOrInit gets the first matched record or initialize a new instance with given conditions (only works with struct or map conditions)
func (orm *OrmAdminMenu) FirstOrInit(dest *AdminMenu, conds ...interface{}) (*AdminMenu, *gorm.DB) {
	return dest, orm.db.FirstOrInit(dest, conds...)
}

// FirstOrCreate gets the first matched record or create a new one with given conditions (only works with struct, map conditions)
func (orm *OrmAdminMenu) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
	return orm.db.FirstOrCreate(dest, conds...)
}

// Update update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminMenu) Update(column string, value interface{}) *gorm.DB {
	return orm.db.Update(column, value)
}

// Updates update attributes with callbacks, refer: https://gorm.io/docs/update.html#Update-Changed-Fields
func (orm *OrmAdminMenu) Updates(values interface{}) *gorm.DB {
	return orm.db.Updates(values)
}

func (orm *OrmAdminMenu) UpdateColumn(column string, value interface{}) *gorm.DB {
	return orm.db.UpdateColumn(column, value)
}

func (orm *OrmAdminMenu) UpdateColumns(values interface{}) *gorm.DB {
	return orm.db.UpdateColumns(values)
}

func (orm *OrmAdminMenu) Where(query interface{}, args ...interface{}) *OrmAdminMenu {
	orm.db.Where(query, args...)
	return orm
}

func (orm *OrmAdminMenu) Select(query interface{}, args ...interface{}) *OrmAdminMenu {
	orm.db.Select(query, args...)
	return orm
}

func (orm *OrmAdminMenu) Sum(field string) int64 {
	type result struct {
		S int64 `json:"s"`
	}
	ret := result{}
	orm.db.Select("SUM(`" + field + "`) AS s").Scan(&ret)
	return ret.S
}

func (orm *OrmAdminMenu) And(fuc func(orm *OrmAdminMenu)) *OrmAdminMenu {
	ormAnd := NewOrmAdminMenu()
	fuc(ormAnd)
	orm.db.Where(ormAnd.db)
	return orm
}

func (orm *OrmAdminMenu) Or(fuc func(orm *OrmAdminMenu)) *OrmAdminMenu {
	ormOr := NewOrmAdminMenu()
	fuc(ormOr)
	orm.db.Or(ormOr.db)
	return orm
}

// Preload preload associations with given conditions
// db.Preload("Orders|orders", "state NOT IN (?)", "cancelled").Find(&users)
func (orm *OrmAdminMenu) Preload(query string, args ...interface{}) *OrmAdminMenu {
	arr := strings.Split(query, ".")
	for i, _ := range arr {
		arr[i] = home.StringToHump(arr[i])
	}
	orm.db.Preload(strings.Join(arr, "."), args...)
	return orm
}

// Joins specify Joins conditions
// db.Joins("Account|account").Find(&user)
// db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
// db.Joins("Account", DB.Select("id").Where("user_id = users.id AND name = ?", "someName").Model(&Account{}))
func (orm *OrmAdminMenu) Joins(query string, args ...interface{}) *OrmAdminMenu {
	if !strings.Contains(query, " ") {
		query = home.StringToHump(query)
	}
	orm.db.Joins(query, args...)
	return orm
}

func (orm *OrmAdminMenu) WhereId(val uint32) *OrmAdminMenu {
	orm.db.Where("`id` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) InsertGetId(row *AdminMenu) uint32 {
	orm.db.Create(row)
	return row.Id
}
func (orm *OrmAdminMenu) WhereIdIn(val []uint32) *OrmAdminMenu {
	orm.db.Where("`id` IN ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereIdGt(val uint32) *OrmAdminMenu {
	orm.db.Where("`id` > ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereIdGte(val uint32) *OrmAdminMenu {
	orm.db.Where("`id` >= ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereIdLt(val uint32) *OrmAdminMenu {
	orm.db.Where("`id` < ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereIdLte(val uint32) *OrmAdminMenu {
	orm.db.Where("`id` <= ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereParentId(val uint32) *OrmAdminMenu {
	orm.db.Where("`parent_id` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereParentIdIn(val []uint32) *OrmAdminMenu {
	orm.db.Where("`parent_id` IN ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereParentIdNe(val uint32) *OrmAdminMenu {
	orm.db.Where("`parent_id` <> ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereName(val string) *OrmAdminMenu {
	orm.db.Where("`name` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereComponent(val string) *OrmAdminMenu {
	orm.db.Where("`component` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WherePath(val string) *OrmAdminMenu {
	orm.db.Where("`path` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereRedirect(val string) *OrmAdminMenu {
	orm.db.Where("`redirect` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereMeta(val database.JSON) *OrmAdminMenu {
	orm.db.Where("`meta` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereSort(val uint32) *OrmAdminMenu {
	orm.db.Where("`sort` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereApiList(val database.JSON) *OrmAdminMenu {
	orm.db.Where("`api_list` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereCreatedAt(val database.Time) *OrmAdminMenu {
	orm.db.Where("`created_at` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereCreatedAtBetween(begin database.Time, end database.Time) *OrmAdminMenu {
	orm.db.Where("`created_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminMenu) WhereCreatedAtLte(val database.Time) *OrmAdminMenu {
	orm.db.Where("`created_at` <= ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereCreatedAtGte(val database.Time) *OrmAdminMenu {
	orm.db.Where("`created_at` >= ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereUpdatedAt(val database.Time) *OrmAdminMenu {
	orm.db.Where("`updated_at` = ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereUpdatedAtBetween(begin database.Time, end database.Time) *OrmAdminMenu {
	orm.db.Where("`updated_at` BETWEEN ? AND ?", begin, end)
	return orm
}
func (orm *OrmAdminMenu) WhereUpdatedAtLte(val database.Time) *OrmAdminMenu {
	orm.db.Where("`updated_at` <= ?", val)
	return orm
}
func (orm *OrmAdminMenu) WhereUpdatedAtGte(val database.Time) *OrmAdminMenu {
	orm.db.Where("`updated_at` >= ?", val)
	return orm
}

type AdminMenuList []*AdminMenu

func (l AdminMenuList) GetIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.Id)
	}
	return got
}
func (l AdminMenuList) GetIdMap() map[uint32]*AdminMenu {
	got := make(map[uint32]*AdminMenu)
	for _, val := range l {
		got[val.Id] = val
	}
	return got
}
func (l AdminMenuList) GetParentIdList() []uint32 {
	got := make([]uint32, 0)
	for _, val := range l {
		got = append(got, val.ParentId)
	}
	return got
}
func (l AdminMenuList) GetParentIdMap() map[uint32]*AdminMenu {
	got := make(map[uint32]*AdminMenu)
	for _, val := range l {
		got[val.ParentId] = val
	}
	return got
}
