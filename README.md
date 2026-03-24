# 📊 Log Analyzer CLI (Go)

Uma ferramenta de linha de comando (CLI) escrita em Go para análise de logs estruturados em JSON.

## 🚀 Funcionalidades

- Leitura de arquivos de log
- Parsing de logs em formato JSON
- Filtro por nível (`ERROR`, `INFO`, etc.)
- Contagem de eventos por nível
- Exibição de resumo ao final

---

## 📦 Estrutura do Log

Cada linha do arquivo deve estar em formato JSON:

```json
{"time":"2026-03-23 10:00:00","level":"INFO","event":"servidor iniciado"}
{"time":"2026-03-23 10:01:00","level":"ERROR","event":"falha no banco"}
{"time":"2026-03-23 10:02:00","level":"ERROR","event":"timeout"}
