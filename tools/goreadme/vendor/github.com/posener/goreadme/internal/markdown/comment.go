// Copyied from https://github.com/golang/go/blob/master/src/go/doc/comment.go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Godoc comment extraction and comment -> HTML formatting.

package markdown

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"text/template" // for HTMLEscape
	"unicode"
	"unicode/utf8"
)

// ToMarkdown converts comment text to formatted Markdown.
// The comment was prepared by DocReader,
// so it is known not to have leading, trailing blank lines
// nor to have trailing spaces at the end of lines.
// The comment markers have already been removed.
//
// Each span of unindented non-blank lines is converted into
// a single paragraph. There is one exception to the rule: a span that
// consists of a single line, is followed by another paragraph span,
// begins with a capital letter, and contains no punctuation
// other than parentheses and commas is formatted as a heading.
//
// A span of indented lines is converted into a <pre> block,
// with the common indent prefix removed.
//
// URLs in the comment text are converted into links; if the URL also appears
// in the words map, the link is taken from the map (if the corresponding map
// value is the empty string, the URL is not converted into a link).
func ToMarkdown(w io.Writer, text string, opts ...Option) {
	var o options
	for _, f := range opts {
		f(&o)
	}

	for _, b := range blocks(text, o.noDiffs) {
		switch b.op {
		case opPara:
			// New paragraph
			for _, line := range b.lines {
				emphasize(w, line, o.words, false)
			}
			fmt.Fprint(w, "\n")
		case opHead:
			// Headline
			fmt.Fprint(w, "## ")
			for _, line := range b.lines {
				fmt.Fprint(w, line)
			}
			fmt.Fprint(w, "\n\n")
		case opPre:
			// Code block
			fmt.Fprintf(w, "```%s\n", b.lang)
			for _, line := range b.lines {
				emphasize(w, line, nil, false)
			}
			fmt.Fprint(w, "```\n\n")
		}
	}
}

// Option is option type for ToMarkdown
type Option func(*options)

// OptWords sets the list of known words.
// Go identifiers that appear in the words map are italicized; if the corresponding
// map value is not the empty string, it is considered a URL and the word is converted
// into a link.
func OptWords(words map[string]string) Option {
	return func(o *options) { o.words = words }
}

// OptNoDiff disables automatic marking of code blocks as diffs.
func OptNoDiff(noDiffs bool) Option {
	return func(o *options) { o.noDiffs = noDiffs }
}

type options struct {
	words   map[string]string
	noDiffs bool
}

const (
	// Regexp for Go identifiers
	identRx = `[\pL_][\pL_0-9]*`

	// Regexp for URLs
	// Match parens, and check in pairedParensPrefixLen for balance - see #5043
	// Match .,:;?! within path, but not at end - see #18139, #16565
	// This excludes some rare yet valid urls ending in common punctuation
	// in order to allow sentences ending in URLs.

	// url title (optional)
	urlTitle = `(\(([^)]+)\)\W)?`
	// protocol (required) e.g. http
	protoPart = `(https?|ftp|file|gopher|mailto|nntp)`
	// host (required) e.g. www.example.com or [::1]:8080
	hostPart = `([a-zA-Z0-9_@\-.\[\]:]+)`
	// path+query+fragment (optional) e.g. /path/index.html?q=foo#bar
	pathPart = `([.,:;?!]*[a-zA-Z0-9$'()*+&#=@~_/\-\[\]%])*`

	urlRx = protoPart + `://` + hostPart + pathPart

	// Regexp for local paths
	localRx = `\.\/[a-zA-Z0-9_@\-\.\/]*`
)

var matchRx = regexp.MustCompile(`(` + urlTitle + `((` + urlRx + `)|(` + localRx + `)))|(` + identRx + `)`)

// pairedParensPrefixLen returns the length of the longest prefix of s containing paired parentheses.
func pairedParensPrefixLen(s string) int {
	parens := 0
	l := len(s)
	for i, ch := range s {
		switch ch {
		case '(':
			if parens == 0 {
				l = i
			}
			parens++
		case ')':
			parens--
			if parens == 0 {
				l = len(s)
			} else if parens < 0 {
				return i
			}
		}
	}
	return l
}

