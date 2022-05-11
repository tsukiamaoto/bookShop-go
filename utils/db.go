package utils

import "github.com/tsukiamaoto/bookShop-go/model"

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
