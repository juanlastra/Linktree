.PHONY: clean site

clean:
    @echo "Cleaning site..."
    rm -rf Resultados/salida.html
    rm -rf assets

site: clean
    @echo "Building site..."
    go run main.go