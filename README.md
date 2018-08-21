云诊所接口文档
===========

**创建时间：2018-08-16**

修改记录
--------
| 修定日期 | 修改内容 | 修改人 | 
| :-: | :-: | :-:  | 

接口列表
--------
>1 诊所模块
>>1.1 添加诊所
>>
>>1.2 用户登录
>>
>>1.3 启用、关闭诊所
>>
>>1.4 更新诊所
>>
>>1.5 获取诊所详情
>>
>>1.6 获取最新的诊所编码

>2 科室模块
>>2.1 添加科室
>>
>>2.2 诊所科室列表
>>
>>2.3 科室删除
>>
>>2.4 科室修改

>3 人员模块
>>3.1 诊所用户登录
>>
>>3.2 添加人员
>>
>>3.3 通过id获取人员
>>
>>3.4 获取人员列表
>>
>>3.5 通过用户查询用户菜单
>>
>>3.6 修改人员
>>
>>3.7 删除人员
>>
>>3.8 有账号的人员列表（包含角色）
>>
>>3.9 修改账号生效状态
>>
>>3.10 修改用户名密码
>>
>>3.11 通过用户查询用户角色
>>
>>3.12 获取医生所属科室

>4 医生排班模块
>>4.1 获取号源列表
>>
>>4.2 号源科室列表
>>
>>4.3 号源科室下医生列表
>>
>>4.4 获取所有医生的号源信息
>>
>>4.5 复制排版
>>
>>4.6 开放号源
>>
>>4.7 插入单个号源
>>
>>4.8 删除单个未开放号源 byid
>>
>>4.9 停诊号源byid

>5 就诊人模块
>>
>>5.1 新增就诊人
>>
>>5.2 就诊人列表
>>
>>5.3 通过id就诊人详情
>>
>>5.4 修改就诊人
>>
>>5.5 通过身份号查就诊人
>>
>>5.6 通过关键字搜索就诊人
>>
>>5.7 会员，就诊人列表
>>
>>5.8 完善患者个人诊前病历
>>
>>5.9 患者个人诊前病历
>>
>>5.10 获取最后一次体征信息
>>
>>5.11 修改身高
>>
>>5.12 修改体重
>>
>>5.13 修改BMI
>>
>>5.14 修改血型
>>
>>5.15 修改RH血型
>>
>>5.16 修改体温
>>
>>5.17 修改呼吸
>>
>>5.18 修改脉搏
>>
>>5.19 修改血压
>>
>>5.20 修改视力
>>
>>5.21 修改血糖
>>
>>5.22 修改氧饱和度
>>
>>5.23 患者身高记录
>>
>>5.24 患者体重记录
>>
>>5.25 患者BMI记录
>>
>>5.26 患者血型记录
>>
>>5.27 患者RH血型记录
>>
>>5.28 患者体温记录
>>
>>5.29 患者呼吸记录
>>
>>5.30 患者脉搏记录
>>
>>5.31 患者血压记录
>>
>>5.32 患者视力记录
>>
>>5.33 患者氧饱和度记录
>>
>>5.34 患者统计 按性别
>>
>>5.35 患者统计 按年龄
>>
>>5.36 患者统计 按渠道

>6 就诊模块
>>6.1 就诊患者登记
>>
>>6.2 登记患者详情
>>
>>6.3 当日登记就诊人列表 分诊记录
>>
>>6.4 接诊就诊人列表
>>
>>6.5 通过id就诊人 查询患者
>>
>>6.6 分诊医生列表
>>
>>6.7 晚上体征信息
>>
>>6.8 完善诊前病历
>>
>>6.9 完善诊前欲诊
>>
>>6.10 获取健康档案
>>
>>6.11 分诊、换诊(选择医生)
>>
>>6.12 医生接诊病人
>>
>>6.13 医生完成接诊
>>
>>6.14 按日期统计挂号记录
>>
>>6.15 开治疗
>>
>>6.16 查询治疗
>>
>>6.17 开检验
>>
>>6.18 获取检验
>>
>>6.19 开西/成药处方
>>
>>6.20 获取西药处方
>>
>>6.21 获取西药历史处方列表
>>
>>6.22 开中药处方
>>
>>6.23 删除中药处方
>>
>>6.24 获取中药处方
>>
>>6.25 获取中药历史处方列表
>>
>>6.26 开检查
>>
>>6.27 获取检查
>>
>>6.28 开其它费用
>>
>>6.29 获取其它费
>>
>>6.30 开材料费
>>
>>6.31 获取其它费
>>
>>6.32 获取病人历史已接诊记录
>>
>>6.33 开诊疗费
>>
>>6.34 获取病人就诊信息详情
>>
>>6.35 分诊记录报告
>>
>>6.36 快速接诊

>7 诊疗模块
>>7.1 创建诊疗缴费项目
>>
>>7.2 更新诊疗缴费项目
>>
>>7.3 启用和停用诊疗项目
>>
>>7.4 诊疗项目列表
>>
>>7.5 诊疗项目详情

>8 门诊缴费状态
>>8.1 门诊待缴费的分诊记录
>>
>>8.2 门诊待缴费订单
>>
>>8.3 创建门诊支付订单
>>
>>8.4 获取门诊支付订单状态
>>
>>8.5 门诊退费
>>
>>8.6 门诊已缴费的分诊记录
>>
>>8.7 门诊已缴费订单
>>
>>8.8 门诊已退费的分诊记录
>>
>>8.9 获取交易流水日报表
>>
>>8.10 获取分析类报表
>>
>>8.11 获取交易流水月报表
>>
>>8.12 获取交易详情
>>
>>8.13 获取交易订单信息

>9 挂账模块
>>9.1 有挂账的分诊记录

>10 预约模块
>>10.1 预约挂号

>11 药品模块
>>11.1 添加药品
>>
>>11.2 更新药品
>>
>>11.3 启用或停止药品
>>
>>11.4 药品列表
>>
>>11.5 查询药品库存信息
>>
>>11.6 药品详情
>>
>>11.7 批量设置药品
>>
>>11.8 药品入库
>>
>>11.9 入库记录列表
>>
>>11.10 入库记录详情
>>
>>11.11 入库记录修改
>>
>>11.12 入库审核
>>
>>11.13 删除入库记录
>>
>>11.14 出库
>>
>>11.15 出库记录
>>
>>11.16 出库记录详情
>>
>>11.17 更新出库
>>
>>11.18 出库审核
>>
>>11.19 删除出库记录
>>
>>11.20 库存列表
>>
>>11.21 创建西药处方模板
>>
>>11.22 西药处方模板列表
>>
>>11.23 查询个人西药处方模板
>>
>>11.24 西药处方模板详情
>>
>>11.25 修改西药处方模板
>>
>>11.26 删除西药处方模板
>>
>>11.27 创建中药处方模板
>>
>>11.28 中药处方模板列表
>>
>>11.29 个人中药处方模板
>>
>>11.30 中药处方模板详情
>>
>>11.31 更新中药处方模板
>>
>>11.32 删除中药模板
>>
>>11.33 新增药房盘点
>>
>>11.34 药房盘点记录表
>>
>>11.35 药房盘点记录详情
>>
>>11.36 修改药房盘点
>>
>>11.37 药房盘点审核
>>
>>11.38 删除盘点记录
>>
>>11.39 库存盘点列表

>12 角色模块
>>12.1 创建角色
>>
>>12.2 更新角色
>>
>>12.3 更新角色
>>
>>12.4 角色列表
>>
>>12.5 角色详情
>>
>>12.6 获取角色未开通的菜单项
>>
>>12.7 获取角色未开通的菜单项
>>
>>12.8 在角色下分配用户
>>
>>12.9 角色分配的用户列表

>13 业务权限模块
>>13.1 添加功能菜单栏
>>
>>13.2 添加功能菜单栏
>>
>>13.3 获取诊所未开通的菜单项
>>
>>13.4 诊所分配业务
>>
>>13.5 诊所分配业务

>14 管理用户模块
>>14.1 诊所用户登录
>>
>>14.2 平台账号添加
>>
>>14.3 获取诊所列表
>>
>>14.4 平台账号修改
>>
>>14.5 获取平台账号信息
>>
>>14.6 启用和停用管理员账号
>>
>>14.7 获取平台未开通的菜单项
>>
>>14.8 获取平台开通菜单项

>15 诊断模块
>>15.1 创建诊断

>16 病历模块
>>16.1 创建主病历
>>
>>16.2 续写病历
>>
>>16.3 续写病历修改
>>
>>16.4 续写病历删除
>>
>>16.5 通过就诊id查找病历
>>
>>16.6 创建病历模板
>>
>>16.7 修改病历模板
>>
>>16.8 删除病历模板
>>
>>16.9 获取患者病历列表
>>
>>16.10 查询模板列表

>17 检查缴费项目模块
>>17.1 创建检查缴费项目
>>
>>17.2 修改检查缴费项目
>>
>>17.3 启用和停用检查项目
>>
>>17.4 诊疗项目列表
>>
>>17.5 检查项目详情
>>
>>17.6 创建检查医嘱模板
>>
>>17.7 查询检查模板列表
>>
>>17.8 查询检查医嘱模板详情
>>
>>17.9 修改检查医嘱模板
>>
>>17.10 删除检查医嘱模板
>>
>>17.11 创建检查报告医嘱模板
>>
>>17.12 查询检查报告医嘱模板
>>
>>17.13 查询检查报告医嘱模板详情
>>
>>17.14 修改检查报告医嘱模板
>>
>>17.15 删除检查报告医嘱模板

>18 治疗收费项目模块
>>18.1 创建治疗缴费项目
>>
>>18.2 更新治疗项目
>>
>>18.3 启用和停用
>>
>>18.4 治疗缴费项目列表
>>
>>18.5 治疗项目详情
>>
>>18.6 创建治疗医嘱模板
>>
>>18.7 查询治疗医嘱模板
>>
>>18.8 查询治疗医嘱模板详情
>>
>>18.9 修改治疗医嘱模板
>>
>>18.10 删除治疗医嘱模板

>19 其他费用收费项目模块
>>19.1 创建其它费用项目
>>
>>19.2 更新其它费用项目
>>
>>19.3 启用和停用
>>
>>19.4 其它费用缴费项目列表
>>
>>19.5 其它费用项目详情

>20 材料费用项目模块
>>20.1 创建材料费用项目
>>
>>20.2 更新材料项目
>>
>>20.3 启用和停用
>>
>>20.4 材料缴费项目列表
>>
>>20.5 材料项目详情
>>
>>20.6 入库
>>
>>20.7 入库记录列表
>>
>>20.8 入库记录详情
>>
>>20.9 入库记录修改
>>
>>20.10 入库审核
>>
>>20.11 删除入库记录
>>
>>20.12 出库
>>
>>20.13 出库记录列表
>>
>>20.14 出库记录详情
>>
>>20.15 出库记录修改
>>
>>20.16 出库审核
>>
>>20.17 删除出库记录
>>
>>20.18 库存列表
>>
>>20.19 新增耗材盘点
>>
>>20.20 耗材盘点记录列表
>>
>>20.21 耗材盘点记录详情
>>
>>20.22 修改耗材盘点
>>
>>20.23 耗材盘点审核
>>
>>20.24 删除盘点记录
>>
>>20.25 耗材盘点列表

>21 检验模块
>>21.1 检验医嘱创建
>>
>>21.2 检验医嘱列表
>>
>>21.3 检验医嘱详情
>>
>>21.4 检验医嘱更新
>>
>>21.5 检验医嘱启用
>>
>>21.6 关联检验项目
>>
>>21.7 检验医嘱关联项目列表
>>
>>21.8 检验项目创建
>>
>>21.9 检验项目创建
>>
>>21.10 检验项目更新
>>
>>21.11 检验项目启用
>>
>>21.12 检验项目列表
>>
>>21.13 创建检验医嘱模板
>>
>>21.14 检验医嘱模板列表
>>
>>21.15 个人检验医嘱模板
>>
>>21.16 检验医嘱模板详情
>>
>>21.17 修改检验医嘱模板
>>
>>21.18 删除检验医嘱模板

>22 基础字典表模块
>>22.1 查询单位列表
>>
>>22.2 药品剂型列表
>>
>>22.3 药物类型列表
>>
>>22.4 药物种类列表
>>
>>22.5 药品别名列表
>>
>>22.6 检查部位列表
>>
>>22.7 使用频率列表
>>
>>22.8 用药途径列表
>>
>>22.9 标本种类列表
>>
>>22.10 试管颜色列表
>>
>>22.11 生产厂商列表
>>
>>22.12 基础检验医嘱项目列表
>>
>>22.13 基础检查医嘱项目列表
>>
>>22.14 基础检验项目列表
>>
>>22.15 基础药品列表
>>
>>22.16 供应商列表
>>
>>22.17 入库方式
>>
>>22.18 出库方式
>>
>>22.19 诊断字典

>23 主诉模板
>>23.1 检验医嘱创建

>24 门诊发药
>>24.1 获取待发药的分诊记录
>>
>>24.2 获取已发药的分诊记录
>>
>>24.3 获取已退药的分诊记录
>>
>>24.4 获取发药药品详情列表
>>
>>24.5 创建发药记录
>>
>>24.6 退药
>>
>>24.7 查询发药记录
>>
>>24.8 查询发药记录详情
>>
>>24.9 查询退药记录
>>
>>24.10 查询发药记录详情

>25 检查模块
>>25.1 获取检查列表
>>
>>25.2 获取待检查的分诊记录
>>
>>25.3 获取已检查的分诊记录
>>
>>25.4 获取检查中的分诊记录
>>
>>25.5 创建检查记录
>>
>>25.6 查询检查记录
>>
>>25.7 更新检查状态
>>
>>25.8 患者历史检查记录

>26 检验报告模板
>>26.1 获取检验列表
>>
>>26.2 获取检验详情
>>
>>26.3 获取待检验的分诊记录
>>
>>26.4 获取已检验的分诊记录
>>
>>26.5 获取检验中的分诊记录
>>
>>26.6 创建检验记录
>>
>>26.7 查询检验记录
>>
>>26.8 查询检验详情
>>
>>26.9 更新检查状态
>>
>>26.10 患者历史检验记录

>27 治疗模板
>>27.1 获取治疗列表
>>
>>27.2 获取待治疗的分诊记录
>>
>>27.3 获取已治疗的分诊记录
>>
>>27.4 获取治疗中的分诊记录
>>
>>27.5 创建治疗记录
>>
>>27.6 查询治疗记录
>>
>>27.7 更新治疗状态
>>
>>27.8 患者历史治疗记录
>>
>>27.9 删除治疗医嘱模板

>28 零售模块
>>28.1 创建药品零售订单
>>
>>28.2 创建支付订单
>>
>>28.3 获取收费状态
>>
>>28.4 获取药品零售表
>>
>>28.5 获取药品详情
>>
>>28.6 退费

>29 财务分析、运营分析模块
>>29.1 收费日报表按支付方式
>>
>>29.2 收费日报表按业务类型
>>
>>29.3 收费月报表按支付方式
>>
>>29.4 利润报表

>30 平台交易总额模块
>>30.1 交易总额

>31 医用报表模块
>>31.1 门诊日志
>>
>>31.2 接诊类型统计
>>
>>31.3 科室统计
>>
>>31.4 医生接诊统计
>>
>>31.5 检查统计
>>
>>31.6 检验统计
>>
>>31.7 治疗统计
>>
>>31.8 登记统计
>>
>>31.9 门诊效率统计

>32 进销存统计模块
>>32.1 药品入库统计
>>
>>32.2 药品出库统计
>>
>>32.3 药品进存销统计
>>
>>32.4 耗材入库统计
>>
>>32.5 耗材出库统计
>>
>>32.6 耗材进存销统计

---

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

8 门诊缴费状态
--------

</br>
<h3>8.1 门诊待缴费的分诊记录

```
请求地址：/charge/traigePatient/unpay
```
**请求包示例**

