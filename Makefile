build:
	go build .

run:
	go run .

#test:
#	go test ...

#jsbuild:
#	...

jswatch:
	npx babel --watch frontend --out-dir public/js --presets react-app/prod
