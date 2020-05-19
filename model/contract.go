package model

import (
	"com.lueey.shop/utils"
)

type Contract struct {
	CustomerID     string  // 身份证号
	SalesAdvisorID string  // 用户昵称
	Price          float32 // 价格
	DisPrice       float32 // 折扣价格
	CarBrand       string  // 品牌
	CarColor       string  // 颜色
	CarInterior    string  // 内饰
	CarSeries      string  // 型号
	Timestamp      int64   // 时间戳
}

func (c Contract) String() (str string) {
	lang, err := json.Marshal(c)
	if err == nil {
		str = string(lang)
	}
	return str
}

func CreateContract(customerID string, salesID string, price float32, disprice float32, brand string, color string, interior string, series string) Contract {
	return Contract{
		CustomerID:     customerID,
		SalesAdvisorID: salesID,                 // 用户昵称
		Price:          price,                   // 价格
		DisPrice:       disprice,                // 折扣价格
		CarBrand:       brand,                   // 品牌
		CarColor:       color,                   // 颜色
		CarInterior:    interior,                // 内饰
		CarSeries:      series,                  // 型号
		Timestamp:      utils.NowMilliseconds(), // 时间戳
	}
}
