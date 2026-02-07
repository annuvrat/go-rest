package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/annuvrat/go-rest/internal/types"
	"github.com/annuvrat/go-rest/internal/utils/response"
)



func New()http.HandlerFunc{
	return  func(w http.ResponseWriter,r *http.Request){
     
		var student types.Student

		err:=json.NewDecoder(r.Body).Decode(&student)
        if errors.Is(err,io.EOF){
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return 
		}

        // if err!= nil{

		// }
{}
		response.WriteJson(w,http.StatusCreated,map[string]string{"success":"ok"})
	// w.Write([]byte("welcome to students api"))
}
}
