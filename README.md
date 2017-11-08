# gluare


## example
```go
package main

import (
	"fmt"

	"github.com/Greyh4t/gluare"
	"github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	L.PreloadModule("re", gluare.Loader)
	err := L.DoString(`
		local re=require("re")
		local re1 = re.match("\\d{3}", "aaa000bbb")
		print(re1)
		p, err = re.compile("bbb(\\d{3})")
		if err~=nil then
			print(err)
			return
		end
		local x = "aaa000bbb111aaa222bbb333ccc444"
		local re2 = p:find(x)
		print(re2)
		local re3 = p:findindex(x)
		print(re3[1],re3[2])
		local re4 = p:findsub(x)
		print(re4[1])
		local re5 = p:findall(x)
		print(re5[1],re5[2])
		local re6 = p:findall(x, 1)
		print(re6[1])
		local re7 = p:findallindex(x, 1)
		print(re7[1][1], re7[1][2])
		local re8 = p:findallindex(x)
		print(re8[1][1], re8[1][2])
		print(re8[2][1], re8[1][2])
		local re9 = p:findallsub(x)
		print(re9[1][1], re9[2][1])
		local re10 = p:split(x)
		print(re10[1], re10[2], re10[3])
		local re11 = p:split(x, 2)
		print(re11[1], re11[2])
		local re12 = p:replace(x, "---")
		print(re12)`,
	)
	if err != nil {
		fmt.Println(err)
	}
}
```