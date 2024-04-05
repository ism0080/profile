# ToDo

- Reduce asset size
- Add resume page
- Fix watchdog

## Getting Started

Install Dependencies

- [Watcher Go - Live Reload](https://github.com/bokwoon95/wgo)
- [Templ - HTML templating](https://github.com/a-h/templ)
- [Tailwindcss - CLI](https://tailwindcss.com/blog/standalone-cli)

```shell
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/bokwoon95/wgo@latest
winget install -e --id TailwindLabs.TailwindCSS
```

Live Reload
```
wgo -file .go -file .templ -xfile _templ.go templ generate :: go run main.go
```