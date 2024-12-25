
import os
import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

def parse_filename(filename):
    """Извлекает параметры A, B и C из имени файла."""
    base_name = os.path.splitext(filename)[0]
    return tuple(map(int, base_name.split('_')))

def read_file(file_path):
    """Читает данные из файла и возвращает их в виде кортежа."""
    with open(file_path, 'r') as file:
        data = [float(line.strip()) for line in file.readlines()]
    return tuple(data)

def main(folder_path):
    data_dict = {}

    # Считываем данные из всех файлов в указанной папке
    for filename in os.listdir(folder_path):
        if filename.endswith('.txt'):
            A, B, C = parse_filename(filename)
            requests, accepted, rejected, avg_time = read_file(os.path.join(folder_path, filename))

            # Рассчитываем проценты и сохраняем данные
            percent_accepted = (accepted / requests) * 100 if requests > 0 else 0
            percent_rejected = (rejected / requests) * 100 if requests > 0 else 0
            
            data_dict[(A, B, C)] = (percent_accepted, percent_rejected, avg_time)

    # Преобразуем данные в массивы для построения графиков
    A_vals, B_vals, C_vals = zip(*data_dict.keys())
    percent_accepted_vals, percent_rejected_vals, avg_time_vals = zip(*data_dict.values())

    # Создаем графики
    fig = plt.figure(figsize=(24, 10))

    # График 1: Процент принятых заявок
    ax1 = fig.add_subplot(231, projection='3d')
    ax1.scatter(A_vals, B_vals, percent_accepted_vals)
    ax1.set_title('Процент принятых заявок')
    ax1.set_xlabel('A (Буфер)')
    ax1.set_ylabel('B (Учителя)')
    ax1.set_zlabel('Процент принятых (%)')

    # График 2: Процент отвергнутых заявок
    ax2 = fig.add_subplot(232, projection='3d')
    ax2.scatter(A_vals, B_vals, percent_rejected_vals)
    ax2.set_title('Загрузка в процентах')
    ax2.set_xlabel('A (Буфер)')
    ax2.set_ylabel('B (Учителя)')
    ax2.set_zlabel('Загрузка (%)')

    # График 3: Среднее время обработки
    ax3 = fig.add_subplot(233, projection='3d')
    ax3.scatter(A_vals, B_vals, avg_time_vals)
    ax3.set_title('Среднее время обработки')
    ax3.set_xlabel('A (Буфер)')
    ax3.set_ylabel('B (Учителя)')
    ax3.set_zlabel('Среднее время (с)')

    # График 4: Процент принятых заявок от учителей и максимальной загрузки
    ax4 = fig.add_subplot(234, projection='3d')
    ax4.scatter(B_vals, C_vals, percent_accepted_vals)
    ax4.set_title('Процент принятых заявок\nот Учителей и Макс. нагрузки')
    ax4.set_xlabel('B (Учителя)')
    ax4.set_ylabel('C (Макс. нагрузка)')
    ax4.set_zlabel('Процент принятых (%)')

    # График 5: Процент отвергнутых заявок от учителей и максимальной загрузки
    ax5 = fig.add_subplot(235, projection='3d')
    ax5.scatter(B_vals, C_vals, percent_rejected_vals)
    ax5.set_title('Загрузка в процентах\nот Учителей и Макс. нагрузки')
    ax5.set_xlabel('B (Учителя)')
    ax5.set_ylabel('C (Макс. нагрузка)')
    ax5.set_zlabel('Загрузка (%)')

    # График 6: Среднее время обработки от учителей и максимальной загрузки
    ax6 = fig.add_subplot(236, projection='3d')
    ax6.scatter(B_vals, C_vals, avg_time_vals)
    ax6.set_title('Среднее время обработки\nот Учителей и Макс. нагрузки')
    ax6.set_xlabel('B (Учителя)')
    ax6.set_ylabel('C (Макс. нагрузка)')
    ax6.set_zlabel('Среднее время (с)')

    # Сохраняем график
    plt.savefig(os.path.join(folder_path, 'bench.png'))

    plt.tight_layout()
    
    # Закрываем фигуру после сохранения
    plt.close(fig)

if __name__ == '__main__':
    folder_path = input("Введите путь к папке с файлами: ")
    main(folder_path)

