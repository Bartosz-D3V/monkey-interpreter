package lexer

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch > 0 && l.ch <= 32 {
		l.readChar()
	}
}

func (l *Lexer) readIdent() string {
	sPos := l.position
	for isChar(l.ch) {
		l.readChar()
	}
	return l.input[sPos:l.position]
}

func (l *Lexer) readLogicOp() string {
	sPos := l.position
	for l.ch == '!' || l.ch == '=' {
		l.readChar()
	}
	return l.input[sPos:l.position]
}

func (l *Lexer) readNum() string {
	sPos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[sPos:l.position]
}

func isChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() string {
	sPost := l.position + 1
	l.readChar()
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[sPost:l.position]
}
