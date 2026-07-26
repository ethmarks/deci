package main

func (m model) handleCursorMove(key string) model {
	switch key {
	case "up":
		if m.cursorY > 0 {
			m.cursorY -= 1
			m.cursorX = m.getClampedCursorX()
		} else {
			m.cursorX = 0
		}
	case "down":
		if m.cursorY < len(m.lines)-1 {
			m.cursorY += 1
			m.cursorX = m.getClampedCursorX()
		} else {
			m.cursorX = len(m.lines[m.cursorY])
		}
	case "left":
		if m.cursorX > 0 {
			m.cursorX -= 1
			m.cursorPrefX = m.cursorX
		} else if m.cursorY > 0 {
			m.cursorY -= 1
			m.cursorPrefX = len(m.lines[m.cursorY])
			m.cursorX = m.cursorPrefX
		}
	case "right":
		if m.cursorX < len(m.lines[m.cursorY]) {
			m.cursorX += 1
			m.cursorPrefX = m.cursorX
		} else if m.cursorY < len(m.lines)-1 {
			m.cursorY += 1
			m.cursorX = 0
			m.cursorPrefX = m.cursorX
		}
	}
	return m
}

func (m model) updateOffsets() model {
	linesToDisplay := m.termHeight - m.reservedFromTop - m.reservedFromBottom
	colsToDisplay := m.termWidth - m.reservedFromLeft - m.reservedFromRight

	if linesToDisplay < 1 || colsToDisplay < 1 {
		return m
	}

	if m.cursorY < m.paneOffsetY {
		m.paneOffsetY = m.cursorY
	} else if m.cursorY >= m.paneOffsetY+linesToDisplay {
		m.paneOffsetY = m.cursorY - linesToDisplay + 1
	}

	if m.cursorX < m.paneOffsetX {
		m.paneOffsetX = m.cursorX
	} else if m.cursorX >= m.paneOffsetX+colsToDisplay {
		m.paneOffsetX = m.cursorX - colsToDisplay + 1
	}

	return m
}

func (m model) getClampedCursorX() int {
	cursorLine := m.lines[m.cursorY]
	clampedX := min(m.cursorPrefX, len(cursorLine))
	return clampedX
}
