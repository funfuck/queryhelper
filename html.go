package queryhelper

import (
	"fmt"
	"html/template"
)

const (
	ALERT_SUCCESS     string = "alert-success"
	ALERT_WARNING     string = "alert-warning"
	ALERT_DANGER      string = "alert-danger"
	TPL_THEAD         string = `<thead>%s</thead>`
	TPL_TBODY         string = `<tbody>%s</tbody>`
	TPL_TR            string = `<tr>%s</tr>`
	TPL_TD            string = `<td>%s</td>`
	TPL_TH            string = `<th>%s</th>`
	TPL_TABLE_CENTER  string = `<table class="table table-striped table-bordered table-hover table-text-center">%s</table>`
	TPL_TH_NO         string = `#`
	TPL_TH_MANAGEMENT string = `การจัดการ`
	TPL_BTN_VIEW      string = `<a href="%s"><i class="fa fa-search fa-fw"></i></a>`
	TPL_BTN_EDIT      string = `<a href="%s"><i class="fa fa-pencil-square-o fa-fw"></i></a>`
	TPL_BTN_DELETE    string = `<a href="%s" title="delete" onclick="return confirm('คุณต้องการลบ, %s');"><i class="fa fa-trash fa-fw"></i></a>`
	TPL_BTN_SORT      string = `<a class="table-sort-icon sort" id="%s"><i class="fa fa-fw fa-sort"></i></a>`
)

func GenSearchForm(q []SearchField) template.HTML {
	t := `<div class="col-lg-4 text-left">
                    <label>%s</label>
					<input type="hidden" class="form-control" name="search.%d.key" value="%s">
					%s
				</div>`

	searchIcon := `<div class="form-group input-group">
						<input type="text" class="form-control" name="search.%d.value" value="%s">
						<span class="input-group-btn">
							<button class="btn btn-default" type="submit"><i class="fa fa-search"></i></button>
						</span>
					</div>`

	search := `<input type="text" class="form-control" name="search.%d.value" value="%s">`

	s := ``
	i := 0
	for k, v := range q {

		input := ``
		if len(q) - 1 == k {
			input = fmt.Sprintf(searchIcon, i, v.Value)
		} else{
			input = fmt.Sprintf(search, i, v.Value)
		}

		s += fmt.Sprintf(t, v.Label, i, v.Key, input)
		i++
	}

	return template.HTML(s)
}

func GenAlertFlashMsg(flash map[string]string) template.HTML {

	if len(flash) == 0 {
		return ""
	}

	var class, message string
	if v, ok := flash["notice"]; ok {
		class = ALERT_SUCCESS
		message = v
	} else if v, ok := flash["warning"]; ok {
		class = ALERT_WARNING
		message = v
	} else if v, ok := flash["error"]; ok {
		class = ALERT_DANGER
		message = v
	}
	t := `<div class="alert %s alert-dismissible" role="alert">
            <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <i class="fa fa-lg fa-exclamation-circle"></i> %s
        </div>`
	s := fmt.Sprintf(t, class, message)
	return template.HTML(s)
}

func GenTH(d interface{}) string {
	return fmt.Sprintf(TPL_TH, d)
}

func GenTD(d interface{}) string {
	return fmt.Sprintf(TPL_TD, d)
}

func GenTR(d string) string {
	return fmt.Sprintf(TPL_TR, d)
}

func GenBtnSort(d interface{}) string {
	return fmt.Sprintf(TPL_BTN_SORT, d)
}

func GenBtnView(d interface{}) string {
	return fmt.Sprintf(TPL_BTN_VIEW, d)
}

func GenBtnEdit(d interface{}) string {
	return fmt.Sprintf(TPL_BTN_EDIT, d)
}

func GenBtnDelete(d interface{}, e interface{}) string {
	return fmt.Sprintf(TPL_BTN_DELETE, d, e)
}

func GenTable(head string, body string) template.HTML {
	h := fmt.Sprintf(TPL_THEAD, head)
	b := fmt.Sprintf(TPL_TBODY, body)
	t := fmt.Sprintf(TPL_TABLE_CENTER, h+b)
	return template.HTML(t)
}

func GetToggleStringByString(val string, targetVal string, trueRes string, falseRes string) string {
	var returnVal string
	if val == targetVal {
		returnVal = trueRes
	} else {
		returnVal = falseRes
	}
	return returnVal
}
