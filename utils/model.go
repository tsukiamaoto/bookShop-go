package utils

func ConvertProductModelColumn2MapString() (map[string]string, map[string]string) {
	var keys = map[string]string{
		"id":        "ID",
		"created":   "CreatedAt",
	}
	var associationKeys = map[string]string{
		"price":     "price",
		"inventory": "inventory",
	}

	return keys, associationKeys
}
