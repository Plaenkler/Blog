# Install dependencies for backend and frontend
install_depends:
	cd ./backend && go mod tidy
	cd ./frontend && npm install

# Bundle Alpine.js and Tailwind CSS
bundle_frontend:
	cd ./frontend && npm run prod-js
	cd ./frontend && npm run prod-css

# Remove .gitkeep files from production
remove_gitkeep:
	rm -f ./backend/pkg/server/routes/web/static/gitkeep
	rm -f ./frontend/dist/gitkeep

# Copy frontend files to static folder
copy_frontend:
	cp -r ./frontend/dist ./backend/pkg/server/routes/web/static
	cp -r ./frontend/src/* ./backend/pkg/server/routes/web/static

# Build Golang executable from backend
build_executable:
	cd ./backend && go mod tidy
	cd ./backend && go build -o ../page.exe ./cmd/main.go

# Remove frontend from backend after
clean_project:
	rm -rf ./backend/pkg/server/routes/web/static/*
	rm -rf ./frontend/dist/*

# Recreate .gitkeep files empty folder
recreate_gitkeep:
	touch ./backend/pkg/server/routes/web/static/gitkeep
	touch ./frontend/dist/gitkeep

# Run prod executable after building it
run_prod:
	./page.exe

prod: install_depends bundle_frontend remove_gitkeep copy_frontend build_executable clean_project recreate_gitkeep
prod-run: prod run_prod