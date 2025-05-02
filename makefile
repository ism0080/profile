.PHONY: dev
dev:
		wgo -file .go -file .templ -xfile _templ.go templ generate :: bunx @tailwindcss/cli -i ./index.css -o public/styles.min.css :: go run main.go