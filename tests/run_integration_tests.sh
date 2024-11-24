# Тест /add_article
echo "Тестирование /add_article"
response_add=$(curl -s -X POST http://0.0.0.0:8000/add_article/ -H "Content-Type: application/json" -d '{"id": 1, "key_words": ["fastapi", "python", "api"]}')
status_add=$?
if [ $status_add -ne 0 ]; then
  echo "Ошибка при вызове /add_article: $response_add"
  exit 1
else
  echo "Успешный вызов /add_article. Ответ: $response_add"
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
  else
    echo "Предупреждение: ids=1 не найден в ответе /search_by_keyword. Проверьте данные."
  fi
fi


echo "Тесты завершены."
exit 0