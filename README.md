云诊所接口文档
===========

**创建时间：2018-08-16**

修改记录
--------
| 修定日期 | 修改内容 | 修改人 | 
| :-: | :-: | :-:  | 

接口列表
--------


1 诊所模块
--------

</br>
<h3>1.1 添加诊所

```
请求地址：/clinic/add
```
**请求包示例**

```
{
	code:1
	name:龙华诊所
	responsible_person:王大锤
	area:北京市朝阳区
	province:北京市
	city:北京市
	district:东城区
	status:
	username:lh_admin
	password:123456
	phone:13300000001
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  诊所编码| |
| name | String | ✅ | 诊所名称 | |
| responsible_person | String | ✅ |  负责人 | |
| area | String | ✅ |  诊所详细地址 | |
| province | String | ❌ |  省 | |
| city | String | ❌ |  市 | |
| district | String | ❌ |  区、县| |
| status | Boolean | ✅ |  是否启用 | |
| username | String | ✅ |  超级管理员账号 | |
| password | String | ✅ |  密码 | |
| phone | String | ✅ |  预留手机号码 | |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>1.2 获取诊所列表

```
请求地址：/clinic/list
```
**请求包示例**

```
{
	keyword:
	start_date:2018-08-01
	end_date:2018-08-08
	status:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  诊所编码/诊所名称| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| status | Boolean | ❌ | 是否启用 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "area": "测试地址来一打178号",
      "city": "深圳市",
      "clinic_id": 8,
      "code": "10004",
      "created_time": "2018-08-08T10:23:45.535638+08:00",
      "deleted_time": null,
      "district": "南山区",
      "id": 8,
      "name": "蛇口店",
      "phone": "15387556262",
      "province": "广东省",
      "responsible_person": "刘一刀",
      "status": true,
      "updated_time": "2018-08-08T10:23:45.535638+08:00",
      "username": "sk_admin"
    }
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 1
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.id | Int | ✅ |  诊所id| |
| data.items.code | String | ✅ |  诊所编码| |
| data.items.city | String | ✅ |  市| |
| data.items.clinic_id | Int | ✅ |  诊所id| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.district | String | ✅ |  区、县| |
| data.items.name | String | ✅ |  诊所名称| |
| data.items.phone | String | ✅ |  预留手机号码| |
| data.items.province | String | ✅ |  省| |
| data.items.responsible_person | String | ✅ |  诊所负责人| |
| data.items.status | Boolean | ✅ |  是否启用| |
| data.items.updated_time | String | ✅ |  更新时间| |
| data.items.username | String | ✅ |  诊所超级管理员账号| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--


</br>
<h3>1.3 启用、关闭诊所

```
请求地址：/clinic/update/status
```
**请求包示例**