```
{
	keyword:‘’
	clinic_id: 1
	offset: 1
	limit: 10
	start_date:'2018-01-10'
	end_date:'2018-02-10'
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ |  诊所id | |
| keyword | String | ❌ |  搜索关键字 | |
| offset |number | ❌ |  分页（跳过个数） | 0|
| limit | number | ❌ |  分页（每页个数） | 6|
| start_date | String | ✅ |  开始时间 | |
| end_date | String | ✅ |  结束时间 | |

**应答包示例**

```
{
    "code": "200",
    "msg": ""
    “data”:[
     {
      "birthday": "19810327",
      "cert_no": null,
      "charge_total_fee": 5200,
      "clinic_patient_id": 9,
      "clinic_triage_patient_id": 15,
      "department_name": "眼科",
      "doctor_name": "华佗",
      "patient_id": 13,
      "patient_name": "林俊杰",
      "phone": "18800000001",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-05-31T21:10:34.157788+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-05-31T22:28:09.862734+08:00",
      "visit_date": "2018-05-31T00:00:00Z"
    }]
    "page_info":{
       limit: 6
       offset: 0
       total:  10
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ |  生日 | |
| cert_no | String | ✅ |  证件号| |
| charge\_total_fee |number | ✅ | 待缴费金额 | |
| clinic\_patient_id | number | ✅ | 诊所病人id  | |
| clinic\_triage\_patient_id | number | ✅ |  分诊记录id| |
| department_name | String | ✅ |  科室名称 | |
| doctor_name | String | ✅ |  医生名称 | |
| patient_id | String | ✅ |  病人id | |
| patient_name | String | ✅ |  病人名称 | |
| phone | String | ✅ |  手机号 | |
| register_personnel_name | String | ✅ |  操作员名称 | |
| register_time | String | ✅ |  挂号时间 | |
| register_type | String | ✅ |  挂号类型（1预约，2线下分诊 3快速接诊） | |
| sex | String | ✅ |  病人性别 | |
| status | String | ✅ |  挂号记录状态 10:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消 | |
| updated_time | String | ✅ |  挂号记录更新状态 | |
| visit_date | String | ✅ |  就诊日期 | |
--

</br>
<h3>8.2 门诊待缴费订单

```
请求地址：/charge/unPay/list
```
**请求包示例**

```
{
	clinic_triage_patient_id:1
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊记录id | |
| offset |number | ❌ |  分页 （跳过条数） | 0|
| limit | number | ❌ |  分页 （每页条数）| 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "amount": 1,
      "department_name": null,
      "discount": 0,
      "doctor_name": "超级管理员",
      "fee": 1200,
      "mz_unpaid_orders_id": 48,
      "name": "感冒灵片",
      "price": 1200,
      "total": 1200
    }
  ],
  "page_info": {
    "charge_total": 5200,
    "charge_total_fee": 5200,
    "discount_total": 0,
    "limit": "10",
    "offset": "0",
    "total": 3,
    "totalIds": "46,47,48"
  },
  "type_total": [
    {
      "charge_project_type_id": 1,
      "type_charge_total": 5200
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| amount | number | ✅ |  数量| |
| department_name | String | ✅ |  科室名称| |
| discount |number | ✅ |  折扣金额（单位：分）| |
| doctor_name | String | ✅ |  医生名称| |
| fee | number | ✅ |  缴费金额| |
| mz\_unpaid\_orders_id | String | ✅ |  待缴费id| |
| name | String | ✅ |  费用项名称| |
| price | number | ✅ |  单价| |
| total | number | ✅ |  总价格| |
| type\_total.charge\_project\_type_id | number | ✅ |  类型id| |
| type\_total.type\_charge_total | number | ✅ | 类型费用 | |
| page\_info.charge_total | number | ✅ |  总金额| |
| page\_info.charge\_total_fee | number | ✅ |  折扣后待缴费的金额| |
| page\_info.discount_total | number | ✅ |  折扣金额| |
| page\_info.totalIds | number | ✅ |  所有待缴费的id| |
|page\_info.total | number | ✅ |  所有条数| |
| page\_info.limit | number | ✅ |  跳过条数| |
| page\_info.offset | number | ✅ | 每页条数 | |
--


</br>
<h3>8.3 创建门诊支付订单

```
请求地址：/charge/payment/create
```
**请求包示例**

```
{
	discount_money:1
	derate_money:1
	medical_money:1
	voucher_money:1
	bonus_points_money:1
	balance_money:1
	auth_code:112122121.....
	clinic_triage_patient_id:1
	orders_ids:2,23,23,43,534
	operation_id:1
	pay_method_code:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| discount_money | number | ❌ | 折扣金额 | 0|
| derate_money | number | ❌ | 减免金额 | 0|
| medical_money | number | ❌ |  医保金额 | 0|
| voucher_money | number | ❌ |  抵金券金额 | 0|
| bonus\_points_money | number | ❌ | 积分金额 |0|
| balance_money | number | ✅ | 应交金额| |
| auth_code | String | ❌ |  授权码（微信和支付宝支付时，比传）| |
| clinic\_triage\_patient_id | number | ✅ |  分诊记录id | |
| orders_ids | String |  ✅ | 门诊未交费id组（以,隔开） | |
| operation_id | number | ✅ | 操作者id | |
| pay\_method_code | String | ✅ |   | |

**应答包示例**

```
{
    "code": "200",
    "msg": "注册成功"
    "data": "12132"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 缴费成功， 300 待用户支付（微信和支付宝支付）| |
| msg | String | ✅ |  返回信息 | |
| data | String | ❌ |  系统交易号 | |
--

</br>
<h3>8.4 获取门诊支付订单状态

```
请求地址：/charge/payment/query
```
**请求包示例**

```
{
	out_trade_no:1212232.....
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| out\_trade_no | String | ✅ |  系统交易号 | |

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
| code | String | ✅ |  返回码， 200 成功. 其他缴费失败| |
| msg | String | ✅ |  返回信息 | |
--

</br>
<h3>8.5 门诊退费

```
请求地址：/charge/payment/refund
```
**请求包示例**

```
{
	out_trade_no:1212232.....
	refundIds: 1,2,3,4
	operation_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| out\_trade_no | String | ✅ |  系统交易号 | |
| refundIds | String | ✅ |  门诊已缴费订单id（已,隔开） | |
| operation_id | String | ✅ |  操作员id | |

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
| code | String | ✅ |  返回码， 200 成功. 其他失败| |
| msg | String | ✅ |  返回信息 | |
--

</br>
<h3>8.6 门诊已缴费的分诊记录

```
请求地址：/charge/traigePatient/paid
```
**请求包示例**

```
{
	keyword:‘’
	clinic_id: 1
	offset: 1
	limit: 10
	start_date:'2018-01-10'
	end_date:'2018-02-10'
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ |  诊所id | |
| keyword | String | ❌ |  搜索关键字 | |
| offset |number | ❌ |  分页（跳过个数） | 0|
| limit | number | ❌ |  分页（每页个数） | 6|
| start_date | String | ✅ |  开始时间 | |
| end_date | String | ✅ |  结束时间 | |

**应答包示例**

```
{
    "code": "200",
    "msg": ""
    “data”:[
     {
      "birthday": "19810327",
      "cert_no": null,
      "charge_total_fee": 5200,
      "clinic_patient_id": 9,
      "clinic_triage_patient_id": 15,
      "department_name": "眼科",
      "mz_paid_record_id": 114,
      "doctor_name": "华佗",
      "patient_id": 13,
      "patient_name": "林俊杰",
      "phone": "18800000001",
      "refund_money": -1,
      "register_personnel_name": "超级管理员",
      "register_time": "2018-05-31T21:10:34.157788+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-05-31T22:28:09.862734+08:00",
      "visit_date": "2018-05-31T00:00:00Z"
    }]
    "page_info":{
       limit: 6
       offset: 0
       total:  10
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ |  生日 | |
| mz\_paid\_record_id | String | ✅ |  门诊缴费记录id  | |
| refund_money | number | ✅ |  已退费金额 | |
| cert_no | String | ✅ |  证件号| |
| charge\_total_fee |number | ✅ | 待缴费金额 | |
| clinic\_patient_id | number | ✅ | 诊所病人id  | |
| clinic\_triage\_patient_id | number | ✅ |  分诊记录id| |
| department_name | String | ✅ |  科室名称 | |
| doctor_name | String | ✅ |  医生名称 | |
| patient_id | String | ✅ |  病人id | |
| patient_name | String | ✅ |  病人名称 | |
| phone | String | ✅ |  手机号 | |
| register\_personnel_name | String | ✅ |  操作员名称 | |
| register_time | String | ✅ |  挂号时间 | |
| register_type | String | ✅ |  挂号类型（1预约，2线下分诊 3快速接诊） | |
| sex | String | ✅ |  病人性别 | |
| status | String | ✅ |  挂号记录状态 10:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消 | |
| updated_time | String | ✅ |  挂号记录更新状态 | |
| visit_date | String | ✅ |  就诊日期 | |
--

</br>
<h3>8.7 门诊已缴费订单

```
请求地址：/charge/paid/list
```
**请求包示例**

```
{
	mz_paid_record_id:1
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊记录id | |
| offset |number | ❌ |  分页 （跳过条数） | 0|
| limit | number | ❌ |  分页 （每页条数）| 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "amount": 1,
      "department_name": null,
      "discount": 0,
      "doctor_name": "超级管理员",
      "fee": 1200,
      "mz_unpaid_orders_id": 48,
      "name": "感冒灵片",
      "price": 1200,
      "total": 1200
    }
  ],
  "page_info": {
    "balance_money": 1,
    "bonus_points_money": 0,
    "charge_total": 1,
    "charge_total_fee": 1,
    "clinic_triage_patient_id": 134,
    "created_time": "2018-08-13T01:10:00.64325+08:00",
    "deleted_time": null,
    "derate_money": 0,
    "discount_money": 0,
    "discount_total": 0,
    "id": 114,
    "limit": "10",
    "medical_money": 0,
    "offset": "0",
    "operation_id": 1,
    "orders_ids": "",
    "out_trade_no": "T2201808130110006428",
    "pay_method_code": "4",
    "refund_money": -1,
    "status": "TRADE_SUCCESS",
    "total": 1,
    "total_money": 1,
    "trade_no": "T2201808130110006428",
    "updated_time": "2018-08-13T01:10:00.745739+08:00",
    "voucher_money": 0
  },
  "type_total": [
    {
      "charge_project_type_id": 1,
      "type_charge_total": 5200
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| amount | number | ✅ |  数量| |
| department_name | String | ✅ |  科室名称| |
| discount |number | ✅ |  折扣金额（单位：分）| |
| doctor_name | String | ✅ |  医生名称| |
| fee | number | ✅ |  缴费金额| |
| mz\_unpaid\_orders_id | String | ✅ |  待缴费id| |
| name | String | ✅ |  费用项名称| |
| price | number | ✅ |  单价| |
| total | number | ✅ |  总价格| |
| type\_total.charge\_project\_type_id | number | ✅ |  类型id| |
| type\_total.type\_charge_total | number | ✅ | 类型费用 | |
| page\_info.charge_total | number | ✅ |  总金额| |
| page\_info.charge\_total_fee | number | ✅ |  折扣后待缴费的金额| |
| page\_info.discount_total | number | ✅ |  折扣金额| |
| page\_info.bonus\_points_money | number | ✅ |  积分金额| |
| page\_info.medical_money | number | ✅ |  医保金额| |
| page\_info.derate_money | number | ✅ |  减免金额| |
| page\_info.medical_money | number | ✅ |  医保金额| |
| page\_info.voucher_money | number | ✅ |  抵金券金额| |
| page\_info.totalIds | number | ✅ |  所有待缴费的id| |
| page\_info.status | number | ✅ |  缴费状态| |
| page\_info.trade_no | number | ✅ |  第三方交易号| |
| page\_info.total | number | ✅ |  所有条数| |
| page\_info.limit | number | ✅ |  跳过条数| |
| page\_info.offset | number | ✅ | 每页条数 | |
--

</br>
<h3>8.8 门诊已退费的分诊记录

```
请求地址：/charge/traigePatient/refund
```
**请求包示例**

```
{
	keyword:‘’
	clinic_id: 1
	offset: 1
	limit: 10
	start_date:'2018-01-10'
	end_date:'2018-02-10'
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ |  诊所id | |
| keyword | String | ❌ |  搜索关键字 | |
| offset |number | ❌ |  分页（跳过个数） | 0|
| limit | number | ❌ |  分页（每页个数） | 6|
| start_date | String | ✅ |  开始时间 | |
| end_date | String | ✅ |  结束时间 | |

**应答包示例**

```
{
    "code": "200",
    "msg": ""
    “data”:[
    {
      "birthday": "19920706",
      "created_time": "2018-08-13T01:13:34.634212+08:00",
      "department_name": "牙科",
      "doctor_name": "胡一天",
      "patient_id": 3,
      "patient_name": "查康",
      "refund_money": -1,
      "refund_people": "超级管理员",
      "sex": 1
    }]
    "page_info":{
       limit: 6
       offset: 0
       total:  10
    }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ |  生日 | |
| department_name | String | ✅ |  科室名称 | |
| doctor_name | String | ✅ |  医生名称 | |
| patient_id | String | ✅ |  病人id | |
| patient_name | String | ✅ |  病人名称 | |
| refund_money |number | ✅ |  退费金额 | |
| refund_people | String | ✅ |  退费人员 | |
| sex | String | ✅ |  病人性别 | |
| created_time | String | ✅ |  挂号记录退费时间 | |
--

</br>
<h3>8.9 获取交易流水日报表

```
请求地址：/charge/business/transaction
```
**请求包示例**

```
{
	oprationName:‘’
	patientName: ''
	offset: 1
	limit: 10
	start_date:'2018-01-10'
	end_date:'2018-02-10'
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| oprationName | String | ❌ |  操作员 | |
| patientName | String | ❌ |  患者名称 | |
| offset |number | ❌ |  分页（跳过个数） | 0|
| limit | number | ❌ |  分页（每页个数） | 6|
| start_date | String | ✅ |  开始时间 | |
| end_date | String | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "alipay": 0,
      "balance_money": 1,
      "bank": 0,
      "bonus_points_money": 0,
      "cash": 1,
      "clinic_patient_id": null,
      "created_time": "2018-07-26T18:37:30.330244+08:00",
      "deleted_time": null,
      "department_id": null,
      "departmentname": null,
      "derate_money": 0,
      "diagnosis_treatment_cost": 0,
      "diagnosis_treatment_fee": 0,
      "discount_money": 0,
      "doctor_id": null,
      "doctorname": null,
      "examination_cost": 0,
      "examination_fee": 0,
      "id": 9,
      "in_out": "in",
      "labortory_cost": 0,
      "labortory_fee": 0,
      "material_cost": 0,
      "material_fee": 0,
      "medical_money": 0,
      "on_credit_money": 0,
      "operation": "超级管理员",
      "operation_id": 1,
      "other_cost": 0,
      "other_fee": 0,
      "out_refund_no": null,
      "out_trade_no": "T1201807261837307928",
      "patient_id": null,
      "patientname": null,
      "pay_record_id": 8,
      "pid": null,
      "record_type": 2,
      "retail_cost": 0,
      "retail_fee": 1,
      "total_money": 1,
      "traditional_medical_cost": 0,
      "traditional_medical_fee": 0,
      "treatment_cost": 0,
      "treatment_fee": 0,
      "updated_time": "2018-07-26T18:37:30.330244+08:00",
      "voucher_money": 0,
      "wechat": 0,
      "western_medicine_cost": 0,
      "western_medicine_fee": 0
    }
  ],
  "page_info": {
    "alipay": 0,
    "balance_money": 39084858,
    "bank": 0,
    "bonus_points_money": 0,
    "cash": 39084856,
    "derate_money": 0,
    "diagnosis_treatment_fee": 20500,
    "discount_money": 0,
    "examination_fee": 627000,
    "labortory_fee": 464500,
    "limit": "1",
    "material_fee": 9024103,
    "medical_money": 0,
    "offset": "0",
    "on_credit_money": 0,
    "other_fee": 12003,
    "retail_fee": 0,
    "total": 52,
    "total_money": 39084858,
    "traditional_medical_fee": 28400000,
    "treatment_fee": 124000,
    "voucher_money": 0,
    "wechat": 2,
    "western_medicine_fee": 412752
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| alipay |number | ✅ |  支付宝金额 | |
| balance_money | number | ✅ |  支付金额 | |
| bank | number | ✅ |  银行卡金额 | |
| cash | number | ✅ |  现金 | |
| wechat | number | ✅ |  微信金额 | |
| bonus\_points_money | number | ✅ |  积分金额 | |
|  derate_money | number | ✅ |  减免金额 | |
|  discount_money | number | ✅ |  折扣金额| |
|  medical_money | number | ✅ |  医保金额| |
|  other_fee | number | ✅ |  其他费用金额 | |
|  voucher_money | number | ✅ |  抵金券金额 | |
|  medical_money | number | ✅ |  医保金额| |
| clinic\_patient_id | number | ✅ |  诊所病人id | |
| departmentname | number | ✅ |  科室名称 | |
| operation | number | ✅ |  操作者 | |
| in_out | number | ✅ |  类型（进，出） | |
| diagnosis\_treatment_cost | number | ✅ | 诊疗金额_成本 | |
| diagnosis\_treatment_fee | number | ✅ |  诊疗金额 | |
| examination_fee | number | ✅ |  检查费 | |
| examination_cost | number | ✅ |  检查费_成本 | |
| labortory_fee | number | ✅ |  检验费 | |
| labortory_cost | number | ✅ |  检验费_成本 | |
| material_fee | number | ✅ |  材料费 | |
| material_cost | number | ✅ |  材料费_成本 | |
| other_fee | number | ✅ |  其他费 | |
| other_cost | number | ✅ |  其他费用_成本 | |
| retail_fee | number | ✅ |  零售费用 | |
| retail_cost | number | ✅ |  零售_成本 | |
| traditional\_medical_fee | number | ✅ |  中药费用 | |
| traditional\_medical_cost | number | ✅ |  中药费用_成本| |
| treatment_fee | number | ✅ |  治疗费 | |
| treatment_cost | number | ✅ | 治疗费_成本 | |
| western\_medicine_fee | number | ✅ |  西药费 | |
| western\_medicine_cost | number | ✅ |  西药费_成本 | |
| page\_info.alipay | number | ✅ |  支付宝（合计） | |
| page\_info.wechat | number | ✅ |  微信（合计） | |
| page\_info.cash | number | ✅ |  现金（合计） | |
| page\_info.bank | number | ✅ |  银行卡（合计） | |
--

</br>
<h3>8.10 获取分析类报表

```
请求地址：/charge/business/transaction/analysis
```
**请求包示例**

```
{
	offset: 1
	limit: 10
	start_date:'2018-01-10'
	end_date:'2018-02-10'
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| offset |number | ❌ |  分页（跳过个数） | 0|
| limit | number | ❌ |  分页（每页个数） | 10|
| start_date | String | ✅ |  开始时间 | |
| end_date | String | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "alipay": 0,
      "balance_money": 4165600,
      "bank": 0,
      "bonus_points_money": 0,
      "cash": 4165600,
      "created_time": "2018-07-26",
      "derate_money": 0,
      "diagnosis_treatment_fee": 1500,
      "discount_money": 0,
      "examination_fee": 72000,
      "labortory_fee": 57000,
      "material_fee": 1000100,
      "medical_money": 0,
      "on_credit_money": 0,
      "other_fee": 2000,
      "retail_fee": 0,
      "total_money": 4165600,
      "traditional_medical_fee": 2950000,
      "treatment_fee": 27000,
      "voucher_money": 0,
      "wechat": 0,
      "western_medicine_fee": 56000
    }
  ],
  "page_info": {
    "alipay": 0,
    "balance_money": 39084858,
    "bank": 0,
    "bonus_points_money": 0,
    "cash": 39084856,
    "derate_money": 0,
    "diagnosis_treatment_fee": 20500,
    "discount_money": 0,
    "examination_fee": 627000,
    "labortory_fee": 464500,
    "limit": "1",
    "material_fee": 9024103,
    "medical_money": 0,
    "offset": "0",
    "on_credit_money": 0,
    "other_fee": 12003,
    "retail_fee": 0,
    "total": 52,
    "total_money": 39084858,
    "traditional_medical_fee": 28400000,
    "treatment_fee": 124000,
    "voucher_money": 0,
    "wechat": 2,
    "western_medicine_fee": 412752
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| alipay |number | ✅ |  支付宝金额 | |
| balance_money | number | ✅ |  支付金额 | |
| bank | number | ✅ |  银行卡金额 | |
| cash | number | ✅ |  现金 | |
| wechat | number | ✅ |  微信金额 | |
| bonus\_points_money | number | ✅ |  积分金额 | |
|  derate_money | number | ✅ |  减免金额 | |
|  discount_money | number | ✅ |  折扣金额| |
|  medical_money | number | ✅ |  医保金额| |
|  other_fee | number | ✅ |  其他费用金额 | |
|  voucher_money | number | ✅ |  抵金券金额 | |
|  medical_money | number | ✅ |  医保金额| |
| operation | number | ✅ |  操作者 | |
| diagnosis\_treatment_fee | number | ✅ |  诊疗金额 | |
| examination_fee | number | ✅ |  检查费 | |
| labortory_fee | number | ✅ |  检验费 | |
| material_fee | number | ✅ |  材料费 | |
| other_fee | number | ✅ |  其他费 | |
| retail_fee | number | ✅ |  零售费用 | |
| traditional\_medical_fee | number | ✅ |  中药费用 | |
| treatment_fee | number | ✅ |  治疗费 | |
| western\_medicine_fee | number | ✅ |  西药费 | |
| page\_info.alipay | number | ✅ |  支付宝（合计） | |
| page\_info.wechat | number | ✅ |  微信（合计） | |
| page\_info.cash | number | ✅ |  现金（合计） | |
| page\_info.bank | number | ✅ |  银行卡（合计） | |
--

</br>
<h3>8.11 获取交易流水月报表

```
请求地址：/charge/business/transaction/month
```
**请求包示例**

```
{
	offset: 1
	limit: 10
	start_date:'2018-01-10'
	end_date:'2018-02-10'
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| offset |number | ❌ |  分页（跳过个数） | 0|
| limit | number | ❌ |  分页（每页个数） | 10|
| start_date | String | ✅ |  开始时间 | |
| end_date | String | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "alipay": 0,
      "balance_money": 1,
      "bank": 0,
      "bonus_points_money": 0,
      "cash": 1,
      "clinic_patient_id": null,
      "created_time": "2018-07-26T18:37:30.330244+08:00",
      "deleted_time": null,
      "department_id": null,
      "departmentname": null,
      "derate_money": 0,
      "diagnosis_treatment_cost": 0,
      "diagnosis_treatment_fee": 0,
      "discount_money": 0,
      "doctor_id": null,
      "doctorname": null,
      "examination_cost": 0,
      "examination_fee": 0,
      "id": 9,
      "in_out": "in",
      "labortory_cost": 0,
      "labortory_fee": 0,
      "material_cost": 0,
      "material_fee": 0,
      "medical_money": 0,
      "on_credit_money": 0,
      "operation": "超级管理员",
      "operation_id": 1,
      "other_cost": 0,
      "other_fee": 0,
      "out_refund_no": null,
      "out_trade_no": "T1201807261837307928",
      "patient_id": null,
      "patientname": null,
      "pay_record_id": 8,
      "pid": null,
      "record_type": 2,
      "retail_cost": 0,
      "retail_fee": 1,
      "total_money": 1,
      "traditional_medical_cost": 0,
      "traditional_medical_fee": 0,
      "treatment_cost": 0,
      "treatment_fee": 0,
      "updated_time": "2018-07-26T18:37:30.330244+08:00",
      "voucher_money": 0,
      "wechat": 0,
      "western_medicine_cost": 0,
      "western_medicine_fee": 0
    }
  ],
  "page_info": {
    "alipay": 0,
    "balance_money": 39084858,
    "bank": 0,
    "bonus_points_money": 0,
    "cash": 39084856,
    "derate_money": 0,
    "diagnosis_treatment_fee": 20500,
    "discount_money": 0,
    "examination_fee": 627000,
    "labortory_fee": 464500,
    "limit": "1",
    "material_fee": 9024103,
    "medical_money": 0,
    "offset": "0",
    "on_credit_money": 0,
    "other_fee": 12003,
    "retail_fee": 0,
    "total": 52,
    "total_money": 39084858,
    "traditional_medical_fee": 28400000,
    "treatment_fee": 124000,
    "voucher_money": 0,
    "wechat": 2,
    "western_medicine_fee": 412752
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| alipay |number | ✅ |  支付宝金额 | |
| balance_money | number | ✅ |  支付金额 | |
| bank | number | ✅ |  银行卡金额 | |
| cash | number | ✅ |  现金 | |
| wechat | number | ✅ |  微信金额 | |
| bonus\_points_money | number | ✅ |  积分金额 | |
|  derate_money | number | ✅ |  减免金额 | |
|  discount_money | number | ✅ |  折扣金额| |
|  medical_money | number | ✅ |  医保金额| |
|  other_fee | number | ✅ |  其他费用金额 | |
|  voucher_money | number | ✅ |  抵金券金额 | |
|  medical_money | number | ✅ |  医保金额| |
| clinic\_patient_id | number | ✅ |  诊所病人id | |
| departmentname | number | ✅ |  科室名称 | |
| operation | number | ✅ |  操作者 | |
| in_out | number | ✅ |  类型（进，出） | |
| diagnosis\_treatment_cost | number | ✅ | 诊疗金额_成本 | |
| diagnosis\_treatment_fee | number | ✅ |  诊疗金额 | |
| examination_fee | number | ✅ |  检查费 | |
| examination_cost | number | ✅ |  检查费_成本 | |
| labortory_fee | number | ✅ |  检验费 | |
| labortory_cost | number | ✅ |  检验费_成本 | |
| material_fee | number | ✅ |  材料费 | |
| material_cost | number | ✅ |  材料费_成本 | |
| other_fee | number | ✅ |  其他费 | |
| other_cost | number | ✅ |  其他费用_成本 | |
| retail_fee | number | ✅ |  零售费用 | |
| retail_cost | number | ✅ |  零售_成本 | |
| traditional\_medical_fee | number | ✅ |  中药费用 | |
| traditional\_medical_cost | number | ✅ |  中药费用_成本| |
| treatment_fee | number | ✅ |  治疗费 | |
| treatment_cost | number | ✅ | 治疗费_成本 | |
| western\_medicine_fee | number | ✅ |  西药费 | |
| western\_medicine_cost | number | ✅ |  西药费_成本 | |
| page\_info.alipay | number | ✅ |  支付宝（合计） | |
| page\_info.wechat | number | ✅ |  微信（合计） | |
| page\_info.cash | number | ✅ |  现金（合计） | |
| page\_info.bank | number | ✅ |  银行卡（合计） | |
--


</br>
<h3>8.12 获取交易详情

```
请求地址：/charge/business/transaction/detail
```
**请求包示例**

```
{
	offset: 1
	limit: 10
	start_date:'2018-01-10'
	end_date:'2018-02-10'
	patientName: ''
	phone: ''
	porjectName: ''
	in_out: ''
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| offset |number | ❌ |  分页（跳过个数） | 0|
| limit | number | ❌ |  分页（每页个数） | 10|
| start_date | String | ✅ |  开始时间 | |
| end_date | String | ✅ |  结束时间 | |
| patientName |number | ❌ |  患者名称 | |
| phone | number | ❌ |  患者手机号 | |
| porjectName | number | ❌ |  项目名称 | |
| in_out | number | ❌ |  进出 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "amount": null,
      "birthday": null,
      "charge_project_type": null,
      "created_time": "2018-07-26T18:37:30.330244+08:00",
      "deptname": null,
      "doctorname": null,
      "drug_mount": 1,
      "drug_name": "维生素AD胶丸",
      "drug_price": 1,
      "drug_total": 1,
      "drug_unit": "瓶",
      "fee": null,
      "name": null,
      "operarion": "超级管理员",
      "out_trade_no": "T1201807261837307928",
      "patientname": null,
      "phone": null,
      "price": null,
      "record_type": 2,
      "sex": null,
      "total": null,
      "unit": null,
      "visit_date": null
    }
  ],
  "page_info": {
    "banance_fee": 39084858,
    "limit": "1",
    "offset": "0",
    "total": 167,
    "total_fee": 39084858
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| amount |number | ✅ |  数量| |
| birthday |number | ✅ |  生日| |
| charge_project_type |number | ✅ |  收费类型| |
| deptname |number | ✅ |  科室名称| |
| doctorname |number | ✅ |  医生名称| |
| drug_mount |number | ✅ |  药品数量| |
| drug_name |number | ✅ |  药品名称| |
| drug_price |number | ✅ | 药品单价| |
| drug_total |number | ✅ |  药品总价钱| |
| fee |number | ✅ |  金额| |
| name |number | ✅ |  名称| |
| operarion |number | ✅ |  操作员| |
| out_trade_no |number | ✅ |  系统交易号| |
| patientname |number | ✅ |  患者名称| |
| phone |number | ✅ |  手机号| |
| price |number | ✅ |  单价| |
| record_type |number | ✅ |  记录类型（1-门诊费用，2-药品零售）| |
| sex |number | ✅ |  性别| |
| total |number | ✅ |  金额| |
| unit |number | ✅ |  单位| |
| visit_date |number | ✅ |  就诊日期| |
| page\_info.banance_fee |number | ✅ |  合计（实收金额）| |
| page\_info.total_fee |number | ✅ |  合计（应收金额| |
--

</br>
<h3>8.13 获取交易订单信息

```
请求地址：/charge/managerment/order
```
**请求包示例**

```
{
	offset: 1
	limit: 10
	start_date:'2018-01-10'
	end_date:'2018-02-10'
	clinic_id: ''
	keyword: ''
	orderType: ''
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| offset |number | ❌ |  分页（跳过个数） | 0|
| limit | number | ❌ |  分页（每页个数） | 10|
| start_date | String | ✅ |  开始时间 | |
| end_date | String | ✅ |  结束时间 | |
| clinic_id |number | ✅ |  诊所id | |
| keyword | number | ❌ |  关键词 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "balance_money": -1,
      "created_time": "2018-08-13T01:13:34.634212+08:00",
      "number": "R2201808130113349443",
      "operation": "超级管理员",
      "order_status": "SUCCESS",
      "order_type": "门诊退费",
      "patient_id": 3,
      "patient_name": "查康",
      "pay_method_code": "cash"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 141
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| balance_money |number | ✅ |  支付金额| |
| created_time |number | ✅ |  创建时间| |
| number |number | ✅ |  订单号| |
| operation |number | ✅ |  操作员| |
| order_status |number | ✅ |  订单状态| |
| order_type |number | ✅ |  订单类型| |
| patient_id |number | ✅ |  就诊人id| |
| patient_name |number | ✅ |  就诊人姓名| |
| pay_method_code |number | ✅ |  支付方式| |

--

9 挂账模块
--------

</br>
<h3>9.1 有挂账的分诊记录

```
请求地址：/onCredit/traigePatient/list
```
**请求包示例**

```
{
	keyword:‘’
	clinic_id: 1
	start_date: ’2018-01-01‘
	end_date: '2018-01-01'
	offset: 0
	limit:1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ✅ |  搜索关键词 | |
| clinic_id | String | ✅ |  诊所id | |
| start_date | String | ✅ |  开始日期 | |
| end_date | String | ✅ |  结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "charge_total_fee": 100,
      "clinic_triage_patient_id": 1,
      "clinic_patient_id":1,
      "operation": "超级管理员",
      "updated_time": "2018-08-13T01:13:34.634212+08:00",
      "visit_date": '2018-01-01'
      "patient_name": '张三'
      ”doctor_name“: '李四'
      ”department_name“ : '骨科'
      ”patient_id“: 1
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 141
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| charge\_total_fee |number | ✅ |  费用| |
| clinic\_triage\_patient_id |number | ✅ |  分诊id| |
| clinic\_patient_id |number | ✅ |  诊所病人id| |
| operation |number | ✅ |  操作员| |
| updated_time |number | ✅ |  更新时间| |
| visit_date |number | ✅ |  就诊日期| |
| patient_name |number | ✅ |  患者姓名| |
| doctor_name |number | ✅ |  医生姓名| |
| department_name |number | ✅ |  科室姓名| |
| patient_id |number | ✅ |  病人id| |


10 预约模块
--------

</br>
<h3>10.1 预约挂号

```
请求地址：/appointment/create
```
**请求包示例**

```
{
	paient_id:‘’
	cert_no: 1
	name: 1
	birthday: 10
	sex:1
	phone: 12881212
	province: ''
	city: ''
	district: ''
	address: ''
	profession: ''
	remark: ''
	patient_channel_id: 1
	clinic_id: 1
	doctor_visit_schedule_id: 1
	visit_type: 1
	personnel_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| name | String | ✅ |  姓名 | |
| birthday | String | ✅ |  生日 | |
| sex | String | ✅ |  性别 | |
| phone | String | ✅ |  手机号 | |
| patientChannelID | String | ✅ |  患者途径 | |
| clinicID | String | ✅ |  诊所id | |
| personnelID | String | ✅ |  操作员id | |
| paient_id | String | ❌ |  病人id | |
| cert_no | String | ❌ |  证件号 | |
| province | String | ❌ |  省 | |
| city | String | ❌ |  市 | |
| district | String | ❌ |  区 | 
| address | String | ❌ |  地址 | 
| profession | String | ❌ |  职业 | 
| remark | String | ❌ |  备注 | 
| doctor_visit_schedule_id | String | ❌ |  排班id | 
| visit_type | String | ❌ |  就诊类型 | 

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
|code | String | ✅ |  200 成功 | |
|msg | String | ✅ |  消息| |
--

11 药品模块
--------
</br>
<h3>11.1 添加药品

```
请求地址：/clinic_drug/ClinicDrugCreate
```
**请求包示例**

```
{
	clinic_id:1
	drug_class_id: 1
	name: 1
	specification: 1
	manu_factory_name: 1
	dose_form_name: 1
	print_name: 1
	license_no: 1
	type: 1
	py_code: 1
	barcode: 1
	status: true
	dosage: 1
	dosage_unit_name: 1
	preparation_count: 1
	preparation_count_unit_name: 1
	packing_unit_name: 1
	ret_price: 1
	buy_price: 1
	mini_dose: 1
	is_discount: true
	is_bulk_sales: 1
	bulk_sales_price: 1
	fetch_address: 1
	once_dose: 1
	once_dose_unit_name: 1
	route_administration_name: 1
	frequency_name: 1
	illustration: 1
	day_warning: 1
	stock_warning: 1
	english_name: 1
	sy_code: 1
	country_flag: 1
	self_flag: 1
	drug_flag: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ |  诊所id| |
| drug\_class_id | String | ✅ |  药品类型编码| |
| name | String | ✅ |  药品名称| |
| specification | String |❌ | 规格| |
| manu\_factory_name | String |❌ | 生产厂商| |
| dose\_form_name | String | ❌| 剂型| |
| print_name | String | ❌ |  打印名称| |
| license_no | String | ❌ |  国药准字| |
| type | String | ✅ |  0-西药 1-中药| |
| py_code | String | ❌ |  拼音码| |
| barcode | String | ❌ |  条形码| |
| status | boolean | ❌ |  启用状态| |
| dosage |number | ❌ |  剂量| |
| dosage\_unit_name | String | ❌ |  剂量单位| |
| preparation_count | number | ❌ |  制剂数量/包装量| |
| preparation\_count\_unit_name | String | ❌ |  制剂数量单位| |
| packing\_unit_name | String | ❌ |  包装单位| |
| ret_price | String | ✅ |  零售价| |
| buy_price | String | ❌ |  成本价| |
| mini_dose | String | ❌ |  最小剂量 | |
| is_discount | boolean | ❌ |  允许打折| false|
| is\_bulk_sales | boolean | ✅ |  是否允许拆零销售| false |
| bulk\_sales_price |number | ❌ |  拆零售价/最小剂量售价| |
| fetch_address | number | ✅ | 取药地点 0 本诊所，1外购 2， 代购| |
| once_dose | number | ❌ | 常用剂量| |
| once\_dose\_unit_name | number | ❌ | 常用剂量单位| |
| route\_administration_name | number | ❌ | 用药途径| |
| frequency_name | number | ❌ | 用药频率 | |
| illustration | number | ❌ | 说明 | |
| day_warning | number | ❌ | 效期预警天数 | |
| stock_warning | number | ❌ | 库存预警数 | |
| english_name | string | ❌ | 英文名称 | |
| sy_code | number | ❌ | 上药编码 | |
| country_flag | boolean | ❌ | 进口标识 | false |
| self_flag | boolean | ❌ | 自费标识 | false |
| drug_flag | boolean | ❌ | 毒麻标志 | false |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
  ”data“ : {
     clinic_drug_id: 1
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |
| clinic_drug_id |number | ✅ |  药品id| |

--

</br>
<h3>11.2 更新药品

```
请求地址：/clinic_drug/ClinicDrugUpdate
```

**请求包示例**

```
{
   clinic_drug_id : 1
	drug_class_id: 1
	name: 1
	specification: 1
	manu_factory_name: 1
	dose_form_name: 1
	print_name: 1
	license_no: 1
	type: 1
	py_code: 1
	barcode: 1
	status: true
	dosage: 1
	dosage_unit_name: 1
	preparation_count: 1
	preparation_count_unit_name: 1
	packing_unit_name: 1
	ret_price: 1
	buy_price: 1
	mini_dose: 1
	is_discount: true
	is_bulk_sales: 1
	bulk_sales_price: 1
	fetch_address: 1
	once_dose: 1
	once_dose_unit_name: 1
	route_administration_name: 1
	frequency_name: 1
	illustration: 1
	day_warning: 1
	stock_warning: 1
	english_name: 1
	sy_code: 1
	country_flag: 1
	self_flag: 1
	drug_flag: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_drug_id | String | ✅ |  药品id| |
| drug\_class_id | String | ❌ |  药品类型编码| |
| name | String | ❌ |  药品名称| |
| specification | String |❌ | 规格| |
| manu\_factory_name | String |❌ | 生产厂商| |
| dose\_form_name | String | ❌| 剂型| |
| print_name | String | ❌ |  打印名称| |
| license_no | String | ❌ |  国药准字| |
| type | String | ❌ |  0-西药 1-中药| |
| py_code | String | ❌ |  拼音码| |
| barcode | String | ❌ |  条形码| |
| status | boolean | ❌ |  启用状态| |
| dosage |number | ❌ |  剂量| |
| dosage\_unit_name | String | ❌ |  剂量单位| |
| preparation_count | number | ❌ |  制剂数量/包装量| |
| preparation\_count\_unit_name | String | ❌ |  制剂数量单位| |
| packing\_unit_name | String | ❌ |  包装单位| |
| ret_price | String | ❌ |  零售价| |
| buy_price | String | ❌ |  成本价| |
| mini_dose | String | ❌ |  最小剂量 | |
| is_discount | boolean | ❌ |  允许打折| false|
| is\_bulk_sales | boolean | ❌ |  是否允许拆零销售| false |
| bulk\_sales_price |number | ❌ |  拆零售价/最小剂量售价| |
| fetch_address | number | ❌ | 取药地点 0 本诊所，1外购 2， 代购| |
| once_dose | number | ❌ | 常用剂量| |
| once\_dose\_unit_name | number | ❌ | 常用剂量单位| |
| route\_administration_name | number | ❌ | 用药途径| |
| frequency_name | number | ❌ | 用药频率 | |
| illustration | number | ❌ | 说明 | |
| day_warning | number | ❌ | 效期预警天数 | |
| stock_warning | number | ❌ | 库存预警数 | |
| english_name | string | ❌ | 英文名称 | |
| sy_code | number | ❌ | 上药编码 | |
| country_flag | boolean | ❌ | 进口标识 | false |
| self_flag | boolean | ❌ | 自费标识 | false |
| drug_flag | boolean | ❌ | 毒麻标志 | false |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
  ”data“ : {
     clinic_drug_id: 1
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |
| clinic\_drug_id |number | ✅ |  药品id| |


</br>
<h3>11.3 启用或停止药品

```
请求地址：/clinic_drug/ClinicDrugOnOff
```
**请求包示例**

```
{
   clinic_id: 1
   clinic_drug_id: 1
   status: true
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| clinic\_drug_id |number | ✅ |  药品id| |
| status |boolean | ✅ |  开启状态| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.4 药品列表

```
请求地址：/clinic_drug/ClinicDrugList
```
**请求包示例**

```
{
   clinic_id: 1
   type: 1
   drug_class_id: true
   keyword: '关键字'
   status: true
   offset: 0
   limit: 10
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| type | number | ❌ |  类型| |
| drug_class_id |number | ❌ | 药品类型id| |
| keyword | string | ❌ |  关键字| |
| status | boolean | ❌ |  状态| |
| offset |number | ❌ | 跳过数 | 0|
| limit |number | ❌ | 每页数| 10 |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "buy_price": 50000,
      "clinic_drug_id": 12,
      "clinic_id": 1,
      "drug_name": "当归",
      "fetch_address": 0,
      "frequency_name": "1次/日 (2pm)",
      "illustration": "收入高达符号化",
      "is_discount": true,
      "manu_factory_name": "北京通县振兴饮片厂",
      "once_dose": null,
      "once_dose_unit_name": null,
      "packing_unit_name": "g",
      "py_code": "DGV",
      "ret_price": 10000,
      "route_administration_name": "口服                  ",
      "specification": "/kg",
      "status": true,
      "stock_amount": 9968,
      "type": 1
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 47
  }
}
```

**应答包参数说明**

同 11.1



</br>
<h3>11.5 查询药品库存信息

```
请求地址：/clinic_drug/ClinicDrugStock
```
**请求包示例**

```
{
   clinic_id: 1
   keyword: '关键字'
   status: true
   offset: 0
   limit: 10
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| keyword | string | ❌ |  关键字| |
| status | boolean | ❌ |  状态| |
| offset |number | ❌ | 跳过数 | 0|
| limit |number | ❌ | 每页数| 10 |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "buy_price": 50000,
      "clinic_drug_id": 12,
      "clinic_id": 1,
      "drug_name": "当归",
      "fetch_address": 0,
      "frequency_name": "1次/日 (2pm)",
      "illustration": "收入高达符号化",
      "is_discount": true,
      "manu_factory_name": "北京通县振兴饮片厂",
      "once_dose": null,
      "once_dose_unit_name": null,
      "packing_unit_name": "g",
      "py_code": "DGV",
      "ret_price": 10000,
      "route_administration_name": "口服                  ",
      "specification": "/kg",
      "status": true,
      "stock_amount": 9968,
      "type": 1
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 47
  }
}
```

**应答包参数说明**（同11.1）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| stock_amount |number | ✅ |  库存数量| |

</br>
<h3>11.6 药品详情 

```
请求地址：/clinic_drug/ClinicDrugDetail
```
**请求包示例**

```
{
   clinic_drug_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_drug_id |number | ✅ |  药品id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "barcode": "735636513513",
    "bulk_sales_price": null,
    "buy_price": 50000,
    "clinic_id": 1,
    "country_flag": null,
    "created_time": "2018-05-27T17:55:21.813179+08:00",
    "day_warning": 10,
    "deleted_time": null,
    "discount_price": 0,
    "dosage": 1,
    "dosage_unit_name": "g",
    "dose_form_name": "根茎类",
    "drug_class_id": null,
    "drug_flag": null,
    "english_name": null,
    "fetch_address": 0,
    "frequency_name": "1次/日 (2pm)",
    "id": 12,
    "illustration": "收入高达符号化",
    "is_bulk_sales": false,
    "is_discount": true,
    "license_no": "京卫药生证字20010132号        ",
    "manu_factory_name": "北京通县振兴饮片厂",
    "mini_dose": null,
    "name": "当归",
    "once_dose": null,
    "once_dose_unit_name": null,
    "packing_unit_name": "g",
    "preparation_count": null,
    "preparation_count_unit_name": null,
    "print_name": null,
    "py_code": "DGV",
    "ret_price": 10000,
    "route_administration_name": "口服                  ",
    "self_flag": null,
    "specification": "/kg",
    "status": true,
    "stock_warning": 50,
    "sy_code": null,
    "type": 1,
    "updated_time": "2018-05-27T17:55:21.813179+08:00"
  }
}
```

**应答包参数说明**

同 11.1

</br>
<h3>11.7 批量设置药品


```
请求地址：/clinic_drug/ClinicDrugBatchSetting
```
**请求包示例**

```
{
   day_warning: 1
   is_discount: false
   items : [
    {
       clinic_drug_id: 1
    }
   ]
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_drug_id |number | ✅ |  药品id| |
| day_warning |number | ✅ |  预警天数| |
| is_discount | boolean | ✅ |  允许折扣 | |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |

</br>
<h3>11.8 药品入库

```
请求地址：/clinic_drug/instock
```
**请求包示例**

```
{
   clinic_id: 1
   instock_operation_id: 1
   instock_way_name: 'dsaf'
   supplier_name: '1232'
   remark: ''
   instock_date: '2018-01-12'
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| instock_operation_id |number | ✅ |  入库操作员| |
| instock_way_name |number | ✅ |  入库方式| |
| supplier_name |number | ✅ |  供应商| |
| instock_date |number | ✅ |  入库日期| |
| remark |number | ✅ |  备注| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |

</br>
<h3>11.9 入库记录列表

```
请求地址：/clinic_drug/instockRecord
```
**请求包示例**

```
{
   clinic_id: 1
   start_date: '2018-01-01'
   end_date: '2019-01-01'
   order_number: 111
   offset: 0
   limit: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| start_date |number | ✅ |  开始日期| |
| end_date |number | ✅ |  结束日期| |
| order_number |number | ❌ |  盘点单号| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "drug_instock_record_id": 66,
      "instock_date": "2018-08-12T00:00:00Z",
      "instock_operation_name": "超级管理员",
      "instock_way_name": "采购入库",
      "order_number": "DRKD-1534084199",
      "supplier_name": "云南白药药厂",
      "verify_operation_name": "超级管理员",
      "verify_status": "02"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 31
  }
}
```

**应答包参数说明**（同11.8）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| verify_status | boolean | ✅ | 审核状态 01 未审核 02 已审核  | |

</br>
<h3>11.10 入库记录详情

```
请求地址：/clinic_drug/instockRecordDetail
```
**请求包示例**

```
{
   drug_instock_record_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_instock\_record_id |number | ✅ |  入库记录id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "created_time": "2018-08-12T22:29:59.104192+08:00",
    "drug_instock_record_id": 66,
    "instock_date": "2018-08-12T00:00:00Z",
    "instock_operation_id": 1,
    "instock_operation_name": "超级管理员",
    "instock_way_name": "采购入库",
    "items": [
      {
        "buy_price": 50000,
        "clinic_drug_id": 12,
        "drug_name": "当归",
        "eff_date": "2018-08-12T00:00:00Z",
        "instock_amount": 10000,
        "manu_factory_name": "北京通县振兴饮片厂",
        "packing_unit_name": "g",
        "ret_price": 10000,
        "serial": "12323"
      }
    ],
    "order_number": "DRKD-1534084199",
    "remark": null,
    "supplier_name": "云南白药药厂",
    "updated_time": "2018-08-12T22:30:16.97446+08:00",
    "verify_operation_id": 1,
    "verify_operation_name": "超级管理员",
    "verify_status": "02"
  }
}
```

**应答包参数说明** （同11.1和11.8）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |

</br>
<h3>11.11 入库记录修改

```
请求地址：/clinic_drug/instockUpdate
```
**请求包示例**

```
{
   clinic_id: 1
   drug_instock_record_id: 1
   instock_operation_id: 1
   instock_way_name: 1
   supplier_name: 1
   remark: 1
   instock_date: ‘2018-03-12’
   items: [
   {
      clinic_drug_id: 1
      instock_amount: 1
      buy_price: 1
      serial: 1
      eff_date: '2018-03-12'
   }
   ]
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| instock\_way_name |number | ✅ |  入库方式| |
| supplier_name |number | ✅ |  供应商| |
| instock_date |number | ✅ |  入库日期| |
| clinic_drug_id |number | ✅ |  药品id| |
| instock_amount |number | ✅ |  入库数量| |
| buy_price |number | ✅ |  成本价| |
| serial |number | ✅ |  入库批号| |
| eff_date |number | ✅ |  有效日期| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |

</br>
<h3>11.12 入库审核

```
请求地址：/clinic_drug/instockCheck
```
**请求包示例**

```
{
   drug_instock_record_id: 1
   verify_operation_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_instock\_record_id |number | ✅ |  入库记录id| |
| verify\_operation_id |number | ✅ |  审核人员id| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |

</br>
<h3>11.13 删除入库记录

```
请求地址：/clinic_drug/instockDelete
```
**请求包示例**

```
{
   drug_instock_record_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_instock\_record_id |number | ✅ |  入库记录id| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |

</br>
<h3>11.14 出库

```
请求地址：/clinic_drug/outstock
```
**请求包示例**

```
{
   clinic_id: 1
   outstock_operation_id: 1
   outstock_way_name: 1
   department_id: 1
   personnel_id: 1
   remark: ''
   outstock_date: '2018-01-01'
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| outstock\_operation_id |number | ✅ |  出库操作者| |
| outstock\_way_name |number | ✅ |  出库方式| |
| department_id |number | ✅ |  科室id| |
| personnel_id |number | ✅ |  医生id| |
| outstock_date |number | ✅ |  出库日期| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |

</br>
<h3>11.15 出库记录

```
请求地址：/clinic_drug/outstockRecord
```
**请求包示例**

```
{
   clinic_id: 1
   start_date: '2018-01-01'
   end_date: '2019-01-01'
   order_number: 111
   offset: 0
   limit: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| start_date |number | ✅ |  开始日期| |
| end_date |number | ✅ |  结束日期| |
| order_number |number | ❌ |  盘点单号| |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "department_name": "眼科",
      "drug_outstock_record_id": 11,
      "order_number": "DCKD-1533830113",
      "outstock_date": "2018-08-09T00:00:00Z",
      "outstock_operation_name": "超级管理员",
      "outstock_way_name": "科室领用",
      "personnel_name": "扁鹊",
      "verify_operation_name": null,
      "verify_status": "01"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 5
  }
}
```

**应答包参数说明** （同11.14）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.16 出库记录详情

```
请求地址：/clinic_drug/outstockRecordDetail
```
**请求包示例**

```
{
   drug_outstock_record_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_outstock\_record_id |number | ✅ |  出库记录idid| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "created_time": "2018-08-09T23:55:13.752401+08:00",
    "department_id": 2,
    "department_name": "眼科",
    "drug_outstock_record_id": 11,
    "items": [
      {
        "buy_price": 600,
        "drug_name": "巴米尔(阿司匹林泡腾片)",
        "drug_stock_id": 26,
        "eff_date": "2018-08-01T00:00:00Z",
        "manu_factory_name": "北京华丰制药公司",
        "outstock_amount": 1,
        "packing_unit_name": "盒",
        "ret_price": 1000,
        "serial": "201808011002",
        "stock_amount": 100,
        "supplier_name": "广州白云药厂"
      }
    ],
    "order_number": "DCKD-1533830113",
    "outstock_date": "2018-08-09T00:00:00Z",
    "outstock_operation_id": 1,
    "outstock_operation_name": "超级管理员",
    "outstock_way_name": "科室领用",
    "personnel_id": 2,
    "personnel_name": "扁鹊",
    "remark": null,
    "updated_time": "2018-08-09T23:55:13.752401+08:00",
    "verify_operation_id": null,
    "verify_operation_name": null,
    "verify_status": "01"
  }
}
```

**应答包参数说明** （同11.1 和 11.14）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.17 更新出库

```
请求地址：/clinic_drug/outstockUpdate
```
**请求包示例**

```
{
   clinic_id: 1
   drug_outstock_record_id: 1
   outstock_operation_id: 1
   outstock_way_name: 1
   department_id: 1
   personnel_id: 1
   remark: ''
   outstock_date: '2018-01-01'
   items: [
   {
      drug_stock_id: 1
      outstock_amount: 1
   }
   ]
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| drug\_outstock\_record_id |number | ✅ |  处理记录id| |
| outstock\_operation_id |number | ✅ |  出库操作员id| |
| department_id |number | ✅ |  科室id| |
| personnel_id |number | ✅ |  医生id| |
| outstock_date |number | ✅ |  出库日期| |
| drug_stock_id |number | ✅ |  库存id | |
| outstock_amount |number | ✅ |  出库数量 | |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.18 出库审核

```
请求地址：/clinic_drug/outstockCheck
```
**请求包示例**

```
{
   drug_outstock_record_id: 1
   verify_operation_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_outstock\_record_id |number | ✅ |  出库记录id| |
| verify\_operation_id |number | ✅ |  审核人员id| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.19 删除出库记录

```
请求地址：/clinic_drug/outstockDelete
```
**请求包示例**

```
{
   drug_outstock_record_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_outstock\_record_id |number | ✅ |  出库记录id| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.20 库存列表

```
请求地址：/clinic_drug/DrugStockList
```
**请求包示例**

```
{
   clinic_id: 1
   keyword: 1
   supplier_name: 1
   amount: 1
   date_warning: 1
   offset: 1
   limit: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| keyword |number | ❌ |  关键字搜索| |
| supplier_name |number | ❌ |  供应商| |
| amount |number | ❌ | 库存数量 | |
| date_warning |number | ❌ | 预警天数  | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "buy_price": 5000,
      "day_warning": 10,
      "drug_stock_id": 20,
      "eff_date": "2018-08-01T00:00:00Z",
      "manu_factory_name": "北京鹤延龄饮片厂",
      "name": "川贝母",
      "packing_unit_name": "g",
      "ret_price": 1000,
      "serial": "20180801004",
      "specification": "H01019/kg",
      "stock_amount": -90,
      "stock_warning": null,
      "supplier_name": "广州白云药厂"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 40
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| buy_price |number | ✅ |  成本价| |
| day_warning |number | ✅ |  预警天数| |
| drug\_stock_id |number | ✅ |  库存id| |
| eff_date | string | ✅ |  有效期| |
| manu\_factory_name |string | ✅ |  生产厂商| |
| name | string | ✅ |  药品名称| |
| packing\_unit_name | string | ✅ | 包装单位| |
| ret_price | string | ✅ | 零售价| |
| specification | string | ✅ | 规格| |
| stock_amount | string | ✅ | 库存| |
| stock_warning | string | ✅ | 库存预警| |
| supplier_name | string | ✅ | 供应商| |


</br>
<h3>11.21 创建西药处方模板

```
请求地址：/clinic_drug/PrescriptionWesternPatientModelCreate
```
**请求包示例**

```
{
   model_name: ‘’
   is_common: ture
   operation_id: 1
   item: [
     {
         clinic_drug_id : 1
         once_dose: 1
         once_dose_unit_name: ''
         route_administration_name: ''
         frequency_name: ''
         amount: 1
         illustration: ''
         fetch_address: ''
         eff_day: '2018-01-01'
     }
   ]
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| model_name |number | ✅ |  模板名称| |
| is_common | boolean | ✅ | 是否通用| |
| operation_id |number | ✅ |  操作员| |
| clinic_drug_id |number | ✅ |  药品id| |
| once_dose |number | ✅ |  常用剂量| |
| once_dose_unit_name |string | ✅ |  剂量单位| |
| route_administration_name | string | ✅ |  用药途径| |
| frequency_name | string | ✅ |  用药频率| |
| amount | string | ✅ |  数量| |
| illustration | string | ✅ | 说明 | |
| fetch_address | string | ✅ | 取药地点 0 本诊所 1 外购 | |
| eff_day | string | ✅ |  有效天数| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.22 西药处方模板列表

```
请求地址：/clinic_drug/PrescriptionWesternPatientModelList
```
**请求包示例**

```
{
   keyword: ’‘
   is_common: false
   operation_id: 1
   offset: 0
   limit 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword |number | ❌ |  关键字| |
| is_common | boolean | ❌ |  是否通用| |
| operation_id | boolean | ❌ |  操作员id| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "model_name": "新入库",
      "prescription_patient_model_id": 16,
      "operation_name": "胡一天",
      "is_common": true,
      "route_administration_name": null,
      "eff_day": null,
      "amount": null,
      "frequency_name": null,
      "fetch_address": null,
      "medicine_illustration": null,
      "created_time": "2018-08-01T14:18:26.605669+08:00",
      "updated_time": "2018-08-01T14:18:26.605669+08:00",
      "items": [
        {
          "clinic_drug_id": 13,
          "type": 0,
          "drug_name": "维生素AD胶丸",
          "specification": "100粒/瓶",
          "stock_amount": 101,
          "once_dose": 1,
          "once_dose_unit_name": "粒",
          "route_administration_name": "口服                  ",
          "frequency_name": "3次/日 (8-12-4)",
          "eff_day": 5,
          "amount": 5,
          "packing_unit_name": "瓶",
          "fetch_address": 0,
          "illustration": "新入库",
          "special_illustration": null
        }
      ]
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 5
  }
}
```

**应答包参数说明** （同11.21和 11.1）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.23 查询个人西药处方模板

```
请求地址：/clinic_drug/PrescriptionWesternPersonalPatientModelList
```
**请求包示例**

```
{
   keyword: ’‘
   is_common: false
   operation_id: 1
   offset: 0
   limit 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword |number | ❌ |  关键字| |
| is_common | boolean | ❌ |  是否通用| |
| operation_id | boolean | ✅ |  操作员id| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "model_name": "新入库",
      "prescription_patient_model_id": 16,
      "operation_name": "胡一天",
      "is_common": true,
      "route_administration_name": null,
      "eff_day": null,
      "amount": null,
      "frequency_name": null,
      "fetch_address": null,
      "medicine_illustration": null,
      "created_time": "2018-08-01T14:18:26.605669+08:00",
      "updated_time": "2018-08-01T14:18:26.605669+08:00",
      "items": [
        {
          "clinic_drug_id": 13,
          "type": 0,
          "drug_name": "维生素AD胶丸",
          "specification": "100粒/瓶",
          "stock_amount": 101,
          "once_dose": 1,
          "once_dose_unit_name": "粒",
          "route_administration_name": "口服                  ",
          "frequency_name": "3次/日 (8-12-4)",
          "eff_day": 5,
          "amount": 5,
          "packing_unit_name": "瓶",
          "fetch_address": 0,
          "illustration": "新入库",
          "special_illustration": null
        }
      ]
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 5
  }
}
```

**应答包参数说明** （同上）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.24 西药处方模板详情

```
请求地址：/clinic_drug/PrescriptionWesternPatientModelDetail
```
**请求包示例**

```
{
   prescription_patient_model_id: 16
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| prescription\_patient\_model_id |number | ✅ |  处方模板id | |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "is_common": true,
    "items": [
      {
        "amount": 5,
        "clinic_drug_id": 58,
        "created_time": "2018-08-01T14:18:26.605669+08:00",
        "deleted_time": null,
        "drug_name": "叶酸片",
        "eff_day": 5,
        "fetch_address": 0,
        "frequency_name": "1次/日 (8am)",
        "id": 12,
        "illustration": "新入库",
        "once_dose": 1,
        "once_dose_unit_name": "片",
        "prescription_western_patient_model_id": 16,
        "route_administration_name": "口服                  ",
        "updated_time": "2018-08-01T14:18:26.605669+08:00"
      }
    ],
    "model_name": "新入库",
    "prescription_patient_model_id": 16,
    "status": true
  }
}
```

**应答包参数说明** （同上）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.25 修改西药处方模板

```
请求地址：/clinic_drug/PrescriptionWesternPatientModelUpdate
```
**请求包示例**

```
{
   prescription_patient_model_id: 1
   model_name: ''
   is_common: false
   operation_id: 1
   items: [
   {
       clinic_drug_id: 1
       once_dose: 1
       once_dose_unit_name: 1
       route_administration_name: 1
       frequency_name: 1
       amount: 1
       illustration: 1
       fetch_address: 1
       eff_day: 1
   }
   ]
}
```

**请求包参数说明** (同11.21)


**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.26 删除西药处方模板

```
请求地址：/clinic_drug/PrescriptionWesternPatientModelDelete
```
**请求包示例**

```
{
   prescription_patient_model_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| prescription\_patient\_model_id |number | ✅ |  西药处方模板id| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.27 创建中药处方模板

```
请求地址：/clinic_drug/PrescriptionChinesePatientModelCreate
```
**请求包示例**

```
{
   model_name: ''
   is_common: true
   route_administration_name: '用药途径'
   frequency_name: '用药频率'
   amount: 1
   fetch_address: '取药地点'
   eff_day: '有效期’
   medicine_illustration: '服药说明'
   operation_id: 1
   items: [{
       clinic_drug_id: 1
       once_dose: 1
       once_dose_unit_name: ''
       special_illustration: ''
   }]
   
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| model_name |number | ✅ |  mo| |
| is_common |boolean | ✅ |  是否通用| |
| route\_administration_name |boolean | ✅ |  用药途径| |
| frequency_name |boolean | ✅ |  用药频率| |
| amount |bumber | ✅ |  数量| |
| fetch_address |string | ✅ |  取药地点| |
| eff_day | string | ✅ |  有效期| |
| medicine_illustration | string | ✅ | 说明| ||
| operation_id | string | ✅ | 操作员id| |
| clinic\_drug_id | string | ✅ |  诊所药品id| |
| once_dose | string | ✅ |  常用剂量| |
| once\_dose\_unit_name | string | ✅ |  剂量单位| |
| special_illustration | string | ✅ |  说明| |

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.28 中药处方模板列表

```
请求地址：/clinic_drug/PrescriptionChinesePatientModelList
```
**请求包示例**

```
{
   keyword: 1
   is_common: 1
   operation_id 1
   offset： 1
   limit: 10
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword |number | ✅ |  关键词| |
| is_common |boolean | ✅ |  是否通用| |
| operation_id |boolean | ✅ | 操作员id| |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "model_name": "新入库",
      "prescription_patient_model_id": 6,
      "operation_name": "胡一天",
      "is_common": true,
      "route_administration_name": "水煎服",
      "eff_day": 5,
      "amount": 5,
      "frequency_name": null,
      "fetch_address": 0,
      "medicine_illustration": "",
      "created_time": "2018-08-01T14:19:41.760702+08:00",
      "updated_time": "2018-08-01T14:19:41.760702+08:00",
      "items": [
        {
          "clinic_drug_id": 12,
          "type": 1,
          "drug_name": "当归",
          "specification": null,
          "stock_amount": 9968,
          "once_dose": 10,
          "once_dose_unit_name": "g",
          "route_administration_name": null,
          "frequency_name": null,
          "eff_day": null,
          "amount": 5,
          "packing_unit_name": null,
          "fetch_address": null,
          "illustration": null,
          "special_illustration": null
        }
      ]
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 5
  }
}
```

**应答包参数说明** (同 11.1 和 11.27)

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.29 个人中药处方模板

```
请求地址：/clinic_drug/PrescriptionChinesePersonalPatientModelList
```
**请求包示例**

```
{
   keyword: 1
   is_common: 1
   operation_id 1
   offset： 1
   limit: 10
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword |number | ✅ |  关键词| |
| is_common |boolean | ✅ |  是否通用| |
| operation_id |boolean | ✅ | 操作员id| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "model_name": "新入库",
      "prescription_patient_model_id": 6,
      "operation_name": "胡一天",
      "is_common": true,
      "route_administration_name": "水煎服",
      "eff_day": 5,
      "amount": 5,
      "frequency_name": null,
      "fetch_address": 0,
      "medicine_illustration": "",
      "created_time": "2018-08-01T14:19:41.760702+08:00",
      "updated_time": "2018-08-01T14:19:41.760702+08:00",
      "items": [
        {
          "clinic_drug_id": 12,
          "type": 1,
          "drug_name": "当归",
          "specification": null,
          "stock_amount": 9968,
          "once_dose": 10,
          "once_dose_unit_name": "g",
          "route_administration_name": null,
          "frequency_name": null,
          "eff_day": null,
          "amount": 5,
          "packing_unit_name": null,
          "fetch_address": null,
          "illustration": null,
          "special_illustration": null
        }
      ]
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 5
  }
}
```

**应答包参数说明** （同上）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.30 中药处方模板详情

```
请求地址：/clinic_drug/PrescriptionChinesePatientModelDetail
```
**请求包示例**

```
{
   prescription_patient_model_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| prescription\_patient\_model_id |number | ✅ |  模板id| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "amount": 5,
    "eff_day": 5,
    "fetch_address": 0,
    "frequency_name": "1次/日 (8am)",
    "is_common": true,
    "items": [
      {
        "amount": 5,
        "clinic_drug_id": 12,
        "created_time": "2018-08-01T14:19:41.760702+08:00",
        "deleted_time": null,
        "drug_name": "当归",
        "id": 22,
        "once_dose": 10,
        "once_dose_unit_name": "g",
        "prescription_chinese_patient_model_id": 6,
        "special_illustration": null,
        "updated_time": "2018-08-01T14:19:41.760702+08:00"
      }
    ],
    "medicine_illustration": "",
    "model_name": "新入库",
    "prescription_patient_model_id": 6,
    "route_administration_name": "水煎服",
    "status": true
  }
}
```

**应答包参数说明** (同上)

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.31 更新中药处方模板

```
请求地址：/clinic_drug/PrescriptionChinesePatientModelUpdate
```
**请求包示例**

```
{
   prescription_patient_model_id: 1
   model_name: ''
   is_common: true
   route_administration_name: '用药途径'
   frequency_name: '用药频率'
   amount: 1
   fetch_address: '取药地点'
   eff_day: '有效期’
   medicine_illustration: '服药说明'
   operation_id: 1
   items: [{
       clinic_drug_id: 1
       once_dose: 1
       once_dose_unit_name: ''
       special_illustration: ''
   }]
   
}
```

**请求包参数说明** (同 11.27)

**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.32 删除中药模板

```
请求地址：/clinic_drug/PrescriptionChinesePatientModelDelete
```
**请求包示例**

```
{
   prescription_patient_model_id: 1   
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| prescription\_patient\_model_id |number | ✅ |  处方模板id| |


**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.33 新增药房盘点

```
请求地址：/clinic_drug/DrugInventoryCreate
```
**请求包示例**

```
{
   clinic_id: 1   
   inventory_operation_id: 1
   items: [{
       drug_stock_id: 1
       actual_amount: 1
   }]
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| inventory_operation_id |number | ✅ |  判断id| |
| drug_stock_id |number | ✅ |  库存id| |
| actual_amount |number | ✅ |  实际数量| |


**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.34 药房盘点记录表

```
请求地址：/clinic_drug/DrugInventoryList
```
**请求包示例**

```
{
   clinic_id: 1 
   start_date: "2018-01-10"
   end_date: '2018-01-30'
   offset: 0
   limit: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| start_date |string | ✅ |  开始日期| |
| end_date | string | ✅ |  结束日期| |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "drug_inventory_record_id": 5,
      "inventory_date": "2018-08-09T00:00:00Z",
      "inventory_operation_name": "超级管理员",
      "order_number": "DPD-1533823257",
      "verify_operation_name": null,
      "verify_status": "01"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 2
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_inventory\_record_id |number | ✅ |  盘点id| |
| inventory_date |number | ✅ |  盘点日期| |
| inventory\_operation_name |number | ✅ |  盘点人员| |
| order_number |number | ✅ | 盘点单号| |
| verify\_operation_name |number | ✅ |  确认人员| |
| verify_status |number | ✅ |  确认状态| |


</br>
<h3>11.35 药房盘点记录详情

```
请求地址：/clinic_drug/DrugInventoryRecordDetail
```
**请求包示例**

```
{
   drug_inventory_record_id: 1   
   clinic_id: 1
   keyword: ''
   status: false
   amount: 1
   offset: 0
   limit: 10
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_inventory\_record_id |number | ✅ |  入库记录id| |
| clinic_id |number | ✅ |  入库记录id| |
| keyword | string | ✅ |  关键词| |
| status |boolean | ✅ |  确认状态| |
| amount |boolean | ✅ |  数量| |


**应答包示例**

```
{
  "code": "200",
  "data": {
    "created_time": "2018-08-09T22:00:57.177251+08:00",
    "drug_inventory_record_id": 5,
    "inventory_date": "2018-08-09T00:00:00Z",
    "inventory_operation_id": 1,
    "inventory_operation_name": "超级管理员",
    "items": [
      {
        "actual_amount": 55,
        "buy_price": 5000,
        "drug_stock_id": 20,
        "eff_date": "2018-08-01T00:00:00Z",
        "manu_factory_name": "北京鹤延龄饮片厂",
        "name": "川贝母",
        "packing_unit_name": "g",
        "serial": "20180801004",
        "specification": "H01019/kg",
        "status": true,
        "stock_amount": -90,
        "supplier_name": "广州白云药厂"
      }
    ],
    "order_number": "DPD-1533823257",
    "updated_time": "2018-08-09T22:00:57.177251+08:00",
    "verify_operation_id": null,
    "verify_operation_name": null,
    "verify_status": "01"
  },
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 40,
    "total_item": [
      {
        "actual_amount": 1133,
        "drug_stock_id": 9
      }
    ]
  }
}
```

**应答包参数说明** （同11.34和 11.1）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| page_info.actual_amount |number | ✅ |  合计（库存剩余数量）| |
| page_info. drug_stock_id |number | ✅ |  合计（库存id）| |


</br>
<h3>11.36 修改药房盘点

```
请求地址：/clinic_drug/DrugInventoryUpdate
```
**请求包示例**

```
{
   drug_inventory_record_id: 1 
   clinic_id: 1
   inventory_operation_id: 1
   items: [{
       drug_stock_id: 1
       actual_amount: 1
   }] 
}
```

**请求包参数说明** （同 11.33）


**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.37 药房盘点审核

```
请求地址：/clinic_drug/DrugInventoryCheck
```
**请求包示例**

```
{
   drug_inventory_record_id: 1   
   verify_operation_id: 1
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_inventory\_record_id |number | ✅ |  药品盘点id| |
| verify\_operation_id |number | ✅ |  盘点确认id| |


**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.38 删除盘点记录

```
请求地址：/clinic_drug/DrugInventoryRecordDelete
```
**请求包示例**

```
{
   drug_inventory_record_id: 1   
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_inventory\_record_id |number | ✅ |  药品盘点id| |


**应答包示例**

```
{
  "code": "200",
  ”msg“: "成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |


</br>
<h3>11.39 库存盘点列表

```
请求地址：/clinic_drug/DrugStockInventoryList
```
**请求包示例**

```
{
   clinic_id: 1  
   keyword: ''
   status: false
   amount: 1
   offset: 10
   limit: 10
}
```

**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id| |
| keyword |number | ✅ |  关键词| |
| status |number | ✅ |  审核状态| |
| amount |number | ✅ |  数量| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "buy_price": 5000,
      "day_warning": 10,
      "drug_stock_id": 20,
      "eff_date": "2018-08-01T00:00:00Z",
      "manu_factory_name": "北京鹤延龄饮片厂",
      "name": "川贝母",
      "packing_unit_name": "g",
      "ret_price": 1000,
      "serial": "20180801004",
      "specification": "H01019/kg",
      "status": true,
      "stock_amount": -90,
      "stock_warning": null,
      "supplier_name": "广州白云药厂"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 40
  }
}
```

**应答包参数说明** （同11.1）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code |number | ✅ |  200时 成功| |



12 角色模块
--------


</br>
<h3>12.1 创建角色

```
请求地址：/role/create
```
**请求包示例**

```
{
	name:‘’
	clinic_id: 1
	items:[{
	  clinic_function_menu_id: 1
	}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| name | String | ✅ |  角色名称 | |
| clinic_id | String | ✅ |  诊所id | |
| items.clinic\_function\_menu_id | String | ✅ | 功能id |  |

**应答包示例**

```
{
    "code": "200",
    "msg": ""
    “data”: 1
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
|code | String | ✅ |  200 成功 | |
|msg | String | ✅ |  消息| |
| data | String | ✅ |  角色id| |
--


</br>
<h3>12.2 更新角色

```
请求地址：/role/update
```
**请求包示例**

```
{
	name:‘’
	role_id: 1
	items:[{
	  clinic_function_menu_id: 1
	}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| name | String | ✅ |  角色名称 | |
| role_id | String | ✅ |  角色id | |
| items.clinic\_function\_menu_id | String | ✅ | 功能id |  |

**应答包示例**

```
{
    "code": "200",
    "msg": ""
    “data”: 1
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
|code | String | ✅ |  200 成功 | |
|msg | String | ✅ |  消息| |
| data | String | ✅ |  角色id| |
--

</br>
<h3>12.3 更新角色

```
请求地址：/role/update
```
**请求包示例**

```
{
	role_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| role_id | String | ✅ |  角色id | |

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
|code | String | ✅ |  200 成功 | |
|msg | String | ✅ |  消息| |
--

</br>
<h3>12.4 列表

```
请求地址：/role/listByClinicID
```
**请求包示例**

```
{
	clinic_id: 1
	keyword: ''
	offset: 1
	limit: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ❌ |  诊所id | |
| keyword | String | ❌ |  关键字 | |
| offset | String | ❌ |  跳过数 | 0 |
| limit | String | ❌ |  每页数 | 10 |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-08-01T13:49:03.304447+08:00",
      "function_menu_name": "就诊流程,药品零售,门诊发药",
      "name": "发药",
      "role_id": 47,
      "status": true
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 12
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| function\_menu\_name | String | ✅ |  功能名称 | |
| status | String | ✅ |  启用状态| |
| name | String | ✅ |  角色名称| |
| role_id | String | ✅ |  角色id| |
| created_time | String | ✅ |  创建时间| |
--


</br>
<h3>12.5 角色详情

```
请求地址：/role/roleDetail
```
**请求包示例**

```
{
	role_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| role_id | String | ✅ |  诊所id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "name": "发药",
      "role_id": 47,
      "status": true,
      "funtionMenus": [
        {
        "ascription": "01",
        "clinic_function_menu_id": 1,
        "function_menu_id": 1,
        "icon": null,
        "level": 0,
        "menu_name": "就诊流程",
        "menu_url": "/treatment",
        "parent_function_menu_id": null,
        "status": true,
        "weight": 0
        }
      ],
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 12
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| status | String | ✅ |  启用状态| |
| name | String | ✅ |  角色名称| |
| role_id | String | ✅ |  角色id| |
| funtionMenus.ascription | String | ✅ |  01 诊所 02 平台| |
| funtionMenus.clinic\_function\_menu_id | String | ✅ | 诊所功能id| |
| funtionMenus.function\_menu_id | String | ✅ | 功能id| |
| funtionMenus.icon | String | ✅ | 图标| |
| funtionMenus.level | String | ✅ | 等级| |
| funtionMenus.menu_name | String | ✅ | 功能名称| |
| funtionMenus.menu_url | String | ✅ | 功能地址| |
| funtionMenus.parent\_function\_menu_id | String | ✅ | 父级id| |
| funtionMenus.weight | String | ✅ | 权重| |
| funtionMenus.status | String | ✅ | 启用状态| |
--


</br>
<h3>12.6 获取角色未开通的菜单项

```
请求地址：/role/RoleFunctionUnset
```
**请求包示例**

```
{
	role_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| role_id | String | ❌ |  角色id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "function_menu_id": 4,,
      "menu_name": "设置管理",
      "level": 0
      "status": true,
      "funtionMenus": [
        {
        "ascription": "01",
        "clinic_function_menu_id": 1,
        "function_menu_id": 1,
        "icon": null,
        "level": 0,
        "menu_name": "就诊流程",
        "menu_url": "/treatment",
        "parent_function_menu_id": null,
        "status": true,
        "weight": 0
        }
      ],
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 12
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| status | String | ✅ |  启用状态| |
| name | String | ✅ |  角色名称| |
| role_id | String | ✅ |  角色id| |
| funtionMenus.ascription | String | ✅ |  01 诊所 02 平台| |
| funtionMenus.clinic\_function\_menu_id | String | ✅ | 诊所功能id| |
| funtionMenus.function\_menu_id | String | ✅ | 功能id| |
| funtionMenus.icon | String | ✅ | 图标| |
| funtionMenus.level | String | ✅ | 等级| |
| funtionMenus.menu_name | String | ✅ | 功能名称| |
| funtionMenus.menu_url | String | ✅ | 功能地址| |
| funtionMenus.parent\_function\_menu_id | String | ✅ | 父级id| |
| funtionMenus.weight | String | ✅ | 权重| |
| funtionMenus.status | String | ✅ | 启用状态| |
--


</br>
<h3>12.7 获取角色未开通的菜单项

```
请求地址：/role/roleDetail
```
**请求包示例**

```
{
	role_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| role_id | String | ✅ |  角色id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "function_menu_id": 4,
      "menu_name": "设置管理",
      "level": 0,
      "ascription": "01",
      "status": true,
      "weight": 3,
      "menu_url": "/setting",
      "list": [
        {
           "function_menu_id": 26,
          "parent_function_menu_id": 4,
          "menu_name": "模板设置",
          "level": 1,
          "ascription": "01",
          "icon": "/static/icons/template.svg",
          "status": true,
          "weight": 1,
          "menu_url": "/setting/template",
          list: [ 
           {
              "function_menu_id": 41,
              "parent_function_menu_id": 26,
              "clinic_function_menu_id": 40,
              "menu_name": "检验模板",
              "level": 2,
              "ascription": "01",
              "status": true,
              "weight": 1,
              "menu_url": "/setting/template/inspectionTemplate"
           }
          ]
        }
      ],
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 12
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| status | String | ✅ |  启用状态| |
| name | String | ✅ |  角色名称| |
| role_id | String | ✅ |  角色id| |
| funtionMenus.ascription | String | ✅ |  01 诊所 02 平台| |
| clinic\_function\_menu_id | String | ✅ | 诊所功能id| |
| function\_menu_id | String | ✅ | 功能id| |
| icon | String | ✅ | 图标| |
| level | String | ✅ | 等级| |
| menu_name | String | ✅ | 功能名称| |
| menu_url | String | ✅ | 功能地址| |
| parent\_function\_menu_id | String | ✅ | 父级id| |
| weight | String | ✅ | 权重| |
--


</br>
<h3>12.8 在角色下分配用户

```
请求地址：/role/RoleAllocation
```
**请求包示例**

```
{
	role_id: 1
	items: [
	 {
	    personnel_id : 1
	 }
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| role_id | String | ✅ |  角色id | |
| personnel_id | String | ✅ |  人id | |

**应答包示例**

```
{
  "code": "200",
  "msg": 
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  200 成功| |
| msg | String | ✅ | 消息| |
--


</br>
<h3>12.9 角色分配的用户列表

```
请求地址：/role/PersonnelsByRole
```
**请求包示例**

```
{
	role_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| role_id | String | ✅ |  角色id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "department_name": "牙科",
      "personnel_id": 33,
      "personnel_name": "王思聪"
    }
  ],
  "msg": "ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| department_name | String | ✅ | 诊所名称| |
| personnel_id | String | ✅ | 人id| |
| personnel_name | String | ✅ | 姓名| |
--


13 业务权限模块
--------

</br>
<h3>13.1 添加功能菜单栏

```
请求地址：/business/menubar/create
```

**请求包示例**

```
{
	url: 1
	level: 1
	icon: 1
	name: 1
	weight: 1
   ascription: 1
   parent_function_menu_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| url | String | ✅ |  地址 | |
| level | String | ✅ |  等级 | |
| icon | String | ✅ |  图标 | |
| name | String | ✅ |  名称 | |
| weight | String | ✅ |  权重 | |
| ascription | String | ✅ |  类型 | |
| parent\_function\_menu_id | String | ✅ |  父级id | |

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
| code | String | ✅ | 200 时成功| |
--

</br>
<h3>13.2 添加功能菜单栏

```
请求地址：/business/menubar/list
```

**请求包示例**

```
{
	ascription: 01
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| ascription | String | ✅ |  类型 | |


**应答包示例**

```
{
  "code": "200",
  "msg": "ok"
  “data”: [    {
      "ascription": "01",
      "function_menu_id": 38,
      "icon": null,
      "level": 2,
      "menu_name": "其他费用",
      "menu_url": "/setting/chargeItemSetting/otherFee",
      "parent_function_menu_id": 25,
      "status": true,
      "weight": 7
    }]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| ascription | String | ✅ |  01 诊所 02 平台| |
| function\_menu_id | String | ✅ | 功能id| |
| icon | String | ✅ | 图标| |
| level | String | ✅ | 等级| |
| menu_name | String | ✅ | 功能名称| |
| menu_url | String | ✅ | 功能地址| |
| weight | String | ✅ | 权重| |
--

</br>
<h3>13.3 获取诊所未开通的菜单项

```
请求地址：/business/menubar/list/clinicUnset
```

**请求包示例**

```
{
	clinic_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ |  诊所id | |


**应答包示例**

```
{
  "code": "200",
  "msg": "ok"
  “data”: [    {
      "ascription": "01",
      "function_menu_id": 38,
      "icon": null,
      "level": 2,
      "menu_name": "其他费用",
      "menu_url": "/setting/chargeItemSetting/otherFee",
      "parent_function_menu_id": 25,
      "status": true,
      "weight": 7
    }]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| ascription | String | ✅ |  01 诊所 02 平台| |
| function\_menu_id | String | ✅ | 功能id| |
| icon | String | ✅ | 图标| |
| level | String | ✅ | 等级| |
| menu_name | String | ✅ | 功能名称| |
| menu_url | String | ✅ | 功能地址| |
| weight | String | ✅ | 权重| |
--

</br>
<h3>13.4 诊所分配业务

```
请求地址：/business/clinic/assign
```

**请求包示例**

```
{
	clinic_id: 1
	items: [
	  {function_menu_id: 1}
	]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ |  诊所id | |
| function_menu_id | String | ✅ |  功能id | |


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
|code | String | ✅ |  200 成功| |
--

</br>
<h3>13.5 诊所分配业务

```
请求地址：/business/clinic/menubar
```

**请求包示例**

```
{
	clinic_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ |  诊所id | |


**应答包示例**

```
{
  "code": "200",
  "msg": "ok"
  data: [
    {
      "ascription": "01",
      "clinic_function_menu_id": 46,
      "function_menu_id": 1,
      "icon": null,
      "level": 0,
      "menu_name": "就诊流程",
      "menu_url": "/treatment",
      "parent_function_menu_id": null,
      "status": true,
      "weight": 0
    }
  
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| ascription | String | ✅ |  01 诊所 02 平台| |
| function\_menu_id | String | ✅ | 功能id| |
| icon | String | ✅ | 图标| |
| level | String | ✅ | 等级| |
| menu_name | String | ✅ | 功能名称| |
| menu_url | String | ✅ | 功能地址| |
| weight | String | ✅ | 权重| |
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
| medical_record_model_id | Int | ✅ | 模板id| |

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

17 检查缴费项目模块
--------

</br>
<h3>17.1 创建检查缴费项目

```
请求地址：/examination/create
```
**请求包示例**

```
{
	clinic_id:1
	name:动态心电图(Holter)
	en_name:
	py_code:DTXDTHOL
	idc_code:
	unit_name:项
	organ:
	remark:备注
	price:110.00
	cost:
	status:true
	is_discount:false
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所id| |
| name | String | ✅ | 检查项目名称 | |
| py_code | String | ❌  |  拼音简码 | |
| idc_code | String | ❌  |  国际名称 | |
| en_name | String | ❌  |  英文名称 | |
| unit_name | String | ❌  |  单位 | 项|
| organ | String | ❌  |  检查部位 | |
| remark | String | ❌  | 备注 | |
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
<h3>17.2 修改检查缴费项目

```
请求地址：/examination/update
```
**请求包示例**

```
{
	clinic_id:1
	name:动态心电图(Holter)
	en_name:
	py_code:DTXDTHOL
	idc_code:
	unit_name:项
	organ:
	remark:备注
	price:110.00
	cost:
	status:true
	is_discount:false
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_examination_id | Int | ✅ |  诊所检查项目id| |
| name | String | ✅ | 检查项目名称 | |
| py_code | String | ❌  |  拼音简码 | |
| idc_code | String | ❌  |  国际名称 | |
| en_name | String | ❌  |  英文名称 | |
| unit_name | String | ❌  |  单位 | 项|
| organ | String | ❌  |  检查部位 | |
| remark | String | ❌  | 备注 | |
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
<h3>17.3 启用和停用检查项目

```
请求地址：/examination/onOff
```
**请求包示例**

```
{
	status:true
	clinic_id:1
	clinic_examination_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| status | Boolean | ✅ |  是否启用 | |
| clinic_id | Int | ✅ |  诊所id| |
| clinic_examination_id | Int | ✅ |  诊所检查项目id| |
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
<h3>17.4 诊疗项目列表

```
请求地址：/examination/list
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
| offset | int | ❌ |  开始条数 | 0 |
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_examination_id": 6,
      "cost": null,
      "en_name": null,
      "idc_code": null,
      "is_discount": false,
      "name": "动态心电图(Holter)",
      "organ": null,
      "price": 11000,
      "py_code": "DTXDTHOL",
      "remark": null,
      "status": true,
      "unit_name": "项"
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
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.clinic_examination_id | int | ✅ |  检查项目id | |
| data.cost | Int | ❌ | 成本价 | |
| data.en_name | String | ❌ |  英文名称| |
| data.idc_code | String | ❌ |  国际编码| |
| data.py_code | String | ❌ |  拼音简码| |
| data.unit_name | String | ❌ |  单位|项 |
| data.organ | String | ❌ | 检查部位||
| data.remark | String | ❌ |  备注| |
| data.is_discount | Boolean | ✅ |  是否折扣 | |
| data.name | String | ✅ |  检查名称 | |
| data.price | int | ✅ |  零售价 | |
| data.status | bolean | ✅ |  是否启用 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |

--

</br>
<h3>17.5 检查项目详情

```
请求地址：/examination/detail
```
**请求包示例**

```
{
	clinic_examination_id:2
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
    "clinic_examination_id": 6,
    "cost": null,
    "en_name": null,
    "idc_code": null,
    "is_discount": false,
    "name": "动态心电图(Holter)",
    "organ": null,
    "price": 11000,
    "py_code": "DTXDTHOL",
    "remark": null,
    "status": true,
    "unit_name": "项"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.clinic_examination_id | int | ✅ |  检查项目id | |
| data.cost | Int | ❌ | 成本价 | |
| data.en_name | String | ❌ |  英文名称| |
| data.idc_code | String | ❌ |  国际编码| |
| data.py_code | String | ❌ |  拼音简码| |
| data.unit_name | String | ❌ |  单位|项 |
| data.organ | String | ❌ | 检查部位||
| data.remark | String | ❌ |  备注| |
| data.is_discount | Boolean | ✅ |  是否折扣 | |
| data.name | String | ✅ |  检查名称 | |
| data.price | int | ✅ |  零售价 | |
| data.status | bolean | ✅ |  是否启用 | |
--

</br>
<h3>17.6 创建检查医嘱模板

```
请求地址：/examination/ExaminationPatientModelCreate
```
**请求包示例**

```
{
	model_name:胸全检
	is_common:true
	operation_id:1
	items:[
	{"clinic_examination_id":"1","times":"1","organ":"","illustration":"二丫头如何"},
	{"clinic_examination_id":"2","times":"1","organ":"","illustration":"地方还是风格还是"}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| model_name | String | ✅ | 模板名称| |
| is_common | Boolean | ✅ | 是否通用 | |
| operation_id | Int |  ✅  |  操作人id | |
| items | Array | ✅ |  检查项 | |
| items.clinic_examination_id | String | ✅  |  个人病史 | |
| items. times | String | ✅  |  次数 | |
| items. organ | String | ❌  |  检查部位 | |
| items. illustration | String | ❌  |  说明 | |

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
<h3>17.7 查询检查模板列表

```
请求地址：/examination/ExaminationPatientModelList
```
**请求包示例**

```
{
	keyword:全
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
      "model_name": "胸全检",
      "examination_patient_model_id": 2,
      "operation_name": "超级管理员",
      "is_common": true,
      "created_time": "2018-05-27T22:07:08.877222+08:00",
      "items": [
        {
          "examination_name": "胸部正位",
          "organ": null,
          "times": 1,
          "clinic_examination_id": 1,
          "illustration": "二丫头如何"
        },
        {
          "examination_name": "胸部正位+侧位",
          "organ": null,
          "times": 1,
          "clinic_examination_id": 2,
          "illustration": "地方还是风格还是"
        }
      ]
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
| data.model_name | String | ✅ | 模板名称 | |
| data.examination_patient_model_id | Int | ✅ |  检查模板id| |
| data.operation_name| String | ✅ |  操作人员名称| |
| data.is_common | Boolean | ✅ | 是否通用 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.items | Array | ✅ |  检查项 | |
| data.items.clinic_examination_id | Int | ✅  |  检查项id | |
| data.items.examination_name | String | ✅  |  检查项名称 | |
| data.items.times | Int | ✅  |  次数 | |
| data.items.organ | String | ❌  |  检查部位 | |
| data.items.illustration | String | ❌  |  说明 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>17.8 查询检查医嘱模板详情

```
请求地址：/examination/ExaminationPatientModelDetail
```
**请求包示例**

```
{
	examination_patient_model_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| examination_patient_model_id | Int | ✅ |  检查模板id | |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "examination_patient_model_id": 2,
    "is_common": true,
    "items": [
      {
        "clinic_examination_id": 1,
        "illustration": "二丫头如何",
        "name": "胸部正位",
        "organ": null,
        "times": 1
      },
      {
        "clinic_examination_id": 2,
        "illustration": "地方还是风格还是",
        "name": "胸部正位+侧位",
        "organ": null,
        "times": 1
      }
    ],
    "model_name": "胸全检",
    "status": true
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Object | ✅ |   | |
| data.model_name | String | ✅ | 主诉 | |
| data.examination_patient_model_id | Int | ✅ |  检查模板id| |
| data.is_common | Boolean | ✅ | 是否通用 | |
| data.status | Boolean | ✅ | 是否启用 | |
| data.items | Array | ✅ |  检查项 | |
| data.items.clinic_examination_id | Int | ✅  |  检查项id | |
| data.items.name | String | ✅  |  检查项名称 | |
| data.items.times | Int | ✅  |  次数 | |
| data.items.organ | String | ❌  |  检查部位 | |
| data.items.illustration | String | ❌  |  说明 | |
--

</br>
<h3>17.9 修改检查医嘱模板

```
请求地址：/examination/ExaminationPatientModelUpdate
```
**请求包示例**

```
{
	examination_patient_model_id:2
	model_name:胸全检
	is_common:true
	operation_id:1
	items:[
	{"clinic_examination_id":"1","times":"1","organ":"","illustration":"二丫头如何"},
	{"clinic_examination_id":"2","times":"1","organ":"","illustration":"地方还是风格还是"}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| examination_patient_model_id | Int | ✅ | 模板id| |
| model_name | String | ✅ | 模板名称| |
| is_common | Boolean | ✅ | 是否通用 | |
| operation_id | String |  ✅  |  操作人id | |
| items | Array | ✅ |  检查项 | |
| items.clinic_examination_id | String | ✅  |  个人病史 | |
| items. times | String | ✅  |  次数 | |
| items. organ | String | ❌  |  检查部位 | |
| items. illustration | String | ❌  |  说明 | |

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
<h3>17.10 删除检查医嘱模板

```
请求地址：/examination/ExaminationPatientModelDelete
```
**请求包示例**

```
{
	examination_patient_model_id:2
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| medical_record_model_id | Int | ✅ | 模板id| |

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
<h3>17.11 创建检查报告医嘱模板

```
请求地址：/examination/ExaminationReportModelCreate
```
**请求包示例**

```
{
	model_name:检查报告模板
	result_examination:检查报告模板描述
	conclusion_examination:检查报告模板结论
	operation_id:17
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| model_name | String | ✅ | 模板名称| |
| result_examination | String | ❌ | 检查结果/描述 | |
| conclusion_examination | String |  ❌  |  检查结论 | |
| operation_id | Int | ✅ |  操作员id | |

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
<h3>17.12 查询检查报告医嘱模板

```
请求地址：/examination/ExaminationReportModelList
```
**请求包示例**

```
{
	keyword:模板
	operation_id:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| operation_id | int | ❌ |  创建人员id | |
| offset | int | ❌ |  开始条数 | |
| limit | int | ❌ |  条数 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "conclusion_examination": "胸部平扫并与2011-8-17本院片比较，左肺上叶斑片样密度增高影范围缩小;其它征象无明显改变。",
      "created_time": "2018-08-01T10:25:50.41798+08:00",
      "deleted_time": null,
      "id": 3,
      "model_name": "8月1日的报告测试模板",
      "operation_id": 1,
      "operation_name": "超级管理员",
      "result_examination": "左肺上叶肺炎消散期改变",
      "status": true,
      "updated_time": "2018-08-01T10:25:50.41798+08:00"
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
| data.model_name | String | ✅ | 模板名称 | |
| data.id | Int | ✅ |  检查报告模板id| |
| data.operation_id| Int | ✅ |  操作人员id| |
| data.operation_name| String | ✅ |  操作人员名称| |
| data.conclusion_examination | String | ✅ | 检查结论 | |
| data.result_examination | String | ✅ | 检查结果/描述 | |
| data.status | Boolean | ✅ | 是否启用 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>17.13 查询检查报告医嘱模板详情

```
请求地址：/examination/ExaminationReportModelDetail
```
**请求包示例**

```
{
	examination_report_model_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| examination_report_model_id | Int | ✅ |  检查模板id | |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "conclusion_examination": "检查报告模板结论修改",
    "created_time": "2018-07-30T22:16:14.546186+08:00",
    "deleted_time": "2018-07-30T22:20:56.813356+08:00",
    "id": 1,
    "model_name": "检查报告模板名称修改",
    "operation_id": 17,
    "result_examination": "检查报告模板描述修改修改",
    "status": true,
    "updated_time": "2018-07-30T22:20:25.951168+08:00"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Object | ✅ |   | |
| data.model_name | String | ✅ | 模板名称 | |
| data.id | Int | ✅ |  检查报告模板id| |
| data.operation_id| Int | ✅ |  操作人员id| |
| data.conclusion_examination | String | ✅ | 检查结论 | |
| data.result_examination | String | ✅ | 检查结果/描述 | |
| data.status | Boolean | ✅ | 是否启用 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
--

</br>
<h3>17.14 修改检查报告医嘱模板

```
请求地址：/examination/ExaminationReportModelUpdate
```
**请求包示例**

```
{
	examination_report_model_id:2
	model_name:检查报告模板
	result_examination:检查报告模板描述
	conclusion_examination:检查报告模板结论
	operation_id:17
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| examination_report_model_id | Int | ✅ | 模板id| |
| model_name | String | ✅ | 模板名称| |
| result_examination | String | ❌ | 检查结果/描述 | |
| conclusion_examination | String |  ❌  |  检查结论 | |
| operation_id | Int | ✅ |  操作员id | |

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
<h3>17.15 删除检查报告医嘱模板

```
请求地址：/examination/ExaminationReportModelDelete
```
**请求包示例**

```
{
	examination_report_model_id:2
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| examination_report_model_id | Int | ✅ | 模板id| |

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
<h3>20.1 创建材料费用项目

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

21 检验模块
--------

</br>
<h3>21.1 检验医嘱创建

```
请求地址：/laboratory/create
```
**请求包示例**

```
{
	clinic_id:1
	name: ''
	en_name: ''
	py_code: ''
	idc_code: ''
	unit_name: ''
	time_report: ''
	clinical_significance: ''
	remark: ''
	laboratory_sample: ''
	cuvette_color_name: ''
	merge_flag: ''
	cost: 1
	price 1
	status : true
	is_discount: flase
	is_delivery: false
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| name | String | ✅ |  医嘱名称 | |
| en_name | String | ❌ |  英文名称 | |
| py_code | String | ❌ |  拼音编码 | |
| idc_code | String | ❌ |  国家准字 | |
| unit_name | String | ❌ |  单位 | |
| time_report | String | ❌ |  报告所需时间 | |
| clinical_significance | String | ❌ |  临床意义 | |
| remark | String | ❌ |  备注 | |
| laboratory_sample | String | ❌ |  检验物 | |
| cuvette\_color_name | String | ❌ |  试管颜色 | |
| merge_flag |number | ❌ |  合并标识 | |
| cost | String | ❌ |  成本 | |
| price | String | ✅ |  单价 | |
| status | boolean | ✅ |  启用状态 | true |
| is_discount |boolean | ✅ |  是否允许折扣 | false |
| is_delivery | boolean | ✅ |  否允许外送| false |

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
| code | String | ✅ | 200 成功 | |

</br>
<h3>21.2 检验医嘱列表

```
请求地址：/laboratory/list
```
**请求包示例**

```
{
	keyword:‘’
	clinic_id： 1
	status： true
	offset: 0
	limit : 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ |  诊所id | |
| keyword | String | ❌ |  关键字 | |
| status | String | ❌ |  启用状态 | |
**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_laboratory_id": 4,
      "discount_price": 0,
      "is_discount": false,
      "laboratory_name": "性激素六项",
      "price": 26000,
      "py_code": "xjslx",
      "remark": null,
      "status": true,
      "unit_name": "项"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 10
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory_id | String | ✅ | 检验id | |
| discount_price | String | ✅ | 折扣金额| |
| is_discount | String | ✅ | 是否折扣 | |
| laboratory_name | String | ✅ | 检验名称 | |
| price | String | ✅ | 单价 | |
| py_code | String | ✅ | 拼音编码 | |
| remark | String | ✅ | 备注 | |
| status | String | ✅ | 启用状态 | |
| unit_name | String | ✅ | 单位 | |

</br>
<h3>21.3 检验医嘱详情

```
请求地址：/laboratory/detail
```
**请求包示例**

```
{
	clinic_laboratory_id:4
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory_id | String | ✅ |  检验医嘱id | |
**应答包示例**

```
{
  "code": "200",
  "data": {
    "clinic_laboratory_id": 4,
    "clinical_significance": null,
    "cost": 10000,
    "cuvette_color_name": "红",
    "discount_price": 0,
    "en_name": null,
    "idc_code": null,
    "is_delivery": false,
    "is_discount": false,
    "laboratory_sample": "血清",
    "merge_flag": null,
    "name": "性激素六项",
    "price": 26000,
    "py_code": "xjslx",
    "remark": null,
    "status": true,
    "time_report": null,
    "unit_name": "项"
  }
}
```

**应答包参数说明** (同 12.1)


</br>
<h3>21.4 检验医嘱更新

```
请求地址：/laboratory/update
```
**请求包示例**

```
{

	clinic_laboratory_id:1
	name: ''
	en_name: ''
	py_code: ''
	idc_code: ''
	unit_name: ''
	time_report: ''
	clinical_significance: ''
	remark: ''
	laboratory_sample: ''
	cuvette_color_name: ''
	merge_flag: ''
	cost: 1
	price 1
	status : true
	is_discount: flase
	is_delivery: false
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory_id |number | ✅ |  检验id | |
| name | String | ✅ |  医嘱名称 | |
| en_name | String | ❌ |  英文名称 | |
| py_code | String | ❌ |  拼音编码 | |
| idc_code | String | ❌ |  国家准字 | |
| unit_name | String | ❌ |  单位 | |
| time_report | String | ❌ |  报告所需时间 | |
| clinical_significance | String | ❌ |  临床意义 | |
| remark | String | ❌ |  备注 | |
| laboratory_sample | String | ❌ |  检验物 | |
| cuvette\_color_name | String | ❌ |  试管颜色 | |
| merge_flag |number | ❌ |  合并标识 | |
| cost | String | ❌ |  成本 | |
| price | String | ✅ |  单价 | |
| status | boolean | ✅ |  启用状态 | true |
| is_discount |boolean | ✅ |  是否允许折扣 | false |
| is_delivery | boolean | ✅ |  否允许外送| false |

**应答包示例**

```
{
    "code": "200",
    "msg": ""
    data: 1
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功 | |
| data | String | ✅ | 检验医嘱id | |

</br>
<h3>21.5 检验医嘱启用

```
请求地址：/laboratory/onOff
```
**请求包示例**

```
{

	clinic_laboratory_id:1
	status: true
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory_id |number | ✅ |  检验id | |
| status | boolean | ✅ |  启用状态 | |


**应答包示例**

```
{
    "code": "200",
    "msg": ""
    “data”: 1
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功 | |
| data | String | ✅ | 检验医嘱id | |

</br>
<h3>21.6 关联检验项目

```
请求地址：/laboratory/association
```
**请求包示例**

```
{

	clinic_laboratory_id:1
	item:[{
	    clinic_laboratory_item_id： 1
	    name: ''
	    default_result: 1
	}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory_id |number | ✅ |  检验医嘱id | |
| clinic\_laboratory\_item_id | boolean | ✅ |  检验项目id | |
| name |string | ✅ |   检验名称| |
| default_result |string | ✅ |   默认结果| |

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
| code | String | ✅ | 200 成功 | |


</br>
<h3>21.7 检验医嘱关联项目列表

```
请求地址：/laboratory/associationList
```
**请求包示例**

```
{

	clinic_laboratory_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory_id |number | ✅ |  检验医嘱id | |
**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_laboratory_item_id": 2,
      "name": "血小板",
      "en_name": "xuexiaoban",
      "unit_name": "/L",
      "status": true,
      "is_special": false,
      "data_type": 2,
      "instrument_code": null,
      "is_delivery": null,
      "result_inspection": null,
      "default_result": "100",
      "clinical_significance": null,
      "references": [
        {
          "reference_sex": "通用",
          "reference_max": "20",
          "reference_min": "10",
          "reference_value": null,
          "isPregnancy": false,
          "stomach_status": "false"
        }
      ]
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory\_item_id | String | ✅ | 检验子项id | |
| name | String | ✅ | 检验名称 | |
| en_name | String | ✅ | 英文名称 | |
| unit_name | String | ✅ | 单位 | |
| status | String | ✅ | 启用状态 | |
| is_special |boolean | ✅ | 参考值是否特殊| |
| data_type | String | ✅ | 数据类型 1 定性 2 定量 | |
| instrument_code | String | ✅ | 仪器编码 | |
| is_delivery | String | ✅ | 是否允许外送 | |
| result_inspection | String | ✅ | 检验结果 | |
| default_result | String | ✅ | 默认结果 | |
| clinical_significance | String | ✅ | 临床意义 | |
| reference_sex | String | ✅ | 参考值性别 男、女、通用 | |
| reference_max | String | ✅ | 参考最大值 | |
| reference_min | String | ✅ | 参考最小值 | |
| reference_value | String | ✅ | 定性参考值 | |
| isPregnancy |boolean | ✅ | 是否妊娠期 | |
| stomach_status |boolean | ✅ | 是否空腹 | |


</br>
<h3>21.8 检验项目创建

```
请求地址：/laboratory/item/create
```
**请求包示例**

```
{

	clinic_id:1
	name: ''
	en_name: ''
	instrument_code: 1
	unit_name: 'w'
	clinical_significance: 1
	data_type: 1
	is_special: 12
	reference_max: 11
	reference_min: 12
	status: true
	is_delivery: 1
	items ; [{
	   reference_sex: 1
	   age_max: '10'
	   age_min: '20'
	   reference_max: 12
	   reference_min: 11
	   stomach_status: false
	   is_pregnancy: 1
	}]
}
```
**请求包参数说明** (同上)

**应答包示例**

```
{
  "code": "200",
  "msg": ''
  "data" : 1
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 检验项目id | |


</br>
<h3>21.9 检验项目创建

```
请求地址：/laboratory/item/detail
```
**请求包示例**

```
{

	clinic_laboratory_item_id: 2
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory\_item_id | String | ✅ | 检验项目id | |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "clinic_laboratory_item_id": 2,
    "name": "血小板",
    "en_name": "xuexiaoban",
    "unit_name": "/L",
    "status": true,
    "is_special": false,
    "data_type": 2,
    "instrument_code": null,
    "is_delivery": false,
    "result_inspection": null,
    "default_result": null,
    "clinical_significance": "符合规范化是否更换",
    "references": [
      {
        "reference_sex": "通用",
        "reference_max": "20",
        "reference_min": "10",
        "reference_value": null,
        "isPregnancy": false,
        "stomach_status": "false"
      }
    ]
  }
}
```

**应答包参数说明** （同21.7）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 检验项目id | |


</br>
<h3>21.10 检验项目更新

```
请求地址：/laboratory/item/update
```
**请求包示例**

```
{

	clinic_laboratory_item_id:1
	name: ''
	en_name: ''
	instrument_code: 1
	unit_name: 'w'
	clinical_significance: 1
	data_type: 1
	is_special: 12
	reference_max: 11
	reference_min: 12
	status: true
	is_delivery: 1
	items ; [{
	   reference_sex: 1
	   age_max: '10'
	   age_min: '20'
	   reference_max: 12
	   reference_min: 11
	   stomach_status: false
	   is_pregnancy: 1
	}]
}
```
**请求包参数说明** (同21.8)

**应答包示例**

```
{
  "code": "200",
  "msg": ''
  "data" : 1
}
```
**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 检验项目id | |


</br>
<h3>21.11 检验项目启用

```
请求地址：/laboratory/item/onOff
```
**请求包示例**

```
{

	clinic_id:1
	clinic_laboratory_item_id： 1
	status: false
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ | 诊所id | |
| clinic\_laboratory\_item_id | String | ✅ | 诊所id | |
| status | String | ✅ | 开启状态 | |

**应答包示例**

```
{
  "code": "200",
  "msg": ''
}
```
**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功| |


</br>
<h3>21.12 检验项目列表

```
请求地址：/laboratory/item/list
```
**请求包示例**

```
{

	clinic_id:1
	name： ''
	status: false
	offset: 0
	limit: 10
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | String | ✅ | 诊所id | |
| name | String | ❌ | 项目名称 | |
| status | String | ❌ | 开启状态 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_laboratory_item_id": 5,
      "name": "白细胞计数",
      "en_name": "WBC",
      "unit_name": "个/L",
      "status": true,
      "is_special": true,
      "data_type": 2,
      "instrument_code": null,
      "is_delivery": false,
      "result_inspection": null,
      "default_result": null,
      "clinical_significance": null,
      "references": [
        {
          "reference_sex": "通用",
          "reference_max": "15",
          "reference_min": "10",
          "reference_value": null,
          "isPregnancy": null,
          "stomach_status": null
        }
      ]
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 4
  }
}
```
**应答包参数说明** （同 21.8）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功| |


</br>
<h3>21.13 创建检验医嘱模板

```
请求地址：/laboratory/LaboratoryPatientModelCreate
```
**请求包示例**

```
{

	model_name:1
	is_common： true
	operation_id: 1
	items : [{
	   clinic_laboratory_id: 1
	   times: 1
	   illustration : ''
	}]
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| model_name | String | ✅ | 模板名称 | |
| is_common | boolean | ❌ | 是否通用 | |
| operation_id | String | ✅ | 操作员id | |
| clinic\_laboratory_id | String | ✅ | 检验医嘱id | |
| times | String | ✅ | 数量 | |
| illustration | String | ✅ | 描述 | |

**应答包示例**

```
{
  "code": "200",
  "msg": “操作成功”
}
```
**应答包参数说明** （同 21.8）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功| |


</br>
<h3>21.14 检验医嘱模板列表

```
请求地址：/laboratory/LaboratoryPatientModelList
```
**请求包示例**

```
{

	keyword:1
	is_common： true
	operation_id: 1
	offset: 0
	limit: 10
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| is_common | boolean | ❌ | 是否通用 | |
| operation_id | String | ✅ | 操作员id | |
| keyword | String | ✅ | 搜索关键词 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "model_name": "检验模板第一个",
      "laboratory_patient_model_id": 2,
      "operation_name": "华佗",
      "is_common": true,
      "created_time": "2018-05-29T16:03:47.095669+08:00",
      "items": [
        {
          "laboratory_name": "血常规",
          "times": 2,
          "clinic_laboratory_id": 1,
          "illustration": "说明222"
        }
      ]
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 2
  }
}
```
**应答包参数说明** （同 21.13）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功| |


</br>
<h3>21.15 个人检验医嘱模板

```
请求地址：/laboratory/LaboratoryPersonalPatientModelList
```
**请求包示例**

```
{

	keyword:1
	is_common： true
	operation_id: 1
	offset: 0
	limit: 10
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| is_common | boolean | ❌ | 是否通用 | |
| operation_id | String | ✅ | 操作员id | |
| keyword | String | ✅ | 搜索关键词 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "model_name": "检验模板第一个",
      "laboratory_patient_model_id": 2,
      "operation_name": "华佗",
      "is_common": true,
      "created_time": "2018-05-29T16:03:47.095669+08:00",
      "items": [
        {
          "laboratory_name": "血常规",
          "times": 2,
          "clinic_laboratory_id": 1,
          "illustration": "说明222"
        }
      ]
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 2
  }
}
```
**应答包参数说明** （同 21.13）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功| |


</br>
<h3>21.16 检验医嘱模板详情

```
请求地址：/laboratory/LaboratoryPatientModelDetail
```
**请求包示例**

```
{

	laboratory_patient_model_id:1
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| laboratory\_patient\_model_id | String | ✅ | 检验模板id | |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "is_common": true,
    "items": [
      {
        "clinic_laboratory_id": 2,
        "illustration": "啥的噶",
        "name": "尿常规",
        "times": 1
      },
      {
        "clinic_laboratory_id": 1,
        "illustration": "啊地方噶哒哈",
        "name": "血常规",
        "times": 1
      }
    ],
    "laboratory_patient_model_id": 1,
    "model_name": "血尿检",
    "status": true
  }
}
```
**应答包参数说明** （同 21.13）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| is_common | String | ✅ | 是否通用| |
| laboratory\_patient\_model_id | String | ✅ | 模板id| |
| model_name | String | ✅ | 模板名称| |
| status | String | ✅ | 启用状态| |
| clinic\_laboratory_id | String | ✅ | 检验医嘱id| |
| illustration | String | ✅ | 说明| |
| name | String | ✅ | 医嘱名称| |
| times | String | ✅ | 次数 | |

</br>
<h3>21.17 修改检验医嘱模板

```
请求地址：/laboratory/LaboratoryPatientModelUpdate
```
**请求包示例**

```
{
	laboratory_patient_model_id:1
	model_name: ''
	is_common: false
	operation_id: 1
	items: [{
	   clinic_laboratory_id: 1
	   times: 1
	   illustration: '描述'
	}]
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| laboratory\_patient\_model_id | String | ✅ | 检验模板id | |
| model_name | String | ✅ | 模板名称 | |
| is_common |boolean | ✅ | 是否通用 | |
| operation_id |boolean | ✅ | 创建人id | |
| clinic\_laboratory_id |boolean | ✅ | 检验医嘱id | |
| times | number | ✅ | 数量 | |
| illustration | number | ✅ | 描述 | |

**应答包示例**

```
{
  "code": "200",
  "msg": ’ok‘
}
```
**应答包参数说明** （同 21.13）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功| |

</br>
<h3>21.18 删除检验医嘱模板

```
请求地址：/laboratory/LaboratoryPatientModelDelete
```
**请求包示例**

```
{
	laboratory_patient_model_id:1
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| laboratory\_patient\_model_id | String | ✅ | 检验模板id | |

**应答包示例**

```
{
  "code": "200",
  "msg": ’ok‘
}
```
**应答包参数说明** （同 21.13）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功| |
--

22 基础字典表模块
--------

</br>
<h3>22.1 查询单位列表

```
请求地址：/dictionaries/ExaminationReportModelList
```
**请求包示例**

```
{
	keyword:盒
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "change_flag": 0,
      "code": "08",
      "created_time": "2018-05-27T00:08:45.40778+08:00",
      "d_code": "XH",
      "deleted_flag": 0,
      "deleted_time": null,
      "id": 1,
      "name": "小盒",
      "py_code": "XH",
      "updated_time": "2018-05-27T00:08:45.40778+08:00"
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
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | 单位id | |
| data.name | String | ✅ | 单位名称 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.change_flag | Int | ❌ | 修改标志 | |
| data.code | String | ✅ |  编码| |
| data.d_code| String | ✅ |  简码| |
| data.deleted_flag| Int | ❌ |  删除标志| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.2 药品剂型列表

```
请求地址：/dictionaries/ExaminationReportModelList
```
**请求包示例**

```
{
	keyword:凝
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "code": "GEL",
      "created_time": "2018-05-27T00:08:54.588596+08:00",
      "d_code": "15",
      "deleted_flag": 0,
      "deleted_time": null,
      "id": 1,
      "name": "凝胶剂",
      "py_code": "NJJ",
      "updated_time": "2018-05-27T00:08:54.588596+08:00"
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
| data.id | Int | ✅ | 单位id | |
| data.name | String | ✅ | 剂型名称 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.code | String | ✅ |  编码| |
| data.d_code| String | ✅ |  简码| |
| data.deleted_flag| Int | ❌ |  删除标志| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.3 药物类型列表

```
请求地址：/dictionaries/DrugClassList
```
**请求包示例**

```
{
	clinic_id:1
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所id | |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "count": 33,
      "id": 2,
      "name": "未分类"
    },
    {
      "count": 1,
      "id": 3,
      "name": "抗感染类药物"
    },
    {
      "count": 1,
      "id": 4,
      "name": "呼吸系统用药"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 33
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 类型名称 | |
| data.count | Int | ✅ | 该分类下的总药品数 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.4 药物种类列表

```
请求地址：/dictionaries/DrugTypeList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "code": "GEL",
      "created_time": "2018-05-27T00:08:54.588596+08:00",
      "d_code": "15",
      "deleted_flag": 0,
      "deleted_time": null,
      "id": 1,
      "name": "123",
      "py_code": "123",
      "updated_time": "2018-05-27T00:08:54.588596+08:00"
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
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.code | String | ✅ |  编码| |
| data.d_code| String | ✅ |  简码| |
| data.deleted_flag| Int | ❌ |  删除标志| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.5 药品别名列表

```
请求地址：/dictionaries/DrugPrintList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "code": "GEL",
      "created_time": "2018-05-27T00:08:54.588596+08:00",
      "d_code": "15",
      "deleted_flag": 0,
      "deleted_time": null,
      "id": 1,
      "name": "123",
      "py_code": "123",
      "print_name":"12323",
      "name_type":,
      "updated_time": "2018-05-27T00:08:54.588596+08:00"
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
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 药品名称 | |
| data.print_name | String | ❌ | 药品别名 | |
| data.name_type | String | ✅ | 类型 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.code | String | ✅ |  编码| |
| data.d_code| String | ✅ |  简码| |
| data.deleted_flag| Int | ❌ |  删除标志| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.6 检查部位列表

```
请求地址：/dictionaries/ExaminationOrganList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-05-28T23:49:04.014708+08:00",
      "deleted_time": null,
      "id": 25,
      "name": "肝",
      "updated_time": "2018-05-28T23:49:04.014708+08:00"
    },
    {
      "created_time": "2018-05-28T23:49:04.014708+08:00",
      "deleted_time": null,
      "id": 26,
      "name": "胆",
      "updated_time": "2018-05-28T23:49:04.014708+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 22
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 部位名称 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.7 使用频率列表

```
请求地址：/dictionaries/FrequencyList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "code": "QD12      ",
      "created_time": "2018-05-27T00:08:36.533288+08:00",
      "days": 1,
      "define_code": "QD12    ",
      "delete_flag": 0,
      "deleted_time": null,
      "doctor_flag": null,
      "id": 1,
      "in_out_flag": 2,
      "input_frequency": "QD",
      "medical_bill_code": 1,
      "name": "1次/日 (12am)",
      "print_code": "Qd12",
      "py_code": "QD12    ",
      "times": 1,
      "update_flag": 0,
      "updated_time": "2018-05-27T00:08:36.533288+08:00",
      "week_day_flag": 0,
      "weight": 33
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 62
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 频率名称 | |
| data.define_code | String | ❌ | 自定义码 | |
| data.print_code | String | ❌| 打印名称 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.code | String | ✅ |  编码| |
| data.input_frequency| String | ❌ |  护士录入频率| |
| data.week_day_flag| Int | ❌ |  周日标志| |
| data.update_flag| Int | ❌ |  允许修改标志| |
| data.deleted_flag| Int | ❌ |  删除标志| |
| data.weight| Int | ❌ |  排序码/权重| |
| data.in_out_flag| Int | ❌ |  门诊住院标记| |
| data.medical_bill_code| Int | ❌ |  医保账单码| |
| data.doctor_flag| Int | ❌ |  医生标记| |
| data.times| Int | ❌ |  次数| |
| data.days| Int | ❌ |  天数| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.8 用药途径列表

```
请求地址：/dictionaries/RouteAdministrationList
```
**请求包示例**

```
{
	keyword:下
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "code": "011",
      "created_time": "2018-05-27T00:09:09.514456+08:00",
      "d_code": "HHIT",
      "deleted_flag": 0,
      "deleted_time": null,
      "id": 9,
      "input_type": "8",
      "is_print": 1,
      "name": "皮下注射",
      "print_name": "皮下注射",
      "py_code": "PXZS",
      "type_code": "8",
      "updated_time": "2018-05-27T00:09:09.514456+08:00",
      "weight": 11
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 34
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.print_name | String | ❌ | 药品别名 | |
| data.input_type | String | ✅ | 护士录入类别| |
| data.py_code | String | ✅ | 拼音简码 | |
| data.code | String | ✅ |  编码| |
| data.d_code| String | ✅ |  简码| |
| data.is_print | Int | ❌ |  是否打印| |
| data.print_name | String | ❌ |  打印名称| |
| data.type_code | String | ❌ |  分类编码| |
| data.weight | Int | ❌ |  排序码/权重| |
| data.deleted_flag| Int | ❌ |  删除标志| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.9 标本种类列表

```
请求地址：/dictionaries/LaboratorySampleList
```
**请求包示例**

```
{
	keyword:清
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "code": "A007  ",
      "created_time": "2018-05-27T00:09:18.218938+08:00",
      "deleted_time": null,
      "id": 1,
      "name": "血清",
      "py_code": "XQ",
      "status": true,
      "updated_time": "2018-05-27T00:09:18.218938+08:00"
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
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.code | String | ✅ |  编码| |
| data.status| Boolean | ✅ |  是否启用| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.10 试管颜色列表

```
请求地址：/dictionaries/CuvetteColorList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-05-27T00:07:13.652046+08:00",
      "deleted_time": null,
      "id": 1,
      "name": "红",
      "status": true,
      "updated_time": "2018-05-27T00:07:13.652046+08:00"
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
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.status| Boolean | ✅ |  是否启用| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.11 生产厂商列表

```
请求地址：/dictionaries/ManuFactoryList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "abbr_name": "北京首儿药厂",
      "address": null,
      "code": "101",
      "comment": null,
      "created_time": "2018-05-27T00:11:39.527606+08:00",
      "d_code": "BJSEYC",
      "deleted_flag": 0,
      "deleted_time": null,
      "id": 1,
      "name": "北京首儿药厂",
      "product_range": null,
      "py_code": "BJSEYC",
      "tel": null,
      "updated_time": "2018-05-27T00:11:39.527606+08:00",
      "zip_code": null
    },
	...
  ],
  "page_info": {
    "total": 5292,
    "limit": "10",
    "offset": "0"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.code | String | ✅ | 编码 | |
| data.d_code | String | ❌ | 简码 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.abbr_name | String | ❌ |  | |
| data.address | String | ❌ | 地址 | |
| data.comment | String | ❌ |  | |
| data.product_range | String | ❌ |  | |
| data.tel | String | ❌ | 电话 | |
| data.zip_code | String | ❌ |  | |
| data.deleted_flag | Int | ❌ | 删除标志 | |
| data.status| Boolean | ✅ |  是否启用| |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.12 基础检验医嘱项目列表

```
请求地址：/dictionaries/Laboratorys
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinical_significance": null,
      "created_time": "2018-05-27T00:09:35.009827+08:00",
      "cuvette_color_name": null,
      "deleted_time": null,
      "en_name": null,
      "id": 1,
      "idc_code": null,
      "laboratory_sample": "尿液",
      "laboratory_sample_dosage": "",
      "name": "尿常规+尿流式沉渣检查",
      "py_code": "NCGNLSCZ",
      "remark": null,
      "time_report": null,
      "unit_name": null,
      "updated_time": "2018-05-27T00:09:35.009827+08:00"
    },
	...
  ],
  "page_info": {
    "total": 1464,
    "limit": "10",
    "offset": "0"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.en_name | String | ❌ | 英文名称 | |
| data.idc_code | String | ❌ | 国际编码 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.clinical_significance | 临床意义 | ❌ |  | |
| data.cuvette_color_name | String | ❌ | 试管颜色 | |
| data.laboratory_sample | String | ❌ | 检验物 | |
| data.laboratory_sample_dosage | String | ❌ | 检验物计量 | |
| data.remark | String | ❌ | 备注 | |
| data.time_report | String | ❌ | 报告所需时间 | |
| data.unit_name | Int | ❌ | 单位名称 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.13 基础检查医嘱项目列表

```
请求地址：/dictionaries/Examinations
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-05-27T00:14:07.58363+08:00",
      "deleted_time": null,
      "en_name": null,
      "id": 1,
      "idc_code": null,
      "name": "腕舟骨位(左)",
      "organ": null,
      "py_code": "WZGWZ",
      "remark": null,
      "unit_name": null,
      "updated_time": "2018-05-27T00:14:07.58363+08:00"
    },
	...
  ],
  "page_info": {
    "total": 1547,
    "limit": "10",
    "offset": "0"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.en_name | String | ❌ | 英文名称 | |
| data.idc_code | String | ❌ | 国际编码 | |
| data.py_code | String | ✅ | 拼音简码 | |
| data.organ | String | ❌ | 检查部位 | |
| data.remark | String | ❌ | 备注 | |
| data.unit_name | Int | ❌ | 单位名称 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.14 基础检验项目列表

```
请求地址：/dictionaries/LaboratoryItems
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_laboratory_item_id": 1,
      "name": "谷丙转氨酶 ",
      "en_name": "ALT ",
      "unit_name": "U/L ",
      "status": null,
      "is_special": false,
      "data_type": 2,
      "instrument_code": null,
      "is_delivery": null,
      "result_inspection": null,
      "default_result": null,
      "clinical_significance": null,
      "references": [
        {
          "reference_sex": "通用",
          "reference_max": "40.00",
          "reference_min": "5.00",
          "reference_value": null,
          "isPregnancy": null,
          "stomach_status": null
        }
      ]
    },
	...
  ],
  "page_info": {
    "count": 150,
    "limit": "10",
    "offset": "0"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.clinic_laboratory_item_id | Int | ✅ | 检验项目id | |
| data.name | String | ✅ | 名称 | |
| data.en_name | String | ❌ | 英文名称 | |
| data.unit_name | Int | ❌ | 单位名称 | |
| data.status | Boolean | ❌ | 是否启用 | |
| data.is_special | Boolean | ✅ | 参考值是否特殊 | |
| data.data_type | Int | ✅ | 数据类型 1 定性 2 定量 | |
| data.instrument_code | String | ❌ | 仪器编码 | |
| data.is_delivery | Boolean | ❌ | 是否允许外送 | |
| data.result_inspection | String | ❌ | 检验结果 | |
| data.default_result | String | ❌ | 默认结果 | |
| data.clinical_significance | String | ❌ |临床意义 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.15 基础药品列表

```
请求地址：/dictionaries/LaboratoryItems
```
**请求包示例**

```
{
	type:0
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| type | Int | ✅ |  类型 0-西药 1-中药 | |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "barcode": null,
      "code": "800656",
      "concentration": null,
      "country_flag": 0,
      "created_time": "2018-05-27T00:22:52.977059+08:00",
      "d_code": null,
      "dcode": "EGQITI",
      "default_remark": null,
      "deleted_time": null,
      "divide_flag": 1,
      "dosage": 1,
      "dosage_unit_name": "支",
      "dose_form_name": "注射剂",
      "drug_class_id": null,
      "drug_flag": 0,
      "english_name": null,
      "extend_code": null,
      "frequency_name": "一次",
      "id": 1,
      "infusion_flag": 2,
      "license_no": "国药准字H32022088",
      "low_dosage_flag": 0,
      "manu_factory_name": "常州千红生化制药股份有限公司",
      "mz_bill_item": null,
      "mz_charge_group": null,
      "name": "肝素钠注射液",
      "national_standard": "0065601",
      "once_dose": 12500,
      "once_dose_unit_name": "单位",
      "packing_unit_name": "盒",
      "preparation_count": null,
      "preparation_count_unit_name": null,
      "print_name": null,
      "py_code": "GSNZSY",
      "route_administration_name": "皮下注射",
      "self_flag": 0,
      "separate_flag": 0,
      "serial": "02",
      "spe_comment": null,
      "specification": "12500单位 2mlx1支/盒",
      "suprice_flag": 0,
      "sy_code": null,
      "type": 0,
      "updated_time": "2018-05-27T00:22:52.977059+08:00",
      "vol_unit_name": "ml",
      "volum": 2,
      "weight": 12500,
      "weight_unit_name": "单位",
      "zy_bill_item": null,
      "zy_charge_group": null
    },
	...
  ],
  "page_info": {
    "count": 150,
    "limit": "10",
    "offset": "0"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | 药品id | |
| data.type | Int | ✅ | 类型 0-西药 1-中药 | |
| data.code | String | ❌ | 编码 | |
| data.name | String | ✅ | 名称 | |
| data.py_code | String | ❌ | 拼音码 | |
| data.barcode | String | ❌ | 条形码 | |
| data.d_code | String | ❌ | 简码 | |
| data.print_name | String | ❌ | 打印名称 | |
| data.specification | String | ❌ | 规格 | |
| data.spe_comment | String | ❌ | 规格备注 | |
| data.manu_factory_name | String | ❌ |生产厂商 | |
| data.drug_class_id | Int | ❌ |药品类型id | |
| data.dose_form_name | String | ❌ |药品剂型 | |
| data.license_no | String | ❌ |国药准字、文号 | |
| data.once_dose | Int | ❌ |常用剂量| |
| data.once_dose_unit_name | String | ❌ |用量单位 常用剂量单位 | |
| data.dosage | Int | ❌ |剂量 | |
| data.dosage_unit_name | String | ❌ |剂量单位 | |
| data.preparation_count | Int | ❌ |制剂数量/包装量 | |
| data.preparation_count_unit_name| String | ❌ |制剂数量单位 | |
| data.packing_unit_name | String | ❌ |药品包装单位 | |
| data.route_administration_name| String | ❌ |用药途径/默认用法 | |
| data.frequency_name | String | ❌ |用药频率/默认频次| |
| data.default_remark| String | ❌ |默认用量用法说明 | |
| data.weight | Int | ❌ |重量 | |
| data.weight_unit_name | String | ❌ |重量单位 | |
| data.volum | Int | ❌ |体积 | |
| data.vol_unit_name | String | ❌ |体积单位 | |
| data.serial | String | ❌ |包装序号 | |
| data.national_standard | String | ❌ |国标分类 | |
| data.concentration | String | ❌ |浓度 | |
| data.dcode | String | ❌ |自定义码 | |
| data.infusion_flag | Int | ❌ |大输液标志,9为并开药 | |
| data.country_flag | Int | ❌ |进口标志 | |
| data.divide_flag| Int | ❌ |分装标志 | |
| data.low_dosage_flag | Int | ❌ |大规格小剂量标志| |
| data.self_flag | Int | ❌ |自费标识 | |
| data.separate_flag | Int | ❌ |单列标志 | |
| data.suprice_flag | Int | ❌ |贵重标志 | |
| data.drug_flag | Int | ❌ |毒麻标志 | |
| data.english_name | String | ❌ |英文名称 | |
| data.sy_code | String | ❌ |上药编码| |
| data.zy_bill_item | String | ❌ |住院帐单码 | |
| data.mz_bill_item | String | ❌ |门诊帐单码 | |
| data.zy_charge_group | String | ❌ |住院用药品分组| |
| data.mz_charge_group | String | ❌ |门诊用药品分组 | |
| data.extend_code | String | ❌ |药品与外界衔接码 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.16 供应商列表

```
请求地址：/dictionaries/SupplierList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-05-27T00:07:13.652046+08:00",
      "deleted_time": null,
      "id": 1,
      "name": "广州白云药厂",
      "updated_time": "2018-05-27T00:07:13.652046+08:00"
    },
    {
      "created_time": "2018-05-27T00:07:13.652046+08:00",
      "deleted_time": null,
      "id": 2,
      "name": "云南白药药厂",
      "updated_time": "2018-05-27T00:07:13.652046+08:00"
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
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.17 入库方式

```
请求地址：/dictionaries/InstockWayList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-05-27T00:07:13.652046+08:00",
      "deleted_time": null,
      "id": 1,
      "name": "采购入库",
      "updated_time": "2018-05-27T00:07:13.652046+08:00"
    },
    {
      "created_time": "2018-05-27T00:07:13.652046+08:00",
      "deleted_time": null,
      "id": 2,
      "name": "公益捐赠",
      "updated_time": "2018-05-27T00:07:13.652046+08:00"
    },
    {
      "created_time": "2018-08-12T22:54:14.528675+08:00",
      "deleted_time": null,
      "id": 3,
      "name": "门诊退药",
      "updated_time": "2018-08-12T22:54:14.528675+08:00"
    },
    {
      "created_time": "2018-08-12T22:54:14.528675+08:00",
      "deleted_time": null,
      "id": 4,
      "name": "零售退药",
      "updated_time": "2018-08-12T22:54:14.528675+08:00"
    }
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
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.18 出库方式

```
请求地址：/dictionaries/OutstockWayList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-05-27T00:07:13.652046+08:00",
      "deleted_time": null,
      "id": 1,
      "name": "科室领用",
      "updated_time": "2018-05-27T00:07:13.652046+08:00"
    },
    {
      "created_time": "2018-05-27T00:07:13.652046+08:00",
      "deleted_time": null,
      "id": 2,
      "name": "退货出库",
      "updated_time": "2018-05-27T00:07:13.652046+08:00"
    },
    {
      "created_time": "2018-05-27T00:07:13.652046+08:00",
      "deleted_time": null,
      "id": 3,
      "name": "报损出库",
      "updated_time": "2018-05-27T00:07:13.652046+08:00"
    },
    {
      "created_time": "2018-08-12T23:15:39.526287+08:00",
      "deleted_time": null,
      "id": 4,
      "name": "门诊发药",
      "updated_time": "2018-08-12T23:15:39.526287+08:00"
    },
    {
      "created_time": "2018-08-12T23:15:39.526287+08:00",
      "deleted_time": null,
      "id": 5,
      "name": "零售发药",
      "updated_time": "2018-08-12T23:15:39.526287+08:00"
    }
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
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>22.19 诊断字典

```
请求地址：/dictionaries/DiagnosisList
```
**请求包示例**

```
{
	keyword:
	offset:
	limit:
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| keyword | String | ❌ |  关键字 | |
| offset | int | ❌ |  开始条数 | 0|
| limit | int | ❌ |  条数 | 10|

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-05-31T22:18:18.87819+08:00",
      "deleted_time": null,
      "icd_code": null,
      "id": 586,
      "name": "鼻部囊肿",
      "py_code": "BBNZ",
      "updated_time": "2018-05-31T22:18:18.87819+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 35917
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String |  ❌ |  错误信息,code不为200时返回 | |
| data | Array | ✅ |   | |
| data.id | Int | ✅ | id | |
| data.name | String | ✅ | 名称 | |
| data.icd_code | String | ❌ | 国际编码 | |
| data.py_code | String | ❌ | 拼音简码 | |
| data.created_time | time | ✅ | 创建时间 | |
| data.updated_time | time | ✅ | 更新时间 | |
| data.deleted_time | time | ❌ | 删除时间 | |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

23 主诉模板
--------

</br>
<h3>23.1 检验医嘱创建

```
请求地址：/chiefComplaint/list
```

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "name": "发热"
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| name | String | ✅ | 主诉名称 | |


24 门诊发药
--------


</br>
<h3>24.1 获取待发药的分诊记录

```
请求地址：/drugDelivery/traige/waiting/list
```
**请求包示例**

```
{
	keyword: ""
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2019-01-01'
	offset: 0
	limit: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword |number | ❌ |  关键字 | |
| start_date |number | ✅ |  开始日期 | |
| end_date |number | ✅ |  结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 2,
      "department_name": "骨科",
      "doctor_name": "扁鹊",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "18701676735",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-05-28T17:02:21.728863+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "total_count": 9,
      "updated_time": "2018-05-31T22:07:56.798037+08:00",
      "visit_date": "2018-05-28T00:00:00Z"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 22
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 患者分诊id | |
| department_name | String | ✅ | 科室名称 | |
| doctor_name | String | ✅ | 医生名称 | |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者姓名 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 挂号操作员 | |
| register_time | String | ✅ | 挂号时间 | |
| register_type | String | ✅ | 1预约，2线下分诊 3快速接诊 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 10:登记，20：分诊(换诊)，30：接诊，40：已就诊 | |
| total_count | String | ✅ | 总的细目数 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 就诊日期 | |

</br>
<h3>24.2 获取已发药的分诊记录

```
请求地址：/drugDelivery/traige/issued/list
```
**请求包示例**

```
{
	keyword: ""
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2019-01-01'
	offset: 0
	limit: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword |number | ❌ |  关键字 | |
| start_date |number | ✅ |  开始日期 | |
| end_date |number | ✅ |  结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 2,
      "department_name": "骨科",
      "doctor_name": "扁鹊",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "18701676735",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-05-28T17:02:21.728863+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "total_count": 9,
      "updated_time": "2018-05-31T22:07:56.798037+08:00",
      "visit_date": "2018-05-28T00:00:00Z"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 22
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 患者分诊id | |
| department_name | String | ✅ | 科室名称 | |
| doctor_name | String | ✅ | 医生名称 | |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者姓名 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 挂号操作员 | |
| register_time | String | ✅ | 挂号时间 | |
| register_type | String | ✅ | 1预约，2线下分诊 3快速接诊 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 10:登记，20：分诊(换诊)，30：接诊，40：已就诊 | |
| total_count | String | ✅ | 总的细目数 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 就诊日期 | |


</br>
<h3>24.3 获取已退药的分诊记录

```
请求地址：/drugDelivery/traige/refund/list
```
**请求包示例**

```
{
	keyword: ""
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2019-01-01'
	offset: 0
	limit: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword |number | ❌ |  关键字 | |
| start_date |number | ✅ |  开始日期 | |
| end_date |number | ✅ |  结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 2,
      "department_name": "骨科",
      "doctor_name": "扁鹊",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "18701676735",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-05-28T17:02:21.728863+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "total_count": 9,
      "updated_time": "2018-05-31T22:07:56.798037+08:00",
      "visit_date": "2018-05-28T00:00:00Z"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 22
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 患者分诊id | |
| department_name | String | ✅ | 科室名称 | |
| doctor_name | String | ✅ | 医生名称 | |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者姓名 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 挂号操作员 | |
| register_time | String | ✅ | 挂号时间 | |
| register_type | String | ✅ | 1预约，2线下分诊 3快速接诊 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 10:登记，20：分诊(换诊)，30：接诊，40：已就诊 | |
| total_count | String | ✅ | 总的细目数 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 就诊日期 | |


</br>
<h3>24.4 获取发药药品详情列表

```
请求地址：/drugDelivery/list
```
**请求包示例**

```
{
	order_status:10
	clinic_triage_patient_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| order_status |number | ✅ |  包括 待发药 - 10，已发药 - 30，已退药 - 40 | |
| clinic\_triage\_patient_id |number | ✅ |  分诊记录id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "amount": 20,
      "charge_project_type_id": 2,
      "dose_form_name": "动物类",
      "id": 6,
      "manu_factory_name": "北京同仁堂饮片厂",
      "name": "穿山甲",
      "order_sn": "20180528222",
      "order_status": "10",
      "prescription_amount": 20,
      "remark": null,
      "specification": "/kg",
      "stock_amount": null
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| amount | String | ✅ | 数量 | |
| charge\_project\_type_id | String | ✅ | 项目类型id | |
| dose\_form_name | String | ✅ | 药品剂型 | |
| name | String | ✅ | 项目名称 | |
| manu\_factory_name | String | ✅ | 生成厂商 | |
| order_sn | String | ✅ | 单号 | |
| order_status | String | ✅ | 发药状态 （10=待发药 30-已发药 40=已退药） | |
| prescription_amount | String | ✅ | 处方数量 | |
| specification | String | ✅ | 规格 | |
| stock_amount | String | ✅ | 库存 | |

</br>
<h3>24.5 创建发药记录

```
请求地址：/drugDelivery/record/create
```
**请求包示例**

```
{
	clinic_triage_patient_id:1
	operation_id: 1
	items:[{
	    mz_paid_orders_id: 1
	    remark: ''
	
	}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊id | |
| operation_id |number | ✅ |  操作员id | |
| mz\_paid\_orders_id |number | ✅ |  门诊已缴费id | |
| remark | string | ✅ |  备注 | |

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
| code | String | ✅ | 200 成功 | |

</br>
<h3>24.6 退药

```
请求地址：/drugDelivery/record/refund
```
**请求包示例**

```
{
	clinic_triage_patient_id:1
	operation_id: 1
	items:[{
	    mz_paid_orders_id: 1
	}]
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊id | |
| operation_id |number | ✅ |  操作员id | |
| mz\_paid\_orders_id |number | ✅ |  门诊已缴费id | |


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
| code | String | ✅ | 200 成功 | |

</br>
<h3>24.7 查询发药记录

```
请求地址：/drugDelivery/record/list
```
**请求包示例**

```
{
	clinic_triage_patient_id:1
	offset: 0
	limit : 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊记录id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-06-04T00:06:19.505112+08:00",
      "doctor_name": "黄飞鸿",
      "drug_delivery_record_id": 17,
      "opration_name": "超级管理员",
      "project_name": "茶苯海明片",
      "visit_date": "2018-06-03T00:00:00Z"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 2
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| created_time | String | ✅ | 创建时间 | |
| doctor_name | String | ✅ | 医生名称 | |
| drug\_delivery\_record_id | String | ✅ | 发药记录id | |
| opration_name | String | ✅ | 操作员 | |
| project_name | String | ✅ | 项目名称 | |
| visit_date | String | ✅ | 就诊日期 | |

</br>
<h3>24.8 查询发药记录详情

```
请求地址：/drugDelivery/record/detail
```
**请求包示例**

```
{
	drug_delivery_record_id:1
   offset: 0
	limit : 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| drug\_delivery\_record_id |number | ✅ |  发药记录id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "amount": 1,
      "charge_project_type_id": 1,
      "manu_factory_name": "浙江东日药业有限公司",
      "name": "茶苯海明片",
      "order_status": "40",
      "remark": "",
      "specification": "25mg*1片"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 2
  }
}
```

**应答包参数说明** （同 24.4）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功 | |
</br>
<h3>24.9 查询退药记录

```
请求地址：/drugDelivery/record/refund/list
```
**请求包示例**

```
{
	clinic_triage_patient_id:1
	offset: 0
	limit : 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊记录id | |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-06-06T01:49:09.70884+08:00",
      "doctor_name": "黄飞鸿",
      "drug_delivery_refund_record_id": 3,
      "opration_name": "超级管理员",
      "project_name": "甘草稍",
      "visit_date": "2018-06-03T00:00:00Z"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 1
  }
}
```

**应答包参数说明** (通 24.7)

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功 | |
</br>
<h3>24.10 查询发药记录详情

```
请求地址：/drugDelivery/record/refund/detail
```
**请求包示例**

```
{
	drug_delivery_refund_record_id:1
	offset: 0
	limit: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "amount": 1,
      "charge_project_type_id": 1,
      "manu_factory_name": "浙江东日药业有限公司",
      "name": "茶苯海明片",
      "order_status": "40",
      "specification": "25mg*1片"
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 7
  }
}
```

**应答包参数说明** (同24.4)

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ | 200 成功 | |
</br>

25 检查模块
--------

</br>
<h3>25.1 获取检查列表

```
请求地址：/examinationTriage/ExaminationTriageList
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	order_status: 0
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊记录id | |
| order_status |number | ✅ |  状态 （10-待检查 20-检查中 30-已检查） | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_examination_id": 3,
      "clinic_examination_name": "盆腔彩色多普勒超声",
      "clinic_triage_patient_id": 30,
      "conclusion_examination": null,
      "examination_patient_id": 9,
      "examination_patient_record_id": null,
      "organ": null,
      "picture_examination": null,
      "report_doctor_name": null,
      "report_time": null,
      "result_examination": null
    }
  ]
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_examination_id | String | ✅ | 检查项目id | |
| clinic\_examination_name | String | ✅ | 检查项目名称 | |
| clinic\_triage\_patient_id | String | ✅ | 分诊记录id | |
| conclusion_examination | String | ✅ | 检查结论 | |
| examination\_patient_id | String | ✅ | 检查记录id | |
| examination\_patient\_record_id | String | ✅ | 检查报告id | |
| organ | String | ✅ | 脏器 | |
| picture_examination | String | ✅ | 检查图片 | |
| report\_doctor_name | String | ✅ | 报告医生 | |
| report_time | String | ✅ | 报告时间 | |
| result_examination | String | ✅ | 检查结果 | |

</br>
<h3>25.2 获取待检查的分诊记录

```
请求地址：/examinationTriage/ExaminationTriageWaiting
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 80,
      "department_name": "牙科",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-12T20:34:17.270198+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "187011735",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-07-19T21:00:26.796976+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:51:49.416355+08:00",
      "visit_date": "2018-07-19T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 8
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检查数量 | |
| checking\_total_count | String | ✅ | 检查中的数量 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间| |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者名称 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 登记操作员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 挂号记录状态 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 接诊日期 | |
| waiting\_total_count | String | ✅ | 待检查数量 | |

</br>
<h3>25.3 获取已检查的分诊记录

```
请求地址：/examinationTriage/ExaminationTriageChecked
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 80,
      "department_name": "牙科",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-12T20:34:17.270198+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "187011735",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-07-19T21:00:26.796976+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:51:49.416355+08:00",
      "visit_date": "2018-07-19T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 8
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检查数量 | |
| checking\_total_count | String | ✅ | 检查中的数量 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间| |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者名称 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 登记操作员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 挂号记录状态 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 接诊日期 | |
| waiting\_total_count | String | ✅ | 待检查数量 | |


</br>
<h3>25.4 获取检查中的分诊记录

```
请求地址：/examinationTriage/ExaminationTriageChecking
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 80,
      "department_name": "牙科",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-12T20:34:17.270198+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "187011735",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-07-19T21:00:26.796976+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:51:49.416355+08:00",
      "visit_date": "2018-07-19T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 8
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检查数量 | |
| checking\_total_count | String | ✅ | 检查中的数量 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间| |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者名称 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 登记操作员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 挂号记录状态 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 接诊日期 | |
| waiting\_total_count | String | ✅ | 待检查数量 | |

</br>
<h3>25.5 创建检查记录

```
请求地址：/examinationTriage/ExaminationTriageRecordCreate
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	examination_patient_id: 1
	operation_id: 1
	picture_examination: ‘’
	result_examination: ‘’
	conclusion_examination: ‘’
}
```
**请求包参数说明** （同 25.1）


**应答包示例**

```
{
  "code": "200",
  "data": “”
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |
</br>
<h3>25.6 查询检查记录

```
请求地址：/examinationTriage/ExaminationTriageRecordList
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
}
```
**请求包参数说明** （同 25.1）


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_examination_name": "胸部正位",
      "conclusion_examination": "检查结论更不错哟",
      "created_time": "2018-06-05T21:38:37.98011+08:00",
      "doctor_name": "华佗",
      "examination_patient_record_id": 1,
      "opration_name": "华佗",
      "picture_examination": "检查图片地址",
      "result_examination": "检查结果很不错哦",
      "visit_date": "2018-06-03T00:00:00Z"
    }
  ]
}
```

**应答包参数说明** (同25.1)

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |

</br>
<h3>25.7 更新检查状态

```
请求地址：/examinationTriage/ExaminationTriageUpdate
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	order_status: 10
	items:[{
	    examination_patient_id
	}]
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| order_status | String | ✅ | 状态 10-待检测 20-检查中 30-已检查 | |
| examination\_patient_id | String | ✅ | 检查记录id | |


**应答包示例**

```
{
  "code": "200",
  "data": 
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |


</br>
<h3>25.8 患者历史检查记录

```
请求地址：/examinationTriage/ExaminationTriagePatientRecordList
```

**请求包示例**

```
{
	clinic_triage_patient_id: 1
	patient_id: 1
	offset: 1
	limit: 1
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| patient_id | String | ✅ | 患者id | |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_examination_name": "胸部正位+侧位，胸部正位",
      "clinic_name": "龙华诊所",
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 97,
      "department_name": "牙科",
      "order_doctor_name": "胡一天",
      "order_time": "2018-07-27T17:36:02.235297+08:00",
      "patient_id": 3
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 3
  }
}
```

**应答包参数说明** （同25.1）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |


26 检验报告模板
--------

</br>
<h3>26.1 获取检验列表

```
请求地址：/laboratoryTriage/LaboratoryTriageList
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	order_status: 0
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊记录id | |
| order_status |number | ✅ |  状态 （10-待检验 20-检验中 30-已检验） | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "checking_time": "2018-06-03T23:34:24.423962+08:00",
      "clinic_laboratory_id": 1,
      "clinic_laboratory_name": "血常规",
      "clinic_triage_patient_id": 31,
      "laboratory_patient_id": 104,
      "laboratory_patient_record_id": null,
      "laboratory_sample": "全血",
      "remark": null,
      "report_doctor_name": null,
      "report_time": null
    }
  ]
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| checking_time | String | ✅ | 检验时间 | |
| clinic\_laboratory_id | String | ✅ | 检验项目id | |
| clinic\_laboratory_name | String | ✅ | 检验项目名称 | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id| |
| laboratory\_patient_id | String | ✅ | 检验记录id | |
| laboratory\_patient\_record_id | String | ✅ | 检验报告id | |
| laboratory_sample | String | ✅ | 检验物 | |
| remark | String | ✅ | 备注 | |
| report\_doctor_name | String | ✅ | 报告医生 | |
| report_time | String | ✅ | 报告时间 | |

</br>
<h3>26.2 获取检验详情

```
请求地址：/laboratoryTriage/LaboratoryTriageDetail
```
**请求包示例**

```
{
	laboratory_patient_id: 1
	clinic_laboratory_id: 1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| laboratory\_patient_id |number | ✅ |  检验记录id | |
| clinic\_laboratory_id | number | ✅ |  检验id | |

**应答包示例**

```
{
  "associationsData": [
    {
      "clinic_laboratory_item_id": 2,
      "name": "血小板",
      "en_name": "xuexiaoban",
      "unit_name": "/L",
      "status": true,
      "is_special": false,
      "data_type": 2,
      "instrument_code": null,
      "is_delivery": null,
      "result_inspection": null,
      "default_result": "100",
      "clinical_significance": null,
      "references": [
        {
          "reference_sex": "通用",
          "reference_max": "20",
          "reference_min": "10",
          "reference_value": null,
          "isPregnancy": false,
          "stomach_status": "false"
        }
      ]
    }
  ],
  "code": "200",
  "resultsData": null
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_laboratory\_item_id | String | ✅ | 子项id | |
| name | String | ✅ | 检查项名称 | |
| en_name | String | ✅ | 英文名称 | |
| unit_name | String | ✅ | 单位 | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| status | String | ✅ |  是否启用 | |
| is_special | String | ✅ | 参考值是否特殊 | |
| result_inspection | String | ✅ |  检验结果 | |
| default_result | String | ✅ |  默认结果 | |
| clinical_significance | String | ✅ |  临床意义 | |
| references.reference_sex | String | ✅ |  参考值性别 男、女、通用 | |
| references.reference_max | String | ✅ |  参考最大值 | |
| references.reference_min | String | ✅ |  参考最小值 | |
| references.reference_value | String | ✅ |  定性参考值 | |
| references.isPregnancy |boolean | ✅ |  是否妊娠期 | |
| references.stomach_status | boolean | ✅ |  时候空腹 | |


</br>
<h3>26.3 获取待检验的分诊记录

```
请求地址：/laboratoryTriage/LaboratoryTriageWaiting
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "allergic_history": "134",
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 80,
      "department_name": "牙科",
      "diagnosis": "皮肤干燥性皮炎",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-12T20:33:51.259219+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "187015",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-07-19T21:00:26.796976+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:51:49.416355+08:00",
      "visit_date": "2018-07-19T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 14
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| allergic_history | String | ✅ | 过敏史 | |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检验的项目 | |
| checking\_total_count | String | ✅ | 检验中的项目 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| diagnosis | String | ✅ | 诊断 | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者姓名 | |
| phone | String | ✅ | 手机号 | |
| register\_personnel_name | String | ✅ | 登记人员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型：1预约，2线下分诊 3快速接诊 | |
| sex | String | ✅ | 患者性别 |
| status | String | ✅ | 接诊状态 0:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消|
| updated_time | String | ✅ | 更新时间 |
| visit_date | String | ✅ | 接诊日期 |
| waiting\_total_count | String | ✅ | 待检验的数量 |


</br>
<h3>26.4 获取已检验的分诊记录

```
请求地址：/laboratoryTriage/LaboratoryTriageChecked
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "allergic_history": "134",
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 80,
      "department_name": "牙科",
      "diagnosis": "皮肤干燥性皮炎",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-12T20:33:51.259219+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "187015",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-07-19T21:00:26.796976+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:51:49.416355+08:00",
      "visit_date": "2018-07-19T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 14
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| allergic_history | String | ✅ | 过敏史 | |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检验的项目 | |
| checking\_total_count | String | ✅ | 检验中的项目 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| diagnosis | String | ✅ | 诊断 | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者姓名 | |
| phone | String | ✅ | 手机号 | |
| register\_personnel_name | String | ✅ | 登记人员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型：1预约，2线下分诊 3快速接诊 | |
| sex | String | ✅ | 患者性别 |
| status | String | ✅ | 接诊状态 0:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消|
| updated_time | String | ✅ | 更新时间 |
| visit_date | String | ✅ | 接诊日期 |
| waiting\_total_count | String | ✅ | 待检验的数量 |


</br>
<h3>26.5 获取检验中的分诊记录

```
请求地址：/laboratoryTriage/LaboratoryTriageChecking
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "allergic_history": "134",
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 80,
      "department_name": "牙科",
      "diagnosis": "皮肤干燥性皮炎",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-12T20:33:51.259219+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "187015",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-07-19T21:00:26.796976+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:51:49.416355+08:00",
      "visit_date": "2018-07-19T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 14
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| allergic_history | String | ✅ | 过敏史 | |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检验的项目 | |
| checking\_total_count | String | ✅ | 检验中的项目 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| diagnosis | String | ✅ | 诊断 | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者姓名 | |
| phone | String | ✅ | 手机号 | |
| register\_personnel_name | String | ✅ | 登记人员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型：1预约，2线下分诊 3快速接诊 | |
| sex | String | ✅ | 患者性别 |
| status | String | ✅ | 接诊状态 0:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消|
| updated_time | String | ✅ | 更新时间 |
| visit_date | String | ✅ | 接诊日期 |
| waiting\_total_count | String | ✅ | 待检验的数量 |


</br>
<h3>26.6 创建检验记录

```
请求地址：/laboratoryTriage/LaboratoryTriageRecordCreate
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	laboratory_patient_id: 1
	operation_id: 1
	remark: ‘’
	item: [{
	    clinic_laboratory_item_id: 1
	    result_inspection: 1
	    reference_max: 1
	    reference_min: 1
	    reference_value: 1
	    data_type: 1
	    is_normal: 1
	}]
}
```
**请求包参数说明** （同 26.2）


**应答包示例**

```
{
  "code": "200",
  "data": “”
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |
</br>
<h3>26.7 查询检验记录

```
请求地址：/laboratoryTriage/LaboratoryTriageRecordList
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
}
```
**请求包参数说明** （同 26.1）

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_laboratory_name": "血常规",
      "created_time": "2018-07-24T00:58:52.265271+08:00",
      "doctor_name": "胡一天",
      "laboratory_patient_record_id": 7,
      "opration_name": "超级管理员",
      "remark": "正常",
      "visit_date": "2018-07-20T00:00:00Z"
    }
  ]
}
```

**应答包参数说明** (同26.1)

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |

</br>
<h3>26.8 查询检验详情

```
请求地址：/laboratoryTriage/LaboratoryTriageRecordDetail
```

**请求包示例**

```
{
	laboratory_patient_record_id: 1
	offset: 1
	limit: 10
}
```
**请求包参数说明** （同 26.1）

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_laboratory_item_id": 5,
      "name": "白细胞计数",
      "en_name": "WBC",
      "unit_name": "个/L",
      "status": true,
      "is_special": true,
      "data_type": 2,
      "instrument_code": null,
      "is_delivery": null,
      "result_inspection": "12",
      "default_result": null,
      "clinical_significance": null,
      "references": [
        {
          "reference_sex": "通用",
          "reference_max": "15",
          "reference_min": "10",
          "reference_value": null,
          "isPregnancy": null,
          "stomach_status": null
        }
      ]
    }
  ],
  "page_info": {
    "limit": "10",
    "offset": "0"
  }
}
```

**应答包参数说明** (同26.2)

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| 

</br>
<h3>26.9 更新检查状态

```
请求地址：/laboratoryTriage/LaboratoryTriageUpdate
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	order_status: 10
	items:[{
	    laboratory_patient_id
	}]
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| order_status | String | ✅ | 状态 10-待检测 20-检查中 30-已检查 | |
| laboratory\_patient_id | String | ✅ | 检验记录id | |


**应答包示例**

```
{
  "code": "200",
  "data": 
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |


</br>
<h3>26.10 患者历史检验记录

```
请求地址：/laboratoryTriage/LaboratoryTriagePatientRecordList
```

**请求包示例**

```
{
	clinic_triage_patient_id: 1
	patient_id: 1
	offset: 1
	limit: 1
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| patient_id | String | ✅ | 患者id | |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_laboratory_name": "尿常规，血常规",
      "clinic_name": "龙华诊所",
      "clinic_patient_id": 16,
      "clinic_triage_patient_id": 81,
      "department_name": "牙科",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-07-24T00:48:06.748606+08:00",
      "patient_id": 20
    }patient_id": 20
    }
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 2
  }
}
```

**应答包参数说明** （同26.1）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |


27 治疗模板
--------

</br>
<h3>27.1 获取治疗列表

```
请求地址：/treatmentTriage/TreatmentTriageList
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	order_status: 0
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id |number | ✅ |  分诊记录id | |
| order_status |number | ✅ |  状态 （10-待检查 20-检查中 30-已检查） | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_treatment_id": 12,
      "clinic_treatment_name": "输液治疗",
      "clinic_triage_patient_id": 31,
      "illustration": "法法",
      "left_times": 2,
      "price": 10000,
      "remark": null,
      "report_doctor_name": null,
      "report_time": null,
      "total_times": 2,
      "treatment_patient_id": 7,
      "treatment_patient_record_id": null,
      "unit_name": "次"
    }
  ]
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_treatment_id | String | ✅ | 治疗项目id | |
| clinic\_treatment_name | String | ✅ | 治疗项目名称 | |
| clinic\_triage\_patient_id | String | ✅ | 分诊记录id | |
| illustration | String | ✅ | 说明 | |
| left_times | String | ✅ | 剩余次数 | |
| price | String | ✅ | 价格 | |
| remark | String | ✅ | 备注 | |
| picture_examination | String | ✅ | 检查图片 | |
| report\_doctor_name | String | ✅ | 报告医生 | |
| report_time | String | ✅ | 报告时间 | |
| total_times | String | ✅ | 总次数 | |
| treatment\_patient_id | String | ✅ | 治疗记录id | |
| treatment\_patient\_record_id | String | ✅ | 治疗报告id | |
| unit_name | String | ✅ | 单位 | |

</br>
<h3>27.2 获取待治疗的分诊记录

```
请求地址：/treatmentTriage/TreatmentTriageWaiting
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 117,
      "department_name": "牙科",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-05T23:19:32.854022+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "18701676735",
      "register_personnel_name": null,
      "register_time": "2018-07-31T23:53:56.924887+08:00",
      "register_type": 3,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:56:54.922227+08:00",
      "visit_date": "2018-07-31T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 8
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检查数量 | |
| checking\_total_count | String | ✅ | 检查中的数量 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间| |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者名称 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 登记操作员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 挂号记录状态 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 接诊日期 | |
| waiting\_total_count | String | ✅ | 待检查数量 | |

</br>
<h3>27.3 获取已治疗的分诊记录

```
请求地址：/treatmentTriage/TreatmentTriageChecked
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 80,
      "department_name": "牙科",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-12T20:34:17.270198+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "187011735",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-07-19T21:00:26.796976+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:51:49.416355+08:00",
      "visit_date": "2018-07-19T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 8
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检查数量 | |
| checking\_total_count | String | ✅ | 检查中的数量 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间| |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者名称 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 登记操作员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 挂号记录状态 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 接诊日期 | |
| waiting\_total_count | String | ✅ | 待检查数量 | |


</br>
<h3>27.4 获取治疗中的分诊记录

```
请求地址：/treatmentTriage/TreatmentTriageChecking
```
**请求包示例**

```
{
	keyword: '关键词'
	clinic_id: 1
	start_date: '2018-01-01'
	end_date: '2018-02-02'
	offset: 0
	limit: 10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id |number | ✅ |  诊所id | |
| keyword | string | ❌ |  关键词 | |
| start_date | string | ✅ |  开始时间 | |
| end_date | string | ✅ |  结束时间 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "19920706",
      "checked_total_count": 0,
      "checking_total_count": 0,
      "clinic_patient_id": 2,
      "clinic_triage_patient_id": 80,
      "department_name": "牙科",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-08-12T20:34:17.270198+08:00",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "187011735",
      "register_personnel_name": "超级管理员",
      "register_time": "2018-07-19T21:00:26.796976+08:00",
      "register_type": 2,
      "sex": 1,
      "status": 40,
      "updated_time": "2018-08-12T22:51:49.416355+08:00",
      "visit_date": "2018-07-19T00:00:00Z",
      "waiting_total_count": 2
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 8
  }
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| birthday | String | ✅ | 患者生日 | |
| checked\_total_count | String | ✅ | 已检查数量 | |
| checking\_total_count | String | ✅ | 检查中的数量 | |
| clinic\_patient_id | String | ✅ | 诊所患者id | |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| department_name | String | ✅ | 科室名称 | |
| order\_doctor_name | String | ✅ | 开单医生 | |
| order_time | String | ✅ | 开单时间| |
| patient_id | String | ✅ | 患者id | |
| patient_name | String | ✅ | 患者名称 | |
| phone | String | ✅ | 患者手机号 | |
| register\_personnel_name | String | ✅ | 登记操作员 | |
| register_time | String | ✅ | 登记时间 | |
| register_type | String | ✅ | 登记类型 | |
| sex | String | ✅ | 患者性别 | |
| status | String | ✅ | 挂号记录状态 | |
| updated_time | String | ✅ | 更新时间 | |
| visit_date | String | ✅ | 接诊日期 | |
| waiting\_total_count | String | ✅ | 待检查数量 | |

</br>
<h3>27.5 创建治疗记录

```
请求地址：/treatmentTriage/TreatmentTriageRecordCreate
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	operation_id: 1
	items : [{
	    treatment_patient_id: 1
	    times: 1
	    remark: ''
	}]
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | String | ✅ |分诊记录id | |
| operation_id | String | ✅ | 操作员 | |
| treatment\_patient_id | String | ✅ | 治疗记录id | |
| times | String | ✅ | 次数| |
| remark | String | ✅ | 备注| |

**应答包示例**

```
{
  "code": "200",
  "data": “”
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |
</br>
<h3>27.6 查询治疗记录

```
请求地址：/treatmentTriage/TreatmentTriageRecordList
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
}
```
**请求包参数说明** （同 27.1）


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_treatment_name": "针灸",
      "created_time": "2018-06-29T00:14:27.139161+08:00",
      "doctor_name": "华佗",
      "left_times": 0,
      "opration_name": "超级管理员",
      "price": 9000,
      "remark": "",
      "times": 3,
      "total_times": 3,
      "treatment_patient_record_id": 1,
      "unit_name": "次",
      "visit_date": "2018-06-26T00:00:00Z"
    }
  ]
}
```

**应答包参数说明** (同27.1)

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |

</br>
<h3>27.7 更新治疗状态

```
请求地址：/treatmentTriage/TreatmentTriageUpdate
```
**请求包示例**

```
{
	clinic_triage_patient_id: 1
	order_status: 10
	items:[{
	    treatment_patient_id
	}]
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| order_status | String | ✅ | 状态 10-待检测 20-检查中 30-已检查 | |
| treatment\_patient_id | String | ✅ | 治疗记录id | |


**应答包示例**

```
{
  "code": "200",
  "data": 
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |


</br>
<h3>27.8 患者历史治疗记录

```
请求地址：/treatmentTriage/TreatmentTriagePatientRecordList
```

**请求包示例**

```
{
	clinic_triage_patient_id: 1
	patient_id: 1
	offset: 1
	limit: 1
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic\_triage\_patient_id | String | ✅ | 分诊id | |
| patient_id | String | ✅ | 患者id | |


**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_name": "龙华诊所",
      "clinic_patient_id": 6,
      "clinic_treatment_name": "打针，针灸",
      "clinic_triage_patient_id": 54,
      "department_name": "眼科",
      "order_doctor_name": "超级管理员",
      "order_time": "2018-06-28T22:26:37.055847+08:00",
      "patient_id": 10
    }
  ],
  "page_info": {
    "limit": "1",
    "offset": "0",
    "total": 1
  }
}
```

**应答包参数说明** （同27.1）

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |


</br>
<h3>27.9 删除治疗医嘱模板

```
请求地址：/treatmentTriage/TreatmentPatientModelDelete
```

**请求包示例**

```
{
	treatment_patient_model_id: 1
}
```
**请求包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| treatment\_patient\_model_id | String | ✅ | 治疗医嘱模板 | ||


**应答包示例**

```
{
  "code": "200",
  "data": “ok”
}
```

**应答包参数说明** 

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| data | String | ✅ | 200 成功 | |
--

28 零售模块
--------

</br>
<h3>28.1 创建药品零售订单

```
请求地址：/drugRetail/createOrder
```
**请求包示例**

```
items:[
	{"clinic_drug_id":"1","amount":"3","total_fee":"15"},
	{"clinic_drug_id":"2","amount":"2","total_fee":"8"}
]
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| items | Array | ✅ |  零售药品信息| |
| items.clinic_drug_id | String | ✅ |  诊所药品id | |
| items.amount | String | ✅ |  数量 | |
| items.total_fee | String | ✅ |  总价 | |

**应答包示例**

```
{
    "code": "200",
    "msg":"ok",
    "data": "O1201808052049528177"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | String | ❌ |  订单号 | |
--

</br>
<h3>28.2 创建支付订单

```
请求地址：/drugRetail/CreateDrugRetailPaymentOrder
```
**请求包示例**

```
{
	record_sn:O1201808052301195376
	pay_method:cash
	auth_code:
	total_money:4
	discount_money:0
	medical_money:0
	balance_money:4
	operation_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| record_sn | String | ✅ |  订单号| |
| pay_method | String | ✅ |  支付方式， wechat - 微信，alipay-支付宝 ,bank - 银行, cash- 现金 | |
| auth_code | String | ❌ |  支付方式为 wechat、alipay 时不能为空 | |
| total_money | Int | ✅ |  交易总金额金额| |
| discount_money | Int | ✅ |  折扣金额 | |
| medical_money | Int | ✅ |  医保金额 | |
| balance_money | Int | ✅ |  实收金额 | |
| operation_id | Int | ✅ |  操作员id | |

**应答包示例**

```
{
    "code": "200",
    "msg":"ok",
    "data": 1
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
| data | Int | ❌ |  code 200时为1 | |
--

</br>
<h3>28.3 获取收费状态

```
请求地址：/drugRetail/payment/status
```
**请求包示例**

```
{
	out_trade_no:T1201808021200256713
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| out_trade_no | String | ✅ |  交易号,一次交易的编码| |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "biz_channel": "99",
    "biz_type": "default",
    "discount_fee": 0,
    "merchant_id": "2017081008129270",
    "nonce_str": "7dn4jypj9htblo8oli06",
    "openid": "2088702374593053",
    "out_trade_no": "T1201808021200256713",
    "pay_mode": "alipay_f2f",
    "pay_order_id": 1808026010800000000,
    "payed_time": "20180802120027",
    "real_total_fee": 2,
    "result_code": "0",
    "result_msg": "查询成功",
    "return_code": "0",
    "return_msg": "OK",
    "sign": "57807DE69445E0E1A6B1FC311E7B5155",
    "sign_type": "MD5",
    "total_fee": 2,
    "trade_status": "SUCCESS",
    "transaction_id": "2018080221001004050534884599"
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ❌ |  返回信息 | |
| data.biz_channel | String | ✅ |  渠道 | |
| data.biz_type | String | ✅ | 类型 | |
| data.discount_fee | String | ✅|  优惠金额| |
| data.merchant_id | String | ✅ |  支付机构商户号 | |
| data.nonce_str | String | ✅ |  随机字符串 | |
| data.openid | String | ✅ |  用户唯一标识 微信为 openid,支付宝为 buyer_id | |
| data.out_trade_no | String | ✅ |  商户订单号 | |
| data.pay_mode | String | ✅ |  支付方式 | |
| data.pay_order_id | String | ✅ |   | |
| data.payed_time | String | ✅ |  支付时间 | |
| data.real_total_fee | String | ✅ |  实际支付金额 | |
| data.result_code | String | ✅ |  业务状态码 业务状态 0-成功其他失败，订单的最终状态 以 trade_status 为准 | |
| data.result_msg | String | ✅ |  业务状态描述 | |
| data.return_code | String | ✅ |  交易结果代码 0:成功其他:失败为 0 时仅表示通讯成功标志 | |
| data.return_msg | String | ✅ |  错诨消息 | |
| data.sign | String | ✅ |  签名 | |
| data.sign_type | String | ✅ |  签名类型 | |
| data.total_fee | String | ✅ |  订单金额 | |
| data.trade_status | String | ✅ |  订单状态 SUCCESS-成功 REFUND-转入退款 USERPAYING-用户支付中 NOTPAY-未支付 CLOSE-已关闭 | |
| data.transaction_id | String | ✅ |  支付机构支付流水号 | |

--

</br>
<h3>28.4 获取药品零售表

```
请求地址：/drugRetail/list
```
**请求包示例**

```
{
	refundStatus:2
	start_date:2018-07-01
	end_date:2018-08-01
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| refundStatus | Int | ✅ |  收费、退费标志 1 收费 2 退费 | |
| start_date | Date | ✅ |  开始时间 | |
| end_date | Date | ✅ | 结束时间 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "balance_money": 8000,
      "name": "超级管理员",
      "out_trade_no": "T1201807271716197099",
      "pay_method": "cash",
      "pay_time": "2018-07-27T17:16:19.715943+08:00",
      "refund_money": -8000
    },
    {
      "balance_money": 1,
      "name": "超级管理员",
      "out_trade_no": "T1201807261837307928",
      "pay_method": "cash",
      "pay_time": "2018-07-26T18:37:30.330244+08:00",
      "refund_money": -1
    }
  ],
  "page_info": {
    "limit": "6",
    "offset": "0",
    "total": 2
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.balance_money | Int | ✅ | 金额| |
| data.items.name | String | ✅ |  收费人员 | |
| data.items.out_trade_no | String | ✅ |  交易单号 | |
| data.items.pay_method | String | ✅ |  支付方式 wechat - 微信，alipay-支付宝 ,bank - 银行, cash- 现金| |
| data.items.pay_time | time | ✅ |  支付时间| |
| data.items.refund_money | Int | ❌ |  退款金额| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>28.5 获取药品详情

```
请求地址：/drugRetail/detail
```
**请求包示例**

```
{
	out_trade_no:T1201807271716197099
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| out_trade_no | String | ✅ |  订单号| |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "amount": 1,
      "eff_date": "2018-10-01T00:00:00Z",
      "name": "当归",
      "packing_unit_name": "g",
      "record_id": 14,
      "ret_price": 10000,
      "serial": "1",
      "specification": "/kg"
    },
    {
      "amount": -1,
      "eff_date": "2018-10-01T00:00:00Z",
      "name": "当归",
      "packing_unit_name": "g",
      "record_id": 15,
      "ret_price": 10000,
      "serial": "1",
      "specification": "/kg"
    }
  ],
  "msg": "ok",
  "payrecordMap": {
    "auth_code": null,
    "balance_money": 8000,
    "created_time": "2018-07-27T17:16:19.624367+08:00",
    "discount_money": 1000,
    "id": 13,
    "medical_money": 1000,
    "operation_id": 1,
    "out_trade_no": "T1201807271716197099",
    "pay_method": "cash",
    "pay_time": "2018-07-27T17:16:19.715943+08:00",
    "record_sn": "O1201807271716087018",
    "status": 2,
    "total_money": 10000,
    "trade_no": ""
  },
  "refundMap": [
    {
      "created_time": "2018-07-27T17:17:46.977639+08:00",
      "id": 6,
      "operation_id": 1,
      "out_refund_no": "R1201807271717478953",
      "out_trade_no": "T1201807271716197099",
      "refund_money": -8000,
      "refund_no": "",
      "refund_time": "2018-07-27T17:17:46.977639+08:00",
      "status": 2
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ✅ |  零售详细信息 | |
| data.amount | int | ✅ |  数量 | |
| data.eff_date | Date | ✅ | 有效期 | |
| data.name | String | ✅ |  药品名称| |
| data.packing_unit_name | String | ✅ |  药品单位 | |
| data.record_id | int | ✅ |  零售记录id | |
| data.ret_price | int | ✅ |  零售价 | |
| data.serial | String | ✅ |  药品批号 | |
| data.specification | String | ✅ |  规格| |
| payrecordMap | Object | ✅ |  缴费交易信息 | |
| payrecordMap.auth_code | String | ✅ |  认证码 | |
| payrecordMap.balance_money | Int | ✅ |实缴金额 | |
| payrecordMap.created_time| time | ✅ |  创建时间| |
| payrecordMap.discount_money | Int | ✅ |  折扣金额 | |
| payrecordMap.id | Int | ✅ | 零售缴费记录id | |
| payrecordMap.medical_money | int | ✅ | 医保金额 | |
| payrecordMap.operation_id | int | ✅ |  收费员id | |
| payrecordMap.out_trade_no | String | ✅ |  系统交易号| |
| payrecordMap.pay_method | String | ✅ |  支付方式 | |
| payrecordMap.pay_time | time | ✅ |  支付时间 | |
| payrecordMap.record_sn | String | ✅ | 订单号 | |
| payrecordMap.status | boolean | ✅ |  订单状态, -1--待确认，1-待支付， 2--已支付， 3-支付失败| |
| payrecordMap.total_money | int | ✅ |  总金额 | |
| payrecordMap.trade_no | String | ❌ |  第三方交易号| |
| refundMap | Object | ✅ |  退费交易信息 | |
| refundMap.created_time | time | ✅ |  创建时间 | |
| refundMap.id| Int | ✅ | 退费记录id | |
| refundMap.operation_id| Int | ✅ |  退费人员id| |
| refundMap.out_refund_no | String | ✅ |  系统退费交易号 | |
| refundMap.out_trade_no | String | ✅ |  交易号，一次交易的编码 | |
| refundMap.refund_money | int | ✅ |  交易总金额金额 | |
| refundMap.refund_no | String | ✅ |  第三方退费交易号 | |
| refundMap.refund_time | time | ❌ | 退费时间| |
| refundMap.status | boolean | ✅ |  订单状态, -1--待确认， 2--退费成功， 3-退费失败 | |
--

</br>
<h3>28.6 退费

```
请求地址：/drugRetail/refund
```
**请求包示例**

```
{
	out_trade_no:T1201807271716197099
	items:[
		{"retail_id":"1","amount":"2"},
		{"retail_id":"1","amount":"2"}
	]
	operation_id:1
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| out_trade_no | String | ✅ |  系统订单号| |
| items | Array | ✅ |  零售药品信息| |
| items.retail_id | String | ✅ |  零售记录id | |
| items.amount | String | ✅ |  数量 | |
| operation_id | Int | ✅ |  退费人员id | |

**应答包示例**

```
{
    "code": "200",
    "msg":"ok"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ❌ |  返回信息 | |
--


29 财务分析、运营分析模块
--------

</br>
<h3>29.1 收费日报表按支付方式

```
请求地址：/financialAnalysis/ChargeDayReportByPayWay
```
**请求包示例**

```
{
	clinic_id:1
	start_date:2018-07-01
	end_date:2018-08-05
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ❌ |  诊所id| 
| start_date | String | ✅ | 开始日期| |
| end_date | String | ✅ |  结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "alipay": 0,
      "balance_money": 30895921,
      "bank": 0,
      "bonus_points_money": 0,
      "cash": 30895919,
      "clinic_id": 1,
      "clinic_name": "龙华诊所",
      "derate_money": 0,
      "discount_money": 0,
      "medical_money": 0,
      "on_credit_money": 0,
      "total_money": 30895921,
      "voucher_money": 0,
      "wechat": 2
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.alipay | Int | ✅ |  支付宝当面付| |
| data.items.balance_money | Int | ✅ |  实收金额| |
| data.items.bank | Int | ✅ |  银行卡| |
| data.items.bonus_points_money | Int | ✅ | 积分抵扣| |
| data.items.cash| Int | ✅ | 现金| |
| data.items.clinic_id | Int | ✅ | 诊所id| |
| data.items.clinic_name | Int | ✅ | 诊所名称| |
| data.items.derate_money| Int | ✅ | 减免| |
| data.items.discount_money | Int | ✅ |折后金额| |
| data.items.medical_money | Int | ✅ | 医保支付| |
| data.items.on_credit_money | Int | ✅ | 挂账| |
| data.items.total_money | Int | ✅ | 费用合计| |
| data.items.voucher_money | Int | ✅ | 卡卷| |
| data.items.wechat | Int | ✅ | 微信当面付| |
--

</br>
<h3>29.2 收费日报表按业务类型

```
请求地址：/financialAnalysis/ChargeDayReportByBusiness
```
**请求包示例**

```
{
	clinic_id:1
	start_date:2018-07-01
	end_date:2018-08-05
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ❌ |  诊所id| 
| start_date | String | ✅ | 开始日期| |
| end_date | String | ✅ |  结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_id": 1,
      "clinic_name": "龙华诊所",
      "diagnosis_treatment_fee": 16000,
      "discount_money": 0,
      "examination_fee": 577000,
      "labortory_fee": 384500,
      "material_fee": 8022003,
      "other_fee": 8003,
      "retail_fee": 14404,
      "total_money": 30895921,
      "traditional_medical_fee": 21425000,
      "treatment_fee": 97000,
      "western_medicine_fee": 352011
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.clinic_id | Int | ✅ | 诊所id| |
| data.items.clinic_name | Int | ✅ | 诊所名称| |
| data.items.diagnosis_treatment_fee| Int | ✅ | 挂号费| |
| data.items.discount_money | Int | ✅ |折后金额| |
| data.items.examination_fee | Int | ✅ | 检查费| |
| data.items.labortory_fee | Int | ✅ |检验费| |
| data.items.material_fee | Int | ✅ | 材料费| |
| data.items.other_fee | Int | ✅ | 其他费用| |
| data.items.retail_fee | Int | ✅ | 药品零售| |
| data.items.total_money | Int | ✅ | 费用合计| |
| data.items.traditional_medical_fee | Int | ✅ | 中药费| |
| data.items.treatment_fee | Int | ✅ | 治疗费| |
| data.items.western_medicine_fee | Int | ✅ | 西药费| |
--

</br>
<h3>29.3 收费月报表按支付方式

```
请求地址：/financialAnalysis/ChargeMonthReportByPayWay
```
**请求包示例**

```
{
	start_date:2018-07
	end_date:2018-08
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| start_date | String | ✅ | 开始日期| |
| end_date | String | ✅ |  结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "alipay": 0,
      "balance_money": -14408,
      "bank": 0,
      "bonus_points_money": 0,
      "business_month": "2018-08",
      "cash": -14408,
      "clinic_id": 2,
      "clinic_name": "最牛诊所",
      "derate_money": 0,
      "discount_money": 0,
      "medical_money": 0,
      "on_credit_money": 0,
      "total_money": -14408,
      "voucher_money": 0,
      "wechat": 0
    },
    {
      "alipay": 0,
      "balance_money": 12746003,
      "bank": 0,
      "bonus_points_money": 0,
      "business_month": "2018-07",
      "cash": 12746001,
      "clinic_id": 1,
      "clinic_name": "龙华诊所",
      "derate_money": 0,
      "discount_money": 0,
      "medical_money": 0,
      "on_credit_money": 0,
      "total_money": 12746003,
      "voucher_money": 0,
      "wechat": 2
    },
    {
      "alipay": 0,
      "balance_money": 22316663,
      "bank": 0,
      "bonus_points_money": 0,
      "business_month": "2018-08",
      "cash": 22316663,
      "clinic_id": 1,
      "clinic_name": "龙华诊所",
      "derate_money": 0,
      "discount_money": 0,
      "medical_money": 0,
      "on_credit_money": 0,
      "total_money": 22316663,
      "voucher_money": 0,
      "wechat": 0
    },
	...
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.alipay | Int | ✅ |  支付宝当面付| |
| data.items.balance_money | Int | ✅ |  实收金额| |
| data.items.bank | Int | ✅ |  银行卡| |
| data.items.bonus_points_money | Int | ✅ | 积分抵扣| |
| data.items.cash| Int | ✅ | 现金| |
| data.items.clinic_id | Int | ✅ | 诊所id| |
| data.items.clinic_name | Int | ✅ | 诊所名称| |
| data.items.derate_money| Int | ✅ | 减免| |
| data.items.discount_money | Int | ✅ |折后金额| |
| data.items.medical_money | Int | ✅ | 医保支付| |
| data.items.on_credit_money | Int | ✅ | 挂账| |
| data.items.total_money | Int | ✅ | 费用合计| |
| data.items.voucher_money | Int | ✅ | 卡卷| |
| data.items.wechat | Int | ✅ | 微信当面付| |
--

</br>
<h3>29.4 利润报表

```
请求地址：/financialAnalysis/ProfitReport
```
**请求包示例**

```
{
	clinic_id:
	start_date:2018-07-01
	end_date:2018-08-05
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ❌  | 开始日期| |
| start_date | String | ✅ | 开始日期| |
| end_date | String | ✅ |  结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": {
    "total_diagnosis_treatment_cost": 0,
    "total_diagnosis_treatment_fee": 17500,
    "total_examination_cost": 10000,
    "total_examination_fee": 577000,
    "total_labortory_cost": 200,
    "total_labortory_fee": 384500,
    "total_material_cost": 4200,
    "total_material_fee": 9022103,
    "total_other_cost": 2400,
    "total_other_fee": 10003,
    "total_retail_cost": 102,
    "total_retail_fee": 14404,
    "total_traditional_medical_cost": 900000,
    "total_traditional_medical_fee": 24375000,
    "total_treatment_cost": 3000,
    "total_treatment_fee": 124000,
    "total_western_medicine_cost": 79,
    "total_western_medicine_fee": 408011
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Object | ✅ |  返回信息 | |
| data.total_diagnosis_treatment_cost | Int | ✅ |  诊疗费成本价| |
| data.total_diagnosis_treatment_fee | Int | ✅ |  诊疗费| |
| data.total_examination_cost | Int | ✅ |  检查费成本价| |
| data.total_examination_fee | Int | ✅ | 检查费| |
| data.total_labortory_cost| Int | ✅ | 检验费成本价| |
| data.total_labortory_fee | Int | ✅ | 检验费| |
| data.total_material_cost | Int | ✅ | 材料费成本价| |
| data.total_material_fee| Int | ✅ | 材料费| |
| data.total_other_cost | Int | ✅ |其他费用成本价| |
| data.total_other_fee | Int | ✅ | 其他费用| |
| data.total_retail_cost | Int | ✅ | 零售费成本价| |
| data.total_retail_fee | Int | ✅ | 零售费用| |
| data.total_traditional_medical_cost | Int | ✅ | 中药费成本价| |
| data.total_traditional_medical_fee | Int | ✅ | 中药费| |
| data.total_treatment_cost | Int | ✅ | 治疗费成本价| |
| data.total_treatment_fee | Int | ✅ | 治疗费| |
| data.total_western_medicine_cost | Int | ✅ |西/成药费成本价| |
| data.total_western_medicine_fee | Int | ✅ | 西/成药费| |
--

30 平台交易总额模块
--------
</br>
<h3>30.1 交易总额

```
请求地址：/invoicingStatistics/DrugInvoicingStatistics
```
**请求包示例**

```
{
	start_date:2018-07-01
	end_date:2018-08-05
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| start_date | String | ✅  | 统计开始日期| |
| end_date | String | ✅  |  统计结束日期 | |

**应答包示例**

```
{
  "clinic": [
    {
      "id": 1,
      "name": "龙华诊所"
    },
    {
      "id": 2,
      "name": "最牛诊所"
    }
  ],
  "code": "200",
  "data": [
    {
      "balance_money": 129000,
      "clinic_id": 1,
      "created_time": "2018-07-26",
      "name": "龙华诊所"
    },
    {
      "balance_money": 4036600,
      "clinic_id": 2,
      "created_time": "2018-07-26",
      "name": "最牛诊所"
    },
    {
      "balance_money": 220000,
      "clinic_id": 1,
      "created_time": "2018-07-27",
      "name": "龙华诊所"
    },
    {
      "balance_money": 12108000,
      "clinic_id": 1,
      "created_time": "2018-07-28",
      "name": "龙华诊所"
    },
	...
  ],
  "total": {
    "total_money": 20845608
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ✅ |  返回信息 | |
| data.items.balance_money | FLOAT | ❌ | 交易金额| |
| data.items.clinic_id | Date | ✅ |  诊所id| |
| data.items.created_time | Date | ✅ |  创建日期| |
| data.items.name | String | ✅ | 诊所名称| |
| clinic | Array | ✅ |  诊所信息| |
| clinic.items.id | Int | ✅ |  诊所id| |
| clinic.items.name | String |✅ | 诊所名称| |
| total | Object | ✅ |  合计信息| |
| total.total_money | FLOAT | ✅ |  合计金额| |

--

31 医用报表模块
--------

</br>
<h3>31.1 门诊日志

```
请求地址：/medicalReport/outPatient/record
```
**请求包示例**

```
{
	clinic_id:1
	patient_name:
	phone:
	doctor_id:
	operation_id:
	start_date:2018-05-01
	end_date:2018-08-13
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| patient_name | String | ❌ | 患者姓名| |
| phone | String | ❌ | 手机号码| |
| doctor_id | Int | ❌ | 接诊医生id| |
| operation_id | Int | ❌ | 登记人员id| |
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "address": "芍药居北里",
      "birthday": "19920706",
      "city": "北京市",
      "dept_name": "骨科",
      "diagnosis": "耳热是从v",
      "district": "东城区",
      "doctor_name": "扁鹊",
      "morbidity_date": "2018-07-21",
      "opreation_name": "超级管理员",
      "patient_id": 3,
      "patient_name": "查康",
      "phone": "18701676735",
      "profession": "22233",
      "province": "北京市",
      "sex": 1,
      "visit_date": "2018-05-28T00:00:00Z",
      "visit_type": 1
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "person_amount": 25,
    "total": 82
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.address | String | ❌ |  地址| |
| data.items.birthday | String | ✅ | 生日| |
| data.items.city | String | ❌ |  市| |
| data.items.dept_name| String | ✅ |  接诊科室名称| |
| data.items.diagnosis | String | ✅ |  诊断| |
| data.items.district | String | ❌ |  区、县| |
| data.items.doctor_name | String | ✅ |  接诊医生名称| |
| data.items.morbidity_date | String | ✅ | 发病日期| |
| data.items.opreation_name | String | ✅ |  登记人员| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.patient_name | String | ✅ |  患者名称| |
| data.items.phone | String | ✅ |手机号| |
| data.items.profession | String | ❌ |  患者职业| |
| data.items.province | String | ❌ |  省| |
| data.items.sex | Int | ✅ |  性别 0 女 1 男| |
| data.items.visit_date | Date | ✅ |  就诊日期| |
| data.items.visit_type | Int | ✅ | 出诊类型 1: 首诊， 2复诊，3：术后复诊| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.person_amount | Int | ✅ |  总接诊人数| |
--

</br>
<h3>31.2 接诊类型统计

```
请求地址：/medicalReport/outPatient/type
```
**请求包示例**

```
{
	clinic_id:1
	start_date:2018-05-01
	end_date:2018-08-13
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| 
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "created_time": "2018-05-28",
      "type1": 2,
      "type2": 0,
      "type3": 0
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 34
  },
  "total": {
    "type1": 42,
    "type2": 37,
    "type3": 0
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.created_time | Date | ✅ |  接诊日期| |
| data.items.type1 | Int | ✅ |  首诊类型人数| |
| data.items.type2 | Int | ✅ |  复诊类型人数| |
| data.items.type3 | Int | ✅ |  术后复诊人数| |
| data.total | Object |  ✅ |  合计数量 | |
| data.total.type1 | Int | ✅ |  总首诊类型人数| |
| data.total.type2 | Int | ✅ |  总复诊类型人数| |
| data.total.type3 | Int | ✅ |  总术后复诊人数| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>31.3 科室统计

```
请求地址：/medicalReport/outPatient/department
```
**请求包示例**

```
{
	clinic_id:1
	start_date:2018-05-01
	end_date:2018-08-13
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| 
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "balance_money": null,
      "date": "2018-05-28",
      "department_id": 1,
      "total": 3
    },
	...
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.balance_money | FLOAT | ✅ |  接诊金额| |
| data.items.date | Date | ✅ |  接诊日期| |
| data.items.department_id | Int | ✅ |  接诊科室id| |
| data.items.total | Int | ✅ | 接诊人次| |
--

</br>
<h3>31.4 医生接诊统计

```
请求地址：/medicalReport/ReceiveTreatment
```
**请求包示例**

```
{
	clinic_id:1
	department_id:1
	start_date:2018-05-01
	end_date:2018-08-13
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| department_id | Int | ❌ | 科室id| |
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "department_name": "骨科",
      "diagnosis_treatment_fee": 47500,
      "east_pre_fee": 2045000,
      "exam_pre_fee": 242000,
      "labora_pre_fee": 2082500,
      "material_fee": 1206000,
      "other_fee": 50000,
      "personnel_name": "扁鹊",
      "tatol_count": 92,
      "total_fee": 5812171,
      "treatement_fee": 105000,
      "west_pre_fee": 34171
    },
    {
      "department_name": "骨科",
      "diagnosis_treatment_fee": 3000,
      "east_pre_fee": 5400000,
      "exam_pre_fee": 3000,
      "labora_pre_fee": 32500,
      "material_fee": 2000,
      "other_fee": 0,
      "personnel_name": "黄飞鸿",
      "tatol_count": 28,
      "total_fee": 5468040,
      "treatement_fee": 20000,
      "west_pre_fee": 7540
    }
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.department_name | String |✅ |  科室名称| |
| data.items.diagnosis_treatment_fee | Float | ✅ | 诊断费用| |
| data.items.east_pre_fee | Float | ✅ |  中药费用| |
| data.items.exam_pre_fee| Float | ✅ |  检查费用| |
| data.items.labora_pre_fee | Float | ✅ |  检验费用| |
| data.items.material_fee | Float | ✅ |  材料费用| |
| data.items.other_fee | Float | ✅ |  其它费用| |
| data.items.personnel_name | String | ✅ | 接诊医生| |
| data.items.tatol_count | Int | ✅ |  总人数| |
| data.items.total_fee | Int | ✅ |  总金额| |
| data.items.treatement_fee | Float | ✅ |  治疗费用| |
| data.items.west_pre_fee | Float | ✅ |西药费用| |
--

</br>
<h3>31.5 检查统计

```
请求地址：/medicalReport/ExaminationStatistics
```
**请求包示例**

```
{
	clinic_id:1
	start_date:2018-06-01
	end_date:2018-08-07
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "20000603",
      "clinic_examination_name": "盆腔彩色多普勒超声",
      "created_time": "2018-06-27T23:55:50.468246+08:00",
      "id": 20,
      "name": "倾世容颜",
      "personnel_name": "超级管理员",
      "phone": "15387556262",
      "price": 3000,
      "sex": 0,
      "times": 1,
      "updated_time": "2018-07-26T09:55:31.136587+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "patient_count": 5,
    "tiems_total": 18,
    "total": 18
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.birthday | String |  ✅ |  生日| |
| data.items.clinic_examination_name | String | ✅ | 检查名称| |
| data.items.created_time | time | ✅ |  创建时间| |
| data.items.id| Int | ✅ |  患者id| |
| data.items.name | String | ✅ |  患者名称| |
| data.items.personnel_name | String | ✅ |  检查医生| |
| data.items.phone | String | ✅ |  患者手机号| |
| data.items.price | String | ✅ | 单价| |
| data.items.sex | Int | ✅ |  性别 0 女 1 男| |
| data.items.times | Int | ✅ |  检查次数| |
| data.items.updated_time | time | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.patient_count | Int | ✅ |  总患者数| |
| data.page_info.tiems_total | Int | ✅ |  总次数| |
--

</br>
<h3>31.6 检验统计

```
请求地址：/medicalReport/LaboratoryStatistics
```
**请求包示例**

```
{
	clinic_id:1
	start_date:2018-06-02
	end_date:2018-08-10
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "20000603",
      "clinic_laboratory_name": "血常规",
      "created_time": "2018-06-27T23:55:14.935973+08:00",
      "id": 20,
      "name": "倾世容颜",
      "personnel_name": "超级管理员",
      "phone": "15387556262",
      "price": 40000,
      "sex": 0,
      "times": 1,
      "updated_time": "2018-06-28T00:04:21.484701+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "patient_count": 5,
    "tiems_total": 18,
    "total": 18
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.birthday | String |  ✅ |  生日| |
| data.items.clinic_laboratory_name | String | ✅ | 检验名称| |
| data.items.created_time | time | ✅ |  创建时间| |
| data.items.id| Int | ✅ |  患者id| |
| data.items.name | String | ✅ |  患者名称| |
| data.items.personnel_name | String | ✅ |  检查医生| |
| data.items.phone | String | ✅ |  患者手机号| |
| data.items.price | String | ✅ | 单价| |
| data.items.sex | Int | ✅ |  性别 0 女 1 男| |
| data.items.times | Int | ✅ |  检查次数| |
| data.items.updated_time | time | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.patient_count | Int | ✅ |  总患者数| |
| data.page_info.tiems_total | Int | ✅ |  总次数| |
--

</br>
<h3>31.7 治疗统计

```
请求地址：/medicalReport/TreatmentStatistics
```
**请求包示例**

```
{
	clinic_id:1
	start_date:2018-06-02
	end_date:2018-08-10
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "birthday": "20000603",
      "clinic_treatment_name": "针灸",
      "created_time": "2018-06-27T23:54:43.93167+08:00",
      "id": 20,
      "name": "倾世容颜",
      "personnel_name": "超级管理员",
      "phone": "15387556262",
      "price": 9000,
      "sex": 0,
      "times": 3,
      "updated_time": "2018-06-29T01:03:54.733934+08:00"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "patient_count": 5,
    "tiems_total": 18,
    "total": 18
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.birthday | String |  ✅ |  生日| |
| data.items.clinic_treatment_name | String | ✅ | 治疗名称| |
| data.items.created_time | time | ✅ |  创建时间| |
| data.items.id| Int | ✅ |  患者id| |
| data.items.name | String | ✅ |  患者名称| |
| data.items.personnel_name | String | ✅ |  检查医生| |
| data.items.phone | String | ✅ |  患者手机号| |
| data.items.price | String | ✅ | 单价| |
| data.items.sex | Int | ✅ |  性别 0 女 1 男| |
| data.items.times | Int | ✅ |  检查次数| |
| data.items.updated_time | time | ✅ |  更新时间| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.patient_count | Int | ✅ |  总患者数| |
| data.page_info.tiems_total | Int | ✅ |  总次数| |
--

</br>
<h3>31.8 登记统计

```
请求地址：/medicalReport/RegisterStatistics
```
**请求包示例**

```
{
	clinic_id:1
	start_date:2018-05-01
	end_date:2018-08-13
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "appointment_count": 0,
      "register_count": 1,
      "total_count": 1,
      "visit_date": "2018-06-02"
    },
    ...
  ]
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.appointment_count | Int | ✅ |  预约就诊人数| |
| data.items.register_count | Int | ✅ |  登记总数| |
| data.items.total_count | Int | ✅ |  合计| |
| data.items.visit_date| Date | ✅ |接诊时间| |
--

</br>
<h3>31.9 门诊效率统计

```
请求地址：/medicalReport/OutPatietnEfficiencyStatistics
```
**请求包示例**

```
{
	clinic_id:1
	department_id:
	start_date:2018-08-01
	end_date:2018-08-05
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所id| |
| department_id | Int | ✅ |  科室id| |
| start_date | String | ❌ | 就诊开始日期| |
| end_date | String | ❌ |  就诊结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "clinic_triage_patient_id": 122,
      "department_id": 6,
      "department_name": "牙科",
      "finish_time": "2018-08-01T10:13:03.995112+08:00",
      "patient_id": 29,
      "patient_name": "赵丽颖",
      "pay_time": "2018-08-01T10:22:36.615245+08:00",
      "reception_time": "2018-08-01T09:59:47.010945+08:00",
      "register_time": "2018-08-01T09:35:43.217418+08:00",
      "status": 40,
      "triage_time": "2018-08-01T09:48:05.457607+08:00",
      "visit_date": "2018-08-01T00:00:00Z"
    },
	...
  ],
  "page_info": {
    "average_pay_finished_time": 121.6,
    "average_reception_finished_time": 6561.9,
    "average_receptioned_time": 34.2,
    "average_triage_finished_time": 11.7,
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
| data | Array | ❌ |  返回信息 | |
| data.items.clinic_triage_patient_id | Int |  ✅ |  就诊id| |
| data.items.department_id | String | ✅ | 科室id| |
| data.items.department_name | time | ✅ |  科室名称| |
| data.items.finish_time| time | ✅ |  接诊结束时间| |
| data.items.patient_id | Int | ✅ |  患者id| |
| data.items.patient_name | String | ✅ |  患者名称| |
| data.items.pay_time | time | ✅ |  收费完成时间| |
| data.items.reception_time | time | ✅ | 接诊开始时间| |
| data.items.register_time | time | ✅ |  登记完成时间| |
| data.items.status | Int | ✅ | 操作类型 10:登记，20：分诊(换诊)，30：接诊，40：已就诊， 100：取消| |
| data.items.triage_time | time | ✅ |  分诊完成时间| |
| data.items.visit_date | Date | ✅ |  接诊日期| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.average_pay_finished_time | Float | ✅ | 平均收费用时(分钟)| |
| data.page_info.average_reception_finished_time | Float | ✅ | 平均接诊用时(分钟)| |
| data.page_info.average_receptioned_time | Float | ✅ | 平均安排接诊用时(分钟)| |
| data.page_info.average_triage_finished_time | Float | ✅ | 平均分诊用时(分钟)| |
--

32 进销存统计模块
--------

</br>
<h3>32.1 药品入库统计

```
请求地址：/invoicingStatistics/DrugInstockStatistics
```
**请求包示例**

```
{
	clinic_id:1
	instock_way_name:零售退药
	supplier_name:
	drug_type:
	drug_name:
	serial:
	start_date:2018-08-01
	end_date:2018-08-12
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| instock_way_name | String | ❌ | 入库方式| |
| supplier_name | String | ❌ | 供应商| |
| drug_type | Int | ❌ | 药品类型 0-西药 1-中药| |
| drug_name | String | ❌ | 药品名称、条形码| |
| serial | String | ❌ | 入库批号| |
| start_date | String | ❌ | 入库开始日期| |
| end_date | String | ❌ |  入库结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "barcode": null,
      "buy_price": 1,
      "drug_name": "测试药品",
      "eff_date": "2018-12-29T00:00:00Z",
      "instock_amount": 2,
      "instock_date": "2018-08-06",
      "instock_operation_name": "超级管理员",
      "instock_way_name": "零售退药",
      "manu_factory_name": "测试生产厂商",
      "order_number": "T1201808062117012284",
      "packing_unit_name": "盒",
      "ret_price": 1,
      "serial": "2018080201",
      "specification": "10粒/每盒",
      "supplier_name": "广州白云药厂"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 8,
    "total_buy_price": 12109,
    "total_instock_amount": 23
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.barcode | String | ❌ |  药品条形码| |
| data.items.buy_price | FLOAT | ❌ | 成本价| |
| data.items.drug_name | String | ✅ |  药品名称| |
| data.items.eff_date | Date | ✅ |  有效期| |
| data.items.instock_amount | Int | ✅ |  入库数量| |
| data.items.instock_date | Date | ✅ |  入库日期| |
| data.items.instock_operation_name | String | ✅ |  入库操作员名称| |
| data.items.instock_way_name | String | ✅ |  入库方式| |
| data.items.manu_factory_name | String | ✅ |  生成产商| |
| data.items.order_number | String | ✅ |  入库单号| |
| data.items.packing_unit_name | String | ✅ |  药品包装单位| |
| data.items.ret_price | FLOAT | ✅ |  零售价| |
| data.items.serial | String | ✅ |  批号| |
| data.items.specification | String | ✅ |  规格| |
| data.items.supplier_name | String | ✅ |  供应商| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.total_buy_price | FLOAT | ✅ |  总成本价| |
| data.page_info.total_instock_amount | Int | ✅ |  总入库数量| |
--

</br>
<h3>32.2 药品出库统计

```
请求地址：/invoicingStatistics/DrugOutstockStatistics
```
**请求包示例**

```
{
	clinic_id:1
	outstock_way_name:报损出库
	supplier_name:
	drug_type:
	drug_name:
	serial:
	personnel_name:
	start_date:2018-08-01
	end_date:2018-08-08
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| outstock_way_name | String | ❌ | 出库方式| |
| supplier_name | String | ❌ | 供应商| |
| drug_type | Int | ❌ | 药品类型 0-西药 1-中药| |
| drug_name | String | ❌ | 药品名称、条形码| |
| serial | String | ❌ | 入库批号| |
| personnel_name | String | ❌ | 领用人员| |
| start_date | String | ❌ | 出库开始日期| |
| end_date | String | ❌ |  出库结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "barcode": "73451412136373134",
      "buy_price": 50,
      "drug_name": "川贝母",
      "eff_date": "2018-08-01T00:00:00Z",
      "manu_factory_name": "北京鹤延龄饮片厂",
      "order_number": "DCKD-1533107589",
      "outstock_amount": 100,
      "outstock_date": "2018-08-01",
      "outstock_operation_name": "超级管理员",
      "outstock_way_name": "报损出库",
      "packing_unit_name": "g",
      "personnel_name": "查康",
      "ret_price": 1000,
      "serial": "20180801002",
      "specification": "H01019/kg",
      "supplier_name": "云南白药药厂"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 23,
    "total_buy_price": 203830995,
    "total_outstock_amount": 14976
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.barcode | String | ❌ |  药品条形码| |
| data.items.buy_price | FLOAT | ❌ | 成本价| |
| data.items.drug_name | String | ✅ |  药品名称| |
| data.items.eff_date | Date | ✅ |  有效期| |
| data.items.outstock_amount | Int | ✅ |  出库数量| |
| data.items.outstock_date | Date | ✅ |  出库日期| |
| data.items.outstock_operation_name | String | ✅ |  出库操作员名称| |
| data.items.outstock_way_name | String | ✅ |  出库方式| |
| data.items.manu_factory_name | String | ✅ |  生成产商| |
| data.items.order_number | String | ✅ |  出库单号| |
| data.items.packing_unit_name | String | ✅ |  药品包装单位| |
| data.items.personnel_name | String | ✅ |  领用人员| |
| data.items.ret_price | FLOAT | ✅ |  零售价| |
| data.items.serial | String | ✅ |  批号| |
| data.items.specification | String | ✅ |  规格| |
| data.items.supplier_name | String | ✅ |  供应商| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.total_buy_price | FLOAT | ✅ |  总成本价| |
| data.page_info.total_outstock_amount | Int | ✅ |  总出库数量| |
--

</br>
<h3>32.3 药品进存销统计

```
请求地址：/invoicingStatistics/DrugInvoicingStatistics
```
**请求包示例**

```
{
	clinic_id:1
	supplier_name:广州白云药厂
	drug_type:
	drug_name:
	serial:
	start_date:2018-05-01
	end_date:2018-08-08
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| supplier_name | String | ❌ | 供应商| |
| drug_type | Int | ❌ | 药品类型 0-西药 1-中药| |
| drug_name | String | ❌ | 药品名称、条形码| |
| serial | String | ❌ | 入库批号| |
| start_date | String | ❌ | 统计开始日期| |
| end_date | String | ❌ |  统计结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "barcode": null,
      "buy_price": 1,
      "eff_date": "2018-12-29T00:00:00Z",
      "id": 35,
      "manu_factory_name": "测试生产厂商",
      "name": "测试药品",
      "packing_unit_name": "盒",
      "serial": "2018080201",
      "specification": "10粒/每盒",
      "stock_amount": 98,
      "supplier_name": "广州白云药厂",
      "total_instock_amount": 104,
      "total_outstock_amount": -2
    },
    ...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 29
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.barcode | String | ❌ |  药品条形码| |
| data.items.buy_price | FLOAT | ❌ | 成本价| |
| data.items.eff_date | Date | ✅ |  有效期| |
| data.items.id | Int | ✅ |  药品库存id| |
| data.items.manu_factory_name | String | ✅ |  生成产商| |
| data.items.name | String | ✅ |  药品名称| |
| data.items.packing_unit_name | String | ✅ |  药品包装单位| |
| data.items.serial | String | ✅ |  批号| |
| data.items.specification | String | ✅ |  规格| |
| data.items.stock_amount | Int | ✅ |  库存数量| |
| data.items.supplier_name | String | ✅ |  供应商| |
| data.items.total_instock_amount | Int | ✅ |  入库数量| |
| data.items.total_outstock_amount | Int | ✅ |  出库数量| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--

</br>
<h3>32.4 耗材入库统计

```
请求地址：/invoicingStatistics/MaterialInstockStatistics
```
**请求包示例**

```
{
	clinic_id:1
	instock_way_name:
	supplier_name:
	material_name:
	serial:
	start_date:2018-08-01
	end_date:2018-08-13
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| instock_way_name | String | ❌ | 入库方式| |
| supplier_name | String | ❌ | 供应商| |
| material_name | String | ❌ | 耗材名称| |
| serial | String | ❌ | 入库批号| |
| start_date | String | ❌ | 入库开始日期| |
| end_date | String | ❌ |  入库结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "buy_price": 100,
      "eff_date": "2019-12-31T00:00:00Z",
      "idc_code": "4545415",
      "instock_amount": -1,
      "instock_date": "2018-08-13",
      "instock_operation_name": "超级管理员",
      "instock_way_name": "门诊退费",
      "manu_factory_name": "生产厂家1",
      "material_name": "假牙",
      "order_number": "MZ20180813011334",
      "ret_price": 1,
      "serial": "PH4165415212",
      "specification": "20g/颗",
      "supplier_name": "广州白云药厂",
      "unit_name": "粒"
    },
	...
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 2,
    "total_buy_price": 9900,
    "total_instock_amount": 99
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.buy_price | FLOAT | ❌ | 成本价| |
| data.items.eff_date | Date | ✅ |  有效期| |
| data.items.instock_amount | Int | ✅ |  入库数量| |
| data.items.instock_date | Date | ✅ |  入库日期| |
| data.items.instock_operation_name | String | ✅ |  入库操作员名称| |
| data.items.instock_way_name | String | ✅ |  入库方式| |
| data.items.manu_factory_name | String | ✅ |  生成产商| |
| data.items.material_name | String | ✅ |  耗材名称| |
| data.items.order_number | String | ✅ |  入库单号| |
| data.items.ret_price | FLOAT | ✅ |  零售价| |
| data.items.serial | String | ✅ |  批号| |
| data.items.specification | String | ✅ |  规格| |
| data.items.supplier_name | String | ✅ |  供应商| |
| data.items.unit_name | String | ✅ |  耗材单位| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.total_buy_price | FLOAT | ✅ |  总成本价| |
| data.page_info.total_instock_amount | Int | ✅ |  总入库数量| |
--

</br>
<h3>32.5 耗材出库统计

```
请求地址：/invoicingStatistics/MaterialOutstockStatistics
```
**请求包示例**

```
{
	clinic_id:1
	outstock_way_name:
	supplier_name:
	material_name:
	serial:
	personnel_name:
	start_date:2018-08-01
	end_date:2018-08-08
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| outstock_way_name | String | ❌ | 出库方式| |
| supplier_name | String | ❌ | 供应商| |
| material_name | String | ❌ | 耗材名称| |
| serial | String | ❌ | 入库批号| |
| personnel_name | String | ❌ | 领用人员| |
| start_date | String | ❌ | 出库开始日期| |
| end_date | String | ❌ |  出库结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "buy_price": 100,
      "eff_date": "2019-12-31T00:00:00Z",
      "idc_code": "4545415",
      "manu_factory_name": "生产厂家1",
      "material_name": "假牙",
      "order_number": "MZ20180813011000",
      "outstock_amount": 1,
      "outstock_date": "2018-08-13",
      "outstock_operation_name": "超级管理员",
      "outstock_way_name": "门诊收费",
      "personnel_name": "查康",
      "ret_price": 1,
      "serial": "PH4165415212",
      "specification": "20g/颗",
      "supplier_name": "广州白云药厂",
      "unit_name": "粒"
    }
  ],
  "page_info": {
    "limit": "10",
    "offset": "0",
    "total": 1,
    "total_buy_price": 100,
    "total_outstock_amount": 1
  }
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| data | Array | ❌ |  返回信息 | |
| data.items.buy_price | FLOAT | ❌ | 成本价| |
| data.items.eff_date | Date | ✅ |  有效期| |
| data.items.idc_code | String | ❌ |  国际编码| |
| data.items.material_name | String | ✅ |  耗材名称| |
| data.items.outstock_amount | Int | ✅ |  出库数量| |
| data.items.outstock_date | Date | ✅ |  出库日期| |
| data.items.outstock_operation_name | String | ✅ |  出库操作员名称| |
| data.items.outstock_way_name | String | ✅ |  出库方式| |
| data.items.manu_factory_name | String | ✅ |  生成产商| |
| data.items.order_number | String | ✅ |  出库单号| |
| data.items.personnel_name | String | ✅ |  领用人员| |
| data.items.ret_price | FLOAT | ✅ |  零售价| |
| data.items.serial | String | ✅ |  批号| |
| data.items.specification | String | ✅ |  规格| |
| data.items.supplier_name | String | ✅ |  供应商| |
| data.items.unit_name | String | ✅ |  耗材单位| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
| data.page_info.total_buy_price | FLOAT | ✅ |  总成本价| |
| data.page_info.total_outstock_amount | Int | ✅ |  总出库数量| |
--

</br>
<h3>32.6 耗材进存销统计

```
请求地址：/invoicingStatistics/DrugInvoicingStatistics
```
**请求包示例**

```
{
	clinic_id:1
	supplier_name:广州白云药厂
	material_name:
	serial:
	start_date:2018-05-01
	end_date:2018-08-08
	offset:0
	limit:10
}
```
**请求包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| clinic_id | Int | ✅ |  诊所编码id| |
| supplier_name | String | ❌ | 供应商| |
| material_name | String | ❌ | 耗材名称| |
| serial | String | ❌ | 入库批号| |
| start_date | String | ❌ | 统计开始日期| |
| end_date | String | ❌ |  统计结束日期 | |
| offset | String | ❌ | 分页查询使用、跳过的数量 | |
| limit | String | ❌ | 分页查询使用、每页数量 | |

**应答包示例**

```
{
  "code": "200",
  "data": [
    {
      "buy_price": 100,
      "eff_date": "2018-07-19T00:00:00Z",
      "id": 6,
      "idc_code": "4545415",
      "manu_factory_name": "生产厂家1",
      "name": "假牙",
      "py_code": "jiaya",
      "serial": "ph20180719",
      "specification": "20g/颗",
      "stock_amount": 50,
      "supplier_name": "广州白云药厂",
      "total_instock_amount": 50,
      "total_outstock_amount": null,
      "unit_name": "粒"
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
| data | Array | ❌ |  返回信息 | |
| data.items.buy_price | FLOAT | ❌ | 成本价| |
| data.items.eff_date | Date | ✅ |  有效期| |
| data.items.id | Int | ✅ |  耗材库存id| |
| data.items.idc_code | String | ❌ |  国际编码| |
| data.items.manu_factory_name | String | ✅ |  生成产商| |
| data.items.name | String | ✅ |  耗材名称| |
| data.items.py_code | String | ✅ |  拼音编码| |
| data.items.serial | String | ✅ |  批号| |
| data.items.specification | String | ✅ |  规格| |
| data.items.stock_amount | Int | ✅ |  库存数量| |
| data.items.supplier_name | String | ✅ |  供应商| |
| data.items.unit_name | String | ✅ |  耗材单位| |
| data.items.total_instock_amount | Int | ✅ |  入库数量| |
| data.items.total_outstock_amount | Int | ✅ |  出库数量| |
| data.page_info | Object | ✅ |  返回的页码和总数| |
| data.page_info.offset | Int | ✅ |  分页使用、跳过的数量| |
| data.page_info.limit | Int | ✅ |  分页使用、每页数量| |
| data.page_info.total | Int | ✅ |  分页使用、总数量| |
--