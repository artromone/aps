Дистанционный сервис для курсов по программированию обрабатывает заявки на курсы по различным направлениям, с учетом времени, которое пользователи готовы ждать, чтобы попасть на курс. Исходя из этого каждый клиент в зависимости от его уровня подготовки и требуемых компетенций попадает к конкретному преподавателю.

---
# Вариант 9

9: ИБ | ИЗ1 | ПЗ2 | Д10З2 | Д10О5 | Д2П1 | Д2Б5 | ОР1 | ОД3

1.1. Источники
- ИБ — бесконечный
- ИЗ1 — пуассоновский для бесконечных, экспоненциальная задержка для конечных

1.2. Приборы
- ПЗ2 — равномерный

2.1. Дисциплины буферизации
- Д1ОЗ2 — в порядке поступления

2.2. Дисциплины отказа
- Д1ОО5 — вновь пришедшая

2.3. Дисциплины постановки на обслуживание
- Д2П1 — приоритет по номеру прибора
- Д2Б5 — приоритет по номеру источника, заявки в пакете

3.1. Динамическое отражение результатов (пошаговый режим)
- ОД3 — временные диаграммы, текущее состояние

3.2. Отражение результатов после сбора статистики ОР1-ОР2 (автоматический режим)
- ОР1 — сводная таблица результатов

---
# Формализованная схема и описание СМО

1) Источник - пользователи, желающие записаться на курсы программирования. Могут быть как новички, так и опытные специалисты, которые хотят улучшить свои навыки.

2) Заявка - запись на курс. Люди подают заявки на определённый курс (например, Python, Java, веб-разработка), указывая свои навыки, желаемый уровень сложности у курса и время обучения.

3) Буфер - очередь заявок. Если все места на курсах заняты, новые заявки помещаются в очередь ожидания. Эта очередь организована по принципу FIFO (первый пришёл — первый обслужен). Если буфер переполнен, отказываются от самых старых заявок.

4) Прибор - преподаватели (специалисты, которые ведут курсы). Каждый преподаватель имеет определённый уровень квалификации и специализацию (веб-разработка, мобильные приложения).

5) Диспетчер постановки - система управления заявками. Обрабатывает поступающие заявки и распределяет их между доступными курсами. Если все места заняты, заявки помещаются в буфер.

6) Диспетчер выборки - система выбора преподавателя. Определяет, какой преподаватель должен взять новую заявку из буфера или освободившееся место на курсе. Выбор осуществляется по приоритету квалификации и доступности (нагрузке) преподавателя.

---

# Диаграмма классов

---

# Сиквенс-диаграмма

---

# BPMN диаграмма


---
---
## Особенности реализации системы

- **Микросервисы**: Разработка микросервисов для управления заявками, курсами и преподавателями.
- **Система рейтингов**: Студенты могут оставлять отзывы о курсах и преподавателях, что поможет улучшить качество обучения.
- **Гибкость расписания**: Возможность выбора времени занятий по взаимной договоренности между студентами и преподавателями.
- **Поддержка различных форматов обучения**: Включение видеоуроков, текстовых материалов и живых сессий.

Также пользователи имеют возможность выбирать, насколько долго они готовы ждать место на курсе, что позволяет им лучше планировать своё обучение и повышает удовлетворённость от процесса записи.
