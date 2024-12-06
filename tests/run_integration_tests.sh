# Тест /add_article
echo "Тестирование /add_article"
response_add=$(curl -s -X POST http://0.0.0.0:8000/add_article/ -H "Content-Type: application/json" -d '{"id": 1, "key_words": ["fastapi", "python", "api"]}')
status_add=$?
if [ $status_add -ne 0 ]; then
  echo "Ошибка при вызове /add_article: $response_add"
  exit 1
else
  echo "Успешный вызов /add_article. Ответ: $response_add "
  echo " "
fi

# Тест /search_by_keyword
echo "Тестирование /search_by_keyword"
response_search=$(curl -s -X GET "http://0.0.0.0:8000/search_by_keyword/?keyword=fastapi")
status_search=$?
if [ $status_search -ne 0 ]; then
  echo "Ошибка при вызове /search_by_keyword: $response_search"
  exit 1
else
  echo "Успешный вызов /search_by_keyword. Ответ: $response_search"
  grep -q '1' <<< "$response_search" 
  if [ $? -eq 0 ]; then
    echo "ids=1 найден в ответе /search_by_keyword"
    echo " "
  else
    echo "Предупреждение: ids=1 не найден в ответе /search_by_keyword. Проверьте данные."
    echo " "
  fi
fi


#Suggest
echo "Тестирование suggest 1/3"
RESPONSE=$(curl -s -X POST -H "accept: application/json" -H "Content-Type: application/json" -d '{ "user_tags": ["string"], "articles": [{"id": "1", "tags": ["string"]}] }' http://localhost:8001/suggest/1)
STATUS=$?

if [ $STATUS -ne 0 ]; then
  echo "Ошибка при вызове /suggest: $RESPONSE"
  exit 1
fi

EXPECTED_OUTPUT='{"ids":["1"]}'

if [[ "$RESPONSE" == "$EXPECTED_OUTPUT" ]]; then
  echo "Тест /suggest пройден успешно"
  echo " "
else
  echo "Тест /suggest провален. Полученный результат: $RESPONSE"
  echo " "
fi


echo "Тестирование suggest 2/3"
RESPONSE=$(curl -s -X POST -H "accept: application/json" -H "Content-Type: application/json" \
  -d '{
    "user_tags": ["осень", "небо"],
    "articles": [
      {
        "id": "1",
        "tags": ["осень", "небо"]
      },
      {
        "id": "2",
        "tags": ["весна", "земля"]
      }
    ]
  }' http://localhost:8001/suggest/1)

EXPECTED_OUTPUT='{"ids":["1","2"]}'
if [[ "$RESPONSE" == "$EXPECTED_OUTPUT" ]]; then
  echo "Тест /suggest: Успешный ответ пройден"
  echo " "
else
  echo "Тест /suggest: Успешный ответ провален. Получено: $RESPONSE"
  echo " "
fi

echo "Тестирование suggest 3/3"
RESPONSE=$(curl -s -X POST -H "accept: application/json" -H "Content-Type: application/json" \
  -d '{
    "user_tags": ["осень", "небо"],
    "articles": [
      {
        "id": "1",
        "tags": ["осень", "небо"]
      },
      {
        "id": "2",
        "tags": ["весна", "земля"]
      },
      {
        "id": "3",
        "tags": ["осень", "земля"]
      },
      {
        "id": "4",
        "tags": ["земля", "небо"]
      }
    ]
  }' http://localhost:8001/suggest/1)

EXPECTED_OUTPUT='{"ids":["1","3","4","2"]}'
if [[ "$RESPONSE" == "$EXPECTED_OUTPUT" ]]; then
  echo "Тест /suggest: Успешный тест сортировки"
  echo " "
else
  echo "Тест /suggest: тест сортировки провален. Получено: $RESPONSE"
  echo " "
fi


#Text Process
echo "Тестирование Text Process 1/3"
RESPONSE=$(curl -s -X POST -H "accept: application/json" -H "Content-Type: application/json" -d '{ "text": "что такое осень это небо" }' http://localhost:8001/text/process)
STATUS=$?

if [ $STATUS -ne 0 ]; then
  echo "Ошибка при вызове text process: $RESPONSE"
  exit 1
fi

EXPECTED_OUTPUT='{"tags":["небо","осень"]}'

if [[ "$RESPONSE" == "$EXPECTED_OUTPUT" ]]; then
  echo "Тест /text/process пройден успешно"
  echo " " 
else
  echo "Тест /text/process провален. Полученный результат: $RESPONSE"
  echo " "
fi

echo "Тестирование Text Process 2/3"
RESPONSE=$(curl -s -X POST -H "accept: application/json" -H "Content-Type: application/json" \
  -d '{"text": ""}' http://localhost:8001/text/process)

EXPECTED_OUTPUT='{"tags":[]}'
if [[ "$RESPONSE" == "$EXPECTED_OUTPUT" ]]; then
  echo "Тест /text/process: Пустой ввод пройден"
  echo " "
else
  echo "Тест /text/process: Пустой ввод провален. Получено: $RESPONSE"
  echo " "
fi

echo "Тестирование Text Process 3/3"
RESPONSE=$(curl -s -X POST -H "accept: application/json" -H "Content-Type: application/json" \
  -d '{"text": "9018038"}' http://localhost:8001/text/process)

EXPECTED_OUTPUT='{"tags":[]}'
if [[ "$RESPONSE" == "$EXPECTED_OUTPUT" ]]; then
  echo "Тест /text/process: Текст без тегов пройден"
  echo " "
else
  echo "Тест /text/process: Текст без тегов провален. Получено: $RESPONSE"
  echo " "
fi


echo "Тесты завершены."
exit 0