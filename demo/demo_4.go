package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const (
	ALL_CITY = "安庆,蚌埠,亳州,巢湖,滁州,池州,阜阳,淮北,合肥,淮南,黄山,六安,马鞍山,宿州,铜陵,芜湖,宣城,澳门,北京,重庆,福州,龙岩,宁德,南平,莆田,泉州,厦门,三明,漳州,潮州,东莞,佛山,广州,河源,惠州,江门,揭阳,茂名,梅州,清远,韶关,汕头,汕尾,深圳,云浮,阳江,珠海,湛江,肇庆,中山,白银,定西,甘南,金昌,酒泉,嘉峪关,陇南,临夏,兰州市,平凉,庆阳,天水,武威,张掖,北海,百色,崇左,防城港,贵港,桂林,河池,贺州,来宾,柳州,南宁,钦州,梧州,玉林,安顺,毕节,贵阳,六盘水,黔东南,黔南,黔西南,铜仁,遵义,保定,承德,沧州,邯郸,衡水,廊坊,秦皇岛,石家庄,唐山,邢台,张家口,恩施,鄂州,黄冈,黄石,荆门,荆州,潜江,神农架,十堰,随州,天门,武汉,孝感,咸宁,仙桃,襄阳,宜昌,大庆,大兴安岭,哈尔滨,鹤岗,黑河,佳木斯,鸡西,牡丹江,齐齐哈尔,七台河,绥化,双鸭山,伊春,白沙,保亭,昌江,澄迈,定安,东方,儋州,海口,乐东,临高,陵水,琼海,琼中,三沙,三亚,屯昌,文昌,万宁,五指山,安阳,鹤壁,济源,焦作,开封,漯河,洛阳,南阳,平顶山,濮阳,三门峡,商丘,许昌,新乡,信阳,周口,驻马店,郑州,常德,长沙,郴州,怀化,衡阳,娄底,邵阳,湘潭,湘西,益阳,岳阳,永州,张家界,株洲,白城,白山,长春,吉林,辽源,四平,松原,通化,延边,常州,淮安,连云港,南京,南通,宿迁,苏州,泰州,无锡,徐州,盐城,扬州,镇江,抚州,赣州,吉安,景德镇,九江,南昌,萍乡,上饶,新余,宜春,鹰潭,鞍山,本溪,朝阳,丹东,大连,抚顺,阜新,葫芦岛,锦州,辽阳,盘锦,沈阳,铁岭,营口,阿拉善,包头,巴彦淖尔,赤峰,鄂尔多斯,呼和浩特,呼伦贝尔,通辽,乌海,乌兰察布,兴安,锡林郭勒,固原,石嘴山,吴忠,银川,中卫,果洛,海北,海东,海南,黄南,海西,西宁,玉树,阿坝,巴中,成都,德阳,达州,广安,广元,甘孜,乐山,凉山,泸州,眉山,绵阳,南充,内江,攀枝花,遂宁,雅安,宜宾,自贡,资阳,滨州,东营,德州,菏泽,济南,济宁,聊城,莱芜,临沂,青岛,日照,泰安,潍坊,威海,烟台,淄博,枣庄,上海,安康,宝鸡,汉中,商洛,铜川,渭南,西安,咸阳,延安,榆林,长治,大同,晋城,晋中,临汾,吕梁,朔州,太原,忻州,运城,阳泉,天津,高雄市,花莲县,基隆市,金门县,嘉义市,嘉义县,连江县,苗栗县,南投县,屏东县,澎湖县,台北市,台东县,台南市,桃园县,台中市,新北市,新竹市,新竹县,云林县,宜兰县,彰化县,阿里,昌都,拉萨,林芝,那曲,日喀则,山南,香港,阿克苏,阿拉尔,阿勒泰,博尔塔拉,巴音郭楞,昌吉,哈密,和田,克拉玛依,喀什,克孜勒苏,石河子,塔城,吐鲁番,图木舒克,五家渠,乌鲁木齐,伊犁,保山,楚雄,德宏,大理,迪庆,红河,昆明,临沧,丽江,怒江,普洱,曲靖,文山,西双版纳,玉溪,昭通,杭州,湖州,金华,嘉兴,丽水,宁波,衢州,绍兴,台州,温州,舟山"
)

type WeChatUser struct {
	City string `json:"city"`
}

func main() {
	str := "香港特别行政区"
	arr := strings.Split(ALL_CITY, ",")
	for _, v := range arr {
		index := strings.Contains(str, v)
		if index {
			fmt.Println(str, v)
			break
		}
	}
	nums := []int{3, 5, 0, -2, -4, 9, 1, 2}
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)

	data := `{"subscribe":1,"openid":"oZCjYwt7lwT2E3HRKqJampUxdueY","nickname":"hl","sex":1,"language":"zh_CN","city":"武汉","province":"湖北","country":"中国","headimgurl":"http:\/\/thirdwx.qlogo.cn\/mmopen\/PiajxSqBRaEI5jsA3ClIZ4iayicZPuxIKGgQp8FIhdBRDgof8pT5fZdeoFlIsnLVicgahTzfIib6nTj6OGO5Ct8LU4A\/132","subscribe_time":1511513566,"remark":"","groupid":0,"tagid_list":[],"subscribe_scene":"ADD_SCENE_PROFILE_CARD","qr_scene":0,"qr_scene_str":""}`
	WeChatUser := new(WeChatUser)
	err := json.Unmarshal([]byte(data), &WeChatUser)
	if err != nil {
		return
	}
	fmt.Println(WeChatUser.City)
}