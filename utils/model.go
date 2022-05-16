package utils

func ConvertProductModelColumn2MapString() (map[string]string, map[string]string) {
	var keys = map[string]string{
		"id":              "ID",
		"created":         "CreatedAt",
		"publicationDate": "PublicationDate",
		"publisher":       "Publisher",
		"editor":          "Editor",
		"name":            "Name",
		"description":     "Description",
	}
	var associationKeys = map[string]string{
		"price":     "price",
		"inventory": "inventory",
		"types":     "types",
	}

	return keys, associationKeys
}