// Emphasize and escape a line of text for HTML. URLs are converted into links;
// if the URL also appears in the words map, the link is taken from the map (if
// the corresponding map value is the empty string, the URL is not converted
// into a link). Go identifiers that appear in the words map are italicized; if
// the corresponding map value is not the empty string, it is considered a URL
// and the word is converted into a link.
func emphasize(w io.Writer, line string, words map[string]string, nice bool) {
	if line[len(line)-1] != '\n' {
		line = line + "\n"
	}
	for {
		m := matchRx.FindStringSubmatchIndex(line)
		if m == nil {
			break
		}
		// m >= 6 (two parenthesized sub-regexps in matchRx, 1st one is urlRx)

		// write text before match
		fmt.Fprint(w, line[0:m[0]])
		// adjust match if necessary
		match := line[m[0]:m[1]]
		if n := pairedParensPrefixLen(match); n < len(match) {
			// match contains unpaired parentheses (rare);
			// redo matching with shortened line for correct indices
			m = matchRx.FindStringSubmatchIndex(line[:m[0]+n])
			match = match[:n]
		}

		// analyze match
		url := ""
		title := ""
		image := false
		italics := false
		if words != nil {
			url, italics = words[match]
		}

		// If the url ends with a punctuation mark, we will hold it here.
		after := ""

		if m[2] >= 0 {
			// A url match against first parenthesized sub-regexp; must be
			// match against urlRx.
			if !italics {
				// no alternative URL in words list, use match instead
				url = match
				title = match
				if m[4] >= 0 {
					title = line[m[6]:m[7]]
					url = line[m[8]:m[9]]
					if strings.HasPrefix(title, "image/") {
						title = strings.TrimPrefix(title, "image/")
						image = true
					}
				}
			}

			// Skip Go path ellipsis.
			if strings.HasSuffix(url, "/...") {
				fmt.Fprint(w, line[1:m[1]])
				line = line[m[1]:]
				continue
			}

			// Remove punctuation mark from url/title suffix.
			switch url[len(url)-1] {
			case '.', ',', ':', ';', '?', '!':
				after = string(url[len(url)-1])
				if title[len(title)-1] == url[len(url)-1] {
					title = title[:len(title)-1]
				}
				url = url[:len(url)-1]
			}

			italics = false // don't italicize URLs
		}

		// write match
		if image {
			fmt.Fprint(w, "!")
		}
		if len(url) > 0 {
			fmt.Fprint(w, "[")
			template.HTMLEscape(w, []byte(title))
			fmt.Fprint(w, "](")
		}
		if italics {
			fmt.Fprint(w, "*")
		}
		if len(url) > 0 {
			fmt.Fprint(w, url)
		} else {
			fmt.Fprint(w, match)
		}
		if italics {
			fmt.Fprint(w, "*")
		}
		if len(url) > 0 {
			fmt.Fprint(w, ")")
		}
		fmt.Fprint(w, after)

		// advance
		line = line[m[1]:]
	}
	fmt.Fprint(w, line)
}

func indentLen(s string) int {
	i := 0
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	return i
}

func isBlank(s string) bool {
	return len(s) == 0 || (len(s) == 1 && s[0] == '\n')
}

func commonPrefix(a, b string) string {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return a[0:i]
}

func unindent(block []string) {
	if len(block) == 0 {
		return
	}

	// compute maximum common white prefix
	prefix := block[0][0:indentLen(block[0])]
	for _, line := range block {
		if !isBlank(line) {
			prefix = commonPrefix(prefix, line[0:indentLen(line)])
		}
	}
	n := len(prefix)

	// remove
	for i, line := range block {
		if !isBlank(line) {
			block[i] = line[n:]
		}
	}
}

