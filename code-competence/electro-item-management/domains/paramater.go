package domains

import (
	"electro-item-management/constants"
	"strconv"
)

type (
	ItemParamQuery struct {
		Name string
		Pages
	}

	Pages struct {
		PageString string
		Page       int
	}
)

func (p *Pages) ToPageINT() error {
	var err error
	if p.PageString != "" {
		p.Page, err = strconv.Atoi(p.PageString)
	}
	return err
}

func (p *Pages) CalcOffsetLimit() (int, int) {
	var offset int
	if p.Page > 1 {
		offset = p.Page * constants.LIMIT
	}
	return offset, constants.LIMIT
}

func (ipq *ItemParamQuery) IsNameSet() bool {
	return ipq.Name != ""
}

func (ipq *ItemParamQuery) SetNameForDBQuery() {
	ipq.Name = "%" + ipq.Name + "%"
}
