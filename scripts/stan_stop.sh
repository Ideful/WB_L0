#!/bin/bash

# Получаем вывод lsof -i :4223
output=$(lsof -i :4223)

# Выводим результат
echo "lsof output:"
echo "$output"

# Получаем PID из вывода lsof -i :4223
pid=$(echo "$output" | awk 'NR==2 {print $2}')

# Выводим PID
echo "PID: $pid"

if [ -n "$pid" ]; then
    echo "Killing process with PID: $pid"
    # Завершаем процесс
    kill -9 $pid
else
    echo "No process found on port 4223"
fi