```
{
	status:true
	clinic_id:8
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| status | Boolean | ✅ |  是否启用 | |
| clinic_id | Int | ✅ |  诊所id| |
**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>1.4 更新诊所

```
请求地址：/clinic/update
```
**请求包示例**

```
{
	clinic_id:8
	name:龙华诊所
	responsible_person:王大锤
	area:北京市朝阳区
	province:北京市
	city:北京市
	district:东城区
	status:
	username:lh_admin
	password:123456
	phone:13300000001
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所id| |
| name | String | ✅ | 诊所名称 | |
| responsible_person | String | ✅ |  负责人 | |
| area | String | ✅ |  诊所详细地址 | |
| province | String | ❌ |  省 | |
| city | String | ❌ |  市 | |
| district | String | ❌ |  区、县| |
| status | Boolean | ✅ |  是否启用 | |
| username | String | ✅ |  超级管理员账号 | |
| password | String | ✅ |  密码 | |
| phone | String | ✅ |  预留手机号码 | |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>1.5 获取诊所详情

```
请求地址：/clinic/getByID
```
**请求包示例**

```
{
	clinic_id:8
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "area": "测试地址来一打178号",
    "clinic_id": 8,
    "code": "10004",
    "created_time": "2018-08-08T10:23:45.535638+08:00",
    "name": "蛇口店",
    "phone": null,
    "responsible_person": "刘一刀",
    "status": true
  },
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.clinic_id | Int | ✅ |  诊所id| |
| data.code | String | ✅ |  诊所编码| |
| data.created_time | Date | ✅ |  创建时间| |
| data.name | String | ✅ |  诊所名称| |
| data.phone | String | ✅ |  预留手机号码| |
| data.responsible_person | String | ✅ |  诊所负责人| |
| data.status | Boolean | ✅ |  是否启用| |
--

</br>
<h3>1.6 获取最新的诊所编码

```
请求地址：/clinic/code
```
**请求包示例**

```
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
**应答包示例**

```
{
  "code": "200",
  "data": {
    "code": "10004"
  },
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.code | String | ✅ |  最新的诊所编码| |
| msg | String | ✅ |  返回的文本信息| |
--


3 人员模块
--------

</br>
<h3>3.1 诊所用户登录

```
请求地址：/personnel/login
```
**请求包示例**

```
{
	username:lh_admin
	password:123456
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| username | String | ✅ |  超级管理员账号 | |
| password | String | ✅ |  密码 | |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "clinic_id": 1,
    "clinic_name": "龙华诊所",
    "code": "10000",
    "id": 1,
    "is_clinic_admin": true,
    "name": "超级管理员",
    "username": "lh_admin"
  },
  "login_times": 2063,
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.clinic_id | Int | ✅ |  诊所id| |
| data.clinic_name | String | ✅ |  诊所名称| |
| data.code | String | ✅ |  诊所编码| |
| data.id | String | ✅ |  诊所编码| |
| data.is_clinic_admin | String | ✅ |  是否超级管理员| |
| data.name | String | ✅ |  登录人员名称| |
| data,username | String | ✅ |  登录账号| |
| login_times | Int | ✅ |  登录次数 | |
| msg | String | ✅ |  返回码， 200 成功| |
--

</br>
<h3>3.2 添加人员

```
请求地址：/personnel/create
```
**请求包示例**

```
{
	code:0007
	name:胡一天
	clinic_id:8
	department_id:5
	weight:13
	title:主治医生
	personnel_type:2
	username:hyt123
	password:123456
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  人员编码| |
| name | String | ✅ | 人员名称 | |
| clinic_id | Int | ✅ |  所属诊所id | |
| department_id | Int | ✅ |  所属科室id | |
| weight | Int | ❌ |  人员权重 | |
| title | String | ✅ | 人员职称 | |
| personnel_type | Int | ✅ |  关系类型 1：人事科室， 2：出诊科室 | |
| username | String | ❌ |  登录账号 | |
| password | String | ❌ |  登录密码 | |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>3.3 通过id获取人员

```
请求地址：/personnel/getById
```
**请求包示例**

```
{
	id:8
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| id | Int | ✅ |  人员id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "clinic_id": 1,
    "clinic_name": "龙华诊所",
    "department_code": "00611111",
    "department_id": 6,
    "department_name": "牙科",
    "id": 20,
    "is_appointment": true,
    "name": "胡一天",
    "status": true,
    "title": "主治医生",
    "username": "hyt123",
    "weight": 110
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.clinic_id | Int | ✅ |  诊所id| |
| data.clinic_name | String | ✅ |  诊所名称 | |
| data.department_code | String | ✅ |  所属科室编码 | |
| data.department_id | Int | ✅ |  所属科室id | |
| data.department_name | String | ✅ |  所属科室名称| |
| data.id | Int | ✅ |  人员id | |
| data.is_appointment | Boolean | ✅ |  是否开放预约/挂号 | |
| data.name | String | ✅ |  人员名称 | |
| data.status | Boolean | ✅ |  是否启用 | |
| data.title | String | ✅ |  人员职称 | |
| data.username | String | ✅ |  登录账号 | |
| data.weight | Int | ✅ |  人员权重 | |
--

</br>
<h3>3.4 获取人员列表

```
请求地址：/personnel/list
```
**请求包示例**

```
{
	clinic_id:1
	personnel_type:2
	department_id:
	offset:
	limit:
	keyword:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所id | |
| personnel_type | Int | ❌ |  关系类型 1：人事科室， 2：出诊科室 | |
| department_id | Int | ❌ | 科室id | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |
| keyword | String | ❌ |  诊所编码/诊所名称| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_id": 1,
      "clinic_name": "龙华诊所",
      "code": "0007",
      "department_code": "00611111",
      "department_id": 6,
      "department_name": "牙科",
      "id": 20,
      "is_appointment": true,
      "name": "胡一天",
      "personnel_type": 2,
      "status": true,
      "title": "主治医生",
      "username": "hyt123",
      "weight": 110
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 17
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.clinic_id | Int | ✅ |  诊所id| |
| data.items.clinic_name | String | ✅ |  诊所名称 | |
| data.items.code | String | ✅ |  人员编码 | |
| data.items.department_code | String | ✅ |  科室编码| |
| data.items.department_id | Int | ✅ |  科室id| |
| data.items.department_name | String | ✅ |  科室名称| |
| data.items.id | Int | ✅ |  人员id| |
| data.items.is_appointment | Boolean | ✅ |  是否开放预约/挂号| |
| data.items.name | String | ✅ |  人员名称| |
| data.items.personnel_type | Int | ✅ |  关系类型 1：人事科室， 2：出诊科室| |
| data.items.status | Boolean | ✅ |  是否启用| |
| data.items.title | String | ✅ | 人员职称| |
| data.items.username | String | ✅ |  人员登录账号| |
| data.items.weight | String | ✅ |  人员权重| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>3.3 通过id获取人员

```
请求地址：/personnel/getById
```
**请求包示例**

```
{
	id:8
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| id | Int | ✅ |  人员id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "clinic_id": 1,
    "clinic_name": "龙华诊所",
    "department_code": "00611111",
    "department_id": 6,
    "department_name": "牙科",
    "id": 20,
    "is_appointment": true,
    "name": "胡一天",
    "status": true,
    "title": "主治医生",
    "username": "hyt123",
    "weight": 110
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.clinic_id | Int | ✅ |  诊所id| |
| data.clinic_name | String | ✅ |  诊所名称 | |
| data.department_code | String | ✅ |  所属科室编码 | |
| data.department_id | Int | ✅ |  所属科室id | |
| data.department_name | String | ✅ |  所属科室名称| |
| data.id | Int | ✅ |  人员id | |
| data.is_appointment | Boolean | ✅ |  是否开放预约/挂号 | |
| data.name | String | ✅ |  人员名称 | |
| data.status | Boolean | ✅ |  是否启用 | |
| data.title | String | ✅ |  人员职称 | |
| data.username | String | ✅ |  登录账号 | |
| data.weight | Int | ✅ |  人员权重 | |
--

</br>
<h3>3.5 通过用户查询用户菜单

```
请求地址：/personnel/FunMenusByPersonnel
```
**请求包示例**

```
{
	id:46
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| id | Int | ✅ |  人员id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "function_menu_id": 1,
      "icon": null,
      "level": 0,
      "menu_name": "就诊流程",
      "menu_url": "/treatment",
      "parent_function_menu_id": null,
      "weight": 0
    },
    {
      "function_menu_id": 2,
      "icon": null,
      "level": 0,
      "menu_name": "诊所管理",
      "menu_url": "/clinic",
      "parent_function_menu_id": null,
      "weight": 1
    },
	...	
  ],
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.function_menu_id | Int | ✅ |  菜单功能项id| |
| data.items.icon | String | ❌ |  功能图标 | |
| data.items.level | Int | ✅ |  菜单等级 | |
| data.items.menu_name | String | ✅ |  菜单名| |
| data.items.menu_url | String | ✅ |  功能路由| |
| data.items.parent_function_menu_id | Int | ✅ |  上级菜单id| |
| data.items.weight | Int | ✅ | 菜单功能项权重 | |
--

</br>
<h3>3.6 修改人员

```
请求地址：/personnel/update
```
**请求包示例**

```
{
	personnel_id:46
	code:0007
	name:胡一天
	clinic_id:8
	department_id:5
	weight:13
	title:主治医生
	personnel_type:2
	username:hyt123
	password:123456
	is_appointment:
	status:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| personnel_id | Int | ✅ |  人员id| |
| code | String | ❌ |  人员编码| |
| name | String | ❌ | 人员名称 | |
| clinic_id | Int | ❌ |  所属诊所id | |
| department_id | Int | ❌ |  所属科室id | |
| weight | Int | ❌ |  人员权重 | |
| title | String | ❌ | 人员职称 | |
| personnel_type | Int | ✅ |  关系类型 1：人事科室， 2：出诊科室 | |
| username | String | ❌ |  登录账号 | |
| password | String | ❌ |  登录密码 | |
| is_appointment | Boolean | ❌ |  是否开放预约/挂号 | |
| status | Boolean | ❌ |  是否启用 | |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>3.7 删除人员

```
请求地址：/personnel/delete
```
**请求包示例**

```
{
	personnel_id:8
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| personnel_id | Int | ✅ |  人员id| |

**应答包示例**

```
{
  "code": "200",
  "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>3.8 有账号的人员列表（包含角色）

```
请求地址：/personnel/PersonnelWithUsername
```
**请求包示例**

```
{
	clinic_id:1
	offset:
	limit:
	keyword:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  诊所编码/诊所名称| |
| clinic_id | Int | ❌ | 诊所id| |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "department_name": null,
      "personnel_id": 1,
      "personnel_name": "超级管理员",
      "personnel_type": null,
      "role_name": "超级管理员",
      "status": true,
      "username": "lh_admin"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 20
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.department_name | String | ❌ |  所属科室名称| |
| data.items.personnel_id | Int | ✅ |  人员id| |
| data.items.personnel_name | String | ✅ |  人员名称| |
| data.items.personnel_type | Int | ❌ |  关系类型 1：人事科室， 2：出诊科室| |
| data.items.role_name | String | ❌ |  所属角色名称| |
| data.items.status | Boolean | ✅ |  是否启用| |
| data.items.username | String | ❌ |  人员登录账号| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>3.9 修改账号生效状态

```
请求地址：/clinic/UpdatePersonnelStatus
```
**请求包示例**

```
{
	status:true
	personnel_id:8
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| status | Boolean | ✅ |  是否启用 | |
| personnel_id | Int | ✅ |  人员id| |
**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>3.10 修改用户名密码

```
请求地址：/clinic/UpdatePersonnelUsername
```
**请求包示例**

```
{
	username:test_admin
	password:111111
	personnel_id:8
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| username | String | ✅ |  账号| |
| password | String | ✅ |  密码| |
| personnel_id | Int | ✅ |  人员id| |
**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>3.11 通过用户查询用户角色

```
请求地址：/personnel/PersonnelRoles
```
**请求包示例**

```
{
	id:46
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| id | Int | ✅ |  人员id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-07-26T21:46:47.215207+08:00",
      "name": "检查",
      "role_id": 40,
      "status": true
    }
  ],
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.name | String | ✅ |  角色名称 | |
| data.items.role_id | Int | ✅ | 角色id | |
| data.items.status | Boolean | ✅ |  是否启用| |
--

</br>
<h3>3.12 获取医生所属科室

```
请求地址：/personnel/PersonnelDepartmentList
```
**请求包示例**

```
{
	personnel_id:46
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| personnel_id | Int | ✅ |  人员id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "department_id": 23,
      "department_name": "医技科室"
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.department_id | Date | ✅ |  科室id| |
| data.items.department_name | String | ✅ |  科室名称 | |
--

5 就诊人模块
--------

</br>
<h3>5.1 新增就诊人

```
请求地址：/patient/create
```
**请求包示例**

```
{
	cert_no:360822199312307090
	name:元洪果
	birthday:1993-12-30
	sex:1
	phone:13211223344
	address:
	profession:
	remark:
	patient_channel_id:1
	clinic_id:1
	personnel_id:46
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| cert_no | String | ✅ |  身份证号| |
| name | String | ✅ | 就诊人姓名| |
| birthday | String | ✅ |  生日 | |
| sex | Int | ✅ |  性别 0：女，1：男 | |
| phone | String | ✅ |  手机号 | |
| address | String | ❌ |  详细地址 | |
| profession | String | ❌ |  职业| |
| remark | String | ❌ |  备注 | |
| patient_channel_id | Int | ✅ |  就诊人来源 | |
| clinic_id | Int | ✅ |  诊所id | |
| personnel_id | Int | ✅ |  录入人员id | |

**应答包示例**

```
{
    "code": "200",
    "data": 27
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Int | ✅ |  返回诊所就诊人id | |
--

</br>
<h3>5.2 就诊人列表

```
请求地址：/patient/list
```
**请求包示例**

```
{
	clinic_id:1
	keyword:
	startDate:
	endDate:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所id| |
| keyword | String | ❌ | 就诊人姓名、身份证号、就诊人手机号 | |
| startDate | String | ❌ | 创建开始日期| |
| endDate | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "address": null,
      "birthday": "19931230",
      "cert_no": "360822199312307090",
      "channel_name": "社区患者",
      "city": "九江市",
      "created_time": "2018-08-05T21:46:02.900841+08:00",
      "deleted_time": null,
      "district": "瑞昌市",
      "id": 31,
      "image": null,
      "name": "元洪果",
      "patient_channel_id": 4,
      "phone": "13211223344",
      "profession": null,
      "province": "江西省",
      "remark": null,
      "sex": 1,
      "status": true,
      "updated_time": "2018-08-05T21:46:41.53629+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 27
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.address | String | ❌ |  详细地址| |
| data.items.birthday | String | ❌ |  生日| |
| data.items.cert_no | String | ✅ |  身份证号| |
| data.items.channel_name | String | ✅ |  患者来源| |
| data.items.city | String | ❌ |  市| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.district | String | ❌ |  区| |
| data.items.id | Int | ✅ |  就诊人id| |
| data.items.image | String | ❌ |  头像 | |
| data.items.name | String | ✅ |  就诊人名称| |
| data.items.patient_channel_id | Int | ✅ |  患者来源id| |
| data.items.phone | String | ✅ |  患者手机号| |
| data.items.profession | String | ❌ |  职业| |
| data.items.province | String | ❌ |  省| |
| data.items.remark | String | ❌ |  备注| |
| data.items.sex | Int | ❌ |  性别 0：女，1：男| |
| data.items.status | Boolean | ✅ |  是否启用| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.3 通过id就诊人详情

```
请求地址：/patient/getById
```
**请求包示例**

```
{
	id:2
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| id | Int | ✅ |  就诊人id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "address": "哈哈哈哈",
    "birthday": "20001010",
    "cert_no": null,
    "city": "北京市",
    "created_time": "2018-05-28T00:26:29.012104+08:00",
    "deleted_time": null,
    "district": "东城区",
    "id": 2,
    "image": null,
    "name": "大乔",
    "patient_channel_id": 1,
    "phone": "15387556262",
    "profession": null,
    "province": "北京市",
    "remark": null,
    "sex": 0,
    "status": true,
    "updated_time": "2018-05-28T00:26:29.012104+08:00"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.address | String | ❌ |  详细地址| |
| data.birthday | String | ❌ |  生日| |
| data.cert_no | String | ✅ |  身份证号| |
| data.city | String | ❌ |  市| |
| data.created_time | Date | ✅ |  创建时间| |
| data.deleted_time | Date | ❌ |  删除时间| |
| data.district | String | ❌ |  区| |
| data.id | Int | ✅ |  就诊人id| |
| data.image | String | ❌ |  头像 | |
| data.name | String | ✅ |  就诊人名称| |
| data.patient_channel_id | Int | ✅ |  患者来源id| |
| data.phone | String | ✅ |  患者手机号| |
| data.profession | String | ❌ |  职业| |
| data.province | String | ❌ |  省| |
| data.remark | String | ❌ |  备注| |
| data.sex | Int | ❌ |  性别 0：女，1：男| |
| data.status | Boolean | ✅ |  是否启用| |
| data.updated_time | Date | ✅ |  更新时间| |
--

</br>
<h3>5.4 修改就诊人

```
请求地址：/patient/update
```
**请求包示例**

```
{
	id:31
	cert_no:360822199312307090
	name:元洪果
	birthday:1993-12-30
	sex:1
	phone:13211223344
	address:
	profession:
	remark:
	patient_channel_id:1
	clinic_id:1
	personnel_id:46
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| id | Int | ✅ |  就诊人id| |
| cert_no | String | ✅ |  身份证号| |
| name | String | ✅ | 就诊人姓名| |
| birthday | String | ✅ |  生日 | |
| sex | Int | ✅ |  性别 0：女，1：男 | |
| phone | String | ✅ |  手机号 | |
| address | String | ❌ |  详细地址 | |
| profession | String | ❌ |  职业| |
| remark | String | ❌ |  备注 | |
| patient_channel_id | Int | ✅ |  就诊人来源 | |
| clinic_id | Int | ✅ |  诊所id | |
| personnel_id | Int | ✅ |  录入人员id | |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.5 通过身份号查就诊人

```
请求地址：/patient/getByCertNo
```
**请求包示例**

```
{
	cert_no:360822199312307090
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| cert_no | String | ✅ |  就诊人身份证 | |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "address": null,
    "birthday": "19931230",
    "cert_no": "360822199312307090",
    "city": "九江市",
    "created_time": "2018-08-05T21:46:02.900841+08:00",
    "deleted_time": null,
    "district": "瑞昌市",
    "id": 31,
    "image": null,
    "name": "元洪果",
    "patient_channel_id": 4,
    "phone": "13211223344",
    "profession": null,
    "province": "江西省",
    "remark": null,
    "sex": 1,
    "status": true,
    "updated_time": "2018-08-05T21:46:41.53629+08:00"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.address | String | ❌ |  详细地址| |
| data.birthday | String | ❌ |  生日| |
| data.cert_no | String | ✅ |  身份证号| |
| data.city | String | ❌ |  市| |
| data.created_time | Date | ✅ |  创建时间| |
| data.deleted_time | Date | ❌ |  删除时间| |
| data.district | String | ❌ |  区| |
| data.id | Int | ✅ |  就诊人id| |
| data.image | String | ❌ |  头像 | |
| data.name | String | ✅ |  就诊人名称| |
| data.patient_channel_id | Int | ✅ |  患者来源id| |
| data.phone | String | ✅ |  患者手机号| |
| data.profession | String | ❌ |  职业| |
| data.province | String | ❌ |  省| |
| data.remark | String | ❌ |  备注| |
| data.sex | Int | ❌ |  性别 0：女，1：男| |
| data.status | Boolean | ✅ |  是否启用| |
| data.updated_time | Date | ✅ |  更新时间| |
--

</br>
<h3>5.6 通过关键字搜索就诊人

```
请求地址：/patient/getByKeyword
```
**请求包示例**

```
{
	keyword:元
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ✅ |  就诊人姓名、身份证、手机号 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "address": null,
      "birthday": "19931230",
      "cert_no": "360822199312307090",
      "city": "九江市",
      "created_time": "2018-08-05T21:46:02.900841+08:00",
      "deleted_time": null,
      "district": "瑞昌市",
      "id": 31,
      "image": null,
      "name": "元洪果",
      "patient_channel_id": 4,
      "phone": "13211223344",
      "profession": null,
      "province": "江西省",
      "remark": null,
      "sex": 1,
      "status": true,
      "updated_time": "2018-08-05T21:46:41.53629+08:00"
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.address | String | ❌ |  详细地址| |
| data.items.birthday | String | ❌ |  生日| |
| data.items.cert_no | String | ✅ |  身份证号| |
| data.items.city | String | ❌ |  市| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.district | String | ❌ |  区| |
| data.items.id | Int | ✅ |  就诊人id| |
| data.items.image | String | ❌ |  头像 | |
| data.items.name | String | ✅ |  就诊人名称| |
| data.items.patient_channel_id | Int | ✅ |  患者来源id| |
| data.items.phone | String | ✅ |  患者手机号| |
| data.items.profession | String | ❌ |  职业| |
| data.items.province | String | ❌ |  省| |
| data.items.remark | String | ❌ |  备注| |
| data.items.sex | Int | ✅ |  性别 0：女，1：男| |
| data.items.status | Boolean | ✅ |  是否启用| |
| data.items.updated_time | Date | ✅ |  更新时间| |
--

</br>
<h3>5.7 会员，就诊人列表

```
请求地址：/patient/MemberPateintList
```
**请求包示例**

```
{
	keyword:
	startDate:
	endDate:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ | 就诊人姓名、身份证号、就诊人手机号 | |
| startDate | String | ❌ | 创建开始日期| |
| endDate | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19931230",
      "created_time": "2018-08-05T21:46:02.900841+08:00",
      "id": 31,
      "name": "元洪果",
      "phone": "13211223344",
      "sex": 1,
      "visited_time": "2018-08-12T22:57:33.24495+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 27
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.birthday | String | ❌ |  生日| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.id | Int | ✅ |  就诊人id| |
| data.items.name | String | ✅ |  就诊人名称| |
| data.items.phone | String | ✅ |  患者手机号| |
| data.items.sex | Int | ✅ |  性别 0：女，1：男| |
| data.items.visited_time | Date | ✅ |  最近就诊时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.8 完善患者个人诊前病历

```
请求地址：/patient/PersonalMedicalRecordUpsert
```
**请求包示例**

```
{
	patient_id:29
	has_allergic_history:true
	allergic_history:鸡蛋、西红柿、海鲜、money
	allergic_reaction:皮肤瘙痒、红肿
	personal_medical_history:生长于香港。文盲。否认外地长期居住史。无疫区、疫水接触史。
	family_medical_history:家中无遗传病病史。
	immunizations:天花
	menarche_age: 13
	menstrual_period_start_day:每月22日
	menstrual_period_end_day:每月28日
	menstrual_cycle_start_day:22
	menstrual_cycle_end_day:28
	menstrual_last_day:20180712
	gestational_weeks: 0
	childbearing_history:0
	remark:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| has_allergic_history | Boolean | ❌  |  是否有过敏| |
| allergic_history | String | ❌  | 过敏史| |
| allergic_reaction | String | ❌  |  过敏反应 | |
| personal_medical_history | String | ❌  |  个人病史 | |
| family_medical_history | String | ❌  |  家族病史 | |
| immunizations | String | ❌ |  接种疫苗| |
| menarche_age | Int | ❌ |  月经初潮年龄| |
| menstrual_period_start_day | String | ❌ |  月经经期开始时间 | |
| menstrual_period_end_day | String | ❌  |  月经经期结束时间 | |
| menstrual_cycle_start_day | String | ❌  |  月经周期结束时间 | |
| menstrual_cycle_end_day | String | ❌  |  月经周期结束时间 | |
| menstrual_last_day | String | ❌  |  末次月经时间 | |
| gestational_weeks | Int | ❌  |  孕周 | |
| childbearing_history | String | ❌  |  生育史 | |
| remark | String | ❌ |  备注 | |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.9 患者个人诊前病历

```
请求地址：/patient/PersonalMedicalRecord
```
**请求包示例**

```
{
	patient_id:29
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "allergic_history": "鸡蛋、西红柿、海鲜、money",
    "allergic_reaction": "皮肤瘙痒、红肿",
    "childbearing_history": "0",
    "created_time": "2018-08-01T09:47:52.688193+08:00",
    "deleted_time": null,
    "family_medical_history": "家中无遗传病病史。",
    "gestational_weeks": 0,
    "has_allergic_history": true,
    "id": 54,
    "immunizations": "天花",
    "menarche_age": 13,
    "menstrual_cycle_end_day": "每月28日",
    "menstrual_cycle_start_day": "每月22日",
    "menstrual_last_day": "20180712",
    "menstrual_period_end_day": "28",
    "menstrual_period_start_day": "22",
    "patient_id": 29,
    "personal_medical_history": "生长于香港。文盲。否认外地长期居住史。无疫区、疫水接触史。否认工业毒物、粉尘及放射性物质接触史。否认牧区、矿山、高氟区、低碘区居住史。平日生活规律，否认吸毒史。否认吸烟嗜好。否认饮酒嗜好。否认冶游史。第N+1次\n",
    "remark": null,
    "updated_time": "2018-08-01T09:47:52.688193+08:00"
  },
  "msg": "查询成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.patient_id | Int | ✅ |  就诊人id| |
| data.id | Int | ✅ |  病历id| |
| data.has_allergic_history | Boolean | ❌  |  是否有过敏| |
| data.allergic_history | String | ❌  | 过敏史| |
| data.allergic_reaction | String | ❌  |  过敏反应 | |
| data.personal_medical_history | String | ❌  |  个人病史 | |
| data.family_medical_history | String | ❌  |  家族病史 | |
| data.immunizations | String | ❌ |  接种疫苗| |
| data.menarche_age | Int | ❌ |  月经初潮年龄| |
| data.menstrual_period_start_day | String | ❌ |  月经经期开始时间 | |
| data.menstrual_period_end_day | String | ❌  |  月经经期结束时间 | |
| data.menstrual_cycle_start_day | String | ❌  |  月经周期结束时间 | |
| data.menstrual_cycle_end_day | String | ❌  |  月经周期结束时间 | |
| data.menstrual_last_day | String | ❌  |  末次月经时间 | |
| data.gestational_weeks | Int | ❌  |  孕周 | |
| data.childbearing_history | String | ❌  |  生育史 | |
| data.remark | String | ❌ |  备注 | |
--

</br>
<h3>5.10 获取最后一次体征信息

```
请求地址：/patient/GetLastBodySign
```
**请求包示例**

```
{
	patient_id:29
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "blood_type": "A",
    "bmi": 14.36,
    "breathe": 72,
    "concentration_after_breakfast": null,
    "concentration_after_dinner": null,
    "concentration_after_lunch": null,
    "concentration_before_breakfast": null,
    "concentration_before_dawn": null,
    "concentration_before_dinner": null,
    "concentration_before_lunch": null,
    "concentration_before_retiring": null,
    "concentration_empty_stomach": null,
    "created_time": null,
    "deleted_time": null,
    "diastolic_blood_pressure": 139,
    "height": 171,
    "id": null,
    "left_vision": "5.2",
    "oxygen_saturation": 20,
    "patient_id": null,
    "pulse": 78,
    "record_time": null,
    "remark": null,
    "rh_blood_type": null,
    "right_vision": "5.2",
    "systolic_blood_pressure": 110,
    "temperature": 37.6,
    "temperature_type": 1,
    "updated_time": null,
    "weight": 42
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.blood_type | String | ❌ |  血型 uc: 未查| |
| data.bmi | FLOAT | ❌ |  体重（千克）/（身高（米）*身高（米））| |
| data.breathe | Int | ❌  |  呼吸(次/分钟)| |
| data.concentration_after_breakfast | FLOAT | ❌  | 早餐后血糖浓度(mmol/I)| |
| data.concentration_after_dinner | FLOAT | ❌  |  晚餐后血糖浓度(mmol/I) | |
| data.concentration_after_lunch | FLOAT | ❌  |  午餐后血糖浓度(mmol/I) | |
| data.concentration_before_breakfast | FLOAT | ❌  |  早餐前血糖浓度(mmol/I) | |
| data.concentration_before_dawn | FLOAT | ❌ |  凌晨血糖浓度(mmol/I)| |
| data.concentration_before_dinner | FLOAT | ❌ |  晚餐前血糖浓度(mmol/I)| |
| data.concentration_before_lunch | FLOAT | ❌ |  午餐血糖浓度(mmol/I) | |
| data.concentration_before_retiring | FLOAT | ❌  |  睡前血糖浓度(mmol/I) | |
| data.concentration_empty_stomach | FLOAT | ❌  |  空腹血糖浓度(mmol/I) | |
| data.diastolic_blood_pressure | Int | ❌  |  血压舒张压 | |
| data.systolic_blood_pressure | Int | ❌  |  血压收缩压 | |
| data.height | FLOAT | ❌  |  升高（m） | |
| data.weight | FLOAT | ❌  |  体重(kg) | |
| data.left_vision | String | ❌ |  左眼视力 | |
| data.right_vision | String | ❌ |  右眼视力 | |
| data.oxygen_saturation | FLOAT | ❌ |  氧饱和度(%) | |
| data.pulse | Int | ❌ |  脉搏(次/分钟) | |
| data.rh_blood_type | Int | ❌ |  RH血型 -1: 阴性，1阳性, 0: 未查 | |
| data.temperature | Int | ❌ | 温度 | |
| data.temperature_type | Int | ❌ |  类型 1: 口温，2：耳温，3：额温，4：腋温，5：肛温 | |
--

</br>
<h3>5.11 修改身高

```
请求地址：/patient/UpsertPatientHeight
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","height":"1.7","upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.height | String | ✅  | 升高（m）| |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.12 修改体重

```
请求地址：/patient/UpsertPatientWeight
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","weight":"65","upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.weight | String | ✅  | 体重(kg)| |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.13 修改BMI

```
请求地址：/patient/UpsertPatientBmi
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","bmi":"15.79","upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.bmi | String | ✅  | 体重（千克）/（身高（米）*身高（米））| |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.14 修改血型

```
请求地址：/patient/UpsertPatientBloodType
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","blood_type":"AB",
	"upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.blood_type | String | ✅  | 血型 uc: 未查| |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.15 修改RH血型

```
请求地址：/patient/UpsertPatientRhBloodType
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","rh_blood_type":"1",
	"upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.rh_blood_type | String | ✅  | RH血型 -1: 阴性，1阳性, 0: 未查| |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.16 修改体温

```
请求地址：/patient/UpsertPatientTemperature
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","temperature_type":"1",
	"temperature":"37.1","upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.temperature_type | String | ✅  | 类型 1: 口温，2：耳温，3：额温，4：腋温，5：肛温| |
| items.item.temperature | String | ✅  | 体温 | |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.17 修改呼吸

```
请求地址：/patient/UpsertPatientBreathe
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","breathe":"71","upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.breathe | String | ✅  | 呼吸(次/分钟) | |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.18 修改脉搏

```
请求地址：/patient/UpsertPatientPulse
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","pulse":"71","upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.pulse | String | ✅  | 脉搏(次/分钟) | |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.19 修改血压

```
请求地址：/patient/UpsertPatientBloodPressure
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","systolic_blood_pressure":"110",
	"diastolic_blood_pressure":"139","upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.systolic_blood_pressure | String | ✅  | 血压收缩压 | |
| items.item.diastolic_blood_pressure | String | ✅  | 血压舒张压 | |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.20 修改视力

```
请求地址：/patient/UpsertPatientVision
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01","left_vision":"1.5",
	"right_vision":"1.7","upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.left_vision | String | ✅  | 左眼视力 | |
| items.item.right_vision | String | ✅  | 右眼视力 | |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.21 修改血糖

```
请求地址：/patient/UpsertPatientBloodSugar
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01",
	"concentration_before_retiring":"15",
	"concentration_after_dinner":"10",
	"concentration_before_dinner":"11",
	"concentration_after_lunch":"10",
	"concentration_before_lunch":"17",
	"concentration_after_breakfast":"16",
	"concentration_before_breakfast":"17",
	"concentration_before_dawn":"17",
	"concentration_empty_stomach":"17",
	"upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.concentration_before_retiring | String | ✅  | 睡前血糖浓度(mmol/I) | |
| items.item.concentration_after_dinner | String | ✅  | 晚餐后血糖浓度(mmol/I) | |
| items.item.concentration_before_dinner | String | ✅  | 晚餐前血糖浓度(mmol/I) | |
| items.item.concentration_after_lunch | String | ✅  | 午餐后血糖浓度(mmol/I) | |
| items.item.concentration_before_lunch | String | ✅  | 午餐前血糖浓度(mmol/I) | |
| items.item.concentration_after_breakfast | String | ✅  | 早餐后血糖浓度(mmol/I) | |
| items.item.concentration_before_breakfast | String | ✅  | 早餐前血糖浓度(mmol/I) | |
| items.item.concentration_before_dawn| String | ✅  | 凌晨血糖浓度(mmol/I) | |
| items.item.concentration_empty_stomach | String | ✅  | 空腹血糖浓度(mmol/I) | |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.22 修改氧饱和度

```
请求地址：/patient/upsertPatientOxygenSaturation
```
**请求包示例**

```
{
	patient_id:29
	items:	[
	{"record_time":"2018-08-01",
	"oxygen_saturation":"15",
	"upsert_type":"insert"}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  就诊人id| |
| items | Array | ✅  |  升高记录| |
| items.item.record_time | String | ✅  | 记录时间| |
| items.item.oxygen_saturation | String | ✅  | 氧饱和度(%) | |
| items.item.upsert_type | String | ✅  | 更新类型:update，insert，delete| |

**应答包示例**

```
{
    "code": "200",
    "data": null
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
--

</br>
<h3>5.23 患者身高记录

```
请求地址：/patient/PatientHeightList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-07-29T20:58:13.118975+08:00",
      "deleted_time": null,
      "height": 1.6,
      "id": 22,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.118975+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.height | FLOAT | ✅ |  身高| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.24 患者体重记录

```
请求地址：/patient/PatientWeightList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-07-29T20:58:13.200503+08:00",
      "deleted_time": null,
      "id": 20,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.200503+08:00",
      "weight": 56
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.weight | FLOAT | ✅ |  体重| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.25 患者BMI记录

```
请求地址：/patient/PatientBmiList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "bmi": 500,
      "created_time": "2018-07-29T20:58:13.280932+08:00",
      "deleted_time": null,
      "id": 12,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.280932+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.bmi | FLOAT | ✅ |  体重（千克）/（身高（米）*身高（米））| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.26 患者血型记录

```
请求地址：/patient/PatientBloodTypeList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "blood_type": "UC",
      "created_time": "2018-07-29T20:58:13.280932+08:00",
      "deleted_time": null,
      "id": 12,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.280932+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.blood_type | String | ✅ | 血型 uc: 未查| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.27 患者RH血型记录

```
请求地址：/patient/PatientRhBloodTypeList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "rh_blood_type": -1,
      "created_time": "2018-07-29T20:58:13.280932+08:00",
      "deleted_time": null,
      "id": 12,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.280932+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.rh_blood_type | String | ✅ | RH血型 -1: 阴性，1阳性, 0: 未查| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.27 患者RH血型记录

```
请求地址：/patient/PatientRhBloodTypeList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "rh_blood_type": -1,
      "created_time": "2018-07-29T20:58:13.280932+08:00",
      "deleted_time": null,
      "id": 12,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.280932+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.rh_blood_type | String | ✅ | RH血型 -1: 阴性，1阳性, 0: 未查| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.28 患者体温记录

```
请求地址：/patient/PatientTemperatureList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "temperature": 30,
      "temperature_type": 2,
      "created_time": "2018-07-29T20:58:13.280932+08:00",
      "deleted_time": null,
      "id": 12,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.280932+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.temperature | FLOAT | ✅ | 体温| |
| data.items.temperature_type | Int | ✅ | 体温| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.29 患者呼吸记录

```
请求地址：/patient/PatientBreatheList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "breathe": 30,
      "created_time": "2018-07-29T20:58:13.280932+08:00",
      "deleted_time": null,
      "id": 12,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.280932+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.breathe | Int | ✅ | 体温| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.30 患者脉搏记录

```
请求地址：/patient/PatientPulseList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "pulse": 31,
      "created_time": "2018-07-29T20:58:13.280932+08:00",
      "deleted_time": null,
      "id": 12,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.280932+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.pulse | Int | ✅ | 体温| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.31 患者血压记录

```
请求地址：/patient/PatientBloodPressureList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-07-29T20:58:13.753903+08:00",
      "deleted_time": null,
      "diastolic_blood_pressure": 110,
      "id": 11,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "systolic_blood_pressure": 70,
      "updated_time": "2018-07-29T20:58:13.753903+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.diastolic_blood_pressure | Int | ✅ | 血压舒张压| |
| data.items.systolic_blood_pressure | Int | ✅ | 血压收缩压| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.32 患者视力记录

```
请求地址：/patient/PatientVisionList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-07-29T20:58:13.829268+08:00",
      "deleted_time": null,
      "id": 10,
      "left_vision": "1.5",
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "right_vision": "1.5",
      "updated_time": "2018-07-29T20:58:13.829268+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.left_vision | String | ✅ | 左眼视力| |
| data.items.right_vision | String | ✅ | 右眼视力| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.33 患者氧饱和度记录

```
请求地址：/patient/PatientOxygenSaturationList
```
**请求包示例**

```
{
	patient_id:3
	start_date:
	end_date:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | Int | ✅ |  患者id| |
| start_date | String | ❌ | 创建开始日期| |
| end_date | String | ❌ |  创建结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-07-29T20:58:13.829268+08:00",
      "deleted_time": null,
      "id": 10,
      "oxygen_saturation": 34,
      "patient_id": 3,
      "record_time": "2018-07-29",
      "remark": null,
      "updated_time": "2018-07-29T20:58:13.829268+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 6
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.oxygen_saturation | FLOAT | ✅ | 氧饱和度(%)| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.id | Int | ✅ |  记录id| |
| data.items.record_time | Date | ✅ |  记录时间| |
| data.items.remark | String | ❌ |  备注| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.deleted_time | Date | ❌ |  删除时间| |
| data.items.updated_time | Date | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>5.34 患者统计 按性别

```
请求地址：/patient/PatientCountBySex
```
**请求包示例**

```
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "sex": 0,
      "total": 11
    },
    {
      "sex": 1,
      "total": 16
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ✅  |  返回信息 | |
| data.items.sex | Int | ✅ | 性别 0：女，1：男| |
| data.items.total | Int | ✅ |  总数| |
--

</br>
<h3>5.35 患者统计 按年龄

```
请求地址：/patient/PatientCountByAge
```
**请求包示例**

```
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "age": "[0 - 10) ",
      "total": 6
    },
    {
      "age": "[10 - 20) ",
      "total": 11
    },
    {
      "age": "[20 - 30) ",
      "total": 5
    },
    {
      "age": "[30 - 40) ",
      "total": 4
    },
    {
      "age": "[40 - 50) ",
      "total": 1
    },
    {
      "age": "[50 - 60) ",
      "total": 0
    },
    {
      "age": "[60 - 70) ",
      "total": 0
    },
    {
      "age": "[70 - 80) ",
      "total": 0
    },
    {
      "age": "[80-) ",
      "total": 0
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ✅  |  返回信息 | |
| data.items.age | Int | ✅ | 年龄段| |
| data.items.total | Int | ✅ |  总数| |
--

</br>
<h3>5.36 患者统计 按渠道

```
请求地址：/patient/PatientCountByChannel
```
**请求包示例**

```
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "patient_channel_name": "网络宣传",
      "total": 2
    },
    {
      "patient_channel_name": "会员介绍",
      "total": 3
    },
    {
      "patient_channel_name": "运营推荐",
      "total": 4
    },
    {
      "patient_channel_name": "未知",
      "total": 16
    },
    {
      "patient_channel_name": "社区患者",
      "total": 2
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ✅  |  返回信息 | |
| data.items.patient_channel_name | String | ✅ | 渠道| |
| data.items.total | Int | ✅ |  总数| |
--