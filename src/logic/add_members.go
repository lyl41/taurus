package logic

import (
	"fmt"
	"taurus/src/model/dao"
	"taurus/src/model/dao/table"
	"time"
)

func AddMembers(datas []*table.Members) (successCnt int, err error) {
	for _, data := range datas {
		if data == nil {
			continue
		}
		data.CreatedAt = time.Time{}
		data.UpdatedAt = time.Time{}
		data.CreatedUser = ""
		data.ID = 0
		if e := dao.InsertMember(data); e != nil {
			fmt.Println(data, e)
			continue
		}
		successCnt++
	}
	return
}
