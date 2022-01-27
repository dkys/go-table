package test

import (
	"github.com/dkys/go-table"
	"github.com/gookit/color"
	"testing"
)

func TestAddRow(t *testing.T) {
	tbl := table.Create("ID", "名称", "备注")
	tbl.AddRow([]interface{}{1, "张三", "20岁，单身！黄叶青苔归路，屧粉衣香何处。"})
	tbl.AddRow([]interface{}{2, "李四", "30 years old"})
	t.Logf("\n%v", tbl)
}

func TestData(t *testing.T) {
	tbl := table.Create("ID", "诗词", "作者", "内容")
	data := [][]interface{}{
		{1, "《如梦令·黄叶青苔归路》", "纳兰性德", "木叶纷纷归路，残月晓风何处。消息竟沉沉，今夜相思几许。秋雨，秋雨，一半因风吹去。"},
		{2, "《逢病军人》", "卢纶", "行多有病住无粮，万里还乡未到乡。 蓬鬓哀吟长城下，不堪秋气入金疮。"},
		{3, "《菩萨蛮·风帘燕舞莺啼柳》", "牛峤", "风帘燕舞莺啼柳，妆台约鬓低纤手。钗重髻盘珊，一枝红牡丹。 门前行乐客，白马嘶春色。故故坠金鞭，回头应眼穿。"},
		{4, "《丑奴儿·书博山道中壁》", "辛弃疾", "少年不识愁滋味，爱上层楼。爱上层楼，为赋新词强说愁。 而今识尽愁滋味，欲说还休。欲说还休，却道天凉好个秋。"},
		{5, "《普天乐·秋江忆别》", "赵善庆", "晚天长，秋水苍。山腰落日，雁背斜阳。壁月词，朱唇唱。犹记当年兰舟上，洒洒风泪湿罗裳。钗分凤凰，杯斟鹦鹉，人拆鸳鸯。"},
	}
	tbl.Data(data)
	tbl.AddRow([]interface{}{6, "《江城子·斗转星移玉漏频》", "和凝", "斗转星移玉漏频。已三更，对栖莺。历历花间，似有马啼声。含笑整衣开绣户，斜敛手，下阶迎。"})
	tbl.SetAlign(table.AlignLeft)   // 设置对齐方式
	tbl.SetBorderColor(color.FgRed) // 设置边框颜色
	t.Logf("\n%v", tbl)
}
