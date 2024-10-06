# Slide 1: Creating Modular Templates in Go

- Use `{{define}}` to create individual, reusable templates
- Store each template in a separate file for better organization
- Example: header.tmpl

```html
{{define "header"}}
<header>
    <h1>{{.Title}}</h1>
    <nav>
        <a href="/">Home</a>
        <a href="/about">About</a>
    </nav>
</header>
{{end}}
```

- Example: footer.tmpl

```html
{{define "footer"}}
<footer>
    <p>&copy; 2024 Our Company</p>
</footer>
{{end}}
```

# Slide 2: Importing and Using Templates

- Use `template.ParseGlob()` to load all template files
- Use `{{template}}` to import and use defined templates
- Example: main page structure (index.tmpl)

```html
<!DOCTYPE html>
<html>
    <body>
        {{template "header" .}}
        
        <main>
            <h2>Welcome to our website</h2>
            <p>{{.Content}}</p>
        </main>
        
        {{template "footer" .}}
    </body>
</html>
```

- In Go code:

```go
tmpl, err := template.ParseGlob("templates/*.tmpl")
if err != nil {
    log.Fatal(err)
}

data := struct {
    Title   string
    Content string
}{
    Title:   "Home Page",
    Content: "This is our main content.",
}

tmpl.ExecuteTemplate(w, "index.tmpl", data)
```

# Slide 3: Advanced Template Features

- Use `range` for iteration:

```html
<ul>
{{range .Items}}
    <li>{{.}}</li>
{{end}}
```

- Use `if` and `eq` for conditionals:

```html
{{if eq .UserRole "admin"}}
    <a href="/admin">Admin Panel</a>
{{end}}
```

- Create custom functions:

```go
funcMap := template.FuncMap{
    "uppercase": strings.ToUpper,
}
tmpl := template.New("").Funcs(funcMap)
tmpl = template.Must(tmpl.ParseGlob("templates/*.tmpl"))
```

- Use in template:

```html
<p>{{uppercase .Username}}</p>
```