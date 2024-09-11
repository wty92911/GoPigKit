# dao

## 顺序
get->create->update->delete
## 一致性
通过gorm.DB的transaction实现事务一致性，为了简便get不需要。
