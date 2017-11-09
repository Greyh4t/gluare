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
		local pattern = "bbb(\\d{3})"
		local str = "aaa000bbb111aaa222bbb333ccc444"

		print("[+] re.match")
		local r, err = re.match(pattern, str)
		print(r, err)

		print("[+] re.find")
		local r, err = re.find(pattern, str)
		print(r, err)

		print("[+] re.findindex")
		local r, err = re.findindex(pattern, str)
		print(r[1], r[2], err)

		print("[+] re.findsub")
		local r, err = re.findsub(pattern, str)
		print(r[1], err)

		print("[+] re.findall")
		local r, err = re.findall(pattern, str)
		print(r[1], err)

		print("[+] re.findallindex")
		local r, err = re.findallindex(pattern, str)
		print(r[1][1], r[1][2], r[2][1], r[1][2], err)

		print("[+] re.findallsub")
		local r, err = re.findallsub(pattern, str)
		print(r[1][1], r[2][1], err)

		print("[+] re.split")
		local r, err = re.split(pattern, str)
		print(r[1], r[2], err)

		print("[+] re.split")
		local r, err = re.split(pattern, str, 2)
		print(r[1], r[2], err)

		print("[+] re.replace")
		local r, err = re.replace(pattern, str, "---")
		print(r, err)

		print("[+] re.compile")
		p, err = re.compile(pattern)
		if err~=nil then
			print(err)
			return
		end

		print("[+] p:find")
		local r = p:find(str)
		print(r)

		print("[+] p:findindex")
		local r = p:findindex(str)
		print(r[1],r[2])

		print("[+] p:findsub")
		local r = p:findsub(str)
		print(r[1])

		print("[+] p:findall")
		local r = p:findall(str)
		print(r[1],r[2])

		print("[+] p:findall")
		local r = p:findall(str, 1)
		print(r[1])

		print("[+] p:findallindex")
		local r = p:findallindex(str, 1)
		print(r[1][1], r[1][2])

		print("[+] p:findallindex")
		local r = p:findallindex(str)
		print(r[1][1], r[1][2], r[2][1], r[1][2])

		print("[+] p:findallsub")
		local r = p:findallsub(str)
		print(r[1][1], r[2][1])

		print("[+] p:split")
		local r = p:split(str)
		print(r[1], r[2], r[3])

		print("[+] p:split")
		local r = p:split(str, 2)
		print(r[1], r[2])

		print("[+] p:replace")
		local r = p:replace(str, "---")
		print(r)`,
	)
	if err != nil {
		fmt.Println(err)
	}
}
```