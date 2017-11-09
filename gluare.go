package gluare

import (
	"regexp"

	"github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	re := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"match":        match,
		"find":         find,
		"findindex":    findIndex,
		"findsub":      findSubmatch,
		"findall":      findAll,
		"findallindex": findAllIndex,
		"findallsub":   findAllSubmatch,
		"replace":      replace,
		"split":        split,
		"compile":      compile,
	})

	mt := L.NewTypeMetatable("Regexp")
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"match":        pMatch,
		"find":         pFind,
		"findindex":    pFindIndex,
		"findsub":      pFindSubmatch,
		"findall":      pFindAll,
		"findallindex": pFindAllIndex,
		"findallsub":   pFindAllSubmatch,
		"replace":      pReplace,
		"split":        pSplit,
	}))
	L.SetField(re, "Regexp", mt)

	L.Push(re)
	return 1
}

func match(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)
	matched, err := regexp.MatchString(p, s)
	L.Push(lua.LBool(matched))
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 2
	}
	return 1
}

func find(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)
	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(pattern.FindString(s)))
	return 1
}

func findIndex(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)

	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	tmp := pattern.FindStringIndex(s)
	loc := L.NewTable()
	if tmp != nil {
		loc.Append(lua.LNumber(tmp[0]))
		loc.Append(lua.LNumber(tmp[1]))
	}
	L.Push(loc)
	return 1
}

func findSubmatch(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)

	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	tmp := pattern.FindStringSubmatch(s)
	matched := L.NewTable()
	if tmp != nil {
		for i, m := range tmp {
			if i == 0 {
				continue
			}
			matched.Append(lua.LString(m))
		}
	}
	L.Push(matched)
	return 1
}

func findAll(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)
	n := L.ToInt(3)
	if n <= 0 {
		n = -1
	}

	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	tmp := pattern.FindAllString(s, n)
	matched := L.NewTable()
	for _, m := range tmp {
		matched.Append(lua.LString(m))
	}
	L.Push(matched)
	return 1
}

func findAllIndex(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)
	n := L.ToInt(3)
	if n <= 0 {
		n = -1
	}

	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	tmp := pattern.FindAllStringIndex(s, n)
	locs := L.NewTable()
	for _, m := range tmp {
		loc := L.NewTable()
		loc.Append(lua.LNumber(m[0]))
		loc.Append(lua.LNumber(m[1]))
		locs.Append(loc)
	}
	L.Push(locs)
	return 1
}

func findAllSubmatch(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)
	n := L.ToInt(3)
	if n <= 0 {
		n = -1
	}

	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	tmp := pattern.FindAllStringSubmatch(s, n)
	matched := L.NewTable()
	for _, each := range tmp {
		match := L.NewTable()
		for i, m := range each {
			if i == 0 {
				continue
			}
			match.Append(lua.LString(m))
		}
		matched.Append(match)
	}
	L.Push(matched)
	return 1
}

func replace(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)
	newStr := L.CheckString(3)
	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LString(s))
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(pattern.ReplaceAllString(s, newStr)))
	return 1
}

func split(L *lua.LState) int {
	p := L.CheckString(1)
	s := L.CheckString(2)
	n := L.ToInt(3)
	if n <= 0 {
		n = -1
	}

	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	table := L.NewTable()
	for _, item := range pattern.Split(s, n) {
		table.Append(lua.LString(item))
	}
	L.Push(table)
	return 1
}

func compile(L *lua.LState) int {
	p := L.CheckString(1)
	pattern, err := regexp.Compile(p)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	ud := L.NewUserData()
	ud.Value = pattern
	L.SetMetatable(ud, L.GetTypeMetatable("Regexp"))
	L.Push(ud)
	return 1
}

func checkRegexp(L *lua.LState) *regexp.Regexp {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*regexp.Regexp); ok {
		return v
	}
	L.ArgError(1, "re.Regexp expected")
	return nil
}

func pMatch(L *lua.LState) int {
	p := checkRegexp(L)
	L.Push(lua.LBool(p.MatchString(L.CheckString(2))))
	return 1
}

func pFind(L *lua.LState) int {
	p := checkRegexp(L)
	L.Push(lua.LString(p.FindString(L.CheckString(2))))
	return 1
}

func pFindIndex(L *lua.LState) int {
	p := checkRegexp(L)
	tmp := p.FindStringIndex(L.CheckString(2))
	loc := L.NewTable()
	if tmp != nil {
		loc.Append(lua.LNumber(tmp[0]))
		loc.Append(lua.LNumber(tmp[1]))
	}
	L.Push(loc)
	return 1
}

func pFindSubmatch(L *lua.LState) int {
	p := checkRegexp(L)
	tmp := p.FindStringSubmatch(L.CheckString(2))
	matched := L.NewTable()
	if tmp != nil {
		for i, m := range tmp {
			if i == 0 {
				continue
			}
			matched.Append(lua.LString(m))
		}
	}
	L.Push(matched)
	return 1
}

func pFindAll(L *lua.LState) int {
	p := checkRegexp(L)
	s := L.CheckString(2)
	n := L.ToInt(3)
	if n <= 0 {
		n = -1
	}

	tmp := p.FindAllString(s, n)
	matched := L.NewTable()
	for _, m := range tmp {
		matched.Append(lua.LString(m))
	}
	L.Push(matched)
	return 1
}

func pFindAllIndex(L *lua.LState) int {
	p := checkRegexp(L)
	s := L.CheckString(2)
	n := L.ToInt(3)
	if n <= 0 {
		n = -1
	}
	tmp := p.FindAllStringIndex(s, n)
	locs := L.NewTable()
	for _, m := range tmp {
		loc := L.NewTable()
		loc.Append(lua.LNumber(m[0]))
		loc.Append(lua.LNumber(m[1]))
		locs.Append(loc)
	}
	L.Push(locs)
	return 1
}

func pFindAllSubmatch(L *lua.LState) int {
	p := checkRegexp(L)
	s := L.CheckString(2)
	n := L.ToInt(3)
	if n <= 0 {
		n = -1
	}
	tmp := p.FindAllStringSubmatch(s, n)
	matched := L.NewTable()
	for _, each := range tmp {
		match := L.NewTable()
		for i, m := range each {
			if i == 0 {
				continue
			}
			match.Append(lua.LString(m))
		}
		matched.Append(match)
	}
	L.Push(matched)
	return 1
}

func pReplace(L *lua.LState) int {
	p := checkRegexp(L)
	L.Push(lua.LString(p.ReplaceAllString(L.CheckString(2), L.CheckString(3))))
	return 1
}

func pSplit(L *lua.LState) int {
	p := checkRegexp(L)
	s := L.CheckString(2)
	n := L.ToInt(3)
	if n <= 0 {
		n = -1
	}
	table := L.NewTable()
	for _, item := range p.Split(s, n) {
		table.Append(lua.LString(item))
	}
	L.Push(table)
	return 1
}
