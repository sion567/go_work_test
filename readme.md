1.在各个模块下新建各自的mod

```shell
work_test\mainModule>go mod init mainModule
work_test\module>go mod init module
work_test\module1>go mod init module1
work_test\module2>go mod init module2
```


2.在项目目录新建work

```shell
work_test>go work init ./mainModule ./module ./module1 ./module2
```