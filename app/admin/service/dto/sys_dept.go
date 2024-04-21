package dto

// SysDeptGetPageReq 列表或者搜索使用结构体
type SysDeptGetPageReq struct {
	DeptId   int    `form:"deptId" search:"type:exact;column:dept_id;table:sys_dept" comment:"id"`       //id
	ParentId int    `form:"parentId" search:"type:exact;column:parent_id;table:sys_dept" comment:"上级部门"` //上级部门
	DeptPath string `form:"deptPath" search:"type:exact;column:dept_path;table:sys_dept" comment:""`     //路径
	DeptName string `form:"deptName" search:"type:exact;column:dept_name;table:sys_dept" comment:"部门名称"` //部门名称
	Sort     int    `form:"sort" search:"type:exact;column:sort;table:sys_dept" comment:"排序"`            //排序
	Leader   string `form:"leader" search:"type:exact;column:leader;table:sys_dept" comment:"负责人"`       //负责人
	Phone    string `form:"phone" search:"type:exact;column:phone;table:sys_dept" comment:"手机"`          //手机
	Email    string `form:"email" search:"type:exact;column:email;table:sys_dept" comment:"邮箱"`          //邮箱
	Status   string `form:"status" search:"type:exact;column:status;table:sys_dept" comment:"状态"`        //状态
}

func (m *SysDeptGetPageReq) GetNeedSearch() interface{} {
	return *m
}
