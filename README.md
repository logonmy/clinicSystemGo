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
| code | String | ✅ |  用户微信id| |
| name | String | ✅ |  手机号 | |
| responsible_person | String | ✅ |  手机号 | |
| area | String | ✅ |  手机号 | |
| province | String | ❌ |  头像地址（可以是 微信 头像地址） | |
| city | String | ❌ |  姓名（可以是 微信 昵称） | |
| district | String | ✅ |  用户微信id| |
| status | String | ✅ |  手机号 | |
| username | String | ❌ |  头像地址（可以是 微信 头像地址） | |
| password | String | ❌ |  姓名（可以是 微信 昵称） | |
| phone | String | ❌ |  姓名（可以是 微信 昵称） | |

**应答包示例**

```
{
    "code": "200",
    "msg": "注册成功"
}
```

**应答包参数说明**

| 参数名称 | 参数类型 | 是否必须 | 说明 | 默认值 |
| :-: | :-: | :-:  | :--: | :--: |
| code | String | ✅ |  返回码， 200 成功| |
| msg | String | ✅ |  返回信息 | |
--