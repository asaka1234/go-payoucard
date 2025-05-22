文档
=============
https://github.com/lifeonearth718/payoucard-doc/blob/main/API-CN.md#%E9%93%B6%E8%A1%8C%E5%8D%A1%E5%85%85%E5%80%BC

鉴权
==============
1. rsa privateKey私钥加密算签名, rsa publicKey公钥解密验证签名
2. 对请求参数算了一个sign签名, 随后作为json里一个字段一起发

回调地址
==============
是提前让payouCard配置好的, 故而无法api中动态修改


Comment
===============
1. only support withdrawl
2. 所有接口都是 application/json 格式的

整体流程
=============
1调用withdraw接口来发起提现. 要对发送的Body里的data参数:排序后用privateKey算一个签名一起发过去
2回调也是有sign签名的,需要用publicKey解密来验证正确性