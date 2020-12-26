controller:
	cp ./stubs/controller.stub ./app/controllers/$(name).go
	cd app/controllers && sed -i '' 's/STUB/$(name)/g' $(name).go

model:
	cp ./stubs/model.stub ./app/models/$(name).go
	cd app/models && sed -i '' 's/STUB/$(name)/g' $(name).go