package request

import "gin-vue-admin/model"

type BusArticleSearch struct{
    model.BusArticle
    PageInfo
}