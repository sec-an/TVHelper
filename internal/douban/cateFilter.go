package douban

import (
	"TVHelper/global"
	"TVHelper/internal/common"
	"encoding/base64"
	"encoding/json"
	"math"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/imroc/req/v3"

	"github.com/tidwall/gjson"
)

var count = 30

func CateFilter(cateType, ext, pg, douban string) (cateFilterResult common.Result, err error) {
	var resp *req.Response
	extDecodeBytes, err := base64.StdEncoding.DecodeString(ext)
	if err != nil {
		global.Logger.Error(ext, zap.Error(err))
	}
	curPage, err := strconv.Atoi(pg)
	if err != nil {
		global.Logger.Error(pg, zap.Error(err))
	}
	switch cateType {
	case "0interests":
		status := GJsonGetDefault(gjson.GetBytes(extDecodeBytes, "status"), "mark")
		subtypeTag := gjson.GetBytes(extDecodeBytes, "subtype_tag").String()
		yearTag := GJsonGetDefault(gjson.GetBytes(extDecodeBytes, "year_tag"), "全部")
		path := strings.Join([]string{"/user/", douban, "/interests"}, "")
		resp, err = global.DouBanClient.R().SetQueryParams(map[string]string{
			"type":        "movie",
			"status":      status,
			"subtype_tag": subtypeTag,
			"year_tag":    yearTag,
			"start":       strconv.Itoa((curPage - 1) * count),
			"count":       strconv.Itoa(count),
		}).Get(path)
	case "1hot_gaia":
		area := GJsonGetDefault(gjson.GetBytes(extDecodeBytes, "area"), "全部")
		sort := GJsonGetDefault(gjson.GetBytes(extDecodeBytes, "sort"), "recommend")
		resp, err = global.DouBanClient.R().SetQueryParams(map[string]string{
			"area":  area,
			"sort":  sort,
			"start": strconv.Itoa((curPage - 1) * count),
			"count": strconv.Itoa(count),
		}).Get("/movie/hot_gaia")
	case "2tv_hot", "3show_hot":
		sType := GJsonGetDefault(gjson.GetBytes(extDecodeBytes, "type"), cateType[1:])
		path := strings.Join([]string{"/subject_collection/", sType, "/items"}, "")
		resp, err = global.DouBanClient.R().SetQueryParams(map[string]string{
			"start": strconv.Itoa((curPage - 1) * count),
			"count": strconv.Itoa(count),
		}).Get(path)
	case "6rank_list_movie", "7rank_list_tv":
		id := GJsonGetDefault(gjson.GetBytes(extDecodeBytes, "榜单"),
			strings.Join([]string{strings.Split(cateType, "_")[2],
				"real_time_hotest"},
				"_"))
		path := strings.Join([]string{"/subject_collection/", id, "/items"}, "")
		resp, err = global.DouBanClient.R().SetQueryParams(map[string]string{
			"start": strconv.Itoa((curPage - 1) * count),
			"count": strconv.Itoa(count),
		}).Get(path)
	default:
		sort := ""
		tags := ""
		selectedCategories := map[string]string{
			"类型": gjson.GetBytes(extDecodeBytes, "类型").String(),
			"地区": gjson.GetBytes(extDecodeBytes, "地区").String(),
		}
		sort = GJsonGetDefault(gjson.GetBytes(extDecodeBytes, "sort"), "T")
		tags = GJsonArrayToStringExcept(gjson.GetBytes(extDecodeBytes, `@values`), sort, ",")
		if cateType[1:] == "tv" {
			selectedCategories["形式"] = gjson.GetBytes(extDecodeBytes, "形式").String()
		}
		selectedCategoriesJson, _ := json.Marshal(selectedCategories)
		path := strings.Join([]string{"/", cateType[1:], "/recommend"}, "")
		resp, err = global.DouBanClient.R().SetQueryParams(map[string]string{
			"tags":                tags,
			"sort":                sort,
			"refresh":             "0",
			"selected_categories": string(selectedCategoriesJson),
			"start":               strconv.Itoa((curPage - 1) * count),
			"count":               strconv.Itoa(count),
		}).Get(path)
	}
	if err != nil {
		return
	}

	respStr := resp.String()
	total, err := strconv.Atoi(gjson.Get(respStr, "total").String())
	if err != nil {
		global.Logger.Error(respStr, zap.Error(err))
	}

	cateFilterResult = common.Result{
		Page:      curPage,
		PageCount: int(math.Ceil(float64(total) / float64(count))),
		Limit:     count,
		Total:     total,
	}

	lists := make([]common.Vod, 0)
	path := ""

	switch cateType {
	case "2tv_hot", "3show_hot", "6rank_list_movie", "7rank_list_tv":
		path = "subject_collection_items"
	case "0interests":
		path = "interests.#.subject"
	default:
		path = "items"
	}

	gjson.Get(resp.String(), path).ForEach(func(_, v gjson.Result) bool {
		itemType := v.Get("type").String()
		if itemType == "movie" || itemType == "tv" {
			rating := GJsonGetDefault(v.Get("rating.value"), v.Get("null_rating_reason"))
			honorInfos := GJsonArrayToString(v.Get("honor_infos.#.title"), " | ")
			lists = append(lists, common.Vod{
				VodId: strings.Join([]string{"msearch:", itemType, "__", v.Get("id").String()},
					""),
				VodName: GJsonGetDefault(v.Get("title"), "暂不支持展示"),
				//VodPic:     strings.Join([]string{v.Get("pic.normal").String(), "@User-Agent=com.douban.frodo"}, ""),
				VodPic:     v.Get("pic.normal").String(),
				VodRemarks: strings.TrimSpace(strings.Join([]string{rating, honorInfos}, " ")),
			})
		}
		return true
	})

	cateFilterResult.List = lists

	return
}
