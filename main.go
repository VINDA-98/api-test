package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type ReturnData struct {
	Members     []Member     `json:"members"`
	Departments []Department `json:"departments"`
}
type Member struct {
	Userid     string   `json:"userid"`
	Name       string   `json:"name"`
	EmployeeNo string   `json:"employee_no"`
	Department []string `json:"department"`
	Email      string   `json:"email"`
}
type Department struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentid"`
	Order    int    `json:"order"`
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/members", func(c *gin.Context) {
		departments, members := GetDepartment10000()
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"members":     members,
			"departments": departments,
		})
	})
	r.GET("/members/plus", func(c *gin.Context) {
		departments, members := GetDepartment100000()
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"members":     members,
			"departments": departments,
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	err := r.Run(":3002")
	if err != nil {
		return
	}
}

func GetDepartment10000() (departments []Department, members []Member) {

	var department1 = Department{
		ID:       "d10000001",
		Name:     "商汤科技",
		ParentId: "",
		Order:    0,
	}
	departments = append(departments, department1)

	var department2 = Department{
		ID:       "d10000002",
		Name:     "事业一部",
		ParentId: department1.ID,
		Order:    0,
	}
	departments = append(departments, department2)

	var department3 = Department{
		ID:       "d10000003",
		Name:     "事业二部",
		ParentId: department1.ID,
		Order:    1,
	}
	departments = append(departments, department3)

	var department4 = Department{
		ID:       "d10000004",
		Name:     "产品部",
		ParentId: department1.ID,
		Order:    2,
	}
	departments = append(departments, department4)

	var department5 = Department{
		ID:       "d10000005",
		Name:     "人事部",
		ParentId: department1.ID,
		Order:    2,
	}
	departments = append(departments, department5)

	//子部门
	var department6 = Department{
		ID:       "d10000006",
		Name:     "产品组",
		ParentId: department4.ID,
		Order:    0,
	}
	departments = append(departments, department6)

	var department7 = Department{
		ID:       "d10000007",
		Name:     "产品一组",
		ParentId: department6.ID,
		Order:    0,
	}
	departments = append(departments, department7)

	var department8 = Department{
		ID:       "d10000008",
		Name:     "产品二组",
		ParentId: department6.ID,
		Order:    1,
	}
	departments = append(departments, department8)

	var department9 = Department{
		ID:       "d10000009",
		Name:     "事业一组",
		ParentId: department2.ID,
		Order:    0,
	}
	departments = append(departments, department9)

	var department10 = Department{
		ID:       "d10000010",
		Name:     "事业二组",
		ParentId: department2.ID,
		Order:    0,
	}
	departments = append(departments, department10)

	//子部门
	var department11 = Department{
		ID:       "d10000011",
		Name:     "绩效组",
		ParentId: department5.ID,
		Order:    0,
	}
	departments = append(departments, department11)

	var department12 = Department{
		ID:       "d10000012",
		Name:     "绩效一组",
		ParentId: department11.ID,
		Order:    0,
	}
	departments = append(departments, department12)

	var department13 = Department{
		ID:       "d10000013",
		Name:     "绩效二组",
		ParentId: department11.ID,
		Order:    1,
	}
	departments = append(departments, department13)

	for i := 0; i < 20; i++ {
		var bianHao = strconv.Itoa(i + 1)
		if i < 10 {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "靓仔" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department7.ID},
				Email:      "liangzai" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		} else {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "靓仔" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department8.ID},
				Email:      "liangzai" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		}

	}

	for i := 20; i < 40; i++ {
		var bianHao = strconv.Itoa(i + 1)
		if i < 30 {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "表锅" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department9.ID},
				Email:      "biaoguo" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		} else {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "表锅" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department10.ID},
				Email:      "biaoguo" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		}

	}

	for i := 40; i < 60; i++ {
		var bianHao = strconv.Itoa(i + 1)
		if i < 50 {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "表妹" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department12.ID},
				Email:      "biaomei" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		} else {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "表妹" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department13.ID},
				Email:      "biaomei" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		}

	}

	for i := 1; i < 9000; i++ {
		var bianHao = strconv.Itoa(i + 1)
		if i < 50 {
			var member = Member{
				Userid:     "1" + bianHao,
				Name:       "小兵" + bianHao,
				EmployeeNo: "e100" + bianHao,
				Department: []string{department2.ID},
				Email:      "xiaobing" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		} else {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "小兵" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department1.ID},
				Email:      "xiaobing" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		}

	}

	for i := 0; i < 10; i++ {
		var bianHao = strconv.Itoa(i*1000 + 1)
		var departmentI = Department{
			ID:       "d10000011" + bianHao,
			Name:     "高校组" + bianHao,
			ParentId: department1.ID,
			Order:    i,
		}
		departments = append(departments, departmentI)

		for j := 1; j < 10; j++ {
			var bianHao = strconv.Itoa(i*1000 + 1 + j*100)
			var departmentJ = Department{
				ID:       "d10000011" + bianHao,
				Name:     "逗比组" + bianHao,
				ParentId: departmentI.ID,
				Order:    j,
			}
			departments = append(departments, departmentJ)
			for z := 1; z < 6; z++ {
				var bianHao = strconv.Itoa(i*1000 + j*100 + z*10)
				var departmentZ = Department{
					ID:       "d10000011" + bianHao,
					Name:     "沙雕组" + bianHao,
					ParentId: departmentJ.ID,
					Order:    z,
				}
				departments = append(departments, departmentZ)
				for k := 1; k < 4; k++ {
					var bianHao = strconv.Itoa(i*1000 + 1 + j*100 + z*10 + k)
					var departmentK = Department{
						ID:       "d10000011" + bianHao,
						Name:     "吃瓜组" + bianHao,
						ParentId: departmentZ.ID,
						Order:    k,
					}
					departments = append(departments, departmentK)
				}
			}

		}
	}
	return
}

