Daily practice, mainly focusing on LeetCode Hot 150 and TopCoder 200

the best method is changing the type each day.
# go test, go tool 分析代码测试情况，测试用例是否完整
但是我只用79wordsearch，参考一下
go test -coverprofile=coverage.out ./src/ -count=1
go tool cover -html=./coverage.out 
go tool cover -func=coverage.out