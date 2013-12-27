package encode

import "strings"

func coding(r rune) rune {

	var c map[rune]rune
	c = make(map[rune]rune)

	c['a'] = '.'
	c['b'] = ':'
	c['c'] = ','
	c['d'] = ';'
	c['e'] = '"'
	c['f'] = '='
	c['g'] = '≈'
	c['h'] = '\''
	c['i'] = '!'
	c['j'] = '?'
	c['k'] = '-'
	c['l'] = '_'
	c['m'] = '+'
	c['n'] = '÷'
	c['o'] = '•'
	c['p'] = '%'
	c['q'] = '$'
	c['r'] = '/'
	c['s'] = '\\'
	c['t'] = '@'
	c['u'] = '#'
	c['v'] = '('
	c['w'] = ')'
	c['x'] = '&'
	c['y'] = '*'
	c['z'] = '~'
	c[' '] = '<'
	c['.'] = '>'

	return c[r]
}

func Encode(s string) string {
	b := []rune(strings.ToLower(s))

	for i := 0; i < len(b); i++ {
		b[i] = coding(b[i])
	}

	return string(b)
}
