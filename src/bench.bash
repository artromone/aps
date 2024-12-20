#!/bin/bash

# Создаем папку out, если она не существует
mkdir -p out

# Максимальное количество параллельных процессов
max_jobs=20 # Установите это значение в зависимости от количества потоков вашего процессора

# Генерация случайных чисел от 1 до 10 для параметров
for i in {1..10}; do
    for j in {1..10}; do
        for k in {1..10}; do
            # Формируем имя файла на основе значений параметров
            filename="out/${i}_${j}_${k}.txt"
            # Запуск программы с параметрами в фоновом режиме
            go run . -buffer "$i" -teachers "$j" -maxload "$k" > "$filename" &

            # Проверка количества фоновых процессов
            while [ $(jobs | wc -l) -ge $max_jobs ]; do
                wait -n  # Ожидание завершения любого из фоновых процессов
            done
        done
    done
done

# Ожидание завершения всех оставшихся процессов
wait