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
<h3>1.5 获取诊所列表

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