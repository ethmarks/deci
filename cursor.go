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
		if m.cursorY < len(*m.lines)-1 {
			m.cursorY += 1
			m.cursorX = m.getClampedCursorX()
		} else {
			m.cursorX = len((*m.lines)[m.cursorY])
		}
	case "left":
		if m.cursorX > 0 {
			m.cursorX -= 1
			m.cursorPrefX = m.cursorX
		} else if m.cursorY > 0 {
			m.cursorY -= 1
			m.cursorPrefX = len((*m.lines)[m.cursorY])
			m.cursorX = m.cursorPrefX
		}
	case "right":
		if m.cursorX < len((*m.lines)[m.cursorY]) {
			m.cursorX += 1
			m.cursorPrefX = m.cursorX
		} else if m.cursorY < len(*m.lines)-1 {
			m.cursorY += 1
			m.cursorX = 0
			m.cursorPrefX = m.cursorX
		}
	}
	return m
}

func (m model) updateOffsets() model {
	contentRows := m.termHeight - m.reservedFromTop - m.reservedFromBottom
	contentCols := m.termWidth - m.reservedFromLeft - m.reservedFromRight

	if contentRows < 1 || contentCols < 1 {
		return m
	}

	if m.cursorY < m.paneOffsetY {
		m.paneOffsetY = m.cursorY
	} else if m.cursorY >= m.paneOffsetY+contentRows {
		m.paneOffsetY = m.cursorY - contentRows + 1
	}

	if m.cursorX < m.paneOffsetX {
		m.paneOffsetX = m.cursorX
	} else if m.cursorX >= m.paneOffsetX+contentCols {
		m.paneOffsetX = m.cursorX - contentCols + 1
	}

	return m
}

func (m model) getClampedCursorX() int {
	cursorLine := (*m.lines)[m.cursorY]
	clampedX := min(m.cursorPrefX, len(cursorLine))
	return clampedX
}
