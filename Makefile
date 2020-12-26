auth:
	cp ./stubs/authentication/controller.stub ./app/controllers/auth.go
	cp ./stubs/authentication/middleware.stub ./app/middlewares/auth.go
	cp ./stubs/authentication/routes.stub ./routes/auth.go
	cd routes && sed -i '' 's/\/\/Auth/Auth/g' setup.go

controller:
	cp ./stubs/controller.stub ./app/controllers/$(name).go
	cd app/controllers && sed -i '' 's/STUB/$(name)/g' $(name).go

model:
	cp ./stubs/model.stub ./app/models/$(name).go
	cd app/models && sed -i '' 's/STUB/$(name)/g' $(name).go