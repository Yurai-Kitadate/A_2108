package area

import (
	"strconv"

	"github.com/jphacks/A_2108/src/domain"
)

func GetPlace(area, pref int) []domain.Place {
	if area == 0 {
		return getAreas()
	}
	if pref == 0 {
		return getPrefsByArea(area)
	}
	return getCitiesByAreaAndPref(area, pref)
}

var areas = [8]string{"北海道", "東北", "関東", "中部", "関西", "中国", "四国", "九州・沖縄県"}
var prefectures = [47]string{"北海道", "青森県", "岩手県", "宮城県", "秋田県", "山形県", "福島県", "茨城県", "栃木県", "群馬県", "埼玉県", "千葉県", "東京都", "神奈川県", "新潟県", "富山県", "石川県", "福井県", "山梨県", "長野県", "岐阜県", "静岡県", "愛知県", "三重県", "滋賀県", "京都府", "大阪府", "兵庫県", "奈良県", "和歌山県", "鳥取県", "島根県", "岡山県", "広島県", "山口県", "徳島県", "香川県", "愛媛県", "高知県", "福岡県", "佐賀県", "長崎県", "熊本県", "大分県", "宮崎県", "鹿児島県", "沖縄県"}

func getAreas() []domain.Place {
	places := []domain.Place{}
	for i, v := range areas {
		places = append(places, domain.Place{
			Area: i + 1,
			Name: v,
		})
	}
	return places
}

func getPrefsByArea(area int) []domain.Place {
	prefs := []int{}
	if area == 1 {
		prefs = append(prefs, 1)
	} else if area == 2 {
		prefs = append(prefs, 2, 3, 4, 5, 6, 7)
	} else if area == 3 {
		prefs = append(prefs, 8, 9, 10, 11, 12, 13, 14)
	} else if area == 4 {
		prefs = append(prefs, 15, 16, 17, 18, 19, 20, 21, 22, 23)
	} else if area == 5 {
		prefs = append(prefs, 24, 25, 26, 27, 28, 29, 30)
	} else if area == 6 {
		prefs = append(prefs, 31, 32, 33, 34, 35)
	} else if area == 7 {
		prefs = append(prefs, 36, 37, 38, 39)
	} else if area == 8 {
		prefs = append(prefs, 40, 41, 42, 43, 44, 45, 46, 47)
	}

	places := []domain.Place{}

	for _, v := range prefs {
		places = append(places, domain.Place{
			Area:       area,
			Prefecture: v,
			Name:       prefectures[v-1],
		})
	}
	return places
}

func getCitiesByAreaAndPref(area, pref int) []domain.Place {
	places := []domain.Place{}
	res, err := getAPIResponse(pref)
	if err != nil {
		return places
	}
	for _, v := range res.Data {
		intCity, _ := strconv.Atoi(v.ID)
		places = append(places, domain.Place{
			Area:       area,
			Prefecture: pref,
			City:       intCity,
			Name:       v.Name,
		})
	}
	return places
}
