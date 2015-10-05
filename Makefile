all: book.tex book.pdf book.epub

book.md:
	go run ./cmd/calculatrice-dont-vous-etes-le-heros/main.go > $@

book.tex book.pdf book.epub: book.md
	pandoc $< -o $@
