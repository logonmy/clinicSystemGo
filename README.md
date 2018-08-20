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

2 科室模块
--------

</br>
<h3>2.1 添加科室

```
请求地址：/department/create
```
**请求包示例**

```
{
	code:1,
	name:骨科,
	clinic_id:1,
	weight:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  科室编码| |
| name | String | ✅ |  科室名称 | |
| clinic_id | Number | ✅ |  诊所id | |
| weight | Number | ✅ |  权重 | |

**应答包示例**

```
{
    "code": "200",
    "data": "1"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | String | ❌ |  科室id | |
--


</br>
<h3>2.2 诊所科室列表

```
请求地址：/department/list
```
**请求包示例**

```
{
	clinic_id:1,
	offset: 0,
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字| |
| clinic_id | int | ✅ |  诊所id | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |

**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "clinic_id": 1,
            "code": "004",
            "created_time": "2018-06-02T15:28:30.443125+08:00",
            "deleted_time": null,
            "id": 4,
            "is_appointment": true,
            "name": "普通外科",
            "status": true,
            "updated_time": "2018-06-04T01:13:17.811669+08:00",
            "weight": 1
        }
    ],
    "page_info": {
        "limit": "1",
        "offset": "10",
        "total": 14
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ✅ |  返回信息 | |
| data | Array | ✅ |   | |
| data. clinic_id | int | ✅ |  诊所id | |
| data. code | string | ✅ | 科室编码 | |
| data. created_time | time | ✅ |  创建时间| |
| data. deleted_time | time | ✅ |  删除时间 | |
| data. id | int | ✅ |  科室id | |
| data. is_appointment | boolean | ✅ |  是否开放预约挂号 | |
| data. name | string | ✅ |  科室名称 | |
| data. status | bolean | ✅ |  是否启用 | |
| data. updated_time | time | ✅ |  修改时间| |
| data. weight | int | ✅ |  权重 | |

--



</br>
<h3>2.3 科室删除

```
请求地址：/department/delete
```
**请求包示例**

```
{
	department_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| department_id | int | ✅ |  科室id| |

**应答包示例**

```
{
    "code": "200",
    "msg": "成功"。
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ✅ |  返回信息 | |
| data | obj | ✅ |   | |

--



</br>
<h3>2.4 科室修改

```
请求地址：/department/update
```
**请求包示例**

```
{
	department_id:1,
	code: "001",
	name: "骨科",
	clinic_id: 1,
	weight: 1,
	is_appointment: true
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| department_id | int | ✅ |  科室id| |
| code | int | ❌ |  科室编码| |
| name | int | ❌ |  科室名称| |
| clinic_id | int | ❌ |  诊所id| |
| weight | int | ❌ |  权重| |
| is_appointment | int | ❌ |  是否开放预约挂号| |

**应答包示例**

```
{
    "code": "200",
    "msg": "成功"。
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ✅ |  返回信息 | |
| data | obj | ✅ |   | |

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
| data.username | String | ✅ |  登录账号| |
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

4 医生排班模块
--------

</br>
<h3>4.1 获取号源列表

```
请求地址：/doctorVisitSchedule/list
```
**请求包示例**

```
{
	personnel_id:1,
	department_id:1,
	clinic_id:1,
	start_date: "2018-05-01",
	end_date: "2018-05-08",
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Number | ✅ |  诊所id | |
| personnel_id | Number | ❌ |  医生id | |
| department_id | Number | ❌ |  科室id | |
| start_date | Number | ✅ |  开始日期 | |
| end_date | Number | ✅ |  结束日期 | |

**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "am_pm": "a",
            "department_id": 1,
            "department_name": "骨科",
            "id": 1850,
            "left_num": 20,
            "personnel_id": 2,
            "personnel_name": "扁鹊",
            "tatal_num": 20,
            "visit_date": "2018-08-20T00:00:00Z"
        },
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | Array | ❌ |   | |
| data. am_pm | string | ✅ |  上午 a, 下午p | |
| data. department_id | int | ✅ |  科室id | |
| data. department_name | string | ✅ |  科室名称 | |
| data. id | int | ✅ |  号源id | |
| data. left_num | int | ✅ |  余号| |
| data. personnel_id | int | ✅ |  医生id | |
| data. personnel_name | string | ✅ |  医生姓名 | |
| data. tatal_num | int | ✅ |  总号源数 | |
| data. visit_date | time | ✅ |  日期 | |
--



</br>
<h3>4.2 号源科室列表

```
请求地址：/doctorVisitSchedule/departments
```
**请求包示例**

```
{
	clinic_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Number | ✅ |  诊所id | |

**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "department_id": 5,
            "name": "普通内科"
        },
        {
            "department_id": 2,
            "name": "眼科"
        },
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | Array | ❌ |   | |
| data. department_id | int | ✅ |  科室id | |
| data. name | string | ✅ |  科室名称 | |
--



</br>
<h3>4.3 号源科室下医生列表

```
请求地址：/doctorVisitSchedule/doctors
```
**请求包示例**

```
{
	department_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| department_id | int | ✅ |  科室id | |

**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "name": "黄飞鸿",
            "personnel_id": 22
        },
        {
            "name": "扁鹊",
            "personnel_id": 2
        }
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | Array | ❌ |   | |
| data. personnel_id | int | ✅ |  医生id | |
| data. name | string | ✅ |  医生姓名 | |
--


</br>
<h3>4.4 获取所有医生的号源信息

```
请求地址：/doctorVisitSchedule/DoctorsWithSchedule
```
**请求包示例**

```
{
	clinic_id:1,
	department_id: 1,
	personnel_id: 1,
	start_date: "2018-08-19",
	end_date: "2018-08-26",
	offset: 0,
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ |  诊所id | |
| department_id | int | ❌ |  科室id | |
| personnel_id | int | ❌ |  医生id | |
| start_date | string | ✅ |  开始日期 | |
| end_date | string | ✅ |  结束日期 | |
| offset | int | ✅ | 开始条数 | |
| limit | int | ✅ |  条数 | |

**应答包示例**

```
{
    "canOverride": false,
    "code": "200",
    "data": [
        {
            "department_id": 1,
            "department_name": "骨科",
            "personnel_id": 2,
            "personnel_name": "扁鹊",
            "schedules": [
                {
                    "am_pm": "a",
                    "department_id": 1,
                    "doctor_visit_schedule_id": 1850,
                    "id": 1850,
                    "open_flag": true,
                    "personnel_id": 2,
                    "stop_flag": false,
                    "visit_date": "2018-08-20T00:00:00Z"
                },
                ...
            ]
        },
        ...
    ],
    "needOpen": false,
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 21
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| canOverride | boolean | ✅ |  该时间段内排版是否可以被覆盖 | |
| data | Array | ❌ |   | |
| data. department_id | int | ✅ |  科室id | |
| data. department_name | string | ✅ |  科室名称 | |
| data. personnel_id | int | ✅ |  医生id | |
| data. personnel_name | string | ✅ |  医生姓名 | |
| data. schedules | Array | ❌ | 号源列表  | |
| data. schedules. am_pm | string | ✅ | 上下午  | |
| data. schedules. department_id | int | ✅ | 科室id  | |
| data. schedules. doctor_visit_schedule_id | int | ✅ | 号源id  | |
| data. schedules. id | int | ✅ | 号源id   | |
| data. schedules. open_flag | booleam | ✅ | 是否开放号源  | |
| data. schedules. personnel_id | int | ✅ | 医生id  | |
| data. schedules. stop_flag | string | ✅ | 是否停诊  | |
| data. schedules. visit_date | time | ✅ | 就诊日期  | |
--



</br>
<h3>4.5 复制排版

```
请求地址：/doctorVisitSchedule/CopyScheduleByDate
```
**请求包示例**

```
{
	clinic_id:1,
	copy_start_date: "2018-08-19",
	insert_start_date: "2018-08-27",
	day_long: 7,
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ |  诊所id | |
| copy_start_date | string | ✅ | 复制开始时间 | |
| insert_start_date | string | ✅ | 新增开始时间  | |
| day_long | int | ✅ | 复制天数 | |

**应答包示例**

```
{
    "code": "200",
    "msg": "复制排版成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
--



</br>
<h3>4.6 开放号源

```
请求地址：/doctorVisitSchedule/OpenScheduleByDate
```
**请求包示例**

```
{
	clinic_id:1,
	start_date: "2018-08-19",
	day_long: 7,
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ |  诊所id | |
| start_date | string | ✅ | 开始时间 | |
| day_long | int | ✅ | 天数 | |

**应答包示例**

```
{
    "code": "200",
    "msg": "开放成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
--


</br>
<h3>4.7 插入单个号源

```
请求地址：/doctorVisitSchedule/CreateOneSchedule
```
**请求包示例**

```
{
	department_id:1,
	personnel_id: 1,
	visit_date: "2018-08-19",
	am_pm: "a",
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| department_id | int | ✅ |  科室id | |
| personnel_id | int | ✅ |  医生id | |
| visit_date | string | ✅ | 就诊时间 | |
| am_pm | int | ✅ | 上下午 | |

**应答包示例**

```
{
    "code": "200",
    "msg": "插入号源成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
--



</br>
<h3>4.8 删除单个未开放号源 byid

```
请求地址：/doctorVisitSchedule/DeleteOneUnOpenScheduleByID
```
**请求包示例**

```
{
	doctor_visit_schedule_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| doctor_visit_schedule_id | int | ✅ |  号源id | |

**应答包示例**

```
{
    "code": "200",
    "msg": "插入号源成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
--


</br>
<h3>4.9 停诊号源byid

```
请求地址：/doctorVisitSchedule/StopScheduleByID
```
**请求包示例**

```
{
	doctor_visit_schedule_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| doctor_visit_schedule_id | int | ✅ |  号源id | |

**应答包示例**

```
{
    "code": "200",
    "msg": "插入号源成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
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

6 就诊模块
--------

</br>
<h3>6.1 就诊患者登记

```
请求地址：/triage/register
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | int | ❌ |  患者id | |
| cert_no | string | ❌ |  身份证号 | |
| name | string | ✅ |  姓名 | |
| birthday | string | ✅ |  生日| |
| sex | int | ✅ |  0:女，1：男| |
| phone | string | ✅ |  手机号| |
| province | string | ❌ |  省份| |
| city | string | ❌ |  市| |
| district | string | ❌ |  区/县| |
| address | string | ❌ |  地址| |
| profession | string | ❌ |  职业| |
| remark | string | ❌ |  备注| |
| patient_channel_id | int | ❌ |  来源渠道id| |
| clinic_id | int | ✅ |  诊所id| |
| visit_type | int | ✅ |  就诊类型| |
| personnel_id | int | ❌ |  医生id | |
| department_id | int | ❌ |  科室id| |


**应答包示例**

```
{
    "code": "200",
    "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
--


</br>
<h3>6.2 登记患者详情

```
请求地址：/triage/TriagePatientDetail
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_triage_patient_id | int | ❌ |  患者就诊id | |


**应答包示例**

```
{
    "code": "200",
    "data": {
        "address": "哈哈哈哈",
        "birthday": "20001010",
        "cert_no": null,
        "city": "北京市",
        "clinic_patient_id": 1,
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
        "updated_time": "2018-05-28T00:26:29.012104+08:00",
        "visit_type": 1
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. address | string | ❌ | 地址 | |
| data. birthday | string | ✅ | 生日 | |
| data. cert_no | string | ❌ | 身份证 | |
| data. city | string | ❌ | 市 | |
| data. clinic_patient_id | int | ❌ | 诊所就诊人id | |
| data. created_time | string | ❌ | 创建时间 | |
| data. deleted_time | string | ❌ | 删除时间 | |
| data. district | string | ❌ | 区、县 | |
| data. id | int | ❌ | 就诊人id | |
| data. image | string | ❌ | 头像url | |
| data. name | string | ❌ | 姓名 | |
| data. patient_channel_id | int | ❌ | 来源渠道id | |
| data. phone | string | ❌ | 手机号 | |
| data. profession | string | ❌ | 职业 | |
| data. province | string | ❌ | 省份 | |
| data. remark | string | ❌ | 备注 | |
| data. sex | int | ❌ | 0:女，1：男 | |
| data. status | boolean | ❌ | 是否启用 | |
| data. updated_time | time | ❌ | 修改时间 | |
| data. visit_type | int | ❌ | 就诊类型 | |
--


</br>
<h3>6.3 当日登记就诊人列表 分诊记录

```
请求地址：/triage/patientlist
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ |  诊所id | |
| keyword | string | ❌ |  关键字 | |
| status_start | string | ❌ |  状态 最小值 | |
| status_end | string | ❌ |  状态 最大值 | |
| register_type | string | ❌ | 登记类型：1预约，2线下分诊 3快速接诊 | |
| personnel_id | string | ❌ |  医生id | |
| department_id | string | ❌ |  科室id | |
| is_today | string | ❌ |  是否当天 | |
| startDate | string | ❌ |  开始时间 | |
| endDate | string | ❌ |  结束时间 | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "birthday": "19920706",
            "clinic_patient_id": 2,
            "clinic_triage_patient_id": 138,
            "department_name": null,
            "doctor_name": null,
            "patient_id": 3,
            "patient_name": "查康",
            "phone": "18701676735",
            "register_personnel_name": "超级管理员",
            "register_time": "2018-08-16T10:57:51.613363+08:00",
            "register_type": 2,
            "sex": 1,
            "status": 10,
            "updated_time": "2018-08-16T10:57:51.613363+08:00",
            "visit_date": "2018-08-16T00:00:00Z"
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 137
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. birthday | string | ✅ | 生日 | |
| data. clinic_patient_id | int | ✅ | 诊所就诊人id | |
| data. clinic_triage_patient_id | int | ✅ | 就诊id | |
| data. department_name | int | ❌ | 就诊科室名称 | |
| data. doctor_name | string | ❌ | 就诊医生姓名 | |
| data. patient_id | int | ✅ | 患者id| |
| data. patient_name | string | ✅ | 患者姓名| |
| data. phone | string | ✅ | 手机号 | |
| data. register_personnel_name | string | ✅ | 登记人员姓名| |
| data. register_time | time | ✅ | 登记时间 | |
| data. register_type | int | ✅ | 登记类型 | |
| data. sex | int | ✅ | 性别 | |
| data. status | int | ✅ | 就诊状态 10:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消| |
| data. updated_time | time | ✅ | 修改时间 | |
| data. visit_date | time | ✅ | 就诊日期 | |
--



</br>
<h3>6.4 接诊就诊人列表

```
请求地址：/triage/RecptionPatientList
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ |  诊所id | |
| personnel_id | int | ✅ |  医生id | |
| query_type | int | ✅ |  待接诊 0, 已接诊 1 | |
| startDate | string | ❌ |  开始时间 | |
| endDate | string | ❌ |  结束时间 | |
| keyword | string | ❌ |  搜索关键字 | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "birthday": "19890402",
            "clinic_patient_id": 25,
            "clinic_triage_patient_id": 126,
            "department_name": "骨科",
            "doctor_name": "扁鹊",
            "patient_id": 29,
            "patient_name": "赵丽颖",
            "phone": "15387556262",
            "register_personnel_name": null,
            "register_time": "2018-08-02T21:19:37.940828+08:00",
            "register_type": 3,
            "sex": 0,
            "status": 40,
            "updated_time": "2018-08-12T22:57:13.559599+08:00",
            "visit_date": "2018-08-02T00:00:00Z"
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 18
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. birthday | string | ✅ | 生日 | |
| data. clinic_patient_id | int | ✅ | 诊所就诊人id | |
| data. clinic_triage_patient_id | int | ✅ | 就诊id | |
| data. department_name | int | ❌ | 就诊科室名称 | |
| data. doctor_name | string | ❌ | 就诊医生姓名 | |
| data. patient_id | int | ✅ | 患者id| |
| data. patient_name | string | ✅ | 患者姓名| |
| data. phone | string | ✅ | 手机号 | |
| data. register_personnel_name | string | ✅ | 登记人员姓名| |
| data. register_time | time | ✅ | 登记时间 | |
| data. register_type | int | ✅ | 登记类型 | |
| data. sex | int | ✅ | 性别 | |
| data. status | int | ✅ | 就诊状态 10:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消| |
| data. updated_time | time | ✅ | 修改时间 | |
| data. visit_date | time | ✅ | 就诊日期 | |
--


</br>
<h3>6.5 通过id就诊人 查询患者

```
请求地址：/triage/getById
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| id | int | ✅ |  患者id | |



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
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. address | string | ❌ | 地址 | |
| data. birthday | string | ✅ | 生日 | |
| data. cert_no | string | ❌ | 身份证 | |
| data. city | string | ❌ | 城市 | |
| data. created_time | time | ✅ | 创建时间 | |
| data. deleted_time | time | ❌ | 删除时间 | |
| data. district | string | ❌ | 区、县| |
| data. id | int | ✅ | 患者id | |
| data. image | string | ✅ | 头像url | |
| data. name | string | ✅ | 姓名 | |
| data. patient_channel_id | int | ❌ | 渠道id | |
| data. phone | string | ✅ | 手机号 | |
| data. profession | string | ❌ | 职业 | |
| data. province | string | ❌ | 省份 | |
| data. remark | string | ❌ | 备注 | |
| data. sex | string | ✅ | 性别 | |
| data. status | string | ✅ | 是否启用 | |
| data. updated_time | string | ✅ | 修改时间 | |
--


</br>
<h3>6.6 分诊医生列表

```
请求地址：/triage/personnelList
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ |  诊所id | |
| department_id | int | ❌ | 科室id | |
| offset | int | ❌ | 开始条数 | |
| limit | int | ❌ | 条数 | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "department_name": "耳鼻喉科",
            "doctor_name": "富察音容",
            "doctor_visit_schedule_id": 784,
            "triaged_total": 0,
            "wait_total": 0
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 7
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. department_name | string | ✅ | 科室名称 | |
| data. doctor_name | string | ✅ | 医生名称 | |
| data. doctor_visit_schedule_id | int | ✅ | 号源id| |
| data. triaged_total | int | ✅ | 已分人数 | |
| data. wait_total | int | ✅ | 等待接诊人数 | |

--


</br>
<h3>6.7 晚上体征信息

```
请求地址：/triage/completeBodySign
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| weight | float | ❌ | 体重 | |
| height | float | ❌ | 身高 | |
| bmi | float | ❌ | bmi | |
| blood_type | string | ❌ | 血型 | |
| rh\_blood_type | int | ❌ | RH血型 -1: 阴性，1阳性, 0: 未查 | |
| temperature_type | int | ❌ | RH血型 1: 口温，2：耳温，3：额温，4：腋温，5：肛温 | |
| temperature | float | ❌ | 温度 | |
| breathe | int | ❌ | 呼吸(次/分钟) | |
| pulse | int | ❌ | 脉搏(次/分钟) | |
| systolic\_blood_pressure | int | ❌ | 血压收缩压 | |
| diastolic\_blood_pressure | int | ❌ | 血压舒张压 | |
| blood\_sugar_time | string | ❌ | 血糖时间 | |
| concentration\_before_retiring | float | ❌ | 睡前血糖浓度(mmol/I) | |
| concentration\_after_dinner | float | ❌ | 晚餐后血糖浓度(mmol/I) | |
| concentration\_before_dinner | float | ❌ | 晚餐前血糖浓度(mmol/I) | |
| concentration\_after_lunch | float | ❌ | 午餐后血糖浓度(mmol/I) | |
| concentration\_before_lunch | float | ❌ | 午餐血糖浓度(mmol/I) | |
| concentration\_after_breakfast | float | ❌ | 早餐后血糖浓度(mmol/I) | |
| concentration\_before_breakfast | float | ❌ | 早餐前血糖浓度(mmol/I) | |
| concentration\_before_dawn | float | ❌ | 凌晨血糖浓度(mmol/I) | |
| concentration\_empty_stomach | float | ❌ | 空腹血糖浓度(mmol/I) | |
| left_vision | string | ❌ | 左眼视力 | |
| right_vision | string | ❌ | 右眼视力 | |
| oxygen_saturation | string | ❌ | 氧饱和度(%) | |
| pain_score | string | ❌ | 疼痛评分 | |
| remark | string | ❌ | 备注 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "保存成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.8 完善诊前病历

```
请求地址：/triage/completePreMedicalRecord
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| has\_allergic_history | boolean | ❌ | 是否有过敏 | |
| allergic_history | string | ❌ | 过敏史 | |
| allergic_reaction | string | ❌ | 过敏反应 | |
| personal\_medical_history | string | ❌ | 个人病史 | |
| family\_medical_history | string | ❌ | 家族病史 | |
| immunizations | string | ❌ | 接种疫苗 | |
| menarche_age | int | ❌ | 月经初潮年龄 | |
| menstrual\_period\_start_day | string | ❌ | 月经经期开始时间 | |
| menstrual\_period\_end_day | string | ❌ | 月经经期结束时间 | |
| menstrual\_cycle\_start_day | string | ❌ | 月经周期结束时间 | |
| menstrual\_cycle\_end_day | string | ❌ | 月经周期结束时间 | |
| menstrual\_last_day | string | ❌ | 末次月经时间 | |
| gestational_weeks | string | ❌ | 孕周 | |
| childbearing_history | string | ❌ | 生育史 | |
| remark | string | ❌ | 备注 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "保存成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.9 完善诊前欲诊

```
请求地址：/triage/completePreDiagnosis
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| chief_complaint | string | ❌ | 主诉 | |
| history\_of\_present_illness | string | ❌ | 现病史 | |
| history\_of\_past_illness | string | ❌ | 既往史 | |
| body_examination | string | ❌ | 体格检查 | |
| remark | string | ❌ | 备注 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "保存成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--



</br>
<h3>6.10 获取健康档案

```
请求地址：/triage/GetHealthRecord
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
 	"code": "200",
    "msg": "ok",
    "body_sign": {
        "blood_sugar_time": null,
        "blood_type": "A",
        "bmi": 14.36,
        "breathe": 72,
        "clinic_triage_patient_id": 120,
        "concentration_after_breakfast": null,
        "concentration_after_dinner": null,
        "concentration_after_lunch": null,
        "concentration_before_breakfast": null,
        "concentration_before_dawn": null,
        "concentration_before_dinner": null,
        "concentration_before_lunch": null,
        "concentration_before_retiring": null,
        "concentration_empty_stomach": null,
        "created_time": "2018-08-01T00:02:39.738378+08:00",
        "deleted_time": null,
        "diastolic_blood_pressure": 139,
        "height": 171,
        "id": 49,
        "left_vision": "5.2",
        "oxygen_saturation": 20,
        "pain_score": 6,
        "pulse": 78,
        "remark": null,
        "rh_blood_type": null,
        "right_vision": "5.2",
        "systolic_blood_pressure": 110,
        "temperature": 37.6,
        "temperature_type": 1,
        "updated_time": "2018-08-01T00:02:39.738378+08:00",
        "weight": 42
    },
    "pre_diagnosis": {
        "body_examination": "T:38.7℃ P:100次/min BP:154/82mmHg,神志恍惚，营养一般，皮肤弹性稍差，呼吸急促，口唇紫绀，胸廓呈桶状，呼吸运动减弱，呼气延长，两肺可听到散在的哮鸣音和干啰音",
        "chief_complaint": "咳嗽、咳痰20年，加重两周，发热1周，神志恍惚1天入院。",
        "clinic_triage_patient_id": 120,
        "created_time": "2018-07-31T23:59:17.624182+08:00",
        "deleted_time": null,
        "history_of_past_illness": "无肺炎、肺结核和过敏史、无高血压、无心脏病史",
        "history_of_present_illness": "自20年前有咳嗽、咳白色泡沫样痰。每逢劳累、气候变化或受凉后，咳嗽咳痰加重。冬季病情复发，持续2-3个月。六年前开始有气喘，起初在体重物和快步行走时气促。",
        "id": 23,
        "remark": "第N次就诊",
        "updated_time": "2018-07-31T23:59:17.624182+08:00"
    },
    "pre_medical_record": {}
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| body_sign | obj | ❌ |  体征同6.7 | |
| pre_medical_record | obj | ❌ | 诊前病历 同6.8  | |
| pre_diagnosis | obj | ❌ |  预诊 同6.9| |

--



</br>
<h3>6.11 分诊、换诊(选择医生)

```
请求地址：/triage/chooseDoctor
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| doctor\_visit\_schedule_id | int | ✅ |  排版id | |
| triage\_personnel_id | int | ✅ | 分诊人员 id | |



**应答包示例**

```
{
 	"code": "200",
    "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.12 医生接诊病人

```
请求地址：/triage/reception
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| recept\_personnel_id | int | ✅ |  接诊人id | |



**应答包示例**

```
{
 	"code": "200",
    "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.13 医生完成接诊

```
请求地址：/triage/complete
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| recept\_personnel_id | int | ✅ |  接诊人id | |



**应答包示例**

```
{
 	"code": "200",
    "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.14 按日期统计挂号记录

```
请求地址：/triage/AppointmentsByDate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ |  诊所id | |
| department_id | int | ❌ | 科室id | |
| personnel_id | int | ❌ | 医生id | |
| start_date | string | ✅ | 开始日期 | |
| offset | int | ❌ | 开始条数 | |
| limit | int | ❌ | 条数 | |
| day_long | int | ✅ | 天数 | |



**应答包示例**

```
{
    "clinic_array": [
        {
            "am_pm": "a",
            "count": 1,
            "visit_date": "2018-06-06T00:00:00Z"
        }
    ],
    "code": "200",
    "doctor_array": [
        {
            "am_pm": null,
            "count": null,
            "department_id": 3,
            "department_name": "全科测试全能科室测试",
            "personnel_id": 11,
            "personnel_name": "人中龙凤",
            "visit_date": null
        },
        ...
    ],
    "msg": "ok",
    "page_info": {
        "limit": 10,
        "offset": 0,
        "total": 21
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| clinic\_array | Array | ❌ |  该诊所 就诊统计 按日期，上下午分组  | |
| clinic\_array. am_pm | string | ❌ |  上下午  | |
| clinic\_array. count | int | ❌ |  数量 | |
| clinic\_array. visit_date | time | ❌ | 时间 | |
| doctor\_array | Array | ❌ |  单个医生 就诊统计 按日期，上下午分组  | |
| doctor\_array. am_pm | string | ❌ |  上下午  | |
| doctor\_array. count | int | ❌ |  数量 | |
| doctor\_array. department_id | int | ❌ | 科室id | |
| doctor\_array. department_name | int | ❌ | 科室名称 | |
| doctor\_array. personnel_id | int | ❌ | 医生id | |
| doctor\_array. personnel_id | string | ❌ | 医生名称 | |
| doctor\_array. visit_date | time | ❌ | 时间 | |

--



</br>
<h3>6.15 开治疗

```
请求地址：/triage/TreatmentPatientCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| personnel_id | int | ✅ |  操作人id | |
| items | string | ✅ |  json 字符串，具体项目 | |
| items. clinic_treatment_id | string | ✅ | 治疗项目id | |
| items. times | string | ✅ | 治疗次数 | |
| items. illustration | string | ❌ | 说明 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.16 查询治疗

```
请求地址：/triage/TreatmentPatientGet
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "clinic_treatment_id": 1,
            "clinic_triage_patient_id": 120,
            "created_time": "2018-08-01T00:28:57.072356+08:00",
            "deleted_time": null,
            "id": 43,
            "illustration": "哈哈哈哈倾世容颜倾世容",
            "left_times": 1,
            "operation_id": 20,
            "order_sn": "201808010028577120775",
            "order_status": "10",
            "paid_status": false,
            "price": 5000,
            "soft_sn": 0,
            "times": 1,
            "treatment_name": "打针",
            "unit_name": "次",
            "updated_time": "2018-08-01T00:28:57.072356+08:00"
        }
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. clinic_treatment_id | int | ✅ | 治疗项目id | |
| data. clinic_triage_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. deleted_time | time | ❌ | 删除时间 | |
| data. id | int | ✅ | id | |
| data. illustration | sting | ❌ | 说明 | |
| data. left_times | int | ✅ | 治疗剩余次数 | |
| data. operation_id | int | ✅ | 开治疗 医生 id | |
| data. order_sn | string | ✅ | 订单号 | |
| data. order_status | string | ✅ | 治疗状态 | |
| data. paid_status | boolean | ✅ | 支付状态 | |
| data. price | int | ✅ | 单价（分） | |
| data. soft_sn | int | ✅ | 订单号 序号 | |
| data. times | int | ✅ | 治疗次数 | |
| data. treatment_name | string | ✅ | 治疗项目名称 | |
| data. unit_name | string | ✅ | 单位 | |
| data. updated_time | time | ✅ | 修改时间 | |

--



</br>
<h3>6.17 开检验

```
请求地址：/triage/LaboratoryPatientCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| personnel_id | int | ✅ |  操作人id | |
| items | string | ✅ |  json 字符串，具体项目 | |
| items. clinic_laboratory_id | string | ✅ | 检验项目id | |
| items. times | string | ✅ | 次数 | |
| items. illustration | string | ❌ | 说明 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.18 获取检验

```
请求地址：/triage/LaboratoryPatientGet
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "checking_time": null,
            "clinic_laboratory_id": 2,
            "clinic_triage_patient_id": 120,
            "created_time": "2018-08-01T00:29:02.465895+08:00",
            "deleted_time": null,
            "id": 135,
            "illustration": "啥的噶",
            "laboratory_name": "尿常规",
            "operation_id": 20,
            "order_sn": "201808010029023120720",
            "order_status": "10",
            "paid_status": false,
            "price": 40000,
            "soft_sn": 0,
            "times": 1,
            "updated_time": "2018-08-01T00:29:02.465895+08:00"
        },
        ...
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. checking_time | time | ❌ | 接收时间(待检验变为检验中的时间) | |
| data. clinic\_laboratory_id | int | ✅ | 检验项目id | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. deleted_time | time | ❌ | 删除时间 | |
| data. id | int | ✅ | id | |
| data. illustration | string | ❌ | 说明 | |
| data. laboratory_name | string | ❌ | 检验项目时间 | |
| data. operation_id | string | ✅ | 操作人 | |
| data. order_sn | string | ✅ | 订单号 | |
| data. order_status | string | ✅ | 检验状态 | |
| data. paid_status | boolean | ✅ | 支付状态 | |
| data. price | int | ✅ | 单价| |
| data. soft_sn | int | ✅ | 订单号序号 | |
| data. times | int | ✅ | 次数 | |
| data. updated_time | time | ✅ | 修改时间 | |

--



</br>
<h3>6.19 开西/成药处方

```
请求地址：/triage/PrescriptionWesternPatientCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| personnel_id | int | ✅ |  操作人id | |
| items | string | ✅ |  json 字符串，具体项目 | |
| items. clinic\_drug_id | string | ✅ | 西药项目id | |
| items. once_dose | string | ✅ | 单次用量 | |
| items. once\_dose\_unit_name | string | ✅ | 单次用量单位 | |
| items. route\_administration_name | string | ✅ | 用法 | |
| items. frequency_name | string | ✅ | 频率 | |
| items. amount | string | ✅ | 总量 | |
| items. illustration | string | ❌ | 说明 | |
| items. fetch_address | string | ✅ | 取药地点 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--



</br>
<h3>6.20 获取西药处方

```
请求地址：/triage/PrescriptionWesternPatientGet
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "amount": 2,
            "clinic_drug_id": 59,
            "clinic_triage_patient_id": 120,
            "drug_name": "可达龙片(盐酸胺碘酮片)",
            "eff_day": 5,
            "fetch_address": 0,
            "frequency_name": "1次/日 (8am)",
            "id": 231,
            "illustration": null,
            "once_dose": 1,
            "once_dose_unit_name": "g",
            "operation_id": 20,
            "order_sn": "201808010028441120721",
            "packing_unit_name": "盒",
            "paid_status": false,
            "route_administration_name": "口服                  ",
            "soft_sn": 1,
            "specification": "0.2g/片",
            "stock_amount": null,
            "type": 0
        },
        ...
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. amount | int | ✅ | 总量 | |
| data. clinic\_drug_id | int | ✅ | 西药收费项目id | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. drug_name | string | ✅ | 药品名称 | |
| data. eff_day | int | ✅ | 用药天数 | |
| data. fetch_address | int | ✅ | 取药地点 | |
| data. frequency_name | string | ✅ | 频率 | |
| data. id | int | ❌ | id | |
| data. illustration | string | ✅ | 说明 | |
| data. once_dose | string | ✅ | 单词用量 | |
| data. once_dose_unit_name | string | ✅ | 单词用量单位 | |
| data. operation_id | int | ✅ | 操作人id | |
| data. order_sn | string | ✅ | 订单号 | |
| data. packing_unit_name | string | ✅ | 包装单位名称 | |
| data. paid_status | boolean | ✅ | 支付状态 | |
| data. route_administration_name | boolean | ✅ | 用法 | |
| data. soft_sn | boolean | ✅ | 订单号 序号 | |
| data. specification | string | ✅ | 规格 | |
| data. stock_amount | int | ✅ | 库存 | |
| data. type | int | ✅ | 类型 0-西药 1-中药 | |

--



</br>
<h3>6.21 获取西药历史处方列表

```
请求地址：/triage/PrescriptionWesternPatientList
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_patient_id | int | ✅ |  诊所就诊人id | |
| keyword | string | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "clinic_triage_patient_id": 3,
            "created_time": "2018-06-03T21:47:51.117251+08:00",
            "department_name": "骨科",
            "diagnosis": "",
            "personnel_name": "扁鹊",
            "prescription_chinese_patient_id": 30,
            "visit_type": 1
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
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. department_name | string | ✅ | 就诊科室 | |
| data. diagnosis | string | ✅ | 诊断 | |
| data. personnel_name | string | ✅ | 就诊医生 | |
| data. prescription_chinese_patient_id | int | ✅ | 西药处方id | |
| data. visit_type | int | ✅ | 就诊类型 | |

--



</br>
<h3>6.22 开中药处方

```
请求地址：/triage/PrescriptionChinesePatientCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| route_administration_name | int | ✅ |  用法 | |
| frequency_name | int | ✅ | 频率 | |
| id | int | ❌ |  中药处方id | |
| amount | int | ✅ | 总付数 | |
| medicine_illustration | string | ❌ | 用药说明 | |
| fetch_address | int | ✅ | 取药地点 | |
| eff_day | int | ✅ | 用药天数 | |
| personnel_id | int | ✅ | 操作人id | |
| items | string | ✅ |  json 字符串，具体项目 | |
| items. clinic\_drug_id | string | ✅ | 中药项目id | |
| items. once_dose | string | ✅ | 单次用量 | |
| items. once\_dose\_unit_name | string | ✅ | 单次用量单位 | |
| items. amount | string | ✅ | 总量 | |
| items. special_illustration | string | ❌ | 说明 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.23 删除中药处方

```
请求地址：/triage/PrescriptionChinesePatientDelete
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| id | int | ✅ | 中药处方id | |
| personnel_id | int | ✅ | 操作人id | |



**应答包示例**

```
{
    "code": "200",
    "msg": ""
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--



</br>
<h3>6.24 获取中药处方

```
请求地址：/triage/PrescriptionChinesePatientGet
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "amount": 5,
            "clinic_triage_patient_id": 120,
            "created_time": "2018-08-01T00:28:48.025089+08:00",
            "deleted_time": null,
            "eff_day": 5,
            "fetch_address": 0,
            "frequency_name": "1次/日 (8am)",
            "id": 72,
            "items": [
                {
                    "amount": 50,
                    "clinic_drug_id": 12,
                    "drug_name": "当归",
                    "id": 125,
                    "once_dose": 10,
                    "once_dose_unit_name": "g",
                    "order_sn": "20180801002848212020",
                    "paid_status": false,
                    "prescription_chinese_patient_id": 72,
                    "soft_sn": 0,
                    "special_illustration": null,
                    "specification": "/kg",
                    "stock_amount": 9968,
                    "type": 1
                },
                ...
            ],
            "medicine_illustration": null,
            "operation_id": 20,
            "order_sn": "20180801002848212020",
            "route_administration_name": "水煎服",
            "updated_time": "2018-08-01T00:28:48.025089+08:00"
        }
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. amount | int | ✅ | 总量 | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. deleted_time | time | ❌ | 删除时间 | |
| data. eff_day | int | ✅ | 用药天数 | |
| data. fetch_address | int | ✅ | 取药地点 | |
| data. frequency_name | string | ✅ | 频率 | |
| data. id | int | ✅ | 中药处方id | |
| data. items | array | ✅ | 中药处方详情| |
| data. items. amount | int | ✅ | 总量 | |
| data. items. clinic_drug_id | int | ✅ | 中药项目id | |
| data. items. drug_name | string | ✅ | 药品名称 | |
| data. items. id | int | ✅ | 细目id | |
| data. items. once_dose | int | ✅ | 单次用量 | |
| data. items. once_dose_unit_name | string | ✅ | 单次用量单位 | |
| data. items. order_sn | string | ✅ | 订单号 | |
| data. items. paid_status | booleam | ✅ | 支付状态 | |
| data. items. prescription\_chinese\_patient_id | int | ✅ | 中药处方id | |
| data. items. soft_sn | int | ✅ | 订单号 序号 | |
| data. items. special_illustration | string | ❌ | 说明 | |
| data. items. specification | string | ✅ | 规格 | |
| data. items. stock_amount | int | ✅ | 库存 | |
| data. items. type | int | ✅ | 类型 | |

--




</br>
<h3>6.25 获取中药历史处方列表

```
请求地址：/triage/PrescriptionChinesePatientList
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_patient_id | int | ✅ |  诊所就诊人id | |
| keyword | string | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "clinic_triage_patient_id": 43,
            "created_time": "2018-08-12T22:48:55.09567+08:00",
            "department_name": "骨科",
            "diagnosis": "11牙体缺损,高血压病 1级 极高危,人感染高致病性禽流感A(H5N1),高血压病 1级 高危,啊啊啊啊啊啊啊啊,aa,阿达萨达,外层渗出性视网膜病(Coats病),Eales病",
            "personnel_name": "扁鹊",
            "prescription_chinese_patient_id": 38,
            "visit_type": 1
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
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. department_name | string | ✅ | 就诊科室 | |
| data. diagnosis | string | ✅ | 诊断 | |
| data. personnel_name | string | ✅ | 就诊医生 | |
| data. prescription_chinese_patient_id | int | ✅ | 中药处方id | |
| data. visit_type | int | ✅ | 就诊类型 | |

--




</br>
<h3>6.26 开检查

```
请求地址：/triage/ExaminationPatientCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| personnel_id | int | ✅ |  操作人id | |
| items | string | ✅ |  json 字符串，具体项目 | |
| items. clinic_examination_id | string | ✅ | 检查项目id | |
| items. times | string | ✅ | 次数 | |
| items. illustration | string | ❌ | 说明 | |
| items. organ | string | ❌ | 部位 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--




</br>
<h3>6.27 获取检查

```
请求地址：/triage/ExaminationPatientGet
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "clinic_examination_id": 1,
            "clinic_triage_patient_id": 120,
            "created_time": "2018-08-01T00:29:08.504874+08:00",
            "deleted_time": null,
            "id": 50,
            "illustration": "agfa",
            "name": "胸部正位",
            "operation_id": 20,
            "order_sn": "201808010029084120627",
            "order_status": "10",
            "organ": null,
            "paid_status": false,
            "price": 50000,
            "soft_sn": 0,
            "times": 1,
            "updated_time": "2018-08-01T00:29:08.504874+08:00"
        },
        ...
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. clinic\_examination_id | int | ✅ | 检查项目id | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. deleted_time | time | ❌ | 删除时间 | |
| data. id | int | ✅ | id | |
| data. illustration | string | ❌ | 说明 | |
| data. name | string | ✅ | 检查名称 | |
| data. operation_id | int | ✅ | 操作人| |
| data. order_sn | string | ✅ | 订单号 | |
| data. order_status | string | ✅ | 检查状态 | |
| data. organ | string | ❌ | 部位 | |
| data. paid_status | boolean | ✅ | 支付状态 | |
| data. price | int | ✅ | 单价 | |
| data. soft_sn | int | ✅ | 订单 序号| |
| data. times | int | ✅ | 次数 | |
| data. updated_time | string | ✅ | 修改时间 | |

--


</br>
<h3>6.28 开其它费用

```
请求地址：/triage/OtherCostPatientCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| personnel_id | int | ✅ |  操作人id | |
| items | string | ✅ |  json 字符串，具体项目 | |
| items. clinic_examination_id | string | ✅ | 其他收费项目id | |
| items. times | string | ✅ | 次数 | |
| items. illustration | string | ❌ | 说明 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--




</br>
<h3>6.29 获取其它费

```
请求地址：/triage/OtherCostPatientGet
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "amount": 1,
            "clinic_other_cost_id": 3,
            "clinic_triage_patient_id": 120,
            "created_time": "2018-08-01T00:29:54.829505+08:00",
            "deleted_time": null,
            "id": 18,
            "illustration": null,
            "name": "快递费",
            "operation_id": 20,
            "order_sn": "201808010029546120982",
            "paid_status": false,
            "price": 2000,
            "soft_sn": 0,
            "unit_name": "次",
            "updated_time": "2018-08-01T00:29:54.829505+08:00"
        }
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. clinic\_other\_cost_id | int | ✅ | 其他收费项目id | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. deleted_time | time | ❌ | 删除时间 | |
| data. id | int | ✅ | id | |
| data. illustration | string | ❌ | 说明 | |
| data. name | string | ✅ | 收费项目名称 | |
| data. operation_id | int | ✅ | 操作人id | |
| data. order_sn | string | ✅ | 订单号 | |
| data. paid_status | booleam | ✅ | 支付状态 | |
| data. price | int | ✅ | 单价 | |
| data. soft_sn | int | ✅ | 订单号序号 | |
| data. unit_name | string | ✅ | 单位名称 | |
| data. updated_time | time | ✅ | 修改时间 | |

--




</br>
<h3>6.30 开材料费

```
请求地址：/triage/MaterialPatientCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| personnel_id | int | ✅ |  操作人id | |
| items | string | ✅ |  json 字符串，具体项目 | |
| items. clinic_material_id | string | ✅ | 材料项目id | |
| items. times | string | ✅ | 次数 | |
| items. illustration | string | ❌ | 说明 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.31 获取其它费

```
请求地址：/triage/MaterialPatientGet
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "amount": 1,
            "clinic_material_id": 1,
            "clinic_triage_patient_id": 120,
            "created_time": "2018-08-01T00:29:32.126779+08:00",
            "deleted_time": null,
            "id": 30,
            "illustration": null,
            "name": "针筒",
            "operation_id": 20,
            "order_sn": "201808010029325120953",
            "paid_status": false,
            "price": 2000,
            "soft_sn": 0,
            "specification": "",
            "stock_amount": 199,
            "unit_name": "个",
            "updated_time": "2018-08-01T00:29:32.126779+08:00"
        }
    ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. amount | int | ✅ | 总量 | |
| data. clinic\_other\_cost_id | int | ✅ | 其他收费项目id | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. deleted_time | time | ❌ | 删除 | |
| data. id | int | ✅ | id | |
| data. illustration | string | ✅ | 说明 | |
| data. name | string | ✅ | 材料名称 | |
| data. operation_id | int | ✅ | 操作人id | |
| data. order_sn | string | ✅ | 订单号 | |
| data. paid_status | boolean | ✅ | 支付状态 | |
| data. price | int | ✅ | 单价 | |
| data. soft_sn | int | ✅ | 订单号序号 | |
| data. specification | string | ✅ | 规格 | |
| data. stock_amount | int | ✅ | 库存 | |
| data. unit_name | string | ✅ | 单位 | |
| data. updated_time | time | ✅ | 修改时间 | |

--



</br>
<h3>6.32 获取病人历史已接诊记录

```
请求地址：/triage/ReceiveRecord
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_patient_id | int | ✅ |  诊所患者id | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |
| keyword | int | ❌ | 搜索关键字 | |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "clinic_triage_patient_id": 5,
            "created_time": "2018-05-30T01:53:20.562036+08:00",
            "department_name": "骨科",
            "diagnosis": "",
            "doctor_name": "扁鹊",
            "ep_count": 0,
            "lp_count": 0,
            "mp_count": 0,
            "ocp_count": 0,
            "pcp_count": 0,
            "pwp_count": 1,
            "tp_count": 0,
            "visit_type": 1
        }
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 2
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. clinic\_triage\_patient_id | int | ✅ | 就诊id | |
| data. created_time | time | ✅ | 创建时间 | |
| data. department_name | string | ✅ | 科室名称 | |
| data. diagnosis | string | ✅ | 诊断 | |
| data. doctor_name | string | ✅ | 医生名称 | |
| data. ep_count | int | ✅ | 检查| |
| data. lp_count | int | ✅ | 检验 | |
| data. mp_count | int | ✅ | 材料费 | |
| data. ocp_count | int | ✅ | 其他费 | |
| data. pcp_count | int | ✅ | 中药处方 | |
| data. pwp_count | int | ✅ | 西药处方 | |
| data. tp_count | int | ✅ | 治疗 | |
| data. visit_type | int | ✅ | 就诊类型 | |

--




</br>
<h3>6.33 开诊疗费

```
请求地址：/triage/DiagnosisTreatmentPatientCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |
| personnel_id | int | ✅ |  操作人id | |
| items | string | ✅ |  json 字符串，具体项目 | |
| items. clinic_diagnosis_treatment_id | string | ✅ | 诊疗项目id | |
| items. amount | string | ✅ | 次数 | |
| items. illustration | string | ❌ | 说明 | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--


</br>
<h3>6.34 获取病人就诊信息详情

```
请求地址：/triage/TriagePatientVisitDetail
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": {
        "allergic_history": "鸡蛋、西红柿、海鲜、money",
        "allergic_reaction": "皮肤瘙痒、红肿",
        "body_examination": "T:38.7℃ P:100次/min BP:154/82mmHg,神志恍惚，营养一般，皮肤弹性稍差，呼吸急促，口唇紫绀，胸廓呈桶状，呼吸运动减弱，呼气延长，两肺可听到散在的哮鸣音和干啰音",
        "chief_complaint": "咳嗽、咳痰20年，加重两周，发热1周，神志恍惚1天入院。",
        "clinic_examination_name": "胸部正位、胸部正位+侧位",
        "clinic_laboratory_name": "尿常规、血常规",
        "clinic_treatment_name": "打针",
        "clinic_triage_patient_id": 120,
        "created_time": "2018-08-01T00:12:07.911785+08:00",
        "cure_suggestion": "清理呼吸道无效：采取坐位或半坐位 给予充足水分或热量，每日饮水一千五百毫升 指导深呼吸和有效咳嗽 遵医嘱施行超声雾化等吸入疗法  给予抗生素、痰液稀释剂等",
        "deleted_time": null,
        "diagnosis": "慢性扁桃体疾病",
        "family_medical_history": "家中无遗传病病史。",
        "files": "[{\"docName\":\"7.29(2).pdf\",\"url\":\"/uploads/7.29(2).pdf\"},{\"docName\":\"随访导图.png\",\"url\":\"/uploads/随访导图.png\"}]",
        "history_of_past_illness": "无肺炎、肺结核和过敏史、无高血压、无心脏病史",
        "history_of_present_illness": "自20年前有咳嗽、咳白色泡沫样痰。每逢劳累、气候变化或受凉后，咳嗽咳痰加重。冬季病情复发，持续2-3个月。六年前开始有气喘，起初在体重物和快步行走时气促。",
        "id": 48,
        "immunizations": "天花",
        "is_default": true,
        "medical_records": null,
        "morbidity_date": "2018-07-31",
        "operation_id": 20,
        "operation_name": "胡一天",
        "personal_medical_history": "生长于香港。文盲。否认外地长期居住史。无疫区、疫水接触史。否认工业毒物、粉尘及放射性物质接触史。否认牧区、矿山、高氟区、低碘区居住史。平日生活规律，否认吸毒史。否认吸烟嗜好。否认饮酒嗜好。否认冶游史。第N次\n",
        "prescription_chinese_patient": [
            {
                "prescription_patient_id": 72,
                "operation_name": null,
                "route_administration_name": "水煎服",
                "eff_day": 5,
                "order_sn": "20180801002848212020",
                "amount": 5,
                "frequency_name": "1次/日 (8am)",
                "fetch_address": null,
                "medicine_illustration": null,
                "created_time": null,
                "updated_time": null,
                "items": [
                    {
                        "clinic_drug_id": null,
                        "type": null,
                        "drug_name": "当归",
                        "specification": null,
                        "stock_amount": null,
                        "once_dose": 10,
                        "once_dose_unit_name": "g",
                        "route_administration_name": null,
                        "frequency_name": null,
                        "eff_day": null,
                        "amount": 50,
                        "packing_unit_name": "g",
                        "fetch_address": null,
                        "illustration": null,
                        "special_illustration": null
                    },
                    {
                        "clinic_drug_id": null,
                        "type": null,
                        "drug_name": "桑白皮",
                        "specification": null,
                        "stock_amount": null,
                        "once_dose": 10,
                        "once_dose_unit_name": "g",
                        "route_administration_name": null,
                        "frequency_name": null,
                        "eff_day": null,
                        "amount": 50,
                        "packing_unit_name": "g",
                        "fetch_address": null,
                        "illustration": null,
                        "special_illustration": null
                    },
                    {
                        "clinic_drug_id": null,
                        "type": null,
                        "drug_name": "甘草稍",
                        "specification": null,
                        "stock_amount": null,
                        "once_dose": 10,
                        "once_dose_unit_name": "g",
                        "route_administration_name": null,
                        "frequency_name": null,
                        "eff_day": null,
                        "amount": 50,
                        "packing_unit_name": "g",
                        "fetch_address": null,
                        "illustration": null,
                        "special_illustration": null
                    }
                ]
            }
        ],
        "prescription_western_patient": [
            {
                "amount": 5,
                "eff_day": 5,
                "illustration": null,
                "name": "肿瘤相关抗原125定标液",
                "once_dose": 1,
                "once_dose_unit_name": "支",
                "order_sn": "201808010028441120721",
                "packing_unit_name": "盒",
                "route_administration_name": "口服<饭前>",
                "specification": "1mlx2支/盒"
            },
            {
                "amount": 2,
                "eff_day": 5,
                "illustration": null,
                "name": "可达龙片(盐酸胺碘酮片)",
                "once_dose": 1,
                "once_dose_unit_name": "g",
                "order_sn": "201808010028441120721",
                "packing_unit_name": "盒",
                "route_administration_name": "口服                  ",
                "specification": "0.2g/片"
            }
        ],
        "remark": "第N次就诊",
        "updated_time": "2018-08-01T00:12:07.911785+08:00"
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. allergic_history | string | ✅ | 过敏史 | |
| data. allergic_reaction | string | ✅ | 过敏反应 | |
| data. body_examination | string | ✅ | 体格检查 | |
| data. chief_complaint | string | ✅ | 主诉 | |
| data. clinic_examination_name | string | ✅ | 检查项目 | |
| data. clinic_laboratory_name | string | ✅ | 检验项目 | |
| data. clinic_treatment_name | string | ✅ | 治疗项目 | |
| data. clinic\_triage\_patient_id | string | ✅ | 就诊id | |
| data. created_time | string | ✅ | 创建时间 | |
| data. cure_suggestion | string | ✅ | 治疗建议 | |
| data. deleted_time | string | ✅ | 删除时间 | |
| data. diagnosis | string | ✅ | 诊断 | |
| data. family_medical_history | string | ✅ | 家族史 | |
| data. files | string | ✅ | 文件 | |
| data. history_of_past_illness | string | ✅ | 既往史 | |
| data. history_of_present_illness | string | ✅ | 现病史 | |
| data. immunizations | string | ✅ | 疫苗接种史 | |
| data. is_default | string | ✅ | 是否主诊断 | |
| data. medical_records | Array | ❌ | 病历续写（同主病历字段） | |
| data. morbidity_date | string | ✅ | 发病日期 | |
| data. operation_id | string | ✅ | 操作人id| |
| data. operation_name | string | ✅ | 操作人姓名| |
| data. personal_medical_history | string | ✅ | 个人史| |
| data. prescription_chinese_patient | Array | ❌ | 中药处方 （同 6.24 获取中药处方）| |
| data. prescription_western_patient | Array | ❌ | 中药处方 （6.20 获取西药处方）| |
--



</br>
<h3>6.35 分诊记录报告

```
请求地址：/triage/TriagePatientReport
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | int | ✅ |  就诊id | |



**应答包示例**

```
{
    "code": "200",
    "data": {
        "examination_results": [
            {
                "clinic_examination_id": 7,
                "clinic_examination_name": "心电图平板运动",
                "clinic_name": "龙华诊所",
                "clinic_triage_patient_id": 95,
                "conclusion_examination": "未见明显异常。",
                "examination_patient_id": 45,
                "examination_patient_record_id": 13,
                "order_doctor_name": "胡一天",
                "order_time": "2018-07-26T22:46:34.980163+08:00",
                "organ": null,
                "picture_examination": "[{\"docName\":\"安排-门诊安排.png\",\"url\":\"/uploads/安排-门诊安排.png\"},{\"docName\":\"安排-门诊安排.png\",\"url\":\"/uploads/安排-门诊安排.png\"}]",
                "report_doctor_name": "何炅",
                "report_time": "2018-07-26T23:14:05.648846+08:00",
                "result_examination": "健康"
            },
        ],
        "laboratory_results": [
            {
                "clinic_laboratory_id": 1,
                "clinic_laboratory_name": "血常规",
                "clinic_name": "龙华诊所",
                "clinic_triage_patient_id": 95,
                "laboratory_patient_id": 129,
                "laboratory_patient_record_id": 9,
                "laboratory_sample": "全血",
                "order_doctor_name": "胡一天",
                "order_time": "2018-07-26T22:47:23.487046+08:00",
                "remark": "注意补充营养啊   ",
                "report_doctor_name": "李维嘉",
                "report_time": "2018-07-26T23:17:15.943164+08:00"
            },
        ]
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. examination_results | array | ❌ | 检查结果 | |
| data. examination\_results. clinic\_examination_id | id | ❌ | 检查项目id | |
| data. examination\_results. clinic\_examination_name | string | ❌ | 检查项目名称 | |
| data. examination\_results. clinic_name | string | ❌ | 诊所名称 | |
| data. examination\_results. clinic\_triage_patient_id | id | ❌ | 就诊id | |
| data. examination\_results. conclusion_examination | string | ❌ | 检查结论 | |
| data. examination\_results. examination\_patient_id | id | ❌ | 检查id | |
| data. examination\_results. examination\_patient_record_id | id | ❌ | 检查记录id| |
| data. examination\_results. order\_doctor_name | string | ❌ | 开单医生 | |
| data. examination\_results. order\_time | time | ❌ | 开单时间 | |
| data. examination\_results. organ | string | ❌ | 检查部位 | |
| data. examination\_results. picture_examination | string | ❌ | 检查图片 | |
| data. examination\_results. report\_doctor_name | string | ❌ | 报告医生 | |
| data. examination\_results. report_time | time | ❌ | 报告时间 | |
| data. examination\_results. result_examination | string | ❌ | 检查结果| |
| data. laboratory_results | array | ❌ | 检验结果 | |
| data. laboratory\_results. clinic\_laboratory_id | int | ❌ | 检验项目id | |
| data. laboratory\_results. clinic\_laboratory_name | string | ❌ | 检验项目名称 | |
| data. laboratory\_results. clinic_name | string | ❌ | 诊所名称 | |
| data. laboratory\_results. clinic\_triage\_patient_id | int | ❌ | 就诊id| |
| data. laboratory\_results. laboratory\_patient_id | int | ❌ | 检验id| |
| data. laboratory\_results. laboratory\_patient\_record_id | int | ❌ | 检验记录id| |
| data. laboratory\_results. laboratory_sample | string | ❌ | 检验物| |
| data. laboratory\_results. order\_doctor_name | string | ❌ | 开单医生| |
| data. laboratory\_results. order_time | time | ❌ | 开单时间| |
| data. laboratory\_results. remark | string | ❌ | 备注| |
| data. laboratory\_results. report_doctor_name | string | ❌ | 报告医生| |
| data. laboratory\_results. report_time | time | ❌ | 报告时间| |
--


</br>
<h3>6.36 快速接诊

```
请求地址：/triage/QuickReception
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | int | ❌ |  患者id | |
| cert_no | int | ❌ | 身份证号 | |
| name | string | ✅ |  姓名 | |
| birthday | string | ✅ | 生日 | |
| sex | int | ✅ | 性别 | |
| phone | string | ✅ | 手机号 | |
| province | string | ❌ | 省 | |
| city | string | ❌ | 市 | |
| district | string | ❌ | 区 | |
| address | string | ❌ | 地址 | |
| clinic_id | int | ✅ | 诊所id | |
| visit_type | int | ✅ | 就诊类型 | |
| personnel_id | int | ✅ | 医生id | |
| department_id | int | ✅ | 科室id | |



**应答包示例**

```
{
    "code": "200",
    "msg": "ok",
    data: {
    	patient_id: "患者id",
    	clinic_triage_patient_id: "就诊id",
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |

--

7 诊疗模块
--------

</br>
<h3>7.1 创建诊疗缴费项目

```
请求地址：/diagnosisTreatment/create
```
**请求包示例**

```
{
	clinic_id:1
	name:紫外线治疗
	en_name:
	price:100
	cost:20
	status:true
	is_discount:false
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所id| |
| name | String | ✅ | 诊疗名称 | |
| en_name | String | ❌  |  英文名称 | |
| price | Int | ✅ |  零售价 | |
| cost | Int | ❌ |  成本价 | |
| status | Boolean | ❌  |  是否启用 | |
| is_discount | Boolean | ❌  |  是否折扣 | |

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
<h3>7.2 更新诊疗缴费项目

```
请求地址：/diagnosisTreatment/update
```
**请求包示例**

```
{
	clinic_diagnosis_treatment_id:1
	name:紫外线治疗
	en_name:
	price:100
	cost:20
	status:true
	is_discount:false
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_diagnosis_treatment_id | Int | ✅ |  诊疗id| |
| name | String | ✅ | 诊疗名称 | |
| en_name | String | ❌  |  英文名称 | |
| price | Int | ✅ |  零售价 | |
| cost | Int | ❌ |  成本价 | |
| status | Boolean | ❌  |  是否启用 | |
| is_discount | Boolean | ❌  |  是否折扣 | |

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
<h3>7.3 启用和停用诊疗项目

```
请求地址：/diagnosisTreatment/onOff
```
**请求包示例**

```
{
	status:true
	clinic_id:8
	clinic_diagnosis_treatment_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| status | Boolean | ✅ |  是否启用 | |
| clinic_id | Int | ✅ |  诊所id| |
| clinic_diagnosis_treatment_id | Int | ✅ |  诊疗id| |
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
<h3>7.4 诊疗项目列表

```
请求地址：/diagnosisTreatment/list
```
**请求包示例**

```
{
	clinic_id:1,
	keyword:,
	status:,
	offset: 0,
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字| |
| clinic_id | int | ✅ |  诊所id | |
| status | Boolean | ❌ |  是否启用 | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_diagnosis_treatment_id": 2,
      "cost": 1200,
      "en_name": null,
      "is_discount": false,
      "name": "主治医生诊疗费",
      "price": 10000,
      "status": true
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 5
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.clinic_diagnosis_treatment_id | int | ✅ |  诊疗id | |
| data.cost | Int | ❌ | 成本价 | |
| data.en_name | String | ❌ |  英文名称| |
| data.is_discount | Boolean | ✅ |  是否折扣 | |
| data.name | String | ✅ |  诊疗名称 | |
| data.price | int | ✅ |  零售价 | |
| data. status | bolean | ✅ |  是否启用 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |

--

</br>
<h3>7.5 诊疗项目详情

```
请求地址：/diagnosisTreatment/detail
```
**请求包示例**

```
{
	clinic_diagnosis_treatment_id:2
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_diagnosis_treatment_id | Int | ✅ |  诊疗id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "clinic_diagnosis_treatment_id": 1,
    "cost": null,
    "en_name": null,
    "is_discount": true,
    "name": "诊疗费",
    "price": 1500,
    "status": true
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.clinic_diagnosis_treatment_id | int | ✅ |  诊疗id | |
| data.cost | Int | ❌ | 成本价 | |
| data.en_name | String | ❌ |  英文名称| |
| data.is_discount | Boolean | ✅ |  是否折扣 | |
| data.name | String | ✅ |  诊疗名称 | |
| data.price | int | ✅ |  零售价 | |
| data. status | bolean | ✅ |  是否启用 | |
--

14 管理用户模块
--------

</br>
<h3>14.1 诊所用户登录

```
请求地址：/admin/login
```
**请求包示例**

```
{
	username:pt_admin
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
    "id": 1,
    "is_clinic_admin": true,
    "name": "平台管理员",
    "phone": "13211112222",
    "title": "平台经理",
    "username": "pt_admin"
  },
  "login_times": 19,
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.id | String | ✅ |  管理人员id| |
| data.is_clinic_admin | String | ✅ |  是否超级管理员| |
| data.name | String | ✅ |  登录人员名称| |
| data.username | String | ✅ |  登录账号| |
| login_times | Int | ✅ |  登录次数 | |
| msg | String | ✅ |  返回码， 200 成功| |
--

</br>
<h3>14.2 平台账号添加

```
请求地址：/admin/create
```
**请求包示例**

```
{
	name:平台管理员
	title:平台经理
	phone:13211112222
	username:pt_amdin
	password:123456
	items:[{"function_menu_id":"1"},{"function_menu_id":"2"}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| name | String | ✅ | 姓名 | |
| title | String | ✅ |  职务 | |
| phone | String | ✅ |  平台账号预留手机号码 | |
| username | String | ✅ |  超级管理员账号 | |
| password | String | ✅ |  密码 | |
| items | Array | ❌ |  预留手机号码 | |
| items.function_menu_id | String | ✅ |  菜单功能项id | |

**应答包示例**

```
{
    "code": "200",
    "data": 1
    "msg":"ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Int | ✅ |  返回管理用户id | |
| msg | String | ✅ |  返回信息 | |

--

</br>
<h3>14.3 获取诊所列表

```
请求地址：/admin/list
```
**请求包示例**

```
{
	keyword:龙
	start_date:
	end_date:
	status:
	offset:0
	limit:10
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
      "admin_id": 2,
      "created_time": "2018-08-04T17:23:19.638049+08:00",
      "name": "深圳龙华店",
      "phone": "15210121021",
      "status": true,
      "title": "店长",
      "username": "lh001"
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
| data.items.admin_id | Int | ✅ |  管理员id| |
| data.items.created_time | Date | ✅ |  创建时间| |
| data.items.name | String | ✅ |  管理员名称| |
| data.items.phone | String | ✅ |  预留手机号码| |
| data.items.status | Boolean | ✅ |  是否启用| |
| data.items.title | String | ✅ |  管理员职称| |
| data.items.username | String | ✅ |  管理员账号| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>14.4 平台账号修改

```
请求地址：/admin/update
```
**请求包示例**

```
{
	admin_id:1
	name:平台管理员
	title:平台经理
	phone:13211112222
	username:pt_amdin
	password:123456
	items:[{"function_menu_id":"1"},{"function_menu_id":"2"}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| admin_id | Int | ✅ | 管理员id | |
| name | String | ✅ | 姓名 | |
| title | String | ✅ |  职务 | |
| phone | String | ✅ |  平台账号预留手机号码 | |
| username | String | ✅ |  超级管理员账号 | |
| password | String | ✅ |  密码 | |
| items | Array | ❌ |  预留手机号码 | |
| items.function_menu_id | String | ✅ |  菜单功能项id | |

**应答包示例**

```
{
    "code": "200",
    "data": 1
    "msg":"ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Int | ✅ |  返回管理用户id | |
| msg | String | ✅ |  返回信息 | |

--

</br>
<h3>14.5 获取平台账号信息

```
请求地址：/admin/getByID
```
**请求包示例**

```
{
	admin_id:2
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| admin_id | Int | ✅ |  管理员id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "admin_id": 1,
    "created_time": "2018-08-04T16:26:13.950835+08:00",
    "funtionMenus": [
      {
        "ascription": "02",
        "function_menu_id": 5,
        "icon": null,
        "level": 0,
        "menu_name": "平台管理",
        "menu_url": "/platform",
        "parent_function_menu_id": null,
        "status": true,
        "weight": 4
      },
		...
    ],
    "name": "平台管理员",
    "phone": "13211112222",
    "status": true,
    "title": "平台经理",
    "username": "pt_admin"
  },
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.admin_id | Int | ✅ |  管理员id| |
| data.created_time | Date | ✅ |  创建时间| |
| data.name | String | ✅ |  管理员名称| |
| data.phone | String | ✅ |  预留手机号码| |
| data.status | Boolean | ✅ |  是否启用| |
| data.title | String | ✅ |  管理员职称| |
| data.username | String | ✅ |  管理员账号| |
| data.funtionMenus | Array | ❌ |  管理员账号菜单项| |
| data.funtionMenus.ascription | String | ✅ | 菜单所属类型 01 诊所 02 平台| |
| data.funtionMenus.function_menu_id | Int | ✅ | 菜单项id| |
| data.funtionMenus.icon | String | ❌ | 菜单图标| |
| data.funtionMenus.level | Int | ✅ | 菜单项等级| |
| data.funtionMenus.menu_name | String | ✅ | 菜单名称| |
| data.funtionMenus.menu_url | String | ✅ | 菜单路由| |
| data.funtionMenus.parent_function_menu_id | Int | ❌ | 父级菜单id| |
| data.funtionMenus.status | Boolean | ✅ | 菜单是否启用| |
| data.funtionMenus.weight | Int | ✅ | 菜单权重| |
--

</br>
<h3>14.6 启用和停用管理员账号

```
请求地址：/admin/onOff
```
**请求包示例**

```
{
	status:true
	admin_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| status | Boolean | ✅ |  是否启用 | |
| admin_id | Int | ✅ |  管理员id| |
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
<h3>14.7 获取平台未开通的菜单项

```
请求地址：/admin/menubarUnset
```
**请求包示例**

```
{
	admin_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| admin_id | Int | ✅ |  管理员id| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "ascription": "02",
      "function_menu_id": 49,
      "icon": "/static/icons/business.svg",
      "level": 1,
      "menu_name": "运营分析",
      "menu_url": "/platform/operation/totalAmount",
      "parent_function_menu_id": 5,
      "status": true,
      "weight": 3
    }
  ],
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.ascription | String | ✅ | 菜单所属类型 01 诊所 02 平台| |
| data.function_menu_id | Int | ✅ | 菜单项id| |
| data.icon | String | ❌ | 菜单图标| |
| data.level | Int | ✅ | 菜单项等级| |
| data.menu_name | String | ✅ | 菜单名称| |
| data.menu_url | String | ✅ | 菜单路由| |
| data.parent_function_menu_id | Int | ❌ | 父级菜单id| |
| data.status | Boolean | ✅ | 菜单是否启用| |
| data.weight | Int | ✅ | 菜单权重| |
| msg | String | ✅ | 文本信息| |
--

</br>
<h3>14.8 获取平台开通菜单项

```
请求地址：/admin/menubarList
```
**请求包示例**

```
{
	admin_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| admin_id | Int | ✅ |  管理员id| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "ascription": "02",
      "function_menu_id": 5,
      "icon": null,
      "level": 0,
      "menu_name": "平台管理",
      "menu_url": "/platform",
      "parent_function_menu_id": null,
      "status": true,
      "weight": 4
    },
    {
      "ascription": "02",
      "function_menu_id": 29,
      "icon": "/static/icons/clinic.svg",
      "level": 1,
      "menu_name": "诊所管理",
      "menu_url": "/platform/clinique/add",
      "parent_function_menu_id": 5,
      "status": true,
      "weight": 0
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
| data | Object | ❌ |  返回信息 | |
| data.ascription | String | ✅ | 菜单所属类型 01 诊所 02 平台| |
| data.function_menu_id | Int | ✅ | 菜单项id| |
| data.icon | String | ❌ | 菜单图标| |
| data.level | Int | ✅ | 菜单项等级| |
| data.menu_name | String | ✅ | 菜单名称| |
| data.menu_url | String | ✅ | 菜单路由| |
| data.parent_function_menu_id | Int | ❌ | 父级菜单id| |
| data.status | Boolean | ✅ | 菜单是否启用| |
| data.weight | Int | ✅ | 菜单权重| |
| msg | String | ✅ | 文本信息| |
--

15 诊断模块
--------

</br>
<h3>15.1 创建诊断

```
请求地址：/diagnosis/create
```
**请求包示例**

```
{
	py_code:BBNZ
	name:鼻部囊肿
	icd_code:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| py_code | String | ✅ |  拼音编码| |
| name | String | ✅ | 诊断名称 | |
| icd_code | String | ❌  |  国际编码 | |

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
| data | Int | ✅ |  诊断id | |

--

16 病历模块
--------

</br>
<h3>16.1 创建主病历

```
请求地址：/medicalRecord/upsert
```
**请求包示例**

```
{
	clinic_triage_patient_id:2
	chief_complaint:主病历内容
	operation_id:6
	morbidity_date:2018-07-21
	personal_medical_history:个人病史
	history_of_present_illness:现病史
	history_of_past_illness:既往史
	family_medical_history:家族史
	allergic_history:过敏史
	allergic_reaction:过敏反应
	immunizations:疫苗接种史
	body_examination:体格检查
	diagnosis:诊断
	cure_suggestion:治疗建议
	remark:备注
	files:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_triage_patient_id | Int | ✅ |  就诊id| |
| chief_complaint | String | ✅ | 主诉 | |
| operation_id | String |  ✅  |  操作人id | |
| morbidity_date | String | ❌  |  国际编码 | |
| personal_medical_history | String | ❌  |  个人病史 | |
| history_of_present_illness | String | ❌  | 现病史 | |
| history_of_past_illness | String | ❌  |  既往史 | |
| family_medical_history | String | ❌  |  家族史 | |
| allergic_history | String | ❌  | 过敏史 | |
| allergic_reaction | String | ❌  |  过敏反应 | |
| immunizations | String | ❌  |  疫苗接种史 | |
| body_examination | String | ❌  |  体格检查 | |
| diagnosis | String | ❌  |  诊断 | |
| cure_suggestion | String | ❌  |  治疗建议 | |
| remark | String | ❌  |  备注| |
| files | String | ❌  |  上传的文件 | |

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
<h3>16.2 续写病历

```
请求地址：/medicalRecord/renew
```
**请求包示例**

```
{
	clinic_triage_patient_id:2
	chief_complaint:续写病历内容
	operation_id:6
	files:上传的文件
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_triage_patient_id | Int | ✅ |  就诊id| |
| chief_complaint | String | ✅ | 续写病历内容 | |
| operation_id | String |  ✅  |  操作人id | |
| files | String | ❌  |  上传的文件 | |

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
<h3>16.3 续写病历修改

```
请求地址：/medicalRecord/MedicalRecordRenewUpdate
```
**请求包示例**

```
{
	medical_record_id:1
	clinic_triage_patient_id:2
	chief_complaint:续写病历内容
	operation_id:6
	files:上传的文件
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| medical_record_id | Int | ✅ |  病历id| |
| chief_complaint | String | ✅ | 续写病历内容 | |
| operation_id | String |  ✅  |  操作人id | |
| files | String | ❌  |  上传的文件 | |

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
<h3>16.4 续写病历删除

```
请求地址：/medicalRecord/MedicalRecordRenewDelete
```
**请求包示例**

```
{
	medical_record_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| medical_record_id | Int | ✅ |  病历id| |

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
<h3>16.5 通过就诊id查找病历

```
请求地址：/medicalRecord/findByTriageId
```
**请求包示例**

```
{
	clinic_triage_patient_id:2
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_triage_patient_id | Int | ✅ |  就诊id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "allergic_history": "胜多负少的范德萨范德萨发",
    "allergic_reaction": "胜多负少的个",
    "body_examination": "发的共同体",
    "chief_complaint": "主病历内容",
    "clinic_triage_patient_id": 2,
    "created_time": "2018-05-31T22:07:42.220501+08:00",
    "cure_suggestion": "共同体热辅导班",
    "deleted_time": null,
    "diagnosis": "耳热是从v",
    "family_medical_history": "是否水电费",
    "files": "",
    "history_of_past_illness": "水电费水电费水电费",
    "history_of_present_illness": "辅导费第三方支付",
    "id": 3,
    "immunizations": "覆盖掉刚刚有个",
    "is_default": true,
    "morbidity_date": "2018-07-21",
    "operation_id": 6,
    "personal_medical_history": null,
    "remark": "地方也个人",
    "updated_time": "2018-07-25T23:25:59.98118+08:00"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.clinic_triage_patient_id | Int | ✅ |  就诊id| |
| data.id | Int | ✅ |  病历id| |
| data.chief_complaint | String | ✅ | 主诉 | |
| data.operation_id | String |  ✅  |  操作人id | |
| data.morbidity_date | String | ❌  |  国际编码 | |
| data.personal_medical_history | String | ❌  |  个人病史 | |
| data.history_of_present_illness | String | ❌  | 现病史 | |
| data.history_of_past_illness | String | ❌  |  既往史 | |
| data.family_medical_history | String | ❌  |  家族史 | |
| data.allergic_history | String | ❌  | 过敏史 | |
| data.allergic_reaction | String | ❌  |  过敏反应 | |
| data.immunizations | String | ❌  |  疫苗接种史 | |
| data.body_examination | String | ❌  |  体格检查 | |
| data.diagnosis | String | ❌  |  诊断 | |
| data.cure_suggestion | String | ❌  |  治疗建议 | |
| data.remark | String | ❌  |  备注| |
| data.files | String | ❌  |  上传的文件 | |
| data.is_default | Boolean | ✅  |  是否是主病历 | |
| data.updated_time | time | ✅ | 修改时间 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.deleted_time | time |  ❌ | 删除时间 | |
--

</br>
<h3>16.6 创建病历模板

```
请求地址：/medicalRecord/model/create
```
**请求包示例**

```
{
	model_name:模板名称
	is_common:true
	chief_complaint:主病历内容
	operation_id:6
	morbidity_date:2018-07-21
	personal_medical_history:个人病史
	history_of_present_illness:现病史
	history_of_past_illness:既往史
	family_medical_history:家族史
	allergic_history:过敏史
	allergic_reaction:过敏反应
	immunizations:疫苗接种史
	body_examination:体格检查
	diagnosis:诊断
	cure_suggestion:治疗建议
	remark:备注
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| model_name | String | ✅ | 模板名称| |
| is_common | Boolean | ✅ | 是否通用 | |
| chief_complaint | String | ❌ | 主诉 | |
| operation_id | String |  ✅  |  操作人id | |
| morbidity_date | String | ❌  |  国际编码 | |
| personal_medical_history | String | ❌  |  个人病史 | |
| history_of_present_illness | String | ❌  | 现病史 | |
| history_of_past_illness | String | ❌  |  既往史 | |
| family_medical_history | String | ❌  |  家族史 | |
| allergic_history | String | ❌  | 过敏史 | |
| allergic_reaction | String | ❌  |  过敏反应 | |
| immunizations | String | ❌  |  疫苗接种史 | |
| body_examination | String | ❌  |  体格检查 | |
| diagnosis | String | ❌  |  诊断 | |
| cure_suggestion | String | ❌  |  治疗建议 | |
| remark | String | ❌  |  备注| |

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
<h3>16.7 修改病历模板

```
请求地址：/medicalRecord/model/update
```
**请求包示例**

```
{
	medical_record_model_id:1
	model_name:模板名称
	is_common:true
	chief_complaint:主病历内容
	operation_id:6
	morbidity_date:2018-07-21
	personal_medical_history:个人病史
	history_of_present_illness:现病史
	history_of_past_illness:既往史
	family_medical_history:家族史
	allergic_history:过敏史
	allergic_reaction:过敏反应
	immunizations:疫苗接种史
	body_examination:体格检查
	diagnosis:诊断
	cure_suggestion:治疗建议
	remark:备注
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| medical_record_model_id | Int | ✅ | 模板名称| |
| model_name | String | ✅ | 模板名称| |
| is_common | Boolean | ✅ | 是否通用 | |
| chief_complaint | String | ❌ | 主诉 | |
| operation_id | String |  ✅  |  操作人id | |
| morbidity_date | String | ❌  |  国际编码 | |
| personal_medical_history | String | ❌  |  个人病史 | |
| history_of_present_illness | String | ❌  | 现病史 | |
| history_of_past_illness | String | ❌  |  既往史 | |
| family_medical_history | String | ❌  |  家族史 | |
| allergic_history | String | ❌  | 过敏史 | |
| allergic_reaction | String | ❌  |  过敏反应 | |
| immunizations | String | ❌  |  疫苗接种史 | |
| body_examination | String | ❌  |  体格检查 | |
| diagnosis | String | ❌  |  诊断 | |
| cure_suggestion | String | ❌  |  治疗建议 | |
| remark | String | ❌  |  备注| |

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
<h3>16.8 删除病历模板

```
请求地址：/medicalRecord/model/delete
```
**请求包示例**

```
{
	medical_record_model_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| medical_record_model_id | Int | ✅ | 模板名称| |

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
<h3>16.9 获取患者病历列表

```
请求地址：/medicalRecord/listByPid
```
**请求包示例**

```
{
	patient_id:2
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| patient_id | int | ✅ |  患者id | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "allergic_history": "",
      "allergic_reaction": "",
      "body_examination": "",
      "chief_complaint": "发热",
      "clinic_name": "龙华诊所",
      "clinic_patient_id": 1,
      "clinic_triage_patient_id": 1,
      "created_time": "2018-06-05T23:08:54.206015+08:00",
      "cure_suggestion": "",
      "deleted_time": null,
      "department_name": "骨科",
      "diagnosis": "",
      "doctor_name": "扁鹊",
      "family_medical_history": "",
      "files": "[{\"docName\":\"微信图片_20180425210127.jpg\",\"url\":\"/uploads/微信图片_20180425210127.jpg\"}]",
      "history_of_past_illness": "aaaa",
      "history_of_present_illness": "现病",
      "id": 9,
      "immunizations": "",
      "is_default": true,
      "morbidity_date": "",
      "operation_id": 1,
      "personal_medical_history": null,
      "registion_time": "2018-05-28T00:26:29.012104+08:00",
      "remark": "",
      "updated_time": "2018-06-07T23:03:27.579995+08:00",
      "visit_type": 1
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
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.clinic_triage_patient_id | Int | ✅ |  就诊id| |
| data.id | Int | ✅ |  病历id| |
| data.chief_complaint | String | ✅ | 主诉 | |
| data.clinic_name | String | ✅  |  诊所名称 | |
| data.clinic_patient_id | Int | ✅  |  患者诊所id | |
| data.department_name | String | ✅  |  科室名称 | |
| data.doctor_name | String | ✅  |  医生名称 | |
| data.operation_id | String |  ✅  |  操作人id | |
| data.morbidity_date | String | ❌  |  国际编码 | |
| data.personal_medical_history | String | ❌  |  个人病史 | |
| data.history_of_present_illness | String | ❌  | 现病史 | |
| data.history_of_past_illness | String | ❌  |  既往史 | |
| data.family_medical_history | String | ❌  |  家族史 | |
| data.allergic_history | String | ❌  | 过敏史 | |
| data.allergic_reaction | String | ❌  |  过敏反应 | |
| data.immunizations | String | ❌  |  疫苗接种史 | |
| data.body_examination | String | ❌  |  体格检查 | |
| data.diagnosis | String | ❌  |  诊断 | |
| data.cure_suggestion | String | ❌  |  治疗建议 | |
| data.remark | String | ❌  |  备注| |
| data.files | String | ❌  |  上传的文件 | |
| data.is_default | Boolean | ✅  |  是否是主病历 | |
| data.visit_type | Int | ✅  |  出诊类型 1: 首诊， 2复诊，3：术后复诊 | |
| data.registion_time | time | ✅ | 等级时间 | |
| data.updated_time | time | ✅ | 修改时间 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.deleted_time | time |  ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |

--

</br>
<h3>16.10 查询模板列表

```
请求地址：/medicalRecord/model/list
```
**请求包示例**

```
{
	keyword:模板
	is_common:
	operation_id:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| is_common | Boolean | ❌ |  是否通用 | |
| operation_id | int | ❌ |  创建人员id | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "allergic_history": "鸡蛋、西红柿、海鲜、money",
      "allergic_reaction": "皮肤瘙痒、红肿",
      "body_examination": "体温36.5℃脉搏100次/分呼吸18次/分血压190/90mmHg发育正常",
      "chief_complaint": "出生并长于原籍，居住及生活环境良好。无酗酒、吸烟、吸毒等不良嗜好。否认到过传染病、地方病流行地区。",
      "created_time": "2018-08-01T10:01:31.853529+08:00",
      "cure_suggestion": "建议最好带宝宝到医院检查，根据病情，确定治疗方案，同时给宝宝暂停母乳，多喝水，并注意观察黄疸值变化。",
      "deleted_time": null,
      "diagnosis": "扁桃体和腺样体肥大",
      "family_medical_history": "家中无遗传病病史。",
      "history_of_past_illness": "糖尿病 高血压 青光眼 颜值癌",
      "history_of_present_illness": "糖尿病 高血压 青光眼 颜值癌",
      "id": 12,
      "immunizations": "天花",
      "is_common": true,
      "model_name": "8月1日模板",
      "operation_id": 20,
      "operation_name": "胡一天",
      "personal_medical_history": "生长于香港。文盲。否认外地长期居住史。无疫区、疫水接触史。否认工业毒物、粉尘及放射性物质接触史。否认牧区、矿山、高氟区、低碘区居住史。平日生活规律，否认吸毒史。否认吸烟嗜好。否认饮酒嗜好。否认冶游史。第N+1次\n",
      "remark": "第N +1次就诊",
      "updated_time": "2018-08-01T10:01:31.853529+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 7
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.clinic_triage_patient_id | Int | ✅ |  就诊id| |
| data.id | Int | ✅ |  病历id| |
| data.chief_complaint | String | ✅ | 主诉 | |
| data.operation_id | Int |  ✅  |  操作人id | |
| data.operation_name | String |  ✅  |  操作人名称 | |
| data.morbidity_date | String | ❌  |  国际编码 | |
| data.personal_medical_history | String | ❌  |  个人病史 | |
| data.history_of_present_illness | String | ❌  | 现病史 | |
| data.history_of_past_illness | String | ❌  |  既往史 | |
| data.family_medical_history | String | ❌  |  家族史 | |
| data.allergic_history | String | ❌  | 过敏史 | |
| data.allergic_reaction | String | ❌  |  过敏反应 | |
| data.immunizations | String | ❌  |  疫苗接种史 | |
| data.body_examination | String | ❌  |  体格检查 | |
| data.diagnosis | String | ❌  |  诊断 | |
| data.cure_suggestion | String | ❌  |  治疗建议 | |
| data.remark | String | ❌  |  备注| |
| data.files | String | ❌  |  上传的文件 | |
| data.is_default | Boolean | ✅  |  是否是主病历 | |
| data.updated_time | time | ✅ | 修改时间 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.deleted_time | time |  ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--