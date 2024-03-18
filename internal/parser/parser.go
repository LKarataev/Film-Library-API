package parser

import (
	"encoding/json"
	"fmt"
	"io"
)

func ParseFilmOptions(body io.ReadCloser) (int, map[string]string, error) {
	reqBody, err := io.ReadAll(body)
	if err != nil {
		return 0, nil, err
	}

	var reqInterface interface{}
	err = json.Unmarshal(reqBody, &reqInterface)
	if err != nil {
		return 0, nil, err
	}

	var values = map[string]string{}
	var filmid int
	for key, val := range reqInterface.(map[string]interface{}) {
		switch key {
		case "id":
			switch val.(type) {
			case float64:
				filmid = int(val.(float64))
			default:
				return 0, nil, fmt.Errorf(`film id should be a number`)
			}
		case "year":
			switch val.(type) {
			case float64:
				values[key] = fmt.Sprintf("%d", int(val.(float64)))
			default:
				return 0, nil, fmt.Errorf(`year should be a number`)
			}
		case "rating":
			switch val.(type) {
			case float64:
				values[key] = fmt.Sprintf("%f", val.(float64))
			default:
				return 0, nil, fmt.Errorf(`rating should be a floating-point number`)
			}
		case "name", "description":
			switch val.(type) {
			case []byte:
				values[key] = string(val.([]byte))
			case string:
				values[key] = val.(string)
			default:
				return 0, nil, fmt.Errorf(`name and description should be a 'text'`)
			}
		default:
			return 0, nil, fmt.Errorf(`some invalid option in request body`)
		}
	}
	return filmid, values, nil
}

func ParseActorOptions(body io.ReadCloser) (int, map[string]string, error) {
	reqBody, err := io.ReadAll(body)
	if err != nil {
		return 0, nil, err
	}

	var reqInterface interface{}
	err = json.Unmarshal(reqBody, &reqInterface)
	if err != nil {
		return 0, nil, err
	}

	var values = map[string]string{}
	var actorId int
	for key, val := range reqInterface.(map[string]interface{}) {
		switch key {
		case "id":
			switch val.(type) {
			case float64:
				actorId = int(val.(float64))
			default:
				return 0, nil, fmt.Errorf(`actor id should be a number`)
			}
		case "name", "gender", "birthday":
			switch val.(type) {
			case []byte:
				values[key] = string(val.([]byte))
			case string:
				values[key] = val.(string)
			default:
				return 0, nil, fmt.Errorf(`name, gender and birthday should be a 'text'`)
			}
		default:
			return 0, nil, fmt.Errorf(`some invalid option in request body`)
		}
	}
	return actorId, values, nil
}

func ParseFilmActorsIdsOptions(body io.ReadCloser) (map[string]string, []int, error) {
	reqBody, err := io.ReadAll(body)
	if err != nil {
		return nil, nil, err
	}

	var reqInterface interface{}
	err = json.Unmarshal(reqBody, &reqInterface)
	if err != nil {
		return nil, nil, err
	}

	var values = map[string]string{}
	var actorsIds []int
	for key, val := range reqInterface.(map[string]interface{}) {
		switch key {
		case "actors":
			switch val.(type) {
			case []interface{}:
				for _, actor := range val.([]interface{}) {
					switch actor.(type) {
					case float64:
						actorsIds = append(actorsIds, int(actor.(float64)))
					default:
						return nil, nil, fmt.Errorf(`actors identificators should be numbers`)
					}
				}
			}
		case "year":
			switch val.(type) {
			case float64:
				values[key] = fmt.Sprintf("%d", int(val.(float64)))
			default:
				return nil, nil, fmt.Errorf(`year should be a number`)
			}
		case "rating":
			switch val.(type) {
			case float64:
				values[key] = fmt.Sprintf("%f", val.(float64))
			default:
				return nil, nil, fmt.Errorf(`rating should be a floating-point number`)
			}
		case "name", "description":
			switch val.(type) {
			case []byte:
				values[key] = string(val.([]byte))
			case string:
				values[key] = val.(string)
			default:
				return nil, nil, fmt.Errorf(`name and description should be a 'text'`)
			}
		default:
			return nil, nil, fmt.Errorf(`some invalid option in request body`)
		}
	}
	return values, actorsIds, nil
}
