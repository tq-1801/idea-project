package cusfun

import (
	"gorm.io/gorm"
	"regexp"
)

type Fuzzy struct {
	Field   []string
	Keyword string
}
type Repeat struct {
	Field   []string
	Keyword string
}
type ParamsPOJO struct {
	Limit  int
	Offset int
	Filter map[string]interface{}
	Sort   map[string]interface{}
	Fuzzy
	Exact map[string]interface{}
	Range map[string][]interface{}
	Repeat
}

func GetSqlByParams(db *gorm.DB, paramsPOJO ParamsPOJO, count *int64) *gorm.DB {
	if len(paramsPOJO.Filter) != 0 {
		filterMap := Map2ModelMap(paramsPOJO.Filter)
		for key, value := range filterMap {
			db.Where(key+" IN ?", value)
		}
	}
	if len(paramsPOJO.Exact) != 0 {
		exactMap := Map2ModelMap(paramsPOJO.Exact)
		db.Where(exactMap)
	}
	if len(paramsPOJO.Range) != 0 {
		for key, value := range paramsPOJO.Range {
			filedName := UnMarshal(key)
			db.Where(filedName+" BETWEEN ? AND ?", value[0], value[1])
		}
	}
	if len(cleanData(paramsPOJO.Fuzzy.Keyword)) != 0 {
		keyword := paramsPOJO.Fuzzy.Keyword
		likeStr := ""
		for index, field := range paramsPOJO.Fuzzy.Field {
			if index == 0 {
				likeStr += UnMarshal(field) + " LIKE '%" + keyword + "%'"
			} else {
				likeStr += " OR " + UnMarshal(field) + " LIKE '%" + keyword + "%'"
			}
		}
		db.Where(likeStr)
	}
	if len(cleanData(paramsPOJO.Repeat.Keyword)) != 0 {
		keyword := paramsPOJO.Repeat.Keyword
		likeStr := ""
		for index, field := range paramsPOJO.Repeat.Field {
			if index == 0 {
				likeStr += UnMarshal(field) + "= '" + keyword + "'"
			} else {
				likeStr += " OR " + UnMarshal(field) + "= '" + keyword + "'"
			}
		}
		db.Where(likeStr)
	}
	if len(paramsPOJO.Sort) != 0 {
		sortMap := Map2ModelMap(paramsPOJO.Sort)
		for filed, rule := range sortMap {
			sortRule := filed + " " + rule.(string)
			db.Order(sortRule)
		}

	}
	db.Count(count)
	if paramsPOJO.Offset != 0 {
		db.Offset(paramsPOJO.Offset)
	}
	if paramsPOJO.Limit != 0 {
		db.Limit(paramsPOJO.Limit)
	}
	return db

}

//type ParamsPOJO struct {
//	Limit  int
//	Offset int
//	Filter map[string][]interface{}
//	Sort   map[string]string
//	Fuzzy
//	Exact map[string]interface{}
//	Range map[string][]string
//}

//func GetSqlByParams(db *gorm.DB, paramsPOJO ParamsPOJO) *gorm.DB {
//	if len(paramsPOJO.Filter) != 0 {
//		for key, value := range paramsPOJO.Filter {
//			db.Where(key+" IN ?", value)
//		}
//	}
//	if len(paramsPOJO.Exact) != 0 {
//		db.Where(paramsPOJO.Exact)
//	}
//	if len(paramsPOJO.Range) != 0 {
//		for fieldName, rangeValue := range paramsPOJO.Range {
//			db.Where(fieldName+" BETWEEN ? AND ?", rangeValue[0], rangeValue[1])
//		}
//	}
//	if len(cleanData(paramsPOJO.Fuzzy.Keyword)) != 0 {
//		keyword := paramsPOJO.Fuzzy.Keyword
//		likeStr := ""
//		for index, field := range paramsPOJO.Fuzzy.Field {
//			if index == 0 {
//				likeStr += field + " LIKE '%" + keyword + "%'"
//			} else {
//				likeStr += " OR " + field + " LIKE '%" + keyword + "%'"
//			}
//		}
//		db.Where(likeStr)
//	}
//	if len(paramsPOJO.Sort) != 0 {
//		for filed, rule := range paramsPOJO.Sort {
//			sortRule := filed + " " + rule
//			db.Order(sortRule)
//		}
//
//	}
//	if paramsPOJO.Limit != 0 {
//		db.Limit(paramsPOJO.Limit)
//	}
//	if paramsPOJO.Offset != 0 {
//		db.Offset(paramsPOJO.Offset)
//	}
//	return db
//
//}

func cleanData(rst string) string {
	rex := regexp.MustCompile("\\s")
	rst = rex.ReplaceAllString(rst, "")
	return rst
}
