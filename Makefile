STD_OUTPUTS = \
	book.asciidoc \
	book.db \
	book.docx \
	book.epub \
	book.html \
	book.json \
	book.opml \
	book.org \
	book.rst \
	book.rtf \
	book.tex \
	book.text \
	book.textile \
	book.icml


CUSTOM_OUTPUTS =


DISABLED_OUTPUTS = \
	book.odt \
	book.pdf


all: $(STD_OUTPUTS) $(CUSTOM_OUTPUTS)


clean:
	rm -f $(STD_OUTPUTS) $(CUSTOM_OUTPUTS) $(DISABLED_OUTPUTS)


book.md: cdvelh.go
	go run ./cmd/calculatrice-dont-vous-etes-le-heros/main.go > $@


book.odt: book.md
	@rm -f tmp-$@
	pandoc +RTS -Ksize -RTS $< -o tmp-$@
	@mv tmp-$@ $@


$(STD_OUTPUTS): book.md
	@rm -f tmp-$@
	pandoc --epub-cover-image=cover2.jpg --epub-metadata=metadata.xml $< -o tmp-$@
	@mv tmp-$@ $@