// heading returns the trimmed line if it passes as a section heading;
// otherwise it returns the empty string.
func heading(line string) string {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return ""
	}

	// a heading must start with an uppercase letter
	r, _ := utf8.DecodeRuneInString(line)
	if !unicode.IsLetter(r) || !unicode.IsUpper(r) {
		return ""
	}

	// it must end in a letter or digit:
	r, _ = utf8.DecodeLastRuneInString(line)
	if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
		return ""
	}

	// exclude lines with illegal characters. we allow "(),"
	if strings.ContainsAny(line, ";:!?+*/=[]{}_^°&§~%#@<\">\\") {
		return ""
	}

	// allow "'" for possessive "'s" only
	for b := line; ; {
		i := strings.IndexRune(b, '\'')
		if i < 0 {
			break
		}
		if i+1 >= len(b) || b[i+1] != 's' || (i+2 < len(b) && b[i+2] != ' ') {
			return "" // not followed by "s "
		}
		b = b[i+2:]
	}

	// allow "." when followed by non-space
	for b := line; ; {
		i := strings.IndexRune(b, '.')
		if i < 0 {
			break
		}
		if i+1 >= len(b) || b[i+1] == ' ' {
			return "" // not followed by non-space
		}
		b = b[i+1:]
	}

	return line
}

type op int

const (
	opPara op = iota
	opHead
	opPre
)

type block struct {
	op    op
	lines []string

	lang string // for opPre, the language of the code block.
}

func blocks(text string, skipDiffs bool) []block {
	var (
		out  []block
		para []string

		lastWasBlank   = false
		lastWasHeading = false
	)

	close := func() {
		if para != nil {
			out = append(out, block{op: opPara, lines: para})
			para = nil
		}
	}

	lines := strings.SplitAfter(text, "\n")
	unindent(lines)
	for i := 0; i < len(lines); {
		line := lines[i]
		if isBlank(line) {
			// close paragraph
			close()
			i++
			lastWasBlank = true
			continue
		}
		if indentLen(line) > 0 {
			// close paragraph
			close()

			// Used to remember if all lines are valid diff code blocks.
			isValidDiff := true
			// Used to remember if there was at least one '+' or '-' signs.
			anyDiff := false
			diffChIdx := diffCharIdx(line)

			// count indented or blank lines
			j := i + 1
			for j < len(lines) && (isBlank(lines[j]) || indentLen(lines[j]) > 0) {
				if isValidDiffLine(lines[j], diffChIdx) {
					isValidDiff = false
				}
				if isDiffLine(lines[j], diffChIdx) {
					anyDiff = true
				}
				j++
			}
			// but not trailing blank lines
			for j > i && isBlank(lines[j-1]) {
				j--
			}
			pre := lines[i:j]
			i = j

			unindent(pre)

			// put those lines in a pre block
			lang := "go"
			if isValidDiff && anyDiff && !skipDiffs {
				lang = "diff"
			}
			out = append(out, block{op: opPre, lines: pre, lang: lang})
			lastWasHeading = false
			continue
		}

		if lastWasBlank && !lastWasHeading && i+2 < len(lines) &&
			isBlank(lines[i+1]) && !isBlank(lines[i+2]) && indentLen(lines[i+2]) == 0 {
			// current line is non-blank, surrounded by blank lines
			// and the next non-blank line is not indented: this
			// might be a heading.
			if head := heading(line); head != "" {
				close()
				out = append(out, block{op: opHead, lines: []string{head}})
				i += 2
				lastWasHeading = true
				continue
			}
		}

		// open paragraph
		lastWasBlank = false
		lastWasHeading = false
		para = append(para, lines[i])
		i++
	}
	close()

	return out
}

// diffCharIdx returns the index of a diff character, given the first line of a code block.
func diffCharIdx(line string) int {
	i := indentLen(line) - 1
	if len(line) > i+1 && (line[i+1] == '+' || line[i+1] == '-') {
		i++
	}
	return i
}

// isDiffLine returns if this is a valid diff line given a code block line, and the expected index
// for the diff character.
func isValidDiffLine(line string, i int) bool {
	if isBlank(line) {
		return false
	}
	return len(line) <= i || (line[i] != ' ' && line[i] != '+' && line[i] != '-')
}

// isDiffLine returns if the character at i is a '+' or a '-' sign.
func isDiffLine(line string, i int) bool {
	return len(line) > i && (line[i] == '+' || line[i] == '-')
}
