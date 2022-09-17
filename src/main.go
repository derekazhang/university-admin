package main

import (
	"apis/common"
	"apis/src/models/admin_user"
	"apis/src/models/course"
	"apis/src/routers"
	"fmt"

	//"github.com/360EntSecGroup-Skylar/excelize"
	//"log"
	"net/http"
	//"strings"
	"time"
)

func main() {

	//43E14B614B173238365172B04C69C509
	//创建excel文件
	//xlsx := excelize.NewFile()
	//index := xlsx.NewSheet("成绩表")
	//
	////创建新表单
	////index := xlsx.NewSheet("新asset.xlsx")
	//f, err := excelize.OpenFile("/Users/huxing/Desktop/ssr 持有2.xlsx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(f.GetSheetName(0))
	//rows := f.GetRows("Table")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//num := 0
	//for _, row := range rows {
	//	//for _, value := range row {
	//	//	fmt.Printf("\t%s", value)
	//	//}
	//
	//	for _, s := range strings.Split(row[3], ";") {
	//		num++
	//		fmt.Println(fmt.Sprintf("C%d", num+2), s)
	//		xlsx.SetCellValue("成绩表", fmt.Sprintf("A%d", num+2), row[0])
	//		//xlsx.SetCellValue("成绩表", fmt.Sprintf("B%d", num+2), row[1])
	//		xlsx.SetCellValue("成绩表", fmt.Sprintf("B%d", num+2), s)
	//	}
	//}
	//xlsx.SetActiveSheet(index)
	//err = xlsx.SaveAs("./成绩表.xlsx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//return
	//fmt.Println(wx.MakeQrCode("test", "pages/design/design", "trial"))
	//return
	//common.GetTempSign()
	//return
	//batchdownImage()
	//return
	//hash, e := bcrypt.GenerateFromPassword([]byte(`123456`), bcrypt.DefaultCost) //加密处理
	//if e != nil {
	//	return
	//}
	//fmt.Println(string(hash))
	//return

	db := common.GetGorm()
	db.AutoMigrate(&course.Course{}, &admin_user.AdminUser{}, &admin_user.AdminUserRole{})
	// redis.GetRedis()

	//cosClient := common.GetCos()
	//_, err := cosClient.Object.Put(context.Background(), "test.log", bytes.NewReader([]byte(`hello`)), nil)
	//if err != nil {
	//	panic(err)
	//}

	// go func() {
	// 	for {
	// 		order_service.CheckOrderExpired()
	// 		time.Sleep(time.Second * 3000)
	// 	}
	// }()

	go func() {
		srv := &http.Server{
			Addr:         fmt.Sprintf(":%d", common.GetConfig().Https.Port),
			Handler:      routers.Run(),
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		}
		//log.Get().Debug("server started at %v", time.Now())
		fmt.Printf("https server started at %v port: %d \n", time.Now(), common.GetConfig().Https.Port)
		// 监听请求
		if err := srv.ListenAndServeTLS(common.GetConfig().Https.CrtPath, common.GetConfig().Https.KeyPath); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", common.GetConfig().Http.Port),
		Handler:      routers.Run(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	fmt.Printf("http server started at %v port：%d \n", time.Now(), common.GetConfig().Http.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
