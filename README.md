
古诗词项目来源于beego官方案例。难度系数四星（★★★★☆☆☆☆☆☆）。
-------

1.安装 go 语言环境  
-------
>安装方式建议参考官方，或者参考我的文档：https://blog.csdn.net/eighthroute/article/details/80318224  


2.安装 beego 框架 & bee 工具  
-------
>安装方式建议参考官方，或者参考我的文档：https://blog.csdn.net/eighthroute/article/details/109708605


3.初始化数据库<  
-------
>将根目录下 gushici.sql 导入到数据库中。  


4.启动项目  
-------
> ``
 bee run
``   


1).访问前台：http://localhost:8080/ 
![图片说明](https://img-blog.csdnimg.cn/20201116163755241.png) 

2).访问后台：http://localhost:8080/login （用户名/密码：admin/admin）  

![图片说明](https://img-blog.csdnimg.cn/20201116163831941.png)  
![图片说明](https://img-blog.csdnimg.cn/202011161638502.png)  

3).代码组织：
![图片说明](https://img-blog.csdnimg.cn/20201116163905630.png)  


说明  
-------
>1).该项目以纯的面向对象方式进行开发，主要涉及 go 基础语法 & beego 使用。<br/>
2).可能有人对封装的方式有质疑，欢迎一起交流，探讨出更好的封装方式。如果后期想到更好的封装方式，可能会对代码进行升级。<br/>
3).下一步将写一个简单 go 语言版的 MVC 框架。<br/>
4).如果学习过程中有疑问可以随时交流。<br/>

感谢【郝大全】同学提供相应资源。  

