package csv_lib

import (
	"bytes"
	"fmt"
	"go-study/utils"
	"time"
)

func WriteCsv() {



	str := utils.Stamp2Str(time.Now().UnixNano()/1e6, utils.DateFormatNoSpan)
	fmt.Println(str)

	//var buf bytes.Buffer
	//setCsvHead(&buf)
	//buf.WriteString("标题")
	//buf.WriteString("\t")
	//buf.WriteString("作者")
	//buf.WriteString("\t")
	//buf.WriteString("时间")
	//buf.WriteString("\t\n")
	//buf.WriteString("标题")
	//buf.WriteString("\t")
	//buf.WriteString("作者")
	//buf.WriteString("\t")
	//buf.WriteString("时间")
	//buf.WriteString("\t")
	//
	//fmt.Println(buf.String())

	//f, err := os.Create("test2.csv")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()
	//
	//var data = make([][]string, 4)
	//data[0] = []string{"标题", "作者", "时间"}
	//data[1] = []string{"羊皮卷", "鲁迅", "2008"}
	//data[2] = []string{"易筋经", "唐生", "665"}
	//f.WriteString("\xEF\xBB\xBF") // 写入一个UTF-8 BOM
	//
	//w := csv.NewWriter(f)
	//w.Comma = '\t'
	//w.WriteAll(data)
	//w.Flush()

}

var columnSep = "\t"
var rowSep = "\n"

func setCsvHead(buf *bytes.Buffer) {
	buf.WriteString("RLZPXXBZ" + columnSep)          // 人脸抓拍信息标识    非空
	buf.WriteString("SPTXXXYYSXDXBZ" + columnSep)    // 抓怕人脸标识  		非空
	buf.WriteString("SPTXXXCJFSDM" + columnSep)      // 人脸采集方式代码
	buf.WriteString("NLSX" + columnSep)              // 年龄上限
	buf.WriteString("NLXX" + columnSep)              //年龄下限
	buf.WriteString("XBDM" + columnSep)              //性别
	buf.WriteString("MZDM" + columnSep)              // 民族
	buf.WriteString("SFBKRY_PDBZ" + columnSep)       // 是否布控人员
	buf.WriteString("SFGZRY_PDBZ" + columnSep)       // 是否关注人员
	buf.WriteString("CJSBXXBZ" + columnSep)          // 采集设备信息标识
	buf.WriteString("ZSJ_ZXDHZB" + columnSep)        // 左上角_中心点横坐标
	buf.WriteString("ZSJ_ZXDZZB" + columnSep)        //左上角_中心点纵坐标
	buf.WriteString("YXJ_ZXDHZB" + columnSep)        // 右下角_中心点横坐标
	buf.WriteString("YXJ_ZXDZZB" + columnSep)        //右下角_中心点横纵标
	buf.WriteString("TXR_RS" + columnSep)            //同行人_人数
	buf.WriteString("SFDKZ_PDBZ" + columnSep)        // 是否戴口罩
	buf.WriteString("ZMTZDM" + columnSep)            // 着帽特征代码(帽子款式)
	buf.WriteString("SFJSY_PDBZ" + columnSep)        // 是否驾驶员
	buf.WriteString("PDYJTZDM" + columnSep)          // 佩戴眼镜特征代码（眼镜款式）
	buf.WriteString("XBTZDM" + columnSep)            //包特征
	buf.WriteString("SSZZTZDM" + columnSep)          // 上身着装特征代码
	buf.WriteString("XSZZTZDM" + columnSep)          // 下身着装特征代码
	buf.WriteString("ZXTZDM" + columnSep)            // 着鞋特征代码
	buf.WriteString("SY_YS" + columnSep)             // 上衣颜色
	buf.WriteString("KZ_YS" + columnSep)             // 裤子颜色
	buf.WriteString("XZ_YS" + columnSep)             // 鞋子颜色
	buf.WriteString("MZ_YS" + columnSep)             // 帽子颜色
	buf.WriteString("KZ_YS" + columnSep)             // 口罩颜色
	buf.WriteString("TP_URL" + columnSep)            // 人脸图片路径
	buf.WriteString("CJSJ" + columnSep)              // 采集时间
	buf.WriteString("SJGXSM" + columnSep)            // 数据更新标识  1新增 2更新 3删除
	buf.WriteString("SJGSDWDM" + columnSep)          // 数据归属单位代码
	buf.WriteString("SHGSDWMC" + columnSep + rowSep) // 数据归属单位名称
}
