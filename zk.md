云诊所接口文档
===========

**创建时间：2018-08-16**

修改记录
--------
| 修定日期 | 修改内容 | 修改人 | 
| :-: | :-: | :-:  | 

接口列表
--------


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




4 医生排班
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
