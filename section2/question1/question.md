# 实现一个简单的加密和解密函数
实现一个int32的数字的加密和解密函数

实现原理:

加密步骤
第一步把数字的每位上的数字加一个随机数37187,第二步把第一步得到的3个数字分别模7，最后把第一位和第三位对调(实际操作是反转)，就得到了加密后的数字。

解密步骤
根据加密步骤反向操作就可以了
