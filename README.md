# URL_shortener
URL shortener (“сокращатель ссылок”) - это сервис, позволяющий пользователю генерировать  для произвольного URL’a его короткую версию, которую удобно вставлять в различные публикации, сообщения, новости, промо-материалы и так далее. Также сервис позволяет получать статистику переходов по каждому сгенерированному URL’у. 

Pre-commit framework используется для работы с git hooks (https://pre-commit.com/). Конфигурация в .pre-commit-config.yaml использует репозиторий с хуками для Go https://github.com/dnephin/pre-commit-golang

Для Github Actions файлы конфигурации размещены в .github/workflows  (lint.yml, pre-commit.yml)