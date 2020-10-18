package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateBusArticle
// @description   create a BusArticle
// @param     busArticle               model.BusArticle
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateBusArticle(busArticle model.BusArticle) (err error) {
	err = global.GVA_DB.Create(&busArticle).Error
	return err
}

// @title    DeleteBusArticle
// @description   delete a BusArticle
// @auth                     （2020/04/05  20:22）
// @param     busArticle               model.BusArticle
// @return                    error

func DeleteBusArticle(busArticle model.BusArticle) (err error) {
	err = global.GVA_DB.Delete(busArticle).Error
	return err
}

// @title    DeleteBusArticleByIds
// @description   delete BusArticles
// @auth                     （2020/04/05  20:22）
// @param     busArticle               model.BusArticle
// @return                    error

func DeleteBusArticleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.BusArticle{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateBusArticle
// @description   update a BusArticle
// @param     busArticle          *model.BusArticle
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateBusArticle(busArticle *model.BusArticle) (err error) {
	err = global.GVA_DB.Save(busArticle).Error
	return err
}

// @title    GetBusArticle
// @description   get the info of a BusArticle
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    BusArticle        BusArticle

func GetBusArticle(id uint) (err error, busArticle model.BusArticle) {
	err = global.GVA_DB.Where("id = ?", id).First(&busArticle).Error
	return
}

// @title    GetBusArticleInfoList
// @description   get BusArticle list by pagination, 分页获取BusArticle
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetBusArticleInfoList(info request.BusArticleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BusArticle{})
	var busArticles []model.BusArticle
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busArticles).Error

	// 标签处理(字符串转数组)
	// for _, value := range busArticles {
	// 	value.Tag = value.Tag
	// }
	return err, busArticles, total
}
