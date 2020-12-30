# ginweb_blog
golang gin

功能一:

1、输入用户名、密码验证登录并返回token

2、新增用户、修改密码

功能二：使用gin框架实现操作数据库,产生含有json数据的接口并部署

接口内容包括：

1、查询学校列表接口（返回“清华大学”、“北京大学”、“浙江大学”），自己增加数据，实现分页功能

2、模糊搜索查询、分地区条件查询（如“北京”、“浙江”）

3、新增、删除学校接口（post请求）

启动:
1.启动前端
  安装node环境
  fe/hello 执行 npm install
  然后执行 npm run serve 
  本地8021端口
2.后端
  在database/mysql下
  修改数据库密码为本地密码
  启动项目即可

go mod创建项目
1.在项目文件夹里初始化项目
go mod init xxx

2.安装插件
例子
go get -u github.com/gogf/gf
