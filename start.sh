#!/bin/bash
echo "🔄 Pornesc PostgreSQL..."
sudo systemctl start postgresql

echo "🔄 Pornesc Backend..."
cd ~/Documents/Clarity-Gym/backend
go run cmd/main.go &

sleep 2

echo "🔄 Pornesc Frontend..."
cd ~/Documents/Clarity-Gym/frontend
npm run dev
