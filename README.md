# Классификатор спама в SMS на основе TF-IDF и Logistic Regression
Этот проект представляет собой простую модель машинного обучения для классификации SMS-сообщений на "спам" и "не спам" (ham). Модель использует векторизацию текста с помощью TF-IDF и логистическую регрессию из библиотеки Scikit-learn. Точность на тестовом наборе данных составляет ~97%.
## Установка
1.  Клонируйте репозиторий:
    ```bash
    git clone https://github.com/your-username/spam-classifier.git
    cd spam-classifier
    ```
2.  Создайте и активируйте виртуальное окружение:
    ```bash
    python -m venv venv
    source venv/bin/activate
    ```
3.  Установите необходимые библиотеки:
    ```bash
    pip install -r requirements.txt
    ```
## Использование
### Обучение модели
Для обучения модели на данных из `data/spam.csv` и сохранения ее в файл `models/model.joblib` запустите:
    ```bash
    python train.py
    ```
### Предсказание для нового сообщения
Чтобы классифицировать новое сообщение, используйте скрипт `predict.py`:
    ```bash
    python predict.py --message "Congratulations! You've won a $1000 gift card. Click here to claim."
    ```
**Ожидаемый вывод:**
    ```
    Сообщение: "Congratulations! You've won a $1000 gift card. Click here to claim."
    Вердикт:  spam
    ```
## Структура проекта
-   `/data/spam.csv`: Исходный датасет.
-   `/models/`: Папка для сохранения обученной модели.
-   `src/`: Вспомогательные функции для обработки текста.
-   `train.py`: Скрипт для обучения и оценки модели.
-   `predict.py`: Скрипт для предсказания на новых данных.
-   `requirements.txt`: Список зависимостей.
## Данные
Используется публичный датасет [SMS Spam Collection](https://www.kaggle.com/datasets/uciml/sms-spam-collection-dataset) с Kaggle. Он содержит 5,574 SMS-сообщения, помеченных как "spam" или "ham".




### Регистрация и аутентификация

| Метод | Эндпоинт | Описание | Тело запроса |
|-------|----------|----------|---------------|
| POST | `/user/create` | Регистрация нового пользователя | `{"login": "user", "password": "pass"}` |
| POST | `/user/login` | Авторизация, возвращает JWT | `{"login": "user", "password": "pass"}` |

**Пример ответа `/user/login`:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}