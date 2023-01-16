package douban

import (
	"TVHelper/common"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

func SubjectRealTimeHotest() (subjectRealTimeHotest []common.Vod) {
	subjectRealTimeHotest = make([]common.Vod, 0)

	resp, err := dbClient.R().Get("/subject_collection/subject_real_time_hotest/items")
	if err != nil {
		fmt.Printf("[Error] - %v\n", err)
		return
	}

	gjson.Get(resp.String(), "subject_collection_items").ForEach(func(_, v gjson.Result) bool {
		rating := GJsonGetDefault(v.Get("rating.value"), v.Get("null_rating_reason"))
		honorInfos := GJsonArrayToString(v.Get("honor_infos.#.title"), " | ")
		subjectRealTimeHotest = append(subjectRealTimeHotest, common.Vod{
			VodId:      "",
			VodName:    GJsonGetDefault(v.Get("title"), "暂不支持展示"),
			VodPic:     v.Get("pic.normal").String(),
			VodRemarks: strings.TrimSpace(strings.Join([]string{rating, honorInfos}, " ")),
		})
		return true
	})

	return
}
