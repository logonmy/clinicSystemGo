云诊所接口文档
===========

**创建时间：2018-08-16**

修改记录
--------
| 修定日期 | 修改内容 | 修改人 | 
| :-: | :-: | :-:  | 

接口列表
--------


18 治疗收费项目模块
--------

</br>
<h3>18.1 创建治疗缴费项目

```
请求地址：/treatment/create
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
| name | string | ✅ |  项目名称 | |
| en_name | string | ❌ |  英文名称 | |
| py_code | string | ❌ |  拼音码 | |
| idc_code | string | ❌ |  国际编码 | |
| unit_name | string | ✅ |  单位名称 | |
| remark | string | ❌ |  备注| |
| price | int | ✅ |  单价| |
| cost | int | ❌ |  成本价 | |
| status | boolean | ❌ |  是否启用 | |
| is_discount | boolean | ❌ |  是否可打折 | |

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
--


</br>
<h3>18.2 更新治疗项目

```
请求地址：/treatment/update
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_treatment_id | int | ✅ |  治疗项目id| |
| name | string | ✅ |  项目名称 | |
| en_name | string | ❌ |  英文名称 | |
| py_code | string | ❌ |  拼音码 | |
| idc_code | string | ❌ |  国际编码 | |
| unit_name | string | ❌ |  单位名称 | |
| remark | string | ❌ |  备注| |
| price | int | ✅ |  单价| |
| cost | int | ❌ |  成本价 | |
| status | boolean | ❌ |  是否启用 | |
| is_discount | boolean | ❌ |  是否可打折 | |

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
--



</br>
<h3>18.3 启用和停用

```
请求地址：/treatment/onOff
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
| clinic_treatment_id | int | ✅ |  治疗项目id| |
| status | boolean | ✅ |  是否启用 | |

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
--


</br>
<h3>18.4 治疗缴费项目列表

```
请求地址：/treatment/list
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
| keyword | string | ❌ |  搜索关键字| |
| status | boolean | ❌ |  是否启用 | |
| offset | int | ❌ |  开始条数| |
| limit | int | ❌ |  条数| |

**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "clinic_treatment_id": 6,
            "cost": 2000,
            "discount_price": 0,
            "en_name": null,
            "idc_code": null,
            "is_discount": false,
            "price": 10000,
            "py_code": "zwxzl",
            "remark": null,
            "status": true,
            "treatment_name": "紫外线治疗",
            "unit_name": "次"
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 12
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ |  | |
| data. clinic\_treatment_id | int | ✅ | 项目id | |
| data. cost | int | ❌ | 成本价 | |
| data. discount_price | ❌ | ✅ | 折扣价 | |
| data. en_name | string | ❌ | 英文名称 | |
| data. idc_code | string | ❌ | 国际编码 | |
| data. is_discount | booleam | ✅ | 是否可折扣 | |
| data. price | int | ✅ | 零售价 | |
| data. py_code | string | ❌ | 拼音码 | |
| data. remark | string | ❌ | 备注 | |
| data. status | boolean | ✅ | 是否启用 | |
| data. treatment_name | string | ✅ | 项目名称 | |
| data. unit_name | int | ✅ | 单位名称 | |
--


</br>
<h3>18.5 治疗项目详情

