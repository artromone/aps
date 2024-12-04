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
# Блок-схема

> https://www.mermaidchart.com/raw/52790184-7643-458d-827a-e09805cc89d9?theme=light&version=v0.1&format=svg

---
# Диаграмма классов

1. `User` – пользователm, который хочет записаться на курс (навыки, желаемый тип курса и время ожидания)
2. `Application` – заявка на курс (ID пользователя, тип курса, уровень навыков и статус)
3. `Buffer` – очередь заявок FIFO. Ограничена максимальным размером. Может удалять старые заявки при переполнении.
4. `Teacher` – преподаватель (пециализация, уровень квалификации и текущая нагрузка)
5. `ApplicationDispatcher` – диспетчер для обработки заявок. Распределяет заявки между преподавателями, управляет буфером.
6. `TeacherSelector` – отвечает за выбор наиболее подходящего преподавателя на основе квалификации и загруженности.

Основной поток работы:

- Пользователь создает заявку
- Диспетчер заявок пытается распределить его к преподавателю, используя TeacherSelector
- Если мест нет, заявка попадает в буфер-очередь

> https://www.mermaidchart.com/raw/3a4af77e-7d47-408d-94c7-dc5b27bc449d?theme=light&version=v0.1&format=svg

---
# Сиквенс-диаграмма

- Пользователь подает заявку
- Диспетчер проверяет буфер
- Если место есть – находит преподавателя
- Если нет места – добавляет в очередь
- При переполнении буфера удаляются старые заявки

> https://www.mermaidchart.com/raw/736a536b-80bf-4b7e-aa07-32b8ecad85ca?theme=light&version=v0.1&format=svg

---
# BPMN диаграмма

> https://www.plantuml.com/plantuml/svg/TPBFxjCm383lUGeVztk5zi0A20a9mUWsSSvIXvcbIKcSeabzU9frqSBAZcxy-VEVRWDBpx4tOZe6z9EunQBQnrDF3Cqm65AIoTbQL3llHkeG3XWO3Z9SQY6c8WK-KXWagtCgLBMea4RyjmPfsBoFMgE7Jz6OTs_0PxImkM9u4m0prBkP-MOw34zwj1rfKNquexh0lKOpIZBoH8RucffPFa-bvglZ_wkQhYi5m0uMf1O0qm5bcWFsQ52nyrD2bW3octR-iGjeSIJya-1appdt3Wq--n-EwHPL2ryk5o3NoewT306xTIc0btTOdglvsRkdQsJq68h4UL-ahv2-WTFGA8K3IwjmIKSJ43RMPMzUUSPRvEZjwyMkXIt-Zcd0p_sjmnujW_Qkdrl8o7BwKhvIftCcqmz9JFOdV8yOKykbRZrTplv2WMkHHQKcmokqnL-eoZzJ17Q3sCv1xCrV

---
