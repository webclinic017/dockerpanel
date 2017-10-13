package controller

import "net/http"
import (
	"github.com/mkreder/dockerpanel/templates"
	"github.com/mkreder/dockerpanel/db"
	"strings"
	"github.com/mkreder/dockerpanel/model"

	"github.com/mkreder/dockerpanel/login"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {
	userName := login.GetUserName(r)
	if userName != "" {
		webs := db.Mgr.GetAllWebs()
		templates.WriteWebTemplate(w,webs,"")
	} else {
		templates.WriteLoginTemplate(w,"")
	}

}

func AddWeb(w http.ResponseWriter, r *http.Request) {
	userName := login.GetUserName(r)
	if userName != "" {
		r.ParseForm()
		dominio := strings.Join(r.Form["dominio"],"")
		id := strings.Join(r.Form["id"],"")
		var err error

		if ( len(id) == 0 ) && ( db.Mgr.CheckIfWebExists(dominio) ){
			templates.WriteWebTemplate(w, db.Mgr.GetAllWebs(),"El sitio web " + dominio + " ya existe")
		} else {
			var web model.Web
			if (len(id) != 0) {
				web = db.Mgr.GetWeb(id)
			}
			web.Dominio = dominio

			cgi := strings.Join(r.Form["cgi"],"")
			if len(cgi) != 0 {
				web.CGI = true
			} else {
				web.CGI = false
			}

			ssl := strings.Join(r.Form["ssl"],"")
			if len(ssl) != 0 {
				web.SSL = true
			} else {
				web.SSL = false
			}

			php := strings.Join(r.Form["php"],"")
			if len(php) != 0 {
				web.PHP = true
			} else {
				web.PHP = false
			}

			phpVersion := strings.Join(r.Form["phpVersion"],"")
			web.PHPversion = phpVersion


			phpMethod := strings.Join(r.Form["phpMethod"],"")
			web.PHPmethod = phpMethod


			webserver := strings.Join(r.Form["webserver"],"")
			web.Webserver = webserver


			web.Estado = 1
			if len(id) == 0 {
				err = db.Mgr.AddWeb(&web)
				if err != nil {
					templates.WriteWebTemplate(w, db.Mgr.GetAllWebs(),"Error al agregar el sitio web")
				} else {
					templates.WriteWebTemplate(w, db.Mgr.GetAllWebs(),"")
				}
			} else {
				err = db.Mgr.UpdateWeb(&web)
				if err != nil {
					templates.WriteWebTemplate(w, db.Mgr.GetAllWebs(),"Error al actualizar el sitio web")
				} else {
					templates.WriteWebTemplate(w, db.Mgr.GetAllWebs(),"")
				}
			}
		}
	} else {
		templates.WriteLoginTemplate(w,"")
	}
}

func RemoveWeb(w http.ResponseWriter, r *http.Request) {
	userName := login.GetUserName(r)
	if userName != "" {
		id := r.URL.Query().Get("id")
		err := db.Mgr.RemoveWeb(id)
		if err != nil {
			templates.WriteWebTemplate(w,db.Mgr.GetAllWebs(),"Error al borrar el sitio web")
		} else {
			templates.WriteWebTemplate(w,db.Mgr.GetAllWebs(),"")
		}
	} else {
		templates.WriteLoginTemplate(w,"")
	}

}