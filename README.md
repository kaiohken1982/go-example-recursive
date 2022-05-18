## HTML Parser

Estrae i link da un documento html

# Argomenti trattati

- Parser html e dom
- Funzioni ricorsive

# Utilizzo

Compilare

```
go build
```

Utilizzo

```
curl -L https://golang.org | ./htmlparser
```

```
echo '<a href="ciao"></a>' | ./htmlparser
```

```
echo '<a href="ciao"></a><a href="test"></a>' | ./htmlparser
```

```
echo '
    <html>
        <body>
            <a href="ciao"></a>
            <a href="test2">
                <div>
                    <a href="inside"></a>
                </div>
            </a>
        </body>
    </html>' | ./htmlparser
```

Note

Il risultato finale NON rispecchia la struttura html passata. Non funziona. 