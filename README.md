# go-shentong

基于神通数据库官方驱动的Demo, 仅运行于linux

docker run -it --rm -v $PWD:/app -w /app \
-e PKG_CONFIG_PATH=/app/shentong \
-e LD_LIBRARY_PATH=/app/shentong/lib/linux64/ \
golang:latest go run -tags=shentong main.go 


## References

https://github.com/Mystery00/go-shentong

https://github.com/Mystery00/gorm-shentong

https://github.com/team-ide/go-driver

https://gitee.com/teamide/teamide