# 简介
zlang是一个基于Go实现的动态语言解释器（aka，z语言），融合了JavaScript和Python的部分语法，
支持传统语言基本流程控制，包括分支判断，标准I/O，函数闭包，递归等功能。
当前正在积极开发中。
# 快速开始
如果你已经安装好了go语言环境，可以选择直接下载编译
```
git clone https://github.com/Chasing1020/zlang.git
go mod tidy
go build
```
或者选择直接使用编译好的版本（win(amd64)：zlang.exe；mac(arm64)：zlang）。尝试运行
```bash
./zlang run main.zjc
# Hello, world!
```
## 1. 数据类型
当前支持了六种基本数据类型：int, string, boolean, array, map, function。
每一次新变量创建，需要使用let，例如
```js
let int = 1;
let string = "a string";
let boolean = true;
let array = [1, 2, 3, 4, 5];
let map = {"k": "v"};
let add = function (a, b) {
    return a + b;
};
```
array 类型和Python的list相同，可以存放任何类型的数据，像这样：
```js
let arr = [1, 2, function (a, b) {println(a + b);}]
arr[2](arr[0], arr[1]) // 3
```
map 类型，为避免哈希冲突key只支持int与string类型，但是在一个map中可以同时混用这两个类型，如：
```js
let map = {"chasing":1020, 1020:"chasing"}
println(map["chasing"], map[1020]) // 1020, chasing
```
## 2. 运算符
当前版本支持：+, -, *, /, %, <, >, <=, >=, !=, ==等基础运算符，优先顺序与C语言相同。
## 3. 流程控制
if和for循环的使用，与C类语言相同
```js
if (true) { print("true"); } else { print("false"); }

let sum = 0;
for (let i = 0; i <= 100; i = i + 1) {
    sum = sum + i;
}
println(sum); // 5050
```
## 4. 内置函数
给定字符串运算出表达式的结果：eval(x) 

标准输出：print(x), println(x), printf(fmt, x) 

标准输入：input()，默认返回值为string类型

取长度：len(x)，返回数组或者字符串长度

新建数组：newArray(x)，新建长度为x的整型数组，默认初始化为0

类型转换：string(x), int(x)，string和int相互转换

比较函数：min(a, b), max(a, b)，返回两个整数的最小值最大值

## 5. 基本命令
直接在命令行输入./zlang运行程序，输入两次ctrl+c或者一次ctrl+D进行退出。
```bash
(base) -> zlang git:(master) ./zlang                         
Welcome to zLang v0.0.1.
Type "help()" for more information.
> ^C
(To exit, press Ctrl+C again or Ctrl+D)
> ^C%                             
```
或者是将已经写好的程序直接通过run命令运行，项目已写好两个测试文件，放在
test_scripts文件夹下：palindrome_number.zjc以及two_sum.zjc

分别为判断回文数、两数之和的算法
```js
./zlang run test_scripts/palindrome_number.zjc
./zlang run test_scripts/two_sum.zjc
```