package utils

import (
	"strconv"

	"github.com/tsukiamaoto/bookShop-go/model"

	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

func RelationMap(data []string) map[string]string {
	relations := make(map[string]string)
	for index := range data {
		if index == 0 {
			relations["root"] = data[index]
		} else {
			child, parent := data[index], data[index-1]
			relations[parent] = child
		}
	}
	return relations
}

func BuildTypes(keys []string, relations map[string]string) []*model.Type {
	Types := make([]*model.Type, 0)
	for index, key := range keys {
		if name, ok := relations[key]; ok {
			Type := new(model.Type)
			Type.Name = name
			Type.Level = index
			Types = append(Types, Type)
		}
	}
	return Types
}

func CreatePaginator(query model.Query, ruleKeys map[string]string) *paginator.Paginator {
	p := paginator.New(&paginator.Config{
		Limit: 30,
		// default order by id DESC
		Order: paginator.DESC,
	})

	if query.Limit != "" {
		limit, _ := strconv.Atoi(query.Limit)
		p.SetLimit(limit)
	}
	if key, ok := ruleKeys[query.SortType]; ok {
		rule := paginator.Rule{
			Key: key,
		}
		p.SetRules(rule)
	}
	if query.Order != "" {
		if query.Order == "ASC" || query.Order == "asc"{
			p.SetOrder(paginator.ASC)
		} else if query.Order == "DESC" || query.Order == "desc" {
			p.SetOrder(paginator.DESC)
		}
	}
	if query.Cursor.PrevPage != "" {
		p.SetBeforeCursor(query.Cursor.PrevPage)
	}
	if query.Cursor.NextPage != "" {
		p.SetAfterCursor(query.Cursor.NextPage)
	}

	return p
}
