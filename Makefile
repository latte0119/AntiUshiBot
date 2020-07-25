build:
	go build -o aub

deploy:
	zip -r aub.zip * -x envfiles; \
	aws lambda \
	update-function-code \
	--function-name aub \
	--zip-file fileb://aub.zip \
	--publish; \
	rm aub.zip