```
请求地址：/treatment/detail
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_treatment_id | int | ✅ |  项目id | |

**应答包示例**

```
{
    "code": "200",
    "data": {
        "clinic_treatment_id": 6,
        "cost": 2000,
        "discount_price": 0,
        "en_name": null,
        "idc_code": null,
        "is_discount": false,
        "name": "紫外线治疗",
        "price": 10000,
        "py_code": "zwxzl",
        "remark": null,
        "status": true,
        "unit_name": "次"
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. clinic\_treatment_id | int | ✅ | 项目id | |
| data. cost | int | ❌ | 成本价 | |
| data. discount_price | ❌ | ✅ | 折扣价 | |
| data. en_name | string | ❌ | 英文名称 | |
| data. idc_code | string | ❌ | 国际编码 | |
| data. is_discount | booleam | ✅ | 是否可折扣 | |
| data. price | int | ✅ | 零售价 | |
| data. py_code | string | ❌ | 拼音码 | |
| data. remark | string | ❌ | 备注 | |
| data. status | boolean | ✅ | 是否启用 | |
| data. treatment_name | string | ✅ | 项目名称 | |
| data. unit_name | int | ✅ | 单位名称 | |
--



</br>
<h3>18.6 创建治疗医嘱模板

```
请求地址：/treatment/TreatmentPatientModelCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| model_name | string | ✅ |  模板名称 | |
| is_common | booleam | ✅ |  通用 或 个人| |
| operation_id | int | ✅ |  操作人id | |
| items | string | ✅ |  详细项目 json 字符串 | |
| items. clinic_treatment_id | string | ✅ | 项目id | |
| items. times | string | ✅ | 次数 | |
| items. illustration | string | ❌ | 说明 | |

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
<h3>18.7 查询治疗医嘱模板

```
请求地址：/treatment/TreatmentPatientModelList
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | string | ❌ |  关键字| |
| is_common | booleam | ❌ |  通用 或 个人| |
| operation_id | int | ❌ |  操作人id | |
| offset | int | ❌ |  开始条数| |
| limit | int | ❌ |  条数| |

**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "model_name": "kid模板",
            "treatment_patient_model_id": 4,
            "operation_name": "超级管理员",
            "is_common": true,
            "created_time": "2018-06-27T23:52:15.79871+08:00",
            "items": [
                {
                    "treatment_name": "针灸",
                    "clinic_treatment_id": 4,
                    "illustration": "一定有很多要求，我不说而已，你说了，你说了",
                    "times": 3,
                    "unit_name": "次"
                }
            ]
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 3
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ |  | |
| data. model_name | string | ✅ | 模板名称 | |
| data. treatment\_patient\_model_id | int | ✅ | 模板id | |
| data. operation_name | string | ✅ | 创建人 | |
| data. is_common | boolean | ✅ | 通用 个人 | |
| data. created_time | time | ✅ | 创建时间 | |
| data. items | array | ✅ | | |
| data. items. treatment_name | string | ✅ | 项目名称 | |
| data. items. clinic_treatment_id | int | ✅ | 项目id | |
| data. items. illustration | string | ✅ | 说明 | |
| data. items. times | int | ✅ | 次数 | |
| data. items. unit_name | int | ✅ | 单位名称 | |
--


</br>
<h3>18.8 查询治疗医嘱模板详情

```
请求地址：/treatment/TreatmentPatientModelDetail
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| treatment_patient_model_id | int | ❌ |  模板id| |

**应答包示例**

```
{
    "code": "200",
    "data": {
        "is_common": false,
        "items": [
            {
                "clinic_treatment_id": 1,
                "illustration": "哈哈哈哈倾世",
                "name": "打针",
                "times": 1
            }
        ],
        "model_name": "测试",
        "status": true,
        "treatment_patient_model_id": 6
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. model_name | string | ✅ | 模板名称 | |
| data. treatment\_patient\_model_id | int | ✅ | 模板id | |
| data. operation_name | string | ✅ | 创建人 | |
| data. is_common | boolean | ✅ | 通用 个人 | |
| data. created_time | time | ✅ | 创建时间 | |
| data. items | array | ✅ | | |
| data. items. treatment_name | string | ✅ | 项目名称 | |
| data. items. clinic_treatment_id | int | ✅ | 项目id | |
| data. items. illustration | string | ✅ | 说明 | |
| data. items. times | int | ✅ | 次数 | |
| data. items. unit_name | int | ✅ | 单位名称 | |
--



</br>
<h3>18.9 修改治疗医嘱模板

```
请求地址：/treatment/TreatmentPatientModelUpdate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| treatment_patient_model_id | string | ✅ |  模板id | |
| model_name | string | ✅ |  模板名称 | |
| is_common | booleam | ✅ |  通用 或 个人| |
| operation_id | int | ✅ |  操作人id | |
| items | string | ✅ |  详细项目 json 字符串 | |
| items. clinic_treatment_id | string | ✅ | 项目id | |
| items. times | string | ✅ | 次数 | |
| items. illustration | string | ❌ | 说明 | |

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
<h3>18.20 删除治疗医嘱模板

```
请求地址：/treatment/TreatmentPatientModelDelete
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| treatment_patient_model_id | string | ✅ |  模板id | |

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




19 其他费用收费项目模块
--------

</br>
<h3>19.1 创建治疗缴费项目

```
请求地址：/otherCost/create
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
| name | string | ✅ |  项目名称 | |
| en_name | string | ❌ |  英文名称 | |
| py_code | string | ❌ |  拼音码 | |
| idc_code | string | ❌ |  国际编码 | |
| unit_name | string | ✅ |  单位名称 | |
| remark | string | ❌ |  备注| |
| price | int | ✅ |  单价| |
| cost | int | ❌ |  成本价 | |
| status | boolean | ❌ |  是否启用 | |
| is_discount | boolean | ❌ |  是否可打折 | |

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
--


</br>
<h3>19.2 更新其它费用项目

```
请求地址：/otherCost/update
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_other_cost_id | int | ✅ |  项目id | |
| name | string | ✅ |  项目名称 | |
| en_name | string | ❌ |  英文名称 | |
| py_code | string | ❌ |  拼音码 | |
| idc_code | string | ❌ |  国际编码 | |
| unit_name | string | ✅ |  单位名称 | |
| remark | string | ❌ |  备注| |
| price | int | ✅ |  单价| |
| cost | int | ❌ |  成本价 | |
| status | boolean | ❌ |  是否启用 | |
| is_discount | boolean | ❌ |  是否可打折 | |

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
--



</br>
<h3>19.3 启用和停用

```
请求地址：/otherCost/onOff
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
| clinic_other_cost_id | int | ✅ |  项目id| |
| status | boolean | ✅ |  是否启用 | |

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
--



</br>
<h3>19.4 其它费用缴费项目列表

```
请求地址：/otherCost/list
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
| keyword | string | ❌ |  搜索关键字| |
| status | boolean | ❌ |  是否启用 | |
| offset | int | ❌ |  开始条数| |
| limit | int | ❌ |  条数| |

**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "clinic_other_cost_id": 1,
            "cost": 1000,
            "discount_price": 0,
            "en_name": null,
            "is_discount": false,
            "name": "主任挂号费",
            "price": 10000,
            "py_code": "ZRGHF",
            "remark": "阿萨德噶发",
            "status": true,
            "unit_name": "次"
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 4
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ |  | |
| data. clinic_other_cost_id | int | ✅ | 项目id | |
| data. cost | int | ❌ | 成本价 | |
| data. discount_price | ❌ | ✅ | 折扣价 | |
| data. en_name | string | ❌ | 英文名称 | |
| data. idc_code | string | ❌ | 国际编码 | |
| data. is_discount | booleam | ✅ | 是否可折扣 | |
| data. price | int | ✅ | 零售价 | |
| data. py_code | string | ❌ | 拼音码 | |
| data. remark | string | ❌ | 备注 | |
| data. status | boolean | ✅ | 是否启用 | |
| data. name | string | ✅ | 项目名称 | |
| data. unit_name | int | ✅ | 单位名称 | |
--


</br>
<h3>19.5 其它费用项目详情

```
请求地址：/otherCost/detail
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_other_cost_id | int | ✅ |  项目id | |

**应答包示例**

```
{
    "code": "200",
    "data": {
        "clinic_other_cost_id": 1,
        "cost": 1000,
        "discount_price": 0,
        "en_name": null,
        "is_discount": false,
        "name": "主任挂号费",
        "price": 10000,
        "py_code": "ZRGHF",
        "remark": "阿萨德噶发",
        "status": true,
        "unit_name": "次"
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. clinic_other_cost_id | int | ✅ | 项目id | |
| data. cost | int | ❌ | 成本价 | |
| data. discount_price | ❌ | ✅ | 折扣价 | |
| data. en_name | string | ❌ | 英文名称 | |
| data. idc_code | string | ❌ | 国际编码 | |
| data. is_discount | booleam | ✅ | 是否可折扣 | |
| data. price | int | ✅ | 零售价 | |
| data. py_code | string | ❌ | 拼音码 | |
| data. remark | string | ❌ | 备注 | |
| data. status | boolean | ✅ | 是否启用 | |
| data. name | string | ✅ | 项目名称 | |
| data. unit_name | int | ✅ | 单位名称 | |
--



20 材料费用项目模块
--------

</br>
<h3>20.1 创建治疗缴费项目

```
请求地址：/material/create
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
| name | string | ✅ |  项目名称 | |
| en_name | string | ❌ |  英文名称 | |
| py_code | string | ❌ |  拼音码 | |
| idc_code | string | ❌ |  国际编码 | |
| unit_name | string | ✅ |  单位名称 | |
| remark | string | ❌ |  备注| |
| manu_factory_name | string | ✅ | 生产厂商 | |
| specification | string | ✅ |  规格| |
| buy_price | int | ✅ |  成本价| |
| ret_price | int | ✅ |  零售价 | |
| status | boolean | ❌ |  是否启用 | |
| is_discount | boolean | ❌ |  是否可打折 | |
| day_warning | int | ❌ | 预警天数| |
| stock_warning | int | ❌ | 库存预警数 | |

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
--


</br>
<h3>20.2 更新材料项目

```
请求地址：/material/update
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_material_id | int | ✅ |  项目id | |
| name | string | ✅ |  项目名称 | |
| en_name | string | ❌ |  英文名称 | |
| py_code | string | ❌ |  拼音码 | |
| idc_code | string | ❌ |  国际编码 | |
| unit_name | string | ✅ |  单位名称 | |
| remark | string | ❌ |  备注| |
| manu_factory_name | string | ✅ | 生产厂商 | |
| specification | string | ✅ |  规格| |
| buy_price | int | ✅ |  成本价| |
| ret_price | int | ✅ |  零售价 | |
| status | boolean | ❌ |  是否启用 | |
| is_discount | boolean | ❌ |  是否可打折 | |
| day_warning | int | ❌ | 预警天数| |
| stock_warning | int | ❌ | 库存预警数 | |

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
--


</br>
<h3>20.3 启用和停用

```
请求地址：/material/onOff
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
| clinic_material_id | int | ✅ |  项目id | |
| status | boolean | ❌ |  是否启用 | |

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
--


</br>
<h3>20.4 材料缴费项目列表

```
请求地址：/material/list
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
| keyword | string | ❌ |  搜索关键字| |
| status | boolean | ❌ |  是否启用 | |
| offset | int | ❌ |  开始条数| |
| limit | int | ❌ |  条数| |

**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "buy_price": 100,
            "clinic_material_id": 4,
            "day_warning": null,
            "discount_price": 0,
            "en_name": "jiashou",
            "idc_code": "gb545415212",
            "is_discount": true,
            "manu_factory_name": "生产厂家1",
            "name": "假手",
            "py_code": "jz",
            "remark": "备注11111",
            "ret_price": 1,
            "specification": "/支",
            "status": false,
            "stock_amount": null,
            "stock_warning": 100000,
            "unit_name": "支"
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
| data | array | ❌ |  | |
| data. buy_price | int | ❌ | 成本价 | |
| data. clinic_material_id | int | ❌ | 项目id | |
| data. day_warning | int | ❌ | 预警天数 | |
| data. discount_price | int | ❌ | 折扣价| |
| data. en_name | string | ❌ | 英文名称 | |
| data. idc_code | string | ❌ | 国际编码 | |
| data. is_discount | boolean | ✅ | 是否折扣| |
| data. manu_factory_name | string | ✅ | 生产厂商 | |
| data. name | string | ✅ | 项目名称 | |
| data. py_code | string | ❌ | 拼音码 | |
| data. remark | string | ❌ | 备注| |
| data. ret_price | int | ✅ | 零售价 | |
| data. specification | string | ✅ | 规格| |
| data. status | boolean | ✅ | 是否启用| |
| data. stock_amount | int | ✅ | 库存| |
| data. stock_warning | int | ❌ | 库存预警数| |
| data. unit_name | stirng | ✅ | 单位名称| |
--


</br>
<h3>20.5 材料项目详情

```
请求地址：/material/detail
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_material_id | int | ✅ |  项目id | |

**应答包示例**

```
{
    "code": "200",
    "data": {
        "buy_price": 1000,
        "clinic_id": 1,
        "created_time": "2018-05-27T21:29:47.680568+08:00",
        "day_warning": null,
        "deleted_time": null,
        "discount_price": 0,
        "en_name": "",
        "id": 1,
        "idc_code": "",
        "is_discount": false,
        "manu_factory_name": "",
        "name": "针筒",
        "py_code": "ZT",
        "remark": "噶发大嘎嘎",
        "ret_price": 2000,
        "specification": "",
        "status": true,
        "stock_warning": 100,
        "unit_name": "个",
        "updated_time": "2018-08-12T22:53:33.577843+08:00"
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | obj | ❌ |  | |
| data. buy_price | int | ❌ | 成本价 | |
| data. clinic_material_id | int | ❌ | 项目id | |
| data. day_warning | int | ❌ | 预警天数 | |
| data. discount_price | int | ❌ | 折扣价| |
| data. en_name | string | ❌ | 英文名称 | |
| data. idc_code | string | ❌ | 国际编码 | |
| data. is_discount | boolean | ✅ | 是否折扣| |
| data. manu_factory_name | string | ✅ | 生产厂商 | |
| data. name | string | ✅ | 项目名称 | |
| data. py_code | string | ❌ | 拼音码 | |
| data. remark | string | ❌ | 备注| |
| data. ret_price | int | ✅ | 零售价 | |
| data. specification | string | ✅ | 规格| |
| data. status | boolean | ✅ | 是否启用| |
| data. stock_amount | int | ✅ | 库存| |
| data. stock_warning | int | ❌ | 库存预警数| |
| data. unit_name | stirng | ✅ | 单位名称| |
--



</br>
<h3>20.6 入库

```
请求地址：/material/instock
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
| instock_operation_id | int | ✅ |  操作人id | |
| instock_way_name | string | ✅ |  入库方式 | |
| supplier_name | string | ✅ |  供应商 | |
| remark | string | ❌ |  备注| |
| instock_date | string | ✅ |  入库日期| |
| items. clinic\_material_id | string | ✅ | 项目id| |
| items. instock_amount | string | ✅ | 入库数量| |
| items. buy_price | string | ✅ | 成本价| |
| items. serial | string | ✅ | 批号| |
| items. eff_date | string | ✅ | 有效期| |



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
<h3>20.7 入库记录列表

```
请求地址：/material/instockRecord
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
| start_date | string | ❌ |  入库开始日期 | |
| end_date | string | ❌ |  入库结束日期 | |
| order_number | string | ❌ | 入库单号 | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  结束条数| |


**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "instock_date": "2018-07-25T00:00:00Z",
            "instock_operation_name": "超级管理员",
            "instock_way_name": "采购入库",
            "material_instock_record_id": 10,
            "order_number": "DRKD-1532522422",
            "supplier_name": "广州白云药厂",
            "verify_operation_name": "超级管理员",
            "verify_status": "02"
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 13
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ |  | |
| data. instock_date | time | ✅ | 入库日期 | |
| data. instock_operation_name | string | ✅ | 入库人| |
| data. instock_way_name | string | ✅ | 入库方式 | |
| data. material_instock_record_id | string | ✅ | 入库记录id | |
| data. order_number | string | ✅ | 入库单号 | |
| data. supplier_name | string | ✅ | 供应商 | |
| data. verify_operation_name | string | ✅ | 审核人 | |
| data. verify_status | string | ✅ | 审核状态 | |
--


</br>
<h3>20.8 入库记录详情

```
请求地址：/material/instockRecordDetail
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material_instock_record_id | int | ✅ |  诊所id | |


**应答包示例**

```
{
    "code": "200",
    "data": {
        "created_time": "2018-07-25T20:40:22.115607+08:00",
        "instock_date": "2018-07-25T00:00:00Z",
        "instock_operation_id": 1,
        "instock_operation_name": "超级管理员",
        "instock_way_name": "采购入库",
        "items": [
            {
                "buy_price": 100,
                "clinic_material_id": 2,
                "eff_date": "2019-12-31T00:00:00Z",
                "instock_amount": 50,
                "manu_factory_name": "生产厂家1",
                "material_name": "假牙",
                "ret_price": 1,
                "serial": "PH4165415212",
                "unit_name": "粒"
            }
        ],
        "material_instock_record_id": 10,
        "order_number": "DRKD-1532522422",
        "remark": null,
        "supplier_name": "广州白云药厂",
        "updated_time": "2018-07-26T00:34:28.425141+08:00",
        "verify_operation_id": 6,
        "verify_operation_name": "超级管理员",
        "verify_status": "02"
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ |  | |
| data. instock_date | time | ✅ | 入库日期 | |
| data. instock_operation_id | string | ✅ | 入库人id | |
| data. instock_operation_name | string | ✅ | 入库人| |
| data. instock_way_name | string | ✅ | 入库方式 | |
| data. material_instock_record_id | string | ✅ | 入库记录id | |
| data. order_number | string | ✅ | 入库单号 | |
| data. supplier_name | string | ✅ | 供应商 | |
| data. verify_operation_id | string | ✅ | 审核人id | |
| data. verify_operation_name | string | ✅ | 审核人 | |
| data. verify_status | string | ✅ | 审核状态 | |
| data. remark | string | ❌ | 备注| |
| data. items | array | ✅ | 详情| |
| data. items. buy_price | int | ✅ | 成本价| |
| data. items. clinic_material_id | int | ✅ | 项目id| |
| data. items. eff_date | int | ✅ | 有效期| |
| data. items. instock_amount | int | ✅ | 入库数量| |
| data. items. manu_factory_name | string | ✅ | 生产厂商| |
| data. items. material_name | string | ✅ | 项目名称| |
| data. items. ret_price | int | ✅ | 零售价| |
| data. items. serial | string | ✅ | 批次号| |
| data. items. unit_name | string | ✅ | 单位名称| |
--


</br>
<h3>20.9 入库记录修改

```
请求地址：/material/instockUpdate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material_instock_record_id | int | ✅ |  入库记录id | |
| clinic_id | int | ✅ |  诊所id | |
| instock_operation_id | int | ✅ |  入库人id | |
| instock_way_name | string | ✅ |  入库方式 | |
| supplier_name | string | ✅ |  供应商 | |
| remark | string | ❌ |  备注| |
| instock_date | string | ✅ |  入库日期| |
| items. clinic\_material_id | string | ✅ | 项目id| |
| items. instock_amount | string | ✅ | 入库数量| |
| items. buy_price | string | ✅ | 成本价| |
| items. serial | string | ✅ | 批号| |
| items. eff_date | string | ✅ | 有效期| |



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
<h3>20.10 入库审核

```
请求地址：/material/instockCheck
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material_instock_record_id | int | ✅ |  入库记录id | |
| verify_operation_id | int | ✅ |  操作人id | |



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
<h3>20.11 删除入库记录

```
请求地址：/material/instockDelete
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material_instock_record_id | int | ✅ |  入库记录id | |



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
<h3>20.12 出库

```
请求地址：/material/outstock
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
| outstock_operation_id | int | ✅ |  出库操作人id | |
| outstock_way_name | string | ✅ |  入库方式 | |
| department_id | int | ✅ |  科室id | |
| personnel_id | int | ✅ |  领用人id | |
| remark | string | ❌ |  备注| |
| outstock_date | string | ✅ |  出库日期| |
| items. clinic\_material_id | string | ✅ | 项目id| |
| items. outstock_amount | string | ✅ | 出库数量| |



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
<h3>20.13 出库记录列表

```
请求地址：/material/outstockRecord
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
| start_date | string | ❌ |  入库开始日期 | |
| end_date | string | ❌ |  入库结束日期 | |
| order_number | string | ❌ | 入库单号 | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  结束条数| |



**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "department_name": "眼科",
            "material_outstock_record_id": 10,
            "order_number": "DCKD-1532537059",
            "outstock_date": "2018-07-26T00:00:00Z",
            "outstock_operation_name": "超级管理员",
            "outstock_way_name": "科室领用",
            "personnel_name": "华佗",
            "verify_operation_name": "超级管理员",
            "verify_status": "02"
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 8
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ | | |
| data. department_name | string | ✅ | 出库科室 | |
| data. material_outstock_record_id | int | ✅ | 出库记录id| |
| data. order_number | string | ✅ | 出库编号 | |
| data. outstock_date | time | ✅ | 出库日期 | |
| data. outstock_operation_name | string | ✅ | 出库人员 | |
| data. outstock_way_name | string | ✅ | 出库方式 | |
| data. personnel_name | string | ✅ | 领用人员| |
| data. verify_operation_name | string | ✅ | 审核人员| |
| data. verify_status | booleam | ✅ | 审核状态 | |
--



</br>
<h3>20.14 出库记录详情

```
请求地址：/material/outstockRecordDetail
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material_outstock_record_id | int | ✅ |  出库记录id | |



**应答包示例**

```
{
    "code": "200",
    "data": {
        "created_time": "2018-07-26T00:44:19.697562+08:00",
        "department_id": 2,
        "department_name": "眼科",
        "items": [
            {
                "buy_price": 100,
                "eff_date": "2019-12-31T00:00:00Z",
                "manu_factory_name": "生产厂家1",
                "material_name": "假牙",
                "material_stock_id": 3,
                "outstock_amount": 14,
                "ret_price": 1,
                "serial": "PH4165415212",
                "stock_amount": 268,
                "supplier_name": "广州白云药厂",
                "unit_name": "粒"
            }
        ],
        "material_outstock_record_id": 10,
        "order_number": "DCKD-1532537059",
        "outstock_date": "2018-07-26T00:00:00Z",
        "outstock_operation_id": 1,
        "outstock_operation_name": "超级管理员",
        "outstock_way_name": "科室领用",
        "personnel_id": 3,
        "personnel_name": "华佗",
        "remark": "2222",
        "updated_time": "2018-07-26T00:52:28.045986+08:00",
        "verify_operation_id": 6,
        "verify_operation_name": "超级管理员"
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ | | |
| data. department_id | int | ✅ | 出库科室id | |
| data. department_name | string | ✅ | 出库科室 | |
| data. material_outstock_record_id | int | ✅ | 出库记录id| |
| data. order_number | string | ✅ | 出库编号 | |
| data. outstock_date | time | ✅ | 出库日期 | |
| data. outstock_operation_id | int | ✅ | 出库人员id | |
| data. outstock_operation_name | string | ✅ | 出库人员 | |
| data. outstock_way_name | string | ✅ | 出库方式 | |
| data. personnel_id | int | ✅ | 领用人员id| |
| data. personnel_name | string | ✅ | 领用人员| |
| data. remark | ❌ | ✅ | 备注| |
| data. verify_operation_id | int | ✅ | 审核人员id| |
| data. verify_operation_name | string | ✅ | 审核人员| |
| data. items | array | ✅ | 详情 | |
| data. items. buy_price | int | ✅ | 成本价 | |
| data. items. eff_date | time | ✅ | 有效期 | |
| data. items. manu_factory_name | string | ✅ | 生成厂商 | |
| data. items. material_name | string | ✅ | 项目名称 | |
| data. items. material_stock_id | int | ✅ | 库存id | |
| data. items. outstock_amount | int | ✅ | 出库数量 | |
| data. items. ret_price | int | ✅ | 零售价 | |
| data. items. serial | string | ✅ | 批次号 | |
| data. items. stock_amount | int | ✅ | 库存量 | |
| data. items. supplier_name | string | ✅ | 供应商 | |
| data. items. unit_name | string | ✅ | 单位名称 | |
--



</br>
<h3>20.15 出库记录修改

```
请求地址：/material/outstockUpdate
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
| material_outstock_record_id | int | ✅ | 出库记录id | |
| outstock_operation_id | int | ✅ |  出库操作人id | |
| outstock_way_name | string | ✅ |  入库方式 | |
| department_id | int | ✅ |  科室id | |
| personnel_id | int | ✅ |  领用人id | |
| remark | string | ❌ |  备注| |
| outstock_date | string | ✅ |  出库日期| |
| items. clinic\_material_id | string | ✅ | 项目id| |
| items. outstock_amount | string | ✅ | 出库数量| |



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
<h3>20.16 出库审核

```
请求地址：/material/outstockCheck
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material_outstock_record_id | int | ✅ | 出库记录id | |
| verify_operation_id | int | ✅ |  审核人id | |



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
<h3>20.17 删除出库记录

```
请求地址：/material/outstockDelete
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material_outstock_record_id | int | ✅ | 出库记录id | |


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
<h3>20.18 库存列表

```
请求地址：/material/MaterialStockList
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ | 诊所id | |
| keyword | string | ❌ | 关键字 | |
| supplier_name | stirng | ❌ | 供应商| |
| amount | boolean | ❌ | 是否大于0 | |
| date_warning | boolean | ❌ | 是否在 预警期内 | |
| offset | int | ❌ | 开始条数 | |
| limit | int | ❌ | 条数| |


**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "buy_price": 99,
            "eff_date": "2018-10-01T00:00:00Z",
            "manu_factory_name": "",
            "material_stock_id": 1,
            "name": "针筒",
            "ret_price": 2000,
            "serial": "1",
            "specification": "",
            "stock_amount": 59,
            "supplier_name": "广州白云药厂",
            "unit_name": "个"
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 9
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ |   | |
| data. buy_price | int | ✅ | 成本价  | |
| data. eff_date | time | ✅ | 有效期  | |
| data. manu_factory_name | string | ✅ | 生产厂商  | |
| data. material_stock_id | int | ✅ | 库存id  | |
| data. name | string | ✅ | 项目名称  | |
| data. ret_price | int | ✅ | 零售价  | |
| data. serial | string | ✅ | 批次  | |
| data. specification | string | ✅ | 规格  | |
| data. stock_amount | int | ✅ | 库存量  | |
| data. supplier_name | string | ✅ | 供应商  | |
| data. unit_name | int | ✅ | 单位名称  | |
--




</br>
<h3>20.19 新增耗材盘点

```
请求地址：/material/MaterialInventoryCreate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ | 诊所id | |
| inventory_operation_id | int | ✅ | 盘点人员id | |
| items | stirng | ✅ | 详情 json 字符串| |
| items. material_stock_id | int | ✅ | 库存id | |
| items. actual_amount | int | ✅ | 实际数量 | |


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
<h3>20.20 耗材盘点记录列表

```
请求地址：/material/MaterialInventoryList
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ | 诊所id | |
| start_date | string | ❌ | 开始日期 | |
| end_date | stirng | ❌ | 结束日期| |
| offset | int | ❌ | 开始条数 | |
| limit | int | ❌ | 条数 | |


**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "inventory_date": "2018-08-07T00:00:00Z",
            "inventory_operation_name": "人中龙凤",
            "material_inventory_record_id": 1,
            "order_number": "DPD-1533653517",
            "verify_operation_name": "人中龙凤",
            "verify_status": "02"
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
| data | array | ❌ |  | |
| data. inventory_date | time | ✅ | 盘点日期 | |
| data. inventory_operation_name | string | ✅ | 盘点人员 | |
| data. material_inventory_record_id | int | ✅ |盘点记录id | |
| data. order_number | string | ✅ | 盘点单号 | |
| data. verify_operation_name | time | ✅ | 审核人员名称 | |
| data. verify_status | time | ✅ | 审核状态 | |
--



</br>
<h3>20.21 耗材盘点记录详情

```
请求地址：/material/MaterialInventoryRecordDetail
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ | 诊所id | |
| start_date | string | ❌ | 开始日期 | |
| end_date | stirng | ❌ | 结束日期| |
| offset | int | ❌ | 开始条数 | |
| limit | int | ❌ | 条数 | |


**应答包示例**

```
{
    "code": "200",
    "data": {
        "created_time": "2018-08-07T22:51:57.388347+08:00",
        "inventory_date": "2018-08-07T00:00:00Z",
        "inventory_operation_id": 11,
        "inventory_operation_name": "人中龙凤",
        "items": [
            {
                "actual_amount": 270,
                "buy_price": 100,
                "eff_date": "2019-12-31T00:00:00Z",
                "manu_factory_name": "生产厂家1",
                "material_name": "假牙",
                "material_stock_id": 3,
                "serial": "PH4165415212",
                "specification": "20g/颗",
                "status": true,
                "stock_amount": 268,
                "supplier_name": "广州白云药厂",
                "unit_name": "粒"
            },
            ...
        ],
        "material_inventory_record_id": 1,
        "order_number": "DPD-1533653517",
        "updated_time": "2018-08-07T23:01:05.519431+08:00",
        "verify_operation_id": 11,
        "verify_operation_name": "人中龙凤",
        "verify_status": "02"
    },
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 9,
        "total_item": [
            {
                "actual_amount": 60,
                "material_stock_id": 1
            },
            ...
        ]
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ |  | |
| data. created_time | time | ✅ | 创建时间 | |
| data. inventory_date | time | ✅ | 盘点日期 | |
| data. inventory\_operation_id | id | ✅ | 盘点人员id | |
| data. inventory\_operation_name | string | ✅ | 盘点人员 | |
| data. material\_inventory\_record_id | int | ✅ |盘点记录id | |
| data. order_number | string | ✅ | 盘点单号 | |
| data. verify\_operation_id | id | ✅ | 审核人员id | |
| data. verify\_operation_name | time | ✅ | 审核人员名称 | |
| data. verify_status | time | ✅ | 审核状态 | |
| data. items | array | ✅ | 盘点详情| |
| data. items. actual_amount | int | ✅ | 实际数量 | |
| data. items. buy_price | int | ✅ | 成本价| |
| data. items. eff_date | time | ✅ | 有效期 | |
| data. items. manu_factory_name | string | ✅ | 生产厂商 | |
| data. items. material_name | string | ✅ | 项目名称 | |
| data. items. material_stock_id | int | ✅ | 库存id | |
| data. items. serial | string | ✅ | 批次号| |
| data. items. specification | string | ✅ | 规格 | |
| data. items. status | boolean | ✅ | 是否启用| |
| data. items. stock_amount | int | ✅ | 库存数量 | |
| data. items. supplier_name | string | ✅ | 供应商 | |
| data. items. unit_name | string | ✅ | 单位名称 | |
| page_info | obj | ✅ | 分页信息 | |
| page_info. limit | string | ✅ | 每页条数 | |
| page_info. offset | string | ✅ | 开始条数 | |
| page_info. total | string | ✅ | 总条数 | |
| page\_info. total_item | array | ✅ |  | |
| page\_info. total\_item. actual_amount| int | ✅ | 真实数量 | |
| page\_info. total\_item. actual_amount| int | ✅ | 真实数量 | |
| page\_info. total\_item. material\_stock_id| int | ✅ | 库存id | |

--



</br>
<h3>20.22 修改耗材盘点

```
请求地址：/material/MaterialInventoryUpdate
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ | 诊所id | |
| material\_inventory\_record_id | int | ✅ | 盘点记录id | |
| inventory\_operation_id | int | ✅ | 盘点人员id | |
| items | stirng | ✅ | 详情 json 字符串| |
| items. material_stock_id | int | ✅ | 库存id | |
| items. actual_amount | int | ✅ | 实际数量 | |


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
<h3>20.23 耗材盘点审核

```
请求地址：/material/MaterialInventoryCheck
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material\_inventory_record_id | int | ✅ | 盘点记录id | |
| verify\_operation_id | int | ✅ | 审核人员id | |


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
<h3>20.24 删除盘点记录

```
请求地址：/material/MaterialInventoryRecordDelete
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| material\_inventory_record_id | int | ✅ | 盘点记录id | |


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
<h3>20.25 耗材盘点列表

```
请求地址：/material/MaterialStockInventoryList
```
**请求包示例**

```
{
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | int | ✅ | 诊所id | |
| keyword | string | ❌ | 搜索关键字| |
| status | booleam | ✅ | 是否启用 | |
| amount | boolean | ✅ | 是否有库存 | |
| offset | int | ✅ | 开始条数 | |
| limit | int | ✅ | 结束条数 | |


**应答包示例**

```
{
    "code": "200",
    "data": [
        {
            "buy_price": 99,
            "day_warning": null,
            "eff_date": "2018-10-01T00:00:00Z",
            "manu_factory_name": "",
            "material_name": "针筒",
            "material_stock_id": 1,
            "ret_price": 2000,
            "serial": "1",
            "specification": "",
            "status": true,
            "stock_amount": 59,
            "stock_warning": 100,
            "supplier_name": "广州白云药厂",
            "unit_name": "个"
        },
        ...
    ],
    "page_info": {
        "limit": "10",
        "offset": "0",
        "total": 9
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | array | ❌ |  | |
| data. buy_price | int | ✅ | 成本价 | |
| data. day_warning | int | ❌ | 预警天数 | |
| data. eff_date | time | ✅ | 有效期 | |
| data. manu\_factory_name | string | ✅ | 生产厂商 | |
| data. material_name | string | ✅ | 项目名称 | |
| data. material\_stock_id | int | ✅ | 库存id | |
| data. ret_price | int | ✅ | 零售价 | |
| data. serial | string | ✅ | 批次号 | |
| data. specification | string | ✅ | 规格 | |
| data. status | boolean | ✅ | 是否启用 | |
| data. stock_amount | int | ✅ | 库存 | |
| data. stock_warning | int | ✅ | 库存预警数 | |
| data. supplier_name | string | ✅ | 供应商| |
| data. unit_name | string | ✅ | 单位名称 | |
--
