@echo off

cd C:\Users\user\Desktop\Диплом\socnet

start "Go" go run server.go handler_db.go requests_handler.go
start "Vue" npm run dev
