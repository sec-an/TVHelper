package douban

import (
	"TVHelper/global"
	"TVHelper/internal/common"
	"strings"

	"go.uber.org/zap"

	"github.com/tidwall/gjson"
)

func SubjectRealTimeHotest() (subjectRealTimeHotest []common.Vod, err error) {
	subjectRealTimeHotest = make([]common.Vod, 0)
	resp, err := global.DouBanClient.R().Get("/subject_collection/subject_real_time_hotest/items")
	if err != nil {
		global.Logger.Error("豆瓣实时热门错误", zap.Error(err))
		return
	}
	gjson.Get(resp.String(), "subject_collection_items").ForEach(func(_, v gjson.Result) bool {
		rating := GJsonGetDefault(v.Get("rating.value"), v.Get("null_rating_reason"))
		honorInfos := GJsonArrayToString(v.Get("honor_infos.#.title"), " | ")
		subjectRealTimeHotest = append(subjectRealTimeHotest, common.Vod{
			VodId:   "",
			VodName: GJsonGetDefault(v.Get("title"), "暂不支持展示"),
			//VodPic:     strings.Join([]string{v.Get("pic.normal").String(), "@User-Agent=com.douban.frodo"}, ""),
			VodPic:     v.Get("pic.normal").String(),
			VodRemarks: strings.TrimSpace(strings.Join([]string{rating, honorInfos}, " ")),
		})
		return true
	})
	return
}
