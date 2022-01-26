package api

import (
	"encoding/json"
	"fmt"
	"github.com/ikbalkazanc/go-todo-rest-api/pkg/models"
	"net/http"
	"reflect"
	"strings"
)

type AlgorithmAPI struct {
}

func NewAlgorithmAPI() AlgorithmAPI {
	return AlgorithmAPI{}
}

func (p AlgorithmAPI) FirstAlgoritm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Query()
		key := p.Get("key")
		word := p.Get("word")
		word = strings.ReplaceAll(word, " ", "")
		result := ""

		for i := 0; i < len(word)-1; i++ {
			if word[i] == key[0] {
				result = result + string(word[i+1])
			}
		}

		RespondWithJSON(w, http.StatusOK, result)
	}
}

func (p AlgorithmAPI) SecondAlgoritm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model models.SecondAlgoritmModel
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&model); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		defer r.Body.Close()

		length := len(model.Arr)
		fmt.Println(length)
		for i := 0; i < length; i++ {
			if reflect.TypeOf(model.Arr[i]).Name() == "float64" {
				if fmt.Sprintf("%v", model.Arr[i]) == "0" {
					model.Arr = append(model.Arr[:i], model.Arr[i+1:]...)
					model.Arr = append(model.Arr, 0)
					i--
					length--
				}
			}
		}

		RespondWithJSON(w, http.StatusOK, model)
	}
}
