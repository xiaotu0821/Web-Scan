欢迎使用初梦的小小工具箱。
     使用方式 
		源码运行方式： go run main.go
		编译后运行方式   main.exe
		参数说明：
			-h    查看帮助
			-alljson  使用所有json文件,默认是poc目录下的,可在config.json里修改目录（就是使用全部漏洞）
			-file   url文件路径
			-url    单个url测试时使用该参数
			-find   后面跟漏洞的关键字 ,比如   泛微   
			-json    find 找到的json路径（或者你自己手动查看的路径），就是使用哪个漏洞json文件
			-thread   线程数量，批量测试时使用，默认为1
			-gettitle   该参数需要和 -file 参数一起使用，用于批量获取url标题，获取的标题保存在同目录下的 url_title.txt（如打开乱码，可使用ue 或vscode打开，或者切换一下编码）
			
	常用搭配方式：
		一个url  一个  漏洞（以泛微举例）
			main.exe   -url  http://127.0.0.1 -json  poc\泛微OA漏洞\泛微-协同办公OA任意管理员登录.json
		一个url  所有漏洞  10个线程
			main.exe  -url http://127.0.0.1  -alljson   -thread  10
		多个url   1个漏洞
			main.exe  -file  urls.txt  -json poc\泛微OA漏洞\泛微-协同办公OA任意管理员登录.json
		多个ulr  多个漏洞
			main.exe  -file  urls.txt  -alljson  -thread 50 （建议不要超过50个线程）
			
		获取url标题
			main.exe  -file  urls.txt -gettitle    （会在url_title.txt中保存结果）
	
	
	
	poc编写说明 ：
		注意：在response_check最少要存在两个，可以一个内容一个返回码，不然无法匹配 ,如无需内容进行匹配，则可以在contain中的value中输入 "renyizifumofamen",使用该字符串恒为true
        在next_check中（该参数用于第二个请求）如没有下一个，next参数要置为空 ，或者写为or也行，但不能为and，否则会匹配失败
		
	优化： 当一个url连续五次超时,可能被waf封禁或本身url不存活，会直接返回结果为false，不存在漏洞

	正则匹配的结果，使用 {{replace{search}replace}}字符串进行替换
	正则匹配方式： 如匹配body，则使用 body:表达式，否则输入 要匹配的头字段，比如Set-Cookie:表达式  ,目前只匹配 set-cookie  Content-type  和body