func GetDepartment100000() (departments []Department, members []Member) {

	var department1 = Department{
		ID:       "d10000001",
		Name:     "商汤科技",
		ParentId: "",
		Order:    0,
	}
	departments = append(departments, department1)

	var department2 = Department{
		ID:       "d10000002",
		Name:     "事业一部",
		ParentId: department1.ID,
		Order:    0,
	}
	departments = append(departments, department2)

	var department3 = Department{
		ID:       "d10000003",
		Name:     "事业二部",
		ParentId: department1.ID,
		Order:    1,
	}
	departments = append(departments, department3)

	var department4 = Department{
		ID:       "d10000004",
		Name:     "产品部",
		ParentId: department1.ID,
		Order:    2,
	}
	departments = append(departments, department4)

	var department5 = Department{
		ID:       "d10000005",
		Name:     "人事部",
		ParentId: department1.ID,
		Order:    2,
	}
	departments = append(departments, department5)

	//子部门
	var department6 = Department{
		ID:       "d10000006",
		Name:     "产品组",
		ParentId: department4.ID,
		Order:    0,
	}
	departments = append(departments, department6)

	var department7 = Department{
		ID:       "d10000007",
		Name:     "产品一组",
		ParentId: department6.ID,
		Order:    0,
	}
	departments = append(departments, department7)

	var department8 = Department{
		ID:       "d10000008",
		Name:     "产品二组",
		ParentId: department6.ID,
		Order:    1,
	}
	departments = append(departments, department8)

	var department9 = Department{
		ID:       "d10000009",
		Name:     "事业一组",
		ParentId: department2.ID,
		Order:    0,
	}
	departments = append(departments, department9)

	var department10 = Department{
		ID:       "d10000010",
		Name:     "事业二组",
		ParentId: department2.ID,
		Order:    0,
	}
	departments = append(departments, department10)

	//子部门
	var department11 = Department{
		ID:       "d10000011",
		Name:     "绩效组",
		ParentId: department5.ID,
		Order:    0,
	}
	departments = append(departments, department11)

	var department12 = Department{
		ID:       "d10000012",
		Name:     "绩效一组",
		ParentId: department11.ID,
		Order:    0,
	}
	departments = append(departments, department12)

	var department13 = Department{
		ID:       "d10000013",
		Name:     "绩效二组",
		ParentId: department11.ID,
		Order:    1,
	}
	departments = append(departments, department13)

	for i := 0; i < 20; i++ {
		var bianHao = strconv.Itoa(i + 1)
		if i < 10 {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "靓仔" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department7.ID},
				Email:      "liangzai" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		} else {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "靓仔" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department8.ID},
				Email:      "liangzai" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		}

	}

	for i := 20; i < 40; i++ {
		var bianHao = strconv.Itoa(i + 1)
		if i < 30 {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "表锅" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department9.ID},
				Email:      "biaoguo" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		} else {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "表锅" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department10.ID},
				Email:      "biaoguo" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		}

	}

	for i := 40; i < 60; i++ {
		var bianHao = strconv.Itoa(i + 1)
		if i < 50 {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "表妹" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department12.ID},
				Email:      "biaomei" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		} else {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "表妹" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department13.ID},
				Email:      "biaomei" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		}

	}

	for i := 1; i < 90000; i++ {
		var bianHao = strconv.Itoa(i + 1)
		if i < 50 {
			var member = Member{
				Userid:     "1" + bianHao,
				Name:       "小兵" + bianHao,
				EmployeeNo: "e100" + bianHao,
				Department: []string{department2.ID},
				Email:      "xiaobing" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		} else {
			var member = Member{
				Userid:     "100" + bianHao,
				Name:       "小兵" + bianHao,
				EmployeeNo: "e100000" + bianHao,
				Department: []string{department1.ID},
				Email:      "xiaobing" + bianHao + "@ones.cn",
			}
			members = append(members, member)
		}

	}

	for i := 0; i < 20; i++ {
		var bianHao = strconv.Itoa(i*1000 + 1)
		var departmentI = Department{
			ID:       "d10000011" + bianHao,
			Name:     "高校组" + bianHao,
			ParentId: department1.ID,
			Order:    i,
		}
		departments = append(departments, departmentI)

		for j := 1; j < 5; j++ {
			var bianHao = strconv.Itoa(i*1000 + 1 + j*100)
			var departmentJ = Department{
				ID:       "d10000011" + bianHao,
				Name:     "逗比组" + bianHao,
				ParentId: departmentI.ID,
				Order:    j,
			}
			departments = append(departments, departmentJ)
			for z := 1; z < 6; z++ {
				var bianHao = strconv.Itoa(i*1000 + j*100 + z*10)
				var departmentZ = Department{
					ID:       "d10000011" + bianHao,
					Name:     "沙雕组" + bianHao,
					ParentId: departmentJ.ID,
					Order:    z,
				}
				departments = append(departments, departmentZ)
				for k := 1; k < 4; k++ {
					var bianHao = strconv.Itoa(i*1000 + 1 + j*100 + z*10 + k)
					var departmentK = Department{
						ID:       "d10000011" + bianHao,
						Name:     "吃瓜组" + bianHao,
						ParentId: departmentZ.ID,
						Order:    k,
					}
					departments = append(departments, departmentK)
				}
			}

		}
	}
	return
}
